// Package shutterevents contains types to represent deserialized shuttermint/tendermint events
package shutterevents

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/shutter-network/shutter/shlib/shcrypto"
	"github.com/shutter-network/shutter/shuttermint/keyper/shutterevents/evtype"
)

/* All of the event types defined here have a "Height" field, that is *not* being serialized when
   calling MakeABCIEvent.  We need this field in the keyper. It's set via passing the height
   argument to MakeEvent.
*/

// Accusation represents a broadcasted accusation message against one or more keypers.
type Accusation struct {
	Height  int64
	Eon     uint64
	Sender  common.Address
	Accused []common.Address
}

func (acc Accusation) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.Accusation,
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", acc.Sender),
			newUintPair("Eon", acc.Eon),
			newAddressesPair("Accused", acc.Accused),
		},
	}
}

func expectAttributes(ev abcitypes.Event, names ...string) error {
	if len(ev.Attributes) < len(names) {
		return errors.Errorf("expected at least %d attributes", len(names))
	}

	for i, n := range names {
		if string(ev.Attributes[i].Key) != n {
			return errors.Errorf(
				"bad attribute, parsing event %s: expected %s, got %s at position %d",
				ev.Type,
				n,
				string(ev.Attributes[i].Key),
				i,
			)
		}
	}
	return nil
}

func makeAccusation(ev abcitypes.Event, height int64) (*Accusation, error) {
	err := expectAttributes(ev, "Sender", "Eon", "Accused")
	if err != nil {
		return nil, err
	}

	sender, err := decodeAddress(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	eon, err := decodeUint64(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	accused, err := decodeAddresses(ev.Attributes[2].Value)
	if err != nil {
		return nil, err
	}

	return &Accusation{
		Height:  height,
		Sender:  sender,
		Eon:     eon,
		Accused: accused,
	}, nil
}

// Apology represents an apology broadcasted in response to a prior accusation.
type Apology struct {
	Height   int64
	Eon      uint64
	Sender   common.Address
	Accusers []common.Address
	PolyEval []*big.Int
}

func (msg Apology) MakeABCIEvent() abcitypes.Event {
	var polyEvalBytes [][]byte
	for _, e := range msg.PolyEval {
		polyEvalBytes = append(polyEvalBytes, e.Bytes())
	}
	return abcitypes.Event{
		Type: evtype.Apology,
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newAddressesPair("Accusers", msg.Accusers),
			newByteSequencePair("PolyEvals", polyEvalBytes),
		},
	}
}

func makeApology(ev abcitypes.Event, height int64) (*Apology, error) {
	err := expectAttributes(ev, "Sender", "Eon", "Accusers", "PolyEvals")
	if err != nil {
		return nil, err
	}

	sender, err := decodeAddress(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	eon, err := decodeUint64(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	accusers, err := decodeAddresses(ev.Attributes[2].Value)
	if err != nil {
		return nil, err
	}
	var polyEval []*big.Int
	polyEvalBytes, err := decodeByteSequence(ev.Attributes[3].Value)
	if err != nil {
		return nil, err
	}
	for _, b := range polyEvalBytes {
		e := new(big.Int)
		e.SetBytes(b)
		polyEval = append(polyEval, e)
	}
	return &Apology{
		Height:   height,
		Sender:   sender,
		Eon:      eon,
		Accusers: accusers,
		PolyEval: polyEval,
	}, nil
}

//

// BatchConfig is the configuration we use for a consecutive sequence of batches.  This should be
// synchronized with the list of BatchConfig structures stored in the ConfigContract deployed on
// the main chain. The keypers use the corresponding shmsg.BatchConfig message to vote on new
// configs. This struct is sent as an event, when a new batch config has enough votes.
type BatchConfig struct {
	Height                int64
	Keypers               []common.Address
	StartBatchIndex       uint64
	Threshold             uint64
	ConfigIndex           uint64
	ConfigContractAddress common.Address
	Started               bool
	ValidatorsUpdated     bool
}

func (bc BatchConfig) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.BatchConfig,
		Attributes: []abcitypes.EventAttribute{
			{
				Key:   []byte("StartBatchIndex"),
				Value: []byte(fmt.Sprintf("%d", bc.StartBatchIndex)),
				Index: true,
			},
			{
				Key:   []byte("Threshold"),
				Value: []byte(fmt.Sprintf("%d", bc.Threshold)),
			},
			{
				Key:   []byte("Keypers"),
				Value: encodeAddresses(bc.Keypers),
			},
			{
				Key:   []byte("ConfigIndex"),
				Value: []byte(fmt.Sprintf("%d", bc.ConfigIndex)),
			},
		},
	}
}

// makeBatchConfig creates a BatchConfigEvent from the given tendermint event of type
// "shutter.batch-config".
func makeBatchConfig(ev abcitypes.Event, height int64) (*BatchConfig, error) {
	err := expectAttributes(ev, "StartBatchIndex", "Threshold", "Keypers", "ConfigIndex")
	if err != nil {
		return nil, err
	}

	startBatchIndex, err := decodeUint64(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	threshold, err := decodeUint64(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}
	keypers, err := decodeAddresses(ev.Attributes[2].Value)
	if err != nil {
		return nil, err
	}

	configIndex, err := decodeUint64(ev.Attributes[3].Value)
	if err != nil {
		return nil, err
	}
	return &BatchConfig{
		Height:          height,
		StartBatchIndex: startBatchIndex,
		Threshold:       threshold,
		Keypers:         keypers,
		ConfigIndex:     configIndex,
	}, nil
}

// CheckIn is emitted by shuttermint when a keyper sends their check in message.
type CheckIn struct {
	Height              int64
	Sender              common.Address
	EncryptionPublicKey *ecies.PublicKey
}

func (msg CheckIn) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.CheckIn,
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			{
				Key:   []byte("EncryptionPublicKey"),
				Value: encodeECIESPublicKey(msg.EncryptionPublicKey),
			},
		},
	}
}

