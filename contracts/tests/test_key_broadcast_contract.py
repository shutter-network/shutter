from typing import Any
from typing import Sequence

import brownie
from brownie.network.account import Account
from eth_utils import encode_hex
from eth_utils import to_canonical_address

from tests.contract_helpers import schedule_config
from tests.factories import make_batch_config
from tests.factories import make_bytes


def test_broadcasting_checks_sender(
    key_broadcast_contract: Any, config_contract: Any, owner: Account, accounts: Sequence[Account]
) -> None:
    keypers = accounts[1:4]
    keyper_addresses = [to_canonical_address(k.address) for k in keypers]
    config = make_batch_config(start_batch_index=0, keypers=keyper_addresses, batch_span=1)
    schedule_config(config_contract, config, owner=owner)

    batch_index = 10
    encryption_key = make_bytes(32)
    signer_indices = [0, 1]
    signatures = [make_bytes() for _ in signer_indices]

    for keyper_index, sender in [
        (1, accounts[5]),
        (1, keypers[2]),
        (3, keypers[0]),
    ]:
        with brownie.reverts():
            key_broadcast_contract.broadcastEncryptionKey(
                keyper_index,
                batch_index,
                encryption_key,
                signer_indices,
                signatures,
                {"from": sender},
            )


def test_emit_event(
    key_broadcast_contract: Any, config_contract: Any, owner: Account, accounts: Sequence[Account]
) -> None:
    keypers = accounts[1:4]
    keyper_addresses = [to_canonical_address(k.address) for k in keypers]
    config = make_batch_config(start_batch_index=0, keypers=keyper_addresses, batch_span=1)
    schedule_config(config_contract, config, owner=owner)

    keyper_index = 1
    sender = keypers[keyper_index]
    batch_index = 10
    encryption_key = make_bytes(32)
    signer_indices = [0, 1]
    signatures = [make_bytes() for _ in signer_indices]

    tx = key_broadcast_contract.broadcastEncryptionKey(
        keyper_index, batch_index, encryption_key, signer_indices, signatures, {"from": sender},
    )

    assert len(tx.events) == 1
    assert tx.events[0] == {
        "sender": sender,
        "batchIndex": batch_index,
        "encryptionKey": encode_hex(encryption_key),
        "signerIndices": signer_indices,
        "signatures": [encode_hex(s) for s in signatures],
    }
