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
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"
)

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

// MakePubkeyGeneratedEvent creates a 'shutter.pubkey-generated' tendermint event.  The given
// BatchIndex and PublicKey are encoded as attributes of the event.
func MakePubkeyGeneratedEvent(batchIndex uint64, pubkey *ecdsa.PublicKey) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.pubkey-generated",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("Pubkey"), Value: []byte(encodePubkeyForEvent(pubkey))}},
	}
}

// MakePrivkeyGeneratedEvent creates a 'shutter.privkey-generated' tendermint event. The given
// BatchIndex and PrivateKey are encoded as attributes of the event
func MakePrivkeyGeneratedEvent(batchIndex uint64, privkey *ecdsa.PrivateKey) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.privkey-generated",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("Privkey"), Value: []byte(encodePrivkeyForEvent(privkey))}},
	}
}

// MakeBatchConfigEvent creates a 'shutter.batch-config' tendermint event. The given
// startBatchIndex, threshhold and list of keyper addresses are encoded as attributes of the event.
func MakeBatchConfigEvent(startBatchIndex uint64, threshhold uint32, keypers []common.Address) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.batch-config",
		Attributes: []kv.Pair{
			{Key: []byte("StartBatchIndex"), Value: []byte(fmt.Sprintf("%d", startBatchIndex))},
			{Key: []byte("Threshhold"), Value: []byte(fmt.Sprintf("%d", threshhold))},
			{Key: []byte("Keypers"), Value: []byte(encodeAddressesForEvent(keypers))},
		},
	}
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

// MakeEncryptionKeySignatureAddedEvent creates a 'shutter.encryption-key-signature-added'
// Tendermint event.
func MakeEncryptionKeySignatureAddedEvent(batchIndex uint64, encryptionKey []byte, signature []byte) abcitypes.Event {
	encodedKey := []byte(base64.RawURLEncoding.EncodeToString(encryptionKey))
	encodedSignature := []byte(base64.RawURLEncoding.EncodeToString(signature))
	return abcitypes.Event{
		Type: "shutter.encryption-key-signature-added",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("EncryptionKey"), Value: encodedKey},
			{Key: []byte("Signature"), Value: encodedSignature},
		},
	}
}
