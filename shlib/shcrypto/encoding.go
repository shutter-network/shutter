package shcrypto

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/bls12381"
)

var (
	ErrInputTooLong             = errors.New("input too long")
	ErrInvalidEonSecretKeyShare = errors.New("invalid eon secret key share")
	ErrVersionMismatch          = func(version_got byte) error {
		return fmt.Errorf("version mismatch. want %d got %d", VersionIdentifier, version_got)
	}
)

const (
	g2EncodingLength = 192
)

// Marshal serializes the EncryptedMessage object. It panics, if C1 is nil.
func (m *EncryptedMessage) Marshal() []byte {
	if m.C1 == nil {
		panic("not a valid encrypted message. C1==nil")
	}
	g2 := bls12381.NewG2()

	buff := bytes.Buffer{}
	buff.WriteByte(VersionIdentifier)
	buff.Write(g2.ToBytes(m.C1))
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
	if d[0] != VersionIdentifier {
		return ErrVersionMismatch(IdentifyVersion(d))
	}

	var err error
	if len(d) < 1+g2EncodingLength+BlockSize ||
		(len(d)-1-g2EncodingLength-BlockSize)%BlockSize != 0 {
		return fmt.Errorf("invalid length %d of encrypted message", len(d))
	}
	g2 := bls12381.NewG2()
	if m.C1 == nil {
		m.C1 = new(bls12381.PointG2)
	}
	m.C1, err = g2.FromBytes(d[1 : 1+g2EncodingLength])
	if err != nil {
		return err
	}
	if !g2.IsOnCurve(m.C1) {
		return errors.New("C1 not on curve")
	}

	copy(m.C2[:], d[1+g2EncodingLength:1+g2EncodingLength+BlockSize])

	m.C3 = nil
	for i := 1 + g2EncodingLength + BlockSize; i < len(d); i += BlockSize {
		b := Block{}
		copy(b[:], d[i:i+BlockSize])
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
	if (*big.Int)(eonSecretKeyShare).Cmp(order) >= 0 {
		return ErrInvalidEonSecretKeyShare
	}
	return nil
}

// Marshal serializes the eon public key share.
func (eonPublicKeyShare *EonPublicKeyShare) Marshal() []byte {
	g2 := bls12381.NewG2()
	return g2.ToBytes((*bls12381.PointG2)(eonPublicKeyShare))
}

// Unmarshal deserializes an eon public key share.
func (eonPublicKeyShare *EonPublicKeyShare) Unmarshal(m []byte) error {
	g2 := bls12381.NewG2()
	p, err := g2.FromBytes(m)
	if err != nil {
		return err
	}
	if !g2.IsOnCurve(p) {
		return errors.New("not on curve")
	}
	(*bls12381.PointG2)(eonPublicKeyShare).Set(p)
	return nil
}

// Marshal serializes the eon public key.
func (eonPublicKey *EonPublicKey) Marshal() []byte {
	g2 := bls12381.NewG2()
	return g2.ToBytes((*bls12381.PointG2)(eonPublicKey))
}

// Unmarshal deserializes an eon public key from the given byte slice.
func (eonPublicKey *EonPublicKey) Unmarshal(m []byte) error {
	g2 := bls12381.NewG2()
	p, err := g2.FromBytes(m)
	if err != nil {
		return err
	}
	if !g2.IsOnCurve(p) {
		return errors.New("not on curve")
	}
	(*bls12381.PointG2)(eonPublicKey).Set(p)
	return nil
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
	g1 := bls12381.NewG1()
	return g1.ToBytes((*bls12381.PointG1)(epochID))
}

// Unmarshal deserializes an epoch id.
func (epochID *EpochID) Unmarshal(m []byte) error {
	g1 := bls12381.NewG1()
	p, err := g1.FromBytes(m)
	if err != nil {
		return err
	}
	if !g1.IsOnCurve(p) {
		return errors.New("not on curve")
	}
	(*bls12381.PointG1)(epochID).Set(p)
	return nil
}

// Marshal serializes the epoch secret key share.
func (epochSecretKeyShare *EpochSecretKeyShare) Marshal() []byte {
	g1 := bls12381.NewG1()
	return g1.ToBytes((*bls12381.PointG1)(epochSecretKeyShare))
}

// Unmarshal deserializes an epoch secret key share.
func (epochSecretKeyShare *EpochSecretKeyShare) Unmarshal(m []byte) error {
	g1 := bls12381.NewG1()
	p, err := g1.FromBytes(m)
	if err != nil {
		return err
	}
	if !g1.IsOnCurve(p) {
		return errors.New("not on curve")
	}
	(*bls12381.PointG1)(epochSecretKeyShare).Set(p)
	return nil
}

// Marshal serializes the epoch secret key.
func (epochSecretKey *EpochSecretKey) Marshal() []byte {
	g1 := bls12381.NewG1()
	return g1.ToBytes((*bls12381.PointG1)(epochSecretKey))
}

// Unmarshal deserializes an epoch secret key.
func (epochSecretKey *EpochSecretKey) Unmarshal(m []byte) error {
	g1 := bls12381.NewG1()
	p, err := g1.FromBytes(m)
	if err != nil {
		return err
	}
	if !g1.IsOnCurve(p) {
		return errors.New("not on curve")
	}
	(*bls12381.PointG1)(epochSecretKey).Set(p)
	return nil
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
