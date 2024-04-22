package shcrypto

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

var (
	ErrInputTooLong             = errors.New("input too long")
	ErrInvalidEonSecretKeyShare = errors.New("invalid eon secret key share")
	ErrVersionMismatch          = func(version_got byte) error {
		return fmt.Errorf("version mismatch. want %d got %d", VersionIdentifier, version_got)
	}
)

// Marshal serializes the EncryptedMessage object. It panics, if C1 is nil.
func (m *EncryptedMessage) Marshal() []byte {
	if m.C1 == nil {
		panic("not a valid encrypted message. C1==nil")
	}

	buff := bytes.Buffer{}
	buff.WriteByte(VersionIdentifier)
	buff.Write(m.C1.Marshal())
	buff.Write(m.C2[:])
	for i := range m.C3 {
		buff.Write(m.C3[i][:])
	}

	return buff.Bytes()
}

// Unmarshal deserializes an EncryptedMessage from the given byte slice.
func (m *EncryptedMessage) Unmarshal(d []byte) error {
	if len(d) == 0 {
		return errors.New("not enough data")
	}
	if m.C1 == nil {
		m.C1 = new(bn256.G2)
	}
	if d[0] != VersionIdentifier {
		return ErrVersionMismatch(m.IdentifyVersion(d))
	}
	d = d[1:]
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

// IdentifyVersion reads the version identifier byte from the given (marshalled) EncryptedMessage.
func (m *EncryptedMessage) IdentifyVersion(d []byte) byte {
	if len(d)%BlockSize == 0 {
		return 0x00
	}
	return d[0]
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

// MarshalText serializes the eon public key to hex.
func (eonPublicKey EonPublicKey) MarshalText() ([]byte, error) {
	return hexutil.Bytes(eonPublicKey.Marshal()).MarshalText()
}

// UnmarshalText deserializes the eon public key from hex.
func (eonPublicKey *EonPublicKey) UnmarshalText(input []byte) error {
	var b hexutil.Bytes
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
	return eonPublicKey.Unmarshal(b)
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

// MarshalText serializes the epoch secret key to hex.
func (epochSecretKey EpochSecretKey) MarshalText() ([]byte, error) { //nolint: unparam
	return []byte(hexutil.Encode(epochSecretKey.Marshal())), nil
}

// UnmarshalText deserializes the epoch secret key from hex.
func (epochSecretKey *EpochSecretKey) UnmarshalText(input []byte) error {
	var b hexutil.Bytes
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
	return epochSecretKey.Unmarshal(b)
}

// MarshalText serializes the block to hex.
func (block Block) MarshalText() ([]byte, error) { //nolint:unparam
	return []byte(hexutil.Encode(block[:])), nil
}

// UnmarshalText deserializes the block from hex.
func (block *Block) UnmarshalText(b []byte) error {
	decoded, err := hexutil.Decode(string(b))
	copy(block[:], decoded)
	return err
}

// MarshalText serializes the encrypted message to hex.
func (m EncryptedMessage) MarshalText() ([]byte, error) { //nolint:unparam
	return []byte(hexutil.Encode(m.Marshal())), nil
}

// UnmarshalText deserializes the encrypted message from hex.
func (m *EncryptedMessage) UnmarshalText(b []byte) error {
	decoded, err := hexutil.Decode(string(b))
	if err != nil {
		return err
	}
	err = m.Unmarshal(decoded)
	return err
}
