from typing import Any
from typing import List

import brownie
from brownie.network.account import Account
from brownie.network.state import Chain
from eth_utils import encode_hex
from eth_utils import to_canonical_address

from tests.contract_helpers import compute_batch_hash
from tests.contract_helpers import mine_until
from tests.contract_helpers import schedule_config
from tests.contract_helpers import ZERO_HASH32
from tests.factories import make_batch
from tests.factories import make_batch_config
from tests.factories import make_bytes


def test_check_and_increment_half_steps(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=10,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span * 10, chain)

    for i in range(3):
        assert executor_contract.numExecutionHalfSteps() == 2 * i

        with brownie.reverts():
            executor_contract.executePlainBatch([])
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})

        assert executor_contract.numExecutionHalfSteps() == 2 * i + 1

        with brownie.reverts():
            executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})
        executor_contract.executePlainBatch([])

    assert executor_contract.numExecutionHalfSteps() == 2 * (i + 1)


def test_check_batching_is_active(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        keypers=keypers,
        threshold=0,
        batch_span=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    with brownie.reverts():
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0)


def test_check_batching_period_is_over(
    executor_contract: Any,
    config_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=10,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)

    for batch_index in range(3):
        mine_until(config.start_block_number + (batch_index + 1) * config.batch_span - 2, chain)
        with brownie.reverts():
            executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})

        plain_batch = make_batch()
        mock_batcher_contract.setBatchHash(batch_index, 1, compute_batch_hash(plain_batch))
        executor_contract.executePlainBatch(plain_batch)


def test_call_target_function(
    executor_contract: Any,
    config_contract: Any,
    mock_target_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
    mock_target_function_selector: bytes,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        keypers=keypers,
        threshold=0,
        target_address=to_canonical_address(mock_target_contract.address),
        target_function_selector=mock_target_function_selector,
        transaction_gas_limit=5000,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    batch = make_batch(3)
    tx = executor_contract.executeCipherBatch(ZERO_HASH32, batch, 0, {"from": keypers[0]})
    assert len(tx.events["Called"]) == len(batch)
    for i, transaction in enumerate(batch):
        assert tx.events["Called"][i]["transaction"] == encode_hex(transaction)
        assert (
            config.transaction_gas_limit - 500
            < tx.events["Called"][0]["gas"]
            <= config.transaction_gas_limit
        )

    batch = make_batch(3)
    mock_batcher_contract.setBatchHash(0, 1, compute_batch_hash(batch))
    tx = executor_contract.executePlainBatch(batch)
    assert len(tx.events["Called"]) == len(batch)
    for i, transaction in enumerate(batch):
        assert tx.events["Called"][i]["transaction"] == encode_hex(transaction)
        assert (
            config.transaction_gas_limit - 500
            < tx.events["Called"][0]["gas"]
            <= config.transaction_gas_limit
        )


def test_emit_event(
    executor_contract: Any,
    config_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span * 10, chain)

    for batch_index in range(3):
        cipher_batch = make_batch(3)
        tx = executor_contract.executeCipherBatch(
            ZERO_HASH32, cipher_batch, 0, {"from": keypers[0]}
        )
        assert len(tx.events["BatchExecuted"]) == 1
        assert tx.events["BatchExecuted"][0] == {
            "numExecutionHalfSteps": batch_index * 2 + 1,
            "batchHash": encode_hex(compute_batch_hash(cipher_batch)),
        }

        plain_batch = make_batch(3)
        mock_batcher_contract.setBatchHash(batch_index, 1, compute_batch_hash(plain_batch))
        tx = executor_contract.executePlainBatch(plain_batch)
        assert len(tx.events["BatchExecuted"]) == 1
        assert tx.events["BatchExecuted"][0] == {
            "numExecutionHalfSteps": batch_index * 2 + 2,
            "batchHash": encode_hex(compute_batch_hash(plain_batch)),
        }


def test_check_plain_batch_hash(
    executor_contract: Any,
    config_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    batch = make_batch(3)
    mock_batcher_contract.setBatchHash(0, 1, compute_batch_hash(batch))

    mine_until(config.start_block_number + config.batch_span - 1, chain)
    executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})

    with brownie.reverts():
        executor_contract.executePlainBatch([])
    with brownie.reverts():
        executor_contract.executePlainBatch(batch[:-1])
    with brownie.reverts():
        executor_contract.executePlainBatch(batch + [make_bytes()])
    executor_contract.executePlainBatch(batch)


