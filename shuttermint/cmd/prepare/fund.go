package prepare

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/contract"
	"github.com/shutter-network/shutter/shuttermint/keyper/gaspricer"
	"github.com/shutter-network/shutter/shuttermint/medley/txbatch"
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

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query balances",
	RunE: func(cmd *cobra.Command, args []string) error {
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
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
		defer w.Flush()

		fmt.Fprintf(w, "#\tAddress\tBalance (Eth)\n")

		ctx := context.Background()
		sum := big.NewInt(0)
		for i, k := range keypers {
			balance, err := client.BalanceAt(ctx, k, nil)
			if err != nil {
				return err
			}

			fmt.Fprintf(w, "%d\t%s\t%f\n", i, k.Hex(), weiToEther(balance))
			sum.Add(sum, balance)
		}
		fmt.Fprintf(w, "---\t---\t---\n")
		fmt.Fprintf(w, "\t\t%f\n", weiToEther(sum))

		return nil
	},
}

func init() {
	PrepareCmd.AddCommand(fundCmd)
	fundCmd.AddCommand(queryCmd)
	fundCmd.Flags().StringVarP(
		&fundFlags.OwnerKey,
		"owner-key",
		"k",
		"",
		"private key of the funding account",
	)
	fundCmd.MarkFlagRequired("owner-key")

	fundCmd.PersistentFlags().StringVarP(
		&fundFlags.EthereumURL,
		"ethereum-url",
		"e",
		"",
		"Ethereum JSON RPC URL",
	)
	fundCmd.MarkPersistentFlagRequired("ethereum-url")

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

type funding struct {
	receiver common.Address
	amount   *big.Int
	balance  *big.Int
}

func determineFundings(
	ctx context.Context,
	client *ethclient.Client,
	addresses []common.Address,
	lowAmount, highAmount *big.Int) ([]funding, error) {
	res := []funding{}
	for _, addr := range addresses {
		balance, err := client.BalanceAt(ctx, addr, nil)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get balance")
		}
		amount := new(big.Int)

		if lowAmount.Cmp(balance) > 0 {
			amount.Sub(highAmount, balance)
		}

		res = append(res, funding{receiver: addr, amount: amount, balance: balance})
	}
	return res, nil
}

func report(fundings []funding) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "Receiver\tBalance\tFund\n")
	for _, f := range fundings {
		fmt.Fprintf(w, "%s\t%f\t%f\n", f.receiver.Hex(), weiToEther(f.balance), weiToEther(f.amount))
	}
}

func fundAddresses(ctx context.Context, client *ethclient.Client, addresses []common.Address) error {
	highAmount := big.NewInt(params.Ether)
	lowAmount := big.NewInt(params.Ether / 2)

	fundings, err := determineFundings(ctx, client, addresses, lowAmount, highAmount)
	if err != nil {
		return err
	}
	report(fundings)
	ownerKey, err := crypto.HexToECDSA(fundFlags.OwnerKey)
	if err != nil {
		return errors.Errorf("invalid owner key")
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query chain id")
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query gas price")
	}
	gasPrice = gaspricer.Adjust(gasPrice)

	signer := types.NewEIP155Signer(chainID)

	batch, err := txbatch.New(ctx, client, ownerKey)
	if err != nil {
		return err
	}

	for _, f := range fundings {
		if f.amount.Sign() <= 0 {
			continue
		}
		unsignedTx := types.NewTransaction(batch.TransactOpts.Nonce.Uint64(), f.receiver, f.amount, 21000, gasPrice, []byte{})
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
