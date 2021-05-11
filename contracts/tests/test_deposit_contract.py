from typing import Any

import brownie
from brownie.network.account import Account
from brownie.network.account import Accounts
from brownie.network.contract import ContractTx
from brownie.network.state import Chain
from brownie.network.web3 import Web3
from eth_utils import to_checksum_address

from tests.contract_helpers import encode_withdrawal_delay
from tests.contract_helpers import mine_until
from tests.contract_helpers import ZERO_ADDRESS
from tests.factories import make_int


def check_deposit(
    deposit_contract: Any,
    account: Account,
    amount: int,
    withdrawal_delay_blocks: int,
    withdrawal_requested_block: int,
    slashed: bool = False,
) -> None:
    assert deposit_contract.getDepositAmount(account) == amount
    assert deposit_contract.getWithdrawalDelayBlocks(account) == withdrawal_delay_blocks
    assert deposit_contract.getWithdrawalRequestedBlock(account) == withdrawal_requested_block
    assert deposit_contract.isSlashed(account) == slashed


def check_deposit_changed_event(
    tx: ContractTx,
    amount: int,
    withdrawal_delay_blocks: int,
    withdrawal_requested_block: int,
    *,
    slashed: bool = False,
    withdrawn: bool = False,
) -> None:
    assert "DepositChanged" in tx.events and len(tx.events["DepositChanged"]) == 1
    event = tx.events["DepositChanged"][0]
    assert event["account"] == tx.sender
    assert event["amount"] == amount
    assert event["withdrawalDelayBlocks"] == withdrawal_delay_blocks
    assert event["withdrawalRequestedBlock"] == withdrawal_requested_block
    assert event["withdrawn"] is withdrawn
    assert event["slashed"] is slashed


def test_deposit(deposit_contract: Any, deposit_token_contract: Any, owner: Account) -> None:
    invalid_data = ["", encode_withdrawal_delay(123)[:-2]]
    for data in invalid_data:
        with brownie.reverts():
            deposit_token_contract.send(deposit_contract, 100, data, {"from": owner})

    withdrawal_delay_blocks = make_int()
    valid_data = encode_withdrawal_delay(withdrawal_delay_blocks)

    tx = deposit_token_contract.send(deposit_contract, 100, valid_data, {"from": owner})
    check_deposit_changed_event(tx, 100, withdrawal_delay_blocks, 0)
    check_deposit(deposit_contract, owner, 100, withdrawal_delay_blocks, 0)

    # depositing twice works, but only if withdrawal delay isn't decreased
    data_too_short = encode_withdrawal_delay(withdrawal_delay_blocks - 1)
    with brownie.reverts():
        deposit_token_contract.send(deposit_contract, 50, data_too_short, {"from": owner})

    tx = deposit_token_contract.send(deposit_contract, 50, valid_data, {"from": owner})
    check_deposit_changed_event(tx, 150, withdrawal_delay_blocks, 0)
    check_deposit(deposit_contract, owner, 150, withdrawal_delay_blocks, 0)

    data_longer = encode_withdrawal_delay(withdrawal_delay_blocks + 1)
    tx = deposit_token_contract.send(deposit_contract, 0, data_longer, {"from": owner})
    check_deposit_changed_event(tx, 150, withdrawal_delay_blocks + 1, 0)
    check_deposit(deposit_contract, owner, 150, withdrawal_delay_blocks + 1, 0)


def test_withdraw(
    deposit_contract: Any,
    deposit_token_contract: Any,
    owner: Account,
    accounts: Accounts,
    chain: Chain,
    web3: Web3,
) -> None:
    withdrawal_delay = 10
    withdrawal_delay_data = encode_withdrawal_delay(withdrawal_delay)
    deposit_token_contract.send(deposit_contract, 100, withdrawal_delay_data)
    recipient = accounts[-1]

    chain.mine(withdrawal_delay)
    with brownie.reverts():
        deposit_contract.withdraw(recipient, {"from": owner})

    tx = deposit_contract.requestWithdrawal({"from": owner})
    check_deposit_changed_event(tx, 100, withdrawal_delay, tx.block_number)
    assert deposit_contract.getWithdrawalRequestedBlock(owner) == tx.block_number
    withdraw_block_number = tx.block_number + withdrawal_delay

    with brownie.reverts():
        deposit_contract.withdraw(recipient, {"from": owner})
    mine_until(withdraw_block_number - 2, chain)
    with brownie.reverts():
        deposit_contract.withdraw(recipient, {"from": owner})
    assert web3.eth.block_number == withdraw_block_number - 1  # one block mined by sending tx
    tx = deposit_contract.withdraw(recipient, {"from": owner})

    check_deposit_changed_event(tx, 0, 0, 0, withdrawn=True)
    check_deposit(deposit_contract, owner, 0, 0, 0)
    assert deposit_token_contract.balanceOf(recipient) == 100


def test_slash_and_burn(
    deposit_contract: Any,
    deposit_token_contract: Any,
    owner: Account,
    accounts: Accounts,
) -> None:
    slashed = accounts[1]
    slasher = accounts[2]
    deposit_contract.setSlasher(slasher, {"from": owner})

    deposit_token_contract.send(slashed, 100, "", {"from": owner})
    deposit_token_contract.send(
        deposit_contract, 100, encode_withdrawal_delay(0), {"from": slashed}
    )

    tx = deposit_contract.slash(slashed, {"from": slasher})
    assert len(tx.events) == 3
    ev1, ev2, ev3 = tx.events

    assert ev1.name == "DepositChanged"
    assert ev1 == {
        "account": slashed,
        "amount": 0,
        "withdrawalDelayBlocks": 0,
        "withdrawalRequestedBlock": 0,
        "withdrawn": False,
        "slashed": True,
    }
    assert ev2.name == "Burned"
    assert ev2 == {
        "operator": deposit_contract,
        "from": deposit_contract,
        "amount": 100,
        "data": "0x00",
        "operatorData": "0x00",
    }
    assert ev3.name == "Transfer"
    assert ev3 == {
        "from": deposit_contract,
        "to": to_checksum_address(ZERO_ADDRESS),
        "value": 100,
    }


def test_slash_and_send(
    deposit_contract: Any,
    deposit_token_contract: Any,
    owner: Account,
    accounts: Accounts,
) -> None:
    slashed = accounts[1]
    slasher = accounts[2]
    slashing_receiver = accounts[3]
    deposit_contract.setSlasher(slasher, {"from": owner})
    deposit_contract.setSlashingReceiver(slashing_receiver, {"from": owner})

    deposit_token_contract.send(slashed, 100, "", {"from": owner})
    deposit_token_contract.send(
        deposit_contract, 100, encode_withdrawal_delay(0), {"from": slashed}
    )

    tx = deposit_contract.slash(slashed, {"from": slasher})
    assert len(tx.events) == 3
    ev1, ev2, ev3 = tx.events

    assert ev1.name == "DepositChanged"
    assert ev1 == {
        "account": slashed,
        "amount": 0,
        "withdrawalDelayBlocks": 0,
        "withdrawalRequestedBlock": 0,
        "withdrawn": False,
        "slashed": True,
    }
    assert ev2.name == "Sent"
    assert ev2 == {
        "operator": deposit_contract,
        "from": deposit_contract,
        "to": slashing_receiver,
        "amount": 100,
        "data": "0x00",
        "operatorData": "0x00",
    }
    assert ev3.name == "Transfer"
    assert ev3 == {
        "from": deposit_contract,
        "to": slashing_receiver,
        "value": 100,
    }

    assert deposit_token_contract.balanceOf(slashing_receiver) == 100
