import random
from typing import List
from typing import Optional

from eth_typing import Address

from tests.contract_helpers import BatchConfig


def make_int(min_value: int = 0, max_value: int = 2 ** 256) -> int:
    return random.randint(min_value, max_value)


def make_bytes(length: Optional[int] = None) -> bytes:
    if length is None:
        length = random.randint(4, 10)
    return bytes(random.randint(0, 255) for _ in range(length))


def make_ecdsa_private_key() -> bytes:
    return make_bytes(32)


def make_address() -> Address:
    return Address(make_bytes(20))


def make_batch_config(
    *,
    start_batch_index: Optional[int] = None,
    start_block_number: Optional[int] = None,
    keypers: Optional[List[Address]] = None,
    threshold: Optional[int] = None,
    batch_span: int = 0,
    batch_size_limit: Optional[int] = None,
    transaction_size_limit: Optional[int] = None,
    transaction_gas_limit: Optional[int] = None,
    fee_receiver: Optional[Address] = None,
    target_address: Optional[Address] = None,
    target_function_selector: Optional[bytes] = None,
    execution_timeout: Optional[int] = None,
) -> BatchConfig:
    if threshold is None:
        if keypers is None:
            threshold = 2
        else:
            threshold = max(1, len(keypers) // 3 * 2)

    if keypers is None:
        num_keypers = threshold // 2 * 3
        keypers = [make_address() for _ in range(num_keypers)]

    return BatchConfig(
        start_batch_index=start_batch_index if start_batch_index is not None else make_int(),
        start_block_number=start_block_number if start_block_number is not None else make_int(),
        keypers=keypers,
        threshold=threshold,
        batch_span=batch_span,
        batch_size_limit=batch_size_limit if batch_size_limit is not None else make_int(),
        transaction_size_limit=transaction_size_limit
        if transaction_size_limit is not None
        else make_int(),
        transaction_gas_limit=transaction_gas_limit
        if transaction_gas_limit is not None
        else make_int(),
        fee_receiver=fee_receiver if fee_receiver else make_address(),
        target_address=target_address if target_address else make_address(),
        target_function_selector=target_function_selector
        if target_function_selector is not None
        else make_bytes(4),
        execution_timeout=execution_timeout if execution_timeout is not None else make_int(),
    )


def make_batch(length: Optional[int] = None) -> List[bytes]:
    if length is None:
        length = make_int(max_value=3)

    return [make_bytes() for _ in range(length)]


def make_signer_indices(num_keypers: int, num_signers: int) -> List[int]:
    assert num_signers <= num_keypers
    keyper_indices = list(range(num_keypers))
    random.shuffle(keyper_indices)
    return sorted(keyper_indices[:num_signers])
