package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

func TestCheckTxStateCouting(t *testing.T) {
	a1 := common.BigToAddress(big.NewInt(0))
	a2 := common.BigToAddress(big.NewInt(1))
	n := uint64(0)

	makeMsg := func() *shmsg.MessageWithNonce {
		n++
		return &shmsg.MessageWithNonce{
			RandomNonce: n,
			Msg:         nil,
		}
	}

	s := NewCheckTxState()
	require.Zero(t, len(s.Members))
	require.Zero(t, len(s.TxCounts))

	require.True(t, s.AddTx(a1, makeMsg())) // no members set yet, so anyone can send
	s.Reset()

	s.SetMembers([]common.Address{a1})
	require.False(t, s.AddTx(a2, makeMsg())) // not a member

	for i := 0; i < MaxTxsPerBlock; i++ {
		require.True(t, s.AddTx(a1, makeMsg()))
	}
	require.False(t, s.AddTx(a1, makeMsg())) // too many txs
	s.Reset()
	require.True(t, s.AddTx(a1, makeMsg()))
}

func TestCheckTxStateNonce(t *testing.T) {
	a1 := common.BigToAddress(big.NewInt(0))
	a2 := common.BigToAddress(big.NewInt(1))

	msg1 := &shmsg.MessageWithNonce{
		RandomNonce: uint64(10),
		Msg:         nil,
	}
	msg2 := &shmsg.MessageWithNonce{
		RandomNonce: uint64(20),
		Msg:         nil,
	}

	s := NewCheckTxState()
	s.SetMembers([]common.Address{a1, a2})

	require.True(t, s.AddTx(a1, msg1))
	require.False(t, s.AddTx(a1, msg1))
	require.True(t, s.AddTx(a1, msg2))
	require.True(t, s.AddTx(a2, msg1))
	s.Reset()
	require.True(t, s.AddTx(a1, msg1))
}
