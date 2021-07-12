package config

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/contract"
	"github.com/shutter-network/shutter/shuttermint/medley/txbatch"
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
	setNextCmd.MarkPersistentFlagRequired("config")

	addOwnerKeyFlag(setNextCmd)
}

func chooseStartBlockAndBatch(ctx context.Context, client *ethclient.Client, cc *contract.ConfigContract) (uint64, uint64, error) {
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	numConfigs, err := cc.NumConfigs(callOpts)
	if err != nil {
		return 0, 0, errors.Errorf("failed to query number of configs")
	}

	var batchSpan uint64
	var startBlockNumber uint64
	var startBatchIndex uint64
	if numConfigs != 0 {
		config, err := cc.GetConfigByIndex(callOpts, numConfigs-1)
		if err != nil {
			return 0, 0, errors.Wrapf(err, "failed to query config %d", numConfigs-1)
		}
		batchSpan = config.BatchSpan
		startBlockNumber = config.StartBlockNumber
		startBatchIndex = config.StartBatchIndex
	} else {
		batchSpan = 0
		startBlockNumber = 0
		startBatchIndex = 0
	}

	headsUp, err := cc.ConfigChangeHeadsUpBlocks(callOpts)
	if err != nil {
		return 0, 0, errors.Wrap(err, "failed to query config change heads up blocks")
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, 0, errors.Wrap(err, "failed to get header")
	}
	minStartBlock := header.Number.Uint64() + headsUp + 20

	if batchSpan == 0 {
		return minStartBlock, startBatchIndex, nil
	}
	delta := minStartBlock - startBlockNumber
	numStartedBatches := (delta + batchSpan + 1) / batchSpan
	newStartBlockNumber := startBlockNumber + numStartedBatches*batchSpan
	newStartBatchIndex := startBatchIndex + numStartedBatches
	return newStartBlockNumber, newStartBatchIndex, nil
}

func setNext(ctx context.Context) error {
	batchConfig := contract.BatchConfig{}
	err := batchConfig.ReadJSONFile(setNextFlags.ConfigPath)
	if err != nil {
		return err
	}

	if batchConfig.StartBlockNumber == 0 {
		startBlock, startBatch, err := chooseStartBlockAndBatch(ctx, client, configContract)
		if err != nil {
			return err
		}
		batchConfig.StartBlockNumber = startBlock
		batchConfig.StartBatchIndex = startBatch
		log.Printf("Using StartBlockNumber %d, StartBatchIndex %d", startBlock, startBatch)
	}

	batch, err := txbatch.New(ctx, client, ownerKey)
	if err != nil {
		return err
	}
	batch.TransactOpts.GasPrice = gasPrice
	err = configContract.SetNextBatchConfig(ctx, batch, batchConfig)
	if err != nil {
		return err
	}

	_, err = batch.WaitMined(ctx)

	return err
}
