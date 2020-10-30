// Deploy the ConfigContract and KeyBroadcastContract to ganache. This uses a hard-coded private
// key that available when ganache is started with the -d flag.

package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
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
	defaultConfigChangeHeadsUpBlocks = 20
	ganacheKeyIdx                    = 9
	numKeypers                       = 3
	threshold                        = 2
	dialDefaultTimeout               = 5 * time.Second
	getconfigDefaultTimeout          = 10 * time.Second
	deployDefaultTimeout             = 30 * time.Second
	scheduleDefaultTimeout           = 60 * time.Second
)

var (
	key    *ecdsa.PrivateKey
	client *ethclient.Client
)

var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Helper tools to deploy and interact with the Shutter contracts",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), dialDefaultTimeout)
		defer cancel()
		key = sandbox.GanacheKey(ganacheKeyIdx)

		cl, err := ethclient.DialContext(ctx, "http://localhost:8545")
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
		ctx, cancel := context.WithTimeout(context.Background(), deployDefaultTimeout)
		defer cancel()
		deploy(ctx)
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

		ctx, cancel := context.WithTimeout(context.Background(), scheduleDefaultTimeout)
		defer cancel()

		schedule(
			ctx,
			configContractAddress,
			uint64(scheduleFlags.StartBatchIndex),
			uint64(scheduleFlags.BatchSpan),
			uint64(scheduleFlags.StartBlockNumber),
		)
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
		ctx, cancel := context.WithTimeout(context.Background(), getconfigDefaultTimeout)
		defer cancel()
		getconfig(ctx, configContractAddress, getconfigFlags.ConfigIndex)
	},
}

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Fund accounts",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), getconfigDefaultTimeout)
		defer cancel()
		fund(ctx)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(getconfigCmd)
	rootCmd.AddCommand(fundCmd)

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
	err := scheduleCmd.MarkFlagRequired("config-contract")
	if err != nil {
		panic(err)
	}

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBatchIndex,
		"start-batch-index",
		0,
		"the start batch index",
	)
	err = scheduleCmd.MarkFlagRequired("batch-span")
	if err != nil {
		panic(err)
	}

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.BatchSpan,
		"batch-span",
		0,
		"the batch span",
	)
	err = scheduleCmd.MarkFlagRequired("batch-span")
	if err != nil {
		panic(err)
	}

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBlockNumber,
		"start-block-number",
		0,
		"the start block number",
	)
	err = scheduleCmd.MarkFlagRequired("start-block-number")
	if err != nil {
		panic(err)
	}
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

func waitForTransactionReceipt(ctx context.Context, cl *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := cl.TransactionReceipt(ctx, txHash)
		if err == ethereum.NotFound {
			time.Sleep(time.Second)
			continue
		}
		return receipt, err
	}
}

