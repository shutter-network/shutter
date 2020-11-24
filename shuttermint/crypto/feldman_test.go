package crypto

import (
	"crypto/rand"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/stretchr/testify/require"
)

func TestNewPolynomial(t *testing.T) {
	validCoefficients := [][]*big.Int{
		[]*big.Int{
			big.NewInt(10),
		},
		[]*big.Int{
			big.NewInt(0),
			big.NewInt(10),
			big.NewInt(20),
		},
		[]*big.Int{
			new(big.Int).Sub(bn256.Order, big.NewInt(1)),
		},
	}

	for _, cs := range validCoefficients {
		p, err := NewPolynomial(cs)
		require.Nil(t, err)
		for i, c := range cs {
			require.Equal(t, c, (*p)[i])
		}
		require.Equal(t, uint64(len(cs)-1), p.Degree())
	}

	invalidCoefficients := [][]*big.Int{
		[]*big.Int{},
		[]*big.Int{
			big.NewInt(-1),
		},
		[]*big.Int{
			bn256.Order,
		},
	}

	for _, cs := range invalidCoefficients {
		_, err := NewPolynomial(cs)
		require.NotNil(t, err)
	}
}

func TestEval(t *testing.T) {
	p1, err := NewPolynomial([]*big.Int{big.NewInt(10), big.NewInt(20), big.NewInt(30)})
	require.Nil(t, err)
	require.Zero(t, p1.Eval(big.NewInt(0)).Cmp(big.NewInt(10)))
	require.Zero(t, p1.Eval(big.NewInt(10)).Cmp(big.NewInt(10+20*10+30*100)))

	p2, err := NewPolynomial([]*big.Int{big.NewInt(0), new(big.Int).Sub(bn256.Order, big.NewInt(1))})
	require.Nil(t, err)
	require.Zero(t, p2.Eval(big.NewInt(0)).Cmp(big.NewInt(0)))
	require.Zero(t, p2.Eval(big.NewInt(1)).Cmp(new(big.Int).Sub(bn256.Order, big.NewInt(1))))
	require.Zero(t, p2.Eval(big.NewInt(2)).Cmp(new(big.Int).Sub(bn256.Order, big.NewInt(2))))

	p3, err := NewPolynomial([]*big.Int{big.NewInt(0), big.NewInt(1)})
	require.Nil(t, err)
	require.Zero(t, p3.Eval(big.NewInt(0)).Cmp(big.NewInt(0)))
	require.Zero(t, p3.Eval(bn256.Order).Cmp(big.NewInt(0)))
	require.Zero(t, p3.Eval(new(big.Int).Mul(bn256.Order, big.NewInt(5))).Cmp(big.NewInt(0)))
}

func TestEvalForKeyper(t *testing.T) {
	p, err := NewPolynomial([]*big.Int{big.NewInt(10), big.NewInt(20), big.NewInt(30)})
	require.Nil(t, err)
	v0 := p.EvalForKeyper(0)
	v1 := p.EvalForKeyper(1)
	require.Zero(t, v0.Cmp(p.Eval(KeyperX(0))))
	require.Zero(t, v1.Cmp(p.Eval(KeyperX(1))))
	require.NotZero(t, v0.Cmp(v1))
}

func TestValidEval(t *testing.T) {
	valid := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		new(big.Int).Sub(bn256.Order, big.NewInt(2)),
		new(big.Int).Sub(bn256.Order, big.NewInt(1)),
	}

	invalid := []*big.Int{
		big.NewInt(-2),
		big.NewInt(-1),
		bn256.Order,
		new(big.Int).Add(bn256.Order, big.NewInt(1)),
	}
	for _, v := range valid {
		require.True(t, ValidEval(v))
	}
	for _, v := range invalid {
		require.False(t, ValidEval(v))
	}
}

func TestRandomPolynomial(t *testing.T) {
	p, err := RandomPolynomial(rand.Reader, uint64(5))
	require.Nil(t, err)
	require.Equal(t, p.Degree(), uint64(5))
}

func TestGammas(t *testing.T) {
	p, err := NewPolynomial([]*big.Int{
		big.NewInt(0),
		big.NewInt(10),
		big.NewInt(20),
	})
	require.Nil(t, err)
	gammas := p.Gammas()
	require.Equal(t, p.Degree(), uint64(len(*gammas))-1)
	require.Equal(t, p.Degree(), gammas.Degree())
	require.Equal(t, new(bn256.G2).ScalarBaseMult(big.NewInt(0)), (*gammas)[0])
	require.Equal(t, new(bn256.G2).ScalarBaseMult(big.NewInt(10)), (*gammas)[1])
	require.Equal(t, new(bn256.G2).ScalarBaseMult(big.NewInt(20)), (*gammas)[2])
}

func TestZeroGammas(t *testing.T) {
	g := ZeroGammas(uint64(3))
	require.Equal(t, 4, len(*g))
	for _, p := range *g {
		require.True(t, EqualG2(p, zeroG2))
	}
}

func TestVerifyPolyEval(t *testing.T) {
	threshold := uint64(2)

	p1, err := RandomPolynomial(rand.Reader, threshold-1)
	require.Nil(t, err)

	p2, err := RandomPolynomial(rand.Reader, threshold-1)
	require.Nil(t, err)

	for i := 0; i < 10; i++ {
		xi := KeyperX(i)
		vi1 := p1.Eval(xi)
		vi2 := p2.Eval(xi)
		require.True(t, VerifyPolyEval(i, vi1, p1.Gammas(), threshold))
		require.True(t, VerifyPolyEval(i, vi2, p2.Gammas(), threshold))
		require.False(t, VerifyPolyEval(i, vi1, p2.Gammas(), threshold))
		require.False(t, VerifyPolyEval(i, vi2, p1.Gammas(), threshold))
		require.False(t, VerifyPolyEval(i+1, vi1, p1.Gammas(), threshold))
		require.False(t, VerifyPolyEval(i+1, vi2, p2.Gammas(), threshold))
	}
}

func TestPi(t *testing.T) {
	g1 := new(bn256.G2).ScalarBaseMult(big.NewInt(2))
	g2 := new(bn256.G2).ScalarBaseMult(big.NewInt(3))
	g3 := new(bn256.G2).ScalarBaseMult(big.NewInt(5))
	gammas := Gammas([]*bn256.G2{g1, g2, g3})

	pi1 := gammas.Pi(big.NewInt(1))
	pi2 := gammas.Pi(big.NewInt(2))

	pi1Exp := new(bn256.G2).Add(g1, g2)
	pi1Exp = new(bn256.G2).Add(pi1Exp, g3)
	pi2Exp := new(bn256.G2).Add(g1, new(bn256.G2).ScalarMult(g2, big.NewInt(2)))
	pi2Exp = new(bn256.G2).Add(pi2Exp, new(bn256.G2).ScalarMult(g3, big.NewInt(4)))

	require.True(t, EqualG2(pi1, pi1Exp))
	require.True(t, EqualG2(pi2, pi2Exp))
}
