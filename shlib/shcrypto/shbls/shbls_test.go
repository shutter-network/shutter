package shbls

import (
	"crypto/rand"
	"testing"

	"gotest.tools/assert"
)

func TestSigning(t *testing.T) {
	secretKey, publicKey, err := RandomKeyPair(rand.Reader)
	assert.NilError(t, err)

	msg := []byte("hello")
	sig := Sign(msg, secretKey)
	assert.Check(t, Verify(sig, publicKey, msg))

	secretKey2, publicKey2, err := RandomKeyPair(rand.Reader)
	assert.NilError(t, err)
	sig2 := Sign(msg, secretKey2)

	assert.Check(t, !Verify(sig, publicKey, []byte("good bye")))
	assert.Check(t, !Verify(sig, publicKey2, msg))
	assert.Check(t, !Verify(sig2, publicKey, msg))
}

func TestAggregation(t *testing.T) {
	msg := []byte("hello")

	allsecretKeys := []*SecretKey{}
	allPublicKeys := []*PublicKey{}
	allSigs := []*Signature{}
	for i := 0; i < 3; i++ {
		secretKey, publicKey, err := RandomKeyPair(rand.Reader)
		assert.NilError(t, err)
		sig := Sign(msg, secretKey)

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
		secretKeys := []*SecretKey{}
		publicKeys := []*PublicKey{}
		sigs := []*Signature{}
		for _, signer := range signers {
			secretKeys = append(secretKeys, allsecretKeys[signer])
			publicKeys = append(publicKeys, allPublicKeys[signer])
			sigs = append(sigs, allSigs[signer])
		}

		aggPublicKey := AggregatePublicKeys(publicKeys)
		aggSig := AggregateSignatures(sigs)
		assert.Check(t, Verify(aggSig, aggPublicKey, msg))
		assert.Check(t, !Verify(aggSig, aggPublicKey, []byte("good bye")))
		assert.Check(t, !Verify(aggSig, allPublicKeys[0], msg))
		assert.Check(t, !Verify(allSigs[0], aggPublicKey, msg))
	}
}
