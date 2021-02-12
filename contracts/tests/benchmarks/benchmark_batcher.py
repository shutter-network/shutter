import statistics
from typing import Any
from typing import Sequence

import pytest
from brownie.network.account import Account
from brownie.network.state import Chain
from brownie.network.transaction import TransactionReceipt

from tests.contract_helpers import mine_until
from tests.contract_helpers import schedule_config
from tests.factories import make_batch_config
from tests.factories import make_bytes


@pytest.fixture
def config_change_heads_up_blocks() -> int:
    return 300


@pytest.fixture(autouse=True)
def configured(
    config_contract: Any, chain: Chain, config_change_heads_up_blocks: int, owner: Account
) -> None:
    config = make_batch_config(
        start_batch_index=0,
        start_block_number=chain.height + config_change_heads_up_blocks + 20,
        batch_span=100,
    )
    schedule_config(config_contract, config, owner=owner)
    mine_until(config.start_block_number - 1, chain)


def print_gas_summary(txs: Sequence[TransactionReceipt]) -> None:
    assert all(tx.status == 1 for tx in txs)

    gas_used = [tx.gas_used for tx in txs]
    min_gas = min(gas_used) / 1000
    max_gas = max(gas_used) / 1000
    avg_gas = sum(gas_used) / len(txs) / 1000
    med_gas = statistics.median(gas_used) / 1000

    print(f"Median:  {med_gas:.1f}k")
    print(f"Average: {avg_gas:.1f}k")
    print(f"Min:     {min_gas:.1f}k")
    print(f"Max:     {max_gas:.1f}k")
    print(f"Samples: {len(txs)}")


@pytest.mark.skip(reason="benchmark")
def benchmark_add_cipher_txs(batcher_contract: Any) -> None:
    txs = []
    for _ in range(10):
        tx = batcher_contract.addTransaction(0, 0, make_bytes(100))
        txs.append(tx)
    print_gas_summary(txs)


@pytest.mark.skip(reason="benchmark")
def benchmark_add_cipher_txs_by_size(batcher_contract: Any) -> None:
    # first tx takes more gas, so add one in advance to not skew results
    batcher_contract.addTransaction(0, 0, b"\x00")

    txs = []
    sizes = list(range(1, 201, 10))
    for size in sizes:
        tx = batcher_contract.addTransaction(0, 0, make_bytes(size))
        txs.append(tx)
    print_gas_summary(txs)


@pytest.mark.skip(reason="benchmark")
def benchmark_add_plain_txs(batcher_contract: Any) -> None:
    txs = []
    for _ in range(10):
        tx = batcher_contract.addTransaction(0, 1, make_bytes(100))
        txs.append(tx)
    print_gas_summary(txs)


@pytest.mark.skip(reason="benchmark")
def benchmark_add_plain_txs_by_size(batcher_contract: Any) -> None:
    # first tx takes more gas, so add one in advance to not skew results
    batcher_contract.addTransaction(0, 1, b"\x00")

    txs = []
    sizes = list(range(1, 201, 10))
    for size in sizes:
        tx = batcher_contract.addTransaction(0, 1, make_bytes(size))
        txs.append(tx)
    print_gas_summary(txs)
