// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;
pragma experimental ABIEncoderV2;

import {Ownable} from "openzeppelin/contracts/access/Ownable.sol";
import {ConfigContract, BatchConfig} from "./ConfigContract.sol";
import {FeeBankContract} from "./FeeBankContract.sol";

enum TransactionType {Cipher, Plain}

/// @title A contract that batches transactions.
contract BatcherContract is Ownable {
    /// @notice The event emitted whenever a transaction is added to a batch.
    /// @param batchIndex The index of the batch to which the transaction has been added.
    /// @param transactionType The type of the transaction (cipher or plain).
    /// @param transaction The encrypted or plaintext transaction (depending on the type).
    /// @param batchHash The batch hash after adding the transaction.
    event TransactionAdded(
        uint64 batchIndex,
        TransactionType transactionType,
        bytes transaction,
        bytes32 batchHash
    );

    // The contract from which batch configs are fetched.
    ConfigContract public configContract;
    // The contract to which fees are sent.
    FeeBankContract public feeBankContract;

    // Stores the current size of the batches by batch index. Note that cipher and plain batches
    // are not tracked separately but in sum.
    mapping(uint64 => uint64) public batchSizes;
    // The current batch hashes by index and type (cipher or plain).
    mapping(uint64 => mapping(TransactionType => bytes32)) public batchHashes;

    // The minimum fee required to add a transaction to a batch.
    uint64 public minFee;

    constructor(
        ConfigContract configContractAddress,
        FeeBankContract feeBankContractAddress
    ) {
        configContract = configContractAddress;
        feeBankContract = feeBankContractAddress;
    }

    /// @notice Add a transaction to a batch.
    /// @param batchIndex The index of the batch to which the transaction should be added. Note
    ///     that this must match the batch corresponding to the current block number.
    /// @param transactionType The type of the transaction (either cipher or plain).
    /// @param transaction The encrypted or plaintext transaction (depending on `transactionType`).
    function addTransaction(
        uint64 batchIndex,
        TransactionType transactionType,
        bytes calldata transaction
    ) external payable {
        uint64 configIndex =
            configContract.configIndexForBatchIndex(batchIndex);
        uint64 transactionSizeLimit =
            configContract.configTransactionSizeLimit(configIndex);
        uint64 batchSizeLimit =
            configContract.configBatchSizeLimit(configIndex);
        address feeReceiver = configContract.configFeeReceiver(configIndex);

        // check batching is active and open
        require(
            configContract.batchingActive(configIndex),
            "BatcherContract: batch not active"
        );
        (uint64 start, uint64 end, ) =
            configContract.batchBoundaryBlocks(configIndex, batchIndex);
        require(block.number >= start, "BatcherContract: batch not open yet");
        require(block.number < end, "BatcherContract: batch already closed");

        // check tx and batch size limits
        require(
            transaction.length > 0,
            "BatcherContract: transaction is empty"
        );
        require(
            transaction.length <= transactionSizeLimit,
            "BatcherContract: transaction too big"
        );
        require(
            batchSizes[batchIndex] + transaction.length <= batchSizeLimit,
            "BatcherContract: batch already full"
        ); // overflow can be ignored here because number of txs and their sizes are both small

        // check fee
        require(msg.value >= minFee, "BatcherContract: fee too small");

        // add tx to batch
        bytes memory batchHashPreimage =
            abi.encodePacked(
                transaction,
                batchHashes[batchIndex][transactionType]
            );
        bytes32 newBatchHash = keccak256(batchHashPreimage);
        batchHashes[batchIndex][transactionType] = newBatchHash;
        batchSizes[batchIndex] += uint64(transaction.length);

        // pay fee to fee bank and emit event
        if (msg.value > 0 && feeReceiver != address(0)) {
            feeBankContract.deposit{value: msg.value}(feeReceiver);
        }
        emit TransactionAdded(
            batchIndex,
            transactionType,
            transaction,
            newBatchHash
        );
    }

    /// @notice Set the minimum fee required to add a transaction to the batch.
    /// @param newMinFee The new value for the minimum fee.
    function setMinFee(uint64 newMinFee) external onlyOwner {
        minFee = newMinFee;
    }
}
