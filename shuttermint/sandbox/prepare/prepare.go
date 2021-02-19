package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/cmd"
	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

var configFlags struct {
	Dir                  string
	NumKeypers           int
	EthereumURL          string
	ShuttermintURLBase   string
	FirstShuttermintPort int
	ContractsPath        string
	Bin                  string
	FixedShuttermintPort bool
}

var scheduleFlags struct {
	Dir                    string
	OwnerKey               string
	TargetContractAddress  string
	TargetFunctionSelector string
	BatchSpan              int
	TransactionSizeLimit   int
	BatchSizeLimit         int
}

var fundFlags struct {
	Dir      string
	OwnerKey string
}

var contractsJSON sandbox.ContractsJSON

var rootCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare everything needed to test shutter.",
}

var configCmd = &cobra.Command{
	Use:   "configs",
	Short: "Generate the keyper config files",
	Run: func(cmd *cobra.Command, args []string) {
		if flag, err := validateConfigFlags(); err != nil {
			fmt.Printf("Invalid flag %s: %s\n", flag, err)
			os.Exit(1)
		}
		if err := loadContractsJSON(configFlags.ContractsPath); err != nil {
			fmt.Printf("Error loading contracts JSON file: %s\n", err)
			os.Exit(1)
		}
		if err := configs(); err != nil {
			fmt.Printf("Error creating configs: %s\n", err)
			os.Exit(1)
		}
		if err := createRunScript(); err != nil {
			fmt.Printf("Error creating run script: %s\n", err)
			os.Exit(1)
		}
	},
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a config in line with a set of earlier prepared keyper set",
	Run: func(cmd *cobra.Command, args []string) {
		if flag, err := validateScheduleFlags(); err != nil {
			fmt.Printf("Invalid flag %s: %s\n", flag, err)
			os.Exit(1)
		}
		if err := schedule(); err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	},
}

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Fund the accounts of an earlier prepared keyper set",
	Run: func(cmd *cobra.Command, args []string) {
		if flag, err := validateFundFlags(); err != nil {
			fmt.Printf("Invalid flag %s: %s\n", flag, err)
			os.Exit(1)
		}
		if err := fund(); err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(fundCmd)

	initConfigFlags()
	initScheduleFlags()
	initFundFlags()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfigFlags() {
	configCmd.Flags().StringVarP(
		&configFlags.Dir,
		"dir",
		"d",
		"testrun",
		"directory in which config files shall be stored",
	)
	configCmd.Flags().IntVarP(
		&configFlags.NumKeypers,
		"num-keypers",
		"n",
		3,
		"number of keypers",
	)
	configCmd.Flags().StringVarP(
		&configFlags.EthereumURL,
		"ethereum-url",
		"e",
		"",
		"Ethereum JSON RPC URL",
	)
	configCmd.Flags().StringVarP(
		&configFlags.ShuttermintURLBase,
		"shuttermint-url",
		"s",
		"http://localhost",
		"Shuttermint RPC URL (without port)",
	)
	configCmd.Flags().IntVarP(
		&configFlags.FirstShuttermintPort,
		"first-shuttermint-port",
		"p",
		26657,
		"port number of the shuttermint node for keyper 0",
	)
	configCmd.Flags().StringVarP(
		&configFlags.ContractsPath,
		"contracts",
		"c",
		"",
		"path to the contracts.json file",
	)
	configCmd.Flags().StringVar(
		&configFlags.Bin,
		"bin",
		"shuttermint",
		"(path to) shuttermint executable",
	)
	configCmd.Flags().BoolVar(
		&configFlags.FixedShuttermintPort,
		"fixed-shuttermint-port",
		false,
		"use a fixed shuttermint port",
	)

	err := configCmd.MarkFlagRequired("contracts")
	if err != nil {
		panic(err)
	}
}

func initScheduleFlags() {
	scheduleCmd.Flags().StringVarP(
		&scheduleFlags.Dir,
		"dir",
		"d",
		"testrun",
		"directory in which config files are stored",
	)
	scheduleCmd.Flags().StringVarP(
		&scheduleFlags.OwnerKey,
		"owner-key",
		"k",
		"b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773",
		"private key of the config contract owner",
	)
	scheduleCmd.Flags().StringVar(
		&scheduleFlags.TargetContractAddress,
		"target-contract",
		"0x25f96B23947F3e57b29d15760Fd8Af926694Fa81",
		"address of the target contract",
	)
	scheduleCmd.Flags().StringVar(
		&scheduleFlags.TargetFunctionSelector,
		"target-selector",
		"0x943d7209",
		"selector for the execute function in the target contract",
	)
	scheduleCmd.Flags().IntVarP(
		&scheduleFlags.BatchSpan,
		"batch-span",
		"b",
		5,
		"the batch span in blocks",
	)
	scheduleCmd.Flags().IntVar(
		&scheduleFlags.TransactionSizeLimit,
		"transaction-size-limit",
		100,
		"the size limit for transactions in bytes",
	)
	scheduleCmd.Flags().IntVar(
		&scheduleFlags.BatchSizeLimit,
		"batch-size-limit",
		100*100,
		"the size limit for batches in bytes",
	)
}

func initFundFlags() {
	fundCmd.Flags().StringVarP(
		&fundFlags.Dir,
		"dir",
		"d",
		"testrun",
		"directory in which config files are stored",
	)
	fundCmd.Flags().StringVarP(
		&fundFlags.OwnerKey,
		"owner-key",
		"k",
		"b0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773",
		"private key of the config contract owner",
	)
}

func validateConfigFlags() (string, error) {
	if configFlags.NumKeypers <= 0 {
		return "num-keypers", fmt.Errorf("must be at least 1")
	}
	if _, err := os.Stat(configFlags.Dir); !os.IsNotExist(err) {
		return "dir", fmt.Errorf("output directory %s already exists", configFlags.Dir)
	}
	var err error
	if configFlags.Bin, err = LookPath(configFlags.Bin); err != nil {
		return "", err
	}
	return "", nil
}

func validateScheduleFlags() (string, error) {
	stats, err := os.Stat(scheduleFlags.Dir)
	if os.IsNotExist(err) {
		return "dir", fmt.Errorf("directory %s does not exists", configFlags.Dir)
	}
	if !stats.IsDir() {
		return "dir", fmt.Errorf("%s is not a directory", configFlags.Dir)
	}

	if err := validatePrivateKey(scheduleFlags.OwnerKey); err != nil {
		return "owner-key", err
	}

	if err := validateAddress(scheduleFlags.TargetContractAddress); err != nil {
		return "target-contract", err
	}
	if err := validateFunctionSelector(scheduleFlags.TargetFunctionSelector); err != nil {
		return "target-selector", err
	}
	if scheduleFlags.BatchSpan < 0 {
		return "batch-span", fmt.Errorf("must not be negative")
	}
	if scheduleFlags.TransactionSizeLimit < 0 {
		return "transaction-size-limt", fmt.Errorf("must not be negative")
	}
	if scheduleFlags.BatchSizeLimit < 0 {
		return "batch-size-limit", fmt.Errorf("must not be negative")
	}

	return "", nil
}

func validateFundFlags() (string, error) {
	stats, err := os.Stat(fundFlags.Dir)
	if os.IsNotExist(err) {
		return "dir", fmt.Errorf("directory %s does not exists", fundFlags.Dir)
	}
	if !stats.IsDir() {
		return "dir", fmt.Errorf("%s is not a directory", fundFlags.Dir)
	}

	if err := validatePrivateKey(fundFlags.OwnerKey); err != nil {
		return "owner-key", err
	}

	return "", nil
}

func validateAddress(address string) error {
	addressParsed := common.HexToAddress(address)
	if addressParsed.Hex() != address {
		return fmt.Errorf("invalid address")
	}
	return nil
}

func validatePrivateKey(key string) error {
	if _, err := crypto.HexToECDSA(key); err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}
	return nil
}

