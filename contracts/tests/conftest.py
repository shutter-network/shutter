from typing import Any
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
def config_contract(
    ConfigContract: ContractContainer, owner: Account, config_change_heads_up_blocks: int
) -> Any:
    config_contract = owner.deploy(ConfigContract, config_change_heads_up_blocks)
    return config_contract


@pytest.fixture
def batcher_contract(
    BatcherContract: ContractContainer, config_contract: Any, owner: Account
) -> Any:
    config_contract = owner.deploy(BatcherContract, config_contract)
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
def mock_target_contract(MockTargetContract: ContractContainer, owner: Account) -> Any:
    mock_target_contract = owner.deploy(MockTargetContract)
    return mock_target_contract


@pytest.fixture
def mock_batcher_contract(MockBatcherContract: ContractContainer, owner: Account) -> Any:
    mock_batcher_contract = owner.deploy(MockBatcherContract)
    return mock_batcher_contract


@pytest.fixture
def test_bls_contract(TestBLS: ContractContainer, accounts: Sequence[Account]) -> Any:
    test_bls_contract = accounts[0].deploy(TestBLS)
    return test_bls_contract


@pytest.fixture
def bls_registration_contract(
    BLSRegistrationContract: ContractContainer, accounts: Sequence[Account]
) -> Any:
    bls_registration_contract = accounts[0].deploy(BLSRegistrationContract)
    return bls_registration_contract


@pytest.fixture
def mock_target_function_selector(MockTargetContract: ContractContainer) -> bytes:
    function_name = "call"
    for selector, name in MockTargetContract.selectors.items():
        if name == function_name:
            return bytes(decode_hex(selector))
    raise AssertionError


@pytest.fixture(autouse=True)
def isolation(fn_isolation: Any) -> None:
    pass
