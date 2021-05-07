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
)

// ExitIfError exits the program with an error message and status if the given error is non-nil.
func ExitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
