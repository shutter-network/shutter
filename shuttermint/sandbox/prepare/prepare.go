package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/brainbot-com/shutter/shuttermint/cmd/config"
	"github.com/brainbot-com/shutter/shuttermint/cmd/deploy"
	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/brainbot-com/shutter/shuttermint/medley"
	"github.com/brainbot-com/shutter/shuttermint/medley/txbatch"
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

var contractsJSON deploy.Contracts

var rootCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare everything needed to test shutter.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return medley.BindFlags(cmd)
	},
	SilenceUsage: true,
	Run:          medley.ShowHelpAndExit,
}

var configCmd = &cobra.Command{
	Use:   "configs",
	Short: "Generate the keyper config files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateConfigFlags(); err != nil {
			return err
		}
		if err := loadContractsJSON(configFlags.ContractsPath); err != nil {
			return errors.WithMessage(err, "load contracts JSON file")
		}
		if err := configs(); err != nil {
			return errors.WithMessage(err, "create config files")
		}
		if err := createRunScript(); err != nil {
			return errors.WithMessage(err, "create run script")
		}
		return nil
	},
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a config in line with a set of earlier prepared keyper set",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateScheduleFlags(); err != nil {
			return err
		}
		return schedule()
	},
}

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Fund the accounts of an earlier prepared keyper set",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateFundFlags(); err != nil {
			return err
		}
		return fund()
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
	status := 0

	if err := rootCmd.Execute(); err != nil {
		status = 1
	}

	os.Exit(status)
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
		1000,
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

func validateConfigFlags() error {
	if configFlags.NumKeypers <= 0 {
		return errors.Errorf("invalid flag --num-keypers: must be at least 1")
	}
	if _, err := os.Stat(configFlags.Dir); !os.IsNotExist(err) {
		return errors.Errorf("invalid flag --dir: output directory %s already exists", configFlags.Dir)
	}
	var err error
	if configFlags.Bin, err = LookPath(configFlags.Bin); err != nil {
		return err
	}
	return nil
}

func validateScheduleFlags() error {
	stats, err := os.Stat(scheduleFlags.Dir)
	if os.IsNotExist(err) {
		return errors.Errorf("invalid flag --dir: directory %s does not exists", configFlags.Dir)
	}
	if !stats.IsDir() {
		return errors.Errorf("invalid flag --dir: %s is not a directory", configFlags.Dir)
	}

	if err := validatePrivateKey(scheduleFlags.OwnerKey); err != nil {
		return errors.WithMessage(err, "invalid flag --owner-key")
	}

	if err := validateAddress(scheduleFlags.TargetContractAddress); err != nil {
		return errors.WithMessage(err, "invalid flag --target-contract")
	}
	if err := validateFunctionSelector(scheduleFlags.TargetFunctionSelector); err != nil {
		return errors.WithMessage(err, "invalid flag --target-selector")
	}
	if scheduleFlags.BatchSpan < 0 {
		return errors.Errorf("invalid flag --batch-span: must not be negative")
	}
	if scheduleFlags.TransactionSizeLimit < 0 {
		return errors.Errorf("invalid flag --transaction-size-limit: must not be negative")
	}
	if scheduleFlags.BatchSizeLimit < 0 {
		return errors.Errorf("invalid flag --batch-size-limit: must not be negative")
	}

	return nil
}

func validateFundFlags() error {
	stats, err := os.Stat(fundFlags.Dir)
	if os.IsNotExist(err) {
		return errors.Errorf("invalid flag --dir: directory %s does not exists", fundFlags.Dir)
	}
	if !stats.IsDir() {
		return errors.Errorf("invalid flag --dir: %s is not a directory", fundFlags.Dir)
	}

	if err := validatePrivateKey(fundFlags.OwnerKey); err != nil {
		return errors.WithMessage(err, "invalid flag --owner-key")
	}

	return nil
}

func validateAddress(address string) error {
	addressParsed := common.HexToAddress(address)
	if addressParsed.Hex() != address {
		return errors.Errorf("invalid address")
	}
	return nil
}

