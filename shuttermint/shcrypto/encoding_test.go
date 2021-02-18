package shcrypto

import (
	"bytes"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/stretchr/testify/require"
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
	require.Nil(t, err)
	require.Equal(t, m1, m2)
}

func TestUnmarshalBroken(t *testing.T) {
	d := encryptedMessage().Marshal()
	m := EncryptedMessage{}

	err := m.Unmarshal(d[:16])
	require.NotNil(t, err)

	err = m.Unmarshal(d[:32])
	require.NotNil(t, err)

	err = m.Unmarshal(d[:65])
	require.NotNil(t, err)

	err = m.Unmarshal(d[:len(d)-1])
	require.NotNil(t, err)
}
