package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestNonceTracker(t *testing.T) {
	tracker := NewNonceTracker()
	a1 := common.BigToAddress(big.NewInt(0))
	a2 := common.BigToAddress(big.NewInt(1))
	r1 := uint64(10)
	r2 := uint64(20)

	require.True(t, tracker.Check(a1, r1))
	tracker.Add(a1, r1)
	require.False(t, tracker.Check(a1, r1))
	require.True(t, tracker.Check(a1, r2))
	require.True(t, tracker.Check(a2, r1))
}
