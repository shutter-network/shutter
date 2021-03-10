package contract

// This file adds some custom methods to the abigen generated ConfigContract class

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/brainbot-com/shutter/shuttermint/medley"
)

// KeyperIndex returns the index of the keyper identified by the given address
func (bc *BatchConfig) KeyperIndex(address common.Address) (uint64, bool) {
	for i, k := range bc.Keypers {
		if k == address {
			return uint64(i), true
		}
	}
	return 0, false
}

// IsKeyper returns true if the given address is part of the keyper set, otherwise false
func (bc *BatchConfig) IsKeyper(address common.Address) bool {
	_, isKeyper := bc.KeyperIndex(address)
	return isKeyper
}

// IsActive checks if the config is active, i.e. the batch span is non-zero
func (bc *BatchConfig) IsActive() bool {
	return bc.BatchSpan > 0
}

// BatchStartBlock returns the StartBlock for the given batch index. This function will panic if
// the batchIndex is less than the BatchConfig's StartBatchIndex.
func (bc *BatchConfig) BatchStartBlock(batchIndex uint64) uint64 {
	if batchIndex < bc.StartBatchIndex {
		panic("BatchStartBlock called with bad parameter")
	}
	relativeBatchIndex := batchIndex - bc.StartBatchIndex
	return bc.StartBlockNumber + relativeBatchIndex*bc.BatchSpan
}

// BatchEndBlock returns the end block for the given batch index. This function will panic if the
// batchIndex is less than the BatchConfig's StartBatchIndex.
func (bc *BatchConfig) BatchEndBlock(batchIndex uint64) uint64 {
	return bc.BatchStartBlock(batchIndex) + bc.BatchSpan
}

// BatchIndex returns the BatchIndex for the given blockNumber. This function will panic if the
// blockNumber is less than the BatchConfig's StartBlockNumber. If the BatchConfig is not active,
// i.e. it's BatchSpan is zero, it will return the StartBatchIndex for all blockNumbers.
func (bc *BatchConfig) BatchIndex(blockNumber uint64) uint64 {
	if blockNumber < bc.StartBlockNumber {
		panic("internal error: BatchIndex called with bad blockNumber")
	}
	if bc.BatchSpan == 0 {
		return bc.StartBatchIndex
	}
	return bc.StartBatchIndex + (blockNumber-bc.StartBlockNumber)/bc.BatchSpan
}

// NextBatchIndex determines the next batch index to be started after the given block number.
func (cc *ConfigContract) NextBatchIndex(blockNumber uint64) (uint64, error) {
	i, err := cc.NumConfigs(nil)
	if err != nil {
		return 0, err
	}
	for {
		if i == 0 {
			return uint64(0), errors.Errorf("contract misconfigured")
		}
		i--

		cfg, err := cc.Configs(nil, big.NewInt(0).SetUint64(i))
		if err != nil {
			return 0, err
		}

		startBlockNumber := cfg.StartBlockNumber
		if startBlockNumber <= blockNumber {
			batchSpan := cfg.BatchSpan
			if batchSpan == 0 {
				return cfg.StartBatchIndex, nil
			}
			next := cfg.StartBatchIndex + (blockNumber-startBlockNumber)/batchSpan + 1
			return next, nil
		}
	}
}

// GetConfigKeypers queries the list of keypers defined in the config given by its index
func (cc *ConfigContract) GetConfigKeypers(opts *bind.CallOpts, configIndex uint64) ([]common.Address, error) {
	var keypers []common.Address

	numKeypers, err := cc.ConfigNumKeypers(opts, configIndex)
	if err != nil {
		return keypers, err
	}

	for i := uint64(0); i < numKeypers; i++ {
		keyper, err := cc.ConfigKeypers(opts, configIndex, i)
		if err != nil {
			return keypers, err
		}
		keypers = append(keypers, keyper)
	}

	return keypers, nil
}

// GetNextConfigKeypers queries the list of keypers set for the next config to be scheduled
func (cc *ConfigContract) GetNextConfigKeypers(opts *bind.CallOpts) ([]common.Address, error) {
	var keypers []common.Address

	numKeypers, err := cc.NextConfigNumKeypers(opts)
	if err != nil {
		return keypers, err
	}

	for i := uint64(0); i < numKeypers; i++ {
		keyper, err := cc.NextConfigKeypers(opts, i)
		if err != nil {
			return keypers, err
		}
		keypers = append(keypers, keyper)
	}

	return keypers, nil
}

// GetConfigByIndex queries the batch config by its index (not the batch index, but the config index)
func (cc *ConfigContract) GetConfigByIndex(opts *bind.CallOpts, configIndex uint64) (BatchConfig, error) {
	config, err := cc.Configs(opts, big.NewInt(0).SetUint64(configIndex))
	if err != nil {
		return BatchConfig{}, err
	}

	keypers, err := cc.GetConfigKeypers(opts, configIndex)
	if err != nil {
		return BatchConfig{}, err
	}

	return BatchConfig{
		StartBatchIndex:        config.StartBatchIndex,
		StartBlockNumber:       config.StartBlockNumber,
		Threshold:              config.Threshold,
		BatchSpan:              config.BatchSpan,
		BatchSizeLimit:         config.BatchSizeLimit,
		TransactionSizeLimit:   config.TransactionSizeLimit,
		TransactionGasLimit:    config.TransactionGasLimit,
		FeeReceiver:            config.FeeReceiver,
		TargetAddress:          config.TargetAddress,
		TargetFunctionSelector: config.TargetFunctionSelector,
		ExecutionTimeout:       config.ExecutionTimeout,
		Keypers:                medley.DedupAddresses(keypers),
	}, nil
}

// GetNextConfig queries the next batch config.
func (cc *ConfigContract) GetNextConfig(opts *bind.CallOpts) (BatchConfig, error) {
	config, err := cc.NextConfig(opts)
	if err != nil {
		return BatchConfig{}, err
	}

	keypers, err := cc.GetNextConfigKeypers(opts)
	if err != nil {
		return BatchConfig{}, err
	}

	return BatchConfig{
		StartBatchIndex:        config.StartBatchIndex,
		StartBlockNumber:       config.StartBlockNumber,
		Threshold:              config.Threshold,
		BatchSpan:              config.BatchSpan,
		BatchSizeLimit:         config.BatchSizeLimit,
		TransactionSizeLimit:   config.TransactionSizeLimit,
		TransactionGasLimit:    config.TransactionGasLimit,
		FeeReceiver:            config.FeeReceiver,
		TargetAddress:          config.TargetAddress,
		TargetFunctionSelector: config.TargetFunctionSelector,
		ExecutionTimeout:       config.ExecutionTimeout,
		Keypers:                medley.DedupAddresses(keypers),
	}, nil
}

// CurrentAndFutureConfigs fetches the config that is active at the given block number as well as
// all following ones. They will be ordered from oldest to newest.
func (cc *ConfigContract) CurrentAndFutureConfigs(opts *bind.CallOpts, blockNumber uint64) ([]BatchConfig, error) {
	emptyResult := []BatchConfig{}

	numConfigs, err := cc.NumConfigs(opts)
	if err != nil {
		return emptyResult, err
	}

	configs := []BatchConfig{}
	for i := numConfigs; i >= 1; i-- {
		config, err := cc.GetConfigByIndex(opts, i-1)
		if err != nil {
			return emptyResult, err
		}

		configs = append([]BatchConfig{config}, configs...) // prepend

		if config.StartBlockNumber <= blockNumber {
			break
		}
	}

	return configs, nil
}
