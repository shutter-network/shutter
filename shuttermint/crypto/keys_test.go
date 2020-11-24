package crypto

import (
	"crypto/rand"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"

	"github.com/stretchr/testify/require"
)

func TestEonSKShare(t *testing.T) {
	zeroKey := ComputeEonSKShare([]*big.Int{})
	require.Zero(t, (*big.Int)(zeroKey).Sign())

	key1 := ComputeEonSKShare([]*big.Int{
		big.NewInt(10),
		big.NewInt(20),
		big.NewInt(30),
	})
	require.Zero(t, big.NewInt(60).Cmp((*big.Int)(key1)))

	key2 := ComputeEonSKShare([]*big.Int{
		bn256.Order,
		big.NewInt(10),
		bn256.Order,
		big.NewInt(20),
		bn256.Order,
	})
	require.Zero(t, big.NewInt(30).Cmp((*big.Int)(key2)))
}

func TestEonPKShare(t *testing.T) {
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

	require.True(t, EqualG2(pks0, (*bn256.G2)(ComputeEonPKShare(0, gammas))))
	require.True(t, EqualG2(pks1, (*bn256.G2)(ComputeEonPKShare(1, gammas))))
	require.True(t, EqualG2(pks2, (*bn256.G2)(ComputeEonPKShare(2, gammas))))
}

func TestEonSharesMatch(t *testing.T) {
	threshold := uint64(2)
	p1, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)

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

	esk1 := ComputeEonSKShare([]*big.Int{v11, v12, v13})
	esk2 := ComputeEonSKShare([]*big.Int{v21, v22, v23})
	esk3 := ComputeEonSKShare([]*big.Int{v31, v32, v33})

	epk1 := ComputeEonPKShare(0, gammas)
	epk2 := ComputeEonPKShare(1, gammas)
	epk3 := ComputeEonPKShare(2, gammas)

	epk1Exp := new(bn256.G2).ScalarBaseMult((*big.Int)(esk1))
	epk2Exp := new(bn256.G2).ScalarBaseMult((*big.Int)(esk2))
	epk3Exp := new(bn256.G2).ScalarBaseMult((*big.Int)(esk3))

	require.True(t, EqualG2((*bn256.G2)(epk1), epk1Exp))
	require.True(t, EqualG2((*bn256.G2)(epk2), epk2Exp))
	require.True(t, EqualG2((*bn256.G2)(epk3), epk3Exp))
}

func TestEonPK(t *testing.T) {
	zeroEPK := ComputeEonPK([]*Gammas{})
	require.True(t, EqualG2((*bn256.G2)(zeroEPK), zeroG2))

	threshold := uint64(2)
	p1, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)

	k1 := ComputeEonPK([]*Gammas{p1.Gammas()})
	require.True(t, EqualG2((*bn256.G2)(k1), ([]*bn256.G2)(*p1.Gammas())[0]))
	k2 := ComputeEonPK([]*Gammas{p2.Gammas()})
	require.True(t, EqualG2((*bn256.G2)(k2), ([]*bn256.G2)(*p2.Gammas())[0]))
	k3 := ComputeEonPK([]*Gammas{p3.Gammas()})
	require.True(t, EqualG2((*bn256.G2)(k3), ([]*bn256.G2)(*p3.Gammas())[0]))
}

func TestEonPKMatchesSK(t *testing.T) {
	threshold := uint64(2)
	p1, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)

	esk := big.NewInt(0)
	for _, p := range []*Polynomial{p1, p2, p3} {
		esk = esk.Add(esk, p.Eval(big.NewInt(0)))
	}
	epkExp := new(bn256.G2).ScalarBaseMult(esk)

	gammas := []*Gammas{p1.Gammas(), p2.Gammas(), p3.Gammas()}
	epk := ComputeEonPK(gammas)
	require.True(t, EqualG2((*bn256.G2)(epk), epkExp))
}

func TestInverse(t *testing.T) {
	testCases := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		new(big.Int).Sub(bn256.Order, big.NewInt(2)),
		new(big.Int).Sub(bn256.Order, big.NewInt(1)),
	}
	for _, test := range testCases {
		inv := invert(test)
		one := new(big.Int).Mul(test, inv)
		one = one.Mod(one, bn256.Order)
		require.Equal(t, big.NewInt(1), one)
	}

	for i := 0; i < 100; i++ {
		x, err := rand.Int(rand.Reader, bn256.Order)
		require.Nil(t, err)
		if x.Sign() == 0 {
			continue
		}

		inv := invert(x)
		one := new(big.Int).Mul(x, inv)
		one = one.Mod(one, bn256.Order)
		require.Equal(t, big.NewInt(1), one)
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
	l01.Mod(l01, bn256.Order)
	require.Equal(t, big.NewInt(1), l01)
	l02.Mul(l02, qMinus2)
	l02.Mod(l02, bn256.Order)
	require.Equal(t, big.NewInt(1), l02)

	require.Equal(t, big.NewInt(2), l10)
	l12.Mul(l12, qMinus1)
	l12.Mod(l12, bn256.Order)
	require.Equal(t, big.NewInt(2), l12)

	l20.Mul(l20, big.NewInt(2))
	l20.Mod(l20, bn256.Order)
	require.Equal(t, big.NewInt(3), l20)
	require.Equal(t, big.NewInt(3), l21)
}

