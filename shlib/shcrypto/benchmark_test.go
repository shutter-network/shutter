package shcrypto

import (
	"crypto/rand"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"gotest.tools/v3/assert"
)

func setup(b *testing.B) (*EonPublicKey, *EpochID, *EpochSecretKey) {
	b.Helper()
	// first generate keys
	n := 3
	threshold := uint64(2)
	epochID := ComputeEpochID([]byte("epoch1"))

	ps := []*Polynomial{}
	gammas := []*Gammas{}
	for i := 0; i < n; i++ {
		p, err := RandomPolynomial(rand.Reader, threshold-1)
		assert.NilError(b, err)
		ps = append(ps, p)
		gammas = append(gammas, p.Gammas())
	}

	eonSecretKeyShares := []*EonSecretKeyShare{}
	epochSecretKeyShares := []*EpochSecretKeyShare{}
	eonSecretKey := big.NewInt(0)
	for i := 0; i < n; i++ {
		eonSecretKey.Add(eonSecretKey, ps[i].Eval(big.NewInt(0)))

		ss := []*big.Int{}
		for j := 0; j < n; j++ {
			s := ps[j].EvalForKeyper(i)
			ss = append(ss, s)
		}
		eonSecretKeyShares = append(eonSecretKeyShares, ComputeEonSecretKeyShare(ss))
		_ = ComputeEonPublicKeyShare(i, gammas)
		epochSecretKeyShares = append(epochSecretKeyShares, ComputeEpochSecretKeyShare(eonSecretKeyShares[i], epochID))
	}
	eonPublicKey := ComputeEonPublicKey(gammas)
	assert.DeepEqual(b, new(bn256.G2).ScalarBaseMult(eonSecretKey), (*bn256.G2)(eonPublicKey), g2Comparer)
	epochSecretKey, err := ComputeEpochSecretKey(
		[]int{0, 1},
		[]*EpochSecretKeyShare{epochSecretKeyShares[0], epochSecretKeyShares[1]},
		threshold)
	assert.NilError(b, err)

	return eonPublicKey, epochID, epochSecretKey
}

func roundtrip(msg []byte, eonPub *EonPublicKey, epochID *EpochID, epochSecret *EpochSecretKey, sigma Block) error {
	encM := Encrypt(msg, eonPub, epochID, sigma)
	_, err := encM.Decrypt(epochSecret)
	return err
}

func marshalRoundtrip(content *EncryptedMessage, result *EncryptedMessage) (*EncryptedMessage, error) {
	m := content.Marshal()
	err := result.Unmarshal(m)
	return result, err
}

func BenchmarkMarshal(b *testing.B) {
	b.StopTimer()
	eonPub, epochID, _ := setup(b)
	sigmas := []Block{}
	msg := []byte("hello")
	// now marshal and unmarshal message
	for i := 0; i < b.N; i++ {
		sigma, err := RandomSigma(rand.Reader)
		sigmas = append(sigmas, sigma)
		assert.NilError(b, err)
	}
	content := Encrypt(msg, eonPub, epochID, sigmas[0])
	content_hex, err := content.MarshalText()
	assert.NilError(b, err)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		template := &EncryptedMessage{}
		result, err := marshalRoundtrip(content, template)
		assert.NilError(b, err)

		result_hex, err := result.MarshalText()
		assert.NilError(b, err)
		assert.Equal(b, string(content_hex), string(result_hex))
	}
}

func BenchmarkRoundTrip(b *testing.B) {
	b.StopTimer()
	eonPub, epochID, epochSecretKey := setup(b)
	sigmas := []Block{}
	msg := []byte("hello")
	// now encrypt and decrypt message
	for i := 0; i < b.N; i++ {
		sigma, err := RandomSigma(rand.Reader)
		sigmas = append(sigmas, sigma)
		assert.NilError(b, err)
	}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := roundtrip(msg, eonPub, epochID, epochSecretKey, sigmas[i])
		assert.NilError(b, err)
	}
}
