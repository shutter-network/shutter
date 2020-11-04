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
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// MakePubkeyGeneratedEvent creates a 'shutter.pubkey-generated' tendermint event.  The given
// BatchIndex and PublicKey are encoded as attributes of the event.
func MakePubkeyGeneratedEvent(batchIndex uint64, pubkey *ecdsa.PublicKey) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.pubkey-generated",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("Pubkey"), Value: []byte(encodePubkeyForEvent(pubkey))},
		},
	}
}

// MakePrivkeyGeneratedEvent creates a 'shutter.privkey-generated' tendermint event. The given
// BatchIndex and PrivateKey are encoded as attributes of the event
func MakePrivkeyGeneratedEvent(batchIndex uint64, privkey *ecdsa.PrivateKey) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.privkey-generated",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("Privkey"), Value: []byte(encodePrivkeyForEvent(privkey))},
		},
	}
}

// MakeBatchConfigEvent creates a 'shutter.batch-config' tendermint event. The given
// startBatchIndex, threshold and list of keyper addresses are encoded as attributes of the event.
func MakeBatchConfigEvent(startBatchIndex uint64, threshold uint64, keypers []common.Address) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.batch-config",
		Attributes: []kv.Pair{
			{Key: []byte("StartBatchIndex"), Value: []byte(fmt.Sprintf("%d", startBatchIndex))},
			{Key: []byte("Threshold"), Value: []byte(fmt.Sprintf("%d", threshold))},
			{Key: []byte("Keypers"), Value: []byte(encodeAddressesForEvent(keypers))},
		},
	}
}

// MakeDecryptionSignatureEvent creates a 'shutter.decryption-signature' event.
func MakeDecryptionSignatureEvent(batchIndex uint64, sender common.Address, signature []byte) abcitypes.Event {
	encodedSignature := base64.RawURLEncoding.EncodeToString(signature)
	return abcitypes.Event{
		Type: "shutter.decryption-signature",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("Sender"), Value: []byte(sender.Hex())},
			{Key: []byte("Signature"), Value: []byte(encodedSignature)},
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
		Attributes: []kv.Pair{
			{Key: []byte("KeyperIndex"), Value: encodedKeyperIndex},
			{Key: []byte("BatchIndex"), Value: encodedBatchIndex},
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
		Attributes: []kv.Pair{
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
		Attributes: []kv.Pair{
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
		Attributes: []kv.Pair{
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
		Attributes: []kv.Pair{
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
		Attributes: []kv.Pair{
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

func newBytesPair(key string, value []byte) kv.Pair {
	return kv.Pair{
		Key:   []byte(key),
		Value: value,
	}
}

func newStringPair(key string, value string) kv.Pair {
	return newBytesPair(key, []byte(value))
}

func newAddressPair(key string, value common.Address) kv.Pair {
	return newStringPair(key, value.Hex())
}

func newUintPair(key string, value uint64) kv.Pair {
	return newStringPair(key, strconv.FormatUint(value, 10))
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
