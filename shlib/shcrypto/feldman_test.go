package shcrypto

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/bls12381"
	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shtest"
)

func TestNewPolynomial(t *testing.T) {
	validCoefficients := [][]*big.Int{
		{
			big.NewInt(10),
		},
		{
			big.NewInt(0),
			big.NewInt(10),
			big.NewInt(20),
		},
		{
			new(big.Int).Sub(order, big.NewInt(1)),
		},
	}

	for _, cs := range validCoefficients {
		p, err := NewPolynomial(cs)
		assert.NilError(t, err)
		for i, c := range cs {
			assert.DeepEqual(t, c, (*p)[i], shtest.BigIntComparer)
		}
		assert.Equal(t, uint64(len(cs)-1), p.Degree())
	}

	invalidCoefficients := [][]*big.Int{
		{},
		{
			big.NewInt(-1),
		},
		{
			order,
		},
	}

	for _, cs := range invalidCoefficients {
		_, err := NewPolynomial(cs)
		assert.Assert(t, err != nil)
	}
}

func TestEval(t *testing.T) {
	p1, err := NewPolynomial([]*big.Int{big.NewInt(10), big.NewInt(20), big.NewInt(30)})
	assert.NilError(t, err)
	assert.DeepEqual(t, big.NewInt(10), p1.Eval(big.NewInt(0)), shtest.BigIntComparer)
	assert.DeepEqual(t, big.NewInt(10+20*10+30*100), p1.Eval(big.NewInt(10)), shtest.BigIntComparer)

	p2, err := NewPolynomial([]*big.Int{big.NewInt(0), new(big.Int).Sub(order, big.NewInt(1))})
	assert.NilError(t, err)
	assert.DeepEqual(t, big.NewInt(0), p2.Eval(big.NewInt(0)), shtest.BigIntComparer)
	assert.DeepEqual(t, new(big.Int).Sub(order, big.NewInt(1)), p2.Eval(big.NewInt(1)), shtest.BigIntComparer)
	assert.DeepEqual(t, new(big.Int).Sub(order, big.NewInt(2)), p2.Eval(big.NewInt(2)), shtest.BigIntComparer)

	p3, err := NewPolynomial([]*big.Int{big.NewInt(0), big.NewInt(1)})
	assert.NilError(t, err)
	assert.DeepEqual(t, big.NewInt(0), p3.Eval(big.NewInt(0)), shtest.BigIntComparer)
	assert.DeepEqual(t, big.NewInt(0), p3.Eval(order), shtest.BigIntComparer)
	assert.DeepEqual(t, big.NewInt(0), p3.Eval(new(big.Int).Mul(order, big.NewInt(5))), shtest.BigIntComparer)
}

func TestEvalForKeyper(t *testing.T) {
	p, err := NewPolynomial([]*big.Int{big.NewInt(10), big.NewInt(20), big.NewInt(30)})
	assert.NilError(t, err)
	v0 := p.EvalForKeyper(0)
	v1 := p.EvalForKeyper(1)
	assert.DeepEqual(t, v0, p.Eval(KeyperX(0)), shtest.BigIntComparer)
	assert.DeepEqual(t, v1, p.Eval(KeyperX(1)), shtest.BigIntComparer)
	assert.Assert(t, v0.Cmp(v1) != 0)
}

func TestValidEval(t *testing.T) {
	valid := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		new(big.Int).Sub(order, big.NewInt(2)),
		new(big.Int).Sub(order, big.NewInt(1)),
	}

	invalid := []*big.Int{
		big.NewInt(-2),
		big.NewInt(-1),
		order,
		new(big.Int).Add(order, big.NewInt(1)),
	}
	for _, v := range valid {
		assert.Assert(t, ValidEval(v))
	}
	for _, v := range invalid {
		assert.Assert(t, !ValidEval(v))
	}
}

func TestRandomPolynomial(t *testing.T) {
	p, err := RandomPolynomial(rand.Reader, uint64(5))
	assert.NilError(t, err)
	assert.Equal(t, p.Degree(), uint64(5))
}

func TestGammas(t *testing.T) {
	p, err := NewPolynomial([]*big.Int{
		big.NewInt(0),
		big.NewInt(10),
		big.NewInt(20),
	})
	assert.NilError(t, err)
	gammas := p.Gammas()
	assert.Equal(t, p.Degree(), uint64(len(*gammas))-1)
	assert.Equal(t, p.Degree(), gammas.Degree())

	expected := Gammas([]*bls12381.PointG2{
		makeTestG2(0),
		makeTestG2(10),
		makeTestG2(20),
	})
	assert.DeepEqual(t, &expected, gammas)
}

func TestZeroGammas(t *testing.T) {
	g2 := bls12381.NewG2()
	g := ZeroGammas(uint64(3))
	assert.Equal(t, 4, len(*g))
	for _, p := range *g {
		assert.DeepEqual(t, p, g2.Zero(), g2Comparer)
	}
}

func TestVerifyPolyEval(t *testing.T) {
	threshold := uint64(2)

	p1, err := RandomPolynomial(rand.Reader, threshold-1)
	assert.NilError(t, err)

	p2, err := RandomPolynomial(rand.Reader, threshold-1)
	assert.NilError(t, err)

	for i := 0; i < 10; i++ {
		xi := KeyperX(i)
		vi1 := p1.Eval(xi)
		vi2 := p2.Eval(xi)
		assert.Assert(t, VerifyPolyEval(i, vi1, p1.Gammas(), threshold))
		assert.Assert(t, VerifyPolyEval(i, vi2, p2.Gammas(), threshold))
		assert.Assert(t, !VerifyPolyEval(i, vi1, p2.Gammas(), threshold))
		assert.Assert(t, !VerifyPolyEval(i, vi2, p1.Gammas(), threshold))
		assert.Assert(t, !VerifyPolyEval(i+1, vi1, p1.Gammas(), threshold))
		assert.Assert(t, !VerifyPolyEval(i+1, vi2, p2.Gammas(), threshold))
	}
}

func TestPi(t *testing.T) {
	g2 := bls12381.NewG2()

	gamma1 := makeTestG2(2)
	gamma2 := makeTestG2(3)
	gamma3 := makeTestG2(5)
	gammas := Gammas([]*bls12381.PointG2{gamma1, gamma2, gamma3})

	pi1 := gammas.Pi(big.NewInt(1))
	pi2 := gammas.Pi(big.NewInt(2))

	pi1Exp := g2.Add(new(bls12381.PointG2), gamma1, gamma2)
	pi1Exp = g2.Add(pi1Exp, pi1Exp, gamma3)
	pi2Exp := g2.Add(new(bls12381.PointG2), gamma1, g2.MulScalar(new(bls12381.PointG2), gamma2, big.NewInt(2)))
	pi2Exp = g2.Add(pi2Exp, pi2Exp, g2.MulScalar(new(bls12381.PointG2), gamma3, big.NewInt(4)))

	assert.DeepEqual(t, pi1, pi1Exp, g2Comparer)
	assert.DeepEqual(t, pi2, pi2Exp, g2Comparer)
}

func TestGammasGobable(t *testing.T) {
	p, err := NewPolynomial([]*big.Int{
		big.NewInt(0),
		big.NewInt(10),
		big.NewInt(20),
	})
	assert.NilError(t, err)

	gammas := p.Gammas()
	deserialized := new(Gammas)
	shtest.EnsureGobable(t, gammas, deserialized)
}
