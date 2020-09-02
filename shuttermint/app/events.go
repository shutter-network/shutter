package app

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

func DecodePubkeyFromEvent(s string) (*ecdsa.PublicKey, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalPubkey(data)
}
func encodePrivkeyForEvent(privkey *ecdsa.PrivateKey) string {
	return base64.RawURLEncoding.EncodeToString(crypto.FromECDSA(privkey))
}

func DecodePrivkeyFromEvent(s string) (*ecdsa.PrivateKey, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return crypto.ToECDSA(data)
}

func MakePubkeyGeneratedEvent(batchIndex uint64, pubkey *ecdsa.PublicKey) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.pubkey-generated",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("Pubkey"), Value: []byte(encodePubkeyForEvent(pubkey))}},
	}
}

func MakePrivkeyGeneratedEvent(batchIndex uint64, privkey *ecdsa.PrivateKey) abcitypes.Event {
	return abcitypes.Event{
		Type: "shutter.privkey-generated",
		Attributes: []kv.Pair{
			{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", batchIndex))},
			{Key: []byte("Privkey"), Value: []byte(encodePrivkeyForEvent(privkey))}},
	}
}

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
func encodeAddressesForEvent(addr []common.Address) string {
	var hex []string
	for _, a := range addr {
		hex = append(hex, a.Hex())
	}
	return strings.Join(hex, ",")
}

func DecodeAddressesFromEvent(s string) []common.Address {
	var res []common.Address
	for _, a := range strings.Split(s, ",") {
		res = append(res, common.HexToAddress(a))
	}
	return res
}
