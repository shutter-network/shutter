// Package txbatch is used to batch transactions for a main chain node
package txbatch

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/brainbot-com/shutter/shuttermint/medley"
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
	txHashes := []common.Hash{}
	for _, tx := range txbatch.transactions {
		txHashes = append(txHashes, tx.Hash())
	}

	return medley.WaitMinedMany(ctx, txbatch.Ethclient, txHashes)
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