// makeCheckIn creates a CheckInEvent from the given tendermint event of type "shutter.check-in".
func makeCheckIn(ev abcitypes.Event, height int64) (*CheckIn, error) {
	err := expectAttributes(ev, "Sender", "EncryptionPublicKey")
	if err != nil {
		return nil, err
	}
	sender, err := decodeAddress(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	publicKey, err := decodeECIESPublicKey(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	return &CheckIn{
		Sender:              sender,
		EncryptionPublicKey: publicKey,
		Height:              height,
	}, nil
}

// DecryptionSignature is generated by shuttermint when a keyper sends a decryption
// signature.
type DecryptionSignature struct {
	Height     int64
	BatchIndex uint64
	Sender     common.Address
	Signature  []byte
}

func (msg DecryptionSignature) MakeABCIEvent() abcitypes.Event {
	encodedSignature := base64.RawURLEncoding.EncodeToString(msg.Signature)
	return abcitypes.Event{
		Type: evtype.DecryptionSignature,
		Attributes: []abcitypes.EventAttribute{
			{
				Key:   []byte("BatchIndex"),
				Value: []byte(fmt.Sprintf("%d", msg.BatchIndex)),
				Index: true,
			},
			{
				Key:   []byte("Sender"),
				Value: []byte(msg.Sender.Hex()),
				Index: true,
			},
			{
				Key:   []byte("Signature"),
				Value: []byte(encodedSignature),
			},
		},
	}
}

// makeDecryptionSignature creates a DecryptionSignatureEvent from the given tendermint event
// of type "shutter.decryption-signature".
func makeDecryptionSignature(ev abcitypes.Event, height int64) (*DecryptionSignature, error) {
	err := expectAttributes(ev, "BatchIndex", "Sender", "Signature")
	if err != nil {
		return nil, err
	}

	batchIndex, err := decodeUint64(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	sender, err := decodeAddress(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	signature, err := base64.RawURLEncoding.DecodeString(string(ev.Attributes[2].Value))
	if err != nil {
		return nil, err
	}

	return &DecryptionSignature{
		Height:     height,
		BatchIndex: batchIndex,
		Sender:     sender,
		Signature:  signature,
	}, nil
}

// EonStarted is generated by shuttermint when a new eon is started.  The batch index identifies
// the first batch that belongs to that eon.
type EonStarted struct {
	Height     int64
	Eon        uint64
	BatchIndex uint64
}

func (msg EonStarted) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.EonStarted,
		Attributes: []abcitypes.EventAttribute{
			newUintPair("Eon", msg.Eon),
			newUintPair("BatchIndex", msg.BatchIndex),
		},
	}
}

// PolyCommitment represents a broadcasted polynomial commitment message.
type PolyCommitment struct {
	Height int64
	Eon    uint64
	Sender common.Address
	Gammas *shcrypto.Gammas
}

func (msg PolyCommitment) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.PolyCommitment,
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newGammas("Gammas", msg.Gammas),
		},
	}
}

