package app

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// GenesisAppState is used to hold the initial list of keypers, who will bootstrap the system by
// providing the first real BatchConfig to be used. We use common.MixedcaseAddress to hold the list
// of keypers as that one serializes as checksum address.
type GenesisAppState struct {
	Keypers   []common.MixedcaseAddress `json:"keypers"`
	Threshold uint64                    `json:"threshold"`
}

func NewGenesisAppState(keypers []common.Address, threshold int) GenesisAppState {
	appState := GenesisAppState{Threshold: uint64(threshold)}
	for _, k := range keypers {
		appState.Keypers = append(appState.Keypers, common.NewMixedcaseAddress(k))
	}
	return appState
}

func (appState *GenesisAppState) GetKeypers() []common.Address {
	var res []common.Address
	for _, k := range appState.Keypers {
		res = append(res, k.Address())
	}
	return res
}

// BatchConfig is the configuration we use for a consecutive sequence of batches.
// This should be synchronized with the list of BatchConfig structures stored in the ConfigContract
// deployed on the main chain.
type BatchConfig struct {
	StartBatchIndex uint64
	Keypers         []common.Address
	Threshold       uint64
	ConfigIndex     uint64
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
// We use this is a map key, so don't use a byte slice
type ValidatorPubkey struct {
	ed25519pubkey string
}

func (vp ValidatorPubkey) String() string {
	return fmt.Sprintf("ed25519:%s", hex.EncodeToString([]byte(vp.ed25519pubkey)))
}

// Powermap maps a ValidatorPubkey to the validators voting power
type Powermap map[ValidatorPubkey]int64

// NewValidatorPubkey creates a new ValidatorPubkey from a 32 byte ed25519 raw pubkey. See
// https://docs.tendermint.com/master/spec/abci/apps.html#validator-updates for more information
func NewValidatorPubkey(pubkey []byte) (ValidatorPubkey, error) {
	if len(pubkey) != ed25519.PublicKeySize {
		return ValidatorPubkey{}, fmt.Errorf("pubkey must be 32 bytes")
	}
	return ValidatorPubkey{ed25519pubkey: string(pubkey)}, nil
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
