package shcrypto

import (
	"bytes"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/pkg/errors"
)

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
		return errors.Errorf("length not a multiple of %d", BlockSize)
	}
	if len(d) < BlockSize {
		return errors.Errorf("short block")
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

// Marshal serialized the eon public key.
func (eonpubkey *EonPublicKey) Marshal() []byte {
	return (*bn256.G2)(eonpubkey).Marshal()
}

// Unmarshal deserializes an eon public key from the given byte slice.
func (eonpubkey *EonPublicKey) Unmarshal(m []byte) error {
	_, err := (*bn256.G2)(eonpubkey).Unmarshal(m)
	return err
}
