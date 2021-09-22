package shcrypto

import (
	"bytes"
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
