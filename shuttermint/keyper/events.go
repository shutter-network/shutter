package keyper

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/brainbot-com/shutter/shuttermint/app"
	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

var EventType = app.EventType

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

func decodeGammasFromEvent(eventValue []byte) (crypto.Gammas, error) {
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
	return decodeGammasFromEvent(attr)
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

	publicKey, err := app.DecodePubkeyFromEvent(s)
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

// MakeCheckInEvent creates a CheckInEvent from the given tendermint event of type "shutter.check-in"
func MakeCheckInEvent(ev abcitypes.Event) (CheckInEvent, error) {
	if ev.Type != EventType.CheckIn {
		return CheckInEvent{}, fmt.Errorf("expected event type shutter.check-in, got %s", ev.Type)
	}

	sender, err := getAddressAttribute(ev, 0, "Sender")
	if err != nil {
		return CheckInEvent{}, err
	}
	publicKey, err := getECIESPublicKeyAttribute(ev, 1, "EncryptionPublicKey")
	if err != nil {
		return CheckInEvent{}, err
	}

	return CheckInEvent{
		Sender:              sender,
		EncryptionPublicKey: publicKey,
	}, nil
}

// MakeBatchConfigEvent creates a BatchConfigEvent from the given tendermint event of type
// "shutter.batch-config"
func MakeBatchConfigEvent(ev abcitypes.Event) (BatchConfigEvent, error) {
	if len(ev.Attributes) < 3 {
		return BatchConfigEvent{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("StartBatchIndex")) ||
		!bytes.Equal(ev.Attributes[1].Key, []byte("Threshold")) ||
		!bytes.Equal(ev.Attributes[2].Key, []byte("Keypers")) {
		return BatchConfigEvent{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	b, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return BatchConfigEvent{}, err
	}

	threshold, err := strconv.Atoi(string(ev.Attributes[1].Value))
	if err != nil {
		return BatchConfigEvent{}, err
	}
	keypers := app.DecodeAddressesFromEvent(string(ev.Attributes[2].Value))
	return BatchConfigEvent{uint64(b), uint64(threshold), keypers}, nil
}

// MakeDecryptionSignatureEvent creates a DecryptionSignatureEvent from the given tendermint event
// of type "shutter.decryption-signature".
func MakeDecryptionSignatureEvent(ev abcitypes.Event) (DecryptionSignatureEvent, error) {
	if len(ev.Attributes) < 3 {
		return DecryptionSignatureEvent{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("BatchIndex")) ||
		!bytes.Equal(ev.Attributes[1].Key, []byte("Sender")) ||
		!bytes.Equal(ev.Attributes[2].Key, []byte("Signature")) {
		return DecryptionSignatureEvent{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	batchIndex, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return DecryptionSignatureEvent{}, err
	}

	encodedSender := string(ev.Attributes[1].Value)
	sender := common.HexToAddress(encodedSender)
	if sender.Hex() != encodedSender {
		return DecryptionSignatureEvent{}, fmt.Errorf("invalid sender address %s", encodedSender)
	}

	signature, err := base64.RawURLEncoding.DecodeString(string(ev.Attributes[2].Value))
	if err != nil {
		return DecryptionSignatureEvent{}, err
	}

	return DecryptionSignatureEvent{
		BatchIndex: uint64(batchIndex),
		Sender:     sender,
		Signature:  signature,
	}, nil
}

// MakeEonStartedEvent creates a EonStartedEvent from the given tendermint event of type
// "shutter.eon-started".
func MakeEonStartedEvent(ev abcitypes.Event) (EonStartedEvent, error) {
	if ev.Type != EventType.EonStarted {
		return EonStartedEvent{}, fmt.Errorf("expected event type %s, got %s", EventType.EonStarted, ev.Type)
	}

	eon, err := getUint64Attribute(ev, 0, "Eon")
	if err != nil {
		return EonStartedEvent{}, err
	}
	batchIndex, err := getUint64Attribute(ev, 1, "BatchIndex")
	if err != nil {
		return EonStartedEvent{}, err
	}

	return EonStartedEvent{
		Eon:        eon,
		BatchIndex: batchIndex,
	}, nil
}

func MakePolyCommitmentRegisteredEvent(ev abcitypes.Event) (PolyCommitmentRegisteredEvent, error) {
	res := PolyCommitmentRegisteredEvent{}
	if ev.Type != EventType.PolyCommitment {
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

func MakePolyEvalRegisteredEvent(ev abcitypes.Event) (PolyEvalRegisteredEvent, error) {
	res := PolyEvalRegisteredEvent{}
	if ev.Type != EventType.PolyEval {
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

// MakeEvent creates an Event from the given tendermint event.
func MakeEvent(ev abcitypes.Event) (IEvent, error) {
	switch ev.Type {
	case EventType.CheckIn:
		return MakeCheckInEvent(ev)
	case EventType.BatchConfig:
		return MakeBatchConfigEvent(ev)
	case EventType.DecryptionSignature:
		return MakeDecryptionSignatureEvent(ev)
	case EventType.EonStarted:
		return MakeEonStartedEvent(ev)
	case EventType.PolyCommitment:
		return MakePolyCommitmentRegisteredEvent(ev)
	case EventType.PolyEval:
		return MakePolyEvalRegisteredEvent(ev)
	default:
		return nil, fmt.Errorf("cannot make event from type %s", ev.Type)
	}
}
