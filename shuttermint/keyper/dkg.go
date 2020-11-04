package keyper

import (
	"crypto/rand"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

// DKGInstance represents the state of a single keyper participating in a DKG process.
type DKGInstance struct {
	Eon    uint64
	Config contract.BatchConfig

	ms MessageSender

	Polynomial *crypto.Polynomial
}

// NewDKGInstance creates a new dkg instance with initialized local random values.
func NewDKGInstance(eon uint64, config contract.BatchConfig, ms MessageSender) (*DKGInstance, error) {
	polyBase, err := crypto.RandomPolynomialBase(rand.Reader)
	if err != nil {
		return nil, err
	}
	polynomial, err := crypto.RandomPolynomial(rand.Reader, config.Threshold, polyBase)
	if err != nil {
		return nil, err
	}

	dkg := DKGInstance{
		Eon:    eon,
		Config: config,

		ms: ms,

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
	for i, keyper := range dkg.Config.Keypers {
		eval := dkg.Polynomial.EvalForKeyper(i)
		encryptedEval := eval.Bytes() // TODO: encrypt, maybe with Tendermint key?
		msg := NewPolyEvalMsg(dkg.Eon, keyper, encryptedEval)
		err := dkg.ms.SendMessage(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
