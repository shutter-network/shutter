package shcrypto

import (
	"crypto/rand"
	"math/big"
)

type TestKeyGen struct {
	n                  uint64
	threshold          uint64
	polynomials        []*Polynomial
	eonSecretKeyShares []*EonSecretKeyShare

	EonPublicKey *EonPublicKey
}

func NewTestKeyGen() (*TestKeyGen, error) {
	n := 3
	threshold := uint64(2)

	polynomials := []*Polynomial{}
	gammas := []*Gammas{}
	for i := 0; i < n; i++ {
		p, err := RandomPolynomial(rand.Reader, threshold-1)
		if err != nil {
			return nil, err
		}
		polynomials = append(polynomials, p)
		gammas = append(gammas, p.Gammas())

	}
	eonPublicKey := ComputeEonPublicKey(gammas)

	eonSecretKeyShares := []*EonSecretKeyShare{}
	for i := 0; i < n; i++ {
		ss := []*big.Int{}
		for j := 0; j < n; j++ {
			s := polynomials[j].EvalForKeyper(i)
			ss = append(ss, s)
		}
		eonSecretKeyShare := ComputeEonSecretKeyShare(ss)
		eonSecretKeyShares = append(eonSecretKeyShares, eonSecretKeyShare)
	}

	return &TestKeyGen{
		n:                  uint64(n),
		threshold:          threshold,
		polynomials:        polynomials,
		eonSecretKeyShares: eonSecretKeyShares,

		EonPublicKey: eonPublicKey,
	}, nil
}

func (g *TestKeyGen) ComputeEpochSecretKey(epochID *EpochID) (*EpochSecretKey, error) {
	indices := []int{}
	epochSecretKeyShares := []*EpochSecretKeyShare{}
	for i := 0; i < int(g.threshold); i++ {
		ss := []*big.Int{}
		for j := 0; j < int(g.n); j++ {
			s := g.polynomials[j].EvalForKeyper(i)
			ss = append(ss, s)
		}
		epochSecretKeyShare := ComputeEpochSecretKeyShare(g.eonSecretKeyShares[i], epochID)

		indices = append(indices, i)
		epochSecretKeyShares = append(epochSecretKeyShares, epochSecretKeyShare)
	}
	return ComputeEpochSecretKey(
		indices,
		epochSecretKeyShares,
		g.threshold,
	)
}
