from typing import Any
from typing import Sequence

import brownie
from brownie.network.account import Account
from brownie.network.state import Chain
from eth_utils import encode_hex
from eth_utils import keccak
from eth_utils import to_canonical_address

from tests.contract_helpers import mine_until
from tests.contract_helpers import schedule_config
from tests.factories import make_batch_config


def test_add_tx_fails_if_not_active(
    batcher_contract: Any,
    config_contract: Any,
    chain: Chain,
    config_change_heads_up_blocks: int,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0, start_block_number=chain.height + config_change_heads_up_blocks + 20,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    with brownie.reverts():
        batcher_contract.addTransaction(0, 0, b"\x00")


def test_add_tx_checks_batching_period_end(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number + 98, chain)

    batcher_contract.addTransaction(0, 0, b"\x00")
    assert chain.height == config.start_block_number + 99
    with brownie.reverts():
        batcher_contract.addTransaction(0, 0, b"\x00")
    batcher_contract.addTransaction(1, 0, b"\x00")


def test_add_tx_checks_current_batch(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number, chain)

    with brownie.reverts():
        batcher_contract.addTransaction(1, 0, b"\x00")
    mine_until(config.start_block_number + config.batch_span, chain)
    batcher_contract.addTransaction(1, 0, b"\x00")


def test_add_tx_checks_tx_size(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        transaction_size_limit=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)

    with brownie.reverts():
        batcher_contract.addTransaction(0, 0, b"")

    batcher_contract.addTransaction(0, 0, b"\x00" * 100)
    with brownie.reverts():
        batcher_contract.addTransaction(0, 0, b"\x00" * 101)


def test_add_tx_checks_batch_size(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        transaction_size_limit=100,
        batch_size_limit=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)

    for _ in range(10):
        batcher_contract.addTransaction(0, 0, b"\x00" * 10)
    with brownie.reverts():
        batcher_contract.addTransaction(0, 0, b"\x00")


def test_add_tx_checks_fee(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)

    batcher_contract.setMinFee(100, {"from": owner})
    with brownie.reverts():
        batcher_contract.addTransaction(0, 0, b"\x00", {"value": 99})
    batcher_contract.addTransaction(0, 0, b"\x00", {"value": 100})


def test_add_tx_updates_hash_chain(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)

    for tx_type in [0, 1]:
        assert bytes(batcher_contract.batchHashes(0, tx_type)) == b"\x00" * 32
        batcher_contract.addTransaction(0, tx_type, b"\x11")
        assert bytes(batcher_contract.batchHashes(0, tx_type)) == keccak(b"\x11" + b"\x00" * 32)
        batcher_contract.addTransaction(0, tx_type, b"\x22")
        assert bytes(batcher_contract.batchHashes(0, tx_type)) == keccak(
            b"\x22" + keccak(b"\x11" + b"\x00" * 32)
        )


def test_add_tx_updates_batch_size(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)

    assert batcher_contract.batchSizes(0) == 0
    batcher_contract.addTransaction(0, 0, b"\x00" * 3)
    assert batcher_contract.batchSizes(0) == 3
    batcher_contract.addTransaction(0, 1, b"\x00" * 5)
    assert batcher_contract.batchSizes(0) == 8


def test_add_tx_emits_event(
    batcher_contract: Any,
    config_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)

    tx = batcher_contract.addTransaction(0, 0, b"\x11")
    assert len(tx.events) == 1
    assert tx.events[0] == {
        "batchIndex": 0,
        "transactionType": 0,
        "transaction": "0x11",
        "batchHash": encode_hex(keccak(b"\x11" + b"\x00" * 32)),
    }

    mine_until(config.start_block_number + 2 * config.batch_span + 1, chain)
    tx = batcher_contract.addTransaction(2, 1, b"\x22")
    assert len(tx.events) == 1
    assert tx.events[0] == {
        "batchIndex": 2,
        "transactionType": 1,
        "transaction": "0x22",
        "batchHash": encode_hex(keccak(b"\x22" + b"\x00" * 32)),
    }


def test_add_tx_pays_fee(
    batcher_contract: Any,
    config_contract: Any,
    fee_bank_contract: Any,
    config_change_heads_up_blocks: int,
    chain: Chain,
    owner: Account,
    accounts: Sequence[Account],
) -> None:
    fee_receiver = accounts[1]
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        fee_receiver=to_canonical_address(fee_receiver.address),
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)

    initial_balance = fee_bank_contract.deposits(config.fee_receiver)
    batcher_contract.addTransaction(0, 0, b"\x00", {"value": 150})

    assert fee_bank_contract.deposits(config.fee_receiver) == initial_balance + 150


def test_set_fee(batcher_contract: Any, owner: Account) -> None:
    assert batcher_contract.minFee() == 0
    batcher_contract.setMinFee(100, {"from": owner})
    assert batcher_contract.minFee() == 100


def test_non_owner_cannot_set_fee(batcher_contract: Any, non_owner: Account) -> None:
    with brownie.reverts():
        batcher_contract.setMinFee(100, {"from": non_owner})
