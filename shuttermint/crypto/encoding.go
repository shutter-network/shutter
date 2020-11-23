package crypto

import (
	"bytes"
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
	return rlp.Encode(w, ([32]byte)(b))
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
