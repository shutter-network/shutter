package app

/*
 This file contains code to generate tendermint events. The abcitypes.Event type allows us to send
 information about an event with a list of key value pairs. The keys and values have to be encoded
 as a utf-8 sequence of bytes.
*/

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

// MakeCheckInEvent creates a shutter.check-in event, to be raised whenever a new keyper sends
// their check in message.
func MakeCheckInEvent(sender common.Address, encryptionPublicKey *ecies.PublicKey) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.check-in",
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", sender),
			newStringPair("EncryptionPublicKey", encodePubkeyForEvent(encryptionPublicKey.ExportECDSA())),
		},
	}
}

// MakeBatchConfigEvent creates a 'shutter.batch-config' tendermint event. The given
// startBatchIndex, threshold and list of keyper addresses are encoded as attributes of the event.
func MakeBatchConfigEvent(startBatchIndex uint64, threshold uint64, keypers []common.Address) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.batch-config",
		Attributes: []abcitypes.EventAttribute{
			{
				Key:   []byte("StartBatchIndex"),
				Value: []byte(fmt.Sprintf("%d", startBatchIndex)),
				Index: true,
			},
			{
				Key:   []byte("Threshold"),
				Value: []byte(fmt.Sprintf("%d", threshold)),
			},
			{
				Key:   []byte("Keypers"),
				Value: []byte(encodeAddressesForEvent(keypers)),
			},
		},
	}
}

// MakeDecryptionSignatureEvent creates a 'shutter.decryption-signature' event.
func MakeDecryptionSignatureEvent(batchIndex uint64, sender common.Address, signature []byte) abcitypes.Event {
	encodedSignature := base64.RawURLEncoding.EncodeToString(signature)
	return abcitypes.Event{
		Type: "shutter.decryption-signature",
		Attributes: []abcitypes.EventAttribute{
			{
				Key:   []byte("BatchIndex"),
				Value: []byte(fmt.Sprintf("%d", batchIndex)),
				Index: true,
			},
			{
				Key:   []byte("Sender"),
				Value: []byte(sender.Hex()),
				Index: true,
			},
			{
				Key:   []byte("Signature"),
				Value: []byte(encodedSignature),
			},
		},
	}
}

// MakeEncryptionKeySignatureAddedEvent creates a 'shutter.encryption-key-signature-added'
// Tendermint event.
func MakeEncryptionKeySignatureAddedEvent(keyperIndex uint64, batchIndex uint64, encryptionKey []byte, signature []byte) abcitypes.Event {
	encodedKeyperIndex := []byte(fmt.Sprintf("%d", keyperIndex))
	encodedBatchIndex := []byte(fmt.Sprintf("%d", batchIndex))
	encodedKey := []byte(base64.RawURLEncoding.EncodeToString(encryptionKey))
	encodedSignature := []byte(base64.RawURLEncoding.EncodeToString(signature))
	return abcitypes.Event{
		Type: "shutter.encryption-key-signature-added",
		Attributes: []abcitypes.EventAttribute{
			{Key: []byte("KeyperIndex"), Value: encodedKeyperIndex, Index: true},
			{Key: []byte("BatchIndex"), Value: encodedBatchIndex, Index: true},
			{Key: []byte("EncryptionKey"), Value: encodedKey},
			{Key: []byte("Signature"), Value: encodedSignature},
		},
	}
}

// MakeNewDKGInstanceEvent creates a new event to be emitted whenever a new dkg process is kicked
// off.
func MakeNewDKGInstanceEvent(eon uint64, configIndex uint64) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.new-dkg-instance",
		Attributes: []abcitypes.EventAttribute{
			newUintPair("Eon", eon),
			newUintPair("ConfigIndex", configIndex),
		},
	}
}

// MakePolyEvalRegisteredEvent creates a new event to be emitted whenever a PolyEval message is
// registered.
func MakePolyEvalRegisteredEvent(msg *PolyEvalMsg) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.poly-eval-registered",
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newAddressPair("Receiver", msg.Receiver),
			newBytesPair("EncryptedEval", msg.EncryptedEval),
		},
	}
}

// MakePolyCommitmentRegisteredEvent creates a new event to be emitted whenever a PolyCommitment
// message is registered.
func MakePolyCommitmentRegisteredEvent(msg *PolyCommitmentMsg) abcitypes.Event {
	// TODO: add gammas
	return abcitypes.Event{
		Type: "shutter.poly-commitment-registered",
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
		},
	}
}

// MakeAccusationRegisteredEvent creates a new event to be emitted whenever an Accusation message
// is registered.
func MakeAccusationRegisteredEvent(msg *AccusationMsg) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.accusation-registered",
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newAddressPair("Accused", msg.Accused),
		},
	}
}

// MakeApologyRegisteredEvent creates a new event to be emitted whenever an Apology message
// is registered.
func MakeApologyRegisteredEvent(msg *ApologyMsg) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.apology-registered",
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newAddressPair("Accuser", msg.Accuser),
			newBytesPair("PolyEval", msg.PolyEval),
		},
	}
}

//
// Encoding/decoding helpers
//

func newBytesPair(key string, value []byte) abcitypes.EventAttribute {
	return abcitypes.EventAttribute{
		Key:   []byte(key),
		Value: value,
	}
}

func newStringPair(key string, value string) abcitypes.EventAttribute {
	return newBytesPair(key, []byte(value))
}

func newAddressPair(key string, value common.Address) abcitypes.EventAttribute {
	p := newStringPair(key, value.Hex())
	p.Index = true
	return p
}

func newUintPair(key string, value uint64) abcitypes.EventAttribute {
	p := newStringPair(key, strconv.FormatUint(value, 10))
	p.Index = true
	return p
}

// encodePubkeyForEvent encodes the PublicKey as a string suitable for putting it into a tendermint
// event, i.e. an utf-8 compatible string
func encodePubkeyForEvent(pubkey *ecdsa.PublicKey) string {
	return base64.RawURLEncoding.EncodeToString(crypto.FromECDSAPub(pubkey))
}

// DecodePubkeyFromEvent decodes a public key from a tendermint event (this is the reverse
// operation of encodePubkeyForEvent )
func DecodePubkeyFromEvent(s string) (*ecdsa.PublicKey, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalPubkey(data)
}

// encodePrivkeyForEvent encodes the given PrivateKey as a string suitable for putting it into a
// tendermint event
func encodePrivkeyForEvent(privkey *ecdsa.PrivateKey) string {
	return base64.RawURLEncoding.EncodeToString(crypto.FromECDSA(privkey))
}

// DecodePrivkeyFromEvent decodes a private key from a tendermint event (this is the reverse
// operation of encodePrivkeyForEvent)
func DecodePrivkeyFromEvent(s string) (*ecdsa.PrivateKey, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return crypto.ToECDSA(data)
}

// encodeAddressesForEvent encodes the given slice of Addresses as comma-separated list of addresses
func encodeAddressesForEvent(addr []common.Address) string {
	var hex []string
	for _, a := range addr {
		hex = append(hex, a.Hex())
	}
	return strings.Join(hex, ",")
}

// DecodeAddressesFromEvent reverses the encodeAddressesForEvent operation, i.e. it parses a list
// of addresses from a comma-separated string.
func DecodeAddressesFromEvent(s string) []common.Address {
	var res []common.Address
	for _, a := range strings.Split(s, ",") {
		res = append(res, common.HexToAddress(a))
	}
	return res
}
