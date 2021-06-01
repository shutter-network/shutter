package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gotest.tools/v3/assert"
)

func TestNonceTracker(t *testing.T) {
	tracker := NewNonceTracker()
	a1 := common.BigToAddress(big.NewInt(0))
	a2 := common.BigToAddress(big.NewInt(1))
	r1 := uint64(10)
	r2 := uint64(20)

	assert.Assert(t, tracker.Check(a1, r1))
	tracker.Add(a1, r1)
	assert.Assert(t, !tracker.Check(a1, r1))
	assert.Assert(t, tracker.Check(a1, r2))
	assert.Assert(t, tracker.Check(a2, r1))
}
