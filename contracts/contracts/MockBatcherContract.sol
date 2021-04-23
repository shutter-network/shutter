// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;
pragma experimental ABIEncoderV2;

contract MockBatcherContract {
    enum TransactionType {Cipher, Plain}

    mapping(uint64 => mapping(TransactionType => bytes32)) public batchHashes;

    function setBatchHash(
        uint64 batchIndex,
        TransactionType transactionType,
        bytes32 batchHash
    ) external {
        batchHashes[batchIndex][transactionType] = batchHash;
    }
}
