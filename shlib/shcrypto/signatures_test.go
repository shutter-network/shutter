package shcrypto

import (
	"crypto/rand"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSigning(t *testing.T) {
	secretKey, publicKey, err := RandomBLSKeyPair(rand.Reader)
	assert.NilError(t, err)

	msg := []byte("hello")
	sig := BLSSign(msg, secretKey)
	assert.Check(t, BLSVerify(sig, publicKey, msg))

	secretKey2, publicKey2, err := RandomBLSKeyPair(rand.Reader)
	assert.NilError(t, err)
	sig2 := BLSSign(msg, secretKey2)

	assert.Check(t, !BLSVerify(sig, publicKey, []byte("good bye")))
	assert.Check(t, !BLSVerify(sig, publicKey2, msg))
	assert.Check(t, !BLSVerify(sig2, publicKey, msg))
}

func TestAggregation(t *testing.T) {
	msg := []byte("hello")

	allsecretKeys := []BLSSecretKey{}
	allPublicKeys := []BLSPublicKey{}
	allSigs := []BLSSignature{}
	for i := 0; i < 3; i++ {
		secretKey, publicKey, err := RandomBLSKeyPair(rand.Reader)
		assert.NilError(t, err)
		sig := BLSSign(msg, secretKey)

		allsecretKeys = append(allsecretKeys, secretKey)
		allPublicKeys = append(allPublicKeys, publicKey)
		allSigs = append(allSigs, sig)
	}

	combinations := [][]int{
		{0, 1},
		{1, 2},
		{0, 2},
		{0, 1, 2},
	}

	for _, signers := range combinations {
		secretKeys := []BLSSecretKey{}
		publicKeys := []BLSPublicKey{}
		sigs := []BLSSignature{}
		for _, signer := range signers {
			secretKeys = append(secretKeys, allsecretKeys[signer])
			publicKeys = append(publicKeys, allPublicKeys[signer])
			sigs = append(sigs, allSigs[signer])
		}

		aggPublicKey := BLSAggregatePublicKeys(publicKeys)
		aggSig := BLSAggregateSignatures(sigs)
		assert.Check(t, BLSVerify(aggSig, aggPublicKey, msg))
		assert.Check(t, !BLSVerify(aggSig, aggPublicKey, []byte("good bye")))
		assert.Check(t, !BLSVerify(aggSig, allPublicKeys[0], msg))
		assert.Check(t, !BLSVerify(allSigs[0], aggPublicKey, msg))
	}
}
