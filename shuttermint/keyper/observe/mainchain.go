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
	CurrentBlock            uint64
	BatchConfigs            []contract.BatchConfig
	Batches                 map[uint64]*Batch
	NumExecutionHalfSteps   uint64
	CipherExecutionReceipts map[uint64]*contract.CipherExecutionReceipt
}

// Batch stores the encrypted and plain transactions submitted to the batching contract for a
// particular index.
type Batch struct {
	BatchIndex            uint64
	EncryptedBatchHash    common.Hash
	EncryptedTransactions [][]byte
	PlainTransactions     [][]byte
	PlainBatchHash        common.Hash
}

// NewMainChain creates an empty MainChain struct
func NewMainChain() *MainChain {
	return &MainChain{
		Batches:                 make(map[uint64]*Batch),
		CipherExecutionReceipts: make(map[uint64]*contract.CipherExecutionReceipt),
	}
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

// ActiveConfigIndex returns the index of the config that is active for the given block number
func (mainchain *MainChain) ActiveConfigIndex(blocknum uint64) int {
	for i := len(mainchain.BatchConfigs) - 1; i >= 0; i-- {
		if mainchain.BatchConfigs[i].StartBlockNumber <= blocknum {
			return i
		}
	}
	panic("illegal values in MainChain.Configs field")
}

// ActiveConfigs returns a slice of the active configs.
func (mainchain *MainChain) ActiveConfigs() []contract.BatchConfig {
	return mainchain.BatchConfigs[mainchain.ActiveConfigIndex(mainchain.CurrentBlock):]
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

func (mainchain *MainChain) syncBatches(batcherContract *contract.BatcherContract, filter *bind.FilterOpts) error {
	it, err := batcherContract.FilterTransactionAdded(filter)
	if err != nil {
		return err
	}

	events := []*contract.BatcherContractTransactionAdded{}
	for it.Next() {
		events = append(events, it.Event)
	}
	if it.Error() != nil {
		return it.Error()
	}

	for _, event := range events {
		mainchain.addTransaction(event)
	}

	return nil
}

// AddTransaction adds a transaction to a batch according to a main chain TransactionAdded event.
func (mainchain *MainChain) addTransaction(event *contract.BatcherContractTransactionAdded) {
	batch, ok := mainchain.Batches[event.BatchIndex]
	if !ok {
		batch.BatchIndex = event.BatchIndex
		// for the rest of the fields, the zero values are fine
	}

	switch event.TransactionType {
	case contract.TransactionTypeCipher:
		batch.EncryptedTransactions = append(batch.EncryptedTransactions, event.Transaction)
		batch.EncryptedBatchHash = event.BatchHash
	case contract.TransactionTypePlain:
		batch.PlainTransactions = append(batch.EncryptedTransactions, event.Transaction)
		batch.PlainBatchHash = event.BatchHash
	default:
		panic("unknown transaction type")
	}

	mainchain.Batches[event.BatchIndex] = batch
}

func (mainchain *MainChain) syncExecutionState(executorContract *contract.ExecutorContract, opts *bind.CallOpts) error {
	lastNumExecutionHalfSteps := mainchain.NumExecutionHalfSteps

	numExecutionHalfSteps, err := executorContract.NumExecutionHalfSteps(opts)
	if err != nil {
		return err
	}

	receipts := []contract.CipherExecutionReceipt{}
	for i := lastNumExecutionHalfSteps; i < numExecutionHalfSteps; i++ {
		if i%2 != 0 {
			// there will only be receipts for cipher execution steps, which are the even ones
			continue
		}

		receipt, err := executorContract.CipherExecutionReceipts(opts, i)
		if err != nil {
			return err
		}
		receipts = append(receipts, receipt)
	}

	mainchain.NumExecutionHalfSteps = numExecutionHalfSteps
	for _, receipt := range receipts {
		mainchain.CipherExecutionReceipts[receipt.HalfStep] = &receipt
	}

	return nil
}

// SyncToHead fetches the latest state from the ethereum node.
// XXX this mutates the object in place. we may want to control mutation of the MainChain struct.
// XXX We can't use keyper.ContractCaller here because we would end up with an import cycle.
func (mainchain *MainChain) SyncToHead(ctx context.Context, ethcl *ethclient.Client, configContract *contract.ConfigContract, batcherContract *contract.BatcherContract, executorContract *contract.ExecutorContract) error {
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
	filter := &bind.FilterOpts{
		Start: mainchain.CurrentBlock,
		End:   &latestBlockNumber,
	}

	err = mainchain.syncConfigs(configContract, opts)
	if err != nil {
		return err
	}

	err = mainchain.syncBatches(batcherContract, filter)
	if err != nil {
		return err
	}

	err = mainchain.syncExecutionState(executorContract, opts)
	if err != nil {
		return err
	}

	mainchain.CurrentBlock = latestBlockNumber
	return nil
}
