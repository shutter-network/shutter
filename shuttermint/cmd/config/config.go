package config

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/cmd/deploy"
	"github.com/shutter-network/shutter/shuttermint/contract"
	"github.com/shutter-network/shutter/shuttermint/medley"
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
	gasPrice       *big.Int
)

var configFlags struct {
	EthereumURL   string
	ContractsPath string
	OwnerKey      string
	GasPrice      string
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

	ConfigCmd.PersistentFlags().StringVar(
		&configFlags.GasPrice,
		"gas-price",
		"",
		"gas price in GWei (default: use suggested one)",
	)
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

	if configFlags.GasPrice != "" {
		gasPriceGWei, ok := new(big.Int).SetString(configFlags.GasPrice, 10)
		if !ok {
			return errors.Errorf("--gas-price: invalid gas price '%s'", configFlags.GasPrice)
		}
		if gasPriceGWei.Sign() < 0 {
			return errors.Errorf("--gas-price: must be non-negative")
		}
		gasPrice = new(big.Int).Mul(gasPriceGWei, big.NewInt(params.GWei))
	} else {
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return errors.Wrapf(err, "failed to get suggested gas price")
		}
	}

	return nil
}
