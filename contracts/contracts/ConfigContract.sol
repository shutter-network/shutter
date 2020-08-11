// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Ownable.sol";

struct BatchConfig {
    uint256 startBatchIndex;
    uint256 startBlockNumber;
    bool active;
    address[] keypers;
    uint256 threshold;
    uint256 batchSpan;
    uint256 batchSizeLimit;
    uint256 transactionSizeLimit;
    uint256 transactionGasLimit;
    address feeReceiver;
    address targetAddress;
    bytes4 targetFunctionSelector;
    uint256 executionTimeout;
}

contract ConfigContract is Ownable {

    event ConfigScheduled(
        uint256 numConfigs
    );
    event ConfigUnscheduled(
        uint256 numConfigs
    );

    BatchConfig[] public configs;
    BatchConfig public nextConfig;

    uint256 public immutable config_change_heads_up_blocks;

    constructor(uint256 _config_change_heads_up_blocks) {
        configs.push(zeroConfig());

        config_change_heads_up_blocks = _config_change_heads_up_blocks;
    }

    function zeroConfig() internal pure returns (BatchConfig memory) {
        return BatchConfig({
            startBatchIndex: 0,
            startBlockNumber: 0,
            active: false,
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

    function numConfigs() external view returns (uint256) {
        return configs.length;
    }

    function getConfig(uint256 _batchIndex) external view returns (BatchConfig memory) {
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
    function configKeypers(uint256 _configIndex, uint256 _keyperIndex) external view returns (address) {
        return configs[_configIndex].keypers[_keyperIndex];
    }

    function configNumKeypers(uint256 _configIndex) external view returns (uint256) {
        return configs[_configIndex].keypers.length;
    }

    //
    // next config setters
    //
    function nextConfigSetStartBatchIndex(uint256 _startBatchIndex) external onlyOwner {
        nextConfig.startBatchIndex = _startBatchIndex;
    }

    function nextConfigSetStartBlockNumber(uint256 _startBlockNumber) external onlyOwner {
        nextConfig.startBlockNumber = _startBlockNumber;
    }

    function nextConfigSetActive(bool _active) external onlyOwner {
        nextConfig.active = _active;
    }

    function nextConfigSetThreshold(uint256 _threshold) external onlyOwner {
        nextConfig.threshold = _threshold;
    }

    function nextConfigSetBatchSpan(uint256 _batchSpan) external onlyOwner {
        nextConfig.batchSpan = _batchSpan;
    }

    function nextConfigSetBatchSizeLimit(uint256 _batchSizeLimit) external onlyOwner {
        nextConfig.batchSizeLimit = _batchSizeLimit;
    }

    function nextConfigSetTransactionSizeLimit(uint256 _transactionSizeLimit) external onlyOwner {
        nextConfig.transactionSizeLimit = _transactionSizeLimit;
    }

    function nextConfigSetTransactionGasLimit(uint256 _transactionGasLimit) external onlyOwner {
        nextConfig.transactionGasLimit = _transactionGasLimit;
    }

    function nextConfigSetFeeReceiver(address _feeReceiver) external onlyOwner {
        nextConfig.feeReceiver = _feeReceiver;
    }

    function nextConfigSetTargetAddress(address _targetAddress) external onlyOwner {
        nextConfig.targetAddress = _targetAddress;
    }

    function nextConfigSetTargetFunctionSelector(bytes4 _targetFunctionSelector) external onlyOwner {
        nextConfig.targetFunctionSelector = _targetFunctionSelector;
    }

    function nextConfigSetExecutionTimeout(uint256 _executionTimeout) external onlyOwner {
        nextConfig.executionTimeout = _executionTimeout;
    }

    function nextConfigAddKeypers(address[] calldata _newKeypers) external onlyOwner {
        for (uint i = 0; i < _newKeypers.length; i++) {
            nextConfig.keypers.push(_newKeypers[i]);
        }
    }

    function nextConfigRemoveKeypers(uint256 n) external onlyOwner {
        uint256 currentLength = nextConfig.keypers.length;
        if (n <= currentLength) {
            for (uint256 i = 0; i < n; i++) {
                nextConfig.keypers.pop();
            }
        } else{
            delete nextConfig.keypers;
        }
    }

    //
    // nextConfig keyper getters
    //
    function nextConfigKeypers(uint256 _index) external view returns (address) {
        return nextConfig.keypers[_index];
    }

    function nextConfigNumKeypers() external view returns (uint256) {
        return nextConfig.keypers.length;
    }

    //
    // Scheduling
    //
    function scheduleNextConfig() external onlyOwner {
        BatchConfig memory config = configs[configs.length - 1];

        require(nextConfig.startBlockNumber > block.number + config_change_heads_up_blocks);

        if (config.active) {
            uint256 batchDelta = nextConfig.startBatchIndex - config.startBatchIndex;
            require(batchDelta > 0);
            require(config.startBlockNumber + config.batchSpan * batchDelta == nextConfig.startBlockNumber);
        } else {
            require(nextConfig.startBatchIndex == config.startBatchIndex);
        }

        if (nextConfig.active) {
            require(nextConfig.batchSpan > 0);
        } else {
            require(nextConfig.batchSpan == 0);
        }

        configs.push(nextConfig);
        nextConfig = zeroConfig();

        emit ConfigScheduled(configs.length);
    }

    function unscheduleConfigs(uint256 _fromStartBlockNumber) external onlyOwner {
        require(_fromStartBlockNumber > block.number + config_change_heads_up_blocks);

        for (uint256 i = configs.length - 1; i > 0; i--) {
            BatchConfig storage config = configs[i];
            if (config.startBlockNumber >= _fromStartBlockNumber) {
                configs.pop();
            } else {
                break;
            }
        }

        emit ConfigUnscheduled(configs.length);
    }
}
