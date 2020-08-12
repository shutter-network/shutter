from typing import Any

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
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=10,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span * 10, chain)

    for i in range(3):
        assert executor_contract.numExecutionHalfSteps() == 2 * i

        with brownie.reverts():
            executor_contract.executePlainBatch([])
        executor_contract.executeCipherBatch(ZERO_HASH32, [], ZERO_HASH32, ZERO_HASH32)

        assert executor_contract.numExecutionHalfSteps() == 2 * i + 1

        with brownie.reverts():
            executor_contract.executeCipherBatch(ZERO_HASH32, [], ZERO_HASH32, ZERO_HASH32)
        executor_contract.executePlainBatch([])

    assert executor_contract.numExecutionHalfSteps() == 2 * (i + 1)


def test_check_batching_is_active(
    executor_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        active=False,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    with brownie.reverts():
        executor_contract.executeCipherBatch(ZERO_HASH32, [], ZERO_HASH32, ZERO_HASH32)


def test_check_batching_period_is_over(
    executor_contract: Any,
    config_contract: Any,
    mock_batcher_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=10,
    )
    schedule_config(config_contract, config, owner=owner)

    for batch_index in range(3):
        mine_until(config.start_block_number + (batch_index + 1) * config.batch_span - 2, chain)
        with brownie.reverts():
            executor_contract.executeCipherBatch(ZERO_HASH32, [], ZERO_HASH32, ZERO_HASH32)
        executor_contract.executeCipherBatch(ZERO_HASH32, [], ZERO_HASH32, ZERO_HASH32)

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
    mock_target_function_selector: bytes,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        target_address=to_canonical_address(mock_target_contract.address),
        target_function_selector=mock_target_function_selector,
        transaction_gas_limit=5000,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span, chain)

    batch = make_batch(3)
    tx = executor_contract.executeCipherBatch(ZERO_HASH32, batch, ZERO_HASH32, ZERO_HASH32)
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
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + config.batch_span * 10, chain)

    for batch_index in range(3):
        cipher_batch = make_batch(3)
        tx = executor_contract.executeCipherBatch(
            ZERO_HASH32, cipher_batch, ZERO_HASH32, ZERO_HASH32
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
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    batch = make_batch(3)
    mock_batcher_contract.setBatchHash(0, 1, compute_batch_hash(batch))

    mine_until(config.start_block_number + config.batch_span - 1, chain)
    executor_contract.executeCipherBatch(ZERO_HASH32, [], ZERO_HASH32, ZERO_HASH32)

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
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    batch_hash = make_bytes(32)
    mock_batcher_contract.setBatchHash(0, 0, batch_hash)

    mine_until(config.start_block_number + config.batch_span - 1, chain)

    with brownie.reverts():
        executor_contract.executeCipherBatch(ZERO_HASH32, [], ZERO_HASH32, ZERO_HASH32)
    with brownie.reverts():
        executor_contract.executeCipherBatch(make_bytes(32), [], ZERO_HASH32, ZERO_HASH32)
    executor_contract.executeCipherBatch(batch_hash, [], ZERO_HASH32, ZERO_HASH32)


def test_executing_cipher_batch_checks_signature() -> None:
    pass
