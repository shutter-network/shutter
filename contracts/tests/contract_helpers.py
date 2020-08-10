from __future__ import annotations

from typing import List
from typing import Sequence
from typing import Tuple

import attr
from eth_typing import Address
from eth_utils import decode_hex
from eth_utils import to_canonical_address

ZERO_ADDRESS = Address(b"\x00" * 20)


@attr.s(auto_attribs=True, frozen=True)
class BatchConfig:
    start_batch_index: int
    start_block_number: int
    active: bool
    keypers: List[Address]
    threshold: int
    batch_span: int
    batch_size_limit: int
    transaction_size_limit: int
    transaction_gas_limit: int
    fee_receiver: Address
    target_address: Address
    target_function: bytes
    execution_timeout: int

    @classmethod
    def from_tuple(cls, t: Tuple, keypers: Sequence[Address]) -> BatchConfig:
        assert len(t) == 12

        return cls(
            start_batch_index=t[0],
            start_block_number=t[1],
            active=t[2],
            keypers=keypers,
            threshold=t[3],
            batch_span=t[4],
            batch_size_limit=t[5],
            transaction_size_limit=t[6],
            transaction_gas_limit=t[7],
            fee_receiver=to_canonical_address(t[8]),
            target_address=to_canonical_address(t[9]),
            target_function=decode_hex(str(t[10])),
            execution_timeout=t[11],
        )


ZERO_CONFIG = BatchConfig(
    start_batch_index=0,
    start_block_number=0,
    active=False,
    keypers=[],
    threshold=0,
    batch_span=0,
    batch_size_limit=0,
    transaction_size_limit=0,
    transaction_gas_limit=0,
    fee_receiver=ZERO_ADDRESS,
    target_address=ZERO_ADDRESS,
    target_function=b"",
    execution_timeout=0,
)


def fetch_config(config_contract, config_index):
    config_tuple = config_contract.configs(config_index)
    config_num_keypers = config_contract.configNumKeypers(config_index)
    config_keypers = []
    for keyper_index in range(config_num_keypers):
        keyper = config_contract.configKeypers(config_index, keyper_index)
        config_keypers.append(to_canonical_address(keyper))

    return BatchConfig.from_tuple(config_tuple, config_keypers)


def fetch_next_config(config_contract):
    next_config_tuple = config_contract.nextConfig()
    next_config_num_keypers = config_contract.nextConfigNumKeypers()
    next_config_keypers = []
    for keyper_index in range(next_config_num_keypers):
        keyper = config_contract.nextConfigKeypers(keyper_index)
        next_config_keypers.append(to_canonical_address(keyper))

    return BatchConfig.from_tuple(next_config_tuple, next_config_keypers)


def set_next_config(config_contract, config, owner):
    for field in attr.fields(BatchConfig):
        name_snake_case = field.name
        name_camel_case = snake_to_camel_case(name_snake_case, capitalize=True)
        if name_snake_case == "keypers":
            continue

        setter_function_name = "nextConfigSet" + name_camel_case
        setter_function = getattr(config_contract, setter_function_name)
        value = getattr(config, name_snake_case)
        setter_function(value, {"from": owner})

    num_existing_keypers = config_contract.nextConfigNumKeypers()
    config_contract.nextConfigRemoveKeypers(num_existing_keypers, {"from": owner})
    config_contract.nextConfigAddKeypers(config.keypers, {"from": owner})


def schedule_config(config_contract, config, owner):
    set_next_config(config_contract, config, owner=owner)
    tx = config_contract.scheduleNextConfig({"from": owner})
    return tx


def snake_to_camel_case(snake_case_string: str, capitalize: bool) -> str:
    if len(snake_case_string) == 0:
        return ""
    parts = snake_case_string.split("_")

    if capitalize:
        first_part = parts[0].capitalize()
    else:
        first_part = parts[0]

    rest_parts = [part.capitalize() for part in parts[1:]]

    return "".join([first_part] + rest_parts)
