from typing import Any

import brownie
from brownie.network.account import Account
from web3 import Web3

from tests.contract_helpers import fetch_config
from tests.contract_helpers import fetch_config_by_index
from tests.contract_helpers import fetch_next_config
from tests.contract_helpers import schedule_config
from tests.contract_helpers import set_next_config
from tests.contract_helpers import ZERO_CONFIG
from tests.factories import make_address
from tests.factories import make_batch_config


def test_constructor_adds_guard_config(config_contract: Any) -> None:
    assert config_contract.numConfigs() == 1
    guard_config = fetch_config_by_index(config_contract, 0)
    assert guard_config == ZERO_CONFIG


def test_constructor_sets_heads_up(
    config_contract: Any, config_change_heads_up_blocks: int
) -> None:
    assert config_contract.configChangeHeadsUpBlocks() == config_change_heads_up_blocks


def test_set_next_config_fields(config_contract: Any, owner: Account) -> None:
    batch_config_template = make_batch_config(batch_span=1)
    set_next_config(config_contract, batch_config_template, owner=owner)
    next_config = fetch_next_config(config_contract)
    assert next_config == batch_config_template


def test_num_configs_returns_number_of_configs(config_contract: Any, owner: Account) -> None:
    assert config_contract.numConfigs() == 1
    schedule_config(
        config_contract,
        make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1),
        owner=owner,
    )
    assert config_contract.numConfigs() == 2
    schedule_config(
        config_contract,
        make_batch_config(start_batch_index=100, start_block_number=600, batch_span=1),
        owner=owner,
    )
    assert config_contract.numConfigs() == 3


def test_scheduling_adds_new_config(config_contract: Any, owner: Account) -> None:
    batch_config = make_batch_config(start_batch_index=0, batch_span=1)
    set_next_config(config_contract, batch_config, owner=owner)

    assert config_contract.numConfigs() == 1
    config_contract.scheduleNextConfig({"from": owner})
    assert config_contract.numConfigs() == 2
    assert fetch_config_by_index(config_contract, 1) == batch_config


def test_scheduling_checks_seamlessness_after_active_config(
    config_contract: Any, owner: Account
) -> None:
    batch_config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=10)
    set_next_config(config_contract, batch_config, owner=owner)
    config_contract.scheduleNextConfig({"from": owner})

    for start_block_number in [598, 599, 601, 602]:
        next_batch_config = make_batch_config(
            start_batch_index=10, start_block_number=start_block_number, batch_span=5
        )
        set_next_config(config_contract, next_batch_config, owner=owner)
        with brownie.reverts():
            config_contract.scheduleNextConfig({"from": owner})

    next_batch_config = make_batch_config(
        start_batch_index=10, start_block_number=600, batch_span=5
    )
    set_next_config(config_contract, next_batch_config, owner=owner)
    config_contract.scheduleNextConfig({"from": owner})

    assert config_contract.numConfigs() == 3
    assert fetch_config_by_index(config_contract, 2) == next_batch_config


def test_scheduling_checks_seamlessness_after_inactive_config(
    config_contract: Any, owner: Account
) -> None:
    config = make_batch_config(
        start_batch_index=1, start_block_number=500, active=True, batch_span=10
    )
    set_next_config(config_contract, config, owner=owner)
    with brownie.reverts():
        config_contract.scheduleNextConfig({"from": owner})

    config = make_batch_config(
        start_batch_index=0, start_block_number=500, active=True, batch_span=10
    )
    set_next_config(config_contract, config, owner=owner)
    config_contract.scheduleNextConfig({"from": owner})

    assert config_contract.numConfigs() == 2
    assert fetch_config_by_index(config_contract, 1) == config


