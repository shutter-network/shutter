package puredkg

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shcrypto"
	"github.com/shutter-network/shutter/shlib/shtest"
)

func TestPureDKGFull(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(3)
	threshold := uint64(2)

	dkgs := []*PureDKG{}
	for i := uint64(0); i < numKeypers; i++ {
		dkg := NewPureDKG(eon, numKeypers, threshold, i)
		dkgs = append(dkgs, &dkg)
	}

	// dealing phase
	for _, dkg := range dkgs {
		polyCommitmentMsg, polyEvalMsgs, err := dkg.StartPhase1Dealing()
		assert.NilError(t, err)

		for _, receiverDKG := range dkgs {
			err := receiverDKG.HandlePolyCommitmentMsg(polyCommitmentMsg)
			assert.NilError(t, err)
		}

		assert.Equal(t, int(numKeypers), 1+len(polyEvalMsgs))
		for _, msg := range polyEvalMsgs {
			err := dkgs[msg.Receiver].HandlePolyEvalMsg(msg)
			assert.NilError(t, err)
		}
	}

	// accusation phase
	for _, dkg := range dkgs {
		accusations := dkg.StartPhase2Accusing()
		assert.Assert(t, len(accusations) == 0)
	}

	// apology phase
	for _, dkg := range dkgs {
		apologies := dkg.StartPhase3Apologizing()
		assert.Assert(t, len(apologies) == 0)
	}

	// finalize
	for _, dkg := range dkgs {
		dkg.Finalize()
	}

	var results []Result

	for _, dkg := range dkgs {
		result, err := dkg.ComputeResult()
		assert.NilError(t, err)
		results = append(results, result)
	}
	for _, r := range results {
		assert.Assert(t, reflect.DeepEqual(r.PublicKey, results[0].PublicKey))
	}
}

// TestPureDKGOfflineSendAccusation tests that we send accusations when we don't receive any
// message from a keyper. See https://github.com/shutter-network/shutter/issues/62
func TestPureDKGOfflineSendAccusation(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(3)
	threshold := uint64(2)

	dkg := NewPureDKG(eon, numKeypers, threshold, 0)

	// dealing phase
	_, _, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	accusations := dkg.StartPhase2Accusing()
	assert.Equal(t, 2, len(accusations))
}

func TestPureDKGCorrupt(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(3)
	threshold := uint64(2)

	dkgs := []*PureDKG{}
	for i := uint64(0); i < numKeypers; i++ {
		dkg := NewPureDKG(eon, numKeypers, threshold, i)
		dkgs = append(dkgs, &dkg)
	}
	honestDKGs := dkgs[:len(dkgs)-1]
	corruptDKG := dkgs[len(dkgs)-1]

	// dealing phase
	for _, dkg := range honestDKGs {
		polyCommitmentMsg, polyEvalMsgs, err := dkg.StartPhase1Dealing()
		assert.NilError(t, err)

		for _, receiverDKG := range dkgs {
			assert.NilError(t, receiverDKG.HandlePolyCommitmentMsg(polyCommitmentMsg))
		}
		for _, msg := range polyEvalMsgs {
			assert.NilError(t, dkgs[msg.Receiver].HandlePolyEvalMsg(msg))
		}
	}

	// corrupt DKG sends invalid poly eval to first keyper
	polyCommitmentMsg, polyEvalMsgs, err := corruptDKG.StartPhase1Dealing()
	assert.NilError(t, err)
	for _, receiverDKG := range dkgs {
		assert.NilError(t, receiverDKG.HandlePolyCommitmentMsg(polyCommitmentMsg))
	}
	for _, msg := range polyEvalMsgs {
		if msg.Receiver == 0 {
			corruptMsg := PolyEvalMsg{
				Eon:      msg.Eon,
				Sender:   msg.Sender,
				Receiver: msg.Receiver,
				Eval:     big.NewInt(666),
			}
			assert.NilError(t, dkgs[msg.Receiver].HandlePolyEvalMsg(corruptMsg))
		} else {
			assert.NilError(t, dkgs[msg.Receiver].HandlePolyEvalMsg(msg))
		}
	}

	// accusation phase
	for _, dkg := range dkgs[1:] {
		accusations := dkg.StartPhase2Accusing()
		assert.Assert(t, len(accusations) == 0)
	}
	accusations := dkgs[0].StartPhase2Accusing()
	assert.Equal(t, 1, len(accusations))
	for _, dkg := range dkgs {
		err := dkg.HandleAccusationMsg(accusations[0])
		assert.NilError(t, err)
	}

	// apology phase
	for _, dkg := range honestDKGs {
		apologies := dkg.StartPhase3Apologizing()
		assert.Assert(t, len(apologies) == 0)
	}
	apologies := corruptDKG.StartPhase3Apologizing()
	assert.Equal(t, 1, len(apologies))
	for _, dkg := range dkgs {
		msg := ApologyMsg{
			Eon:     apologies[0].Eon,
			Accuser: apologies[0].Accuser,
			Accused: apologies[0].Accused,
			Eval:    big.NewInt(121212),
		}
		err := dkg.HandleApologyMsg(msg)
		assert.NilError(t, err)
	}

	// finalize
	for _, dkg := range dkgs {
		dkg.Finalize()
	}

	var results []Result
	for _, dkg := range dkgs {
		result, err := dkg.ComputeResult()
		assert.NilError(t, err)
		results = append(results, result)
	}
	for _, r := range results {
		assert.Assert(t, reflect.DeepEqual(r.PublicKey, results[0].PublicKey))
	}
}

