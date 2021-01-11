/*puredkg implements the DKG protocol. It's independent of any transport, but rather expects to be
driven from the outside.
*/

//go:generate stringer -type=Phase
package puredkg

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

type (
	KeyperIndex = uint64
	Phase       int
)

const (
	Off = Phase(iota)
	Dealing
	Accusing
	Apologizing
	Finalized
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
	Phase       Phase
	Eon         uint64
	NumKeypers  uint64
	Threshold   uint64
	Keyper      KeyperIndex
	Polynomial  *crypto.Polynomial
	Commitments []*crypto.Gammas
	Evals       []*big.Int
	Accusations map[accusationKey]struct{}
	Apologies   map[accusationKey]*big.Int
}

func NewPureDKG(eon uint64, numKeypers uint64, threshold uint64, keyper KeyperIndex) PureDKG {
	return PureDKG{
		Phase:       Off,
		Eon:         eon,
		NumKeypers:  numKeypers,
		Threshold:   threshold,
		Keyper:      keyper,
		Commitments: make([]*crypto.Gammas, numKeypers),
		Evals:       make([]*big.Int, numKeypers),
		Accusations: make(map[accusationKey]struct{}),
		Apologies:   make(map[accusationKey]*big.Int),
	}
}

func (pure *PureDKG) setPhase(p Phase) {
	if p != pure.Phase+1 {
		panic("wrong phase")
	}
	pure.Phase = p
}

func (pure *PureDKG) checkPhase(maxPhase Phase) error {
	if pure.Phase > maxPhase {
		return fmt.Errorf("received msg for phase '%s' in phase '%s'", maxPhase, pure.Phase)
	}
	return nil
}

func (pure *PureDKG) checkEon(eon uint64) error {
	if pure.Eon != eon {
		return fmt.Errorf("received msg for eon %d instead of %d", eon, pure.Eon)
	}
	return nil
}

func (pure *PureDKG) checkEonAndPhase(eon uint64, maxPhase Phase) error {
	if err := pure.checkEon(eon); err != nil {
		return err
	}
	if err := pure.checkPhase(maxPhase); err != nil {
		return err
	}
	return nil
}

func (pure *PureDKG) StartPhase1Dealing() (PolyCommitmentMsg, []PolyEvalMsg, error) {
	pure.setPhase(Dealing)
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
		msg := PolyEvalMsg{
			Eon:      pure.Eon,
			Sender:   pure.Keyper,
			Receiver: receiver,
			Eval:     pure.Polynomial.EvalForKeyper(int(receiver)),
		}
		if receiver == pure.Keyper {
			// XXX maybe use pure.Evals[msg.Sender] = msg.Eval instead
			err = pure.HandlePolyEvalMsg(msg)
			if err != nil {
				panic(err)
			}
		} else {
			polyEvalMsgs = append(polyEvalMsgs, msg)
		}
	}
	return polyCommitmentMsg, polyEvalMsgs, nil

	// XXX: should we self-send or rely on the owner to call HandleXXX?
}

