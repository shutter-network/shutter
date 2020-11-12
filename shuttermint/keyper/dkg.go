package keyper

import (
	"crypto/rand"
	"fmt"

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
	}
	return &dkg, nil
}

// Run everything.
func (dkg *DKGInstance) Run() error {
	err := dkg.sendGammas()
	if err != nil {
		return err
	}
	err = dkg.sendPolyEvals()
	if err != nil {
		return err
	}

	return nil
}

// sendGammas broadcasts the gamma values.
func (dkg *DKGInstance) sendGammas() error {
	msg := NewPolyCommitmentMsg(dkg.Eon, dkg.Polynomial.Gammas())
	return dkg.ms.SendMessage(msg)
}

// sendPolyEvals sends the corresponding polynomial evaluation to each keyper, including ourselves.
func (dkg *DKGInstance) sendPolyEvals() error {
	for i, keyper := range dkg.BatchConfig.Keypers {
		encryptionKey, ok := dkg.keyperEncryptionKeys[keyper]
		if !ok {
			continue // don't send them a message if there encryption public key is unknown
		}

		eval := dkg.Polynomial.EvalForKeyper(i)
		evalBytes := eval.Bytes()
		encryptedEval, err := ecies.Encrypt(rand.Reader, encryptionKey, evalBytes, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to encrypt message: %w", err)
		}

		msg := NewPolyEvalMsg(dkg.Eon, keyper, encryptedEval)
		err = dkg.ms.SendMessage(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
