package keyper

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ComputeBatchHash computes the batch hash from the sequence of transactions the batch consists
// of.
func ComputeBatchHash(txs [][]byte) common.Hash {
	var hash common.Hash
	for _, tx := range txs {
		hash = crypto.Keccak256Hash(tx, hash.Bytes())
	}
	return hash
}
