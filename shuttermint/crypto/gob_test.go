package crypto

import (
	"bytes"
	"encoding/gob"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/stretchr/testify/require"
)

func ensureGobable(t *testing.T, src, dst interface{}) {
	buff := bytes.Buffer{}
	err := gob.NewEncoder(&buff).Encode(src)
	require.Nil(t, err)
	err = gob.NewDecoder(&buff).Decode(dst)
	require.Nil(t, err)
	require.Equal(t, src, dst)
}

func TestEonSecretKeyShareGobable(t *testing.T) {
	share := (*EonSecretKeyShare)(big.NewInt(1111))
	ensureGobable(t, share, new(EonSecretKeyShare))
}

func TestEonPublicKeyShareGobable(t *testing.T) {
	share := (*EonPublicKeyShare)(new(bn256.G2).ScalarBaseMult(big.NewInt(5)))
	ensureGobable(t, share, new(EonPublicKeyShare))
}

func TestEonPublicKeyGobable(t *testing.T) {
	pubkey := (*EonPublicKey)(new(bn256.G2).ScalarBaseMult(big.NewInt(5)))
	ensureGobable(t, pubkey, new(EonPublicKey))
}

func TestEpochIDGobable(t *testing.T) {
	epochid := ComputeEpochID(1111)
	ensureGobable(t, epochid, new(EpochID))
}

func TestEpochSecretKeyShareGobable(t *testing.T) {
	share := (*EpochSecretKeyShare)(new(bn256.G1).ScalarBaseMult(big.NewInt(1111)))
	ensureGobable(t, share, new(EpochSecretKeyShare))
}

func TestEpochSecretKeyGobable(t *testing.T) {
	key := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(1111)))
	ensureGobable(t, key, new(EpochSecretKey))
}
