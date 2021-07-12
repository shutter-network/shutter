package shcrypto

import (
	"crypto/rand"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	gocmp "github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shuttermint/internal/shtest"
)

func TestEonSecretKeyShare(t *testing.T) {
	zeroKey := ComputeEonSecretKeyShare([]*big.Int{})
	assert.DeepEqual(t, big.NewInt(0), (*big.Int)(zeroKey), shtest.BigIntComparer)

	key1 := ComputeEonSecretKeyShare([]*big.Int{
		big.NewInt(10),
		big.NewInt(20),
		big.NewInt(30),
	})
	assert.DeepEqual(t, big.NewInt(60), (*big.Int)(key1), shtest.BigIntComparer)

	key2 := ComputeEonSecretKeyShare([]*big.Int{
		bn256.Order,
		big.NewInt(10),
		bn256.Order,
		big.NewInt(20),
		bn256.Order,
	})
	assert.DeepEqual(t, big.NewInt(30), (*big.Int)(key2), shtest.BigIntComparer)
}

func TestEonPublicKeyShare(t *testing.T) {
	gammas0 := Gammas{
		new(bn256.G2).ScalarBaseMult(big.NewInt(1)),
		new(bn256.G2).ScalarBaseMult(big.NewInt(2)),
	}
	gammas1 := Gammas{
		new(bn256.G2).ScalarBaseMult(big.NewInt(3)),
		new(bn256.G2).ScalarBaseMult(big.NewInt(4)),
	}
	gammas2 := Gammas{
		new(bn256.G2).ScalarBaseMult(big.NewInt(5)),
		new(bn256.G2).ScalarBaseMult(big.NewInt(6)),
	}
	gammas := []*Gammas{
		&gammas0,
		&gammas1,
		&gammas2,
	}

	x0 := KeyperX(0)
	x1 := KeyperX(1)
	x2 := KeyperX(2)

	mu00 := new(bn256.G2).Add(gammas0[0], new(bn256.G2).ScalarMult(gammas0[1], x0))
	mu01 := new(bn256.G2).Add(gammas1[0], new(bn256.G2).ScalarMult(gammas1[1], x0))
	mu02 := new(bn256.G2).Add(gammas2[0], new(bn256.G2).ScalarMult(gammas2[1], x0))
	mu10 := new(bn256.G2).Add(gammas0[0], new(bn256.G2).ScalarMult(gammas0[1], x1))
	mu11 := new(bn256.G2).Add(gammas1[0], new(bn256.G2).ScalarMult(gammas1[1], x1))
	mu12 := new(bn256.G2).Add(gammas2[0], new(bn256.G2).ScalarMult(gammas2[1], x1))
	mu20 := new(bn256.G2).Add(gammas0[0], new(bn256.G2).ScalarMult(gammas0[1], x2))
	mu21 := new(bn256.G2).Add(gammas1[0], new(bn256.G2).ScalarMult(gammas1[1], x2))
	mu22 := new(bn256.G2).Add(gammas2[0], new(bn256.G2).ScalarMult(gammas2[1], x2))

	pks0 := new(bn256.G2).Add(mu00, mu01)
	pks0.Add(pks0, mu02)
	pks1 := new(bn256.G2).Add(mu10, mu11)
	pks1.Add(pks1, mu12)
	pks2 := new(bn256.G2).Add(mu20, mu21)
	pks2.Add(pks2, mu22)

	assert.DeepEqual(t, pks0, (*bn256.G2)(ComputeEonPublicKeyShare(0, gammas)), G2Comparer)
	assert.DeepEqual(t, pks1, (*bn256.G2)(ComputeEonPublicKeyShare(1, gammas)), G2Comparer)
	assert.DeepEqual(t, pks2, (*bn256.G2)(ComputeEonPublicKeyShare(2, gammas)), G2Comparer)
}

