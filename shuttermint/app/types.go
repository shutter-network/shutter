package app

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shuttermint/keyper/shutterevents"
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

// GetKeypers returns the keypers defined in the GenesisAppState.
func (appState *GenesisAppState) GetKeypers() []common.Address {
	var res []common.Address
	for _, k := range appState.Keypers {
		res = append(res, k.Address())
	}
	return res
}

// Voting is a struct storing votes for arbitrary indices.
type Voting struct {
	Votes map[common.Address]int
}

// ConfigVoting is used to let the keypers vote on new BatchConfigs to be added
// Each keyper can vote exactly once.
type ConfigVoting struct {
	Voting
	Candidates []BatchConfig
}

// EonStartVoting is used to vote on the batch index at which the next eon should be started.
type EonStartVoting struct {
	Voting
	Candidates []uint64
}

// DecryptionSignature stores the decryption key signature created by one of the keypers.
type DecryptionSignature struct {
	Sender    common.Address
	Signature []byte
}

// BatchState is used to manage the key generation process for a certain batch.
type BatchState struct {
	BatchIndex           uint64
	Config               *BatchConfig
	DecryptionSignatures []DecryptionSignature
}

// ValidatorPubkey holds the raw 32 byte ed25519 public key to be used as tendermint validator key
// We use this is a map key, so don't use a byte slice.
type ValidatorPubkey struct {
	Ed25519pubkey string
}

func (vp ValidatorPubkey) String() string {
	return fmt.Sprintf("ed25519:%s", hex.EncodeToString([]byte(vp.Ed25519pubkey)))
}

// Powermap maps a ValidatorPubkey to the validators voting power.
type Powermap map[ValidatorPubkey]int64

// NewValidatorPubkey creates a new ValidatorPubkey from a 32 byte ed25519 raw pubkey. See
// https://docs.tendermint.com/master/spec/abci/apps.html#validator-updates for more information
func NewValidatorPubkey(pubkey []byte) (ValidatorPubkey, error) {
	if len(pubkey) != ed25519.PublicKeySize {
		return ValidatorPubkey{}, errors.Errorf("pubkey must be 32 bytes")
	}
	return ValidatorPubkey{Ed25519pubkey: string(pubkey)}, nil
}

// ShutterApp holds our data structures used for the tendermint app.
type ShutterApp struct {
	Configs         []*BatchConfig
	BatchStates     map[uint64]BatchState
	DKGMap          map[uint64]*DKGInstance
	ConfigVoting    ConfigVoting
	EonStartVotings map[uint64]*EonStartVoting
	Gobpath         string
	LastSaved       time.Time
	LastBlockHeight int64
	Identities      map[common.Address]ValidatorPubkey
	StartedVotes    map[common.Address]struct{}
	Validators      Powermap
	EONCounter      uint64
	DevMode         bool
	CheckTxState    *CheckTxState
	NonceTracker    *NonceTracker
	ChainID         string
}

// CheckTxState is a part of the state used by CheckTx calls that is reset at every commit.
type CheckTxState struct {
	Members      map[common.Address]bool
	TxCounts     map[common.Address]int
	NonceTracker *NonceTracker
}

// NonceTracker tracks which nonces have been used and which have not.
type NonceTracker struct {
	RandomNonces map[common.Address]map[uint64]bool
}

type SenderReceiverPair struct {
	Sender, Receiver common.Address
}

// DKGInstance manages the state of one eon key generation instance.
type DKGInstance struct {
	Config BatchConfig
	Eon    uint64

	PolyEvalsSeen       map[SenderReceiverPair]struct{}
	PolyCommitmentsSeen map[common.Address]struct{}
	AccusationsSeen     map[common.Address]struct{}
	ApologiesSeen       map[common.Address]struct{}
}

type (
	Accusation          = shutterevents.Accusation
	Apology             = shutterevents.Apology
	BatchConfig         = shutterevents.BatchConfig
	PolyCommitment      = shutterevents.PolyCommitment
	PolyEval            = shutterevents.PolyEval
	EpochSecretKeyShare = shutterevents.EpochSecretKeyShare
)
