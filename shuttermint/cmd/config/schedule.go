package config

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/medley"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule the next batch config",
	Long: `This command schedules the next batch config. The next batch config can be
configured using the 'shuttermint config set-next' command and queried with
'shuttermint config query -i next'.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		err := processConfigFlags(ctx)
		if err != nil {
			return err
		}
		return schedule(ctx)
	},
}

func init() {
	addOwnerKeyFlag(scheduleCmd)
}

func schedule(ctx context.Context) error {
	o, err := sandbox.InitTransactOpts(ctx, client, ownerKey)
	if err != nil {
		return err
	}

	tx, err := configContract.ScheduleNextConfig(o)
	if err != nil {
		return err
	}
	_, err = medley.WaitMined(ctx, client, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}
