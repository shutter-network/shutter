package keyper

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/ecies"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

// DKGInstance represents the state of a single keyper participating in a DKG process.
type DKGInstance struct {
	Eon          uint64
	BatchConfig  contract.BatchConfig
	KeyperConfig KeyperConfig

	ms                   MessageSender
	keyperEncryptionKeys map[common.Address]*ecies.PublicKey

	Polynomial *crypto.Polynomial
	Commitment map[common.Address]crypto.Gammas
}

// NewDKGInstance creates a new dkg instance with initialized local random values.
func NewDKGInstance(
	eon uint64,
	batchConfig contract.BatchConfig,
	keyperConfig KeyperConfig,
	ms MessageSender,
	keyperEncryptionKeys map[common.Address]*ecies.PublicKey,
) (*DKGInstance, error) {
	polynomial, err := crypto.RandomPolynomial(rand.Reader, batchConfig.Threshold)
	if err != nil {
		return nil, err
	}

	dkg := DKGInstance{
		Eon:          eon,
		BatchConfig:  batchConfig,
		KeyperConfig: keyperConfig,

		ms:                   ms,
		keyperEncryptionKeys: keyperEncryptionKeys,

		Polynomial: polynomial,
		Commitment: make(map[common.Address]crypto.Gammas),
	}
	return &dkg, nil
}

// Run everything.
func (dkg *DKGInstance) Run(ctx context.Context) error {
	err := dkg.sendGammas(ctx)
	if err != nil {
		return err
	}
	err = dkg.sendPolyEvals(ctx)
	if err != nil {
		return err
	}

	return nil
}

// sendGammas broadcasts the gamma values.
func (dkg *DKGInstance) sendGammas(ctx context.Context) error {
	log.Printf("Sending Gammas")
	msg := NewPolyCommitmentMsg(dkg.Eon, dkg.Polynomial.Gammas())
	return dkg.ms.SendMessage(ctx, msg)
}

// sendPolyEvals sends the corresponding polynomial evaluation to each keyper, including ourselves.
func (dkg *DKGInstance) sendPolyEvals(ctx context.Context) error {
	log.Printf("Sending PolyEvals to keypers")
	receivers := []common.Address{}
	evals := [][]byte{}
	for i, keyper := range dkg.BatchConfig.Keypers {
		if keyper == dkg.KeyperConfig.Address() {
			continue
		}

		encryptionKey, ok := dkg.keyperEncryptionKeys[keyper]
		if !ok {
			log.Printf("no key available, cannot send message to keyper %s", keyper.Hex())
			continue // don't send them a message if their encryption public key is unknown
		}

		eval := dkg.Polynomial.EvalForKeyper(i)
		evalBytes := eval.Bytes()
		encryptedEval, err := ecies.Encrypt(rand.Reader, encryptionKey, evalBytes, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to encrypt message: %w", err)
		}

		receivers = append(receivers, keyper)
		evals = append(evals, encryptedEval)
	}
	msg := NewPolyEvalMsg(dkg.Eon, receivers, evals)
	err := dkg.ms.SendMessage(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (dkg *DKGInstance) dispatchShuttermintEvent(ev IEvent) {
	switch e := ev.(type) {
	case PolyCommitmentRegisteredEvent:
		// XXX we should handle the case where we already have a commitment
		dkg.Commitment[e.Sender] = *e.Gammas
	default:
		panic("unknown event type, cannot dispatch")
	}
}
