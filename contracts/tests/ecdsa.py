from typing import NewType

from eth_keys import keys
from eth_typing import Address
from eth_utils import int_to_big_endian
from eth_utils import keccak


ECDSAPrivateKey = NewType("ECDSAPrivateKey", bytes)


def private_key_to_address(private_key: bytes) -> Address:
    private_key_object = keys.PrivateKey(private_key)
    return Address(private_key_object.public_key.to_canonical_address())


def sign(private_key: bytes, message: bytes) -> bytes:
    private_key_object = keys.PrivateKey(private_key)
    msg_hash = keccak(message)
    signature_object = private_key_object.sign_msg_hash(msg_hash)
    r_bytes = int_to_big_endian(signature_object.r).rjust(32, b"\x00")
    s_bytes = int_to_big_endian(signature_object.s).rjust(32, b"\x00")
    v_bytes = bytes([signature_object.v + 27])
    return r_bytes + s_bytes + v_bytes
