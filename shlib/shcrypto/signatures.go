package shcrypto

import (
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

type (
	BLSSecretKey big.Int
	BLSPublicKey bn256.G1
	BLSSignature bn256.G2
)

// RandomBLSKeyPair generates a random BLS secret and corresponding public key.
func RandomBLSKeyPair(r io.Reader) (*BLSSecretKey, *BLSPublicKey, error) {
	i, g1, err := bn256.RandomG1(r)
	if err != nil {
		return nil, nil, err
	}
	return (*BLSSecretKey)(i), (*BLSPublicKey)(g1), nil
}

// BLSSecretToPublicKey returns the BLS public key corresponding to the given secret key.
func BLSSecretToPublicKey(secretKey *BLSSecretKey) *BLSPublicKey {
	secretKeyInt := (*big.Int)(secretKey)
	publicKeyG1 := new(bn256.G1).ScalarBaseMult(secretKeyInt)
	return (*BLSPublicKey)(publicKeyG1)
}

// BLSSign creates a signature over a message using a secret key.
func BLSSign(msg []byte, secretKey *BLSSecretKey) *BLSSignature {
	h := hashToG2(msg)
	secretKeyInt := (*big.Int)(secretKey)
	sigG2 := new(bn256.G2).ScalarMult(h, secretKeyInt)
	return (*BLSSignature)(sigG2)
}

// BLSVerify checks that a signature over a certain message was created with the secret key
// corresponding to the given public key.
func BLSVerify(sig *BLSSignature, publicKey *BLSPublicKey, msg []byte) bool {
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

// BLSAggregateSignatures aggregates a set of BLS signatures into one.
func BLSAggregateSignatures(sigs []*BLSSignature) *BLSSignature {
	aggSig := new(bn256.G2).ScalarBaseMult(big.NewInt(0))
	for _, sig := range sigs {
		aggSig.Add(aggSig, (*bn256.G2)(sig))
	}
	return (*BLSSignature)(aggSig)
}

// BLSAggregatePublicKeys aggregates a set of BLS public keys into one.
func BLSAggregatePublicKeys(publicKeys []*BLSPublicKey) *BLSPublicKey {
	aggPublicKey := new(bn256.G1).ScalarBaseMult(big.NewInt(0))
	for _, publicKey := range publicKeys {
		aggPublicKey.Add(aggPublicKey, (*bn256.G1)(publicKey))
	}
	return (*BLSPublicKey)(aggPublicKey)
}

func hashToG2(data []byte) *bn256.G2 {
	// FIXME: use proper hash to curve algorithm
	// see https://github.com/shutter-network/rolling-shutter/issues/38
	h := crypto.Keccak256Hash(data)
	return new(bn256.G2).ScalarBaseMult(h.Big())
}
