package keyper

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/ecies"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/puredkg"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// DKGInstance represents the state of a single keyper participating in a DKG process.
type DKGInstance struct {
	Eon          uint64
	BatchConfig  contract.BatchConfig
	KeyperConfig KeyperConfig

	ms                   MessageSender
	keyperEncryptionKeys map[common.Address]*ecies.PublicKey
	pure                 puredkg.PureDKG
}

// FindAddressIndex returns the index of the given address inside the slice of addresses or returns
// an error, if the slice does not contain the given address
func FindAddressIndex(addresses []common.Address, addr common.Address) (int, error) {
	for i, a := range addresses {
		if a == addr {
			return i, nil
		}
	}
	return -1, errors.New("address not found")
}

// NewDKGInstance creates a new dkg instance with initialized local random values.
func NewDKGInstance(
	eon uint64,
	batchConfig contract.BatchConfig,
	keyperConfig KeyperConfig,
	ms MessageSender,
	keyperEncryptionKeys map[common.Address]*ecies.PublicKey,
) (*DKGInstance, error) {
	keyperIndex, err := FindAddressIndex(batchConfig.Keypers, keyperConfig.Address())
	if err != nil {
		return nil, fmt.Errorf("internal error: not a keyper: %s", keyperConfig.Address().Hex())
	}
	dkg := DKGInstance{
		Eon:          eon,
		BatchConfig:  batchConfig,
		KeyperConfig: keyperConfig,

		ms:                   ms,
		keyperEncryptionKeys: keyperEncryptionKeys,
		pure:                 puredkg.NewPureDKG(eon, uint64(len(batchConfig.Keypers)), batchConfig.Threshold, uint64(keyperIndex)),
	}
	return &dkg, nil
}

// Run everything.
func (dkg *DKGInstance) Run(ctx context.Context) error {
	err := dkg.startPhase1Dealing(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (dkg *DKGInstance) FindKeyperIndex(addr common.Address) (int, error) {
	return FindAddressIndex(dkg.BatchConfig.Keypers, addr)
}

// startPhase1Dealing starts the dealing phase. It will send the gammas (commitment) and the
// encrypted poly eval messages to each keyper
func (dkg *DKGInstance) startPhase1Dealing(ctx context.Context) error {
	polyCommitment, polyEvals, err := dkg.pure.StartPhase1Dealing()
	if err != nil {
		return err
	}
	msg := NewPolyCommitmentMsg(dkg.Eon, polyCommitment.Gammas)
	log.Printf("Sending Gammas")
	err = dkg.ms.SendMessage(ctx, msg)
	if err != nil {
		return err
	}
	return dkg.sendPolyEvals(ctx, polyEvals)
}

func (dkg *DKGInstance) makePolyEvalMessage(polyEvals []puredkg.PolyEvalMsg) (*shmsg.Message, error) {
	receivers := []common.Address{}
	evals := [][]byte{}

	for i, polyEval := range polyEvals {
		keyper := dkg.BatchConfig.Keypers[i]
		if keyper == dkg.KeyperConfig.Address() {
			continue
		}

		encryptionKey, ok := dkg.keyperEncryptionKeys[keyper]
		if !ok {
			log.Printf("no key available, cannot send message to keyper %s", keyper.Hex())
			continue // don't send them a message if their encryption public key is unknown
		}

		evalBytes := polyEval.Eval.Bytes()
		encryptedEval, err := ecies.Encrypt(rand.Reader, encryptionKey, evalBytes, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt message: %w", err)
		}

		receivers = append(receivers, keyper)
		evals = append(evals, encryptedEval)
	}
	if uint64(len(evals)) < dkg.BatchConfig.Threshold {
		return nil, fmt.Errorf("makePolyEvalMessage: need at least %d keys, only have %d", dkg.BatchConfig.Threshold, len(evals))
	}
	return NewPolyEvalMsg(dkg.Eon, receivers, evals), nil
}

// sendPolyEvals sends the corresponding polynomial evaluation to each keyper, including ourselves.
func (dkg *DKGInstance) sendPolyEvals(ctx context.Context, polyEvals []puredkg.PolyEvalMsg) error {
	log.Printf("Sending PolyEvals to keypers")
	msg, err := dkg.makePolyEvalMessage(polyEvals)
	if err != nil {
		return err
	}
	err = dkg.ms.SendMessage(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (dkg *DKGInstance) dispatchShuttermintEvent(ev IEvent) {
	switch e := ev.(type) {
	case PolyCommitmentRegisteredEvent:
		senderIndex, err := dkg.FindKeyperIndex(e.Sender)
		if err != nil {
			log.Printf("Could not handle poly commitment message. sender is not a keyper")
			return
		}
		m := puredkg.PolyCommitmentMsg{
			Eon:    e.Eon,
			Sender: uint64(senderIndex),
			Gammas: e.Gammas,
		}
		err = dkg.pure.HandlePolyCommitmentMsg(m)
		if err != nil {
			log.Printf("Could not handle poly commitment message: %+v %s", m, err)
			return
		}
	case PolyEvalRegisteredEvent:
		senderIndex, err := dkg.FindKeyperIndex(e.Sender)
		if err != nil {
			log.Printf("Could not handle poly eval message. sender is not a keyper")
			return
		}
		_ = senderIndex
		log.Printf("XXX handle poly eval registered: %+v", e)
	default:
		panic("unknown event type, cannot dispatch")
	}
}
