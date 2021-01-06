// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/access/Ownable.sol";
import "./ConfigContract.sol";

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
