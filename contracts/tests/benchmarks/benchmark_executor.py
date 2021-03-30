#! pytest -s

from typing import Any

import pytest
from brownie.network.account import Account
from brownie.network.account import Accounts
from brownie.network.contract import ContractContainer
from brownie.network.state import Chain
from brownie.network.transaction import TransactionReceipt

from tests.contract_helpers import compute_batch_hash
from tests.contract_helpers import mine_until
from tests.contract_helpers import schedule_config
from tests.factories import make_batch_config
from tests.factories import make_bytes


@pytest.fixture
def config_change_heads_up_blocks() -> int:
    return 300


@pytest.fixture
def executor_contract(
    ExecutorContract: ContractContainer,
    config_contract: Any,
    batcher_contract: Any,
    owner: Account,
) -> Any:
    # overrides executor_contract fixture in tests/conftest.py which uses mock_batcher_contract
    executor_contract = owner.deploy(ExecutorContract, config_contract, batcher_contract)
    return executor_contract


def execute_cipher_batch(
    *,
    config_contract: Any,
    batcher_contract: Any,
    executor_contract: Any,
    chain: Chain,
    owner: Account,
    keypers: Accounts,
    config_change_heads_up_blocks: int,
    threshold: int,
    batch_size: int,
    tx_size: int,
) -> TransactionReceipt:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
        keypers=keypers,
        threshold=1,
    )
    schedule_config(config_contract, config, owner=owner)

    mine_until(config.start_block_number - 1, chain)
    batch = [make_bytes(tx_size) for _ in range(batch_size)]
    for tx in batch:
        batcher_contract.addTransaction(0, 0, tx)
    cipher_batch_hash = compute_batch_hash(batch)

    mine_until(config.start_block_number + config.batch_span, chain)
    decrypted_transactions = [make_bytes(tx_size) for _ in range(batch_size)]

    return executor_contract.executeCipherBatch(
        0, cipher_batch_hash, decrypted_transactions, 0, {"from": keypers[0]}
    )


@pytest.mark.skip(reason="benchmark")
def benchmark_execute_cipher_batch(
    config_contract: Any,
    batcher_contract: Any,
    executor_contract: Any,
    chain: Chain,
    owner: Account,
    keypers: Accounts,
    config_change_heads_up_blocks: int,
) -> None:
    batch_sizes = [0, 1, 10, 100]
    thresholds = [1, 10, 100]
    chain.snapshot()
    for batch_size in batch_sizes:
        for threshold in thresholds:
            chain.revert()
            tx = execute_cipher_batch(
                config_contract=config_contract,
                batcher_contract=batcher_contract,
                executor_contract=executor_contract,
                chain=chain,
                owner=owner,
                keypers=keypers,
                config_change_heads_up_blocks=config_change_heads_up_blocks,
                threshold=threshold,
                batch_size=batch_size,
                tx_size=100,
            )
            print(
                f"Batch size: {batch_size:>3d}  "
                f"Threshold: {threshold:>3d}  "
                f"Gas: {tx.gas_used / 1000:.1f}k"
            )
