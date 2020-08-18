from typing import Any
from typing import Sequence

import brownie
from brownie.network.account import Account
from brownie.network.state import Chain

from tests.bls import make_private_key
from tests.bls import private_to_public_key
from tests.bls import sign
from tests.contract_helpers import public_key_to_contract_format
from tests.factories import make_bytes


def test_registering_stores_key(
    bls_registration_contract: Any, chain: Chain, accounts: Sequence[Account]
) -> None:
    sender = accounts[1]
    block_number = chain.height
    block_hash = chain[block_number].hash
    private_key = make_private_key()
    public_key = private_to_public_key(private_key)
    public_key_contract = public_key_to_contract_format(public_key)
    signature = sign(private_key, block_hash)
    bls_registration_contract.registerPublicKey(
        public_key_contract, block_number, block_hash, signature, {"from": sender}
    )

    assert bls_registration_contract.getPublicKey(sender) == public_key_contract


def test_registering_emits_event(
    bls_registration_contract: Any, chain: Chain, accounts: Sequence[Account]
) -> None:
    sender = accounts[1]
    block_number = chain.height
    block_hash = chain[block_number].hash
    private_key = make_private_key()
    public_key = private_to_public_key(private_key)
    public_key_contract = public_key_to_contract_format(public_key)
    signature = sign(private_key, block_hash)
    tx = bls_registration_contract.registerPublicKey(
        public_key_contract, block_number, block_hash, signature, {"from": sender}
    )

    assert len(tx.events) == 1
    assert tx.events[0] == {
        "sender": sender.address,
        "publicKey": public_key_contract,
    }


def test_cannot_register_twice(
    bls_registration_contract: Any, chain: Chain, accounts: Sequence[Account]
) -> None:
    sender = accounts[1]
    block_number = chain.height
    block_hash = chain[block_number].hash
    private_key = make_private_key()
    public_key = private_to_public_key(private_key)
    public_key_contract = public_key_to_contract_format(public_key)
    signature = sign(private_key, block_hash)
    bls_registration_contract.registerPublicKey(
        public_key_contract, block_number, block_hash, signature, {"from": sender}
    )

    with brownie.reverts():
        bls_registration_contract.registerPublicKey(
            public_key_contract, block_number, block_hash, signature, {"from": sender}
        )


def test_check_key_is_non_zero(bls_registration_contract: Any, chain: Chain) -> None:
    block_number = chain.height
    block_hash = chain[block_number].hash
    private_key = make_private_key()
    signature = sign(private_key, block_hash)
    with brownie.reverts():
        bls_registration_contract.registerPublicKey(
            ((0, 0), (0, 0)), block_number, block_hash, signature
        )


def test_check_block_hash(bls_registration_contract: Any, chain: Chain) -> None:
    block_number = chain.height
    block_hash = make_bytes(32)
    private_key = make_private_key()
    public_key = private_to_public_key(private_key)
    public_key_contract = public_key_to_contract_format(public_key)
    signature = sign(private_key, block_hash)
    with brownie.reverts():
        bls_registration_contract.registerPublicKey(
            public_key_contract, block_number, block_hash, signature
        )


def test_check_block_number_is_recent(bls_registration_contract: Any, chain: Chain) -> None:
    chain.mine(256)
    block_number = chain.height - 256
    block_hash = chain[block_number].hash
    private_key = make_private_key()
    public_key = private_to_public_key(private_key)
    public_key_contract = public_key_to_contract_format(public_key)
    signature = sign(private_key, block_hash)
    with brownie.reverts():
        bls_registration_contract.registerPublicKey(
            public_key_contract, block_number, block_hash, signature
        )


def test_check_signature(bls_registration_contract: Any, chain: Chain) -> None:
    block_number = chain.height
    block_hash = chain[block_number].hash
    private_key = make_private_key()
    public_key = private_to_public_key(private_key)
    public_key_contract = public_key_to_contract_format(public_key)
    signature = sign(private_key, make_bytes(32))
    with brownie.reverts():
        bls_registration_contract.registerPublicKey(
            public_key_contract, block_number, block_hash, signature
        )


def test_getting_missing_key_returns_zero(
    bls_registration_contract: Any, chain: Chain, accounts: Sequence[Account]
) -> None:
    assert bls_registration_contract.getPublicKey(accounts[0]) == ((0, 0), (0, 0))