func validateFunctionSelector(selector string) error {
	selectorBytes, err := hexutil.Decode(selector)
	if err != nil {
		return err
	}
	if len(selectorBytes) != 4 {
		return fmt.Errorf("function selector must be 4 bytes, got %d", len(selectorBytes))
	}
	return nil
}

func loadContractsJSON(path string) error {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(d, &contractsJSON)
	if err != nil {
		return err
	}

	err = contractsJSON.Validate()
	if err != nil {
		return err
	}

	return nil
}

func configs() error {
	configs := []*cmd.RawKeyperConfig{}
	for i := 0; i < configFlags.NumKeypers; i++ {
		config, err := rawConfig(i)
		if err != nil {
			return err
		}
		configs = append(configs, config)
	}
	return saveConfigs(configs)
}

func rawConfig(keyperIndex int) (*cmd.RawKeyperConfig, error) {
	signingKey, err := randomSigningKey()
	if err != nil {
		return nil, err
	}
	encryptionKey, err := randomEncryptionKey()
	if err != nil {
		return nil, err
	}
	validatorSeed, err := randomValidatorSeed()
	if err != nil {
		return nil, err
	}

	var shuttermintPort int
	if configFlags.FixedShuttermintPort {
		shuttermintPort = configFlags.FirstShuttermintPort
	} else {
		shuttermintPort = configFlags.FirstShuttermintPort + keyperIndex
	}

	shuttermintURL := configFlags.ShuttermintURLBase + ":" + strconv.Itoa(shuttermintPort)

	config := cmd.RawKeyperConfig{
		ShuttermintURL:       shuttermintURL,
		EthereumURL:          configFlags.EthereumURL,
		SigningKey:           hex.EncodeToString(crypto.FromECDSA(signingKey)),
		ValidatorSeed:        validatorSeed,
		EncryptionKey:        hex.EncodeToString(crypto.FromECDSA(encryptionKey.ExportECDSA())),
		ConfigContract:       contractsJSON.ConfigContract,
		BatcherContract:      contractsJSON.BatcherContract,
		KeyBroadcastContract: contractsJSON.KeyBroadcastContract,
		ExecutorContract:     contractsJSON.ExecutorContract,
		DepositContract:      contractsJSON.DepositContract,
		KeyperSlasher:        contractsJSON.KeyperSlasherContract,
		ExecutionStaggering:  "5",
	}
	return &config, nil
}

func randomSigningKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func randomEncryptionKey() (*ecies.PrivateKey, error) {
	encryptionKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	encryptionKey := ecies.ImportECDSA(encryptionKeyECDSA)
	return encryptionKey, nil
}

func randomValidatorSeed() (string, error) {
	seed := make([]byte, ed25519.SeedSize)
	if _, err := rand.Read(seed); err != nil {
		return "", err
	}
	return hex.EncodeToString(seed), nil
}

func saveConfigs(configs []*cmd.RawKeyperConfig) error {
	for i, c := range configs {
		configToml, err := toml.Marshal(*c)
		if err != nil {
			return err
		}

		dir := filepath.Join(configFlags.Dir, "keyper"+strconv.Itoa(i))
		if err = os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create keyper directory: %w", err)
		}
		path := filepath.Join(dir, "config.toml")

		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create keyper config file: %w", err)
		}
		if _, err = file.Write(configToml); err != nil {
			return fmt.Errorf("failed to write keyper config file: %w", err)
		}
	}
	return nil
}

func schedule() error {
	configPaths, err := findConfigFiles(scheduleFlags.Dir)
	if err != nil {
		return err
	}
	configs, err := loadConfigs(configPaths)
	if err != nil {
		return err
	}

	if len(configs) < 3 {
		return fmt.Errorf("3 keypers required, but only %d config files found in %s", len(configs), scheduleFlags.Dir)
	}

	ownerKey, err := crypto.HexToECDSA(scheduleFlags.OwnerKey)
	if err != nil {
		return fmt.Errorf("invalid owner key")
	}

	ethereumURL := configs[0].EthereumURL
	client, err := ethclient.DialContext(context.Background(), ethereumURL)
	if err != nil {
		return fmt.Errorf("faild to connect to Ethereum node at %s: %w", ethereumURL, err)
	}

	if err := scheduleForKeyperConfigs(context.Background(), client, ownerKey, configs); err != nil {
		return err
	}

	return nil
}

func findConfigFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, fmt.Errorf("failed to read directory %s", dir)
	}

	configPaths := []string{}
	for _, file := range files {
		if file.IsDir() {
			configPath := filepath.Join(scheduleFlags.Dir, file.Name(), "config.toml")
			configPaths = append(configPaths, configPath)
		}
	}
	return configPaths, nil
}

func loadConfigs(paths []string) ([]keyper.KeyperConfig, error) {
	configs := []keyper.KeyperConfig{}
	for _, path := range paths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return []keyper.KeyperConfig{}, fmt.Errorf("failed to read config file at %s: %w", path, err)
		}

		var r cmd.RawKeyperConfig
		if err := toml.Unmarshal(data, &r); err != nil {
			return []keyper.KeyperConfig{}, fmt.Errorf("failed to read config file %s: %w", path, err)
		}

		c, err := cmd.ValidateKeyperConfig(r)
		if err != nil {
			return []keyper.KeyperConfig{}, fmt.Errorf("failed to parse config file %s: %w", path, err)
		}

		configs = append(configs, c)
	}
	return configs, nil
}

func scheduleForKeyperConfigs(ctx context.Context, client *ethclient.Client, ownerKey *ecdsa.PrivateKey, configs []keyper.KeyperConfig) error {
	numKeypers := len(configs)
	threshold := uint64((numKeypers*2 + 2) / 3) // 2/3 rounded up
	keypers := []common.Address{}
	for _, c := range configs {
		keypers = append(keypers, c.Address())
	}

	if err := checkContractExists(ctx, client, configs[0].ConfigContractAddress); err != nil {
		return err
	}

	auth, err := makeAuth(ctx, client, ownerKey)
	if err != nil {
		return err
	}
	auth.Context = ctx
	var txs []*types.Transaction
	var tx *types.Transaction
	addTx := func() {
		txs = append(txs, tx)
		auth.Nonce.SetInt64(auth.Nonce.Int64() + 1)
	}

	cc, err := contract.NewConfigContract(configs[0].ConfigContractAddress, client)
	if err != nil {
		return err
	}

	tx, err = cc.NextConfigSetBatchSpan(auth, uint64(scheduleFlags.BatchSpan))
	if err != nil {
		return err
	}
	addTx()

	tx, err = cc.NextConfigAddKeypers(auth, keypers)
	if err != nil {
		return err
	}
	addTx()

	tx, err = cc.NextConfigSetThreshold(auth, threshold)
	if err != nil {
		return err
	}
	addTx()

	tx, err = cc.NextConfigSetTargetAddress(auth, common.HexToAddress(scheduleFlags.TargetContractAddress))
	if err != nil {
		return err
	}
	addTx()

	selectorBytesSlice, err := hexutil.Decode(scheduleFlags.TargetFunctionSelector)
	if err != nil {
		return err // this should already be catched during flag validation
	}
	var selectorBytes [4]byte
	copy(selectorBytes[:], selectorBytesSlice)
	tx, err = cc.NextConfigSetTargetFunctionSelector(auth, selectorBytes)
	if err != nil {
		return err
	}
	addTx()

	tx, err = cc.NextConfigSetExecutionTimeout(auth, uint64(scheduleFlags.BatchSpan))
	if err != nil {
		return err
	}
	addTx()

	tx, err = cc.NextConfigSetTransactionSizeLimit(auth, uint64(scheduleFlags.TransactionSizeLimit))
	if err != nil {
		return err
	}
	addTx()

	tx, err = cc.NextConfigSetBatchSizeLimit(auth, uint64(scheduleFlags.BatchSizeLimit))
	if err != nil {
		return err
	}
	addTx()

	startBlockNumber, startBatchIndex, err := chooseStartBlockAndBatch(ctx, client, cc)
	if err != nil {
		return err
	}
	tx, err = cc.NextConfigSetStartBlockNumber(auth, startBlockNumber)
	if err != nil {
		return err
	}
	addTx()
	tx, err = cc.NextConfigSetStartBatchIndex(auth, startBatchIndex)
	if err != nil {
		return err
	}
	addTx()

	tx, err = cc.ScheduleNextConfig(auth)
	if err != nil {
		return err
	}
	addTx()

	_, err = waitForTransactions(ctx, client, txs)
	if err != nil {
		return err
	}

	return nil
}