func TestDealingSendsCorrectMsgs(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(3)
	threshold := uint64(2)
	keyper := uint64(1)
	dkg := NewPureDKG(eon, numKeypers, threshold, keyper)
	commitmentMsg, evalMsgs, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	assert.Equal(t, eon, commitmentMsg.Eon)
	assert.Equal(t, keyper, commitmentMsg.Sender)
	expectedGammas := dkg.Polynomial.Gammas()
	assert.DeepEqual(t, expectedGammas, commitmentMsg.Gammas)

	assert.Equal(t, int(numKeypers), 1+len(evalMsgs))
	for _, msg := range evalMsgs {
		assert.Equal(t, eon, msg.Eon)
		assert.Equal(t, keyper, msg.Sender)
		assert.DeepEqual(t, dkg.Polynomial.EvalForKeyper(int(msg.Receiver)), msg.Eval, shtest.BigIntComparer)
	}
}

func TestAccusing(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(4)
	threshold := uint64(3)
	keyper := uint64(3)
	dkg := NewPureDKG(eon, numKeypers, threshold, keyper)
	_, _, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	polys := []*shcrypto.Polynomial{}
	for i := 0; i < int(numKeypers); i++ {
		p, err := shcrypto.RandomPolynomial(rand.Reader, shcrypto.DegreeFromThreshold(threshold))
		assert.NilError(t, err)
		polys = append(polys, p)
	}

	makeCommitmentMsg := func(sender int) PolyCommitmentMsg {
		return PolyCommitmentMsg{
			Eon:    eon,
			Sender: uint64(sender),
			Gammas: polys[sender].Gammas(),
		}
	}
	makeEvalMsg := func(sender int) PolyEvalMsg {
		return PolyEvalMsg{
			Eon:      eon,
			Sender:   uint64(sender),
			Receiver: keyper,
			Eval:     polys[sender].EvalForKeyper(int(keyper)),
		}
	}

	// first keyper is honest
	commitmentMsg0 := makeCommitmentMsg(0)
	evalMsg0 := makeEvalMsg(0)
	err = dkg.HandlePolyCommitmentMsg(commitmentMsg0)
	assert.NilError(t, err)
	err = dkg.HandlePolyEvalMsg(evalMsg0)
	assert.NilError(t, err)

	// second keyper is dishonest
	commitmentMsg1 := makeCommitmentMsg(1)
	evalMsg1 := makeEvalMsg(1)
	evalMsg1.Eval = big.NewInt(123)
	err = dkg.HandlePolyCommitmentMsg(commitmentMsg1)
	assert.NilError(t, err)
	err = dkg.HandlePolyEvalMsg(evalMsg1)
	assert.NilError(t, err)

	// third keyper misses eval
	commitmentMsg2 := makeCommitmentMsg(2)
	err = dkg.HandlePolyCommitmentMsg(commitmentMsg2)
	assert.NilError(t, err)
	accusations := dkg.StartPhase2Accusing()
	assert.Equal(t, 2, len(accusations))
	accDishonest := accusations[0]
	accMissed := accusations[1]

	assert.Equal(t, eon, accDishonest.Eon)
	assert.Equal(t, keyper, accDishonest.Accuser)
	assert.Equal(t, KeyperIndex(1), accDishonest.Accused)

	assert.Equal(t, eon, accMissed.Eon)
	assert.Equal(t, keyper, accMissed.Accuser)
	assert.Equal(t, KeyperIndex(2), accMissed.Accused)
}

