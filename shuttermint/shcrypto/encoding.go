package shcrypto

import (
	"bytes"
	"fmt"
	"io"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/ethereum/go-ethereum/rlp"
)

type encryptedMessageRLP struct {
	C1 []byte
	C2 Block
	C3 []Block
}

// EncodeRLP implements RLP encoding for blocks.
func (b Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, [32]byte(b))
}

// DecodeRLP implements RLP decoding for blocks.
func (b *Block) DecodeRLP(s *rlp.Stream) error {
	return s.Decode((*[32]byte)(b))
}

// encodeRLPG2 implements RLP encoding for points on G2.
func encodeRLPG2(w io.Writer, p *bn256.G2) error {
	return rlp.Encode(w, p.Marshal())
}

// decodeRLPG2 implements RLP decoding for points on G2.
func decodeRLPG2(s *rlp.Stream, p *bn256.G2) error {
	bs, err := s.Bytes()
	if err != nil {
		return err
	}

	_, err = p.Unmarshal(bs)
	if err != nil {
		return err
	}

	return nil
}

// EncodeRLP implements RLP encoding for encrypted messages.
func (m EncryptedMessage) EncodeRLP(w io.Writer) error {
	c1Buf := new(bytes.Buffer)
	err := encodeRLPG2(c1Buf, m.C1)
	if err != nil {
		return err
	}
	c1Bytes := c1Buf.Bytes()

	mRLP := encryptedMessageRLP{
		C1: c1Bytes,
		C2: m.C2,
		C3: m.C3,
	}
	return rlp.Encode(w, mRLP)
}

// DecodeRLP implements RLP decoding for encrypted messages.
func (m *EncryptedMessage) DecodeRLP(s *rlp.Stream) error {
	mRLP := new(encryptedMessageRLP)
	err := s.Decode(mRLP)
	if err != nil {
		return err
	}

	c1Stream := rlp.NewStream(bytes.NewBuffer(mRLP.C1), 0)
	c1 := new(bn256.G2)
	err = decodeRLPG2(c1Stream, c1)
	if err != nil {
		return err
	}

	m.C1 = c1
	m.C2 = mRLP.C2
	m.C3 = mRLP.C3
	return nil
}

/* The following implements serialization via custom Marshal/Unmarshal methods.  Unless we need RLP
   encoding on the contract side, we will use these custom methods, as they are much simpler.
*/

// Marshal serializes the EncryptedMessage object. It panics, if C1 is nil.
func (m *EncryptedMessage) Marshal() []byte {
	if m.C1 == nil {
		panic("not a valid encrypted message. C1==nil")
	}

	buff := bytes.Buffer{}
	buff.Write(m.C1.Marshal())
	buff.Write(m.C2[:])
	for i := range m.C3 {
		buff.Write(m.C3[i][:])
	}

	return buff.Bytes()
}

// Unmarshal deserializes an EncryptedMessage from the given byte slice.
func (m *EncryptedMessage) Unmarshal(d []byte) error {
	if m.C1 == nil {
		m.C1 = new(bn256.G2)
	}
	d, err := m.C1.Unmarshal(d)
	if err != nil {
		return err
	}
	if len(d)%BlockSize != 0 {
		return fmt.Errorf("length not a multiple of %d", BlockSize)
	}
	if len(d) < BlockSize {
		return fmt.Errorf("short block")
	}
	copy(m.C2[:], d)
	d = d[BlockSize:]
	m.C3 = nil
	for len(d) > 0 {
		b := Block{}
		copy(b[:], d)
		d = d[BlockSize:]
		m.C3 = append(m.C3, b)
	}
	return nil
}
