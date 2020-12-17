package shutterevents

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// KeyperIndex returns the index of the keyper identified by the given address
func (bc *BatchConfig) KeyperIndex(address common.Address) (uint64, bool) {
	for i, k := range bc.Keypers {
		if k == address {
			return uint64(i), true
		}
	}
	return 0, false
}

// IsKeyper checks if the given address is a keyper
func (bc *BatchConfig) IsKeyper(candidate common.Address) bool {
	_, ok := bc.KeyperIndex(candidate)
	return ok
}

// EnsureValid checks if the BatchConfig is valid and returns an error if it's not valid
func (bc *BatchConfig) EnsureValid() error {
	if len(bc.Keypers) == 0 {
		return fmt.Errorf("no keypers in batch config")
	}
	if bc.Threshold == 0 {
		return fmt.Errorf("threshold must not be zero")
	}
	if int(bc.Threshold) > len(bc.Keypers) {
		return fmt.Errorf("threshold too high")
	}
	// XXX maybe we should check for duplicate addresses
	return nil
}

// Message converts the batch config to a shmsg.Message
func (bc *BatchConfig) Message() *shmsg.Message {
	return shmsg.NewBatchConfig(
		bc.StartBatchIndex,
		bc.Keypers,
		bc.Threshold,
		bc.ConfigContractAddress,
		bc.ConfigIndex,
		bc.Started,
		bc.ValidatorsUpdated)
}

// BatchConfigFromMessage extracts the batch config received in a message
func BatchConfigFromMessage(m *shmsg.BatchConfig) (BatchConfig, error) {
	var keypers []common.Address
	for _, b := range m.Keypers {
		if len(b) != common.AddressLength {
			return BatchConfig{}, fmt.Errorf("keyper address has invalid length")
		}
		keypers = append(keypers, common.BytesToAddress(b))
	}

	if len(m.ConfigContractAddress) != common.AddressLength {
		return BatchConfig{}, fmt.Errorf(
			"config contract address has invalid length (%d instead of %d)",
			len(m.ConfigContractAddress),
			common.AddressLength,
		)
	}
	configContractAddress := common.BytesToAddress(m.ConfigContractAddress)

	bc := BatchConfig{
		StartBatchIndex:       m.StartBatchIndex,
		Keypers:               keypers,
		Threshold:             m.Threshold,
		ConfigContractAddress: configContractAddress,
		ConfigIndex:           m.ConfigIndex,
		Started:               m.Started,
		ValidatorsUpdated:     m.ValidatorsUpdated,
	}
	return bc, nil
}
