package keyper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// SleepUntil pauses the current goroutine until the given time is reached
func SleepUntil(t time.Time) {
	time.Sleep(time.Until(t))
}

// NewBatchState created a new BatchState object with the given parameters.
func NewBatchState(
	bp BatchParams,
	kc KeyperConfig,
	ms MessageSender,
	cc *ContractCaller,
	cipherExecutionParams chan CipherExecutionParams,
) BatchState {
	numKeypers := len(bp.BatchConfig.Keypers)
	decryptionSignatureAdded := make(chan shutterevents.DecryptionSignatureEvent, numKeypers)
	return BatchState{
		BatchParams:               bp,
		KeyperConfig:              kc,
		MessageSender:             ms,
		ContractCaller:            cc,
		decryptionSignatureAdded:  decryptionSignatureAdded,
		cipherExecutionParams:     cipherExecutionParams,
		startBlockSeen:            make(chan struct{}),
		endBlockSeen:              make(chan struct{}),
		executionTimeoutBlockSeen: make(chan struct{}),
	}
}

func (batch *BatchState) dispatchShuttermintEvent(ev shutterevents.IEvent) {
	switch e := ev.(type) {
	case shutterevents.DecryptionSignatureEvent:
		select {
		case batch.decryptionSignatureAdded <- e:
		default:
			fmt.Printf("Unexpected error: decryptionSignatureAdded channel blocked")
		}
	default:
		panic("unknown type")
	}
}

func (batch *BatchState) collectDecryptionSignatureEvents(
	cipherBatchHash common.Hash,
	decryptionKey *ecdsa.PrivateKey,
	batchHash common.Hash,
) ([]shutterevents.DecryptionSignatureEvent, error) {
	events := []shutterevents.DecryptionSignatureEvent{}
	for {
		select {
		case ev := <-batch.decryptionSignatureAdded:
			recoveredSigner, err := RecoverDecryptionSignatureSigner(
				ev.Signature,
				batch.KeyperConfig.BatcherContractAddress,
				cipherBatchHash,
				decryptionKey,
				batchHash,
			)
			if err != nil {
				log.Printf("failed to recover signer of decryption signature message from %s", ev.Sender.Hex())
				continue
			}
			if recoveredSigner != ev.Sender {
				log.Printf("failed to verify decryption signature message from %s", ev.Sender.Hex())
				continue
			}

			events = append(events, ev)
			if len(events) >= int(batch.BatchParams.BatchConfig.Threshold) {
				return events, nil
			}
		case <-batch.executionTimeoutBlockSeen:
			return events, fmt.Errorf("timeout while waiting for decryption signatures")
		}
	}
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

func (batch *BatchState) sendDecryptionSignature(cipherTxs [][]byte, decryptedTxs [][]byte, decryptionKey *ecdsa.PrivateKey) error {
	decryptionSignature, err := ComputeDecryptionSignature(
		batch.KeyperConfig.SigningKey,
		batch.KeyperConfig.BatcherContractAddress,
		ComputeBatchHash(cipherTxs),
		decryptionKey,
		ComputeBatchHash(decryptedTxs),
	)
	if err != nil {
		return fmt.Errorf("error computing dercyption signature: %s", err)
	}

	msg := shmsg.NewDecryptionSignature(batch.BatchParams.BatchIndex, decryptionSignature)
	return batch.MessageSender.SendMessage(context.TODO(), msg)
}

// NewBlockHeader is called whenever a new block header arrives.
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
	log.Printf("Waiting for start block %d for batch #%d", batch.BatchParams.StartBlock, batch.BatchParams.BatchIndex)
	batch.waitForStartBlock()
	log.Printf("Starting key generation process for batch #%d", batch.BatchParams.BatchIndex)

	batch.waitForEndBlock()

	cipherTxs, err := batch.downloadTransactions()
	cipherBatchHash := ComputeBatchHash(cipherTxs)
	if err != nil {
		log.Printf("Error while downloading transactions: %s", err)
		return
	}

	// XXX use random key for decryption until proper key generation is implemented
	decryptionKey, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Error generating decryption key")
	}

	decryptedTxs := DecryptTransactions(decryptionKey, cipherTxs)
	batchHash := ComputeBatchHash(decryptedTxs)
	err = batch.sendDecryptionSignature(cipherTxs, decryptedTxs, decryptionKey)
	if err != nil {
		log.Printf("Error sending decryption signature: %s", err)
		return
	}

	decryptionSignatureEvents, err := batch.collectDecryptionSignatureEvents(cipherBatchHash, decryptionKey, batchHash)
	if err != nil {
		log.Printf("Error collecting decryption signatures: %s", err)
		return
	}

	decryptionSignerIndices, decryptionSignatures, err := batch.signerIndicesAndSignaturesFromEvents(decryptionSignatureEvents)
	if err != nil {
		log.Printf("Error processing decryption signature events: %s", err)
		return
	}
	cipherExecutionParams := CipherExecutionParams{
		BatchIndex:              batch.BatchParams.BatchIndex,
		CipherBatchHash:         cipherBatchHash,
		DecryptionKey:           decryptionKey,
		DecryptedTxs:            decryptedTxs,
		DecryptionSignerIndices: decryptionSignerIndices,
		DecryptionSignatures:    decryptionSignatures,
	}
	batch.cipherExecutionParams <- cipherExecutionParams
}

// KeyperAddress returns the keyper's Ethereum address.
func (batch *BatchState) KeyperAddress() common.Address {
	return crypto.PubkeyToAddress(batch.KeyperConfig.SigningKey.PublicKey)
}

func (batch *BatchState) signerIndicesAndSignaturesFromEvents(events []shutterevents.DecryptionSignatureEvent) ([]uint64, [][]byte, error) {
	sliceIndices := []uint64{}
	signerIndices := []uint64{}
	signatures := [][]byte{}
	for i, ev := range events {
		index, isKeyper := batch.BatchParams.BatchConfig.KeyperIndex(ev.Sender)
		if !isKeyper {
			return []uint64{}, [][]byte{}, fmt.Errorf("%s is not a keyper", ev.Sender.Hex())
		}

		sliceIndices = append(sliceIndices, uint64(i))
		signerIndices = append(signerIndices, index)
		signatures = append(signatures, ev.Signature)
	}

	// sort by signer index
	sort.Slice(sliceIndices, func(i, j int) bool {
		return signerIndices[i] < signerIndices[j]
	})
	signerIndicesSorted := []uint64{}
	signaturesSorted := [][]byte{}
	for _, i := range sliceIndices {
		signerIndicesSorted = append(signerIndicesSorted, signerIndices[i])
		signaturesSorted = append(signaturesSorted, signatures[i])
	}

	return signerIndicesSorted, signaturesSorted, nil
}
