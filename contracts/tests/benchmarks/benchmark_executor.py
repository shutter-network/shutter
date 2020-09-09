from typing import Any

import pytest
from brownie.network.account import Account
from brownie.network.contract import ContractContainer
from brownie.network.state import Chain
from brownie.network.transaction import TransactionReceipt
from eth_utils import to_canonical_address

from tests import ecdsa
from tests.contract_helpers import compute_batch_hash
from tests.contract_helpers import compute_decryption_signature_preimage
from tests.contract_helpers import mine_until
from tests.contract_helpers import schedule_config
from tests.factories import make_batch_config
from tests.factories import make_bytes
from tests.factories import make_ecdsa_private_key
from tests.factories import make_signer_indices


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
    config_change_heads_up_blocks: int,
    threshold: int,
    batch_size: int,
    tx_size: int,
) -> TransactionReceipt:
    num_keypers = threshold * 3 // 2
    keyper_private_keys = [make_ecdsa_private_key() for _ in range(num_keypers)]
    keypers = [ecdsa.private_key_to_address(key) for key in keyper_private_keys]
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
    decryption_key = make_bytes(32)
    decrypted_transactions = [make_bytes(tx_size) for _ in range(batch_size)]
    signer_indices = make_signer_indices(len(keypers), threshold)
    decryption_signature_preimage = compute_decryption_signature_preimage(
        batcher_contract_address=to_canonical_address(batcher_contract.address),
        cipher_batch_hash=cipher_batch_hash,
        decryption_key=decryption_key,
        decrypted_transactions=decrypted_transactions,
    )
    signatures = [
        ecdsa.sign(keyper_private_keys[signer_index], decryption_signature_preimage)
        for signer_index in signer_indices
    ]

    return executor_contract.executeCipherBatch(
        cipher_batch_hash, decrypted_transactions, decryption_key, signer_indices, signatures
    )


def benchmark_execute_cipher_batch(
    config_contract: Any,
    batcher_contract: Any,
    executor_contract: Any,
    chain: Chain,
    owner: Account,
    config_change_heads_up_blocks: int,
) -> None:
    batch_sizes = [1, 10, 100]
    thresholds = [1, 10, 100]
    for batch_size in batch_sizes:
        for threshold in thresholds:
            chain.snapshot()
            tx = execute_cipher_batch(
                config_contract=config_contract,
                batcher_contract=batcher_contract,
                executor_contract=executor_contract,
                chain=chain,
                owner=owner,
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
            chain.revert()