func TestLagrangeCoefficients(t *testing.T) {
	require.Equal(t, big.NewInt(1), lagrangeCoefficient(0, []int{0}))
	require.Equal(t, big.NewInt(1), lagrangeCoefficient(1, []int{1}))
	require.Equal(t, big.NewInt(1), lagrangeCoefficient(2, []int{2}))

	require.Equal(t, lagrangeCoefficientFactor(1, 0), lagrangeCoefficient(0, []int{0, 1}))
	require.Equal(t, lagrangeCoefficientFactor(0, 1), lagrangeCoefficient(1, []int{0, 1}))

	l0 := lagrangeCoefficient(0, []int{0, 1, 2})
	l0Exp := lagrangeCoefficientFactor(1, 0)
	l0Exp.Mul(l0Exp, lagrangeCoefficientFactor(2, 0))
	l0Exp.Mod(l0Exp, bn256.Order)
	require.Equal(t, l0Exp, l0)

	l1 := lagrangeCoefficient(1, []int{0, 1, 2})
	l1Exp := lagrangeCoefficientFactor(0, 1)
	l1Exp.Mul(l1Exp, lagrangeCoefficientFactor(2, 1))
	l1Exp.Mod(l1Exp, bn256.Order)
	require.Equal(t, l1Exp, l1)

	l2 := lagrangeCoefficient(2, []int{0, 1, 2})
	l2Exp := lagrangeCoefficientFactor(0, 2)
	l2Exp.Mul(l2Exp, lagrangeCoefficientFactor(1, 2))
	l2Exp.Mod(l2Exp, bn256.Order)
	require.Equal(t, l2Exp, l2)
}

func TestLagrangeReconstruct(t *testing.T) {
	p, err := RandomPolynomial(rand.Reader, uint64(2))
	require.Nil(t, err)

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

	require.Equal(t, p.Eval(big.NewInt(0)), y)
}

func TestComputeEpochSKShare(t *testing.T) {
	eonSKShare := (*EonSKShare)(big.NewInt(123))
	epochID := ComputeEpochID(uint64(456))
	epochSKShare := ComputeEpochSKShare(eonSKShare, epochID)
	expectedEpochSKShare := new(bn256.G1).ScalarMult((*bn256.G1)(epochID), (*big.Int)(eonSKShare))
	require.Equal(t, (*bn256.G1)(epochSKShare).String(), expectedEpochSKShare.String())
}

func TestVerifyEpochSKShare(t *testing.T) {
	threshold := uint64(2)
	epochID := ComputeEpochID(uint64(10))
	p1, err := RandomPolynomial(rand.Reader, threshold-1)
	require.Nil(t, err)
	p2, err := RandomPolynomial(rand.Reader, threshold-1)
	require.Nil(t, err)
	p3, err := RandomPolynomial(rand.Reader, threshold-1)
	require.Nil(t, err)

	gammas := []*Gammas{
		p1.Gammas(),
		p2.Gammas(),
		p3.Gammas(),
	}

	epk1 := ComputeEonPKShare(0, gammas)
	epk2 := ComputeEonPKShare(1, gammas)
	epk3 := ComputeEonPKShare(2, gammas)
	esk1 := ComputeEonSKShare([]*big.Int{p1.EvalForKeyper(0), p2.EvalForKeyper(0), p3.EvalForKeyper(0)})
	esk2 := ComputeEonSKShare([]*big.Int{p1.EvalForKeyper(1), p2.EvalForKeyper(1), p3.EvalForKeyper(1)})
	esk3 := ComputeEonSKShare([]*big.Int{p1.EvalForKeyper(2), p2.EvalForKeyper(2), p3.EvalForKeyper(2)})
	epsk1 := ComputeEpochSKShare(esk1, epochID)
	epsk2 := ComputeEpochSKShare(esk2, epochID)
	epsk3 := ComputeEpochSKShare(esk3, epochID)

	require.True(t, VerifyEpochSKShare(epsk1, epk1, epochID))
	require.True(t, VerifyEpochSKShare(epsk2, epk2, epochID))
	require.True(t, VerifyEpochSKShare(epsk3, epk3, epochID))

	require.False(t, VerifyEpochSKShare(epsk1, epk2, epochID))
	require.False(t, VerifyEpochSKShare(epsk2, epk1, epochID))
	require.False(t, VerifyEpochSKShare(epsk1, epk1, ComputeEpochID(uint64(11))))
}

