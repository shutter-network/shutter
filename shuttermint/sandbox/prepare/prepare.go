package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/cmd"
	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper"
)

var configFlags struct {
	Dir                         string
	NumKeypers                  int
	EthereumURL                 string
	ShuttermintURLBase          string
	FirstShuttermintPort        int
	ConfigContractAddress       string
	BatcherContractAddress      string
	KeyBroadcastContractAddress string
	ExecutorContractAddress     string
	Bin                         string
}

var scheduleFlags struct {
	Dir                  string
	OwnerKey             string
	BatchSpan            int
	TransactionSizeLimit int
	BatchSizeLimit       int
}

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

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(scheduleCmd)

	initConfigFlags()
	initScheduleFlags()
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
		"ws://localhost:8545/websocket",
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
	configCmd.Flags().StringVar(
		&configFlags.ConfigContractAddress,
		"config-contract",
		"0x07a457d878BF363E0Bb5aa0B096092f941e19962",
		"address of the config contract",
	)
	configCmd.Flags().StringVar(
		&configFlags.BatcherContractAddress,
		"batcher-contract",
		"0x27D44c7337ce4D67b7cd573e9c36bDEED2b2162a",
		"address of the batcher contract",
	)
	configCmd.Flags().StringVar(
		&configFlags.KeyBroadcastContractAddress,
		"key-broadcast-contract",
		"0xFA33c8EF8b5c4f3003361c876a298D1DB61ccA4e",
		"address of the key broadcast contract",
	)
	configCmd.Flags().StringVar(
		&configFlags.ExecutorContractAddress,
		"executor-contract",
		"0x5d18dED3c0A476fCbc9E67Fc1C613cfc5DD0d34B",
		"address of the executor contract",
	)
	configCmd.Flags().StringVar(
		&configFlags.Bin,
		"bin",
		"../../../bin/shuttermint",
		"path to the shuttermint executable",
	)
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

func validateConfigFlags() (string, error) {
	if configFlags.NumKeypers <= 0 {
		return "num-keypers", fmt.Errorf("must be at least 1")
	}
	if err := validateAddress(configFlags.ConfigContractAddress); err != nil {
		return "config-contract", err
	}
	if err := validateAddress(configFlags.BatcherContractAddress); err != nil {
		return "batcher-contract", err
	}
	if err := validateAddress(configFlags.KeyBroadcastContractAddress); err != nil {
		return "key-broadcast-contract", err
	}
	if err := validateAddress(configFlags.ExecutorContractAddress); err != nil {
		return "executor-contract", err
	}
	if _, err := os.Stat(configFlags.Dir); !os.IsNotExist(err) {
		return "dir", fmt.Errorf("output directory %s already exists", configFlags.Dir)
	}
	if _, err := os.Stat(configFlags.Bin); err != nil {
		return "bin-dir", fmt.Errorf("shuttermint executable not found: %w", err)
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

	shuttermintPort := configFlags.FirstShuttermintPort + keyperIndex
	shuttermintURL := configFlags.ShuttermintURLBase + ":" + strconv.Itoa(shuttermintPort)

	config := cmd.RawKeyperConfig{
		ShuttermintURL:       shuttermintURL,
		EthereumURL:          configFlags.EthereumURL,
		SigningKey:           hex.EncodeToString(crypto.FromECDSA(signingKey)),
		ValidatorSeed:        validatorSeed,
		EncryptionKey:        hex.EncodeToString(crypto.FromECDSA(encryptionKey.ExportECDSA())),
		ConfigContract:       configFlags.ConfigContractAddress,
		BatcherContract:      configFlags.BatcherContractAddress,
		KeyBroadcastContract: configFlags.KeyBroadcastContractAddress,
		ExecutorContract:     configFlags.ExecutorContractAddress,
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
		toml, err := toml.Marshal(*c)
		if err != nil {
			return err
		}

		dir := filepath.Join(configFlags.Dir, "keyper"+strconv.Itoa(i))
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create keyper directory: %w", err)
		}
		path := filepath.Join(dir, "config.toml")

		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create keyper config file: %w", err)
		}
		if _, err = file.Write(toml); err != nil {
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
		// TODO: check if two (or even three) keypers are required
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

	auth := bind.NewKeyedTransactor(privateKey)
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

func createRunScript() error {
	path := filepath.Join(configFlags.Dir, "run.sh")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create run script file: %w", err)
	}
	err = os.Chmod(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to make run script executable: %w", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
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
		ShuttermintCmd: filepath.Join(cwd, configFlags.Bin),
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
