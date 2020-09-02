// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "./ConfigContract.sol";
import "./Ownable.sol";

enum TransactionType {Cipher, Plain}

contract BatcherContract is Ownable {
    event TransactionAdded(
        uint256 batchIndex,
        TransactionType transactionType,
        bytes transaction,
        bytes32 batchHash
    );

    ConfigContract public configContract;
    mapping(uint256 => uint256) public batchSizes;
    mapping(uint256 => mapping(TransactionType => bytes32)) public batchHashes;

    uint256 public minFee;

    constructor(address _configContractAddress) public {
        configContract = ConfigContract(_configContractAddress);
    }

    function addTransaction(
        uint256 _batchIndex,
        TransactionType _type,
        bytes calldata _transaction
    ) external payable {
        BatchConfig memory config = configContract.getConfig(_batchIndex);
        assert(block.number >= config.startBlockNumber);

        require(config.active);

        // TODO: over/underflows?
        uint256 _relativeBatchIndex = _batchIndex - config.startBatchIndex;
        require(
            block.number <
                config.startBlockNumber +
                    config.batchSpan *
                    (_relativeBatchIndex + 1)
        );

        require(_transaction.length > 0);
        require(_transaction.length <= config.transactionSizeLimit);

        require(
            batchSizes[_batchIndex] + _transaction.length <=
                config.batchSizeLimit
        );

        require(msg.value >= minFee);

        bytes memory _batchHashPreimage = abi.encodePacked(
            _transaction,
            batchHashes[_batchIndex][_type]
        );
        bytes32 _newBatchHash = keccak256(_batchHashPreimage);
        batchHashes[_batchIndex][_type] = _newBatchHash;
        batchSizes[_batchIndex] += _transaction.length;

        emit TransactionAdded(_batchIndex, _type, _transaction, _newBatchHash);
        (bool success, ) = config.feeReceiver.call{value: msg.value}(""); // TODO: check if reentrancy could be an issue
        require(success);
    }

    function setMinFee(uint256 _minFee) external onlyOwner {
        minFee = _minFee;
    }
}
