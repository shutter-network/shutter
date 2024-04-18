package shcrypto

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/bls12381"
	gocmp "github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shtest"
)

func makeTestG1(n int64) *bls12381.PointG1 {
	g1 := bls12381.NewG1()
	return g1.MulScalar(new(bls12381.PointG1), g1.One(), big.NewInt(n))
}

func makeTestG2(n int64) *bls12381.PointG2 {
	g2 := bls12381.NewG2()
	return g2.MulScalar(new(bls12381.PointG2), g2.One(), big.NewInt(n))
}

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
		order,
		big.NewInt(10),
		order,
		big.NewInt(20),
		order,
	})
	assert.DeepEqual(t, big.NewInt(30), (*big.Int)(key2), shtest.BigIntComparer)
}

func TestEonPublicKeyShare(t *testing.T) {
	g2 := bls12381.NewG2()
	gammas0 := Gammas{
		makeTestG2(1),
		makeTestG2(2),
	}
	gammas1 := Gammas{
		makeTestG2(3),
		makeTestG2(4),
	}
	gammas2 := Gammas{
		makeTestG2(5),
		makeTestG2(6),
	}
	gammas := []*Gammas{
		&gammas0,
		&gammas1,
		&gammas2,
	}

	x0 := KeyperX(0)
	x1 := KeyperX(1)
	x2 := KeyperX(2)

	mu00 := g2.Add(new(bls12381.PointG2), gammas0[0], g2.MulScalar(new(bls12381.PointG2), gammas0[1], x0))
	mu01 := g2.Add(new(bls12381.PointG2), gammas1[0], g2.MulScalar(new(bls12381.PointG2), gammas1[1], x0))
	mu02 := g2.Add(new(bls12381.PointG2), gammas2[0], g2.MulScalar(new(bls12381.PointG2), gammas2[1], x0))
	mu10 := g2.Add(new(bls12381.PointG2), gammas0[0], g2.MulScalar(new(bls12381.PointG2), gammas0[1], x1))
	mu11 := g2.Add(new(bls12381.PointG2), gammas1[0], g2.MulScalar(new(bls12381.PointG2), gammas1[1], x1))
	mu12 := g2.Add(new(bls12381.PointG2), gammas2[0], g2.MulScalar(new(bls12381.PointG2), gammas2[1], x1))
	mu20 := g2.Add(new(bls12381.PointG2), gammas0[0], g2.MulScalar(new(bls12381.PointG2), gammas0[1], x2))
	mu21 := g2.Add(new(bls12381.PointG2), gammas1[0], g2.MulScalar(new(bls12381.PointG2), gammas1[1], x2))
	mu22 := g2.Add(new(bls12381.PointG2), gammas2[0], g2.MulScalar(new(bls12381.PointG2), gammas2[1], x2))

	pks0 := g2.Add(new(bls12381.PointG2), mu00, mu01)
	g2.Add(pks0, pks0, mu02)
	pks1 := g2.Add(new(bls12381.PointG2), mu10, mu11)
	g2.Add(pks1, pks1, mu12)
	pks2 := g2.Add(new(bls12381.PointG2), mu20, mu21)
	g2.Add(pks2, pks2, mu22)

	assert.DeepEqual(t, pks0, (*bls12381.PointG2)(ComputeEonPublicKeyShare(0, gammas)), g2Comparer)
	assert.DeepEqual(t, pks1, (*bls12381.PointG2)(ComputeEonPublicKeyShare(1, gammas)), g2Comparer)
	assert.DeepEqual(t, pks2, (*bls12381.PointG2)(ComputeEonPublicKeyShare(2, gammas)), g2Comparer)
}

