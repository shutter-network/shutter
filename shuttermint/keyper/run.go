package keyper

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/brainbot-com/shutter/shuttermint/app"
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

// NewEncryptionKeyAttestation creates a new EncryptionKeyAttestation with the given values wrapped
// in a shmsg.Message
func NewEncryptionKeyAttestation(
	batchIndex uint64,
	encryptionKey *ecdsa.PublicKey,
	configContractAddress common.Address,
	signature []byte,
) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_EncryptionKeyAttestation{
			EncryptionKeyAttestation: &shmsg.EncryptionKeyAttestation{
				BatchIndex:            batchIndex,
				Key:                   crypto.FromECDSAPub(encryptionKey),
				ConfigContractAddress: configContractAddress.Bytes(),
				Signature:             signature,
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
func NewBatchState(bp BatchParams, kc KeyperConfig, ms *MessageSender, cc *ContractCaller) BatchState {
	numKeypers := len(bp.BatchConfig.Keypers)
	pubkeyGenerated := make(chan PubkeyGeneratedEvent, 1)
	privkeyGenerated := make(chan PrivkeyGeneratedEvent, 1)
	encryptionKeySignatureAdded := make(chan EncryptionKeySignatureAddedEvent, numKeypers)
	return BatchState{
		BatchParams:                 bp,
		KeyperConfig:                kc,
		MessageSender:               ms,
		ContractCaller:              cc,
		pubkeyGenerated:             pubkeyGenerated,
		privkeyGenerated:            privkeyGenerated,
		encryptionKeySignatureAdded: encryptionKeySignatureAdded,
		startBlockSeen:              make(chan struct{}),
		endBlockSeen:                make(chan struct{}),
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
	case EncryptionKeySignatureAddedEvent:
		select {
		case batch.encryptionKeySignatureAdded <- e:
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
	case <-batch.endBlockSeen:
		log.Print("Timeout while waiting for public key generation to finish", batch.BatchParams)
		return PubkeyGeneratedEvent{}, fmt.Errorf("timeout while waiting for public key generation to finish")
	}
}

func (batch *BatchState) sendPublicKeyCommitment(key *ecdsa.PrivateKey) error {
	msg := NewPublicKeyCommitment(batch.BatchParams.BatchIndex, key)
	log.Print("Generated pubkey", batch.BatchParams)
	return batch.MessageSender.SendMessage(msg)
}

func (batch *BatchState) sendEncryptionKeySignature(encryptionKey *ecdsa.PublicKey) error {
	preimage := app.EncryptionKeyPreimage(
		crypto.FromECDSAPub(encryptionKey),
		batch.BatchParams.BatchIndex,
		batch.KeyperConfig.ConfigContractAddress,
	)
	hash := crypto.Keccak256Hash(preimage)
	sig, err := crypto.Sign(hash.Bytes(), batch.KeyperConfig.SigningKey)
	if err != nil {
		return err
	}
	msg := NewEncryptionKeyAttestation(
		batch.BatchParams.BatchIndex,
		encryptionKey,
		batch.KeyperConfig.ConfigContractAddress,
		sig,
	)
	return batch.MessageSender.SendMessage(msg)
}

func (batch *BatchState) collectEncryptionKeySignatureAddedEvents() ([]EncryptionKeySignatureAddedEvent, error) {
	events := []EncryptionKeySignatureAddedEvent{}
	for {
		select {
		case ev := <-batch.encryptionKeySignatureAdded:
			events = append(events, ev)
			if len(events) >= int(batch.BatchParams.BatchConfig.Threshold.Int64()) {
				return events, nil
			}
			// case <-time.After(time.Until(batch.BatchParams.PrivateKeyGenerationStartTime)):
			//	return events, fmt.Errorf("timeout while waiting for encryption key signatures")
		}
	}
}

func (batch *BatchState) broadcastEncryptionKey(key *ecdsa.PrivateKey) error {
	encryptionKey := crypto.FromECDSAPub(&key.PublicKey)
	bp := batch.BatchParams
	keyperIndex, ok := bp.BatchConfig.KeyperIndex(batch.KeyperAddress())
	if !ok {
		return fmt.Errorf("not a keyper") // XXX we really should lookup our address earlier!
	}

	events, err := batch.collectEncryptionKeySignatureAddedEvents()
	if err != nil {
		return err
	}
	indices := []uint64{}
	sigs := [][]byte{}
	for _, ev := range events {
		sigs = append(sigs, ev.Signature)
		indices = append(indices, ev.KeyperIndex)
	}

	return batch.ContractCaller.BroadcastEncryptionKey(
		keyperIndex,
		batch.BatchParams.BatchIndex,
		encryptionKey,
		indices,
		sigs,
	)
}

func (batch *BatchState) sendSecretShare(key *ecdsa.PrivateKey) error {
	msg := NewSecretShare(batch.BatchParams.BatchIndex, key)
	log.Print("Generated privkey", batch.BatchParams)
	return batch.MessageSender.SendMessage(msg)
}

func (batch *BatchState) NewBlockHeader(header *types.Header) {
	blockNumber := header.Number.Uint64()
	if blockNumber >= batch.BatchParams.StartBlock {
		batch.startBlockSeenOnce.Do(func() { close(batch.startBlockSeen) })
	}
	if blockNumber >= batch.BatchParams.EndBlock {
		batch.endBlockSeenOnce.Do(func() { close(batch.endBlockSeen) })
	}
}

func (batch *BatchState) waitForStartBlock() {
	<-batch.startBlockSeen
}

func (batch *BatchState) waitForEndBlock() {
	<-batch.endBlockSeen
}

// Run runs the key generation for the given batch
func (batch *BatchState) Run() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Error while generating key: %s", err)
		return
	}
	log.Printf("Waiting for start block %d for batch #%d", batch.BatchParams.StartBlock, batch.BatchParams.BatchIndex)
	batch.waitForStartBlock()
	log.Print("Starting key generation process", batch.BatchParams)

	err = batch.sendPublicKeyCommitment(key)
	if err != nil {
		log.Printf("Error while trying to send message: %s", err)
		return
	}

	ev, err := batch.waitPubkeyGenerated()
	if err != nil {
		log.Printf("Error while waiting for public key generation: %s", err)
		return
	}

	err = batch.sendEncryptionKeySignature(ev.Pubkey)
	if err != nil {
		log.Printf("Error while trying to send encryption key signature: %s", err)
	}

	err = batch.broadcastEncryptionKey(key)
	if err != nil {
		log.Printf("Error while trying to broadcast encryption key: %s", err)
		return
	}

	batch.waitForEndBlock()

	err = batch.sendSecretShare(key)
	if err != nil {
		log.Printf("Error while trying to send secret share: %s", err)
		return
	}
}

// KeyperAddress returns the keyper's Ethereum address.
func (batch *BatchState) KeyperAddress() common.Address {
	return crypto.PubkeyToAddress(batch.KeyperConfig.SigningKey.PublicKey)
}