func TestEonSharesMatch(t *testing.T) {
	threshold := uint64(2)
	p1, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)

	x1 := KeyperX(0)
	x2 := KeyperX(1)
	x3 := KeyperX(2)

	gammas := []*Gammas{p1.Gammas(), p2.Gammas(), p3.Gammas()}

	v11 := p1.Eval(x1)
	v21 := p1.Eval(x2)
	v31 := p1.Eval(x3)
	v12 := p2.Eval(x1)
	v22 := p2.Eval(x2)
	v32 := p2.Eval(x3)
	v13 := p3.Eval(x1)
	v23 := p3.Eval(x2)
	v33 := p3.Eval(x3)

	esk1 := ComputeEonSecretKeyShare([]*big.Int{v11, v12, v13})
	esk2 := ComputeEonSecretKeyShare([]*big.Int{v21, v22, v23})
	esk3 := ComputeEonSecretKeyShare([]*big.Int{v31, v32, v33})

	epk1 := ComputeEonPublicKeyShare(0, gammas)
	epk2 := ComputeEonPublicKeyShare(1, gammas)
	epk3 := ComputeEonPublicKeyShare(2, gammas)

	epk1Exp := (*EonPublicKeyShare)(new(bn256.G2).ScalarBaseMult((*big.Int)(esk1)))
	epk2Exp := (*EonPublicKeyShare)(new(bn256.G2).ScalarBaseMult((*big.Int)(esk2)))
	epk3Exp := (*EonPublicKeyShare)(new(bn256.G2).ScalarBaseMult((*big.Int)(esk3)))

	assert.DeepEqual(t, epk1, epk1Exp)
	assert.DeepEqual(t, epk2, epk2Exp)
	assert.DeepEqual(t, epk3, epk3Exp)
}

func TestEonPublicKey(t *testing.T) {
	zeroEPK := ComputeEonPublicKey([]*Gammas{})
	assert.DeepEqual(t, (*bn256.G2)(zeroEPK), zeroG2, G2Comparer)

	threshold := uint64(2)
	p1, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)

	k1 := ComputeEonPublicKey([]*Gammas{p1.Gammas()})
	assert.DeepEqual(t, (*bn256.G2)(k1), []*bn256.G2(*p1.Gammas())[0], G2Comparer)
	k2 := ComputeEonPublicKey([]*Gammas{p2.Gammas()})
	assert.DeepEqual(t, (*bn256.G2)(k2), []*bn256.G2(*p2.Gammas())[0], G2Comparer)
	k3 := ComputeEonPublicKey([]*Gammas{p3.Gammas()})
	assert.DeepEqual(t, (*bn256.G2)(k3), []*bn256.G2(*p3.Gammas())[0], G2Comparer)
}

func TestEonPublicKeyMatchesSecretKey(t *testing.T) {
	threshold := uint64(2)
	p1, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)

	esk := big.NewInt(0)
	for _, p := range []*Polynomial{p1, p2, p3} {
		esk = esk.Add(esk, p.Eval(big.NewInt(0)))
	}
	epkExp := new(bn256.G2).ScalarBaseMult(esk)

	gammas := []*Gammas{p1.Gammas(), p2.Gammas(), p3.Gammas()}
	epk := ComputeEonPublicKey(gammas)
	assert.DeepEqual(t, (*bn256.G2)(epk), epkExp, G2Comparer)
}

var modbn256Comparer = gocmp.Comparer(func(x, y *big.Int) bool {
	d := new(big.Int).Sub(x, y)
	return d.Mod(d, bn256.Order).Sign() == 0
})

func TestInverse(t *testing.T) {
	testCases := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		new(big.Int).Sub(bn256.Order, big.NewInt(2)),
		new(big.Int).Sub(bn256.Order, big.NewInt(1)),
	}
	for i := 0; i < 100; i++ {
		x, err := rand.Int(rand.Reader, bn256.Order)
		assert.NilError(t, err)
		if x.Sign() == 0 {
			continue
		}
		testCases = append(testCases, x)
	}

	for _, test := range testCases {
		inv := invert(test)
		one := new(big.Int).Mul(test, inv)
		assert.DeepEqual(t, big.NewInt(1), one, modbn256Comparer)
	}
}

