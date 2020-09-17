package app

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

// Message converts the batch config to a shmsg.Message
func (bc *BatchConfig) Message() shmsg.Message {
	var keypers [][]byte
	for _, keyperAddress := range bc.Keypers {
		keypers = append(keypers, keyperAddress.Bytes())
	}

	msg := shmsg.Message_BatchConfig{
		BatchConfig: &shmsg.BatchConfig{
			StartBatchIndex: bc.StartBatchIndex,
			Keypers:         keypers,
			Threshold:       bc.Threshold,
		},
	}
	return shmsg.Message{Payload: &msg}
}

// BatchConfigFromMessage extracts the batch config received in a message
func BatchConfigFromMessage(m *shmsg.BatchConfig) (BatchConfig, error) {
	var keypers []common.Address
	for _, b := range m.Keypers {
		if len(b) != 20 {
			return BatchConfig{}, fmt.Errorf("Keyper address has invalid length")
		}
		keypers = append(keypers, common.BytesToAddress(b))
	}

	bc := BatchConfig{
		StartBatchIndex: m.StartBatchIndex,
		Keypers:         keypers,
		Threshold:       m.Threshold,
	}
	return bc, nil
}
