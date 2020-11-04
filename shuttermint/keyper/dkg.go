package keyper

import (
	"crypto/rand"
	"math/big"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

// DKGInstance represents the state of a single keyper participating in a DKG process.
type DKGInstance struct {
	Eon uint64

	PolyBase *big.Int
}

// NewDKGInstance creates a new dkg instance with initialized local random values.
func NewDKGInstance(eon uint64) (*DKGInstance, error) {
	polyBase, err := crypto.RandomPolynomialBase(rand.Reader)
	if err != nil {
		return nil, err
	}

	dkg := DKGInstance{
		Eon:      eon,
		PolyBase: polyBase,
	}
	return &dkg, nil
}