func TestEonSharesMatch(t *testing.T) {
	g2 := bls12381.NewG2()
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

	epk1Exp := (*EonPublicKeyShare)(g2.MulScalar(new(bls12381.PointG2), g2.One(), (*big.Int)(esk1)))
	epk2Exp := (*EonPublicKeyShare)(g2.MulScalar(new(bls12381.PointG2), g2.One(), (*big.Int)(esk2)))
	epk3Exp := (*EonPublicKeyShare)(g2.MulScalar(new(bls12381.PointG2), g2.One(), (*big.Int)(esk3)))

	assert.DeepEqual(t, epk1, epk1Exp)
	assert.DeepEqual(t, epk2, epk2Exp)
	assert.DeepEqual(t, epk3, epk3Exp)
}

func TestEonPublicKey(t *testing.T) {
	g2 := bls12381.NewG2()
	zeroEPK := ComputeEonPublicKey([]*Gammas{})
	assert.DeepEqual(t, (*bls12381.PointG2)(zeroEPK), g2.Zero(), g2Comparer)

	threshold := uint64(2)
	p1, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)

	k1 := ComputeEonPublicKey([]*Gammas{p1.Gammas()})
	assert.DeepEqual(t, (*bls12381.PointG2)(k1), []*bls12381.PointG2(*p1.Gammas())[0], g2Comparer)
	k2 := ComputeEonPublicKey([]*Gammas{p2.Gammas()})
	assert.DeepEqual(t, (*bls12381.PointG2)(k2), []*bls12381.PointG2(*p2.Gammas())[0], g2Comparer)
	k3 := ComputeEonPublicKey([]*Gammas{p3.Gammas()})
	assert.DeepEqual(t, (*bls12381.PointG2)(k3), []*bls12381.PointG2(*p3.Gammas())[0], g2Comparer)
}

func TestEonPublicKeyMatchesSecretKey(t *testing.T) {
	g2 := bls12381.NewG2()
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
	epkExp := g2.MulScalar(new(bls12381.PointG2), g2.One(), esk)

	gammas := []*Gammas{p1.Gammas(), p2.Gammas(), p3.Gammas()}
	epk := ComputeEonPublicKey(gammas)
	assert.DeepEqual(t, (*bls12381.PointG2)(epk), epkExp, g2Comparer)
}

var modOrderComparer = gocmp.Comparer(func(x, y *big.Int) bool {
	d := new(big.Int).Sub(x, y)
	return d.Mod(d, order).Sign() == 0
})

func TestInverse(t *testing.T) {
	testCases := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		new(big.Int).Sub(order, big.NewInt(2)),
		new(big.Int).Sub(order, big.NewInt(1)),
	}
	for i := 0; i < 100; i++ {
		x, err := rand.Int(rand.Reader, order)
		assert.NilError(t, err)
		if x.Sign() == 0 {
			continue
		}
		testCases = append(testCases, x)
	}

	for _, test := range testCases {
		inv := invert(test)
		one := new(big.Int).Mul(test, inv)
		assert.DeepEqual(t, big.NewInt(1), one, modOrderComparer)
	}
}

func TestLagrangeCoefficientFactors(t *testing.T) {
	l01 := lagrangeCoefficientFactor(0, 1)
	l02 := lagrangeCoefficientFactor(0, 2)
	l10 := lagrangeCoefficientFactor(1, 0)
	l12 := lagrangeCoefficientFactor(1, 2)
	l20 := lagrangeCoefficientFactor(2, 0)
	l21 := lagrangeCoefficientFactor(2, 1)

	qMinus1 := new(big.Int).Sub(order, big.NewInt(1))
	qMinus2 := new(big.Int).Sub(order, big.NewInt(2))

	l01.Mul(l01, qMinus1)
	assert.DeepEqual(t, big.NewInt(1), l01, modOrderComparer)
	l02.Mul(l02, qMinus2)
	assert.DeepEqual(t, big.NewInt(1), l02, modOrderComparer)

	assert.DeepEqual(t, big.NewInt(2), l10, modOrderComparer)
	l12.Mul(l12, qMinus1)
	assert.DeepEqual(t, big.NewInt(2), l12, modOrderComparer)

	l20.Mul(l20, big.NewInt(2))
	assert.DeepEqual(t, big.NewInt(3), l20, modOrderComparer)
	assert.DeepEqual(t, big.NewInt(3), l21, modOrderComparer)
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
	assert.DeepEqual(t, l0Exp, l0, modOrderComparer)

	l1 := lagrangeCoefficient(1, []int{0, 1, 2})
	l1Exp := lagrangeCoefficientFactor(0, 1)
	l1Exp.Mul(l1Exp, lagrangeCoefficientFactor(2, 1))
	assert.DeepEqual(t, l1Exp, l1, modOrderComparer)

	l2 := lagrangeCoefficient(2, []int{0, 1, 2})
	l2Exp := lagrangeCoefficientFactor(0, 2)
	l2Exp.Mul(l2Exp, lagrangeCoefficientFactor(1, 2))
	assert.DeepEqual(t, l2Exp, l2, modOrderComparer)
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
	y1.Mod(y1, order)
	y2.Mod(y2, order)
	y3.Mod(y3, order)

	y := new(big.Int).Add(y1, y2)
	y.Add(y, y3)
	y.Mod(y, order)

	assert.DeepEqual(t, p.Eval(big.NewInt(0)), y, shtest.BigIntComparer)
}

