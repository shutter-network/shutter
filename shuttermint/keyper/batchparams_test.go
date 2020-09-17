package keyper

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestNewBatchParams(t *testing.T) {
	b := NewBatchParams(0, common.BigToAddress(big.NewInt(0)))
	t.Logf("%+v", b)
	require.Equal(t, int64(0), b.PublicKeyGenerationStartTime.Unix())
}

func TestNextBatchIndex(t *testing.T) {
	b := NewBatchParams(319773219, common.BigToAddress(big.NewInt(0)))
	t.Logf("%+v", b)
	require.Equal(t, b.BatchIndex, NextBatchIndex(b.PublicKeyGenerationStartTime))
	ts := b.PublicKeyGenerationStartTime.Add(time.Duration(1))
	b2 := NextBatchIndex(ts)
	require.Equal(t, b.BatchIndex+1, b2)
}
