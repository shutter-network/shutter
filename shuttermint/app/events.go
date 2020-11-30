package app

/*
 This file contains code to generate tendermint events. The abcitypes.Event type allows us to send
 information about an event with a list of key value pairs. The keys and values have to be encoded
 as a utf-8 sequence of bytes.
*/

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
			newAddressesPair("Receivers", msg.Receivers),
			newByteSequencePair("EncryptedEvals", msg.EncryptedEvals),
		},
	}
}

// MakePolyCommitmentRegisteredEvent creates a new event to be emitted whenever a PolyCommitment
// message is registered.
func MakePolyCommitmentRegisteredEvent(msg *PolyCommitmentMsg) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.poly-commitment-registered",
		Attributes: []abcitypes.EventAttribute{
			newAddressPair("Sender", msg.Sender),
			newUintPair("Eon", msg.Eon),
			newGammas("Gammas", msg.Gammas),
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
			newAddressesPair("Accused", msg.Accused),
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
			newAddressesPair("Accusers", msg.Accusers),
			newByteSequencePair("PolyEvals", msg.PolyEvals),
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

func newAddressesPair(key string, value []common.Address) abcitypes.EventAttribute {
	s := encodeAddressesForEvent(value)
	return newStringPair(key, s)
}

func newByteSequencePair(key string, value [][]byte) abcitypes.EventAttribute {
	return newStringPair(key, encodeByteSequenceForEvent(value))
}

func newUintPair(key string, value uint64) abcitypes.EventAttribute {
	p := newStringPair(key, strconv.FormatUint(value, 10))
	p.Index = true
	return p
}

func newGammas(key string, gammas [][]byte) abcitypes.EventAttribute {
	var encoded []string
	for _, i := range gammas {
		encoded = append(encoded, hex.EncodeToString(i))
	}
	return abcitypes.EventAttribute{
		Key:   []byte(key),
		Value: []byte(strings.Join(encoded, ",")),
	}
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

func encodeByteSequenceForEvent(v [][]byte) string {
	var hex []string
	for _, a := range v {
		hex = append(hex, hexutil.Encode(a))
	}
	return strings.Join(hex, ",")
}

// DecodeByteSequenceFromEvent parses a list of hex encoded, comma-separated byte slices.
func DecodeByteSequenceFromEvent(s string) ([][]byte, error) {
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
