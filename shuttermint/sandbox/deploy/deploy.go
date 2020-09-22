// Deploy the ConfigContract and KeyBroadcastContract to ganache. This uses a hard-coded private
// key that available when ganache is started with the -d flag.

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	startBlockNumberOffset           = 30
	defaultGasLimit                  = 5000000
	defaultConfigChangeHeadsUpBlocks = 10
	ganacheKeyIdx                    = 9
	defaultBatchSpan                 = 5
	numKeypers                       = 3
)

func waitForTransactionReceipt(cl *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := cl.TransactionReceipt(context.Background(), txHash)
		if err == ethereum.NotFound {
			time.Sleep(time.Second)
			continue
		}
		return receipt, err
	}
}

func waitForTransactions(client *ethclient.Client, txs []*types.Transaction) ([]*types.Receipt, error) {
	defer fmt.Print("\n")
	var res []*types.Receipt
	for _, tx := range txs {
		receipt, err := waitForTransactionReceipt(client, tx.Hash())
		if err != nil {
			return res, err
		}
		res = append(res, receipt)
		if receipt.Status != 1 {
			fmt.Print("X")
		} else {
			fmt.Print(".")
		}
	}

	return res, nil
}

func makeKeypers() []common.Address {
	var keypers []common.Address
	for i := 0; i < numKeypers; i++ {
		keypers = append(keypers, crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey))
	}
	return keypers
}

func makeAuth(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = defaultGasLimit
	auth.GasPrice = gasPrice
	return auth, nil
}

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	auth, err := makeAuth(client, sandbox.GanacheKey(ganacheKeyIdx))
	if err != nil {
		panic(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	var txs []*types.Transaction
	var tx *types.Transaction

	addTx := func() {
		if err != nil {
			panic(err)
		}
		txs = append(txs, tx)
		auth.Nonce.SetInt64(auth.Nonce.Int64() + 1)
	}

	startBlockNumber := big.NewInt(startBlockNumberOffset)
	startBlockNumber.Add(startBlockNumber, header.Number)

	configAddress, tx, cc, err := contract.DeployConfigContract(
		auth,
		client,
		big.NewInt(defaultConfigChangeHeadsUpBlocks))
	addTx()

	tx, err = cc.NextConfigSetBatchSpan(auth, big.NewInt(defaultBatchSpan))
	addTx()

	tx, err = cc.NextConfigSetStartBatchIndex(auth, big.NewInt(0))
	addTx()

	tx, err = cc.NextConfigSetStartBlockNumber(auth, startBlockNumber)
	addTx()

	tx, err = cc.NextConfigAddKeypers(auth, makeKeypers())
	addTx()

	tx, err = cc.ScheduleNextConfig(auth)
	addTx()

	broadcastAddress, tx, _, err := contract.DeployKeyBroadcastContract(auth, client, configAddress)
	addTx()

	_, err = waitForTransactions(client, txs)
	if err != nil {
		panic(err)
	}
	fmt.Println("ConfigContract address:", configAddress.Hex())
	fmt.Println("KeyBroadcastContract address:", broadcastAddress.Hex())
	fmt.Printf("start block of config: %s\n", startBlockNumber)
}
