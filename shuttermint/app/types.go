package app

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"fmt"
	"time"

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
	EncryptionKey         []byte
	Signature             []byte
	BatchIndex            uint64
	Sender                common.Address
	ConfigContractAddress common.Address
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

// ValidatorPubkey holds the raw 32 byte ed25519 public key to be used as tendermint validator key
type ValidatorPubkey struct {
	Data []byte
}

// NewValidatorPubkey creates a new ValidatorPubkey from a 32 byte ed25519 raw pubkey. See
// https://docs.tendermint.com/master/spec/abci/apps.html#validator-updates for more information
func NewValidatorPubkey(pubkey []byte) (ValidatorPubkey, error) {
	if len(pubkey) != ed25519.PublicKeySize {
		return ValidatorPubkey{}, fmt.Errorf("pubkey must be 32 bytes")
	}
	return ValidatorPubkey{Data: pubkey}, nil
}

// ShutterApp holds our data structures used for the tendermint app.
type ShutterApp struct {
	Configs         []*BatchConfig
	BatchStates     map[uint64]BatchState
	Voting          ConfigVoting
	Gobpath         string
	LastSaved       time.Time
	LastBlockHeight int64
	Identities      map[common.Address]ValidatorPubkey
}
