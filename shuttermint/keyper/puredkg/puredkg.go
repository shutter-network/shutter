/*puredkg implements the DKG protocol. It's independent of any transport, but rather expects to be
driven from the outside.
*/
package puredkg

import (
	"crypto/rand"
	"fmt"
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
	Commitments []*crypto.Gammas
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
		Commitments: make([]*crypto.Gammas, numKeypers),
		Evals:       make([]*big.Int, numKeypers),
		Accusations: make(map[accusationKey]bool),
		Apologies:   make(map[accusationKey]*big.Int),
	}
}

func (pure *PureDKG) setPhase(p phase) {
	if p != pure.Phase+1 {
		panic("wrong phase")
	}
	pure.Phase = p
}

func (pure *PureDKG) checkPhase(maxPhase phase) error {
	if pure.Phase > maxPhase {
		return fmt.Errorf("received msg for phase %d in phase %d", maxPhase, pure.Phase)
	}
	return nil
}

func (pure *PureDKG) checkEon(eon uint64) error {
	if pure.Eon != eon {
		return fmt.Errorf("received msg for eon %d instead of %d", eon, pure.Eon)
	}
	return nil
}

func (pure *PureDKG) checkEonAndPhase(eon uint64, maxPhase phase) error {
	if err := pure.checkEon(eon); err != nil {
		return err
	}
	if err := pure.checkPhase(maxPhase); err != nil {
		return err
	}
	return nil
}

func (pure *PureDKG) StartPhase1Dealing() (PolyCommitmentMsg, []PolyEvalMsg, error) {
	pure.setPhase(dealing)
	degree := crypto.DegreeFromThreshold(pure.Threshold)
	polynomial, err := crypto.RandomPolynomial(rand.Reader, degree)
	if err != nil {
		return PolyCommitmentMsg{}, []PolyEvalMsg{}, err
	}
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

	// XXX: should we self-send or rely on the owner to call HandleXXX?
}

func (pure *PureDKG) StartPhase2Accusing() []AccusationMsg {
	pure.setPhase(accusing)
	var accusations []AccusationMsg
	var dealer KeyperIndex
	for dealer = 0; dealer < pure.NumKeypers; dealer++ {
		if pure.present(dealer) && !pure.verifyPolyEval(dealer) {
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

func (pure *PureDKG) present(i KeyperIndex) bool {
	return pure.Commitments[i] != nil
}

func (pure *PureDKG) verifyPolyEval(dealer KeyperIndex) bool {
	if pure.Evals[dealer] == nil || pure.Commitments[dealer] == nil {
		return false
	}
	return crypto.VerifyPolyEval(int(pure.Keyper), pure.Evals[dealer], pure.Commitments[dealer], pure.Threshold)
}

// HandlePolyCommitmentMsg
func (pure *PureDKG) HandlePolyCommitmentMsg(msg PolyCommitmentMsg) error {
	if err := pure.checkEonAndPhase(msg.Eon, dealing); err != nil {
		return err
	}
	if pure.Commitments[msg.Sender] != nil {
		return fmt.Errorf("received duplicate poly commitment msg")
	}
	if msg.Gammas.Degree() != crypto.DegreeFromThreshold(pure.Threshold) {
		return fmt.Errorf(
			"received poly commitment with unexpected degree %d instead of %d",
			msg.Gammas.Degree(),
			crypto.DegreeFromThreshold(pure.Threshold),
		)
	}

	pure.Commitments[msg.Sender] = msg.Gammas
	return nil
}

// HandlePolyEvalMsg
func (pure *PureDKG) HandlePolyEvalMsg(msg PolyEvalMsg) error {
	if err := pure.checkEonAndPhase(msg.Eon, dealing); err != nil {
		return err
	}
	if msg.Receiver != pure.Keyper {
		return fmt.Errorf("received poly eval msg for keyper %d instead of %d", msg.Receiver, pure.Keyper)
	}
	if pure.Evals[msg.Sender] != nil {
		return fmt.Errorf("received duplicate poly eval msg")
	}
	if !crypto.ValidEval(msg.Eval) {
		return fmt.Errorf("received invalid poly eval %d", msg.Eval)
	}

	pure.Evals[msg.Sender] = msg.Eval
	return nil
}

// HandleAccusationMsg
func (pure *PureDKG) HandleAccusationMsg(msg AccusationMsg) error {
	if err := pure.checkEonAndPhase(msg.Eon, accusing); err != nil {
		return err
	}
	key := accusationKey{Accuser: msg.Accuser, Accused: msg.Accused}
	if _, ok := pure.Accusations[key]; ok {
		return fmt.Errorf("received duplicate accusation")
	}

	pure.Accusations[key] = true
	return nil
}

// HandleApologyMsg
func (pure *PureDKG) HandleApologyMsg(msg ApologyMsg) error {
	if err := pure.checkEonAndPhase(msg.Eon, apologizing); err != nil {
		return err
	}
	key := accusationKey{Accuser: msg.Accuser, Accused: msg.Accused}
	if _, ok := pure.Apologies[key]; ok {
		return fmt.Errorf("received duplicate apology")
	}
	if !crypto.ValidEval(msg.Eval) {
		return fmt.Errorf("received apology with invalid poly eval %d", msg.Eval)
	}

	pure.Apologies[key] = msg.Eval
	return nil
}
