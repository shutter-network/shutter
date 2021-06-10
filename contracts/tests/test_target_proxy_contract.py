from typing import Any

import brownie
import eth_abi
import eth_utils
from brownie.network.account import Account


def test_call_proxy(
    target_proxy_contract: Any,
    test_proxy_receiver: Any,
    owner: Account,
) -> None:
    message = b"hello"
    data = eth_abi.encode_abi(
        ("address", "bytes"),
        (eth_utils.to_canonical_address(test_proxy_receiver.address), message),
    )
    tx = target_proxy_contract.executeTransaction(data, {"from": owner})
    assert len(tx.events["Called"]) == 1
    assert dict(tx.events["Called"][0]) == {
        "data": eth_utils.encode_hex(message),
    }

    invalid_message = b"\x00hello"
    invalid_data = eth_abi.encode_abi(
        ("address", "bytes"),
        (eth_utils.to_canonical_address(test_proxy_receiver.address), invalid_message),
    )
    with brownie.reverts("TargetProxyContract: call reverted"):
        target_proxy_contract.executeTransaction(invalid_data, {"from": owner})
