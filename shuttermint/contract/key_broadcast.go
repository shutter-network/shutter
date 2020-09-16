package contract

// This file adds some custom methods to the abigen generated KeyBroadcastContract class
import (
	"math/big"

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

	signerIndicesBig := []*big.Int{}
	for _, signerIndex := range signerIndices {
		signerIndexBig := big.NewInt(int64(signerIndex))
		signerIndicesBig = append(signerIndicesBig, signerIndexBig)
	}

	return kbc.BroadcastEncryptionKey(
		auth,
		big.NewInt(int64(keyperIndex)),
		big.NewInt(int64(batchIndex)),
		encryptionKeySized,
		signerIndicesBig,
		signatures,
	)
}