func (pure *PureDKG) StartPhase2Accusing() []AccusationMsg {
	pure.setPhase(Accusing)
	var accusations []AccusationMsg
	var dealer KeyperIndex
	for dealer = 0; dealer < pure.NumKeypers; dealer++ {
		eval := pure.Evals[dealer]
		c := pure.Commitments[dealer]
		if pure.present(dealer) && (eval == nil || !crypto.VerifyPolyEval(int(pure.Keyper), eval, c, pure.Threshold)) {
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
	pure.setPhase(Apologizing)
	var apologies []ApologyMsg
	for key := range pure.Accusations {
		if key.Accused == pure.Keyper {
			apologies = append(apologies, ApologyMsg{
				Eon:     pure.Eon,
				Accuser: key.Accuser,
				Accused: key.Accused,
				Eval:    pure.Polynomial.EvalForKeyper(int(key.Accuser)),
			})
		}
	}
	return apologies
}

func (pure *PureDKG) Finalize() {
	pure.setPhase(Finalized)
}

func (pure *PureDKG) present(i KeyperIndex) bool {
	return pure.Commitments[i] != nil
}

// ShortInfo returns a short string to be used in log output, which describes the current state of the DKG
func (pure *PureDKG) ShortInfo() string {
	var numCommitments, numCorrupt, numAccusations, numApologies int
	numAccusations = len(pure.Accusations)
	numApologies = len(pure.Apologies)

	for dealer := uint64(0); dealer < pure.NumKeypers; dealer++ {
		c := pure.Commitments[dealer]
		eval := pure.polyEval(dealer)

		if c != nil && len(*c) != 0 {
			numCommitments += 1
		}

		if pure.isCorrupt(dealer) {
			numCorrupt += 1
		} else {
			if pure.Phase > Dealing && (eval == nil || !crypto.VerifyPolyEval(int(pure.Keyper), eval, c, pure.Threshold)) {
				numCorrupt += 1
			}
		}
	}

	return fmt.Sprintf("phase=%s(commitments=%d, corrupt=%d, accusations=%d, apologies=%d)", pure.Phase, numCommitments, numCorrupt, numAccusations, numApologies)
}

// ComputeResult computes the eon secret key share and public key output of the DKG process. An
// error is returned if this is called before finalization or if too few keypers participated.
func (pure *PureDKG) ComputeResult() (*crypto.EonSecretKeyShare, *crypto.EonPublicKey, error) {
	if pure.Phase < Finalized {
		return nil, nil, fmt.Errorf("dkg is not finalized yet")
	}

	numParticipants := 0
	commitments := []*crypto.Gammas{}
	evals := []*big.Int{}
	for dealer := uint64(0); dealer < pure.NumKeypers; dealer++ {
		var c *crypto.Gammas
		var eval *big.Int

		if !pure.isCorrupt(dealer) {
			numParticipants++
			c = pure.Commitments[dealer]
			eval = pure.polyEval(dealer)

			if eval == nil || !crypto.VerifyPolyEval(int(pure.Keyper), eval, c, pure.Threshold) {
				// when we receive no or an invalid poly eval, we send an accusation. If this
				// accusation does not end up in the chain, the keyper will not be considered
				// corrupt by the other keypers. We know  they should be though, so we abort.
				return nil, nil, fmt.Errorf("corrupt keyper %d not considered corrupt", dealer)
			}
		} else {
			c = crypto.ZeroGammas(crypto.DegreeFromThreshold(pure.Threshold))
			eval = big.NewInt(0)
		}
		commitments = append(commitments, c)
		evals = append(evals, eval)
	}

	if uint64(numParticipants) < pure.Threshold {
		return nil, nil, fmt.Errorf("only %d keypers participated, but threshold is %d", numParticipants, pure.Threshold)
	}
	eonSKShare := crypto.ComputeEonSecretKeyShare(evals)
	eonPK := crypto.ComputeEonPublicKey(commitments)
	return eonSKShare, eonPK, nil
}

// isCorrupt checks if the given keyper is considered corrupt. Note that this might change when
// new messages are received.
func (pure *PureDKG) isCorrupt(dealer KeyperIndex) bool {
	// a keyper is corrupt if they haven't sent a commitment
	c := pure.Commitments[dealer]
	if c == nil {
		return true
	}

	// a keyper is corrupt if they have apologized with a poly eval that doesn't match their
	// commitment
	for accusationKey, polyEval := range pure.Apologies {
		if accusationKey.Accused != dealer {
			continue
		}
		if !crypto.VerifyPolyEval(int(accusationKey.Accuser), polyEval, c, pure.Threshold) {
			return true
		}
	}

	// a keyper is corrupt if they haven't apologized in case of an accusation
	for accusationKey := range pure.Accusations {
		if accusationKey.Accused != dealer {
			continue
		}
		_, apologized := pure.Apologies[accusationKey]
		if !apologized {
			return true
		}
	}

	return false
}

// polyEval returns the poly eval received from the given dealer. Only call this function if the
// dealer has been determined to not be corrupt, otherwise the result is not necessarily unique.
func (pure *PureDKG) polyEval(dealer KeyperIndex) *big.Int {
	for accusationKey, polyEval := range pure.Apologies {
		if accusationKey.Accuser == pure.Keyper && accusationKey.Accused == dealer {
			return polyEval
		}
	}
	return pure.Evals[dealer]
}

// HandlePolyCommitmentMsg
func (pure *PureDKG) HandlePolyCommitmentMsg(msg PolyCommitmentMsg) error {
	if err := pure.checkEonAndPhase(msg.Eon, Dealing); err != nil {
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
	if err := pure.checkEonAndPhase(msg.Eon, Dealing); err != nil {
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
	if err := pure.checkEonAndPhase(msg.Eon, Accusing); err != nil {
		return err
	}
	key := accusationKey{Accuser: msg.Accuser, Accused: msg.Accused}
	if _, ok := pure.Accusations[key]; ok {
		return fmt.Errorf("received duplicate accusation")
	}

	pure.Accusations[key] = struct{}{}
	return nil
}

// HandleApologyMsg
func (pure *PureDKG) HandleApologyMsg(msg ApologyMsg) error {
	if err := pure.checkEonAndPhase(msg.Eon, Apologizing); err != nil {
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
