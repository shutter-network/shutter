package observe

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

// MainChain let's a keyper fetch all necessary information from an ethereum node to do it's
// work. The only source for the data stored in this struct should be the ethereum node.  The
// SyncToHead method can be used to update the data. All other accesses should be read-only.
type MainChain struct {
	CurrentBlock uint64
	BatchConfigs []contract.BatchConfig
}

// NewMainChain creates an empty MainChain struct
func NewMainChain() *MainChain {
	return &MainChain{}
}

// IsActiveKeyper checks if the given address is registered as a keyper in one of the active batch
// configs.
func (mainchain *MainChain) IsActiveKeyper(addr common.Address) bool {
	for _, config := range mainchain.ActiveConfigs() {
		if config.IsKeyper(addr) {
			return true
		}
	}
	return false
}

// ActiveConfigIndex returns the index of the first config that is active.
func (mainchain *MainChain) ActiveConfigIndex() int {
	for i := len(mainchain.BatchConfigs) - 1; i >= 0; i-- {
		if mainchain.BatchConfigs[i].StartBlockNumber <= mainchain.CurrentBlock {
			return i
		}
	}
	panic("illegal values in MainChain.Configs field")
}

// ActiveConfigs returns a slice of the active configs.
func (mainchain *MainChain) ActiveConfigs() []contract.BatchConfig {
	return mainchain.BatchConfigs[mainchain.ActiveConfigIndex():]
}

func (mainchain *MainChain) syncConfigs(configContract *contract.ConfigContract, opts *bind.CallOpts) error {
	numConfigs, err := configContract.NumConfigs(opts)
	if err != nil {
		return err
	}
	for configIndex := uint64(len(mainchain.BatchConfigs)); configIndex < numConfigs; configIndex++ {
		config, err := configContract.GetConfigByIndex(opts, configIndex)
		if err != nil {
			return err
		}
		mainchain.BatchConfigs = append(mainchain.BatchConfigs, config)
	}
	return nil
}

// SyncToHead fetches the latest state from the ethereum node.
// XXX this mutates the object in place. we may want to control mutation of the MainChain struct.
// XXX We can't use keyper.ContractCaller here because we would end up with an import cycle.
func (mainchain *MainChain) SyncToHead(ctx context.Context, ethcl *ethclient.Client, configContract *contract.ConfigContract) error {
	latestBlockHeader, err := ethcl.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	latestBlockNumber := latestBlockHeader.Number.Uint64()
	if latestBlockNumber == mainchain.CurrentBlock {
		return nil
	}

	opts := &bind.CallOpts{
		BlockNumber: latestBlockHeader.Number,
		Context:     ctx,
	}

	err = mainchain.syncConfigs(configContract, opts)
	if err != nil {
		return err
	}

	mainchain.CurrentBlock = latestBlockNumber
	return nil
}
