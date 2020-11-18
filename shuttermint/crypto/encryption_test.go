package crypto

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	// first generate keys
	n := 3
	threshold := uint64(2)
	epochID := ComputeEpochID(uint64(10))
	epochID = (*EpochID)(new(bn256.G1).ScalarBaseMult(big.NewInt(1)))

	ps := []*Polynomial{}
	gammas := []*Gammas{}
	for i := 0; i < n; i++ {
		p, err := RandomPolynomial(rand.Reader, threshold-1)
		require.Nil(t, err)
		ps = append(ps, p)
		gammas = append(gammas, p.Gammas())
	}

	eonSKShares := []*EonSKShare{}
	eonPKShares := []*EonPKShare{}
	epochSKShares := []*EpochSKShare{}
	eonSK := big.NewInt(0)
	for i := 0; i < n; i++ {
		eonSK.Add(eonSK, ps[i].Eval(big.NewInt(0)))

		ss := []*big.Int{}
		for j := 0; j < n; j++ {
			s := ps[j].EvalForKeyper(i)
			ss = append(ss, s)
		}
		eonSKShares = append(eonSKShares, ComputeEonSKShare(ss))
		eonPKShares = append(eonPKShares, ComputeEonPKShare(i, gammas))
		epochSKShares = append(epochSKShares, ComputeEpochSKShare(eonSKShares[i], epochID))
	}
	eonPK := ComputeEonPK(gammas)
	require.True(t, EqualG2(new(bn256.G2).ScalarBaseMult(eonSK), (*bn256.G2)(eonPK)))
	epochSK, err := ComputeEpochSK([]int{0, 1}, []*EpochSKShare{epochSKShares[0], epochSKShares[1]}, threshold)
	require.Nil(t, err)

	// now encrypt and decrypt message
	m := []byte("hello")
	r, err := RandomR(rand.Reader)
	require.Nil(t, err)

	c := Encrypt(m, r, eonPK, epochID)
	c1Exp := new(bn256.G2).ScalarBaseMult(r)
	require.True(t, EqualG2(c.RandomnessWitness, c1Exp))

	encKey := computeEncryptionKey(r, eonPK, epochID)
	decKey := computeDecryptionKey(epochSK, c)
	require.Equal(t, encKey, decKey)

	m2 := Decrypt(c, epochSK)
	require.True(t, bytes.Equal(m, m2))
}
