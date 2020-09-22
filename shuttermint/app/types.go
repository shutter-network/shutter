package app

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

// BatchConfig is the configuration we use for a consecutive sequence of batches.
// This should be synchronized with the list of BatchConfig structures stored in the ConfigContract
// deployed on the main chain.
type BatchConfig struct {
	StartBatchIndex uint64
	Keypers         []common.Address
	Threshold       uint32
}

// ConfigVoting is used to let the keypers vote on new BatchConfigs to be added
// Each keyper can vote exactly once
type ConfigVoting struct {
	Candidates []BatchConfig
	Votes      map[common.Address]int
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

// EncryptionKeyAttestation stores an attestation to an encryption key signed by one of the keypers.
type EncryptionKeyAttestation struct {
	Sender                common.Address
	EncryptionKey         []byte
	BatchIndex            uint64
	ConfigContractAddress common.Address
	Signature             []byte
}

// The BatchState structure is used to manage the key generation process for a certain batch
type BatchState struct {
	BatchIndex                uint64
	Config                    *BatchConfig
	Commitments               []PublicKeyCommitment
	SecretShares              []SecretShare
	PublicKey                 *ecdsa.PublicKey
	PrivateKey                *ecdsa.PrivateKey
	EncryptionKeyAttestations []EncryptionKeyAttestation
}

// ShutterApp holds our data structures used for the tendermint app.  At the moment we don't
// persist anything on disk. When starting tendermint, it will 'feed' us with all of the messages
// received via deliverMessage
type ShutterApp struct {
	Configs     []*BatchConfig
	BatchStates map[uint64]BatchState
	Voting      ConfigVoting
}
