package shbls

import (
	"bytes"
	"crypto/rand"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"gotest.tools/assert"
)

func TestMarshalUnmarshalSecretKey(t *testing.T) {
	secretKey, _, err := RandomKeyPair(rand.Reader)
	assert.NilError(t, err)

	marshaled := secretKey.Marshal()
	assert.Check(t, len(marshaled) == 32)

	unmarshaled := new(SecretKey)
	err = unmarshaled.Unmarshal(marshaled)
	assert.NilError(t, err)

	assert.Check(t, unmarshaled.Equal(secretKey))
}

func TestUnmarshalInvalidSecretKey(t *testing.T) {
	m1 := make([]byte, 32)
	bn256.Order.FillBytes(m1)
	sk1 := new(SecretKey)
	err := sk1.Unmarshal(m1)
	assert.Error(t, err, "secret key must be smaller than the order")

	m2 := make([]byte, 31)
	sk2 := new(SecretKey)
	err = sk2.Unmarshal(m2)
	assert.Error(t, err, "secret key must be 32 bytes, got 31")

	m3 := make([]byte, 33)
	sk3 := new(SecretKey)
	err = sk3.Unmarshal(m3)
	assert.Error(t, err, "secret key must be 32 bytes, got 33")
}

func TestMarshalUnmarshalPublicKey(t *testing.T) {
	_, publicKey, err := RandomKeyPair(rand.Reader)
	assert.NilError(t, err)

	marshaled := publicKey.Marshal()
	assert.Check(t, len(marshaled) == 64)

	unmarshaled := new(PublicKey)
	err = unmarshaled.Unmarshal(marshaled)
	assert.NilError(t, err)

	assert.Check(t, unmarshaled.Equal(publicKey))
}

func TestUnmarshalInvalidPublicKey(t *testing.T) {
	m1 := bytes.Repeat([]byte{255}, 64)
	pk1 := new(PublicKey)
	err := pk1.Unmarshal(m1)
	assert.ErrorContains(t, err, "public key invalid")

	m2 := make([]byte, 63)
	pk2 := new(PublicKey)
	err = pk2.Unmarshal(m2)
	assert.Error(t, err, "public key must be 64 bytes, got 63")

	m3 := make([]byte, 65)
	pk3 := new(PublicKey)
	err = pk3.Unmarshal(m3)
	assert.Error(t, err, "public key must be 64 bytes, got 65")
}

func TestMarshalUnmarshalSignature(t *testing.T) {
	secretKey, _, err := RandomKeyPair(rand.Reader)
	assert.NilError(t, err)
	sig := Sign([]byte("hello"), secretKey)

	marshaled := sig.Marshal()
	assert.Check(t, len(marshaled) == 128)

	unmarshaled := new(Signature)
	err = unmarshaled.Unmarshal(marshaled)
	assert.NilError(t, err)

	assert.Check(t, unmarshaled.Equal(sig))
}

func TestUnmarshalInvalidSignature(t *testing.T) {
	m1 := bytes.Repeat([]byte{255}, 128)
	sig1 := new(Signature)
	err := sig1.Unmarshal(m1)
	assert.ErrorContains(t, err, "signature invalid")

	m2 := make([]byte, 127)
	sig2 := new(Signature)
	err = sig2.Unmarshal(m2)
	assert.Error(t, err, "signature must be 128 bytes, got 127")

	m3 := make([]byte, 129)
	sig3 := new(Signature)
	err = sig3.Unmarshal(m3)
	assert.Error(t, err, "signature must be 128 bytes, got 129")
}
