package shcrypto

import (
	"bytes"
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/pkg/errors"
)

var (
	ErrInputTooLong             = errors.New("input too long")
	ErrInvalidEonSecretKeyShare = errors.New("invalid eon secret key share")
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

// Marshal serializes the eon secret key share.
func (eonSecretKeyShare *EonSecretKeyShare) Marshal() []byte {
	return (*big.Int)(eonSecretKeyShare).Bytes()
}

// Unarshal deserializes an eon secret key share.
func (eonSecretKeyShare *EonSecretKeyShare) Unmarshal(m []byte) error {
	(*big.Int)(eonSecretKeyShare).SetBytes(m)
	if (*big.Int)(eonSecretKeyShare).Cmp(bn256.Order) >= 0 {
		return ErrInvalidEonSecretKeyShare
	}
	return nil
}

// Marshal serializes the eon public key share.
func (eonPublicKeyShare *EonPublicKeyShare) Marshal() []byte {
	return (*bn256.G2)(eonPublicKeyShare).Marshal()
}

// Unmarshal deserializes an eon public key share.
func (eonPublicKeyShare *EonPublicKeyShare) Unmarshal(m []byte) error {
	mLeft, err := (*bn256.G2)(eonPublicKeyShare).Unmarshal(m)
	if len(mLeft) > 0 {
		return ErrInputTooLong
	}
	return err
}

// Marshal serializes the eon public key.
func (eonPublicKey *EonPublicKey) Marshal() []byte {
	return (*bn256.G2)(eonPublicKey).Marshal()
}

// Unmarshal deserializes an eon public key from the given byte slice.
func (eonPublicKey *EonPublicKey) Unmarshal(m []byte) error {
	mLeft, err := (*bn256.G2)(eonPublicKey).Unmarshal(m)
	if len(mLeft) > 0 {
		return ErrInputTooLong
	}
	return err
}

// Marshal serializes the epoch id.
func (epochID *EpochID) Marshal() []byte {
	return (*bn256.G1)(epochID).Marshal()
}

// Unmarshal deserializes an epoch id.
func (epochID *EpochID) Unmarshal(m []byte) error {
	mLeft, err := (*bn256.G1)(epochID).Unmarshal(m)
	if len(mLeft) > 0 {
		return ErrInputTooLong
	}
	return err
}

// Marshal serializes the epoch secret key share.
func (epochSecretKeyShare *EpochSecretKeyShare) Marshal() []byte {
	return (*bn256.G1)(epochSecretKeyShare).Marshal()
}

// Unmarshal deserializes an epoch secret key share.
func (epochSecretKeyShare *EpochSecretKeyShare) Unmarshal(m []byte) error {
	mLeft, err := (*bn256.G1)(epochSecretKeyShare).Unmarshal(m)
	if len(mLeft) > 0 {
		return ErrInputTooLong
	}
	return err
}

// Marshal serializes the epoch secret key.
func (epochSecretKey *EpochSecretKey) Marshal() []byte {
	return (*bn256.G1)(epochSecretKey).Marshal()
}

// Unmarshal deserializes an epoch secret key.
func (epochSecretKey *EpochSecretKey) Unmarshal(m []byte) error {
	mLeft, err := (*bn256.G1)(epochSecretKey).Unmarshal(m)
	if len(mLeft) > 0 {
		return ErrInputTooLong
	}
	return err
}
