package config

import (
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/medley"
	"github.com/shutter-network/shutter/shuttermint/medley/txbatch"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule the next batch config",
	Long: `This command schedules the next batch config. The next batch config can be configured using the
'shuttermint config set-next' command and queried with 'shuttermint config query -i next'. If
--config is specified, this command will set the next config and schedule it.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		ctx := context.Background()
		err = processConfigFlags(ctx)
		if err != nil {
			return err
		}
		if setNextFlags.ConfigPath != "" {
			err = setNext(ctx)
			if err != nil {
				return err
			}
		}
		return schedule(ctx)
	},
}

func init() {
	addOwnerKeyFlag(scheduleCmd)
	scheduleCmd.PersistentFlags().StringVar(
		&setNextFlags.ConfigPath,
		"config",
		"",
		"path to the config JSON file",
	)
}

func schedule(ctx context.Context) error {
	o, err := txbatch.InitTransactOpts(ctx, client, ownerKey)
	if err != nil {
		return err
	}
	o.GasPrice = gasPrice

	tx, err := configContract.ScheduleNextConfig(o)
	if err != nil {
		return errors.WithMessage(err, "ScheduleNextConfig")
	}
	receipt, err := medley.WaitMined(ctx, client, tx.Hash())
	if err != nil {
		return err
	}

	if receipt.Status == 1 {
		return nil
	}

	from := crypto.PubkeyToAddress(ownerKey.PublicKey)
	reason := medley.GetRevertReason(ctx, client, from, tx, receipt.BlockNumber)
	return errors.Errorf("ScheduleNextConfig, tx %s: %s", receipt.TxHash.Hex(), reason)
}