def test_scheduling_checks_batch_span(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(
        start_batch_index=0, start_block_number=500, active=True, batch_span=0
    )
    set_next_config(config_contract, config, owner=owner)
    with brownie.reverts():
        config_contract.scheduleNextConfig({"from": owner})

    config = make_batch_config(
        start_batch_index=0, start_block_number=500, active=False, batch_span=1
    )
    set_next_config(config_contract, config, owner=owner)
    with brownie.reverts():
        config_contract.scheduleNextConfig({"from": owner})


def test_scheduling_resets_new_config(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    schedule_config(config_contract, config, owner=owner)
    next_config = fetch_next_config(config_contract)
    assert next_config == ZERO_CONFIG


def test_only_owner_can_set_next_config(
    config_contract: Any, owner: Account, non_owner: Account
) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    with brownie.reverts():
        set_next_config(config_contract, config, owner=non_owner)
    set_next_config(config_contract, config, owner=owner)


def test_only_owner_can_schedule_config(
    config_contract: Any, owner: Account, non_owner: Account
) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    set_next_config(config_contract, config, owner=owner)
    with brownie.reverts():
        config_contract.scheduleNextConfig({"from": non_owner})
    config_contract.scheduleNextConfig({"from": owner})


def test_scheduling_must_happen_with_heads_up(
    web3: Web3, config_contract: Any, config_change_heads_up_blocks: int, owner: Account
) -> None:
    config = make_batch_config(start_batch_index=0, batch_span=1)
    set_next_config(config_contract, config, owner=owner)

    block_number = web3.eth.blockNumber
    block_number_at_schedule = block_number + 2
    config_contract.nextConfigSetStartBlockNumber(
        block_number_at_schedule + config_change_heads_up_blocks
    )
    with brownie.reverts():
        config_contract.scheduleNextConfig({"from": owner})

    block_number = web3.eth.blockNumber
    block_number_at_schedule = block_number + 2
    config_contract.nextConfigSetStartBlockNumber(
        block_number_at_schedule + config_change_heads_up_blocks + 1
    )
    config_contract.scheduleNextConfig({"from": owner})


def test_scheduling_emits_event(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    tx = schedule_config(config_contract, config, owner=owner)
    assert len(tx.events) == 1
    assert tx.events["ConfigScheduled"] == {"numConfigs": 2}


def test_remove_some_keypers_from_next_config(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(
        keypers=[make_address() for _ in range(4)], batch_span=1, start_block_number=500
    )
    set_next_config(config_contract, config, owner=owner)
    assert config_contract.nextConfigNumKeypers() == 4
    config_contract.nextConfigRemoveKeypers(0)
    assert config_contract.nextConfigNumKeypers() == 4
    config_contract.nextConfigRemoveKeypers(1)
    assert config_contract.nextConfigNumKeypers() == 3
    config_contract.nextConfigRemoveKeypers(2)
    assert config_contract.nextConfigNumKeypers() == 1


def test_remove_all_keypers_from_next_config(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(
        keypers=[make_address() for _ in range(4)], batch_span=1, start_block_number=500
    )
    set_next_config(config_contract, config, owner=owner)
    config_contract.nextConfigRemoveKeypers(5)
    assert config_contract.nextConfigNumKeypers() == 0


def test_unschedule_single_config(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    schedule_config(config_contract, config, owner=owner)

    assert config_contract.numConfigs() == 2
    config_contract.unscheduleConfigs(400, {"from": owner})
    assert config_contract.numConfigs() == 1


def test_unschedule_multiple_configs(config_contract: Any, owner: Account) -> None:
    for start_batch_index in [0, 100, 200]:
        start_block_number = start_batch_index + 500
        config = make_batch_config(
            start_batch_index=start_batch_index,
            start_block_number=start_block_number,
            batch_span=1,
        )
        schedule_config(config_contract, config, owner=owner)

    assert config_contract.numConfigs() == 4
    config_contract.unscheduleConfigs(501, {"from": owner})
    assert config_contract.numConfigs() == 2


def test_unscheduling_nothing(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    schedule_config(config_contract, config, owner=owner)

    with brownie.reverts():
        config_contract.unscheduleConfigs(501, {"from": owner})
    config_contract.unscheduleConfigs(500, {"from": owner})


def test_only_owner_can_unschedule(
    config_contract: Any, owner: Account, non_owner: Account
) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    schedule_config(config_contract, config, owner=owner)

    with brownie.reverts():
        config_contract.unscheduleConfigs(400, {"from": non_owner})
    config_contract.unscheduleConfigs(400, {"from": owner})


def test_unscheduling_must_happen_with_heads_up(
    config_contract: Any, owner: Account, web3: Web3, config_change_heads_up_blocks: int
) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    schedule_config(config_contract, config, owner=owner)

    block_number = web3.eth.blockNumber
    with brownie.reverts():
        config_contract.unscheduleConfigs(block_number + config_change_heads_up_blocks)
    config_contract.unscheduleConfigs(web3.eth.blockNumber + config_change_heads_up_blocks + 2)


def test_unscheduling_emits_event(config_contract: Any, owner: Account) -> None:
    config = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=1)
    schedule_config(config_contract, config, owner=owner)
    tx = config_contract.unscheduleConfigs(400, {"from": owner})
    assert len(tx.events) == 1
    assert tx.events["ConfigUnscheduled"] == {"numConfigs": 1}


def test_get_guard_config(config_contract: Any) -> None:
    config = fetch_config(config_contract, 0)
    assert config == ZERO_CONFIG
    assert fetch_config(config_contract, 123) == config


def test_get_active_config(config_contract: Any, owner: Account) -> None:
    config1 = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=2)
    config2 = make_batch_config(start_batch_index=5, start_block_number=510, batch_span=10)
    schedule_config(config_contract, config1, owner=owner)
    schedule_config(config_contract, config2, owner=owner)

    for batch_index in range(5):
        assert fetch_config(config_contract, batch_index) == config1

    for batch_index in range(5, 10):
        assert fetch_config(config_contract, batch_index) == config2


def test_get_inactive_config(config_contract: Any, owner: Account) -> None:
    config1 = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=2)
    config2 = make_batch_config(start_batch_index=5, start_block_number=510, active=False)
    schedule_config(config_contract, config1, owner=owner)
    schedule_config(config_contract, config2, owner=owner)

    for batch_index in range(5):
        assert fetch_config(config_contract, batch_index) == config1

    for batch_index in range(5, 10):
        assert fetch_config(config_contract, batch_index) == config2


def test_get_active_after_inactive_config(config_contract: Any, owner: Account) -> None:
    config1 = make_batch_config(start_batch_index=0, start_block_number=500, batch_span=2)
    config2 = make_batch_config(start_batch_index=5, start_block_number=510, active=False)
    config3 = make_batch_config(
        start_batch_index=5, start_block_number=520, active=True, batch_span=10
    )
    schedule_config(config_contract, config1, owner=owner)
    schedule_config(config_contract, config2, owner=owner)
    schedule_config(config_contract, config3, owner=owner)

    for batch_index in range(5):
        assert fetch_config(config_contract, batch_index) == config1

    for batch_index in range(5, 10):
        assert fetch_config(config_contract, batch_index) == config3