func TestApologizing(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(4)
	threshold := uint64(3)
	keyper := uint64(3)
	dkg := NewPureDKG(eon, numKeypers, threshold, keyper)
	_, _, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	_ = dkg.StartPhase2Accusing()

	accusation := AccusationMsg{
		Eon:     eon,
		Accuser: uint64(1),
		Accused: keyper,
	}
	err = dkg.HandleAccusationMsg(accusation)
	assert.NilError(t, err)

	apologies := dkg.StartPhase3Apologizing()
	assert.Equal(t, 1, len(apologies))
	assert.Equal(t, eon, apologies[0].Eon)
	assert.Equal(t, keyper, apologies[0].Accused)
	assert.Equal(t, accusation.Accuser, apologies[0].Accuser)
	assert.DeepEqual(t, dkg.Polynomial.EvalForKeyper(int(accusation.Accuser)), apologies[0].Eval, shtest.BigIntComparer)
}

func TestInvalidCommitmentHandling(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(4)
	threshold := uint64(3)
	keyper := uint64(3)
	dkg := NewPureDKG(eon, numKeypers, threshold, keyper)
	_, _, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	makeCommitmentMsg := func() PolyCommitmentMsg {
		p, err := shcrypto.RandomPolynomial(rand.Reader, shcrypto.DegreeFromThreshold(threshold))
		assert.NilError(t, err)
		return PolyCommitmentMsg{
			Eon:    eon,
			Sender: uint64(1),
			Gammas: p.Gammas(),
		}
	}

	c1 := makeCommitmentMsg()
	c1.Eon = uint64(10)
	assert.Assert(t, dkg.HandlePolyCommitmentMsg(c1) != nil)

	c2 := makeCommitmentMsg()
	gammas := (*c2.Gammas)[:len(*c2.Gammas)-1]
	c2.Gammas = &gammas
	assert.Assert(t, dkg.HandlePolyCommitmentMsg(c2) != nil)

	c3 := makeCommitmentMsg()
	assert.NilError(t, dkg.HandlePolyCommitmentMsg(c3))
	assert.Assert(t, dkg.HandlePolyCommitmentMsg(c3) != nil)
}

