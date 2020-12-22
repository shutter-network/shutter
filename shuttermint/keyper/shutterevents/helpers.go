package shutterevents

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	shcrypto "github.com/brainbot-com/shutter/shuttermint/crypto"
)

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

func newGammas(key string, gammas *shcrypto.Gammas) abcitypes.EventAttribute {
	var encoded []string
	if gammas != nil {
		for _, g := range *gammas {
			encoded = append(encoded, hex.EncodeToString(g.Marshal()))
		}
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

// encodeAddressesForEvent encodes the given slice of Addresses as comma-separated list of addresses
func encodeAddressesForEvent(addr []common.Address) string {
	var hexstrings []string
	for _, a := range addr {
		hexstrings = append(hexstrings, a.Hex())
	}
	return strings.Join(hexstrings, ",")
}

func encodeByteSequenceForEvent(v [][]byte) string {
	var hexstrings []string
	for _, a := range v {
		hexstrings = append(hexstrings, hexutil.Encode(a))
	}
	return strings.Join(hexstrings, ",")
}
