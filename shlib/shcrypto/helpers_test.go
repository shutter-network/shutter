package shcrypto

import (
	gocmp "github.com/google/go-cmp/cmp"
	blst "github.com/supranational/blst/bindings/go"
)

func equalG1(a, b *blst.P1Affine) bool {
	return a.Equals(b)
}

func equalG2(a, b *blst.P2Affine) bool {
	return a.Equals(b)
}

var (
	g1Comparer = gocmp.Comparer(equalG1)
	g2Comparer = gocmp.Comparer(equalG2)
)
