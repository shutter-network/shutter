package shcrypto

import (
	"encoding/binary"
	"math/rand"

	"github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// Shuffle shuffles the order of the transactions using the epoch secret key as a source of
// randomness.
func Shuffle(txs [][]byte, key *EpochSecretKey) [][]byte {
	keyBytes := (*bn256.G1)(key).Marshal()
	keyHash := crypto.Keccak256(keyBytes)

	seed, n := binary.Varint(keyHash[:8])
	if n <= 0 {
		panic("failed to decode hash as int64")
	}
	source := rand.NewSource(seed)
	rand := rand.New(source)

	shuffledTxs := append([][]byte{}, txs...)
	swap := func(i, j int) {
		shuffledTxs[i], shuffledTxs[j] = shuffledTxs[j], shuffledTxs[i]
	}
	rand.Shuffle(len(shuffledTxs), swap)

	return shuffledTxs
}
