// Package shutterevents contains types to represent deserialized shuttermint/tendermint events
package shutterevents

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/brainbot-com/shutter/shuttermint/app/evtype"
	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

// Accusation represents a broadcasted accusation message against one or more keypers.
type Accusation struct {
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

func makeAccusation(ev abcitypes.Event) (Accusation, error) {
	if ev.Type != evtype.Accusation {
		return Accusation{}, fmt.Errorf("expected event type %s, got %s", evtype.Accusation, ev.Type)
	}
	sender, err := getAddressAttribute(ev, 0, "Sender")
	if err != nil {
		return Accusation{}, err
	}

	eon, err := getUint64Attribute(ev, 1, "Eon")
	if err != nil {
		return Accusation{}, err
	}

	accused, err := decodeAddresses(string(ev.Attributes[2].GetValue()))
	if err != nil {
		return Accusation{}, err
	}

	return Accusation{
		Sender:  sender,
		Eon:     eon,
		Accused: accused,
	}, nil
}

// Apology represents an apology broadcasted in response to a prior accusation.
type Apology struct {
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

func makeApology(ev abcitypes.Event) (Apology, error) {
	if ev.Type != evtype.Apology {
		return Apology{}, fmt.Errorf("expected event type %s, got %s", evtype.Apology, ev.Type)
	}
	sender, err := getAddressAttribute(ev, 0, "Sender")
	if err != nil {
		return Apology{}, err
	}

	eon, err := getUint64Attribute(ev, 1, "Eon")
	if err != nil {
		return Apology{}, err
	}

	accusers, err := decodeAddresses(string(ev.Attributes[2].GetValue()))
	if err != nil {
		return Apology{}, err
	}
	var polyEval []*big.Int // XXX
	polyEvalBytes, err := decodeByteSequenceFromEvent(string(ev.Attributes[3].GetValue()))
	if err != nil {
		return Apology{}, err
	}
	for _, b := range polyEvalBytes {
		e := new(big.Int)
		e.SetBytes(b)
		polyEval = append(polyEval, e)
	}
	return Apology{
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
	StartBatchIndex       uint64
	Keypers               []common.Address
	Threshold             uint64
	ConfigContractAddress common.Address
	ConfigIndex           uint64
	Started               bool
	ValidatorsUpdated     bool
}

func (msg BatchConfig) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.BatchConfig,
		Attributes: []abcitypes.EventAttribute{
			{
				Key:   []byte("StartBatchIndex"),
				Value: []byte(fmt.Sprintf("%d", msg.StartBatchIndex)),
				Index: true,
			},
			{
				Key:   []byte("Threshold"),
				Value: []byte(fmt.Sprintf("%d", msg.Threshold)),
			},
			{
				Key:   []byte("Keypers"),
				Value: []byte(encodeAddressesForEvent(msg.Keypers)),
			},
			{
				Key:   []byte("ConfigIndex"),
				Value: []byte(fmt.Sprintf("%d", msg.ConfigIndex)),
			},
		},
	}
}

// makeBatchConfig creates a BatchConfigEvent from the given tendermint event of type
// "shutter.batch-config"
func makeBatchConfig(ev abcitypes.Event) (BatchConfig, error) {
	if len(ev.Attributes) < 4 {
		return BatchConfig{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("StartBatchIndex")) ||
		!bytes.Equal(ev.Attributes[1].Key, []byte("Threshold")) ||
		!bytes.Equal(ev.Attributes[2].Key, []byte("Keypers")) ||
		!bytes.Equal(ev.Attributes[3].Key, []byte("ConfigIndex")) {
		return BatchConfig{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	b, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return BatchConfig{}, err
	}

	threshold, err := strconv.Atoi(string(ev.Attributes[1].Value))
	if err != nil {
		return BatchConfig{}, err
	}
	keypers, err := decodeAddresses(string(ev.Attributes[2].Value))
	if err != nil {
		return BatchConfig{}, err
	}

	configIndex, err := strconv.ParseUint(string(ev.Attributes[3].Value), 10, 64)
	if err != nil {
		return BatchConfig{}, err
	}
	return BatchConfig{
		StartBatchIndex: uint64(b),
		Threshold:       uint64(threshold),
		Keypers:         keypers,
		ConfigIndex:     configIndex,
	}, nil
}

// CheckIn is emitted by shuttermint when a keyper sends their check in message.
type CheckIn struct {
	Sender              common.Address
	EncryptionPublicKey *ecies.PublicKey
}

func (msg CheckIn) MakeABCIEvent() abcitypes.Event {
	return abcitypes.Event{
		Type: evtype.CheckIn,
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newStringPair("EncryptionPublicKey", encodePubkeyForEvent(msg.EncryptionPublicKey.ExportECDSA())),
		},
	}
}

// makeCheckIn creates a CheckInEvent from the given tendermint event of type "shutter.check-in"
func makeCheckIn(ev abcitypes.Event) (CheckIn, error) {
	if ev.Type != evtype.CheckIn {
		return CheckIn{}, fmt.Errorf("expected event type shutter.check-in, got %s", ev.Type)
	}

	sender, err := getAddressAttribute(ev, 0, "Sender")
	if err != nil {
		return CheckIn{}, err
	}
	publicKey, err := getECIESPublicKeyAttribute(ev, 1, "EncryptionPublicKey")
	if err != nil {
		return CheckIn{}, err
	}

	return CheckIn{
		Sender:              sender,
		EncryptionPublicKey: publicKey,
	}, nil
}

// DecryptionSignature is generated by shuttermint when a keyper sends a decryption
// signature.
type DecryptionSignature struct {
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
func makeDecryptionSignature(ev abcitypes.Event) (DecryptionSignature, error) {
	if len(ev.Attributes) < 3 {
		return DecryptionSignature{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("BatchIndex")) ||
		!bytes.Equal(ev.Attributes[1].Key, []byte("Sender")) ||
		!bytes.Equal(ev.Attributes[2].Key, []byte("Signature")) {
		return DecryptionSignature{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	batchIndex, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return DecryptionSignature{}, err
	}

	encodedSender := string(ev.Attributes[1].Value)
	sender := common.HexToAddress(encodedSender)
	if sender.Hex() != encodedSender {
		return DecryptionSignature{}, fmt.Errorf("invalid sender address %s", encodedSender)
	}

	signature, err := base64.RawURLEncoding.DecodeString(string(ev.Attributes[2].Value))
	if err != nil {
		return DecryptionSignature{}, err
	}

	return DecryptionSignature{
		BatchIndex: uint64(batchIndex),
		Sender:     sender,
		Signature:  signature,
	}, nil
}

// EonStarted is generated by shuttermint when a new eon is started.
type EonStarted struct {
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
	Eon    uint64
	Sender common.Address
	Gammas *crypto.Gammas
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

func makePolyCommitment(ev abcitypes.Event) (PolyCommitment, error) {
	res := PolyCommitment{}
	if ev.Type != evtype.PolyCommitment {
		return res, fmt.Errorf("expected event type shutter.poly-commitment-registered, got %s", ev.Type)
	}

	sender, err := getAddressAttribute(ev, 0, "Sender")
	if err != nil {
		return res, err
	}
	res.Sender = sender

	eon, err := getUint64Attribute(ev, 1, "Eon")
	if err != nil {
		return res, err
	}
	res.Eon = eon

	gammas, err := getGammasAttribute(ev, 2, "Gammas")
	if err != nil {
		return res, err
	}
	res.Gammas = &gammas

	return res, nil
}

// PolyEval represents an encrypted polynomial evaluation message from one keyper to another.
type PolyEval struct {
	Eon            uint64
	Sender         common.Address
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

func makePolyEval(ev abcitypes.Event) (PolyEval, error) {
	res := PolyEval{}
	if ev.Type != evtype.PolyEval {
		return res, fmt.Errorf("expected event type shutter.poly-eval-registered, got %s", ev.Type)
	}

	sender, err := getAddressAttribute(ev, 0, "Sender")
	if err != nil {
		return res, err
	}
	res.Sender = sender

	eon, err := getUint64Attribute(ev, 1, "Eon")
	if err != nil {
		return res, err
	}
	res.Eon = eon

	return res, nil
}

// IEvent is an interface for the event types declared above
type IEvent interface {
	MakeABCIEvent() abcitypes.Event
}

func getBytesAttribute(ev abcitypes.Event, index int, key string) ([]byte, error) {
	if len(ev.Attributes) <= index {
		return []byte{}, fmt.Errorf("event does not have enough attributes")
	}
	attr := ev.Attributes[index]
	if string(attr.Key) != key {
		return []byte{}, fmt.Errorf("expected attribute key %s at index %d, got %s", key, index, attr.Key)
	}
	return attr.Value, nil
}

func getUint64Attribute(ev abcitypes.Event, index int, name string) (uint64, error) {
	attr, err := getBytesAttribute(ev, index, name)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(string(attr))
	if err != nil {
		return 0, fmt.Errorf("failed to parse event: %w", err)
	}
	return uint64(v), nil
}

// decodeAddresses reverses the encodeAddressesForEvent operation, i.e. it parses a list
// of addresses from a comma-separated string.
func decodeAddresses(s string) ([]common.Address, error) {
	var res []common.Address
	for _, a := range strings.Split(s, ",") {
		if !common.IsHexAddress(a) {
			return nil, fmt.Errorf("malformed address: %q", s)
		}

		res = append(res, common.HexToAddress(a))
	}
	return res, nil
}

// decodeByteSequenceFromEvent parses a list of hex encoded, comma-separated byte slices.
func decodeByteSequenceFromEvent(s string) ([][]byte, error) {
	var res [][]byte
	for _, v := range strings.Split(s, ",") {
		bs, err := hexutil.Decode(v)
		if err != nil {
			return [][]byte{}, err
		}
		res = append(res, bs)
	}
	return res, nil
}

// DecodePubkey decodes a public key from a tendermint event (this is the reverse
// operation of app.encodePubkeyForEvent )
// XXX the is only needed by a shuttermint app test, should eventually end up private like all
// other methods
func DecodePubkey(s string) (*ecdsa.PublicKey, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return ethcrypto.UnmarshalPubkey(data)
}

func decodeGammas(eventValue []byte) (crypto.Gammas, error) {
	parts := strings.Split(string(eventValue), ",")
	var res crypto.Gammas
	for _, p := range parts {
		marshaledG2, err := hex.DecodeString(p)
		if err != nil {
			return crypto.Gammas{}, err
		}
		g := new(bn256.G2)
		_, err = g.Unmarshal(marshaledG2)
		if err != nil {
			return crypto.Gammas{}, err
		}
		res = append(res, g)
	}
	return res, nil
}

func getGammasAttribute(ev abcitypes.Event, index int, name string) (crypto.Gammas, error) {
	attr, err := getBytesAttribute(ev, index, name)
	if err != nil {
		return crypto.Gammas{}, err
	}
	return decodeGammas(attr)
}

func getStringAttribute(ev abcitypes.Event, index int, key string) (string, error) {
	b, err := getBytesAttribute(ev, index, key)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getAddressAttribute(ev abcitypes.Event, index int, key string) (common.Address, error) {
	s, err := getStringAttribute(ev, index, key)
	if err != nil {
		return common.Address{}, err
	}
	a := common.HexToAddress(s)
	if a.Hex() != s {
		return common.Address{}, fmt.Errorf("invalid address %s", s)
	}
	return a, nil
}

func getPublicKeyAttribute(ev abcitypes.Event, index int, key string) (*ecdsa.PublicKey, error) {
	s, err := getStringAttribute(ev, index, key)
	if err != nil {
		return nil, err
	}

	publicKey, err := DecodePubkey(s)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func getECIESPublicKeyAttribute(ev abcitypes.Event, index int, key string) (*ecies.PublicKey, error) {
	publicKeyECDSA, err := getPublicKeyAttribute(ev, index, key)
	if err != nil {
		return nil, err
	}
	return ecies.ImportECDSAPublic(publicKeyECDSA), nil
}

// makeEonStarted creates a EonStartedEvent from the given tendermint event of type
// "shutter.eon-started".
func makeEonStarted(ev abcitypes.Event) (EonStarted, error) {
	if ev.Type != evtype.EonStarted {
		return EonStarted{}, fmt.Errorf("expected event type %s, got %s", evtype.EonStarted, ev.Type)
	}

	eon, err := getUint64Attribute(ev, 0, "Eon")
	if err != nil {
		return EonStarted{}, err
	}
	batchIndex, err := getUint64Attribute(ev, 1, "BatchIndex")
	if err != nil {
		return EonStarted{}, err
	}

	return EonStarted{
		Eon:        eon,
		BatchIndex: batchIndex,
	}, nil
}

// MakeEvent creates an Event from the given tendermint event.
func MakeEvent(ev abcitypes.Event) (IEvent, error) {
	switch ev.Type {
	case evtype.CheckIn:
		return makeCheckIn(ev)
	case evtype.BatchConfig:
		return makeBatchConfig(ev)
	case evtype.DecryptionSignature:
		return makeDecryptionSignature(ev)
	case evtype.EonStarted:
		return makeEonStarted(ev)
	case evtype.PolyCommitment:
		return makePolyCommitment(ev)
	case evtype.PolyEval:
		return makePolyEval(ev)
	case evtype.Accusation:
		return makeAccusation(ev)
	case evtype.Apology:
		return makeApology(ev)

	default:
		return nil, fmt.Errorf("cannot make event from type %s", ev.Type)
	}
}
