// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import {ConfigContract} from "./ConfigContract.sol";
import {
    BatcherContract,
    BatchConfig,
    TransactionType
} from "./BatcherContract.sol";

struct CipherExecutionReceipt {
    bool executed;
    address executor;
    uint64 halfStep;
    bytes32 cipherBatchHash;
    bytes32 batchHash;
}

/// @title A contract that serves as the entry point of batch execution
/// @dev Batch execution is carried out in two separate steps: Execution of the encrypted portion,
///     followed by execution of the plaintext portion. Thus, progress is counted in half steps (0
///     and 1 for batch 0, 2 and 3 for batch 1, and so on).
contract ExecutorContract {
    /// @notice The event emitted after a batch execution half step has been carried out.
    /// @param numExecutionHalfSteps The total number of finished execution half steps, including
    ///     the one responsible for emitting the event.
    /// @param batchHash The hash of the executed batch (consisting of plaintext transactions).
    event BatchExecuted(uint64 numExecutionHalfSteps, bytes32 batchHash);

    /// @notice The event emitted after execution of the cipher portion of a batch has been skipped.
    /// @param numExecutionHalfSteps The total number of finished execution half steps, including
    ///     this one.
    event CipherExecutionSkipped(uint64 numExecutionHalfSteps);

    event TransactionFailed(uint64 txIndex, bytes32 txHash, bytes data);

    ConfigContract public configContract;
    BatcherContract public batcherContract;

    uint64 public numExecutionHalfSteps;
    mapping(uint64 => CipherExecutionReceipt) public cipherExecutionReceipts;

    constructor(
        ConfigContract _configContract,
        BatcherContract _batcherContract
    ) {
        configContract = _configContract;
        batcherContract = _batcherContract;
    }

    /// @notice Execute the cipher portion of a batch.
    /// @param _batchIndex The index of the batch
    /// @param _cipherBatchHash The hash of the batch (consisting of encrypted transactions)
    /// @param _transactions The sequence of (decrypted) transactions to execute.
    /// @param _keyperIndex The index of the keyper calling the function.
    /// @notice Execution is only performed if `_cipherBatchHash` matches the hash in the batcher
    ///     contract and the batch is active and completed.
    function executeCipherBatch(
        uint64 _batchIndex,
        bytes32 _cipherBatchHash,
        bytes[] calldata _transactions,
        uint64 _keyperIndex
    ) external {
        require(
            numExecutionHalfSteps / 2 == _batchIndex,
            "ExecutorContract: unexpected batch index"
        );
        // Check that it's a cipher batch turn
        require(
            numExecutionHalfSteps % 2 == 0,
            "ExecutorContract: unexpected half step"
        );

        BatchConfig memory _config = configContract.getConfig(_batchIndex);

        // Check that batching is active and the batch is closed
        require(_config.batchSpan > 0, "ExecutorContract: config is inactive");

        // skip cipher execution if we reached the execution timeout.
        if (
            block.number >=
            _config.startBlockNumber +
                _config.batchSpan *
                (_batchIndex + 1) +
                _config.executionTimeout
        ) {
            numExecutionHalfSteps++;
            emit CipherExecutionSkipped(numExecutionHalfSteps);
            return;
        }
        require(
            block.number >=
                _config.startBlockNumber +
                    _config.batchSpan *
                    (_batchIndex + 1),
            "ExecutorContract: batch is not closed yet"
        );

        // Check that caller is keyper
        require(
            _keyperIndex < _config.keypers.length,
            "ExecutorContract: keyper index out of bounds"
        );
        require(
            msg.sender == _config.keypers[_keyperIndex],
            "ExecutorContract: sender is not specified keyper"
        );

        // Check the cipher batch hash is correct
        require(
            _cipherBatchHash ==
                batcherContract.batchHashes(
                    _batchIndex,
                    TransactionType.Cipher
                ),
            "ExecutorContract: incorrect cipher batch hash"
        );

        // Execute the batch
        bytes32 _batchHash =
            executeTransactions(
                _config.targetAddress,
                _config.targetFunctionSelector,
                _config.transactionGasLimit,
                _transactions
            );

        cipherExecutionReceipts[
            numExecutionHalfSteps
        ] = CipherExecutionReceipt({
            executed: true,
            executor: msg.sender,
            halfStep: numExecutionHalfSteps,
            cipherBatchHash: _cipherBatchHash,
            batchHash: _batchHash
        });
        numExecutionHalfSteps++;
        emit BatchExecuted(numExecutionHalfSteps, _batchHash);
    }

    /// @notice Skip execution of the cipher portion of a batch.
    /// @notice This is only possible if successful execution has not been carried out in time
    ///     (according to the execution timeout defined in the config)
    function skipCipherExecution(uint64 _batchIndex) external {
        require(
            numExecutionHalfSteps / 2 == _batchIndex,
            "ExecutorContract: unexpected batch index"
        );

        require(
            numExecutionHalfSteps % 2 == 0,
            "ExecutorContract: unexpected half step"
        );

        BatchConfig memory _config = configContract.getConfig(_batchIndex);

        require(_config.batchSpan > 0, "ExecutorContract: config is inactive");
        require(
            block.number >=
                _config.startBlockNumber +
                    _config.batchSpan *
                    (_batchIndex + 1) +
                    _config.executionTimeout,
            "ExecutorContract: execution timeout not reached yet"
        );

        numExecutionHalfSteps++;

        emit CipherExecutionSkipped(numExecutionHalfSteps);
    }

    /// @notice Execute the plaintext portion of a batch.
    /// @param _batchIndex The index of the batch
    /// @param _transactions The array of plaintext transactions in the batch.
    /// @notice This is a trustless operation since `_transactions` will be checked against the
    ///     (plaintext) batch hash from the batcher contract.
    function executePlainBatch(
        uint64 _batchIndex,
        bytes[] calldata _transactions
    ) external {
        require(
            numExecutionHalfSteps / 2 == _batchIndex,
            "ExecutorContract: unexpected batch index"
        );
        require(
            numExecutionHalfSteps % 2 == 1,
            "ExecutorContract: unexpected half step"
        );

        BatchConfig memory _config = configContract.getConfig(_batchIndex);

        // Since the cipher part of the batch has already been executed or skipped and the
        // config cannot be changed anymore (since the batching period is over), the following
        // checks remain true.
        assert(_config.batchSpan > 0);
        assert(
            block.number >=
                _config.startBlockNumber + _config.batchSpan * (_batchIndex + 1)
        );

        bytes32 _batchHash =
            executeTransactions(
                _config.targetAddress,
                _config.targetFunctionSelector,
                _config.transactionGasLimit,
                _transactions
            );

        require(
            _batchHash ==
                batcherContract.batchHashes(_batchIndex, TransactionType.Plain),
            "ExecutorContract: batch hash does not match"
        );

        numExecutionHalfSteps++;

        emit BatchExecuted(numExecutionHalfSteps, _batchHash);
    }

    function executeTransactions(
        address _targetAddress,
        bytes4 _targetFunctionSelector,
        uint64 _gasLimit,
        bytes[] calldata _transactions
    ) private returns (bytes32) {
        bytes32 _batchHash;
        for (uint64 _i = 0; _i < _transactions.length; _i++) {
            bytes memory _calldata =
                abi.encodeWithSelector(
                    _targetFunctionSelector,
                    _transactions[_i]
                );

            // call target function, ignoring any errors
            (bool success, bytes memory data) =
                _targetAddress.call{gas: _gasLimit}(_calldata);
            if (!success) {
                emit TransactionFailed({
                    txIndex: _i,
                    txHash: keccak256(_transactions[_i]),
                    data: data
                });
            }

            _batchHash = keccak256(
                abi.encodePacked(_transactions[_i], _batchHash)
            );
        }
        return _batchHash;
    }

    function getReceipt(uint64 _halfStep)
        public
        view
        returns (CipherExecutionReceipt memory)
    {
        return cipherExecutionReceipts[_halfStep];
    }
}
