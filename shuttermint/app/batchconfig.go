package app

import "github.com/ethereum/go-ethereum/common"
import "github.com/brainbot-com/shutter/shuttermint/shmsg"

// IsKeyper checks if the given address is a keyper
func (bc *BatchConfig) IsKeyper(candidate common.Address) bool {
	for _, k := range bc.Keypers {
		if k == candidate {
			return true
		}
	}
	return false
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