func waitForTransactions(ctx context.Context, client *ethclient.Client, txs []*types.Transaction) ([]*types.Receipt, error) {
	defer fmt.Print("\n")
	var res []*types.Receipt

	failedTxs := []int{}
	for i, tx := range txs {
		receipt, err := waitForTransactionReceipt(ctx, client, tx.Hash())
		if err != nil {
			return res, err
		}
		res = append(res, receipt)
		if receipt.Status != 1 {
			fmt.Print("X")
			failedTxs = append(failedTxs, i)
		} else {
			fmt.Print(".")
		}
	}

	if len(failedTxs) > 0 {
		return res, errors.New("Some txs have failed")
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

func makeAuth(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
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

func deploy(ctx context.Context) {
	auth, err := makeAuth(ctx, client, key)
	if err != nil {
		panic(err)
	}
	auth.Context = ctx
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
		defaultConfigChangeHeadsUpBlocks,
	)
	addTx()

	broadcastAddress, tx, _, err := contract.DeployKeyBroadcastContract(auth, client, configAddress)
	addTx()

	feeAddress, tx, _, err := contract.DeployFeeBankContract(auth, client)
	addTx()

	batcherAddress, tx, _, err := contract.DeployBatcherContract(auth, client, configAddress, feeAddress)
	addTx()

	executorAddress, tx, _, err := contract.DeployExecutorContract(auth, client, configAddress, batcherAddress)
	addTx()

	_, err = waitForTransactions(ctx, client, txs)
	if err != nil {
		panic(err)
	}
	fmt.Println("ConfigContract address:", configAddress.Hex())
	fmt.Println("KeyBroadcastContract address:", broadcastAddress.Hex())
	fmt.Println("FeeBankContract address:", feeAddress.Hex())
	fmt.Println("BatcherContract address:", batcherAddress.Hex())
	fmt.Println("ExecutorContract address:", executorAddress.Hex())
}

func checkContractExists(ctx context.Context, configContractAddress common.Address) {
	code, err := client.CodeAt(ctx, configContractAddress, nil)
	if err != nil {
		panic(err)
	}
	if len(code) == 0 {
		log.Fatalf("No contract deployed at address %s", configContractAddress.Hex())
	}
}

func schedule(
	ctx context.Context,
	configContractAddress common.Address,
	startBatchIndex, batchSpan, startBlockNumber uint64,
) {
	auth, err := makeAuth(ctx, client, sandbox.GanacheKey(ganacheKeyIdx))
	if err != nil {
		panic(err)
	}
	auth.Context = ctx

	var txs []*types.Transaction
	var tx *types.Transaction

	addTx := func() {
		if err != nil {
			panic(err)
		}
		txs = append(txs, tx)
		auth.Nonce.SetInt64(auth.Nonce.Int64() + 1)
	}

	checkContractExists(ctx, configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	if err != nil {
		panic(err)
	}

	tx, err = cc.NextConfigSetStartBatchIndex(auth, startBatchIndex)
	addTx()

	tx, err = cc.NextConfigSetBatchSpan(auth, batchSpan)
	addTx()

	tx, err = cc.NextConfigAddKeypers(auth, makeKeypers())
	addTx()

	tx, err = cc.NextConfigSetThreshold(auth, threshold)
	addTx()

	tx, err = cc.NextConfigSetExecutionTimeout(auth, batchSpan)
	addTx()

	tx, err = cc.NextConfigSetTransactionSizeLimit(auth, 100)
	addTx()

	tx, err = cc.NextConfigSetBatchSizeLimit(auth, 100*100)
	addTx()

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		panic(err)
	}
	minStartBlockNumber := header.Number.Uint64() + startBlockNumberOffset
	if startBlockNumber < minStartBlockNumber {
		log.Fatalf(
			"Start block number %d is too close to current head %d (required offset %d)",
			startBlockNumber,
			header.Number,
			startBlockNumberOffset,
		)
	}
	tx, err = cc.NextConfigSetStartBlockNumber(auth, startBlockNumber)
	addTx()

	tx, err = cc.ScheduleNextConfig(auth)
	addTx()

	_, err = waitForTransactions(ctx, client, txs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("start block of config: %d\n", startBlockNumber)
}

func getconfig(ctx context.Context, configContractAddress common.Address, index int) {
	checkContractExists(ctx, configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	if err != nil {
		panic(err)
	}
	c, err := cc.GetConfigByIndex(&bind.CallOpts{Context: ctx}, uint64(index))
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
	fmt.Printf("\nKeypers:\n")
	if len(config.Keypers) == 0 {
		fmt.Printf("  None\n")
	} else {
		for i, k := range config.Keypers {
			fmt.Printf("  %d: %s\n", i, k.Hex())
		}
	}
}

func fund(ctx context.Context) {
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	var txs []*types.Transaction
	var tx *types.Transaction

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		panic(err)
	}

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		panic(err)
	}

	amount := big.NewInt(1000000000000000000) // 1 eth
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	addTx := func() {
		if err != nil {
			panic(err)
		}
		txs = append(txs, tx)
		nonce++
	}

	signer := types.NewEIP155Signer(chainID)

	for i := 0; i < sandbox.NumGanacheKeys()-1; i++ {
		receiver := crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey)
		var data []byte
		tx = types.NewTransaction(nonce, receiver, amount, gasLimit, gasPrice, data)

		tx, err = types.SignTx(tx, signer, key)
		if err != nil {
			panic(err)
		}

		err = client.SendTransaction(ctx, tx)
		if err != nil {
			panic(err)
		}
		addTx()
	}

	_, err = waitForTransactions(ctx, client, txs)
	if err != nil {
		panic(err)
	}

	for i := 0; i < sandbox.NumGanacheKeys(); i++ {
		addr := crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey)
		balance, err := client.BalanceAt(context.Background(), addr, nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %s\n", addr.Hex(), balance)
	}
}
