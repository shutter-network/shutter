package prepare

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/medley/txbatch"
)

var fundFlags struct {
	EthereumURL string
	OwnerKey    string
	ConfigPath  string
}

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Fund the accounts of an earlier prepared keyper set",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateFundFlags(); err != nil {
			return err
		}
		return fund()
	},
}

func initFundFlags() {
	fundCmd.Flags().StringVarP(
		&fundFlags.OwnerKey,
		"owner-key",
		"k",
		"",
		"private key of the funding account",
	)
	fundCmd.MarkFlagRequired("owner-key")

	fundCmd.Flags().StringVarP(
		&fundFlags.EthereumURL,
		"ethereum-url",
		"e",
		"",
		"Ethereum JSON RPC URL",
	)
	fundCmd.MarkFlagRequired("ethereum-url")

	fundCmd.PersistentFlags().StringVar(
		&fundFlags.ConfigPath,
		"config",
		"",
		"path to the config JSON file",
	)
	fundCmd.MarkPersistentFlagRequired("config")
}

func validateFundFlags() error {
	if err := validatePrivateKey(fundFlags.OwnerKey); err != nil {
		return errors.WithMessage(err, "invalid flag --owner-key")
	}

	return nil
}

func fund() error {
	bc := contract.BatchConfig{}
	err := bc.ReadJSONFile(fundFlags.ConfigPath)
	if err != nil {
		return err
	}
	keypers := bc.Keypers

	client, err := ethclient.DialContext(context.Background(), fundFlags.EthereumURL)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Ethereum node at %s", configFlags.EthereumURL)
	}

	return fundAddresses(context.Background(), client, keypers)
}

func fundAddresses(ctx context.Context, client *ethclient.Client, addresses []common.Address) error {
	ownerKey, err := crypto.HexToECDSA(fundFlags.OwnerKey)
	if err != nil {
		return errors.Errorf("invalid owner key")
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query chain id")
	}
	signer := types.NewEIP155Signer(chainID)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query gas price")
	}

	amount, ok := new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		panic("unexpected error")
	}

	batch, err := txbatch.New(ctx, client, ownerKey)
	if err != nil {
		return err
	}

	for _, addr := range addresses {
		unsignedTx := types.NewTransaction(batch.TransactOpts.Nonce.Uint64(), addr, amount, 21000, gasPrice, []byte{})
		tx, err := types.SignTx(unsignedTx, signer, ownerKey)
		if err != nil {
			return errors.Wrap(err, "failed to sign transaction")
		}

		err = client.SendTransaction(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "failed to send transaction")
		}
		batch.Add(tx)
	}
	_, err = batch.WaitMined(ctx)
	return err
}

func validatePrivateKey(key string) error {
	if _, err := crypto.HexToECDSA(key); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
