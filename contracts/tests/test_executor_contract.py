from typing import Any
from typing import List

import brownie
from brownie.network.account import Account
from brownie.network.state import Chain
from eth_typing import Hash32
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

    for batch_index in range(3):
        assert executor_contract.numExecutionHalfSteps() == 2 * batch_index

        with brownie.reverts("ExecutorContract: unexpected half step"):
            executor_contract.executePlainBatch(batch_index, [])
        executor_contract.executeCipherBatch(batch_index, ZERO_HASH32, [], 0, {"from": keypers[0]})

        assert executor_contract.numExecutionHalfSteps() == 2 * batch_index + 1

        with brownie.reverts("ExecutorContract: unexpected half step"):
            executor_contract.executeCipherBatch(
                batch_index, ZERO_HASH32, [], 0, {"from": keypers[0]}
            )
        executor_contract.executePlainBatch(batch_index, [])

    assert executor_contract.numExecutionHalfSteps() == 2 * (batch_index + 1)


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

    with brownie.reverts("ExecutorContract: config is inactive"):
        executor_contract.executeCipherBatch(0, ZERO_HASH32, [], 0)


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
            executor_contract.executeCipherBatch(
                batch_index, ZERO_HASH32, [], 0, {"from": keypers[0]}
            )
        executor_contract.executeCipherBatch(batch_index, ZERO_HASH32, [], 0, {"from": keypers[0]})

        plain_batch = make_batch()
        mock_batcher_contract.setBatchHash(batch_index, 1, compute_batch_hash(plain_batch))
        executor_contract.executePlainBatch(batch_index, plain_batch)


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
        batch_span=10,
        keypers=keypers,
        threshold=0,
        target_address=to_canonical_address(mock_target_contract.address),
        target_function_selector=mock_target_function_selector,
        transaction_gas_limit=5000,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    batch = make_batch(3)
    cipher_batch_hash = Hash32(b"\xde" * 32)
    mock_batcher_contract.setBatchHash(0, 0, cipher_batch_hash)
    tx = executor_contract.executeCipherBatch(0, cipher_batch_hash, batch, 0, {"from": keypers[0]})
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
    tx = executor_contract.executePlainBatch(0, batch)
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
        batch_span=10,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span * 10, chain)

    for batch_index in range(3):
        cipher_batch = make_batch(3)
        cipher_batch_hash = Hash32(b"\xfc" * 32)
        mock_batcher_contract.setBatchHash(batch_index, 0, cipher_batch_hash)
        tx = executor_contract.executeCipherBatch(
            batch_index, cipher_batch_hash, cipher_batch, 0, {"from": keypers[0]}
        )
        assert len(tx.events["BatchExecuted"]) == 1
        assert tx.events["BatchExecuted"][0] == {
            "numExecutionHalfSteps": batch_index * 2 + 1,
            "batchHash": encode_hex(compute_batch_hash(cipher_batch)),
        }

        plain_batch = make_batch(3)
        mock_batcher_contract.setBatchHash(batch_index, 1, compute_batch_hash(plain_batch))
        tx = executor_contract.executePlainBatch(batch_index, plain_batch)
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
        batch_span=10,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    batch = make_batch(3)
    mock_batcher_contract.setBatchHash(0, 1, compute_batch_hash(batch))

    mine_until(config.start_block_number + config.batch_span - 1, chain)
    executor_contract.executeCipherBatch(0, ZERO_HASH32, [], 0, {"from": keypers[0]})

    with brownie.reverts("ExecutorContract: batch hash does not match"):
        executor_contract.executePlainBatch(0, [])
    with brownie.reverts("ExecutorContract: batch hash does not match"):
        executor_contract.executePlainBatch(0, batch[:-1])
    with brownie.reverts("ExecutorContract: batch hash does not match"):
        executor_contract.executePlainBatch(0, batch + [make_bytes()])
    executor_contract.executePlainBatch(0, batch)


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
        batch_span=10,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    batch_hash = make_bytes(32)
    mock_batcher_contract.setBatchHash(0, 0, batch_hash)

    mine_until(config.start_block_number + config.batch_span - 1, chain)

    with brownie.reverts("ExecutorContract: incorrect cipher batch hash"):
        executor_contract.executeCipherBatch(0, ZERO_HASH32, [], 0, {"from": keypers[0]})
    with brownie.reverts("ExecutorContract: incorrect cipher batch hash"):
        executor_contract.executeCipherBatch(0, make_bytes(32), [b""], 0, {"from": keypers[0]})
    executor_contract.executeCipherBatch(0, batch_hash, [b""], 0, {"from": keypers[0]})


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
        batch_span=10,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)
    batch_index = 0
    with brownie.reverts("ExecutorContract: sender is not specified keyper"):
        executor_contract.executeCipherBatch(batch_index, ZERO_HASH32, [], 1, {"from": keypers[0]})
    with brownie.reverts("ExecutorContract: sender is not specified keyper"):
        executor_contract.executeCipherBatch(batch_index, ZERO_HASH32, [], 0, {"from": keypers[1]})
    with brownie.reverts("ExecutorContract: sender is not specified keyper"):
        executor_contract.executeCipherBatch(batch_index, ZERO_HASH32, [], 0, {"from": non_owner})
    executor_contract.executeCipherBatch(batch_index, ZERO_HASH32, [], 1, {"from": keypers[1]})


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
        batch_span=10,
        keypers=keypers,
        threshold=0,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    executor_contract.executeCipherBatch(0, ZERO_HASH32, [], 0, {"from": keypers[0]})
    receipt = executor_contract.cipherExecutionReceipts(0)
    assert receipt[0] is True  # executed
    assert receipt[1] == keypers[0]  # executor
    assert receipt[2] == 0  # half step
    assert bytes(receipt[3]) == ZERO_HASH32  # batch hash


