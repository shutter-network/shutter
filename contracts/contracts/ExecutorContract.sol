// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "./ConfigContract.sol";
import "./BatcherContract.sol";

contract ExecutorContract {
    event BatchExecuted(uint256 numExecutionHalfSteps, bytes32 batchHash);

    ConfigContract public configContract;
    BatcherContract public batcherContract;

    uint256 public numExecutionHalfSteps;

    constructor(
        ConfigContract _configContract,
        BatcherContract _batcherContract
    ) public {
        configContract = _configContract;
        batcherContract = _batcherContract;
    }

    function executeCipherBatch(
        bytes32 _cipherBatchHash,
        bytes[] calldata _transactions,
        bytes32 _decryptionKey,
        bytes32 _aggregatedSignature
    ) external {
        require(numExecutionHalfSteps % 2 == 0);

        uint256 _batchIndex = numExecutionHalfSteps / 2;
        BatchConfig memory _config = configContract.getConfig(_batchIndex);

        require(
            _cipherBatchHash ==
                batcherContract.batchHashes(_batchIndex, TransactionType.Cipher)
        );

        require(_config.active);
        require(
            block.number >=
                _config.startBlockNumber + _config.batchSpan * (_batchIndex + 1)
        );

        bytes32 _batchHash = executeTransactions(
            _config.targetAddress,
            _config.targetFunctionSelector,
            _config.transactionGasLimit,
            _transactions
        );

        numExecutionHalfSteps++;

        emit BatchExecuted(numExecutionHalfSteps, _batchHash);
    }

    function executePlainBatch(bytes[] calldata _transactions) external {
        require(numExecutionHalfSteps % 2 == 1);

        uint256 _batchIndex = numExecutionHalfSteps / 2;
        BatchConfig memory _config = configContract.getConfig(_batchIndex);

        // Since the cipher part of the batch has already been executed successfully, and the
        // config cannot be changed anymore (since the batching period is over), the following
        // checks remain true. (TODO: double check)
        assert(_config.active);
        assert(
            block.number >=
                _config.startBlockNumber + _config.batchSpan * (_batchIndex + 1)
        );

        bytes32 _batchHash = executeTransactions(
            _config.targetAddress,
            _config.targetFunctionSelector,
            _config.transactionGasLimit,
            _transactions
        );

        require(
            _batchHash ==
                batcherContract.batchHashes(_batchIndex, TransactionType.Plain)
        );

        numExecutionHalfSteps++;

        emit BatchExecuted(numExecutionHalfSteps, _batchHash);
    }

    function executeTransactions(
        address _targetAddress,
        bytes4 _targetFunctionSelector,
        uint256 _gasLimit,
        bytes[] calldata _transactions
    ) private returns (bytes32) {
        bytes32 _batchHash;
        for (uint256 i = 0; i < _transactions.length; i++) {
            bytes memory _calldata = abi.encodeWithSelector(
                _targetFunctionSelector,
                _transactions[i]
            );
            _targetAddress.call{gas: _gasLimit}(_calldata);

            _batchHash = keccak256(
                abi.encodePacked(_transactions[i], _batchHash)
            );
        }
        return _batchHash;
    }
}
