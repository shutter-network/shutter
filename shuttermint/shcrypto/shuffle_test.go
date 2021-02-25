package shcrypto

import (
	"bytes"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/stretchr/testify/require"
)

func TestShuffleEmpty(t *testing.T) {
	txs := [][]byte{}
	key := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(5)))
	shuffledTxs := Shuffle(txs, key)
	require.Zero(t, len(shuffledTxs))
}

func TestShuffleSingle(t *testing.T) {
	txs := [][]byte{{1, 2, 3}}
	key := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(5)))
	shuffledTxs := Shuffle(txs, key)
	require.Equal(t, 1, len(shuffledTxs))
	require.True(t, bytes.Equal(txs[0], shuffledTxs[0]))
}

func TestShuffleMany(t *testing.T) {
	txs := [][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	key1 := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(5)))
	key2 := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(10)))
	shuffledTxs1 := Shuffle(txs, key1)
	shuffledTxs2 := Shuffle(txs, key2)

	require.Equal(t, len(txs), len(shuffledTxs1))
	require.Equal(t, len(txs), len(shuffledTxs2))

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
	require.False(t, equalOriginal)
	require.False(t, equalShuffled)
}
