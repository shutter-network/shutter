package contract

// This file adds some custom methods to the abigen generated KeyBroadcastContract class
import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// BroadcastEncryptionKey2 calls BroadcastEncryptionKey with the arguments converted into the right types
func (kbc *KeyBroadcastContract) BroadcastEncryptionKey2(
	auth *bind.TransactOpts,
	keyperIndex uint64,
	batchIndex uint64,
	encryptionKey []byte,
	signerIndices []uint64,
	signatures [][]byte) (*types.Transaction, error) {
	encryptionKeySized := [32]byte{}
	copy(encryptionKeySized[:], encryptionKey[:32])

	return kbc.BroadcastEncryptionKey(
		auth,
		keyperIndex,
		batchIndex,
		encryptionKeySized,
		signerIndices,
		signatures,
	)
}
