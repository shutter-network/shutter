package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/contract"
	"github.com/shutter-network/shutter/shuttermint/sandbox"
)

const (
	ganacheKeyIdx = 8

	txSize     = 2
	txInterval = 2000 * time.Millisecond

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
	err := sendtxsCmd.MarkFlagRequired("batcher-contract")
	if err != nil {
		log.Fatal(err)
	}
}

func initKey() {
	key = sandbox.GanacheKey(ganacheKeyIdx)
}

func initClient() {
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	cl, err := ethclient.DialContext(ctx, "http://localhost:8545")
	cancel()
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
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Printf("failed to query chain ID: %+v", err)
		return
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Printf("failed to create transactor: %+v", err)
		return
	}

	transactOpts.Context = ctx
	transactOpts.GasLimit = 100000

	lastTxTime := time.Now()

	for {
		dt := txInterval - time.Since(lastTxTime)
		time.Sleep(dt)
		lastTxTime = time.Now()

		cipherTx := makeTx()

		blockNumber, err := client.BlockNumber(ctx)
		if err != nil {
			log.Printf("failed to determine current block number: %+v", err)
			continue
		}
		nextBatchIndex, err := configContract.NextBatchIndex(blockNumber)
		if err != nil {
			log.Printf("failed to determine batch index for next tx: %+v", err)
			continue
		}
		if nextBatchIndex == 0 {
			log.Printf("Waiting for first batch to start (block %d)", blockNumber)
			time.Sleep(1 * time.Second)
			continue
		}
		batchIndex := nextBatchIndex - 1

		txTypeBig, err := rand.Int(rand.Reader, big.NewInt(contract.TransactionTypePlain+1))
		if err != nil {
			panic("failed to generate random number")
		}
		txType := uint8(txTypeBig.Uint64())

		tx, err := batcherContract.AddTransaction(transactOpts, batchIndex, txType, cipherTx)
		if err != nil {
			log.Printf("failed to send tx: %+v", err)
			continue
		}
		log.Printf("Sent tx %s of type %d for batch %d in tx %s", hexutil.Encode(cipherTx), txType, batchIndex, tx.Hash().Hex())

		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			log.Printf("failed to wait for receipt of tx %+v: %v", tx.Hash().Hex(), err)
			continue
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Printf("tx failed")
		}
	}
}

func makeTx() []byte {
	b := make([]byte, txSize)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func main() {
	if err := sendtxsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
