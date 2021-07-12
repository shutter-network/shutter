// Package txbatch is used to batch transactions for a main chain node
package txbatch

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shuttermint/medley"
)

type TXBatch struct {
	Ethclient    *ethclient.Client
	TransactOpts *bind.TransactOpts

	key          *ecdsa.PrivateKey
	transactions []*types.Transaction
}

func New(ctx context.Context, client *ethclient.Client, key *ecdsa.PrivateKey) (*TXBatch, error) {
	opts, err := InitTransactOpts(ctx, client, key)
	if err != nil {
		return nil, err
	}
	return &TXBatch{
		Ethclient:    client,
		TransactOpts: opts,
		key:          key,
		transactions: nil,
	}, nil
}

func (txbatch *TXBatch) Add(tx *types.Transaction) {
	txbatch.transactions = append(txbatch.transactions, tx)
	txbatch.TransactOpts.Nonce.SetInt64(txbatch.TransactOpts.Nonce.Int64() + 1)
}

func (txbatch *TXBatch) WaitMined(ctx context.Context) ([]*types.Receipt, error) {
	defer fmt.Print("\n")
	var res []*types.Receipt
	numFailed := 0
	for i, tx := range txbatch.transactions {
		receipt, err := medley.WaitMined(ctx, txbatch.Ethclient, tx.Hash())
		if err != nil {
			return res, err
		}
		res = append(res, receipt)
		if receipt.Status != 1 {
			err = medley.GetRevertReason(ctx, txbatch.Ethclient, crypto.PubkeyToAddress(txbatch.key.PublicKey), tx, receipt.BlockNumber)
			fmt.Print("\n")
			log.Printf("tx #%d %s reverted: %s", i, tx.Hash(), err)
			numFailed++
		} else {
			fmt.Print(".")
		}
	}
	if numFailed > 0 {
		return res, errors.Errorf("%d transactions failed.", numFailed)
	}

	return res, nil
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
