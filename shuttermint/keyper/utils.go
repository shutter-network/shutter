package keyper

import (
	"crypto/ecdsa"

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

// DecryptTransaction decrypts a single transaction with the given key.
func DecryptTransaction(key *ecdsa.PrivateKey, cipherTx []byte) []byte {
	return cipherTx
}

// DecryptTransactions decrypts a slice of transactions, presumably from a batch, with the given
// key.
func DecryptTransactions(key *ecdsa.PrivateKey, cipherTxs [][]byte) [][]byte {
	var txs [][]byte
	for _, cipherTx := range cipherTxs {
		txs = append(txs, cipherTx)
	}
	return txs
}

// DecryptionKeyToBytes converts a decryption key to its byte representation.
func DecryptionKeyToBytes(key *ecdsa.PrivateKey) []byte {
	keyBytes := key.D.Bytes()
	result := make([]byte, 32)
	copy(result[32-len(keyBytes):], keyBytes)
	return result
}

func computeDecryptionSignatureHash(
	batcherContractAddress common.Address,
	cipherBatchHash common.Hash,
	decryptionKey *ecdsa.PrivateKey,
	batchHash common.Hash,
) []byte {
	return crypto.Keccak256(
		batcherContractAddress.Bytes(),
		cipherBatchHash.Bytes(),
		DecryptionKeyToBytes(decryptionKey),
		batchHash.Bytes(),
	)
}

// ComputeDecryptionSignature computes the signature to be submitted for each keyper when
// executing a batch.
func ComputeDecryptionSignature(
	key *ecdsa.PrivateKey,
	batcherContractAddress common.Address,
	cipherBatchHash common.Hash,
	decryptionKey *ecdsa.PrivateKey,
	batchHash common.Hash,
) ([]byte, error) {
	hash := computeDecryptionSignatureHash(
		batcherContractAddress,
		cipherBatchHash,
		decryptionKey,
		batchHash,
	)
	sig, err := crypto.Sign(hash, key)
	if err != nil {
		return []byte{}, err
	}

	copy(sig[64:], []byte{sig[64] + 27})
	return sig, nil
}

// RecoverDecryptionSignatureSigner recovers the address of the account that has signed the
// decryption message.
func RecoverDecryptionSignatureSigner(
	signature []byte,
	batcherContractAddress common.Address,
	cipherBatchHash common.Hash,
	decryptionKey *ecdsa.PrivateKey,
	batchHash common.Hash,
) (common.Address, error) {
	hash := computeDecryptionSignatureHash(
		batcherContractAddress,
		cipherBatchHash,
		decryptionKey,
		batchHash,
	)

	sig := make([]byte, 65)
	copy(sig[:64], signature[:64])
	copy(sig[64:], []byte{signature[64] - 27})

	pubkey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}