func validatePrivateKey(key string) error {
	if _, err := crypto.HexToECDSA(key); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func validateFunctionSelector(selector string) error {
	selectorBytes, err := hexutil.Decode(selector)
	if err != nil {
		return err
	}
	if len(selectorBytes) != 4 {
		return errors.Errorf("function selector must be 4 bytes, got %d", len(selectorBytes))
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

	return nil
}

func generateConfigJSON(keypers []common.Address) error {
	cfg := contract.BatchConfig{
		StartBatchIndex:        0,
		StartBlockNumber:       0,
		Keypers:                keypers,
		Threshold:              uint64(twothirds(len(keypers))),
		BatchSpan:              10,
		BatchSizeLimit:         100000,
		TransactionSizeLimit:   1000,
		TransactionGasLimit:    10000,
		FeeReceiver:            common.HexToAddress("0x1111111111111111111111111111111111111111"),
		TargetAddress:          contractsJSON.TargetContract,
		TargetFunctionSelector: [4]byte{0x94, 0x3d, 0x72, 0x09},
		ExecutionTimeout:       15,
	}

	j, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	path := filepath.Join(configFlags.Dir, "config.json")
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "create %s", path)
	}
	defer file.Close()
	_, err = file.Write(j)
	if err != nil {
		return errors.Wrapf(err, "write to %s", path)
	}

	fmt.Printf(
		"Please adapt the following and use it as config.json. A copy has also been saved to %s.\n\n",
		path,
	)
	fmt.Println(string(j))
	return nil
}

func configs() error {
	keypers := []common.Address{}
	for i := 0; i < configFlags.NumKeypers; i++ {
		config, err := newConfig(i)
		if err != nil {
			return err
		}
		dir := filepath.Join(configFlags.Dir, "keyper"+strconv.Itoa(i))
		err = saveConfig(config, dir)
		if err != nil {
			return err
		}
		keypers = append(keypers, config.Address())
	}
	return generateConfigJSON(keypers)
}

func newConfig(keyperIndex int) (*keyper.Config, error) {
	var shuttermintPort int
	if configFlags.FixedShuttermintPort {
		shuttermintPort = configFlags.FirstShuttermintPort
	} else {
		shuttermintPort = configFlags.FirstShuttermintPort + keyperIndex
	}

	shuttermintURL := configFlags.ShuttermintURLBase + ":" + strconv.Itoa(shuttermintPort)

	config := keyper.Config{
		ShuttermintURL:              shuttermintURL,
		EthereumURL:                 configFlags.EthereumURL,
		DBDir:                       "",
		ConfigContractAddress:       contractsJSON.ConfigContract,
		BatcherContractAddress:      contractsJSON.BatcherContract,
		KeyBroadcastContractAddress: contractsJSON.KeyBroadcastContract,
		ExecutorContractAddress:     contractsJSON.ExecutorContract,
		DepositContractAddress:      contractsJSON.DepositContract,
		KeyperSlasherAddress:        contractsJSON.KeyperSlasherContract,
		MainChainFollowDistance:     0,
		ExecutionStaggering:         5,
		DKGPhaseLength:              30,
	}
	err := config.GenerateNewKeys()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func saveConfig(c *keyper.Config, dir string) error {
	var err error
	if err = os.MkdirAll(dir, 0o755); err != nil {
		return errors.Wrap(err, "failed to create keyper directory")
	}
	path := filepath.Join(dir, "config.toml")

	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "failed to create keyper config file")
	}
	if err = c.WriteTOML(file); err != nil {
		return errors.Wrap(err, "failed to write keyper config file")
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
		return errors.Errorf("3 keypers required, but only %d config files found in %s", len(configs), scheduleFlags.Dir)
	}

	ownerKey, err := crypto.HexToECDSA(scheduleFlags.OwnerKey)
	if err != nil {
		return errors.Errorf("invalid owner key")
	}

	ethereumURL := configs[0].EthereumURL
	client, err := ethclient.DialContext(context.Background(), ethereumURL)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Ethereum node at %s", ethereumURL)
	}

	if err := scheduleForKeyperConfigs(context.Background(), client, ownerKey, configs); err != nil {
		return err
	}

	return nil
}

func findConfigFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, errors.Errorf("failed to read directory %s", dir)
	}

	configPaths := []string{}
	for _, file := range files {
		if file.IsDir() {
			configPath := filepath.Join(dir, file.Name(), "config.toml")
			configPaths = append(configPaths, configPath)
		}
	}
	return configPaths, nil
}

func loadConfig(path string) (keyper.Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		return keyper.Config{}, errors.Wrapf(err, "failed to read config file %s", path)
	}

	config := keyper.Config{}
	err = config.Unmarshal(v)
	if err != nil {
		return keyper.Config{}, errors.Wrapf(err, "failed to unmarshal config file %s", path)
	}
	return config, nil
}

