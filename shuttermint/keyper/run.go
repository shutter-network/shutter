package keyper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/brainbot-com/shutter/shuttermint/app"
	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// SleepUntil pauses the current goroutine until the given time is reached
func SleepUntil(t time.Time) {
	time.Sleep(time.Until(t))
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
		executionTimeoutBlockSeen:   make(chan struct{}),
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
		log.Printf("Received PubkeyGenerated event")
		return ev, nil
	case <-batch.endBlockSeen:
		log.Print("Timeout while waiting for public key generation to finish", batch.BatchParams)
		return PubkeyGeneratedEvent{}, fmt.Errorf("timeout while waiting for public key generation to finish")
	}
}

// waitPrivkeyGenerated waits for a PrivkeyGeneratedEvent to be put into the privkeyGenerated
// channel. We do expect exactly one instance to be put there via a call to
// dispatchShuttermintEvent by the keyper from a different goroutine. The call times out when with
// the execution timeout.
func (batch *BatchState) waitPrivkeyGenerated() (PrivkeyGeneratedEvent, error) {
	select {
	case ev := <-batch.privkeyGenerated:
		log.Printf("Received PrivkeyGenerated event")
		return ev, nil
	case <-batch.executionTimeoutBlockSeen:
		log.Print("Timeout while waiting for private key generation to finish", batch.BatchParams)
		return PrivkeyGeneratedEvent{}, fmt.Errorf("timeout while waiting for private key generation to finish")
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
			if len(events) >= int(batch.BatchParams.BatchConfig.Threshold) {
				return events, nil
			}
		case <-batch.endBlockSeen:
			return events, fmt.Errorf("timeout while waiting for encryption key signatures")
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

	auth, err := batch.ContractCaller.Auth()
	if err != nil {
		return err
	}
	auth.GasLimit = 100000
	tx, err := batch.ContractCaller.KeyBroadcastContract.BroadcastEncryptionKey2(
		auth,
		keyperIndex,
		batch.BatchParams.BatchIndex,
		encryptionKey,
		indices,
		sigs,
	)
	if err != nil {
		return err
	}
	log.Printf("Encryption key broadcasted with tx %s", tx.Hash().Hex())

	return nil
}

func (batch *BatchState) sendSecretShare(key *ecdsa.PrivateKey) error {
	msg := NewSecretShare(batch.BatchParams.BatchIndex, key)
	log.Print("Generated privkey", batch.BatchParams)
	return batch.MessageSender.SendMessage(msg)
}

func (batch *BatchState) downloadTransactions() ([][]byte, error) {
	filterOpts := &bind.FilterOpts{
		Context: context.TODO(),
		Start:   batch.BatchParams.StartBlock,
		End:     &batch.BatchParams.EndBlock,
	}
	itr, err := batch.ContractCaller.BatcherContract.FilterTransactionAdded(filterOpts)
	if err != nil {
		return [][]byte{}, err
	}

	txs := [][]byte{}
	var batchHash common.Hash
	for itr.Next() {
		if itr.Event.BatchIndex != batch.BatchParams.BatchIndex {
			continue
		}
		if itr.Event.TransactionType != contract.TransactionTypeCipher {
			continue
		}

		batchHash = itr.Event.BatchHash
		txs = append(txs, itr.Event.Transaction)
	}
	if batchHash != ComputeBatchHash(txs) {
		return [][]byte{}, fmt.Errorf("failed to verify batch hash of batch #%d", batch.BatchParams.BatchIndex)
	}

	return txs, nil
}

func (batch *BatchState) NewBlockHeader(header *types.Header) {
	blockNumber := header.Number.Uint64()
	if blockNumber >= batch.BatchParams.StartBlock {
		batch.startBlockSeenOnce.Do(func() { close(batch.startBlockSeen) })
	}
	if blockNumber >= batch.BatchParams.EndBlock {
		batch.endBlockSeenOnce.Do(func() { close(batch.endBlockSeen) })
	}
	if blockNumber >= batch.BatchParams.ExecutionTimeoutBlock {
		batch.executionTimeoutBlockSeenOnce.Do(func() { close(batch.executionTimeoutBlockSeen) })
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

	_, err = batch.waitPrivkeyGenerated()
	if err != nil {
		log.Printf("Error while waiting for decryption key generation: %s", err)
	}

	_, err = batch.downloadTransactions()
	if err != nil {
		log.Printf("Error while downloading transactions: %s", err)
	}
}

// KeyperAddress returns the keyper's Ethereum address.
func (batch *BatchState) KeyperAddress() common.Address {
	return crypto.PubkeyToAddress(batch.KeyperConfig.SigningKey.PublicKey)
}
