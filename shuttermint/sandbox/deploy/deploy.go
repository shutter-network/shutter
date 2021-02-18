// Deploy the ConfigContract and KeyBroadcastContract to ganache. This uses a hard-coded private
// key that available when ganache is started with the -d flag.

package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/contract/erc1820"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	startBlockNumberOffset           = 30
	defaultConfigChangeHeadsUpBlocks = 20
	defaultAppealBlocks              = 20
	ganacheKeyIdx                    = 9
	numKeypers                       = 3
	threshold                        = 2
	transactionSizeLimit             = 100
	batchSizeLimit                   = 100 * 100
	baseGasLimit                     = 21000
	fundAmount                       = 1 * params.Ether
	dialDefaultTimeout               = 5 * time.Second
	getconfigDefaultTimeout          = 10 * time.Second
	deployDefaultTimeout             = 300 * time.Second
	scheduleDefaultTimeout           = 600 * time.Second
)

var (
	key      *ecdsa.PrivateKey
	client   *ethclient.Client
	gasPrice *big.Int
)

var rootFlags struct {
	Key         string
	GasPrice    string
	EthereumURL string
}

var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Helper tools to deploy and interact with the Shutter contracts",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), dialDefaultTimeout)
		defer cancel()
		var err error

		key, err = crypto.HexToECDSA(strings.TrimPrefix(rootFlags.Key, "0x"))
		failIfError(err)

		cl, err := ethclient.DialContext(ctx, rootFlags.EthereumURL)
		failIfError(err)
		client = cl

		if rootFlags.GasPrice == "" {
			gasPrice, err = client.SuggestGasPrice(ctx)
			failIfError(err)
		} else {
			gasPriceGWei, ok := new(big.Int).SetString(rootFlags.GasPrice, 10)
			if !ok {
				log.Fatalf("Invalid gas price %s", rootFlags.GasPrice)
			}
			gasPrice = gweiToWei(gasPriceGWei)
		}

		address := crypto.PubkeyToAddress(key.PublicKey)
		balance, err := client.BalanceAt(ctx, address, nil)
		failIfError(err)

		log.Printf("Deploy Address: %s", address)
		log.Printf("Balance: %f ETH", weiToEther(balance))
		log.Printf("Gas Price: %f GWei", weiToGwei(gasPrice))
		log.Printf("Available gas: %d", new(big.Int).Quo(balance, gasPrice))
	},
}

var deployFlags struct {
	NoERC1820  bool
	OutputFile string
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

	initRootFlags()
	initDeployFlags()
	initScheduleFlags()
	initGetconfigFlags()
}

func initRootFlags() {
	rootCmd.PersistentFlags().StringVarP(
		&rootFlags.Key,
		"key",
		"k",
		hexutil.Encode(crypto.FromECDSA(sandbox.GanacheKey(ganacheKeyIdx))),
		"private key of deployer account",
	)
	rootCmd.PersistentFlags().StringVar(
		&rootFlags.GasPrice,
		"gas-price",
		"",
		"gas price in GWei (default: use suggested one)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&rootFlags.EthereumURL,
		"ethereum-url",
		"e",
		"ws://localhost:8545/websocket",
		"Ethereum RPC URL",
	)
}

func initDeployFlags() {
	deployCmd.Flags().BoolVar(
		&deployFlags.NoERC1820,
		"no-erc1820",
		false,
		"don't deploy the ERC1820 contract",
	)
	deployCmd.Flags().StringVarP(
		&deployFlags.OutputFile,
		"output",
		"o",
		"",
		"if given, store the contract addresses as a JSON file at this path",
	)
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
	failIfError(err)

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBatchIndex,
		"start-batch-index",
		0,
		"the start batch index",
	)
	err = scheduleCmd.MarkFlagRequired("start-batch-index")
	failIfError(err)

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.BatchSpan,
		"batch-span",
		0,
		"the batch span",
	)
	err = scheduleCmd.MarkFlagRequired("batch-span")
	failIfError(err)

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBlockNumber,
		"start-block-number",
		0,
		"the start block number",
	)
	err = scheduleCmd.MarkFlagRequired("start-block-number")
	failIfError(err)
}

func initGetconfigFlags() {
	getconfigCmd.Flags().StringVarP(
		&getconfigFlags.ConfigContractAddress,
		"config-contract",
		"c",
		"",
		"address of config contract",
	)
	err := getconfigCmd.MarkFlagRequired("config-contract")
	failIfError(err)

	getconfigCmd.Flags().IntVarP(
		&getconfigFlags.ConfigIndex,
		"config-index",
		"i",
		0,
		"the config index",
	)
	err = getconfigCmd.MarkFlagRequired("config-index")
	failIfError(err)
}

func main() {
	failIfError(rootCmd.Execute())
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
		log.Printf("The following transactions have failed:")
		for _, i := range failedTxs {
			log.Printf("%s", txs[i].Hash().Hex())
		}
		return res, errors.New("some txs have failed")
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

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	return auth, nil
}

func maybeDeployERC1820(ctx context.Context) {
	deployed, err := erc1820.ERC1820Deployed(ctx, client)
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	if deployed {
		log.Print("erc1820 contract already deployed")
		return
	}
	log.Print("Deploying erc1820 contract")
	err = erc1820.DeployERC1820Contract(ctx, client, key)
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
}

