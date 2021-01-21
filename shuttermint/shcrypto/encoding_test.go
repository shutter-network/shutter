package shcrypto

import (
	"bytes"
	"math/big"
	"reflect"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

func TestBlockEncoding(t *testing.T) {
	blockSlice := bytes.Repeat([]byte("a"), 32)
	var b1 Block
	copy(b1[:], blockSlice)

	encoded, err := rlp.EncodeToBytes(b1)
	require.Nil(t, err)

	var b2 Block
	err = rlp.DecodeBytes(encoded, &b2)
	require.Nil(t, err)
	require.Equal(t, b1, b2)

	tooShort := bytes.Repeat([]byte("a"), 31)
	tooShortEncoded, err := rlp.EncodeToBytes(tooShort)
	require.Nil(t, err)
	err = rlp.DecodeBytes(tooShortEncoded, new(Block))
	require.NotNil(t, err)

	tooLong := bytes.Repeat([]byte("a"), 33)
	tooLongEncoded, err := rlp.EncodeToBytes(tooLong)
	require.Nil(t, err)
	err = rlp.DecodeBytes(tooLongEncoded, new(Block))
	require.NotNil(t, err)

	notBytes := []uint32{1, 2, 3}
	notBytesEncoded, err := rlp.EncodeToBytes(notBytes)
	require.Nil(t, err)
	err = rlp.DecodeBytes(notBytesEncoded, new(Block))
	require.NotNil(t, err)
}

func TestPointEncoding(t *testing.T) {
	buf := new(bytes.Buffer)

	p1 := new(bn256.G2).ScalarBaseMult(big.NewInt(5))
	err := encodeRLPG2(buf, p1)
	require.Nil(t, err)

	p2 := new(bn256.G2)
	s := rlp.NewStream(buf, 0)
	err = decodeRLPG2(s, p2)
	require.Nil(t, err)
	require.Equal(t, p1.Marshal(), p2.Marshal())

	tooShort := []byte("abc")
	tooShortEncoded, err := rlp.EncodeToBytes(tooShort)
	require.Nil(t, err)
	s = rlp.NewStream(bytes.NewBuffer(tooShortEncoded), 0)
	err = decodeRLPG2(s, new(bn256.G2))
	require.NotNil(t, err)

	tooLong := bytes.Repeat([]byte("x"), 32*4+1)
	tooLongEncoded, err := rlp.EncodeToBytes(tooLong)
	require.Nil(t, err)
	s = rlp.NewStream(bytes.NewBuffer(tooLongEncoded), 0)
	err = decodeRLPG2(s, new(bn256.G2))
	require.NotNil(t, err)

	notBytes := []uint32{1, 2, 3}
	notBytesEncoded, err := rlp.EncodeToBytes(notBytes)
	require.Nil(t, err)
	s = rlp.NewStream(bytes.NewBuffer(notBytesEncoded), 0)
	err = decodeRLPG2(s, new(bn256.G2))
	require.NotNil(t, err)
}

func TestMessageEncoding(t *testing.T) {
	blocks := []Block{}
	for i := 0; i < 3; i++ {
		s := bytes.Repeat([]byte{byte(i)}, 32)
		var b Block
		copy(b[:], s)
		blocks = append(blocks, b)
	}

	m1 := EncryptedMessage{
		C1: new(bn256.G2).ScalarBaseMult(big.NewInt(5)),
		C2: blocks[0],
		C3: blocks[1:],
	}
	encoded, err := rlp.EncodeToBytes(m1)
	require.Nil(t, err)

	m2 := new(EncryptedMessage)
	err = rlp.DecodeBytes(encoded, &m2)
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(m1.C1, m2.C1))
	require.True(t, reflect.DeepEqual(m1.C2, m2.C2))
	require.True(t, reflect.DeepEqual(m1.C3, m2.C3))
}
