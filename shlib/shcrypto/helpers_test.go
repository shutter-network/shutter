package shcrypto

import (
	"github.com/ethereum/go-ethereum/crypto/bls12381"
	gocmp "github.com/google/go-cmp/cmp"
)

func equalG1(a, b *bls12381.PointG1) bool {
	g1 := bls12381.NewG1()
	return g1.Equal(a, b)
}

func equalG2(a, b *bls12381.PointG2) bool {
	g2 := bls12381.NewG2()
	return g2.Equal(a, b)
}

var (
	g1Comparer = gocmp.Comparer(equalG1)
	g2Comparer = gocmp.Comparer(equalG2)
)
