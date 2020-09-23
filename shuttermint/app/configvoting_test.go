package app

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
)

var (
	addr  []common.Address
	votes []BatchConfig
)

func init() {
	for i := 0; i < 10; i++ {
		addr = append(addr, common.BigToAddress(big.NewInt(int64(i))))
		votes = append(votes, BatchConfig{StartBatchIndex: uint64(i)})
	}
}

func TestConfigVoting(t *testing.T) {
	cfgv := NewConfigVoting()
	// Make sure we don't have an outcome yet.
	_, ok := cfgv.Outcome(0)
	require.Equal(t, false, ok)
	_, ok = cfgv.Outcome(1)
	require.Equal(t, false, ok)

	err := cfgv.AddVote(addr[0], votes[0])
	require.Nil(t, err)

	outcome, ok := cfgv.Outcome(0)
	require.Equal(t, true, ok)
	require.Equal(t, votes[0], outcome)

	outcome, ok = cfgv.Outcome(1)
	require.Equal(t, true, ok)
	require.Equal(t, votes[0], outcome)

	_, ok = cfgv.Outcome(2)
	require.Equal(t, false, ok)

	err = cfgv.AddVote(addr[0], votes[0]) // duplicate vote, same vote
	require.NotNil(t, err, "voting two times should be prohibited")

	err = cfgv.AddVote(addr[0], votes[1]) // duplicate vote, different vote
	require.NotNil(t, err, "voting two times should be prohibited")

	err = cfgv.AddVote(addr[1], votes[2])
	require.Nil(t, err)

	_, ok = cfgv.Outcome(2)
	require.Equal(t, false, ok)

	err = cfgv.AddVote(addr[2], votes[2])
	require.Nil(t, err)

	outcome, ok = cfgv.Outcome(2)
	require.Equal(t, true, ok)
	require.Equal(t, votes[2], outcome)
}
