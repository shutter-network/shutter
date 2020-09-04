from typing import Any
from typing import Sequence

import brownie
import pytest
from brownie.network.account import Account


@pytest.fixture
def depositor(accounts: Sequence[Account]) -> Account:
    return accounts[1]


@pytest.fixture
def receiver(accounts: Sequence[Account]) -> Account:
    return accounts[2]


def test_deposit(fee_bank_contract: Any, depositor: Account, receiver: Account) -> None:
    assert fee_bank_contract.deposits(receiver) == 0

    amounts = [100, 200, 1, 500]
    for i, amount in enumerate(amounts):
        tx = fee_bank_contract.deposit(receiver, {"from": depositor, "value": amount})
        assert fee_bank_contract.deposits(receiver) == sum(amounts[: i + 1])
        assert len(tx.events) == 1
        assert tx.events[0] == {
            "depositor": depositor,
            "receiver": receiver,
            "amount": amount,
            "totalAmount": sum(amounts[: i + 1]),
        }


def test_withdraw(
    fee_bank_contract: Any, depositor: Account, receiver: Account, accounts: Sequence[Account]
) -> None:
    fee_bank_contract.deposit(receiver, {"from": depositor, "value": 1000})

    with brownie.reverts():
        fee_bank_contract.withdraw({"from": depositor})
    with brownie.reverts():
        fee_bank_contract.withdraw(receiver, 1001, {"from": receiver})

    different_receiver = accounts[-1]
    assert different_receiver != receiver
    receiver_pre_balance = receiver.balance()
    different_receiver_pre_balance = different_receiver.balance()

    tx = fee_bank_contract.withdraw(different_receiver, 100, {"from": receiver})
    assert fee_bank_contract.deposits(receiver) == 900
    assert different_receiver.balance() == different_receiver_pre_balance + 100
    assert receiver.balance() == receiver_pre_balance
    assert len(tx.events) == 1
    assert tx.events[0] == {
        "sender": receiver,
        "receiver": different_receiver,
        "amount": 100,
        "totalAmount": 900,
    }

    tx = fee_bank_contract.withdraw({"from": receiver})
    assert fee_bank_contract.deposits(receiver) == 0
    assert receiver.balance() == receiver_pre_balance + 900
    assert different_receiver.balance() == different_receiver_pre_balance + 100
    assert len(tx.events) == 1
    assert tx.events[0] == {
        "sender": receiver,
        "receiver": receiver,
        "amount": 900,
        "totalAmount": 0,
    }
