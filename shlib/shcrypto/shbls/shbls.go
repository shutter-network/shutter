package shbls

import (
	"bytes"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

type (
	SecretKey big.Int
	PublicKey bn256.G1
	Signature bn256.G2
)

// RandomKeyPair generates a random BLS secret and corresponding public key.
func RandomKeyPair(r io.Reader) (*SecretKey, *PublicKey, error) {
	i, g1, err := bn256.RandomG1(r)
	if err != nil {
		return nil, nil, err
	}
	return (*SecretKey)(i), (*PublicKey)(g1), nil
}

// SecretToPublicKey returns the BLS public key corresponding to the given secret key.
func SecretToPublicKey(secretKey *SecretKey) *PublicKey {
	secretKeyInt := (*big.Int)(secretKey)
	publicKeyG1 := new(bn256.G1).ScalarBaseMult(secretKeyInt)
	return (*PublicKey)(publicKeyG1)
}

// Sign creates a signature over a message using a secret key.
func Sign(msg []byte, secretKey *SecretKey) *Signature {
	h := hashToG2(msg)
	secretKeyInt := (*big.Int)(secretKey)
	sigG2 := new(bn256.G2).ScalarMult(h, secretKeyInt)
	return (*Signature)(sigG2)
}

// Verify checks that a signature over a certain message was created with the secret key
// corresponding to the given public key.
func Verify(sig *Signature, publicKey *PublicKey, msg []byte) bool {
	h := hashToG2(msg)
	g1 := new(bn256.G1).ScalarBaseMult(big.NewInt(1))
	return bn256.PairingCheck(
		[]*bn256.G1{
			new(bn256.G1).Neg(g1),
			(*bn256.G1)(publicKey),
		},
		[]*bn256.G2{
			(*bn256.G2)(sig),
			h,
		},
	)
}

// AggregateSignatures aggregates a set of BLS signatures into one.
func AggregateSignatures(sigs []*Signature) *Signature {
	aggSig := new(bn256.G2).ScalarBaseMult(big.NewInt(0))
	for _, sig := range sigs {
		aggSig.Add(aggSig, (*bn256.G2)(sig))
	}
	return (*Signature)(aggSig)
}

// AggregatePublicKeys aggregates a set of BLS public keys into one.
func AggregatePublicKeys(publicKeys []*PublicKey) *PublicKey {
	aggPublicKey := new(bn256.G1).ScalarBaseMult(big.NewInt(0))
	for _, publicKey := range publicKeys {
		aggPublicKey.Add(aggPublicKey, (*bn256.G1)(publicKey))
	}
	return (*PublicKey)(aggPublicKey)
}

func hashToG2(data []byte) *bn256.G2 {
	// FIXME: use proper hash to curve algorithm
	// see https://github.com/shutter-network/rolling-shutter/issues/38
	h := crypto.Keccak256Hash(data)
	return new(bn256.G2).ScalarBaseMult(h.Big())
}

// Equal checks if two secret keys are equal to each other.
func (secretKey *SecretKey) Equal(otherSecretKey *SecretKey) bool {
	return (*big.Int)(secretKey).Cmp((*big.Int)(otherSecretKey)) == 0
}

// Equal checks if two public keys are equal to each other.
func (publicKey *PublicKey) Equal(otherPublicKey *PublicKey) bool {
	publicKeyMarshaled := (*bn256.G1)(publicKey).Marshal()
	otherPublicKeyMarshaled := (*bn256.G1)(otherPublicKey).Marshal()
	return bytes.Equal(publicKeyMarshaled, otherPublicKeyMarshaled)
}

// Equal checks if two signatures are equal to each other.
func (sig *Signature) Equal(otherSig *Signature) bool {
	sigMarshaled := (*bn256.G2)(sig).Marshal()
	otherSigMarshaled := (*bn256.G2)(otherSig).Marshal()
	return bytes.Equal(sigMarshaled, otherSigMarshaled)
}