func TestLagrangeCoefficientFactors(t *testing.T) {
	l01 := lagrangeCoefficientFactor(0, 1)
	l02 := lagrangeCoefficientFactor(0, 2)
	l10 := lagrangeCoefficientFactor(1, 0)
	l12 := lagrangeCoefficientFactor(1, 2)
	l20 := lagrangeCoefficientFactor(2, 0)
	l21 := lagrangeCoefficientFactor(2, 1)

	qMinus1 := new(big.Int).Sub(bn256.Order, big.NewInt(1))
	qMinus2 := new(big.Int).Sub(bn256.Order, big.NewInt(2))

	l01.Mul(l01, qMinus1)
	assert.DeepEqual(t, big.NewInt(1), l01, modbn256Comparer)
	l02.Mul(l02, qMinus2)
	assert.DeepEqual(t, big.NewInt(1), l02, modbn256Comparer)

	assert.DeepEqual(t, big.NewInt(2), l10, modbn256Comparer)
	l12.Mul(l12, qMinus1)
	assert.DeepEqual(t, big.NewInt(2), l12, modbn256Comparer)

	l20.Mul(l20, big.NewInt(2))
	assert.DeepEqual(t, big.NewInt(3), l20, modbn256Comparer)
	assert.DeepEqual(t, big.NewInt(3), l21, modbn256Comparer)
}

func TestLagrangeCoefficients(t *testing.T) {
	assert.DeepEqual(t, big.NewInt(1), lagrangeCoefficient(0, []int{0}), shtest.BigIntComparer)
	assert.DeepEqual(t, big.NewInt(1), lagrangeCoefficient(1, []int{1}), shtest.BigIntComparer)
	assert.DeepEqual(t, big.NewInt(1), lagrangeCoefficient(2, []int{2}), shtest.BigIntComparer)

	assert.DeepEqual(t, lagrangeCoefficientFactor(1, 0), lagrangeCoefficient(0, []int{0, 1}), shtest.BigIntComparer)
	assert.DeepEqual(t, lagrangeCoefficientFactor(0, 1), lagrangeCoefficient(1, []int{0, 1}), shtest.BigIntComparer)

	l0 := lagrangeCoefficient(0, []int{0, 1, 2})
	l0Exp := lagrangeCoefficientFactor(1, 0)
	l0Exp.Mul(l0Exp, lagrangeCoefficientFactor(2, 0))
	assert.DeepEqual(t, l0Exp, l0, modbn256Comparer)

	l1 := lagrangeCoefficient(1, []int{0, 1, 2})
	l1Exp := lagrangeCoefficientFactor(0, 1)
	l1Exp.Mul(l1Exp, lagrangeCoefficientFactor(2, 1))
	assert.DeepEqual(t, l1Exp, l1, modbn256Comparer)

	l2 := lagrangeCoefficient(2, []int{0, 1, 2})
	l2Exp := lagrangeCoefficientFactor(0, 2)
	l2Exp.Mul(l2Exp, lagrangeCoefficientFactor(1, 2))
	assert.DeepEqual(t, l2Exp, l2, modbn256Comparer)
}

func TestLagrangeReconstruct(t *testing.T) {
	p, err := RandomPolynomial(rand.Reader, uint64(2))
	assert.NilError(t, err)

	l1 := lagrangeCoefficient(0, []int{0, 1, 2})
	l2 := lagrangeCoefficient(1, []int{0, 1, 2})
	l3 := lagrangeCoefficient(2, []int{0, 1, 2})
	v1 := p.EvalForKeyper(0)
	v2 := p.EvalForKeyper(1)
	v3 := p.EvalForKeyper(2)

	y1 := new(big.Int).Mul(l1, v1)
	y2 := new(big.Int).Mul(l2, v2)
	y3 := new(big.Int).Mul(l3, v3)
	y1.Mod(y1, bn256.Order)
	y2.Mod(y2, bn256.Order)
	y3.Mod(y3, bn256.Order)

	y := new(big.Int).Add(y1, y2)
	y.Add(y, y3)
	y.Mod(y, bn256.Order)

	assert.DeepEqual(t, p.Eval(big.NewInt(0)), y, shtest.BigIntComparer)
}

