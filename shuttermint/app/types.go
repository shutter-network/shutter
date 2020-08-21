package app

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

// BatchConfig is the configuration we use for a consecutive sequence of batches.
// This should be synchronized with the list of BatchConfig structures stored in the ConfigContract
// deployed on the main chain.
type BatchConfig struct {
	Keypers    []common.Address
	Threshhold int32
}

// PublicKeyCommitment from one of the keypers. Since we only implement our 'fake' key generation
// this already holds the public key
type PublicKeyCommitment struct {
	Sender common.Address
	Pubkey []byte
}

// SecretShare stores a private key from one of the keypers.
type SecretShare struct {
	Sender  common.Address
	Privkey []byte
}

// The BatchKeys structure is used to manage the key generation process for a certain batch
type BatchKeys struct {
	Config       *BatchConfig
	Commitments  []PublicKeyCommitment
	SecretShares []SecretShare
	PublicKey    *ecdsa.PublicKey
	PrivateKey   *ecdsa.PrivateKey
}

// ShutterApp holds our data structures used for the tendermint app.  At the moment we don't
// persist anything on disk. When starting tendermint, it will 'feed' us with all of the messages
// received via deliverMessage
type ShutterApp struct {
	Batches map[uint64]BatchKeys
}
