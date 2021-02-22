package sandbox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

const selectorLength = 4

// chksumAddr is used internally to serialize addresses as checksum addresses when writing JSON
// files
type chksumAddr common.Address

func (addr chksumAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(common.Address(addr).Hex())
}

func (addr *chksumAddr) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	a := common.HexToAddress(s)
	if a.Hex() != s {
		return fmt.Errorf("invalid address: %s", s)
	}
	copy(addr[:], a[:])
	return nil
}

// ContractsJSON stores the addresses of all contracts.
type ContractsJSON struct {
	ConfigContract        common.Address
	KeyBroadcastContract  common.Address
	FeeBankContract       common.Address
	BatcherContract       common.Address
	ExecutorContract      common.Address
	TokenContract         common.Address
	DepositContract       common.Address
	KeyperSlasherContract common.Address
	TargetContract        common.Address
}

type contractsJSON struct {
	ConfigContract        chksumAddr
	KeyBroadcastContract  chksumAddr
	FeeBankContract       chksumAddr
	BatcherContract       chksumAddr
	ExecutorContract      chksumAddr
	TokenContract         chksumAddr
	DepositContract       chksumAddr
	KeyperSlasherContract chksumAddr
	TargetContract        chksumAddr
}

// MarshalJSON makes us output checksum addresses when marshaling as json
func (c ContractsJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		contractsJSON{
			chksumAddr(c.ConfigContract),
			chksumAddr(c.KeyBroadcastContract),
			chksumAddr(c.FeeBankContract),
			chksumAddr(c.BatcherContract),
			chksumAddr(c.ExecutorContract),
			chksumAddr(c.TokenContract),
			chksumAddr(c.DepositContract),
			chksumAddr(c.KeyperSlasherContract),
			chksumAddr(c.TargetContract),
		},
	)
}

func (c *ContractsJSON) UnmarshalJSON(data []byte) error {
	tmp := contractsJSON{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	c.ConfigContract = common.Address(tmp.ConfigContract)
	c.KeyBroadcastContract = common.Address(tmp.KeyBroadcastContract)
	c.FeeBankContract = common.Address(tmp.FeeBankContract)
	c.BatcherContract = common.Address(tmp.BatcherContract)
	c.ExecutorContract = common.Address(tmp.ExecutorContract)
	c.TokenContract = common.Address(tmp.TokenContract)
	c.DepositContract = common.Address(tmp.DepositContract)
	c.KeyperSlasherContract = common.Address(tmp.KeyperSlasherContract)
	c.TargetContract = common.Address(tmp.TargetContract)
	return nil
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
		return nil, errors.Wrap(err, "malformed contracts JSON")
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
