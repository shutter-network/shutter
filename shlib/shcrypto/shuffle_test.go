package shcrypto

import (
	"bytes"
	"testing"

	"gotest.tools/v3/assert"
)

func TestShuffleEmpty(t *testing.T) {
	txs := [][]byte{}
	shuffledTxs := Shuffle(txs, (*EpochSecretKey)(makeTestG1(5)))
	assert.Assert(t, len(shuffledTxs) == 0)
}

func TestShuffleSingle(t *testing.T) {
	txs := [][]byte{{1, 2, 3}}
	shuffledTxs := Shuffle(txs, (*EpochSecretKey)(makeTestG1(5)))
	assert.Equal(t, 1, len(shuffledTxs))
	assert.DeepEqual(t, txs[0], shuffledTxs[0])
}

func TestShuffleMany(t *testing.T) {
	txs := [][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	shuffledTxs1 := Shuffle(txs, (*EpochSecretKey)(makeTestG1(5)))
	shuffledTxs2 := Shuffle(txs, (*EpochSecretKey)(makeTestG1(10)))

	assert.Equal(t, len(txs), len(shuffledTxs1))
	assert.Equal(t, len(txs), len(shuffledTxs2))

	equalOriginal := true
	equalShuffled := true
	for i := 0; i < len(txs); i++ {
		if !bytes.Equal(txs[i], shuffledTxs1[i]) {
			equalOriginal = false
		}
		if !bytes.Equal(shuffledTxs1[i], shuffledTxs2[i]) {
			equalShuffled = false
		}
	}
	assert.Assert(t, !equalOriginal)
	assert.Assert(t, !equalShuffled)
}
