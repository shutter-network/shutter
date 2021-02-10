package sandbox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

const selectorLength = 4

// ContractJSON stores the hex encoded addresses of all contracts.
type ContractsJSON struct {
	ConfigContract        string
	KeyBroadcastContract  string
	FeeBankContract       string
	BatcherContract       string
	ExecutorContract      string
	TokenContract         string
	DepositContract       string
	KeyperSlasherContract string
	TargetContract        string
}

// ConfigJSON represents a batch config in JSON format
type ConfigJSON struct {
	StartBatchIndex        uint64
	StartBlockNumber       uint64
	Keypers                []string
	Threshold              uint64
	BatchSpan              uint64
	BatchSizeLimit         uint64
	TransactionSizeLimit   uint64
	TransactionGasLimit    uint64
	FeeReceiver            string
	TargetAddress          string
	TargetFunctionSelector string
	ExecutionTimeout       uint64
}

func (c ContractsJSON) Validate() error {
	addresses := []string{
		c.ConfigContract,
		c.KeyBroadcastContract,
		c.FeeBankContract,
		c.BatcherContract,
		c.ExecutorContract,
		c.TokenContract,
		c.DepositContract,
		c.KeyperSlasherContract,
		c.TargetContract,
	}
	for _, address := range addresses {
		if err := validateAddress(address); err != nil {
			return err
		}
	}
	return nil
}

func (c ConfigJSON) Validate() error {
	addresses := []string{
		c.FeeReceiver,
		c.TargetAddress,
	}
	addresses = append(addresses, c.Keypers...)
	for _, address := range addresses {
		if err := validateAddress(address); err != nil {
			return err
		}
	}

	if err := validateTargetFunctionSelector(c.TargetFunctionSelector); err != nil {
		return err
	}

	if err := validateThreshold(c.Threshold, len(c.Keypers)); err != nil {
		return err
	}

	return nil
}

func validateAddress(address string) error {
	addressParsed := common.HexToAddress(address)
	if addressParsed.Hex() != address {
		return fmt.Errorf("invalid address")
	}
	return nil
}

func validateTargetFunctionSelector(selector string) error {
	decoded, err := hexutil.Decode(selector)
	if err != nil {
		return err
	}
	if len(decoded) != selectorLength {
		return fmt.Errorf("target function selector must be 4 bytes long")
	}
	return nil
}

func validateThreshold(threshold uint64, numKeypers int) error {
	if numKeypers <= 0 {
		return fmt.Errorf("there must be at least one keyper")
	}
	if threshold > uint64(numKeypers) {
		return fmt.Errorf("threshold must not be greater than number of keypers")
	}
	if threshold == 0 {
		return fmt.Errorf("threshold must not be zero")
	}
	return nil
}

// LoadContractsJSON loads and validates a contracts json file.
func LoadContractsJSON(path string) (*ContractsJSON, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &ContractsJSON{}
	err = json.Unmarshal(d, c)
	if err != nil {
		return nil, err
	}

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func ConfigToJSON(config *contract.BatchConfig) *ConfigJSON {
	keypers := []string{}
	for _, keyper := range config.Keypers {
		keypers = append(keypers, keyper.Hex())
	}

	return &ConfigJSON{
		StartBatchIndex:        config.StartBatchIndex,
		StartBlockNumber:       config.StartBlockNumber,
		Keypers:                keypers,
		Threshold:              config.Threshold,
		BatchSpan:              config.BatchSpan,
		BatchSizeLimit:         config.BatchSizeLimit,
		TransactionSizeLimit:   config.TransactionSizeLimit,
		TransactionGasLimit:    config.TransactionGasLimit,
		FeeReceiver:            config.FeeReceiver.Hex(),
		TargetAddress:          config.TargetAddress.Hex(),
		TargetFunctionSelector: hexutil.Encode(config.TargetFunctionSelector[:]),
		ExecutionTimeout:       config.ExecutionTimeout,
	}
}

func LoadConfigJSON(path string) (*ConfigJSON, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &ConfigJSON{}
	err = json.Unmarshal(d, c)
	if err != nil {
		return nil, err
	}

	err = c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// ToBatchConfig converts the ConfigJSON to a contract.BatchConfig. It assumes `Validate` has been
// called prior.
func (c ConfigJSON) ToBatchConfig() *contract.BatchConfig {
	keypers := []common.Address{}
	for _, keyper := range c.Keypers {
		keypers = append(keypers, common.HexToAddress(keyper))
	}

	selectorSlice := hexutil.MustDecode(c.TargetFunctionSelector)
	var selector [4]byte
	copy(selector[:], selectorSlice)

	return &contract.BatchConfig{
		StartBatchIndex:        c.StartBatchIndex,
		StartBlockNumber:       c.StartBlockNumber,
		Keypers:                keypers,
		Threshold:              c.Threshold,
		BatchSpan:              c.BatchSpan,
		BatchSizeLimit:         c.BatchSizeLimit,
		TransactionSizeLimit:   c.TransactionSizeLimit,
		TransactionGasLimit:    c.TransactionGasLimit,
		FeeReceiver:            common.HexToAddress(c.FeeReceiver),
		TargetAddress:          common.HexToAddress(c.TargetAddress),
		TargetFunctionSelector: selector,
		ExecutionTimeout:       c.ExecutionTimeout,
	}
}
