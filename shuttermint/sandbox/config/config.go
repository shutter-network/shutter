package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const keyFlagName = "key"

var rootCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure and query the Shutter config contract",
}

var (
	client         *ethclient.Client
	contractsJSON  *sandbox.ContractsJSON
	configContract *contract.ConfigContract
)

var rootFlags struct {
	EthereumURL   string
	ContractsPath string
}

func init() {
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(setNextCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(listCmd)

	initRootFlags()
	initQueryFlags()
	initSetNextFlags()
	initScheduleFlags()
	initListFlags()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initRootFlags() {
	rootCmd.PersistentFlags().StringVarP(
		&rootFlags.EthereumURL,
		"ethereum-url",
		"e",
		"ws://localhost:8546/",
		"Ethereum JSON RPC URL",
	)

	rootCmd.PersistentFlags().StringVarP(
		&rootFlags.ContractsPath,
		"contracts",
		"c",
		"",
		"path to the contracts.json file",
	)
	sandbox.MarkFlagRequired(rootCmd, "contracts")
}

func validateRootFlags() (string, error) {
	return "", nil
}

func processRootFlags(ctx context.Context) error {
	if flag, err := validateRootFlags(); err != nil {
		return errors.Wrapf(err, "invalid flag %s", flag)
	}

	var err error

	client, err = ethclient.DialContext(ctx, rootFlags.EthereumURL)
	if err != nil {
		return errors.Wrapf(err, "faild to connect to Ethereum node at %s", rootFlags.EthereumURL)
	}

	contractsJSON, err = sandbox.LoadContractsJSON(rootFlags.ContractsPath)
	if err != nil {
		return errors.Wrapf(err, "failed to load contracts JSON file at %s", rootFlags.ContractsPath)
	}

	configContract, err = contract.NewConfigContract(common.HexToAddress(contractsJSON.ConfigContract), client)
	if err != nil {
		return err
	}

	return nil
}
