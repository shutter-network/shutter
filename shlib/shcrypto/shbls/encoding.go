package shbls

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/pkg/errors"
)

// Marshal converts a BLS secret key to bytes. The result is a 32 byte slice. Make sure the input
// is a valid key, otherwise the method might panic or return something that when decoded would not
// match the original value.
func (secretKey *SecretKey) Marshal() []byte {
	b := make([]byte, 32)
	(*big.Int)(secretKey).FillBytes(b)
	return b
}

// Unmarshal sets the secret key to the value stored in the given byte slice. The input byte slice
// must be 32 bytes and contain a valid key, otherwise an error is returned.
func (secretKey *SecretKey) Unmarshal(b []byte) error {
	if len(b) != 32 {
		return errors.Errorf("secret key must be 32 bytes, got %d", len(b))
	}

	secretKeyInt := (*big.Int)(secretKey)
	secretKeyInt.SetBytes(b)
	if secretKeyInt.Cmp(bn256.Order) > -1 {
		secretKeyInt.SetUint64(0) // set value to zero to avoid using a corrupt one
		return errors.New("secret key must be smaller than the order")
	}

	return nil
}

// Marshal converts a BLS public key to bytes.
func (publicKey *PublicKey) Marshal() []byte {
	return (*bn256.G1)(publicKey).Marshal()
}

// Unmarshal sets the public key to the value stored in the given byte slice. The input byte slice
// must be 64 bytes and contain a valid key, otherwise an error is returned.
func (publicKey *PublicKey) Unmarshal(b []byte) error {
	if len(b) != 64 {
		return errors.Errorf("public key must be 64 bytes, got %d", len(b))
	}
	publicKeyG1 := (*bn256.G1)(publicKey)
	_, err := publicKeyG1.Unmarshal(b)
	return errors.Wrap(err, "public key invalid")
}

// Marshal converts a BLS signature to bytes.
func (sig *Signature) Marshal() []byte {
	return (*bn256.G2)(sig).Marshal()
}

// Unmarshal sets the signature to the value stored in the given byte slice. The input byte slice
// must be 128 bytes and contain a well formed signature, otherwise an error is returned.
func (sig *Signature) Unmarshal(b []byte) error {
	if len(b) != 128 {
		return errors.Errorf("signature must be 128 bytes, got %d", len(b))
	}
	sigG2 := (*bn256.G2)(sig)
	_, err := sigG2.Unmarshal(b)
	return errors.Wrap(err, "signature invalid")
}
