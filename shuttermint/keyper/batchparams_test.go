package keyper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewBatchParams(t *testing.T) {
	b := NewBatchParams(0)
	t.Logf("%+v", b)
	require.Equal(t, int64(0), b.PublicKeyGenerationStartTime.Unix())
}

func TestNextBatchIndex(t *testing.T) {
	b := NewBatchParams(319773219)
	t.Logf("%+v", b)
	require.Equal(t, b.BatchIndex, NextBatchIndex(b.PublicKeyGenerationStartTime))
	ts := b.PublicKeyGenerationStartTime.Add(time.Duration(1))
	b2 := NextBatchIndex(ts)
	require.Equal(t, b.BatchIndex+1, b2)
}
