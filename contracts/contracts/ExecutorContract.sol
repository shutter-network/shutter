// SPDX-License-Identifier: MIT

pragma solidity =0.7.6;
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
        ConfigContract configContractAddress,
        BatcherContract batcherContractAddress
    ) {
        configContract = configContractAddress;
        batcherContract = batcherContractAddress;
    }

    /// @notice Execute the cipher portion of a batch.
    /// @param batchIndex The index of the batch
    /// @param cipherBatchHash The hash of the batch (consisting of encrypted transactions)
    /// @param transactions The sequence of (decrypted) transactions to execute.
    /// @param keyperIndex The index of the keyper calling the function.
    /// @notice Execution is only performed if `cipherBatchHash` matches the hash in the batcher
    ///     contract and the batch is active and completed.
    function executeCipherBatch(
        uint64 batchIndex,
        bytes32 cipherBatchHash,
        bytes[] calldata transactions,
        uint64 keyperIndex
    ) external {
        require(
            numExecutionHalfSteps / 2 == batchIndex,
            "ExecutorContract: unexpected batch index"
        );
        // Check that it's a cipher batch turn
        require(
            numExecutionHalfSteps % 2 == 0,
            "ExecutorContract: unexpected half step"
        );

        BatchConfig memory config = configContract.getConfig(batchIndex);

        // Check that batching is active and the batch is closed
        require(config.batchSpan > 0, "ExecutorContract: config is inactive");

        // skip cipher execution if we reached the execution timeout.
        if (
            block.number >=
            config.startBlockNumber +
                config.batchSpan *
                (batchIndex + 1) +
                config.executionTimeout
        ) {
            numExecutionHalfSteps++;
            emit CipherExecutionSkipped(numExecutionHalfSteps);
            return;
        }
        require(
            block.number >=
                config.startBlockNumber + config.batchSpan * (batchIndex + 1),
            "ExecutorContract: batch is not closed yet"
        );

        // Check that caller is keyper
        require(
            keyperIndex < config.keypers.length,
            "ExecutorContract: keyper index out of bounds"
        );
        require(
            msg.sender == config.keypers[keyperIndex],
            "ExecutorContract: sender is not specified keyper"
        );

        // Check the cipher batch hash is correct
        require(
            cipherBatchHash ==
                batcherContract.batchHashes(batchIndex, TransactionType.Cipher),
            "ExecutorContract: incorrect cipher batch hash"
        );

        // Check the number of transactions is zero iff we provide the ZERO_HASH
        require(
            (cipherBatchHash == bytes32(0) && transactions.length == 0) ||
                (cipherBatchHash != bytes32(0) && transactions.length > 0),
            "ExecutorContract: cipherBatchHash should be zero iff transactions is empty"
        );

        // Execute the batch
        bytes32 batchHash =
            executeTransactions(
                config.targetAddress,
                config.targetFunctionSelector,
                config.transactionGasLimit,
                transactions
            );

        cipherExecutionReceipts[
            numExecutionHalfSteps
        ] = CipherExecutionReceipt({
            executed: true,
            executor: msg.sender,
            halfStep: numExecutionHalfSteps,
            cipherBatchHash: cipherBatchHash,
            batchHash: batchHash
        });
        numExecutionHalfSteps++;
        emit BatchExecuted(numExecutionHalfSteps, batchHash);
    }

    /// @notice Skip execution of the cipher portion of a batch.
    /// @notice This is only possible if successful execution has not been carried out in time
    ///     (according to the execution timeout defined in the config)
    function skipCipherExecution(uint64 batchIndex) external {
        require(
            numExecutionHalfSteps / 2 == batchIndex,
            "ExecutorContract: unexpected batch index"
        );

        require(
            numExecutionHalfSteps % 2 == 0,
            "ExecutorContract: unexpected half step"
        );

        BatchConfig memory config = configContract.getConfig(batchIndex);

        require(config.batchSpan > 0, "ExecutorContract: config is inactive");
        require(
            block.number >=
                config.startBlockNumber +
                    config.batchSpan *
                    (batchIndex + 1) +
                    config.executionTimeout,
            "ExecutorContract: execution timeout not reached yet"
        );

        numExecutionHalfSteps++;

        emit CipherExecutionSkipped(numExecutionHalfSteps);
    }

    /// @notice Execute the plaintext portion of a batch.
    /// @param batchIndex The index of the batch
    /// @param transactions The array of plaintext transactions in the batch.
    /// @notice This is a trustless operation since `transactions` will be checked against the
    ///     (plaintext) batch hash from the batcher contract.
    function executePlainBatch(uint64 batchIndex, bytes[] calldata transactions)
        external
    {
        require(
            numExecutionHalfSteps / 2 == batchIndex,
            "ExecutorContract: unexpected batch index"
        );
        require(
            numExecutionHalfSteps % 2 == 1,
            "ExecutorContract: unexpected half step"
        );

        BatchConfig memory config = configContract.getConfig(batchIndex);

        // Since the cipher part of the batch has already been executed or skipped and the
        // config cannot be changed anymore (since the batching period is over), the following
        // checks remain true.
        assert(config.batchSpan > 0);
        assert(
            block.number >=
                config.startBlockNumber + config.batchSpan * (batchIndex + 1)
        );

        bytes32 batchHash =
            executeTransactions(
                config.targetAddress,
                config.targetFunctionSelector,
                config.transactionGasLimit,
                transactions
            );

        require(
            batchHash ==
                batcherContract.batchHashes(batchIndex, TransactionType.Plain),
            "ExecutorContract: batch hash does not match"
        );

        numExecutionHalfSteps++;

        emit BatchExecuted(numExecutionHalfSteps, batchHash);
    }

    function executeTransactions(
        address targetAddress,
        bytes4 targetFunctionSelector,
        uint64 gasLimit,
        bytes[] calldata transactions
    ) private returns (bytes32) {
        bytes32 batchHash;
        for (uint64 i = 0; i < transactions.length; i++) {
            bytes memory callData =
                abi.encodeWithSelector(targetFunctionSelector, transactions[i]);

            // call target function, ignoring any errors
            (bool success, bytes memory returnData) =
                targetAddress.call{gas: gasLimit}(callData);
            if (!success) {
                emit TransactionFailed({
                    txIndex: i,
                    txHash: keccak256(transactions[i]),
                    data: returnData
                });
            }

            batchHash = keccak256(abi.encodePacked(transactions[i], batchHash));
        }
        return batchHash;
    }

    function getReceipt(uint64 halfStep)
        public
        view
        returns (CipherExecutionReceipt memory)
    {
        return cipherExecutionReceipts[halfStep];
    }
}
