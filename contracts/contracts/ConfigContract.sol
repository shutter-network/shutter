// SPDX-License-Identifier: MIT

pragma solidity =0.8.4;

import {Ownable} from "openzeppelin/contracts/access/Ownable.sol";

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

    // Stores all scheduled configs, plus the next config at the end
    BatchConfig[] public configs;

    uint64 public immutable configChangeHeadsUpBlocks;

    constructor(uint64 headsUp) {
        configs.push(_zeroConfig()); // guard
        configs.push(_zeroConfig()); // next config

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

    function numConfigs() public view returns (uint64) {
        return uint64(configs.length) - 1; // exclude next config
    }

    function nextConfigIndex() public view returns (uint64) {
        return uint64(configs.length) - 1;
    }

    function nextConfig() public view returns (BatchConfig memory) {
        return configs[nextConfigIndex()];
    }

    function lastScheduledConfigIndex() public view returns (uint64) {
        return uint64(configs.length) - 2;
    }

    function lastScheduledConfig() public view returns (BatchConfig memory) {
        return configs[lastScheduledConfigIndex()];
    }

    /// @notice Get the index of the config for the batch with the given index.
    /// @param batchIndex The index of the batch.
    function configIndexForBatchIndex(uint64 batchIndex)
        public
        view
        returns (uint64)
    {
        for (uint256 i = configs.length - 2; i >= 0; i--) {
            if (configs[i].startBatchIndex <= batchIndex) {
                return uint64(i);
            }
        }
        assert(false);
        return 0; // only for the linter
    }

    /// @notice Get the config for a certain batch.
    /// @param batchIndex The index of the batch.
    function configForBatchIndex(uint64 batchIndex)
        public
        view
        returns (BatchConfig memory)
    {
        uint64 configIndex = configIndexForBatchIndex(batchIndex);
        return configs[configIndex];
    }

    function configForConfigIndex(uint64 configIndex)
        public
        view
        returns (BatchConfig memory)
    {
        return configs[configIndex];
    }

    //
    // Config field getters
    //
    function configKeypers(uint64 configIndex, uint64 keyperIndex)
        public
        view
        returns (address)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].keypers[keyperIndex];
    }

    function configNumKeypers(uint64 configIndex) public view returns (uint64) {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return uint64(configs[configIndex].keypers.length);
    }

    function configStartBatchIndex(uint64 configIndex)
        public
        view
        returns (uint64)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].startBatchIndex;
    }

    function configStartBlockNumber(uint64 configIndex)
        public
        view
        returns (uint64)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].startBlockNumber;
    }

    function configThreshold(uint64 configIndex) public view returns (uint64) {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].threshold;
    }

    function configBatchSpan(uint64 configIndex) public view returns (uint64) {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].batchSpan;
    }

    function configBatchSizeLimit(uint64 configIndex)
        public
        view
        returns (uint64)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].batchSizeLimit;
    }

    function configTransactionSizeLimit(uint64 configIndex)
        public
        view
        returns (uint64)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].transactionSizeLimit;
    }

    function configTransactionGasLimit(uint64 configIndex)
        public
        view
        returns (uint64)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].transactionGasLimit;
    }

    function configFeeReceiver(uint64 configIndex)
        public
        view
        returns (address)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].feeReceiver;
    }

    function configTargetAddress(uint64 configIndex)
        public
        view
        returns (address)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].targetAddress;
    }

    function configTargetFunctionSelector(uint64 configIndex)
        public
        view
        returns (bytes4)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].targetFunctionSelector;
    }

    function configExecutionTimeout(uint64 configIndex)
        public
        view
        returns (uint64)
    {
        require(
            configIndex < numConfigs(),
            "ConfigContract: config index out of range"
        );
        return configs[configIndex].executionTimeout;
    }

    //
    // next config setters
    //
    function nextConfigSetStartBatchIndex(uint64 startBatchIndex)
        public
        onlyOwner
    {
        configs[nextConfigIndex()].startBatchIndex = startBatchIndex;
    }

    function nextConfigSetStartBlockNumber(uint64 startBlockNumber)
        public
        onlyOwner
    {
        configs[nextConfigIndex()].startBlockNumber = startBlockNumber;
    }

    function nextConfigSetThreshold(uint64 threshold) public onlyOwner {
        configs[nextConfigIndex()].threshold = threshold;
    }

    function nextConfigSetBatchSpan(uint64 batchSpan) public onlyOwner {
        configs[nextConfigIndex()].batchSpan = batchSpan;
    }

    function nextConfigSetBatchSizeLimit(uint64 batchSizeLimit)
        public
        onlyOwner
    {
        configs[nextConfigIndex()].batchSizeLimit = batchSizeLimit;
    }

    function nextConfigSetTransactionSizeLimit(uint64 transactionSizeLimit)
        public
        onlyOwner
    {
        configs[nextConfigIndex()].transactionSizeLimit = transactionSizeLimit;
    }

    function nextConfigSetTransactionGasLimit(uint64 transactionGasLimit)
        public
        onlyOwner
    {
        configs[nextConfigIndex()].transactionGasLimit = transactionGasLimit;
    }

    function nextConfigSetFeeReceiver(address feeReceiver) public onlyOwner {
        configs[nextConfigIndex()].feeReceiver = feeReceiver;
    }

    function nextConfigSetTargetAddress(address targetAddress)
        public
        onlyOwner
    {
        configs[nextConfigIndex()].targetAddress = targetAddress;
    }

    function nextConfigSetTargetFunctionSelector(bytes4 targetFunctionSelector)
        public
        onlyOwner
    {
        configs[nextConfigIndex()]
            .targetFunctionSelector = targetFunctionSelector;
    }

    function nextConfigSetExecutionTimeout(uint64 executionTimeout)
        public
        onlyOwner
    {
        configs[nextConfigIndex()].executionTimeout = executionTimeout;
    }

    function nextConfigAddKeypers(address[] calldata newKeypers)
        public
        onlyOwner
    {
        require(
            configs[nextConfigIndex()].keypers.length <=
                type(uint64).max - newKeypers.length,
            "ConfigContract: number of keypers exceeds uint64"
        );
        for (uint64 i = 0; i < newKeypers.length; i++) {
            configs[nextConfigIndex()].keypers.push(newKeypers[i]);
        }
    }

    function nextConfigRemoveKeypers(uint64 n) public onlyOwner {
        uint256 currentLength = nextConfig().keypers.length;
        if (n <= currentLength) {
            for (uint64 i = 0; i < n; i++) {
                configs[nextConfigIndex()].keypers.pop();
            }
        } else {
            delete configs[nextConfigIndex()].keypers;
        }
    }

    //
    // nextConfig getters
    //
    function nextConfigKeypers(uint64 keyperIndex)
        public
        view
        returns (address)
    {
        return nextConfig().keypers[keyperIndex];
    }

    function nextConfigNumKeypers() public view returns (uint64) {
        return uint64(nextConfig().keypers.length);
    }

    function nextConfigStartBatchIndex() public view returns (uint64) {
        return nextConfig().startBatchIndex;
    }

    function nextConfigStartBlockNumber() public view returns (uint64) {
        return nextConfig().startBlockNumber;
    }

    function nextConfigThreshold() public view returns (uint64) {
        return nextConfig().threshold;
    }

    function nextConfigBatchSpan() public view returns (uint64) {
        return nextConfig().batchSpan;
    }

    function nextConfigBatchSizeLimit() public view returns (uint64) {
        return nextConfig().batchSizeLimit;
    }

    function nextConfigTransactionSizeLimit() public view returns (uint64) {
        return nextConfig().transactionSizeLimit;
    }

    function nextConfigTransactionGasLimit() public view returns (uint64) {
        return nextConfig().transactionGasLimit;
    }

    function nextConfigFeeReceiver() public view returns (address) {
        return nextConfig().feeReceiver;
    }

    function nextConfigTargetAddress() public view returns (address) {
        return nextConfig().targetAddress;
    }

    function nextConfigTargetFunctionSelector() public view returns (bytes4) {
        return nextConfig().targetFunctionSelector;
    }

    function nextConfigExecutionTimeout() public view returns (uint64) {
        return nextConfig().executionTimeout;
    }

    //
    // Scheduling
    //

    /// @notice Finalize the `nextConfig` object and add it to the end of the config sequence.
    /// @notice `startBlockNumber` of the next config must be at least `configChangeHeadsUpBlocks`
    ///     blocks or the batch span of the current config in the future, whatever is greater.
    /// @notice The transition between the next config and the config currently at the end of the
    ///     config sequence must be seamless, i.e., the batches must not be cut short.
    function scheduleNextConfig() public onlyOwner {
        require(
            configs.length < type(uint64).max - 1,
            "ConfigContract: number of configs exceeds uint64"
        );
        BatchConfig memory config1 = lastScheduledConfig();
        BatchConfig memory config2 = nextConfig();

        // check start block is not too early
        uint64 headsUp = configChangeHeadsUpBlocks;
        if (config1.batchSpan > headsUp) {
            headsUp = config1.batchSpan;
        }
        uint64 earliestStart = uint64(block.number) + headsUp + 1;
        require(
            config2.startBlockNumber >= earliestStart,
            "ConfigContract: start block too early"
        );

        // check transition is seamless
        if (config1.batchSpan > 0) {
            require(
                config2.startBatchIndex > config1.startBatchIndex,
                "ConfigContract: start batch index too small"
            );
            uint64 batchDelta =
                config2.startBatchIndex - config1.startBatchIndex;
            require(
                config1.startBlockNumber + config1.batchSpan * batchDelta ==
                    config2.startBlockNumber,
                "ConfigContract: config transition not seamless"
            );
        } else {
            require(
                config2.startBatchIndex == config1.startBatchIndex,
                "ConfigContract: transition from inactive config with wrong start index"
            );
        }

        configs.push(_zeroConfig());

        emit ConfigScheduled(numConfigs());
    }

    /// @notice Remove configs from the end.
    /// @param fromStartBlockNumber All configs with a start block number greater than or equal
    ///     to this will be removed.
    /// @notice `fromStartBlockNumber` must be `configChangeHeadsUpBlocks` blocks in the future.
    /// @notice This method can remove one or more configs. If no config would be removed, an error
    ///     is thrown.
    /// @notice This invalidates the currently set next config.
    function unscheduleConfigs(uint64 fromStartBlockNumber) public onlyOwner {
        require(
            fromStartBlockNumber > block.number + configChangeHeadsUpBlocks,
            "ConfigContract: from start block too early"
        );

        uint64 lengthBefore = uint64(configs.length);

        BatchConfig memory nextConfig_ = nextConfig();
        configs.pop();

        for (uint256 i = configs.length - 1; i > 0; i--) {
            BatchConfig storage config = configs[i];
            if (config.startBlockNumber >= fromStartBlockNumber) {
                configs.pop();
            } else {
                break;
            }
        }

        configs.push(nextConfig_); // add the next config again

        require(
            configs.length < lengthBefore,
            "ConfigContract: no configs unscheduled"
        );
        emit ConfigUnscheduled(numConfigs());
    }

    function batchingActive(uint64 configIndex) public view returns (bool) {
        return configs[configIndex].batchSpan > 0;
    }

    function batchBoundaryBlocks(uint64 configIndex, uint64 batchIndex)
        public
        view
        returns (
            uint64,
            uint64,
            uint64
        )
    {
        uint64 startBlockNumber = configs[configIndex].startBlockNumber;
        uint64 startBatchIndex = configs[configIndex].startBatchIndex;
        uint64 batchSpan = configs[configIndex].batchSpan;
        uint64 executionTimeout = configs[configIndex].executionTimeout;

        assert(batchSpan > 0);
        assert(batchIndex >= startBatchIndex);

        uint64 relativeBatchIndex = batchIndex - startBatchIndex;
        uint64 end =
            startBlockNumber + relativeBatchIndex * batchSpan + batchSpan;
        uint64 start = end - batchSpan;
        if (relativeBatchIndex >= 1) {
            start -= batchSpan;
        }
        uint64 timeout = end + executionTimeout;

        return (start, end, timeout);
    }
}
