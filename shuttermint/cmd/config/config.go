package config

import (
	"context"
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/cmd/deploy"
	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/medley"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Send batch configs to and query them from Shutter's config contract",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Root().PersistentPreRunE(cmd, args)
		if err != nil {
			return err
		}
		return parseOwnerKey(cmd)
	},

	// We want to bail out in PreRun, because otherwise cobra checks for required arguments. We
	// need to set 'Run' here, otherwise PreRun is not called by cobra.
	PreRun: medley.ShowHelpAndExit,
	Run:    medley.ShowHelpAndExit,
}

var (
	client         *ethclient.Client
	configContract *contract.ConfigContract
	ownerKey       *ecdsa.PrivateKey
)

var configFlags struct {
	EthereumURL   string
	ContractsPath string
	OwnerKey      string
}

func parseOwnerKey(cmd *cobra.Command) error {
	pflag := cmd.PersistentFlags().Lookup("owner-key")
	if pflag == nil || !pflag.Changed {
		return nil
	}

	val := pflag.Value.String()
	key, err := crypto.HexToECDSA(strings.TrimPrefix(val, "0x"))
	if err != nil {
		return errors.Wrapf(err, "parse -k / --owner-key '%s'", val)
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
	cmd.MarkPersistentFlagRequired("owner-key")
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

	ConfigCmd.MarkPersistentFlagRequired("contracts")
}

func processConfigFlags(ctx context.Context) error {
	var err error

	client, err = ethclient.DialContext(ctx, configFlags.EthereumURL)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Ethereum node at %s", configFlags.EthereumURL)
	}

	contractsJSON, err := deploy.LoadContractsJSON(configFlags.ContractsPath)
	if err != nil {
		return errors.Wrapf(err, "failed to load contracts JSON file at %s", configFlags.ContractsPath)
	}

	configContract, err = contract.NewConfigContract(contractsJSON.ConfigContract, client)
	if err != nil {
		return err
	}

	return nil
}
