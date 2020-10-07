package keyper

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestBatchHash(t *testing.T) {
	var zeroHash common.Hash
	tx1 := []byte{0xaa, 0xaa}
	tx2 := []byte{0xbb, 0xbb}

	// computed with Python contract test contract utils
	bh1 := common.HexToHash("15455e7675bd6d7a73a2cb8eb54a354f8180881d4c6b410556b505814b94cda1")
	bh2 := common.HexToHash("5081bbf18eac5fd3c87fb408bd54dca6008ec51aca62e0e4ca4631b96aa43f58")
	bh12 := common.HexToHash("bebf2f754c45520216ca08d65c4fef1b2dc98762f85dbabc3ce6abe23086ddbc")
	bh21 := common.HexToHash("2cb6f9ae6401f16792872bdeabc5f29cb0e52a949e35399e322a5af8bddd1489")

	require.Equal(t, ComputeBatchHash([][]byte{}), zeroHash)
	require.Equal(t, ComputeBatchHash([][]byte{tx1}), bh1)
	require.Equal(t, ComputeBatchHash([][]byte{tx2}), bh2)
	require.Equal(t, ComputeBatchHash([][]byte{tx1, tx2}), bh12)
	require.Equal(t, ComputeBatchHash([][]byte{tx2, tx1}), bh21)
}
