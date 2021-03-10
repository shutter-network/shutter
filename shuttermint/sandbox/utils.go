package sandbox

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// MarkFlagRequired marks a cobra flag as required and panics if this fails.
func MarkFlagRequired(cmd *cobra.Command, name string) {
	if cmd.Flags().Lookup(name) != nil {
		if err := cmd.MarkFlagRequired(name); err != nil {
			panic(err)
		}
	}

	if cmd.PersistentFlags().Lookup(name) != nil {
		if err := cmd.MarkPersistentFlagRequired(name); err != nil {
			panic(err)
		}
	}
}

// ExitIfError exits the program with an error message and status if the given error is non-nil.
func ExitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ValidatePrivateKey checks that the argument is a valid hex-encoded private key.
func ValidatePrivateKey(key string) error {
	if _, err := crypto.HexToECDSA(key); err != nil {
		return errors.Wrap(err, "invalid private key")
	}
	return nil
}

// InitTransactOpts initializes the transact options struct.
func InitTransactOpts(ctx context.Context, client *ethclient.Client, key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}

	sender := crypto.PubkeyToAddress(key.PublicKey)
	nonce, err := client.PendingNonceAt(ctx, sender)
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))

	return opts, nil
}
