package crypto

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

var (
	zeroG1 *bn256.G1
	zeroG2 *bn256.G2
)

// Polynomial represents a polynomial over Z_q
type Polynomial []*big.Int

// Gammas is a sequence of G2 points based on a polynomial.
type Gammas []*bn256.G2

func init() {
	zeroG1 = new(bn256.G1).ScalarBaseMult(big.NewInt(0))
	zeroG2 = new(bn256.G2).ScalarBaseMult(big.NewInt(0))
}

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
		if v.Cmp(bn256.Order) >= 0 {
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

// Eval evaluates the polynomial at the given coordinate.
func (p *Polynomial) Eval(x *big.Int) *big.Int {
	// uses Horner's method
	res := new(big.Int).Set((*p)[p.Degree()])
	for i := int(p.Degree()) - 1; i >= 0; i-- {
		res.Mul(res, x)
		res.Add(res, (*p)[i])
		res.Mod(res, bn256.Order)
	}
	return res
}

// EvalForKeyper evaluates the polynomial at the position designated for the given keyper.
func (p *Polynomial) EvalForKeyper(keyperIndex int) *big.Int {
	x := KeyperX(keyperIndex)
	return p.Eval(x)
}

// Gammas computes the gamma values for a given polynomial.
func (p *Polynomial) Gammas() *Gammas {
	gammas := Gammas{}
	for _, c := range *p {
		gamma := new(bn256.G2).ScalarBaseMult(c)
		gammas = append(gammas, gamma)
	}
	return &gammas
}

// Pi computes the pi value at the given x coordinate.
func (g *Gammas) Pi(xi *big.Int) *bn256.G2 {
	xiToJ := big.NewInt(1)
	res := new(bn256.G2).Set(zeroG2)
	for _, gamma := range *g {
		p := new(bn256.G2).ScalarMult(gamma, xiToJ)
		res = new(bn256.G2).Add(res, p)
		xiToJ.Mul(xiToJ, xi)
		xiToJ.Mod(xiToJ, bn256.Order)
	}
	return res
}

// KeyperX computes the x value assigned to the keyper identified by its index.
func KeyperX(keyperIndex int) *big.Int {
	keyperIndexBig := big.NewInt(int64(keyperIndex))
	return new(big.Int).Add(big.NewInt(1), keyperIndexBig)
}

// EqualG1 checks if two points on G1 are requal.
func EqualG1(p1, p2 *bn256.G1) bool {
	p1Bytes := new(bn256.G1).Set(p1).Marshal()
	p2Bytes := new(bn256.G1).Set(p2).Marshal()
	return bytes.Equal(p1Bytes, p2Bytes)
}

// EqualG2 checks if two points on G2 are requal.
func EqualG2(p1, p2 *bn256.G2) bool {
	p1Bytes := new(bn256.G2).Set(p1).Marshal()
	p2Bytes := new(bn256.G2).Set(p2).Marshal()
	return bytes.Equal(p1Bytes, p2Bytes)
}

// VerifyPolyEval checks that the evaluation of a polynomial is consistent with the public gammas.
func VerifyPolyEval(keyperIndex int, polyEval *big.Int, gammas *Gammas, threshold uint64) bool {
	if gammas.Degree() != threshold-1 {
		return false
	}
	rhs := new(bn256.G2).ScalarBaseMult(polyEval)
	lhs := gammas.Pi(KeyperX(keyperIndex))
	return EqualG2(lhs, rhs)
}

// RandomPolynomial generates a random polynomial of given degree.
func RandomPolynomial(r io.Reader, degree uint64) (*Polynomial, error) {
	coefficients := []*big.Int{}
	for i := uint64(0); i < degree+1; i++ {
		c, err := rand.Int(r, bn256.Order)
		if err != nil {
			return nil, err
		}
		coefficients = append(coefficients, c)
	}
	return NewPolynomial(coefficients)
}