func checkContractExists(ctx context.Context, client *ethclient.Client, address common.Address) error {
	code, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		return fmt.Errorf("failed to check contract code: %w", err)
	}
	if len(code) == 0 {
		return fmt.Errorf("no contract exists at address %s", address.Hex())
	}
	return nil
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
	auth.GasLimit = 1000000
	return auth, nil
}

func chooseStartBlockAndBatch(ctx context.Context, client *ethclient.Client, cc *contract.ConfigContract) (uint64, uint64, error) {
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	numConfigs, err := cc.NumConfigs(callOpts)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to query number of configs")
	}

	var batchSpan uint64
	var startBlockNumber uint64
	var startBatchIndex uint64
	if numConfigs != 0 {
		config, err := cc.GetConfigByIndex(callOpts, numConfigs-1)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to query config %d: %w", numConfigs-1, err)
		}
		batchSpan = config.BatchSpan
		startBlockNumber = config.StartBlockNumber
		startBatchIndex = config.StartBatchIndex
	} else {
		batchSpan = 0
		startBlockNumber = 0
		startBatchIndex = 0
	}

	headsUp, err := cc.ConfigChangeHeadsUpBlocks(callOpts)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to query config change heads up blocks: %w", err)
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get header: %w", err)
	}
	minStartBlock := header.Number.Uint64() + headsUp + 10

	if batchSpan == 0 {
		return minStartBlock, startBatchIndex, nil
	}
	delta := minStartBlock - startBlockNumber
	numStartedBatches := (delta + batchSpan + 1) / batchSpan
	newStartBlockNumber := startBlockNumber + numStartedBatches*batchSpan
	newStartBatchIndex := startBatchIndex + numStartedBatches
	return newStartBlockNumber, newStartBatchIndex, nil
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
		return res, errors.New("some txs have failed")
	}

	return res, nil
}

func fund() error {
	ctx := context.Background()

	configPaths, err := findConfigFiles(fundFlags.Dir)
	if err != nil {
		return err
	}
	configs, err := loadConfigs(configPaths)
	if err != nil {
		return err
	}

	ownerKey, err := crypto.HexToECDSA(fundFlags.OwnerKey)
	if err != nil {
		return fmt.Errorf("invalid owner key")
	}
	ownerAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)

	ethereumURL := configs[0].EthereumURL
	client, err := ethclient.DialContext(context.Background(), ethereumURL)
	if err != nil {
		return fmt.Errorf("faild to connect to Ethereum node at %s: %w", ethereumURL, err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to query chain id: %w", err)
	}
	signer := types.NewEIP155Signer(chainID)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to query gas price: %w", err)
	}

	amount, ok := new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		panic("unexpected error")
	}

	nonce, err := client.PendingNonceAt(ctx, ownerAddress)
	if err != nil {
		return fmt.Errorf("failed to query nonce: %w", err)
	}

	txs := []*types.Transaction{}
	for _, config := range configs {
		unsignedTx := types.NewTransaction(nonce, config.Address(), amount, 21000, gasPrice, []byte{})
		tx, err := types.SignTx(unsignedTx, signer, ownerKey)
		if err != nil {
			return fmt.Errorf("failed to sign transaction: %w", err)
		}

		err = client.SendTransaction(ctx, tx)
		if err != nil {
			return fmt.Errorf("failed to send transaction: %w", err)
		}
		nonce++

		txs = append(txs, tx)
	}

	_, err = waitForTransactions(ctx, client, txs)
	return err
}

// LookPath searches for an executable in $PATH. The difference to os.LookPath is, that this
// function makes sure to return an absolute path.
func LookPath(file string) (string, error) {
	p, err := exec.LookPath(file)
	if err != nil {
		return "", fmt.Errorf("cannot find executable %s: %w", file, err)
	}
	return filepath.Abs(p)
}

