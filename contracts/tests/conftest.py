from typing import Any
from typing import List
from typing import Sequence

import pytest
from brownie.network.account import Account
from brownie.network.contract import ContractContainer
from eth_utils import decode_hex


@pytest.fixture
def config_change_heads_up_blocks() -> int:
    return 200


@pytest.fixture
def owner(accounts: Sequence[Account]) -> Account:
    return accounts[0]


@pytest.fixture
def non_owner(accounts: Sequence[Account], owner: Account) -> Account:
    non_owner = accounts[1]
    assert non_owner != owner
    return non_owner


@pytest.fixture
def keypers(accounts: Sequence[Account]) -> List[Account]:
    # as opposed to existing accounts, added ones have private keys known private keys which we
    # need to sign things
    return [accounts.add() for _ in range(3)]  # type: ignore


@pytest.fixture
def keyper_private_keys(keypers: Sequence[Account]) -> List[bytes]:
    return [decode_hex(keyper.private_key) for keyper in keypers]


@pytest.fixture
def appeal_blocks() -> int:
    return 10


@pytest.fixture
def config_contract(
    ConfigContract: ContractContainer, owner: Account, config_change_heads_up_blocks: int
) -> Any:
    config_contract = owner.deploy(ConfigContract, config_change_heads_up_blocks)
    return config_contract


@pytest.fixture
def batcher_contract(
    BatcherContract: ContractContainer,
    config_contract: Any,
    fee_bank_contract: Any,
    owner: Account,
) -> Any:
    config_contract = owner.deploy(BatcherContract, config_contract, fee_bank_contract)
    return config_contract


@pytest.fixture
def executor_contract(
    ExecutorContract: ContractContainer,
    config_contract: Any,
    mock_batcher_contract: Any,
    owner: Account,
) -> Any:
    executor_contract = owner.deploy(ExecutorContract, config_contract, mock_batcher_contract)
    return executor_contract


@pytest.fixture
def keyper_slasher(
    KeyperSlasher: ContractContainer,
    config_contract: Any,
    executor_contract: Any,
    owner: Account,
    appeal_blocks: int,
) -> Any:
    keyper_slasher = owner.deploy(KeyperSlasher, appeal_blocks, config_contract, executor_contract)
    return keyper_slasher


@pytest.fixture
def mock_target_contract(MockTargetContract: ContractContainer, owner: Account) -> Any:
    mock_target_contract = owner.deploy(MockTargetContract)
    return mock_target_contract


@pytest.fixture
def mock_batcher_contract(MockBatcherContract: ContractContainer, owner: Account) -> Any:
    mock_batcher_contract = owner.deploy(MockBatcherContract)
    return mock_batcher_contract


@pytest.fixture
def mock_target_function_selector(MockTargetContract: ContractContainer) -> bytes:
    function_name = "call"
    for selector, name in MockTargetContract.selectors.items():
        if name == function_name:
            return bytes(decode_hex(selector))
    raise AssertionError


@pytest.fixture
def fee_bank_contract(FeeBankContract: ContractContainer, accounts: Sequence[Account]) -> Any:
    fee_bank_contract = accounts[0].deploy(FeeBankContract)
    return fee_bank_contract


@pytest.fixture
def key_broadcast_contract(
    KeyBroadcastContract: ContractContainer, config_contract: Any, accounts: Sequence[Account]
) -> Any:
    key_broadcast_contract = accounts[0].deploy(KeyBroadcastContract, config_contract.address)
    return key_broadcast_contract


@pytest.fixture(autouse=True)
def isolation(fn_isolation: Any) -> None:
    pass
