package config

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const keyFlagName = "key"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure and query the Shutter config contract",
}

var (
	client         *ethclient.Client
	contractsJSON  *sandbox.ContractsJSON
	configContract *contract.ConfigContract
)

var configFlags struct {
	EthereumURL   string
	ContractsPath string
}

func init() {
	ConfigCmd.AddCommand(queryCmd)
	ConfigCmd.AddCommand(setNextCmd)
	ConfigCmd.AddCommand(scheduleCmd)
	ConfigCmd.AddCommand(listCmd)

	initConfigRootFlags()
}

func initConfigRootFlags() {
	ConfigCmd.PersistentFlags().StringVarP(
		&configFlags.EthereumURL,
		"ethereum-url",
		"e",
		"",
		"Ethereum JSON RPC URL",
	)

	ConfigCmd.PersistentFlags().StringVarP(
		&configFlags.ContractsPath,
		"contracts",
		"c",
		"",
		"path to the contracts.json file",
	)
	sandbox.MarkFlagRequired(ConfigCmd, "contracts")
}

func validateConfigFlags() (string, error) {
	return "", nil
}

func processConfigFlags(ctx context.Context) error {
	if flag, err := validateConfigFlags(); err != nil {
		return errors.Wrapf(err, "invalid flag %s", flag)
	}

	var err error

	client, err = ethclient.DialContext(ctx, configFlags.EthereumURL)
	if err != nil {
		return errors.Wrapf(err, "faild to connect to Ethereum node at %s", configFlags.EthereumURL)
	}

	contractsJSON, err = sandbox.LoadContractsJSON(configFlags.ContractsPath)
	if err != nil {
		return errors.Wrapf(err, "failed to load contracts JSON file at %s", configFlags.ContractsPath)
	}

	configContract, err = contract.NewConfigContract(common.HexToAddress(contractsJSON.ConfigContract), client)
	if err != nil {
		return err
	}

	return nil
}
