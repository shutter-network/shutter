package app

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Add a prefix to avoid accidentally signing data with special meaning in different context, in
// particular Ethereum transactions (c.f. EIP191 https://eips.ethereum.org/EIPS/eip-191).
var hashPrefix = []byte{0x19, 'e', 'n', 'c', 's', 'i', 'g'}

// VerifySignature checks that the signature of the attestation is correct.
func (att *EncryptionKeyAttestation) VerifySignature() bool {
	preimage := EncryptionKeyPreimage(att.EncryptionKey, att.BatchIndex)
	hash := crypto.Keccak256Hash(preimage)

	pubkey, err := crypto.SigToPub(hash.Bytes(), att.Signature)
	if err != nil {
		return false
	}

	address := crypto.PubkeyToAddress(*pubkey)
	return address == att.Sender
}

// EncryptionKeyPreimage computes the preimage of the hash to be signed as part of the encryption
// key attestation.
func EncryptionKeyPreimage(key []byte, batchIndex uint64) []byte {
	// TODO: include config contract address
	batchIndexBig := new(big.Int).SetUint64(batchIndex)

	parts := [][]byte{
		hashPrefix,
		key,
		common.LeftPadBytes(batchIndexBig.Bytes(), 32),
	}

	var result []byte
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}
