package shutterevents

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

func encodeUint64(val uint64) []byte {
	return []byte(strconv.FormatUint(val, 10))
}

func decodeUint64(val []byte) (uint64, error) {
	v, err := strconv.ParseUint(string(val), 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse event")
	}
	return v, nil
}

// encodeAddresses encodes the given slice of Addresses as comma-separated list of addresses.
func encodeAddresses(addr []common.Address) []byte {
	var hexstrings []string
	for _, a := range addr {
		hexstrings = append(hexstrings, a.Hex())
	}
	return []byte(strings.Join(hexstrings, ","))
}

// decodeAddresses reverses the encodeAddressesForEvent operation, i.e. it parses a list
// of addresses from a comma-separated string.
func decodeAddresses(val []byte) ([]common.Address, error) {
	s := string(val)
	var res []common.Address
	if s == "" {
		return res, nil
	}
	for _, a := range strings.Split(s, ",") {
		if !common.IsHexAddress(a) {
			return nil, errors.Errorf("malformed address: %q", s)
		}

		res = append(res, common.HexToAddress(a))
	}
	return res, nil
}

func encodeBytes(v []byte) []byte {
	return []byte(hexutil.Encode(v))
}

func decodeBytes(val []byte) ([]byte, error) {
	return hexutil.Decode(string(val))
}

func encodeEpochSecretKeyShare(v *shcrypto.EpochSecretKeyShare) []byte {
	d, _ := v.GobEncode()
	return encodeBytes(d)
}

func decodeEpochSecretKeyShare(v []byte) (*shcrypto.EpochSecretKeyShare, error) {
	decoded, err := decodeBytes(v)
	if err != nil {
		return nil, err
	}
	share := new(shcrypto.EpochSecretKeyShare)
	err = share.GobDecode(decoded)
	if err != nil {
		return nil, err
	}
	return share, nil
}

// encodeByteSequence encodes a slice o byte strings as a comma separated string.
func encodeByteSequence(v [][]byte) []byte {
	var hexstrings []string
	for _, a := range v {
		hexstrings = append(hexstrings, hexutil.Encode(a))
	}
	return []byte(strings.Join(hexstrings, ","))
}

// decodeByteSequence parses a list of hex encoded, comma-separated byte slices.
func decodeByteSequence(val []byte) ([][]byte, error) {
	s := string(val)
	var res [][]byte
	if s == "" {
		return res, nil
	}
	for _, v := range strings.Split(s, ",") {
		bs, err := hexutil.Decode(v)
		if err != nil {
			return [][]byte{}, err
		}
		res = append(res, bs)
	}
	return res, nil
}

// encodePubkey encodes the PublicKey as a string suitable for putting it into a tendermint
// event, i.e. an utf-8 compatible string.
func encodePubkey(pubkey *ecdsa.PublicKey) []byte {
	return []byte(base64.RawURLEncoding.EncodeToString(ethcrypto.FromECDSAPub(pubkey)))
}

// decodePubkey decodes a public key from a tendermint event.
func decodePubkey(val []byte) (*ecdsa.PublicKey, error) {
	data, err := base64.RawURLEncoding.DecodeString(string(val))
	if err != nil {
		return nil, err
	}
	return ethcrypto.UnmarshalPubkey(data)
}

func encodeGammas(gammas *shcrypto.Gammas) []byte {
	var encoded []string
	if gammas != nil {
		for _, g := range *gammas {
			encoded = append(encoded, hex.EncodeToString(g.Marshal()))
		}
	}
	return []byte(strings.Join(encoded, ","))
}

func decodeGammas(eventValue []byte) (shcrypto.Gammas, error) {
	parts := strings.Split(string(eventValue), ",")
	var res shcrypto.Gammas
	for _, p := range parts {
		marshaledG2, err := hex.DecodeString(p)
		if err != nil {
			return shcrypto.Gammas{}, err
		}
		g := new(bn256.G2)
		_, err = g.Unmarshal(marshaledG2)
		if err != nil {
			return shcrypto.Gammas{}, err
		}
		res = append(res, g)
	}
	return res, nil
}

func encodeAddress(a common.Address) []byte {
	return []byte(a.Hex())
}

func decodeAddress(v []byte) (common.Address, error) {
	s := string(v)
	a := common.HexToAddress(s)
	if a.Hex() != s {
		return common.Address{}, errors.Errorf("invalid address %s", s)
	}
	return a, nil
}

func encodeECIESPublicKey(key *ecies.PublicKey) []byte {
	return encodePubkey(key.ExportECDSA())
}

func decodeECIESPublicKey(val []byte) (*ecies.PublicKey, error) {
	publicKeyECDSA, err := decodePubkey(val)
	if err != nil {
		return nil, err
	}
	return ecies.ImportECDSAPublic(publicKeyECDSA), nil
}
