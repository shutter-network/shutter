package prepare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/cmd/deploy"
	"github.com/shutter-network/shutter/shuttermint/contract"
	"github.com/shutter-network/shutter/shuttermint/keyper"
)

var contractsJSON deploy.Contracts

var configFlags struct {
	Dir            string
	NumKeypers     int
	EthereumURL    string
	ShuttermintURL string
	ContractsPath  string
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
		return nil
	},
}

func init() {
	PrepareCmd.AddCommand(configCmd)

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
		&configFlags.ShuttermintURL,
		"shuttermint-url",
		"s",
		"http://localhost:26657",
		"Shuttermint RPC URL",
	)
	configCmd.Flags().StringVarP(
		&configFlags.ContractsPath,
		"contracts",
		"c",
		"",
		"path to the contracts.json file",
	)

	configCmd.MarkFlagRequired("contracts")
}

func validateConfigFlags() error {
	if configFlags.NumKeypers <= 0 {
		return errors.Errorf("invalid flag --num-keypers: must be at least 1")
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
		config, err := newConfig()
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

func newConfig() (*keyper.Config, error) {
	config := keyper.Config{
		ShuttermintURL:              configFlags.ShuttermintURL,
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
		GasPriceMultiplier:          1.5,
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

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return errors.Wrap(err, "failed to create keyper config file")
	}
	if err = c.WriteTOML(file); err != nil {
		return errors.Wrap(err, "failed to write keyper config file")
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

func twothirds(numKeypers int) int {
	return (2*numKeypers + 2) / 3
}
