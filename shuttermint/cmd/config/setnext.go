package config

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/medley/txbatch"
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

	batch, err := txbatch.New(ctx, client, ownerKey)
	if err != nil {
		return err
	}
	err = configContract.SetNextBatchConfig(ctx, batch, *batchConfig)
	if err != nil {
		return err
	}

	_, err = batch.WaitMined(ctx)

	return err
}
