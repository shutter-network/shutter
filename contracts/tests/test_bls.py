from typing import Any

from eth_utils import decode_hex
from py_ecc.fields import bn128_FQ
from py_ecc.fields import bn128_FQ2

from tests.bls import aggregate_public_keys
from tests.bls import aggregate_signatures
from tests.bls import BLSPublicKey
from tests.bls import BLSSignature
from tests.bls import make_private_key
from tests.bls import private_to_public_key
from tests.bls import sign
from tests.bls import verify
from tests.contract_helpers import public_key_to_contract_format


def test_round_trip() -> None:
    sk = make_private_key()
    pk = private_to_public_key(sk)
    msg = b"msg"
    sig = sign(sk, msg)
    assert verify(pk, msg, sig)
    assert not verify(private_to_public_key(make_private_key()), msg, sig)
    assert not verify(pk, b"different msg", sig)
    assert not verify(pk, msg, sign(make_private_key(), msg))
    assert not verify(pk, msg, sign(sk, b"different msg"))


def test_contract_data() -> None:
    # uses data from https://github.com/kfichter/solidity-bls/blob/master/test/bls.js and
    # https://github.com/kfichter/solidity-bls/blob/master/contracts/BLSTest.sol
    msg_hex = (
        "0x7b0a2020226f70656e223a207b0a20202020227072696365223a2039353931372c0a"
        "202020202274696d65223a207b0a20202020202022756e6978223a2031343833313432"
        "3430302c0a2020202020202269736f223a2022323031362d31322d33315430303a3030"
        "3a30302e3030305a220a202020207d0a20207d2c0a202022636c6f7365223a207b0a20"
        "202020227072696365223a2039363736302c0a202020202274696d65223a207b0a2020"
        "2020202022756e6978223a20313438333232383830302c0a2020202020202269736f22"
        "3a2022323031372d30312d30315430303a30303a30302e3030305a220a202020207d0a"
        "20207d2c0a2020226c6f6f6b7570223a207b0a20202020227072696365223a20393637"
        "36302c0a20202020226b223a20312c0a202020202274696d65223a207b0a2020202020"
        "2022756e6978223a20313438333232383830302c0a2020202020202269736f223a2022"
        "323031372d30312d30315430303a30303a30302e3030305a220a202020207d0a20207d"
        "0a7d0a6578616d706c652e636f6d2f6170692f31"
    )
    sig = BLSSignature(
        (
            bn128_FQ(
                11181692345848957662074290878138344227085597134981019040735323471731897153462
            ),
            bn128_FQ(6479746447046570360435714249272776082787932146211764251347798668447381926167),
        )
    )
    pk_contract = (
        (
            18523194229674161632574346342370534213928970227736813349975332190798837787897,
            5725452645840548248571879966249653216818629536104756116202892528545334967238,
        ),
        (
            3816656720215352836236372430537606984911914992659540439626020770732736710924,
            677280212051826798882467475639465784259337739185938192379192340908771705870,
        ),
    )
    pk = BLSPublicKey((bn128_FQ2(pk_contract[0][::-1]), bn128_FQ2(pk_contract[1][::-1]),))
    msg = decode_hex(msg_hex)

    assert verify(pk, msg, sig)


def test_bls_contract(test_bls_contract: Any) -> None:
    sk = make_private_key()
    pk = private_to_public_key(sk)
    msg = b"msg"
    sig = sign(sk, msg)

    test_bls_contract.verify(public_key_to_contract_format(pk), msg, sig)


def test_aggregate_single_public_key() -> None:
    single_public_key = private_to_public_key(make_private_key())
    assert aggregate_public_keys([single_public_key]) == single_public_key


def test_aggregate_two_public_keys() -> None:
    two_private_keys = [make_private_key() for _ in range(2)]
    two_public_keys = [private_to_public_key(sk) for sk in two_private_keys]
    aggregated_public_key = aggregate_public_keys(two_public_keys)
    assert aggregate_public_keys(list(reversed(two_public_keys))) == aggregated_public_key


def test_aggregate_single_signature() -> None:
    sk = make_private_key()
    sig = sign(sk, b"msg")
    aggregated_signature = aggregate_signatures([sig])
    assert aggregated_signature == sig


def test_aggregate_two_signatures() -> None:
    two_private_keys = [make_private_key() for _ in range(2)]
    two_public_keys = [private_to_public_key(sk) for sk in two_private_keys]
    msg = b"msg"
    individual_signatures = [sign(sk, msg) for sk in two_private_keys]

    aggregated_signature = aggregate_signatures(individual_signatures)
    aggregated_public_key = aggregate_public_keys(two_public_keys)
    assert verify(aggregated_public_key, msg, aggregated_signature)
