package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCheckTxState(t *testing.T) {
	a1 := common.BigToAddress(big.NewInt(0))
	a2 := common.BigToAddress(big.NewInt(1))

	s := NewCheckTxState()
	require.Zero(t, len(s.Members))
	require.Zero(t, len(s.TxCounts))

	require.True(t, s.AddTx(a1, nil)) // no members set yet, so anyone can send
	s.Reset()

	s.SetMembers([]common.Address{a1})
	require.False(t, s.AddTx(a2, nil)) // not a member

	for i := 0; i < MaxTxsPerBlock; i++ {
		require.True(t, s.AddTx(a1, nil))
	}
	require.False(t, s.AddTx(a1, nil)) // too many txs
	s.Reset()
	require.True(t, s.AddTx(a1, nil))
}
