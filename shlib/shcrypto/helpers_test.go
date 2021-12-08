package shcrypto

import gocmp "github.com/google/go-cmp/cmp"

var (
	g1Comparer = gocmp.Comparer(EqualG1)
	g2Comparer = gocmp.Comparer(EqualG2)
)