func makePolyCommitment(ev abcitypes.Event, height int64) (*PolyCommitment, error) {
	err := expectAttributes(ev, "Sender", "Eon", "Gammas")
	if err != nil {
		return nil, err
	}
	sender, err := decodeAddress(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	eon, err := decodeUint64(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	gammas, err := decodeGammas(ev.Attributes[2].Value)
	if err != nil {
		return nil, err
	}

	return &PolyCommitment{
		Height: height,
		Sender: sender,
		Eon:    eon,
		Gammas: &gammas,
	}, nil
}

// PolyEval represents an encrypted polynomial evaluation message from one keyper to another.
type PolyEval struct {
	Height         int64
	Sender         common.Address
	Eon            uint64
	Receivers      []common.Address
	EncryptedEvals [][]byte
}

func (msg PolyEval) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.PolyEval,
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newAddressesPair("Receivers", msg.Receivers),
			newByteSequencePair("EncryptedEvals", msg.EncryptedEvals),
		},
	}
}

func makePolyEval(ev abcitypes.Event, height int64) (*PolyEval, error) {
	err := expectAttributes(ev, "Sender", "Eon", "Receivers", "EncryptedEvals")
	if err != nil {
		return nil, err
	}

	sender, err := decodeAddress(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	eon, err := decodeUint64(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	receivers, err := decodeAddresses(ev.Attributes[2].GetValue())
	if err != nil {
		return nil, err
	}

	encryptedEvals, err := decodeByteSequence(ev.Attributes[3].Value)
	if err != nil {
		return nil, err
	}

	return &PolyEval{
		Height:         height,
		Eon:            eon,
		Sender:         sender,
		Receivers:      receivers,
		EncryptedEvals: encryptedEvals,
	}, nil
}

// EpochSecretKeyShare represents a message containing an epoch secret key.
type EpochSecretKeyShare struct {
	Height int64
	Sender common.Address
	Eon    uint64
	Epoch  uint64
	Share  *shcrypto.EpochSecretKeyShare
}

func (msg EpochSecretKeyShare) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.EpochSecretKeyShare,
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newUintPair("Epoch", msg.Epoch),
			newEpochSecretKeyShare("Share", msg.Share),
		},
	}
}

func makeEpochSecretKeyShare(ev abcitypes.Event, height int64) (*EpochSecretKeyShare, error) {
	err := expectAttributes(ev, "Sender", "Eon", "Epoch", "Share")
	if err != nil {
		return nil, err
	}

	sender, err := decodeAddress(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}

	eon, err := decodeUint64(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	epoch, err := decodeUint64(ev.Attributes[2].Value)
	if err != nil {
		return nil, err
	}
	share, err := decodeEpochSecretKeyShare(ev.Attributes[3].Value)
	if err != nil {
		return nil, err
	}

	return &EpochSecretKeyShare{
		Height: height,
		Sender: sender,
		Eon:    eon,
		Epoch:  epoch,
		Share:  share,
	}, nil
}

// IEvent is an interface for the event types declared above.
type IEvent interface {
	MakeABCIEvent() abcitypes.Event
}

// makeEonStarted creates a EonStartedEvent from the given tendermint event of type
// "shutter.eon-started".
func makeEonStarted(ev abcitypes.Event, height int64) (*EonStarted, error) {
	err := expectAttributes(ev, "Eon", "BatchIndex")
	if err != nil {
		return nil, err
	}

	eon, err := decodeUint64(ev.Attributes[0].Value)
	if err != nil {
		return nil, err
	}
	batchIndex, err := decodeUint64(ev.Attributes[1].Value)
	if err != nil {
		return nil, err
	}

	return &EonStarted{
		Height:     height,
		Eon:        eon,
		BatchIndex: batchIndex,
	}, nil
}

// MakeEvent creates an Event from the given tendermint event.
func MakeEvent(ev abcitypes.Event, height int64) (IEvent, error) {
	switch ev.Type {
	case evtype.CheckIn:
		return makeCheckIn(ev, height)
	case evtype.BatchConfig:
		return makeBatchConfig(ev, height)
	case evtype.DecryptionSignature:
		return makeDecryptionSignature(ev, height)
	case evtype.EonStarted:
		return makeEonStarted(ev, height)
	case evtype.PolyCommitment:
		return makePolyCommitment(ev, height)
	case evtype.PolyEval:
		return makePolyEval(ev, height)
	case evtype.Accusation:
		return makeAccusation(ev, height)
	case evtype.Apology:
		return makeApology(ev, height)
	case evtype.EpochSecretKeyShare:
		return makeEpochSecretKeyShare(ev, height)
	default:
		return nil, errors.Errorf("cannot make event from type %s", ev.Type)
	}
}