def test_check_cipher_batch_hash(
    executor_contract: Any,
    config_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    batch_hash = make_bytes(32)
    mock_batcher_contract.setBatchHash(0, 0, batch_hash)

    mine_until(config.start_block_number + config.batch_span - 1, chain)

    with brownie.reverts():
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})
    with brownie.reverts():
        executor_contract.executeCipherBatch(make_bytes(32), [], 0, {"from": keypers[0]})
    executor_contract.executeCipherBatch(batch_hash, [], 0, {"from": keypers[0]})


def test_check_keyper(
    executor_contract: Any,
    config_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    non_owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    with brownie.reverts():
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 1, {"from": keypers[0]})
    with brownie.reverts():
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[1]})
    with brownie.reverts():
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": non_owner})
    executor_contract.executeCipherBatch(ZERO_HASH32, [], 1, {"from": keypers[1]})


def test_cipher_execution_stores_receipt(
    executor_contract: Any,
    config_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})
    receipt = executor_contract.cipherExecutionReceipts(0)
    assert receipt[0] is True  # executed
    assert receipt[1] == keypers[0]  # executor
    assert receipt[2] == 0  # half step
    assert bytes(receipt[3]) == ZERO_HASH32  # batch hash


def test_cipher_execution_forbidden(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=20,
        execution_timeout=30,
        keypers=keypers,
    )
    schedule_config(config_contract, config, owner=owner)

    batch_index = 0
    forbidden_from_block = (
        config.start_block_number
        + (batch_index + 1) * config.batch_span
        + config.execution_timeout
    )

    mine_until(
        forbidden_from_block - 2, chain,
    )
    chain.snapshot()

    # test that skipCipherExecution still fails
    with brownie.reverts("ExecutorContract: execution timeout not reached yet"):
        executor_contract.skipCipherExecution()

    # test that we could still executeCipherBatch for block forbidden_from_block -1
    chain.revert()
    executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})

    # test that we cannot call executeCipherBatch for block forbidden_from_block
    chain.revert()
    mine_until(
        forbidden_from_block - 1, chain,
    )
    with brownie.reverts("ExecutorContract: execution timeout already reached"):
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})


def test_skip_cipher_execution(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        execution_timeout=300,
    )
    schedule_config(config_contract, config, owner=owner)

    for batch_index in range(3):
        mine_until(
            config.start_block_number
            + (batch_index + 1) * config.batch_span
            + config.execution_timeout
            - 1,
            chain,
        )
        tx = executor_contract.skipCipherExecution()
        assert executor_contract.numExecutionHalfSteps() == batch_index * 2 + 1
        assert len(tx.events) == 1
        assert tx.events["CipherExecutionSkipped"][0] == {
            "numExecutionHalfSteps": batch_index * 2 + 1
        }
        executor_contract.executePlainBatch([])


def test_skip_cipher_execution_checks_half_step(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
    keypers: List[Account],
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        execution_timeout=300,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)

    for batch_index in range(3):
        mine_until(
            config.start_block_number
            + batch_index * config.batch_span
            + config.execution_timeout
            - 1,
            chain,
        )
        executor_contract.executeCipherBatch(ZERO_HASH32, [], 0, {"from": keypers[0]})
        with brownie.reverts():
            executor_contract.skipCipherExecution()
        executor_contract.executePlainBatch([])


def test_skip_cipher_execution_checks_timeout(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        execution_timeout=300,
    )
    schedule_config(config_contract, config, owner=owner)

    for batch_index in range(3):
        mine_until(
            config.start_block_number
            + (batch_index + 1) * config.batch_span
            + config.execution_timeout
            - 2,
            chain,
        )
        with brownie.reverts():
            executor_contract.skipCipherExecution()
        executor_contract.skipCipherExecution()
        executor_contract.executePlainBatch([])


def test_skip_cipher_execution_checks_active(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        execution_timeout=300,
        batch_span=0,
    )
    schedule_config(config_contract, config, owner=owner)
    with brownie.reverts():
        executor_contract.skipCipherExecution()
