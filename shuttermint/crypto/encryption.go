package crypto

import (
	"crypto/rand"
	"io"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// SymmetricKey is the key used to encrypt and decrypt messages.
type SymmetricKey common.Hash

// EncryptionResult represents the full output of the encryption procedure.
type EncryptionResult struct {
	RandomnessWitness *bn256.G2
	CipherText        []byte
}

// RandomR returns a random value to be used during encryption.
func RandomR(r io.Reader) (*big.Int, error) {
	return rand.Int(r, bn256.Order)
}

func computeEncryptionKey(r *big.Int, eonPK *EonPK, epochID *EpochID) *SymmetricKey {
	g := bn256.Pair((*bn256.G1)(epochID), (*bn256.G2)(eonPK))
	g = new(bn256.GT).ScalarMult(g, r)
	log.Println("enc")
	log.Println(g.String())
	gBytes := g.Marshal()
	h := crypto.Keccak256Hash(gBytes)
	k := SymmetricKey(h)
	return &k
}

func computeCipherText(message []byte, key *SymmetricKey) []byte {
	// TODO: encryption
	return message
}

// Encrypt encrypts a message for a given epoch using the provided eon key and randomness.
func Encrypt(message []byte, r *big.Int, eonPK *EonPK, epochID *EpochID) *EncryptionResult {
	key := computeEncryptionKey(r, eonPK, epochID)
	cipherText := computeCipherText(message, key)
	c := EncryptionResult{
		RandomnessWitness: new(bn256.G2).ScalarBaseMult(r),
		CipherText:        cipherText,
	}
	return &c
}

func computeDecryptionKey(epochSK *EpochSK, c *EncryptionResult) *SymmetricKey {
	g := bn256.Pair((*bn256.G1)(epochSK), c.RandomnessWitness)
	log.Println("dec")
	log.Println(g.String())
	gBytes := g.Marshal()
	h := crypto.Keccak256Hash(gBytes)
	k := SymmetricKey(h)
	return &k
}

// Decrypt decrypts an encrypted message using the given epoch secret key.
func Decrypt(c *EncryptionResult, epochSK *EpochSK) []byte {
	key := computeDecryptionKey(epochSK, c)
	p := computePlainText(c.CipherText, key)
	return p
}

func computePlainText(cipherText []byte, key *SymmetricKey) []byte {
	// TODO: decryption
	return cipherText
}
