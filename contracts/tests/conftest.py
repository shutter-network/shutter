import pytest


@pytest.fixture
def config_change_heads_up_blocks():
    return 5


@pytest.fixture
def owner(accounts):
    return accounts[0]


@pytest.fixture
def non_owner(accounts, owner):
    non_owner = accounts[1]
    assert non_owner != owner


@pytest.fixture
def config_contract(ConfigContract, owner, config_change_heads_up_blocks):
    config_contract = owner.deploy(ConfigContract, config_change_heads_up_blocks)
    return config_contract


@pytest.fixture(autouse=True)
def isolation(fn_isolation):
    pass
