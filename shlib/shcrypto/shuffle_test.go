package shcrypto

import (
	"bytes"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"gotest.tools/v3/assert"
)

func TestShuffleEmpty(t *testing.T) {
	txs := [][]byte{}
	key := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(5)))
	shuffledTxs := Shuffle(txs, key)
	assert.Assert(t, len(shuffledTxs) == 0)
}

func TestShuffleSingle(t *testing.T) {
	txs := [][]byte{{1, 2, 3}}
	key := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(5)))
	shuffledTxs := Shuffle(txs, key)
	assert.Equal(t, 1, len(shuffledTxs))
	assert.DeepEqual(t, txs[0], shuffledTxs[0])
}

func TestShuffleMany(t *testing.T) {
	txs := [][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	key1 := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(5)))
	key2 := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(10)))
	shuffledTxs1 := Shuffle(txs, key1)
	shuffledTxs2 := Shuffle(txs, key2)

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