func TestComputeEpochSK(t *testing.T) {
	n := 3
	threshold := uint64(2)
	epochID := ComputeEpochID(uint64(10))

	ps := []*Polynomial{}
	for i := 0; i < n; i++ {
		p, err := RandomPolynomial(rand.Reader, threshold-1)
		require.Nil(t, err)
		ps = append(ps, p)
	}

	eonSKShares := []*EonSKShare{}
	epochSKShares := []*EpochSKShare{}
	for i := 0; i < n; i++ {
		vs := []*big.Int{}
		for _, p := range ps {
			v := p.EvalForKeyper(i)
			vs = append(vs, v)
		}
		eonSKShare := ComputeEonSKShare(vs)
		epochSKShare := ComputeEpochSKShare(eonSKShare, epochID)

		eonSKShares = append(eonSKShares, eonSKShare)
		epochSKShares = append(epochSKShares, epochSKShare)
	}

	var err error
	_, err = ComputeEpochSK([]int{0}, epochSKShares[:1], threshold)
	require.NotNil(t, err)
	_, err = ComputeEpochSK([]int{0, 1, 2}, epochSKShares[:2], threshold)
	require.NotNil(t, err)
	_, err = ComputeEpochSK([]int{0, 1}, epochSKShares[:1], threshold)
	require.NotNil(t, err)
	_, err = ComputeEpochSK([]int{0}, epochSKShares[:2], threshold)
	require.NotNil(t, err)

	epochSK12, err := ComputeEpochSK([]int{0, 1}, []*EpochSKShare{epochSKShares[0], epochSKShares[1]}, threshold)
	require.Nil(t, err)
	epochSK13, err := ComputeEpochSK([]int{0, 2}, []*EpochSKShare{epochSKShares[0], epochSKShares[2]}, threshold)
	require.Nil(t, err)
	epochSK23, err := ComputeEpochSK([]int{1, 2}, []*EpochSKShare{epochSKShares[1], epochSKShares[2]}, threshold)
	require.Nil(t, err)

	require.True(t, EqualG1((*bn256.G1)(epochSK12), (*bn256.G1)(epochSK13)))
	require.True(t, EqualG1((*bn256.G1)(epochSK12), (*bn256.G1)(epochSK23)))
}

func TestFull(t *testing.T) {
	n := 3
	threshold := uint64(2)
	epochID := ComputeEpochID(uint64(10))

	ps := []*Polynomial{}
	gammas := []*Gammas{}
	for i := 0; i < n; i++ {
		p, err := RandomPolynomial(rand.Reader, threshold-1)
		require.Nil(t, err)
		ps = append(ps, p)
		gammas = append(gammas, p.Gammas())
	}

	eonSKShares := []*EonSKShare{}
	for i := 0; i < n; i++ {
		vs := []*big.Int{}
		for j := 0; j < n; j++ {
			v := ps[j].EvalForKeyper(i)
			vs = append(vs, v)
		}
		eonSKShare := ComputeEonSKShare(vs)
		eonSKShares = append(eonSKShares, eonSKShare)
	}

	eonPKShares := []*EonPKShare{}
	for i := 0; i < n; i++ {
		eonPKShare := ComputeEonPKShare(i, gammas)
		eonPKShares = append(eonPKShares, eonPKShare)
	}

	epochSKShares := []*EpochSKShare{}
	for i := 0; i < n; i++ {
		epochSKShare := ComputeEpochSKShare(eonSKShares[i], epochID)
		epochSKShares = append(epochSKShares, epochSKShare)
	}

	// verify (published) epoch sk shares
	for i := 0; i < n; i++ {
		require.True(t, VerifyEpochSKShare(epochSKShares[i], eonPKShares[i], epochID))
	}

	epochSK, err := ComputeEpochSK([]int{0, 1}, []*EpochSKShare{epochSKShares[0], epochSKShares[1]}, threshold)
	require.Nil(t, err)

	epochSK13, err := ComputeEpochSK([]int{0, 2}, []*EpochSKShare{epochSKShares[0], epochSKShares[2]}, threshold)
	require.Nil(t, err)
	epochSK23, err := ComputeEpochSK([]int{1, 2}, []*EpochSKShare{epochSKShares[1], epochSKShares[2]}, threshold)
	require.Nil(t, err)
	require.True(t, EqualG1((*bn256.G1)(epochSK), (*bn256.G1)(epochSK13)))
	require.True(t, EqualG1((*bn256.G1)(epochSK), (*bn256.G1)(epochSK23)))
}