func TestComputeEpochSecretKeyShare(t *testing.T) {
	eonSecretKeyShare := (*EonSecretKeyShare)(big.NewInt(123))
	epochID := ComputeEpochID(uint64(456))
	epochSecretKeyShare := ComputeEpochSecretKeyShare(eonSecretKeyShare, epochID)
	expectedEpochSecretKeyShare := new(bn256.G1).ScalarMult((*bn256.G1)(epochID), (*big.Int)(eonSecretKeyShare))
	assert.DeepEqual(t, expectedEpochSecretKeyShare, (*bn256.G1)(epochSecretKeyShare), G1Comparer)
}

func TestVerifyEpochSecretKeyShare(t *testing.T) {
	threshold := uint64(2)
	epochID := ComputeEpochID(uint64(10))
	p1, err := RandomPolynomial(rand.Reader, threshold-1)
	assert.NilError(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold-1)
	assert.NilError(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold-1)
	assert.NilError(t, err)

	gammas := []*Gammas{
		p1.Gammas(),
		p2.Gammas(),
		p3.Gammas(),
	}

	epk1 := ComputeEonPublicKeyShare(0, gammas)
	epk2 := ComputeEonPublicKeyShare(1, gammas)
	epk3 := ComputeEonPublicKeyShare(2, gammas)
	esk1 := ComputeEonSecretKeyShare([]*big.Int{p1.EvalForKeyper(0), p2.EvalForKeyper(0), p3.EvalForKeyper(0)})
	esk2 := ComputeEonSecretKeyShare([]*big.Int{p1.EvalForKeyper(1), p2.EvalForKeyper(1), p3.EvalForKeyper(1)})
	esk3 := ComputeEonSecretKeyShare([]*big.Int{p1.EvalForKeyper(2), p2.EvalForKeyper(2), p3.EvalForKeyper(2)})
	epsk1 := ComputeEpochSecretKeyShare(esk1, epochID)
	epsk2 := ComputeEpochSecretKeyShare(esk2, epochID)
	epsk3 := ComputeEpochSecretKeyShare(esk3, epochID)

	assert.Assert(t, VerifyEpochSecretKeyShare(epsk1, epk1, epochID))
	assert.Assert(t, VerifyEpochSecretKeyShare(epsk2, epk2, epochID))
	assert.Assert(t, VerifyEpochSecretKeyShare(epsk3, epk3, epochID))

	assert.Assert(t, !VerifyEpochSecretKeyShare(epsk1, epk2, epochID))
	assert.Assert(t, !VerifyEpochSecretKeyShare(epsk2, epk1, epochID))
	assert.Assert(t, !VerifyEpochSecretKeyShare(epsk1, epk1, ComputeEpochID(uint64(11))))
}