func createRunScript() error {
	path := filepath.Join(configFlags.Dir, "run.sh")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create run script file: %w", err)
	}
	err = os.Chmod(path, 0o755)
	if err != nil {
		return fmt.Errorf("failed to make run script executable: %w", err)
	}

	shuttermintCmd, err := LookPath(configFlags.Bin)
	if err != nil {
		return err
	}
	configPaths, err := findConfigFiles(configFlags.Dir)
	if err != nil {
		return err
	}
	keyperDirs := []string{}
	for _, path := range configPaths {
		dir := filepath.Join("..", filepath.Dir(path))
		keyperDirs = append(keyperDirs, dir)
	}

	t, err := template.New("script").Parse(runScriptTemplate)
	if err != nil {
		return fmt.Errorf("failed to create template: %w", err)
	}

	d := runScriptTemplateData{
		ShuttermintCmd: shuttermintCmd,
		KeyperDirs:     keyperDirs,
	}
	err = t.Execute(file, d)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// err := ioutil.WriteFile(path, []byte(runScriptTemplate), 0755)
	// if err != nil {
	// 	return fmt.Errorf("failed to write run script: %w", err)
	// }
	return nil
}

type runScriptTemplateData struct {
	ShuttermintCmd string
	KeyperDirs     []string
}

//nolint:lll
const runScriptTemplate = `#! /usr/bin/env bash
set -euxo pipefail

SHUTTERMINT={{$.ShuttermintCmd}}

PARENT_PATH=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd ${PARENT_PATH}

function getRPCURL() {
	python -c "import configparser; f = '[_main]\n' + open('$PARENT_PATH/$1/config/config.toml').read(); c = configparser.ConfigParser(); c.read_string(f); print(c['rpc']['laddr'][7:-1])"
}

SESSION=shutter
SHUTTERMINT_START_TIME=5

# create shuttermint configs
{{range $i, $d := $.KeyperDirs}}
${SHUTTERMINT} init --dev --root ${PARENT_PATH}/{{$d}} --index {{$i}}{{end}}

# make all keypers use same genesis file
{{range $i, $d := $.KeyperDirs}}{{if $i}}
cp {{index $.KeyperDirs 0}}/config/genesis.json ${PARENT_PATH}/{{$d}}/config/genesis.json{{end}}{{end}}

# start first keyper in tmux session
tmux new -s ${SESSION} -d
tmux send-keys "${SHUTTERMINT} run --config ${PARENT_PATH}/{{index $.KeyperDirs 0}}/config/config.toml" C-m
sleep ${SHUTTERMINT_START_TIME}

# query p2p address of keyper0 via RPC
RPC_ADDR_0=$(getRPCURL {{index $.KeyperDirs 0}})
P2P_ADDR_0=$(curl "${RPC_ADDR_0}"/status | python -c "import json, sys; i = json.load(sys.stdin)['result']['node_info']; print(i['id'] + '@127.0.0.1:' + i['listen_addr'].split(':')[-1])")

# add keyper 0 as persistent peer for the other keypers
{{range $i, $d := $.KeyperDirs}}{{if $i}}
sed -i "s/persistent_peers = \"\"/persistent_peers = \"${P2P_ADDR_0}\"/" {{$d}}/config/config.toml{{end}}{{end}}

{{range $i, $d := $.KeyperDirs}}{{if $i}}
tmux split-window -h
tmux send-keys "${SHUTTERMINT} run --config ${PARENT_PATH}/{{$d}}/config/config.toml" C-m{{end}}{{end}}
sleep ${SHUTTERMINT_START_TIME}

# start keypers and make them connect to their assigned shuttermint node
{{range $i, $d := $.KeyperDirs}}
RPC_URL=$(getRPCURL {{$d}})
sed -i "s/ShuttermintURL = \".*/ShuttermintURL = \"http:\/\/${RPC_URL}\"/" {{$d}}/config.toml
tmux split-window -h
tmux send-keys "${SHUTTERMINT} keyper --config {{$d}}/config.toml" C-m{{end}}

tmux select-layout even-horizontal
exec tmux attach -t ${SESSION}
`