func loadConfigs(paths []string) ([]keyper.Config, error) {
	configs := []keyper.Config{}
	for _, path := range paths {
		config, err := loadConfig(path)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func twothirds(numKeypers int) int {
	return (2*numKeypers + 2) / 3
}

func scheduleForKeyperConfigs(ctx context.Context, client *ethclient.Client, ownerKey *ecdsa.PrivateKey, configs []keyper.Config) error {
	numKeypers := len(configs)
	threshold := uint64(twothirds(numKeypers))
	keypers := []common.Address{}
	for i := range configs {
		keypers = append(keypers, configs[i].Address())
	}

	if err := checkContractExists(ctx, client, configs[0].ConfigContractAddress); err != nil {
		return err
	}

	cc, err := contract.NewConfigContract(configs[0].ConfigContractAddress, client)
	if err != nil {
		return err
	}

	batch, err := txbatch.New(ctx, client, ownerKey)
	if err != nil {
		return err
	}

	batch.TransactOpts.GasLimit = 1000000
	batch.TransactOpts.Context = ctx

	auth := batch.TransactOpts
	var tx *types.Transaction

	selectorBytesSlice, err := hexutil.Decode(scheduleFlags.TargetFunctionSelector)
	if err != nil {
		return err // this should already be catched during flag validation
	}
	var selectorBytes [4]byte
	copy(selectorBytes[:], selectorBytesSlice)

	startBlockNumber, startBatchIndex, err := config.ChooseStartBlockAndBatch(ctx, client, cc)
	if err != nil {
		return err
	}

	bc := contract.BatchConfig{
		BatchSpan:              uint64(scheduleFlags.BatchSpan),
		Keypers:                keypers,
		Threshold:              threshold,
		TargetAddress:          common.HexToAddress(scheduleFlags.TargetContractAddress),
		TargetFunctionSelector: selectorBytes,
		ExecutionTimeout:       uint64(scheduleFlags.BatchSpan),
		TransactionSizeLimit:   uint64(scheduleFlags.TransactionSizeLimit),
		BatchSizeLimit:         uint64(scheduleFlags.BatchSizeLimit),
		StartBlockNumber:       startBlockNumber,
		StartBatchIndex:        startBatchIndex,
	}
	err = cc.SetNextBatchConfig(ctx, batch, bc)
	if err != nil {
		return err
	}

	tx, err = cc.ScheduleNextConfig(auth)
	if err != nil {
		return err
	}
	batch.Add(tx)
	_, err = batch.WaitMined(ctx)
	return err
}

func checkContractExists(ctx context.Context, client *ethclient.Client, address common.Address) error {
	code, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		return errors.Wrap(err, "failed to check contract code")
	}
	if len(code) == 0 {
		return errors.Errorf("no contract exists at address %s", address.Hex())
	}
	return nil
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
		return errors.Errorf("invalid owner key")
	}

	ethereumURL := configs[0].EthereumURL
	client, err := ethclient.DialContext(context.Background(), ethereumURL)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Ethereum node at %s", ethereumURL)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query chain id")
	}
	signer := types.NewEIP155Signer(chainID)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to query gas price")
	}

	amount, ok := new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		panic("unexpected error")
	}

	keypers := []common.Address{}
	for i := range configs {
		keypers = append(keypers, configs[i].Address())
	}

	batch, err := txbatch.New(ctx, client, ownerKey)
	if err != nil {
		return err
	}

	for _, addr := range keypers {
		unsignedTx := types.NewTransaction(batch.TransactOpts.Nonce.Uint64(), addr, amount, 21000, gasPrice, []byte{})
		tx, err := types.SignTx(unsignedTx, signer, ownerKey)
		if err != nil {
			return errors.Wrap(err, "failed to sign transaction")
		}

		err = client.SendTransaction(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "failed to send transaction")
		}
		batch.Add(tx)
	}
	_, err = batch.WaitMined(ctx)
	return err
}

// LookPath searches for an executable in $PATH. The difference to os.LookPath is, that this
// function makes sure to return an absolute path.
func LookPath(file string) (string, error) {
	p, err := exec.LookPath(file)
	if err != nil {
		return "", errors.Wrapf(err, "cannot find executable %s", file)
	}
	return filepath.Abs(p)
}

func createRunScript() error {
	path := filepath.Join(configFlags.Dir, "run.sh")
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "failed to create run script file")
	}
	err = os.Chmod(path, 0o755)
	if err != nil {
		return errors.Wrap(err, "failed to make run script executable")
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
		return errors.Wrap(err, "failed to create template")
	}

	d := runScriptTemplateData{
		ShuttermintCmd: shuttermintCmd,
		KeyperDirs:     keyperDirs,
	}
	err = t.Execute(file, d)
	if err != nil {
		return errors.Wrap(err, "failed to execute template")
	}

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
