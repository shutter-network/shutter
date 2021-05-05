package config

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/medley"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

var setNextCmd = &cobra.Command{
	Use:   "set-next",
	Short: "Set the next config in order to schedule it later",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		err := processConfigFlags(ctx)
		if err != nil {
			return err
		}
		return setNext(ctx)
	},
}

var setNextFlags struct {
	ConfigPath string
}

func init() {
	setNextCmd.PersistentFlags().StringVar(
		&setNextFlags.ConfigPath,
		"config",
		"",
		"path to the config JSON file",
	)
	sandbox.MarkFlagRequired(setNextCmd, "config")

	addOwnerKeyFlag(setNextCmd)
}

func loadConfigJSON(path string) (*contract.BatchConfig, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &contract.BatchConfig{}
	err = json.Unmarshal(d, c)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal json")
	}

	return c, nil
}

func setNext(ctx context.Context) error {
	batchConfig, err := loadConfigJSON(setNextFlags.ConfigPath)
	if err != nil {
		return err
	}

	callOpts := &bind.CallOpts{Context: ctx}
	currentNextConfig, err := configContract.GetNextConfig(callOpts)
	if err != nil {
		return err
	}

	txs, err := sendSetNextTransactions(ctx, ownerKey, batchConfig, &currentNextConfig)
	if err != nil {
		return err
	}

	txHashes := []common.Hash{}
	for _, tx := range txs {
		txHashes = append(txHashes, tx.Hash())
	}
	_, err = medley.WaitMinedMany(ctx, client, txHashes)
	if err != nil {
		return err
	}

	return nil
}

func sendSetNextTransactions(ctx context.Context, key *ecdsa.PrivateKey, newNextConfig, currentNextConfig *contract.BatchConfig) ([]*types.Transaction, error) {
	o, err := sandbox.InitTransactOpts(ctx, client, key)
	if err != nil {
		return []*types.Transaction{}, err
	}

	var tx *types.Transaction
	var txs []*types.Transaction
	addTx := func(tx *types.Transaction) {
		txs = append(txs, tx)
		o.Nonce.SetInt64(o.Nonce.Int64() + 1)
	}

	cc := configContract // for brevity
	if newNextConfig.StartBatchIndex != currentNextConfig.StartBatchIndex {
		if tx, err = cc.NextConfigSetStartBatchIndex(o, newNextConfig.StartBatchIndex); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.StartBlockNumber != currentNextConfig.StartBlockNumber {
		if tx, err = cc.NextConfigSetStartBlockNumber(o, newNextConfig.StartBlockNumber); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.Threshold != currentNextConfig.Threshold {
		if tx, err = cc.NextConfigSetThreshold(o, newNextConfig.Threshold); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.BatchSpan != currentNextConfig.BatchSpan {
		if tx, err = cc.NextConfigSetBatchSpan(o, newNextConfig.BatchSpan); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.BatchSizeLimit != currentNextConfig.BatchSizeLimit {
		if tx, err = cc.NextConfigSetBatchSizeLimit(o, newNextConfig.BatchSizeLimit); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.TransactionSizeLimit != currentNextConfig.TransactionSizeLimit {
		if tx, err = cc.NextConfigSetTransactionSizeLimit(o, newNextConfig.TransactionSizeLimit); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.TransactionGasLimit != currentNextConfig.TransactionGasLimit {
		if tx, err = cc.NextConfigSetTransactionGasLimit(o, newNextConfig.TransactionGasLimit); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.FeeReceiver != currentNextConfig.FeeReceiver {
		if tx, err = cc.NextConfigSetFeeReceiver(o, newNextConfig.FeeReceiver); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.TargetAddress != currentNextConfig.TargetAddress {
		if tx, err = cc.NextConfigSetTargetAddress(o, newNextConfig.TargetAddress); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.TargetFunctionSelector != currentNextConfig.TargetFunctionSelector {
		if tx, err = cc.NextConfigSetTargetFunctionSelector(o, newNextConfig.TargetFunctionSelector); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if newNextConfig.ExecutionTimeout != currentNextConfig.ExecutionTimeout {
		if tx, err = cc.NextConfigSetExecutionTimeout(o, newNextConfig.ExecutionTimeout); err != nil {
			return txs, err
		}
		addTx(tx)
	}
	if !addressesEqual(newNextConfig.Keypers, currentNextConfig.Keypers) {
		// TODO: remove and keypers in groups if there are too many of them
		if tx, err = cc.NextConfigRemoveKeypers(o, uint64(len(currentNextConfig.Keypers))); err != nil {
			return txs, err
		}
		addTx(tx)
		if tx, err = cc.NextConfigAddKeypers(o, newNextConfig.Keypers); err != nil {
			return txs, err
		}
		addTx(tx)
	}

	return txs, nil
}

func addressesEqual(s1, s2 []common.Address) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, a1 := range s1 {
		if a1 != s2[i] {
			return false
		}
	}
	return true
}
