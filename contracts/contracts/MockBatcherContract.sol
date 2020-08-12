// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "./ConfigContract.sol";
import "./Ownable.sol";

contract MockBatcherContract {
    enum TransactionType {Cipher, Plain}

    mapping(uint256 => mapping(TransactionType => bytes32)) public batchHashes;

    function setBatchHash(
        uint256 batchIndex,
        TransactionType _type,
        bytes32 hash
    ) external {
        batchHashes[batchIndex][_type] = hash;
    }
}
