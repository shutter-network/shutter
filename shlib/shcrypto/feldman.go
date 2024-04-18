package shcrypto

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/bls12381"
)

// Polynomial represents a polynomial over Z_q.
type Polynomial []*big.Int

// Gammas is a sequence of G2 points based on a polynomial.
type Gammas []*bls12381.PointG2

var order = bls12381.NewG2().Q()

// NewPolynomial creates a new polynomial from the given coefficients. It verifies the number and
// range of them.
func NewPolynomial(coefficients []*big.Int) (*Polynomial, error) {
	if len(coefficients) == 0 {
		return nil, fmt.Errorf("no coefficients given")
	}
	for i, v := range coefficients {
		if v.Sign() < 0 {
			return nil, fmt.Errorf("coefficient %d is negative (%d)", i, v)
		}
		if v.Cmp(order) >= 0 {
			return nil, fmt.Errorf("coefficient %d is too big (%d)", i, v)
		}
	}
	p := Polynomial(coefficients)
	return &p, nil
}

// Degree returns the degree of the polynomial.
func (p *Polynomial) Degree() uint64 {
	return uint64(len(*p)) - 1
}

// Degree returns the degree of the underlying polynomial.
func (g *Gammas) Degree() uint64 {
	return uint64(len(*g)) - 1
}

func (g Gammas) Equal(otherG Gammas) bool {
	g2 := bls12381.NewG2()
	gs := []*bls12381.PointG2(g)
	otherGs := []*bls12381.PointG2(otherG)

	if len(gs) != len(otherGs) {
		return false
	}
	for i := range gs {
		if !g2.Equal(gs[i], otherGs[i]) {
			return false
		}
	}
	return true
}

// ZeroGammas returns the zero value for gammas.
func ZeroGammas(degree uint64) *Gammas {
	g2 := bls12381.NewG2()
	points := []*bls12381.PointG2{}
	for i := uint64(0); i < degree+1; i++ {
		points = append(points, g2.Zero())
	}
	gammas := Gammas(points)
	return &gammas
}

// DegreeFromThreshold returns the degree polynomials should have for the given threshold.
func DegreeFromThreshold(threshold uint64) uint64 {
	return threshold - 1
}

// Eval evaluates the polynomial at the given coordinate.
func (p *Polynomial) Eval(x *big.Int) *big.Int {
	// uses Horner's method
	res := new(big.Int).Set((*p)[p.Degree()])
	for i := int(p.Degree()) - 1; i >= 0; i-- {
		res.Mul(res, x)
		res.Add(res, (*p)[i])
		res.Mod(res, order)
	}
	return res
}

// EvalForKeyper evaluates the polynomial at the position designated for the given keyper.
func (p *Polynomial) EvalForKeyper(keyperIndex int) *big.Int {
	x := KeyperX(keyperIndex)
	return p.Eval(x)
}

// ValidEval checks if the given value is a valid polynomial evaluation, i.e., if it is in Z_q.
func ValidEval(v *big.Int) bool {
	if v.Sign() < 0 {
		return false
	}
	if v.Cmp(order) >= 0 {
		return false
	}
	return true
}

// Gammas computes the gamma values for a given polynomial.
func (p *Polynomial) Gammas() *Gammas {
	g2 := bls12381.NewG2()
	gammas := Gammas{}
	for _, c := range *p {
		gamma := g2.One()
		g2.MulScalar(gamma, gamma, c)
		gammas = append(gammas, gamma)
	}
	return &gammas
}

// Pi computes the pi value at the given x coordinate.
func (g *Gammas) Pi(xi *big.Int) *bls12381.PointG2 {
	g2 := bls12381.NewG2()
	xiToJ := big.NewInt(1)
	res := g2.Zero()
	p := new(bls12381.PointG2)
	for _, gamma := range *g {
		g2.MulScalar(p, gamma, xiToJ)
		g2.Add(res, res, p)
		xiToJ.Mul(xiToJ, xi)
		xiToJ.Mod(xiToJ, order)
	}
	return res
}

// GobEncode encodes a Gammas value. See https://golang.org/pkg/encoding/gob/#GobEncoder
func (g *Gammas) GobEncode() ([]byte, error) {
	g2 := bls12381.NewG2()
	buff := bytes.Buffer{}
	if g != nil {
		for _, p := range *g {
			buff.Write(g2.ToBytes(p))
		}
	}
	return buff.Bytes(), nil
}

// GobDecode decodes a Gammas value. See https://golang.org/pkg/encoding/gob/#GobDecoder
func (g *Gammas) GobDecode(data []byte) error {
	g2 := bls12381.NewG2()
	for i := 0; i < len(data); i += g2EncodingLength {
		p, err := g2.FromBytes(data[i : i+g2EncodingLength])
		if err != nil {
			return err
		}
		if !g2.IsOnCurve(p) {
			return errors.New("not on curve")
		}
		*g = append(*g, p)
	}
	return nil
}

// KeyperX computes the x value assigned to the keyper identified by its index.
func KeyperX(keyperIndex int) *big.Int {
	keyperIndexBig := big.NewInt(int64(keyperIndex))
	return new(big.Int).Add(big.NewInt(1), keyperIndexBig)
}

// VerifyPolyEval checks that the evaluation of a polynomial is consistent with the public gammas.
func VerifyPolyEval(keyperIndex int, polyEval *big.Int, gammas *Gammas, threshold uint64) bool {
	if gammas.Degree() != threshold-1 {
		return false
	}
	g2 := bls12381.NewG2()
	rhs := g2.One()
	g2.MulScalar(rhs, rhs, polyEval)
	lhs := gammas.Pi(KeyperX(keyperIndex))
	return g2.Equal(lhs, rhs)
}

// RandomPolynomial generates a random polynomial of given degree.
func RandomPolynomial(r io.Reader, degree uint64) (*Polynomial, error) {
	coefficients := []*big.Int{}
	for i := uint64(0); i < degree+1; i++ {
		c, err := rand.Int(r, order)
		if err != nil {
			return nil, err
		}
		coefficients = append(coefficients, c)
	}
	return NewPolynomial(coefficients)
}
