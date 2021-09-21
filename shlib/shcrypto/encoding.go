package shcrypto

import (
	"bytes"
	"math/big"

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

// Marshal converts a BLS secret key to bytes. The result is a 32 byte slice. Make sure the input
// is a valid key, otherwise the method might panic or return something that when decoded would not
// match the original value.
func (blsSecretKey *BLSSecretKey) Marshal() []byte {
	b := make([]byte, 32)
	(*big.Int)(blsSecretKey).FillBytes(b)
	return b
}

// Unmarshal sets the secret key to the value stored in the given byte slice. The input byte slice
// must be 32 bytes and contain a valid key, otherwise an error is returned.
func (blsSecretKey *BLSSecretKey) Unmarshal(b []byte) error {
	if len(b) != 32 {
		return errors.Errorf("secret key must be 32 bytes, got %d", len(b))
	}

	secretKeyInt := (*big.Int)(blsSecretKey)
	secretKeyInt.SetBytes(b)
	if secretKeyInt.Cmp(bn256.Order) > -1 {
		secretKeyInt.SetUint64(0) // set value to zero to avoid using a corrupt one
		return errors.New("secret key must be smaller than the order")
	}

	return nil
}

// Marshal converts a BLS public key to bytes.
func (blsPublicKey *BLSPublicKey) Marshal() []byte {
	return (*bn256.G1)(blsPublicKey).Marshal()
}

// Unmarshal sets the public key to the value stored in the given byte slice. The input byte slice
// must be 64 bytes and contain a valid key, otherwise an error is returned.
func (blsPublicKey *BLSPublicKey) Unmarshal(b []byte) error {
	if len(b) != 64 {
		return errors.Errorf("public key must be 64 bytes, got %d", len(b))
	}
	publicKeyG1 := (*bn256.G1)(blsPublicKey)
	_, err := publicKeyG1.Unmarshal(b)
	return errors.Wrap(err, "public key invalid")
}

// Marshal converts a BLS signature to bytes.
func (blsSignature *BLSSignature) Marshal() []byte {
	return (*bn256.G2)(blsSignature).Marshal()
}

// Unmarshal sets the signature to the value stored in the given byte slice. The input byte slice
// must be 128 bytes and contain a valid key, otherwise an error is returned.
func (blsSignature *BLSSignature) Unmarshal(b []byte) error {
	if len(b) != 128 {
		return errors.Errorf("signature must be 128 bytes, got %d", len(b))
	}
	signatureG2 := (*bn256.G2)(blsSignature)
	_, err := signatureG2.Unmarshal(b)
	return errors.Wrap(err, "signature invalid")
}
