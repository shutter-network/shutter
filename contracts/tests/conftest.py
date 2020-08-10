from typing import Any
from typing import Sequence

from brownie import ContractContainer
from eth_typing import Address
import pytest


@pytest.fixture
def config_change_heads_up_blocks() -> int:
    return 5


@pytest.fixture
def owner(accounts: Address) -> Address:
    return accounts[0]


@pytest.fixture
def non_owner(accounts: Sequence[Address], owner: Address) -> Address:
    non_owner = accounts[1]
    assert non_owner != owner
    return non_owner


@pytest.fixture
def config_contract(
    ConfigContract: ContractContainer, owner: Address, config_change_heads_up_blocks: int
) -> Any:
    config_contract = owner.deploy(ConfigContract, config_change_heads_up_blocks)
    return config_contract


@pytest.fixture(autouse=True)
def isolation(fn_isolation: Any) -> None:
    pass
