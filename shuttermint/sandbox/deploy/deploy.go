// Deploy the ConfigContract and KeyBroadcastContract to ganache. This uses a hard-coded private
// key that available when ganache is started with the -d flag.

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	startBlockNumberOffset           = 30
	defaultGasLimit                  = 5000000
	defaultConfigChangeHeadsUpBlocks = 10
	ganacheKeyIdx                    = 9
	numKeypers                       = 3
)

var (
	key    *ecdsa.PrivateKey
	client *ethclient.Client
)

var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Helper tools to deploy and interact with the Shutter contracts",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		key = sandbox.GanacheKey(ganacheKeyIdx)

		cl, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}
		client = cl
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy all contracts",
	Run: func(cmd *cobra.Command, args []string) {
		deploy()
	},
}

var scheduleFlags struct {
	ConfigContractAddress string
	StartBatchIndex       int
	BatchSpan             int
	StartBlockNumber      int
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a new config",
	Run: func(cmd *cobra.Command, args []string) {
		configContractAddress := common.HexToAddress(scheduleFlags.ConfigContractAddress)
		if scheduleFlags.ConfigContractAddress != configContractAddress.Hex() {
			log.Fatalf("Invalid config contract address %s", scheduleFlags.ConfigContractAddress)
		}

		schedule(configContractAddress, scheduleFlags.StartBatchIndex, scheduleFlags.BatchSpan, scheduleFlags.StartBlockNumber)
	},
}

var getconfigFlags struct {
	ConfigContractAddress string
	ConfigIndex           int
}
var getconfigCmd = &cobra.Command{
	Use:   "getconfig",
	Short: "Get the config with the given index",
	Run: func(cmd *cobra.Command, args []string) {
		configContractAddress := common.HexToAddress(getconfigFlags.ConfigContractAddress)
		if getconfigFlags.ConfigContractAddress != configContractAddress.Hex() {
			log.Fatalf("Invalid config contract address %s", getconfigFlags.ConfigContractAddress)
		}

		getconfig(configContractAddress, getconfigFlags.ConfigIndex)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(getconfigCmd)

	initScheduleFlags()
	initGetconfigFlags()
}

func initScheduleFlags() {
	scheduleCmd.Flags().StringVarP(
		&scheduleFlags.ConfigContractAddress,
		"config-contract",
		"c",
		"",
		"address of config contract",
	)
	scheduleCmd.MarkFlagRequired("config-contract")

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBatchIndex,
		"start-batch-index",
		0,
		"the start batch index",
	)
	scheduleCmd.MarkFlagRequired("batch-span")

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.BatchSpan,
		"batch-span",
		0,
		"the batch span",
	)
	scheduleCmd.MarkFlagRequired("batch-span")

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBlockNumber,
		"start-block-number",
		0,
		"the start block number",
	)
	scheduleCmd.MarkFlagRequired("start-block-number")
}

func initGetconfigFlags() {
	getconfigCmd.Flags().StringVarP(
		&getconfigFlags.ConfigContractAddress,
		"config-contract",
		"c",
		"",
		"address of config contract",
	)
	getconfigCmd.MarkFlagRequired("config-contract")

	getconfigCmd.Flags().IntVarP(
		&getconfigFlags.ConfigIndex,
		"config-index",
		"i",
		0,
		"the config index",
	)
	getconfigCmd.MarkFlagRequired("config-index")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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

func deploy() {
	auth, err := makeAuth(client, sandbox.GanacheKey(ganacheKeyIdx))
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

	configAddress, tx, _, err := contract.DeployConfigContract(
		auth,
		client,
		big.NewInt(defaultConfigChangeHeadsUpBlocks))
	addTx()

	broadcastAddress, tx, _, err := contract.DeployKeyBroadcastContract(auth, client, configAddress)
	addTx()

	_, err = waitForTransactions(client, txs)
	if err != nil {
		panic(err)
	}
	fmt.Println("ConfigContract address:", configAddress.Hex())
	fmt.Println("KeyBroadcastContract address:", broadcastAddress.Hex())
}

func checkContractExists(configContractAddress common.Address) {
	code, err := client.CodeAt(context.Background(), configContractAddress, nil)
	if err != nil {
		panic(err)
	}
	if len(code) == 0 {
		log.Fatalf("No contract deployed at address %s", configContractAddress.Hex())
	}
}

func schedule(configContractAddress common.Address, startBatchIndex int, batchSpan int, startBlockNumber int) {
	auth, err := makeAuth(client, sandbox.GanacheKey(ganacheKeyIdx))
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

	checkContractExists(configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	if err != nil {
		panic(err)
	}

	tx, err = cc.NextConfigSetStartBatchIndex(auth, big.NewInt(int64(startBatchIndex)))
	addTx()

	tx, err = cc.NextConfigSetBatchSpan(auth, big.NewInt(int64(batchSpan)))
	addTx()

	tx, err = cc.NextConfigAddKeypers(auth, makeKeypers())
	addTx()

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	minStartBlockNumber := big.NewInt(0).Add(header.Number, big.NewInt(startBlockNumberOffset))
	startBlockNumberBig := big.NewInt(int64(startBlockNumber))
	if startBlockNumberBig.Cmp(minStartBlockNumber) < 1 {
		log.Fatalf(
			"Start block number %d is too close to current head %d (required offset %d)",
			startBlockNumber,
			header.Number,
			startBlockNumberOffset,
		)
	}
	tx, err = cc.NextConfigSetStartBlockNumber(auth, startBlockNumberBig)
	addTx()

	tx, err = cc.ScheduleNextConfig(auth)
	addTx()

	_, err = waitForTransactions(client, txs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("start block of config: %d\n", startBlockNumber)
}

func getconfig(configContractAddress common.Address, index int) {
	checkContractExists(configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	if err != nil {
		panic(err)
	}

	c, err := cc.GetConfigByIndex(nil, uint64(index))
	if err != nil {
		panic(err)
	}

	printConfig(c)
}

func printConfig(config contract.BatchConfig) {
	fmt.Printf("     Start batch index: %d\n", config.StartBatchIndex)
	fmt.Printf("    Start block number: %d\n", config.StartBlockNumber)
	fmt.Printf("            Batch span: %d\n", config.BatchSpan)
	fmt.Printf("           Num keypers: %d\n", len(config.Keypers))
	fmt.Printf("             Threshold: %d\n", config.Threshold)
	fmt.Printf("        BatchSizeLimit: %d\n", config.BatchSizeLimit)
	fmt.Printf("  TransactionSizeLimit: %d\n", config.TransactionSizeLimit)
	fmt.Printf("   TransactionGasLimit: %d\n", config.TransactionGasLimit)
	fmt.Printf("           FeeReceiver: %s\n", config.FeeReceiver.Hex())
	fmt.Printf("         TargetAddress: %s\n", config.TargetAddress.Hex())
	fmt.Printf("TargetFunctionSelector: %x\n", config.TargetFunctionSelector)
	fmt.Printf("      ExecutionTimeout: %d\n", config.ExecutionTimeout)
}
