package prepare

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/shutter-network/shutter/shuttermint/keyper"
	"github.com/shutter-network/shutter/shuttermint/medley"
)

type EthSnatcher struct {
	gasPrice *big.Int
	chainID  *big.Int
	receiver common.Address
	signer   types.EIP155Signer
	client   *ethclient.Client
}

func (snatcher *EthSnatcher) Init(ctx context.Context, client *ethclient.Client) error {
	var err error
	snatcher.chainID, err = client.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query chain id")
	}

	snatcher.gasPrice, err = client.SuggestGasPrice(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query gas price")
	}
	snatcher.signer = types.NewEIP155Signer(snatcher.chainID)
	snatcher.client = client
	return nil
}

func (snatcher *EthSnatcher) Drain(ctx context.Context, key *ecdsa.PrivateKey) (*types.Transaction, error) {
	const gasLimit = 21000

	addr := crypto.PubkeyToAddress(key.PublicKey)
	nonce, err := snatcher.client.NonceAt(ctx, addr, nil)
	if err != nil {
		return nil, err
	}
	balance, err := snatcher.client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return nil, err
	}

	amount := big.NewInt(-gasLimit)
	amount.Mul(amount, snatcher.gasPrice)
	amount.Add(amount, balance)

	if amount.Sign() <= 0 {
		return nil, errors.Errorf("no funds to drain from %s", addr.Hex())
	}
	unsignedTx := types.NewTransaction(nonce, snatcher.receiver, amount, gasLimit, snatcher.gasPrice, []byte{})
	tx, err := types.SignTx(unsignedTx, snatcher.signer, key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction")
	}

	err = snatcher.client.SendTransaction(ctx, tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send transaction")
	}
	log.Printf("Draining %s, balance=%s, amount=%s, gasPrice=%s", addr, balance, amount, snatcher.gasPrice)
	return tx, nil
}

func readKeys(configFiles []string) ([]*ecdsa.PrivateKey, error) {
	keys := []*ecdsa.PrivateKey{}
	for _, path := range configFiles {
		k, err := readKeyFromConfig(path)
		if err != nil {
			return nil, err
		}
		keys = append(keys, k)
	}
	return keys, nil
}

var unfundCmd = &cobra.Command{
	Use:   "unfund",
	Short: "Drain the account of one or more keypers",
	RunE: func(cmd *cobra.Command, args []string) error {
		receiver, err := getReceiver()
		if err != nil {
			return err
		}
		keys, err := readKeys(args)
		if err != nil {
			return err
		}
		ctx := context.Background()
		client, err := ethclient.DialContext(ctx, unfundFlags.EthereumURL)
		if err != nil {
			return errors.Wrapf(err, "failed to connect to Ethereum node at %s", unfundFlags.EthereumURL)
		}
		snatcher := EthSnatcher{receiver: receiver}
		err = snatcher.Init(ctx, client)
		if err != nil {
			return err
		}

		txhashes := []common.Hash{}
		for _, k := range keys {
			var tx *types.Transaction
			tx, err = snatcher.Drain(ctx, k)
			if err != nil {
				log.Printf("Cannot drain keyper: %s", err)
				continue
			}
			txhashes = append(txhashes, tx.Hash())
		}
		_, err = medley.WaitMinedMany(ctx, client, txhashes)
		return err
	},
}

var unfundFlags struct {
	EthereumURL string
	Receiver    string
}

func getReceiver() (common.Address, error) {
	r := common.HexToAddress(unfundFlags.Receiver)
	if r.Hex() != unfundFlags.Receiver {
		return common.Address{}, errors.Errorf("bad receiver address: %s", unfundFlags.Receiver)
	}
	return r, nil
}

func readKeyFromConfig(path string) (*ecdsa.PrivateKey, error) {
	v := viper.New()
	v.SetConfigType("toml")
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = v.ReadConfig(f)
	if err != nil {
		return nil, err
	}
	config := keyper.Config{}
	err = config.Unmarshal(v)
	if err != nil {
		return nil, err
	}
	return config.SigningKey, nil
}

func init() {
	PrepareCmd.AddCommand(unfundCmd)
	unfundCmd.PersistentFlags().StringVarP(
		&unfundFlags.EthereumURL,
		"ethereum-url",
		"e",
		"",
		"Ethereum JSON RPC URL",
	)
	unfundCmd.MarkPersistentFlagRequired("ethereum-url")

	unfundCmd.PersistentFlags().StringVarP(
		&unfundFlags.Receiver,
		"to",
		"",
		"",
		"Receiver address",
	)
	unfundCmd.MarkPersistentFlagRequired("to")
}
