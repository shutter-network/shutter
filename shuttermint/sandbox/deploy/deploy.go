// Deploy the ConfigContract and KeyBroadcastContract to ganache. This uses a hard-coded private
// key that available when ganache is started with the -d flag.

package main

import (
	"context"
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
)

const (
	defaultGasLimit                  = 5000000
	defaultConfigChangeHeadsUpBlocks = 5
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

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// This is the last key used by 'ganache-cli -d'
	privateKey, err := crypto.HexToECDSA("b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	newAuth := func() *bind.TransactOpts {
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0) // in wei
		auth.GasLimit = defaultGasLimit
		auth.GasPrice = gasPrice
		return auth
	}

	configAddress, tx, _, err := contract.DeployConfigContract(
		newAuth(),
		client,
		big.NewInt(defaultConfigChangeHeadsUpBlocks))
	if err != nil {
		log.Fatal(err)
	}
	_, err = waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ConfigContract address:", configAddress.Hex())
	fmt.Println(tx.Hash().Hex())

	address, tx, _, err := contract.DeployKeyBroadcastContract(newAuth(), client, configAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nDeployKeyBroadcastContract address:", address.Hex())
	fmt.Println(tx.Hash().Hex())
	_, err = waitForTransactionReceipt(client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
}
