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

    start_batch_index = 10
    key = make_bytes(32)

    for keyper_index, sender in [
        (1, accounts[5]),
        (1, keypers[2]),
    ]:
        with brownie.reverts("KeyBroadcastContract: sender is not keyper"):
            key_broadcast_contract.vote(
                keyper_index, start_batch_index, key, {"from": sender},
            )
    with brownie.reverts("KeyBroadcastContract: keyper index out of range"):
        key_broadcast_contract.vote(
            3, start_batch_index, key, {"from": keypers[0]},
        )


def test_vote(
    key_broadcast_contract: Any, config_contract: Any, owner: Account, accounts: Sequence[Account]
) -> None:
    keypers = accounts[1:4]
    keyper_addresses = [to_canonical_address(k.address) for k in keypers]
    config = make_batch_config(start_batch_index=0, keypers=keyper_addresses, batch_span=1)
    schedule_config(config_contract, config, owner=owner)

    index1 = 1
    index2 = 2
    sender1 = keypers[index1]
    sender2 = keypers[index2]
    start_batch_index = 10
    key = make_bytes(32)

    tx1 = key_broadcast_contract.vote(index1, start_batch_index, key, {"from": sender1})
    assert len(tx1.events) == 1
    assert tx1.events["Voted"][0] == {
        "keyper": sender1,
        "startBatchIndex": start_batch_index,
        "key": encode_hex(key),
        "numVotes": 1,
    }

    tx2 = key_broadcast_contract.vote(index2, start_batch_index, key, {"from": sender2})
    assert len(tx2.events) == 1
    assert tx2.events["Voted"][0] == {
        "keyper": sender2,
        "startBatchIndex": start_batch_index,
        "key": encode_hex(key),
        "numVotes": 2,
    }

    with brownie.reverts("KeyBroadcastContract: keyper has already voted"):
        key_broadcast_contract.vote(index1, start_batch_index, make_bytes(32), {"from": sender1})
