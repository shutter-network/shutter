package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	ganacheKeyIdx = 8

	txSize     = 2
	txInterval = 1000 * time.Millisecond

	dialTimeout = 5 * time.Second
)

var (
	key             *ecdsa.PrivateKey
	client          *ethclient.Client
	batcherContract *contract.BatcherContract
	configContract  *contract.ConfigContract

	batcherContractAddressFlag string
)

var sendtxsCmd = &cobra.Command{
	Use:   "sendtxs",
	Short: "A tool to simulate a user sending random transactions",
	Run: func(cmd *cobra.Command, args []string) {
		sendtxs()
	},
}

func init() {
	initFlags()
}

func initFlags() {
	sendtxsCmd.Flags().StringVarP(
		&batcherContractAddressFlag,
		"batcher-contract",
		"b",
		"",
		"address of batcher contract",
	)
	sendtxsCmd.MarkFlagRequired("batcher-contract")
}

func initKey() {
	key = sandbox.GanacheKey(ganacheKeyIdx)
}

func initClient() {
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	cl, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	client = cl
}

func initBatcherContract() {
	address := common.HexToAddress(batcherContractAddressFlag)
	if batcherContractAddressFlag != address.Hex() {
		log.Fatalf("invalid batcher contract address %s", batcherContractAddressFlag)
	}

	bc, err := contract.NewBatcherContract(address, client)
	if err != nil {
		log.Fatalf("failed to initialize batcher contract: %v", err)
	}
	batcherContract = bc
}

func initConfigContract() {
	address, err := batcherContract.ConfigContract(nil)
	if err != nil {
		log.Fatalf("failed to retrieve config contract address: %v", err)
	}

	cc, err := contract.NewConfigContract(address, client)
	if err != nil {
		log.Fatalf("failed to initialize config contract: %v", err)
	}
	configContract = cc
}

func sendtxs() {
	initKey()
	initClient()
	initBatcherContract()
	initConfigContract()

	ctx := context.Background()
	transactOpts := bind.NewKeyedTransactor(key)
	transactOpts.Context = ctx
	transactOpts.GasLimit = 100000

	for {
		time.Sleep(txInterval)

		cipherTx := makeTxs()

		blockNumber, err := client.BlockNumber(ctx)
		if err != nil {
			log.Printf("failed to determine current block number: %s", err)
			continue
		}
		batchIndex, err := configContract.NextBatchIndex(blockNumber)
		if err != nil {
			log.Printf("failed to determine batch index for next tx: %v", err)
			continue
		}

		tx, err := batcherContract.AddTransaction(transactOpts, batchIndex, contract.TransactionTypeCipher, cipherTx)
		if err != nil {
			log.Printf("failed to send tx: %v", err)
			continue
		}

		log.Printf("Sent cipher tx %s for batch %d in tx %s", hexutil.Encode(cipherTx), batchIndex, tx.Hash().Hex())
	}
}

func makeTxs() []byte {
	b := make([]byte, txSize)
	rand.Read(b)
	return b
}

func main() {
	if err := sendtxsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
