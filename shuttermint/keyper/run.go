package keyper

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

// NewBatchState created a new BatchState object with the given parameters.
func NewBatchState(bp BatchParams, ms *MessageSender, cc *ContractCaller) BatchState {
	pubkeyGenerated := make(chan PubkeyGeneratedEvent, 1)
	privkeyGenerated := make(chan PrivkeyGeneratedEvent, 1)
	return BatchState{
		BatchParams:      bp,
		MessageSender:    ms,
		ContractCaller:   cc,
		pubkeyGenerated:  pubkeyGenerated,
		privkeyGenerated: privkeyGenerated,
	}
}

func (batch *BatchState) dispatchShuttermintEvent(ev IEvent) {
	switch e := ev.(type) {
	case PubkeyGeneratedEvent:
		select {
		case batch.pubkeyGenerated <- e:
		default:
			// this should not happen
		}
	case PrivkeyGeneratedEvent:
		select {
		case batch.privkeyGenerated <- e:
		default:
			// this should not happen
		}
	default:
		panic("unknown type")
	}
}

// waitPubkeyGenerated waits for a PubkeyGeneratedEvent to be put into the pubkeyGenerated channel.
// We do expect exactly one instance to be put there via a call to dispatchShuttermintEvent by the
// keyper from a different goroutine.
func (batch *BatchState) waitPubkeyGenerated() (PubkeyGeneratedEvent, error) {
	select {
	case ev := <-batch.pubkeyGenerated:
		log.Printf("Got event %+v", ev)
		return ev, nil
	case <-time.After(time.Until(batch.BatchParams.PrivateKeyGenerationStartTime)):
		log.Print("Timeout while waiting for public key generation to finish", batch.BatchParams)
		return PubkeyGeneratedEvent{}, fmt.Errorf("timeout while waiting for public key generation to finish")
	}
}

func (batch *BatchState) sendPublicKeyCommitment(key *ecdsa.PrivateKey) error {
	msg := NewPublicKeyCommitment(batch.BatchParams.BatchIndex, key)
	log.Print("Generated pubkey", batch.BatchParams)
	return batch.MessageSender.SendMessage(msg)
}

func (batch *BatchState) broadcastEncryptionKey(key *ecdsa.PrivateKey) error {
	encryptionKey := crypto.FromECDSAPub(&key.PublicKey)
	bp := batch.BatchParams
	keyperIndex, ok := bp.BatchConfig.KeyperIndex(bp.KeyperAddress)
	if !ok {
		return fmt.Errorf("not a keyper") // XXX we really should lookup our address earlier!
	}
	return batch.ContractCaller.BroadcastEncryptionKey(keyperIndex, batch.BatchParams.BatchIndex, encryptionKey, []uint64{}, [][]byte{})
}

func (batch *BatchState) sendSecretShare(key *ecdsa.PrivateKey) error {
	msg := NewSecretShare(batch.BatchParams.BatchIndex, key)
	log.Print("Generated privkey", batch.BatchParams)
	return batch.MessageSender.SendMessage(msg)
}

func (batch *BatchState) NewBlockHeader(header *types.Header) {
}

// Run runs the key generation for the given batch
func (batch *BatchState) Run() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Error while generating key: %s", err)
		return
	}

	// Wait for the start time
	SleepUntil(batch.BatchParams.PublicKeyGenerationStartTime)
	log.Print("Starting key generation process", batch.BatchParams)

	err = batch.sendPublicKeyCommitment(key)
	if err != nil {
		log.Printf("Error while trying to send message: %s", err)
		return
	}

	_, err = batch.waitPubkeyGenerated()
	if err != nil {
		log.Printf("Error while waiting for public key generation: %s", err)
		return
	}

	err = batch.broadcastEncryptionKey(key)
	if err != nil {
		log.Printf("Error while trying to broadcast encryption key: %s", err)
		return
	}

	SleepUntil(batch.BatchParams.PrivateKeyGenerationStartTime)

	err = batch.sendSecretShare(key)
	if err != nil {
		log.Printf("Error while trying to send secret share: %s", err)
		return
	}
}
