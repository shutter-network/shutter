package keyper

import (
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// SleepUntil pauses the current goroutine until the given time is reached
func SleepUntil(t time.Time) {
	time.Sleep(time.Until(t))
}

// NewBatchConfig creates a new BatchConfig with the given values wrapped in a shmsg.Message
func NewBatchConfig(startBatchIndex uint64, keypers []common.Address, threshold uint32) *shmsg.Message {
	var addresses [][]byte
	for _, k := range keypers {
		addresses = append(addresses, k.Bytes())
	}
	return &shmsg.Message{
		Payload: &shmsg.Message_BatchConfig{
			BatchConfig: &shmsg.BatchConfig{
				StartBatchIndex: startBatchIndex,
				Keypers:         addresses,
				Threshold:       threshold,
			},
		},
	}
}

// NewPublicKeyCommitment creates a new PublicKeyCommitment with the given values wrapped in a shmsg.Message
func NewPublicKeyCommitment(batchIndex uint64, privkey *ecdsa.PrivateKey) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_PublicKeyCommitment{
			PublicKeyCommitment: &shmsg.PublicKeyCommitment{
				BatchIndex: batchIndex,
				Commitment: crypto.FromECDSAPub(&privkey.PublicKey),
			},
		},
	}
}

// NewSecretShare creates a new SecretShare with the given values wrapped in a shmsg.Message
func NewSecretShare(batchIndex uint64, privkey *ecdsa.PrivateKey) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_SecretShare{
			SecretShare: &shmsg.SecretShare{
				BatchIndex: batchIndex,
				Privkey:    crypto.FromECDSA(privkey),
			},
		},
	}
}

// Run runs the key generation for the given batch
func Run(params BatchParams, ms MessageSender, cc ContractCaller, events <-chan IEvent) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return
	}

	// Wait for the start time
	SleepUntil(params.PublicKeyGenerationStartTime)
	log.Print("Starting key generation process", params)
	msg := NewPublicKeyCommitment(params.BatchIndex, key)
	log.Print("Generated pubkey", params)
	err = ms.SendMessage(msg)
	if err != nil {
		log.Print("Error while trying to send message:", err)
		return
	}

	select {
	case ev := <-events:
		log.Printf("Got event %+v", ev)
	case <-time.After(time.Until(params.PrivateKeyGenerationStartTime)):
		log.Print("Timeout while waiting for public key generation to finish", params)
		return
	}

	encryptionKey := crypto.FromECDSAPub(&key.PublicKey)
	err = cc.BroadcastEncryptionKey(0, params.BatchIndex, encryptionKey, []uint64{}, [][]byte{})
	if err != nil {
		log.Print("Error while trying to broadcast encryption key:", err)
		return
	}

	SleepUntil(params.PrivateKeyGenerationStartTime)
	msg = NewSecretShare(params.BatchIndex, key)
	log.Print("Generated privkey", params)
	err = ms.SendMessage(msg)
	if err != nil {
		log.Println("Error while trying to send message:", err)
		return
	}
}
