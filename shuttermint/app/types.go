package app

import (
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

// GetKeypers returns the keypers defined in the GenesisAppState
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
	Keypers         []common.Address
	StartBatchIndex uint64
	Threshold       uint64

	ConfigIndex           uint64
	ConfigContractAddress common.Address

	Started           bool
	ValidatorsUpdated bool
}

// ConfigVoting is used to let the keypers vote on new BatchConfigs to be added
// Each keyper can vote exactly once
type ConfigVoting struct {
	Candidates []BatchConfig
	Votes      map[common.Address]int
}

// EncryptionKeyAttestation stores an attestation to an encryption key signed by one of the keypers.
type EncryptionKeyAttestation struct {
	EncryptionKey         []byte
	Signature             []byte
	BatchIndex            uint64
	Sender                common.Address
	ConfigContractAddress common.Address
}

// DecryptionSignature stores the decryption key signature created by one of the keypers.
type DecryptionSignature struct {
	Sender    common.Address
	Signature []byte
}

// The BatchState structure is used to manage the key generation process for a certain batch
type BatchState struct {
	BatchIndex                uint64
	Config                    *BatchConfig
	EncryptionKeyAttestations []EncryptionKeyAttestation
	DecryptionSignatures      []DecryptionSignature
}

// ValidatorPubkey holds the raw 32 byte ed25519 public key to be used as tendermint validator key
// We use this is a map key, so don't use a byte slice
type ValidatorPubkey struct {
	Ed25519pubkey string
}

func (vp ValidatorPubkey) String() string {
	return fmt.Sprintf("ed25519:%s", hex.EncodeToString([]byte(vp.Ed25519pubkey)))
}

// Powermap maps a ValidatorPubkey to the validators voting power
type Powermap map[ValidatorPubkey]int64

// NewValidatorPubkey creates a new ValidatorPubkey from a 32 byte ed25519 raw pubkey. See
// https://docs.tendermint.com/master/spec/abci/apps.html#validator-updates for more information
func NewValidatorPubkey(pubkey []byte) (ValidatorPubkey, error) {
	if len(pubkey) != ed25519.PublicKeySize {
		return ValidatorPubkey{}, fmt.Errorf("pubkey must be 32 bytes")
	}
	return ValidatorPubkey{Ed25519pubkey: string(pubkey)}, nil
}

// ShutterApp holds our data structures used for the tendermint app.
type ShutterApp struct {
	Configs         []*BatchConfig
	BatchStates     map[uint64]BatchState
	DKGMap          map[uint64]*DKGInstance
	Voting          ConfigVoting
	Gobpath         string
	LastSaved       time.Time
	LastBlockHeight int64
	Identities      map[common.Address]ValidatorPubkey
	StartedVotes    map[common.Address]bool
	Validators      Powermap
	EONCounter      uint64
}

// DKGInstance manages the state of one eon key generation instance.
type DKGInstance struct {
	Config BatchConfig
	Eon    uint64

	PolyEvalMsgs       map[common.Address]map[common.Address]PolyEvalMsg
	PolyCommitmentMsgs map[common.Address]PolyCommitmentMsg
	AccusationMsgs     map[common.Address]map[common.Address]AccusationMsg
	ApologyMsgs        map[common.Address]map[common.Address]ApologyMsg

	SubmissionsClosed bool
	AccusationsClosed bool
	ApologiesClosed   bool
}

// PolyEvalMsg represents an encrypted polynomial evaluation message from one keyper to another.
type PolyEvalMsg struct {
	EncryptedEval []byte
	Eon           uint64
	Sender        common.Address
	Receiver      common.Address
}

// PolyCommitmentMsg represents a broadcasted polynomial commitment message.
type PolyCommitmentMsg struct {
	Sender common.Address
	Eon    uint64
	Gammas [][]byte
}

// AccusationMsg represents a broadcasted accusation message against a keyper.
type AccusationMsg struct {
	Eon     uint64
	Sender  common.Address
	Accused common.Address
}

// ApologyMsg represents an apology broadcasted in response to a prior accusation.
type ApologyMsg struct {
	PolyEval []byte
	Eon      uint64
	Sender   common.Address
	Accuser  common.Address
}

// EpochSKShareMsg represents a message containing an epoch secret key.
type EpochSKShareMsg struct {
	Sender       common.Address
	Eon          uint64
	Epoch        uint64
	EpochSKShare []byte
}
