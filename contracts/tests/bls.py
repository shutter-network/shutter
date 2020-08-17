import secrets
from typing import NewType
from typing import Tuple

from eth_utils import keccak
from py_ecc import bn128
from py_ecc.fields import bn128_FQ
from py_ecc.fields import bn128_FQ2
from py_ecc.typing import Point2D


BLSPrivateKey = NewType("BLSPrivateKey", int)
BLSPublicKey = NewType("BLSPublicKey", Tuple[bn128_FQ2, bn128_FQ2])
BLSSignature = NewType("BLSSignature", Tuple[bn128_FQ, bn128_FQ])


def hash_to_g1(message: bytes) -> Point2D[bn128_FQ]:
    message_hash_bytes = keccak(message)
    message_hash_int = int.from_bytes(message_hash_bytes, "big")
    message_hash_point = bn128.multiply(bn128.G1, message_hash_int)
    return message_hash_point


def make_private_key() -> BLSPrivateKey:
    key_int = secrets.randbelow(2 ** 256)
    return BLSPrivateKey(key_int)


def private_to_public_key(private_key: BLSPrivateKey) -> BLSPublicKey:
    public_key = bn128.multiply(bn128.G2, int(private_key))
    assert public_key is not None
    return BLSPublicKey(public_key)


def sign(private_key: BLSPrivateKey, message: bytes) -> BLSSignature:
    message_g1 = hash_to_g1(message)
    signature = bn128.multiply(message_g1, int(private_key))
    assert signature is not None
    return BLSSignature(signature)


def verify(public_key: BLSPublicKey, message: bytes, signature: BLSSignature) -> bool:
    message_g1 = hash_to_g1(message)
    p1 = bn128.pairing(public_key, message_g1)
    p2 = bn128.pairing(bn128.G2, signature)
    return p1 == p2