func TestComputeEpochSecretKey(t *testing.T) {
	n := 3
	threshold := uint64(2)
	epochID := ComputeEpochID(uint64(10))

	ps := []*Polynomial{}
	for i := 0; i < n; i++ {
		p, err := RandomPolynomial(rand.Reader, threshold-1)
		assert.NilError(t, err)
		ps = append(ps, p)
	}

	epochSecretKeyShares := []*EpochSecretKeyShare{}
	for i := 0; i < n; i++ {
		vs := []*big.Int{}
		for _, p := range ps {
			v := p.EvalForKeyper(i)
			vs = append(vs, v)
		}
		eonSecretKeyShare := ComputeEonSecretKeyShare(vs)
		epochSecretKeyShare := ComputeEpochSecretKeyShare(eonSecretKeyShare, epochID)

		epochSecretKeyShares = append(epochSecretKeyShares, epochSecretKeyShare)
	}

	var err error
	_, err = ComputeEpochSecretKey([]int{0}, epochSecretKeyShares[:1], threshold)
	assert.Assert(t, err != nil)
	_, err = ComputeEpochSecretKey([]int{0, 1, 2}, epochSecretKeyShares[:2], threshold)
	assert.Assert(t, err != nil)
	_, err = ComputeEpochSecretKey([]int{0, 1}, epochSecretKeyShares[:1], threshold)
	assert.Assert(t, err != nil)
	_, err = ComputeEpochSecretKey([]int{0}, epochSecretKeyShares[:2], threshold)
	assert.Assert(t, err != nil)

	epochSecretKey12, err := ComputeEpochSecretKey(
		[]int{0, 1},
		[]*EpochSecretKeyShare{epochSecretKeyShares[0], epochSecretKeyShares[1]},
		threshold)
	assert.NilError(t, err)
	epochSecretKey13, err := ComputeEpochSecretKey(
		[]int{0, 2},
		[]*EpochSecretKeyShare{epochSecretKeyShares[0], epochSecretKeyShares[2]},
		threshold)
	assert.NilError(t, err)
	epochSecretKey23, err := ComputeEpochSecretKey(
		[]int{1, 2},
		[]*EpochSecretKeyShare{epochSecretKeyShares[1], epochSecretKeyShares[2]},
		threshold)
	assert.NilError(t, err)

	assert.DeepEqual(t, epochSecretKey12, epochSecretKey13)
	assert.DeepEqual(t, epochSecretKey12, epochSecretKey23)
}

func TestFull(t *testing.T) {
	n := 3
	threshold := uint64(2)
	epochID := ComputeEpochID(uint64(10))

	ps := []*Polynomial{}
	gammas := []*Gammas{}
	for i := 0; i < n; i++ {
		p, err := RandomPolynomial(rand.Reader, threshold-1)
		assert.NilError(t, err)
		ps = append(ps, p)
		gammas = append(gammas, p.Gammas())
	}

	eonSecretKeyShares := []*EonSecretKeyShare{}
	for i := 0; i < n; i++ {
		vs := []*big.Int{}
		for j := 0; j < n; j++ {
			v := ps[j].EvalForKeyper(i)
			vs = append(vs, v)
		}
		eonSecretKeyShare := ComputeEonSecretKeyShare(vs)
		eonSecretKeyShares = append(eonSecretKeyShares, eonSecretKeyShare)
	}

	eonPublicKeyShares := []*EonPublicKeyShare{}
	for i := 0; i < n; i++ {
		eonPublicKeyShare := ComputeEonPublicKeyShare(i, gammas)
		eonPublicKeyShares = append(eonPublicKeyShares, eonPublicKeyShare)
	}

	epochSecretKeyShares := []*EpochSecretKeyShare{}
	for i := 0; i < n; i++ {
		epochSecretKeyShare := ComputeEpochSecretKeyShare(eonSecretKeyShares[i], epochID)
		epochSecretKeyShares = append(epochSecretKeyShares, epochSecretKeyShare)
	}

	// verify (published) epoch sk shares
	for i := 0; i < n; i++ {
		assert.Assert(t, VerifyEpochSecretKeyShare(epochSecretKeyShares[i], eonPublicKeyShares[i], epochID))
	}

	epochSecretKey, err := ComputeEpochSecretKey(
		[]int{0, 1},
		[]*EpochSecretKeyShare{epochSecretKeyShares[0], epochSecretKeyShares[1]},
		threshold)
	assert.NilError(t, err)

	epochSecretKey13, err := ComputeEpochSecretKey(
		[]int{0, 2},
		[]*EpochSecretKeyShare{epochSecretKeyShares[0], epochSecretKeyShares[2]},
		threshold)
	assert.NilError(t, err)
	epochSecretKey23, err := ComputeEpochSecretKey(
		[]int{1, 2},
		[]*EpochSecretKeyShare{epochSecretKeyShares[1], epochSecretKeyShares[2]},
		threshold)
	assert.NilError(t, err)
	assert.DeepEqual(t, epochSecretKey, epochSecretKey13)
	assert.DeepEqual(t, epochSecretKey, epochSecretKey23)
}
