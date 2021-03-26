// SPDX-License-Identifier: MIT

pragma solidity =0.7.6;
pragma experimental ABIEncoderV2;

contract MockBatcherContract {
    enum TransactionType {Cipher, Plain}

    mapping(uint64 => mapping(TransactionType => bytes32)) public batchHashes;

    function setBatchHash(
        uint64 _batchIndex,
        TransactionType _type,
        bytes32 _hash
    ) external {
        batchHashes[_batchIndex][_type] = _hash;
    }
}