func TestComputeEpochSecretKeyShare(t *testing.T) {
	g1 := bls12381.NewG1()
	eonSecretKeyShare := (*EonSecretKeyShare)(big.NewInt(123))
	epochID := ComputeEpochID([]byte("epoch1"))
	epochSecretKeyShare := ComputeEpochSecretKeyShare(eonSecretKeyShare, epochID)
	expectedEpochSecretKeyShare := g1.MulScalar(new(bls12381.PointG1), (*bls12381.PointG1)(epochID), (*big.Int)(eonSecretKeyShare))
	assert.DeepEqual(t, expectedEpochSecretKeyShare, (*bls12381.PointG1)(epochSecretKeyShare), g1Comparer)
}

func TestVerifyEpochSecretKeyShare(t *testing.T) {
	threshold := uint64(2)
	epochID := ComputeEpochID([]byte("epoch1"))
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
	assert.Assert(t, !VerifyEpochSecretKeyShare(epsk1, epk1, ComputeEpochID([]byte("epoch2"))))
}

func TestVerifyEpochSecretKey(t *testing.T) {
	p, err := RandomPolynomial(rand.Reader, 0)
	assert.NilError(t, err)
	eonPublicKey := ComputeEonPublicKey([]*Gammas{p.Gammas()})

	epochIDBytes := []byte("epoch1")
	epochID := ComputeEpochID(epochIDBytes)

	v := p.EvalForKeyper(0)
	eonSecretKeyShare := ComputeEonSecretKeyShare([]*big.Int{v})
	epochSecretKeyShare := ComputeEpochSecretKeyShare(eonSecretKeyShare, epochID)
	epochSecretKey, err := ComputeEpochSecretKey(
		[]int{0},
		[]*EpochSecretKeyShare{epochSecretKeyShare},
		1,
	)
	assert.NilError(t, err)

	ok, err := VerifyEpochSecretKey(epochSecretKey, eonPublicKey, epochIDBytes)
	assert.NilError(t, err)
	assert.Check(t, ok)

	ok, err = VerifyEpochSecretKey(epochSecretKey, eonPublicKey, append(epochIDBytes, 0xab))
	assert.NilError(t, err)
	assert.Check(t, !ok)

	var sigma Block
	message := []byte("msg")
	ok, err = VerifyEpochSecretKeyDeterministic(epochSecretKey, eonPublicKey, epochIDBytes, sigma, message)
	assert.NilError(t, err)
	assert.Check(t, ok)

	ok, err = VerifyEpochSecretKeyDeterministic(epochSecretKey, eonPublicKey, append(epochIDBytes, 0xab), sigma, message)
	assert.NilError(t, err)
	assert.Check(t, !ok)
}

func TestComputeEpochSecretKey(t *testing.T) {
	n := 3
	threshold := uint64(2)
	epochID := ComputeEpochID([]byte("epoch1"))

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
	epochID := ComputeEpochID([]byte("epoch1"))

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
