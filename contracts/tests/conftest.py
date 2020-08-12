from typing import Any
from typing import Sequence

import pytest
from brownie.network.account import Account
from brownie.network.contract import ContractContainer


@pytest.fixture
def config_change_heads_up_blocks() -> int:
    return 5


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


@pytest.fixture(autouse=True)
def isolation(fn_isolation: Any) -> None:
    pass
