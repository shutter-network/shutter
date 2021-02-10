package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configs",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		sandbox.ExitIfError(processRootFlags(ctx))
		sandbox.ExitIfError(list(ctx))
	},
}

func initListFlags() {
	// no flags
}

func list(ctx context.Context) error {
	currentBlock, err := client.BlockNumber(ctx)
	if err != nil {
		return err
	}
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: new(big.Int).SetUint64(currentBlock),
	}

	numConfigs, err := configContract.NumConfigs(opts)
	if err != nil {
		return err
	}

	configs := []*contract.BatchConfig{}
	for i := uint64(0); i < numConfigs; i++ {
		c, err := configContract.GetConfigByIndex(opts, i)
		if err != nil {
			return err
		}
		configs = append(configs, &c)
	}

	nextConfig, err := configContract.GetNextConfig(opts)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	fmt.Fprintf(w, "Index\tStart Batch Index \tStartBlockNumber\n")
	for i, c := range configs {
		fmt.Fprintf(w, "%d\t%d\t%d\t\n", i, c.StartBatchIndex, c.StartBlockNumber)
	}
	fmt.Fprintf(w, "Next\t%d\t%d\t\n", nextConfig.StartBatchIndex, nextConfig.StartBlockNumber)
	w.Flush()

	return nil
}
