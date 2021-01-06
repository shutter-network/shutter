// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "OpenZeppelin/openzeppelin-contracts@3.3.0/contracts/access/Ownable.sol";

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

    constructor(uint64 _configChangeHeadsUpBlocks) public {
        configs.push(zeroConfig());

        configChangeHeadsUpBlocks = _configChangeHeadsUpBlocks;
    }

    function zeroConfig() internal pure returns (BatchConfig memory) {
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
    /// @param _batchIndex The index of the batch.
    function getConfig(uint64 _batchIndex)
        external
        view
        returns (BatchConfig memory)
    {
        for (uint256 i = configs.length - 1; i >= 0; i--) {
            BatchConfig storage config = configs[i];
            if (config.startBatchIndex <= _batchIndex) {
                return config;
            }
        }
        assert(false);
    }

    //
    // Config keyper getters
    //
    function configKeypers(uint64 _configIndex, uint64 _keyperIndex)
        external
        view
        returns (address)
    {
        return configs[_configIndex].keypers[_keyperIndex];
    }

    function configNumKeypers(uint64 _configIndex)
        external
        view
        returns (uint64)
    {
        return uint64(configs[_configIndex].keypers.length);
    }

    //
    // next config setters
    //
    function nextConfigSetStartBatchIndex(uint64 _startBatchIndex)
        external
        onlyOwner
    {
        nextConfig.startBatchIndex = _startBatchIndex;
    }

    function nextConfigSetStartBlockNumber(uint64 _startBlockNumber)
        external
        onlyOwner
    {
        nextConfig.startBlockNumber = _startBlockNumber;
    }

    function nextConfigSetThreshold(uint64 _threshold) external onlyOwner {
        nextConfig.threshold = _threshold;
    }

    function nextConfigSetBatchSpan(uint64 _batchSpan) external onlyOwner {
        // make sure the heads up is at least one batch
        require(_batchSpan < configChangeHeadsUpBlocks);
        nextConfig.batchSpan = _batchSpan;
    }

    function nextConfigSetBatchSizeLimit(uint64 _batchSizeLimit)
        external
        onlyOwner
    {
        nextConfig.batchSizeLimit = _batchSizeLimit;
    }

    function nextConfigSetTransactionSizeLimit(uint64 _transactionSizeLimit)
        external
        onlyOwner
    {
        nextConfig.transactionSizeLimit = _transactionSizeLimit;
    }

    function nextConfigSetTransactionGasLimit(uint64 _transactionGasLimit)
        external
        onlyOwner
    {
        nextConfig.transactionGasLimit = _transactionGasLimit;
    }

    function nextConfigSetFeeReceiver(address _feeReceiver) external onlyOwner {
        nextConfig.feeReceiver = _feeReceiver;
    }

    function nextConfigSetTargetAddress(address _targetAddress)
        external
        onlyOwner
    {
        nextConfig.targetAddress = _targetAddress;
    }

    function nextConfigSetTargetFunctionSelector(bytes4 _targetFunctionSelector)
        external
        onlyOwner
    {
        nextConfig.targetFunctionSelector = _targetFunctionSelector;
    }

    function nextConfigSetExecutionTimeout(uint64 _executionTimeout)
        external
        onlyOwner
    {
        nextConfig.executionTimeout = _executionTimeout;
    }

    function nextConfigAddKeypers(address[] calldata _newKeypers)
        external
        onlyOwner
    {
        require(
            nextConfig.keypers.length <= type(uint64).max - _newKeypers.length
        );
        for (uint64 i = 0; i < _newKeypers.length; i++) {
            nextConfig.keypers.push(_newKeypers[i]);
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
    function nextConfigKeypers(uint64 _index) external view returns (address) {
        return nextConfig.keypers[_index];
    }

    function nextConfigNumKeypers() external view returns (uint64) {
        return uint64(nextConfig.keypers.length);
    }

    //
    // Scheduling
    //

    /// @notice Finalize the `nextConfig` object and add it to the end of the config sequence.
    /// @notice `startBlockNumber` of the next config must be `configChangeHeadsUpBlocks`
    ///     blocks in the future. Note that the batch spans are smaller than or equal to
    ///     `configChangeHeadsUpBlocks`, so the heads up corresponds to at least one batch.
    /// @notice The transition between the next config and the config currently at the end of the
    ///     config sequence must be seamless, i.e., there batches must not be cut short.
    function scheduleNextConfig() external onlyOwner {
        require(configs.length < type(uint64).max - 1);
        BatchConfig memory _config = configs[configs.length - 1];

        require(
            nextConfig.startBlockNumber >
                block.number + configChangeHeadsUpBlocks
        );

        if (_config.batchSpan > 0) {
            require(nextConfig.startBatchIndex > _config.startBatchIndex);
            uint64 _batchDelta = nextConfig.startBatchIndex -
                _config.startBatchIndex;
            require(
                _config.startBlockNumber + _config.batchSpan * _batchDelta ==
                    nextConfig.startBlockNumber
            );
        } else {
            require(nextConfig.startBatchIndex == _config.startBatchIndex);
        }

        configs.push(nextConfig);
        nextConfig = zeroConfig();

        emit ConfigScheduled(uint64(configs.length));
    }

    /// @notice Remove configs from the end.
    /// @param _fromStartBlockNumber All configs with a start block number greater than or equal
    ///     to this will be removed.
    /// @notice `_fromStartBlockNumber` must be `configChangeHeadsUpBlocks` blocks in the future.
    /// @notice This method can remove one or more configs. If no config would be removed, an error
    ///     is thrown.
    function unscheduleConfigs(uint64 _fromStartBlockNumber)
        external
        onlyOwner
    {
        require(
            _fromStartBlockNumber > block.number + configChangeHeadsUpBlocks
        );

        uint64 _lengthBefore = uint64(configs.length);

        for (uint256 i = configs.length - 1; i > 0; i--) {
            BatchConfig storage config = configs[i];
            if (config.startBlockNumber >= _fromStartBlockNumber) {
                configs.pop();
            } else {
                break;
            }
        }

        require(configs.length < _lengthBefore);
        emit ConfigUnscheduled(uint64(configs.length));
    }
}
