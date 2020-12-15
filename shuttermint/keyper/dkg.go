package keyper

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/ecies"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/puredkg"
	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
	"github.com/brainbot-com/shutter/shuttermint/medley"
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

// NewDKGInstance creates a new dkg instance with initialized local random values.
func NewDKGInstance(
	eon uint64,
	batchConfig contract.BatchConfig,
	keyperConfig KeyperConfig,
	ms MessageSender,
	keyperEncryptionKeys map[common.Address]*ecies.PublicKey,
) (*DKGInstance, error) {
	keyperIndex, err := medley.FindAddressIndex(batchConfig.Keypers, keyperConfig.Address())
	if err != nil {
		return nil, fmt.Errorf("new dkg: not a keyper: %s", keyperConfig.Address().Hex())
	}
	// The following checks if we have enough encryption keys.
	// XXX Make sure we are not off by one here and check if keyperEncryptionKeys contains our
	// own key.
	if uint64(len(keyperEncryptionKeys)) < batchConfig.Threshold {
		return nil, fmt.Errorf("new dkg: need at least %d encryption keys, only have %d", batchConfig.Threshold, len(keyperEncryptionKeys))
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
	return medley.FindAddressIndex(dkg.BatchConfig.Keypers, addr)
}

// startPhase1Dealing starts the dealing phase. It will send the gammas (commitment) and the
// encrypted poly eval messages to each keyper
func (dkg *DKGInstance) startPhase1Dealing(ctx context.Context) error {
	polyCommitment, polyEvals, err := dkg.pure.StartPhase1Dealing()
	if err != nil {
		return err
	}
	msg := shmsg.NewPolyCommitment(dkg.Eon, polyCommitment.Gammas)
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
	return shmsg.NewPolyEval(dkg.Eon, receivers, evals), nil
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

func (dkg *DKGInstance) dispatchShuttermintEvent(ev shutterevents.IEvent) {
	switch e := ev.(type) {
	case shutterevents.PolyCommitment:
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
	case shutterevents.PolyEval:
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
