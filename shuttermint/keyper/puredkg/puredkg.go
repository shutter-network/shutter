/*puredkg implements the DKG protocol. It's independent of any transport, but rather expects to be
driven from the outside.
*/
package puredkg

import (
	"crypto/rand"
	"math/big"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

type (
	KeyperIndex = uint64
	phase       int
)

const (
	off = phase(iota)
	dealing
	accusing
	apologizing
	finalized
)

// XXX All of the messages here carry the Eon field, which we could also remove. It's not needed
// here.

// PolyCommitmentMsg is broadcast to all keypers
type PolyCommitmentMsg struct {
	Eon    uint64
	Sender KeyperIndex
	Gammas *crypto.Gammas
}

// PolyEvalMsg is sent over a secure channel to a single receiver
type PolyEvalMsg struct {
	Eon      uint64
	Sender   KeyperIndex
	Receiver KeyperIndex
	Eval     *big.Int
}

// AccusationMsg is broadcast, Accuser is the sender
type AccusationMsg struct {
	Eon              uint64
	Accuser, Accused KeyperIndex
}

// ApologyMsg is broadcast, Accused is the sender
type ApologyMsg struct {
	Eon              uint64
	Accuser, Accused KeyperIndex
	Eval             *big.Int
}

type accusationKey struct {
	Accuser, Accused KeyperIndex
}

// PureDKG implements the distributed key generation process for a single keyper
type PureDKG struct {
	Phase       phase
	Eon         uint64
	NumKeypers  uint64
	Threshold   uint64
	Keyper      KeyperIndex
	Polynomial  *crypto.Polynomial
	Commitment  []*crypto.Gammas
	Evals       []*big.Int
	Accusations map[accusationKey]bool
	Apologies   map[accusationKey]*big.Int
}

func NewPureDKG(eon uint64, numKeypers uint64, threshold uint64, keyper KeyperIndex) PureDKG {
	return PureDKG{
		Phase:       off,
		Eon:         eon,
		NumKeypers:  numKeypers,
		Threshold:   threshold,
		Keyper:      keyper,
		Commitment:  make([]*crypto.Gammas, numKeypers),
		Evals:       make([]*big.Int, numKeypers),
		Accusations: make(map[accusationKey]bool),
		Apologies:   make(map[accusationKey]*big.Int),
	}
}

func (pure *PureDKG) setPhase(p phase) {
	if p != pure.Phase+1 {
		panic("wrong phase")
	}
}

func (pure *PureDKG) StartPhase1Dealing() (PolyCommitmentMsg, []PolyEvalMsg, error) {
	polynomial, err := crypto.RandomPolynomial(rand.Reader, pure.Threshold)
	if err != nil {
		return PolyCommitmentMsg{}, []PolyEvalMsg{}, err
	}
	pure.setPhase(dealing)
	pure.Polynomial = polynomial

	polyCommitmentMsg := PolyCommitmentMsg{
		Eon:    pure.Eon,
		Sender: pure.Keyper,
		Gammas: pure.Polynomial.Gammas(),
	}

	var polyEvalMsgs []PolyEvalMsg
	var receiver KeyperIndex
	for receiver = 0; receiver < pure.NumKeypers; receiver++ {
		polyEvalMsgs = append(polyEvalMsgs, PolyEvalMsg{
			Eon:      pure.Eon,
			Sender:   pure.Keyper,
			Receiver: receiver,
			Eval:     pure.Polynomial.EvalForKeyper(int(receiver)),
		})
	}
	return polyCommitmentMsg, polyEvalMsgs, nil
}

func (pure *PureDKG) StartPhase2Accusing() []AccusationMsg {
	pure.setPhase(accusing)
	var accusations []AccusationMsg
	var dealer KeyperIndex
	for dealer = 0; dealer < pure.NumKeypers; dealer++ {
		if !pure.verifyPolyEval(dealer) {
			accusations = append(accusations, AccusationMsg{
				Eon:     pure.Eon,
				Accuser: pure.Keyper,
				Accused: dealer,
			})
		}
	}
	return accusations
}

func (pure *PureDKG) StartPhase3Apologizing() []ApologyMsg {
	pure.setPhase(apologizing)
	var apologies []ApologyMsg
	return apologies
}

func (pure *PureDKG) Finalize() {
	pure.setPhase(finalized)
}

func (pure *PureDKG) verifyPolyEval(dealer KeyperIndex) bool {
	if pure.Evals[dealer] == nil || pure.Commitment[dealer] == nil {
		return false
	}
	return crypto.VerifyPolyEval(int(dealer), pure.Evals[dealer], pure.Commitment[dealer], pure.Threshold)
}

// HandlePolyCommitmentMsg
func (pure *PureDKG) HandlePolyCommitmentMsg(msg PolyCommitmentMsg) error {
	// XXX implement error handling
	pure.Commitment[msg.Sender] = msg.Gammas
	return nil
}

// HandlePolyEvalMsg
func (pure *PureDKG) HandlePolyEvalMsg(msg PolyEvalMsg) error {
	// XXX implement error handling
	pure.Evals[msg.Sender] = msg.Eval
	return nil
}

// HandleAccusationMsg
func (pure *PureDKG) HandleAccusationMsg(msg AccusationMsg) error {
	pure.Accusations[accusationKey{Accuser: msg.Accuser, Accused: msg.Accused}] = true
	return nil
}

// HandleApologyMsg
func (pure *PureDKG) HandleApologyMsg(msg ApologyMsg) error {
	pure.Apologies[accusationKey{Accuser: msg.Accuser, Accused: msg.Accused}] = msg.Eval
	return nil
}
