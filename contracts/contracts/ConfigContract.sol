// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;
pragma experimental ABIEncoderV2;

import {
    Ownable
} from "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/access/Ownable.sol";

struct BatchConfig {
    uint64 startBatchIndex; // the index of the first batch using this config
    uint64 startBlockNumber; // the block number from which on this config is applicable
    address[] keypers; // the keyper addresses
    uint64 threshold; // the threshold parameter
    uint64 batchSpan; // the duration of one batch in blocks
    uint64 batchSizeLimit; // the maximum size of a batch in bytes
    uint64 transactionSizeLimit; // the maximum size of each transaction in the batch in bytes
    uint64 transactionGasLimit; // the maximum amount of gas each transaction may use
    address feeReceiver; // the address receiving the collected fees
    address targetAddress; // the address of the contract responsible of executing transactions
    bytes4 targetFunctionSelector; // function of the target contract that executes transactions
    uint64 executionTimeout; // the number of blocks after which execution can be skipped
}

/// @title A contract that manages `BatchConfig` objects.
/// @dev The config objects are stored in sequence, with configs applicable to later batches being
///     lined up behind configs applicable to earlier batches (according to
///     `config.startBlockNumber`). The contract owner is entitled to add or remove configs at the
///     end at will as long as a notice of at least `configChangeHeadsUpBlocks` is given.
/// @dev To add a new config, first populate the `nextConfig` object accordingly and then schedule
///     it with `scheduleNextConfig`.
contract ConfigContract is Ownable {
    /// @notice The event emitted after a new config object has been scheduled.
    /// @param numConfigs The new number of configs stored.
    event ConfigScheduled(uint64 numConfigs);

    /// @notice The event emitted after the owner has unscheduled one or more config objects.
    /// @param numConfigs The new number of configs stored.
    event ConfigUnscheduled(uint64 numConfigs);

    BatchConfig[] public configs;
    BatchConfig public nextConfig;

    uint64 public immutable configChangeHeadsUpBlocks;

    constructor(uint64 headsUp) {
        configs.push(_zeroConfig());

        configChangeHeadsUpBlocks = headsUp;
    }

    function _zeroConfig() internal pure returns (BatchConfig memory) {
        return
            BatchConfig({
                startBatchIndex: 0,
                startBlockNumber: 0,
                keypers: new address[](0),
                threshold: 0,
                batchSpan: 0,
                batchSizeLimit: 0,
                transactionSizeLimit: 0,
                transactionGasLimit: 0,
                feeReceiver: address(0),
                targetAddress: address(0),
                targetFunctionSelector: bytes4(0),
                executionTimeout: 0
            });
    }

    function numConfigs() external view returns (uint64) {
        return uint64(configs.length);
    }

    /// @notice Get the config for a certain batch.
    /// @param batchIndex The index of the batch.
    function getConfig(uint64 batchIndex)
        external
        view
        returns (BatchConfig memory)
    {
        for (uint256 i = configs.length - 1; i >= 0; i--) {
            BatchConfig storage config = configs[i];
            if (config.startBatchIndex <= batchIndex) {
                return config;
            }
        }
        assert(false);
    }

    //
    // Config keyper getters
    //
    function configKeypers(uint64 configIndex, uint64 keyperIndex)
        external
        view
        returns (address)
    {
        return configs[configIndex].keypers[keyperIndex];
    }

    function configNumKeypers(uint64 configIndex)
        external
        view
        returns (uint64)
    {
        return uint64(configs[configIndex].keypers.length);
    }

    //
    // next config setters
    //
    function nextConfigSetStartBatchIndex(uint64 startBatchIndex)
        external
        onlyOwner
    {
        nextConfig.startBatchIndex = startBatchIndex;
    }

    function nextConfigSetStartBlockNumber(uint64 startBlockNumber)
        external
        onlyOwner
    {
        nextConfig.startBlockNumber = startBlockNumber;
    }

    function nextConfigSetThreshold(uint64 threshold) external onlyOwner {
        nextConfig.threshold = threshold;
    }

    function nextConfigSetBatchSpan(uint64 batchSpan) external onlyOwner {
        nextConfig.batchSpan = batchSpan;
    }

    function nextConfigSetBatchSizeLimit(uint64 batchSizeLimit)
        external
        onlyOwner
    {
        nextConfig.batchSizeLimit = batchSizeLimit;
    }

    function nextConfigSetTransactionSizeLimit(uint64 transactionSizeLimit)
        external
        onlyOwner
    {
        nextConfig.transactionSizeLimit = transactionSizeLimit;
    }

    function nextConfigSetTransactionGasLimit(uint64 transactionGasLimit)
        external
        onlyOwner
    {
        nextConfig.transactionGasLimit = transactionGasLimit;
    }

    function nextConfigSetFeeReceiver(address feeReceiver) external onlyOwner {
        nextConfig.feeReceiver = feeReceiver;
    }

    function nextConfigSetTargetAddress(address targetAddress)
        external
        onlyOwner
    {
        nextConfig.targetAddress = targetAddress;
    }

    function nextConfigSetTargetFunctionSelector(bytes4 targetFunctionSelector)
        external
        onlyOwner
    {
        nextConfig.targetFunctionSelector = targetFunctionSelector;
    }

    function nextConfigSetExecutionTimeout(uint64 executionTimeout)
        external
        onlyOwner
    {
        nextConfig.executionTimeout = executionTimeout;
    }

    function nextConfigAddKeypers(address[] calldata newKeypers)
        external
        onlyOwner
    {
        require(
            nextConfig.keypers.length <= type(uint64).max - newKeypers.length,
            "ConfigContract: number of keypers exceeds uint64"
        );
        for (uint64 i = 0; i < newKeypers.length; i++) {
            nextConfig.keypers.push(newKeypers[i]);
        }
    }

    function nextConfigRemoveKeypers(uint64 n) external onlyOwner {
        uint256 currentLength = nextConfig.keypers.length;
        if (n <= currentLength) {
            for (uint64 i = 0; i < n; i++) {
                nextConfig.keypers.pop();
            }
        } else {
            delete nextConfig.keypers;
        }
    }

    //
    // nextConfig keyper getters
    //
    function nextConfigKeypers(uint64 index) external view returns (address) {
        return nextConfig.keypers[index];
    }

    function nextConfigNumKeypers() external view returns (uint64) {
        return uint64(nextConfig.keypers.length);
    }

    //
    // Scheduling
    //

    /// @notice Finalize the `nextConfig` object and add it to the end of the config sequence.
    /// @notice `startBlockNumber` of the next config must be at least `configChangeHeadsUpBlocks`
    ///     blocks or the batch span of the current config in the future, whatever is greater.
    /// @notice The transition between the next config and the config currently at the end of the
    ///     config sequence must be seamless, i.e., the batches must not be cut short.
    function scheduleNextConfig() external onlyOwner {
        require(
            configs.length < type(uint64).max - 1,
            "ConfigContract: number of configs exceeds uint64"
        );
        BatchConfig memory config = configs[configs.length - 1];

        // check start block is not too early
        uint64 headsUp = configChangeHeadsUpBlocks;
        if (config.batchSpan > headsUp) {
            headsUp = config.batchSpan;
        }
        uint64 earliestStart = uint64(block.number) + headsUp + 1;
        require(
            nextConfig.startBlockNumber >= earliestStart,
            "ConfigContract: start block too early"
        );

        // check transition is seamless
        if (config.batchSpan > 0) {
            require(
                nextConfig.startBatchIndex > config.startBatchIndex,
                "ConfigContract: start batch index too small"
            );
            uint64 batchDelta =
                nextConfig.startBatchIndex - config.startBatchIndex;
            require(
                config.startBlockNumber + config.batchSpan * batchDelta ==
                    nextConfig.startBlockNumber,
                "ConfigContract: config transition not seamless"
            );
        } else {
            require(
                nextConfig.startBatchIndex == config.startBatchIndex,
                "ConfigContract: transition from inactive config with wrong start index"
            );
        }

        configs.push(nextConfig);
        nextConfig = _zeroConfig();

        emit ConfigScheduled(uint64(configs.length));
    }

    /// @notice Remove configs from the end.
    /// @param fromStartBlockNumber All configs with a start block number greater than or equal
    ///     to this will be removed.
    /// @notice `fromStartBlockNumber` must be `configChangeHeadsUpBlocks` blocks in the future.
    /// @notice This method can remove one or more configs. If no config would be removed, an error
    ///     is thrown.
    function unscheduleConfigs(uint64 fromStartBlockNumber) external onlyOwner {
        require(
            fromStartBlockNumber > block.number + configChangeHeadsUpBlocks,
            "ConfigContract: from start block too early"
        );

        uint64 lengthBefore = uint64(configs.length);

        for (uint256 i = configs.length - 1; i > 0; i--) {
            BatchConfig storage config = configs[i];
            if (config.startBlockNumber >= fromStartBlockNumber) {
                configs.pop();
            } else {
                break;
            }
        }

        require(
            configs.length < lengthBefore,
            "ConfigContract: no configs unscheduled"
        );
        emit ConfigUnscheduled(uint64(configs.length));
    }
}
