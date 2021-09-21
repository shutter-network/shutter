package shcrypto

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"gotest.tools/v3/assert"
)

func encryptedMessage() *EncryptedMessage {
	blocks := []Block{}
	for i := 0; i < 3; i++ {
		s := bytes.Repeat([]byte{byte(i)}, 32)
		var b Block
		copy(b[:], s)
		blocks = append(blocks, b)
	}

	return &EncryptedMessage{
		C1: new(bn256.G2).ScalarBaseMult(big.NewInt(5)),
		C2: blocks[0],
		C3: blocks[1:],
	}
}

func TestMarshalUnmarshal(t *testing.T) {
	m1 := encryptedMessage()
	m2 := &EncryptedMessage{}
	err := m2.Unmarshal(m1.Marshal())
	assert.NilError(t, err)
	assert.DeepEqual(t, m1, m2, G2Comparer)
}

func TestUnmarshalBroken(t *testing.T) {
	d := encryptedMessage().Marshal()
	m := EncryptedMessage{}

	err := m.Unmarshal(d[:16])
	assert.Assert(t, err != nil)

	err = m.Unmarshal(d[:32])
	assert.Assert(t, err != nil)

	err = m.Unmarshal(d[:65])
	assert.Assert(t, err != nil)

	err = m.Unmarshal(d[:len(d)-1])
	assert.Assert(t, err != nil)
}

func TestMarshalUnmarshalBLSSecretKey(t *testing.T) {
	secretKey, _, err := RandomBLSKeyPair(rand.Reader)
	assert.NilError(t, err)

	marshaled := secretKey.Marshal()
	assert.Check(t, len(marshaled) == 32)

	unmarshaled := new(BLSSecretKey)
	err = unmarshaled.Unmarshal(marshaled)
	assert.NilError(t, err)

	assert.Check(t, unmarshaled.Equal(secretKey))
}

func TestUnmarshalInvalidBLSSecretKey(t *testing.T) {
	m1 := make([]byte, 32)
	bn256.Order.FillBytes(m1)
	sk1 := new(BLSSecretKey)
	err := sk1.Unmarshal(m1)
	assert.Error(t, err, "secret key must be smaller than the order")

	m2 := make([]byte, 31)
	sk2 := new(BLSSecretKey)
	err = sk2.Unmarshal(m2)
	assert.Error(t, err, "secret key must be 32 bytes, got 31")

	m3 := make([]byte, 33)
	sk3 := new(BLSSecretKey)
	err = sk3.Unmarshal(m3)
	assert.Error(t, err, "secret key must be 32 bytes, got 33")
}

func TestMarshalUnmarshalBLSPublicKey(t *testing.T) {
	_, publicKey, err := RandomBLSKeyPair(rand.Reader)
	assert.NilError(t, err)

	marshaled := publicKey.Marshal()
	assert.Check(t, len(marshaled) == 64)

	unmarshaled := new(BLSPublicKey)
	err = unmarshaled.Unmarshal(marshaled)
	assert.NilError(t, err)

	assert.Check(t, unmarshaled.Equal(publicKey))
}

func TestUnmarshalInvalidBLSPublicKey(t *testing.T) {
	m1 := bytes.Repeat([]byte{255}, 64)
	pk1 := new(BLSPublicKey)
	err := pk1.Unmarshal(m1)
	assert.ErrorContains(t, err, "public key invalid")

	m2 := make([]byte, 63)
	pk2 := new(BLSPublicKey)
	err = pk2.Unmarshal(m2)
	assert.Error(t, err, "public key must be 64 bytes, got 63")

	m3 := make([]byte, 65)
	pk3 := new(BLSPublicKey)
	err = pk3.Unmarshal(m3)
	assert.Error(t, err, "public key must be 64 bytes, got 65")
}

func TestMarshalUnmarshalBLSSignature(t *testing.T) {
	secretKey, _, err := RandomBLSKeyPair(rand.Reader)
	assert.NilError(t, err)
	sig := BLSSign([]byte("hello"), secretKey)

	marshaled := sig.Marshal()
	assert.Check(t, len(marshaled) == 128)

	unmarshaled := new(BLSSignature)
	err = unmarshaled.Unmarshal(marshaled)
	assert.NilError(t, err)

	assert.Check(t, unmarshaled.Equal(sig))
}

func TestUnmarshalInvalidBLSSignature(t *testing.T) {
	m1 := bytes.Repeat([]byte{255}, 128)
	sig1 := new(BLSSignature)
	err := sig1.Unmarshal(m1)
	assert.ErrorContains(t, err, "signature invalid")

	m2 := make([]byte, 127)
	sig2 := new(BLSSignature)
	err = sig2.Unmarshal(m2)
	assert.Error(t, err, "signature must be 128 bytes, got 127")

	m3 := make([]byte, 129)
	sig3 := new(BLSSignature)
	err = sig3.Unmarshal(m3)
	assert.Error(t, err, "signature must be 128 bytes, got 129")
}