def test_cipher_execution_skips_after_timeout(
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
        execution_timeout=20,
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
        executor_contract.skipCipherExecution(batch_index)

    # test that we could still executeCipherBatch for block forbidden_from_block -1
    chain.revert()
    tx = executor_contract.executeCipherBatch(
        batch_index, ZERO_HASH32, [], 0, {"from": keypers[0]}
    )
    assert tx.events["BatchExecuted"]
    # test that calling executeCipherBatch after the timeout results in skipping the cipher
    # execution
    chain.revert()
    mine_until(
        forbidden_from_block - 1, chain,
    )
    tx = executor_contract.executeCipherBatch(
        batch_index, ZERO_HASH32, [], 0, {"from": keypers[0]}
    )
    assert len(tx.events) == 1
    assert tx.events["CipherExecutionSkipped"]["numExecutionHalfSteps"] == 1


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
        batch_span=10,
        execution_timeout=20,
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
        tx = executor_contract.skipCipherExecution(batch_index)
        assert executor_contract.numExecutionHalfSteps() == batch_index * 2 + 1
        assert len(tx.events) == 1
        assert tx.events["CipherExecutionSkipped"][0] == {
            "numExecutionHalfSteps": batch_index * 2 + 1
        }
        executor_contract.executePlainBatch(batch_index, [])


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
        batch_span=10,
        execution_timeout=20,
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
        executor_contract.executeCipherBatch(batch_index, ZERO_HASH32, [], 0, {"from": keypers[0]})
        with brownie.reverts("ExecutorContract: unexpected half step"):
            executor_contract.skipCipherExecution(batch_index)
        executor_contract.executePlainBatch(batch_index, [])


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
        batch_span=10,
        execution_timeout=20,
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
        with brownie.reverts("ExecutorContract: execution timeout not reached yet"):
            executor_contract.skipCipherExecution(batch_index)
        executor_contract.skipCipherExecution(batch_index)
        executor_contract.executePlainBatch(batch_index, [])


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
        execution_timeout=20,
        batch_span=0,
    )
    schedule_config(config_contract, config, owner=owner)
    with brownie.reverts("ExecutorContract: config is inactive"):
        executor_contract.skipCipherExecution(0)