func TestInvalidEvalHandling(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(4)
	threshold := uint64(3)
	keyper := uint64(3)
	dkg := NewPureDKG(eon, numKeypers, threshold, keyper)
	_, _, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	makeEvalMsg := func() PolyEvalMsg {
		p, err := shcrypto.RandomPolynomial(rand.Reader, shcrypto.DegreeFromThreshold(threshold))
		assert.NilError(t, err)
		return PolyEvalMsg{
			Eon:      eon,
			Sender:   uint64(1),
			Receiver: keyper,
			Eval:     p.EvalForKeyper(int(keyper)),
		}
	}

	m1 := makeEvalMsg()
	m1.Eon = uint64(10)
	assert.Assert(t, dkg.HandlePolyEvalMsg(m1) != nil)

	m2 := makeEvalMsg()
	m2.Receiver = uint64(0)
	assert.Assert(t, dkg.HandlePolyEvalMsg(m2) != nil)

	m3 := makeEvalMsg()
	m3.Eval = big.NewInt(-1)
	assert.Assert(t, dkg.HandlePolyEvalMsg(m3) != nil)

	m4 := makeEvalMsg()
	assert.NilError(t, dkg.HandlePolyEvalMsg(m4))
	assert.Assert(t, dkg.HandlePolyEvalMsg(m4) != nil)

	dkg.StartPhase2Accusing()
	m5 := makeEvalMsg()
	m5.Sender = uint64(2)
	assert.Assert(t, dkg.HandlePolyEvalMsg(m5) != nil)
}

func TestInvalidAccusationHandling(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(4)
	threshold := uint64(3)
	keyper := uint64(3)
	dkg := NewPureDKG(eon, numKeypers, threshold, keyper)
	_, _, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	makeAccusationMsg := func() AccusationMsg {
		return AccusationMsg{
			Eon:     eon,
			Accuser: uint64(1),
			Accused: keyper,
		}
	}

	m1 := makeAccusationMsg()
	m1.Eon = uint64(10)
	assert.Assert(t, dkg.HandleAccusationMsg(m1) != nil)

	m2 := makeAccusationMsg()
	assert.NilError(t, dkg.HandleAccusationMsg(m2))
	assert.Assert(t, dkg.HandleAccusationMsg(m2) != nil)

	dkg.StartPhase2Accusing()
	dkg.StartPhase3Apologizing()

	m3 := makeAccusationMsg()
	m3.Accuser = uint64(0)
	assert.Assert(t, dkg.HandleAccusationMsg(m3) != nil)
}

func TestInvalidApologyHandling(t *testing.T) {
	eon := uint64(5)
	numKeypers := uint64(4)
	threshold := uint64(3)
	keyper := uint64(3)
	dkg := NewPureDKG(eon, numKeypers, threshold, keyper)
	_, _, err := dkg.StartPhase1Dealing()
	assert.NilError(t, err)

	makeApologyMsg := func() ApologyMsg {
		return ApologyMsg{
			Eon:     eon,
			Accuser: uint64(1),
			Accused: keyper,
			Eval:    big.NewInt(1),
		}
	}

	m1 := makeApologyMsg()
	m1.Eon = uint64(10)
	assert.Assert(t, dkg.HandleApologyMsg(m1) != nil)

	m2 := makeApologyMsg()
	m2.Eval = big.NewInt(-1)
	assert.Assert(t, dkg.HandleApologyMsg(m2) != nil)

	m3 := makeApologyMsg()
	assert.NilError(t, dkg.HandleApologyMsg(m3))
	assert.Assert(t, dkg.HandleApologyMsg(m3) != nil)

	dkg.StartPhase2Accusing()
	dkg.StartPhase3Apologizing()
	dkg.Finalize()

	m4 := makeApologyMsg()
	m4.Accuser = uint64(0)
	assert.Assert(t, dkg.HandleApologyMsg(m4) != nil)
}

func TestGetResultErrors(t *testing.T) {
	dkg := NewPureDKG(uint64(5), uint64(3), uint64(2), 1)
	_, err := dkg.ComputeResult()
	assert.Assert(t, err != nil)
	_, _, err = dkg.StartPhase1Dealing()
	assert.NilError(t, err)
	_, err = dkg.ComputeResult()
	assert.Assert(t, err != nil)
	dkg.StartPhase2Accusing()
	_, err = dkg.ComputeResult()
	assert.Assert(t, err != nil)
	dkg.StartPhase3Apologizing()
	_, err = dkg.ComputeResult()
	assert.Assert(t, err != nil)
	dkg.Finalize()
	_, err = dkg.ComputeResult()
	assert.Assert(t, err != nil)
}
