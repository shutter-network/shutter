package contract

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// ExecuteCipherBatch2 calls ExecuteCipherBatch with the arguments converted into the right types
func (ec *ExecutorContract) ExecuteCipherBatch2(
	opts *bind.TransactOpts,
	_cipherBatchHash [32]byte,
	_transactions [][]byte,
	_decryptionKey *ecdsa.PrivateKey,
	_signerIndices []uint64,
	_signatures [][]byte,
) (*types.Transaction, error) {
	decryptionKeyBytes := crypto.FromECDSA(_decryptionKey)
	decryptionKeySized := [32]byte{}
	copy(decryptionKeySized[:], decryptionKeyBytes[:32])

	return ec.ExecuteCipherBatch(
		opts,
		_cipherBatchHash,
		_transactions,
		decryptionKeySized,
		_signerIndices,
		_signatures,
	)
}
