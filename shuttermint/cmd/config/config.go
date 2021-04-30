package config

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Send batch configs to and query them from Shutter's config contract",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		if cmd.PersistentFlags().Lookup("owner-key") != nil {
			err := parseOwnerKey()
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var (
	client         *ethclient.Client
	contractsJSON  *sandbox.ContractsJSON
	configContract *contract.ConfigContract
	ownerKey       *ecdsa.PrivateKey
)

var configFlags struct {
	EthereumURL   string
	ContractsPath string
	OwnerKey      string
}

func parseOwnerKey() error {
	key, err := crypto.HexToECDSA(configFlags.OwnerKey)
	if err != nil {
		return errors.Wrap(err, "parse -k / --owner-key flag")
	}
	ownerKey = key
	return nil
}

func addOwnerKeyFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&configFlags.OwnerKey,
		"owner-key",
		"k",
		"",
		"private key of the owner",
	)
	sandbox.MarkFlagRequired(cmd, "owner-key")
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

func processConfigFlags(ctx context.Context) error {
	var err error

	client, err = ethclient.DialContext(ctx, configFlags.EthereumURL)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Ethereum node at %s", configFlags.EthereumURL)
	}

	contractsJSON, err = sandbox.LoadContractsJSON(configFlags.ContractsPath)
	if err != nil {
		return errors.Wrapf(err, "failed to load contracts JSON file at %s", configFlags.ContractsPath)
	}

	configContract, err = contract.NewConfigContract(contractsJSON.ConfigContract, client)
	if err != nil {
		return err
	}

	return nil
}