func deploy(ctx context.Context) {
	if !deployFlags.NoERC1820 {
		maybeDeployERC1820(ctx)
	}

	auth, err := makeAuth(ctx, client, key)
	failIfError(err)
	auth.Context = ctx
	var txs []*types.Transaction
	var tx *types.Transaction

	addTx := func() {
		failIfError(err)
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

	tokenAddress, tx, _, err := contract.DeployTestDepositTokenContract(auth, client)
	addTx()

	depositAddress, tx, _, err := contract.DeployDepositContract(auth, client, tokenAddress)
	addTx()

	targetAddress, tx, _, err := contract.DeployTestTargetContract(auth, client, executorAddress)
	addTx()

	// The keyper slasher requires the deposit contract to exist. Since at this point it doesn't,
	// gas estimation will fail, so we simply set it manually.
	auth.GasLimit = 2000000
	keyperSlasherAddress, tx, _, err := contract.DeployKeyperSlasher(
		auth,
		client,
		big.NewInt(defaultAppealBlocks),
		configAddress,
		executorAddress,
		depositAddress,
	)
	auth.GasLimit = 0
	addTx()

	receipts, err := waitForTransactions(ctx, client, txs)
	failIfError(err)

	totalGasUsed := uint64(0)
	for _, receipt := range receipts {
		totalGasUsed += receipt.GasUsed
	}
	fmt.Println("Gas used:", totalGasUsed)

	fmt.Println("ConfigContract address:", configAddress.Hex())
	fmt.Println("KeyBroadcastContract address:", broadcastAddress.Hex())
	fmt.Println("FeeBankContract address:", feeAddress.Hex())
	fmt.Println("BatcherContract address:", batcherAddress.Hex())
	fmt.Println("ExecutorContract address:", executorAddress.Hex())
	fmt.Println("TokenContract address:", tokenAddress.Hex())
	fmt.Println("DepositContract address:", depositAddress.Hex())
	fmt.Println("KeyperSlasher address:", keyperSlasherAddress.Hex())
	fmt.Println("TargetContract address:", targetAddress.Hex())

	if deployFlags.OutputFile != "" {
		j := sandbox.ContractsJSON{
			ConfigContract:        configAddress.Hex(),
			KeyBroadcastContract:  broadcastAddress.Hex(),
			FeeBankContract:       feeAddress.Hex(),
			BatcherContract:       batcherAddress.Hex(),
			ExecutorContract:      executorAddress.Hex(),
			TokenContract:         tokenAddress.Hex(),
			DepositContract:       depositAddress.Hex(),
			KeyperSlasherContract: keyperSlasherAddress.Hex(),
			TargetContract:        targetAddress.Hex(),
		}
		s, err := json.MarshalIndent(j, "", "    ")
		failIfError(err)
		err = ioutil.WriteFile(deployFlags.OutputFile, s, 0o644)
		failIfError(err)
		fmt.Println("addresses written to", deployFlags.OutputFile)
	}
}

func checkContractExists(ctx context.Context, configContractAddress common.Address) {
	code, err := client.CodeAt(ctx, configContractAddress, nil)
	failIfError(err)
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
	failIfError(err)
	auth.Context = ctx

	var txs []*types.Transaction
	var tx *types.Transaction

	addTx := func() {
		failIfError(err)
		txs = append(txs, tx)
		auth.Nonce.SetInt64(auth.Nonce.Int64() + 1)
	}

	checkContractExists(ctx, configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	failIfError(err)

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

	tx, err = cc.NextConfigSetTransactionSizeLimit(auth, transactionSizeLimit)
	addTx()

	tx, err = cc.NextConfigSetBatchSizeLimit(auth, batchSizeLimit)
	addTx()

	header, err := client.HeaderByNumber(ctx, nil)
	failIfError(err)
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
	failIfError(err)
	fmt.Printf("start block of config: %d\n", startBlockNumber)
}

func getconfig(ctx context.Context, configContractAddress common.Address, index int) {
	checkContractExists(ctx, configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	failIfError(err)
	c, err := cc.GetConfigByIndex(&bind.CallOpts{Context: ctx}, uint64(index))
	failIfError(err)

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
	failIfError(err)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	failIfError(err)

	amount := big.NewInt(fundAmount) // 1 eth
	gasLimit := uint64(baseGasLimit)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	failIfError(err)

	addTx := func() {
		failIfError(err)
		txs = append(txs, tx)
		nonce++
	}

	signer := types.NewEIP155Signer(chainID)

	for i := 0; i < sandbox.NumGanacheKeys()-1; i++ {
		receiver := crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey)
		var data []byte
		tx = types.NewTransaction(nonce, receiver, amount, gasLimit, gasPrice, data)

		tx, err = types.SignTx(tx, signer, key)
		failIfError(err)

		err = client.SendTransaction(ctx, tx)
		failIfError(err)
		addTx()
	}

	_, err = waitForTransactions(ctx, client, txs)
	failIfError(err)

	for i := 0; i < sandbox.NumGanacheKeys(); i++ {
		addr := crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey)
		balance, err := client.BalanceAt(context.Background(), addr, nil)
		failIfError(err)
		fmt.Printf("%s: %s\n", addr.Hex(), balance)
	}
}

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

func weiToGwei(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.GWei))
}

func gweiToWei(gwei *big.Int) *big.Int {
	return new(big.Int).Mul(gwei, big.NewInt(params.GWei))
}

func failIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
