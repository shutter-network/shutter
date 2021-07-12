package shcrypto

import (
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"

	"github.com/shutter-network/shutter/shuttermint/internal/shtest"
)

func TestEonSecretKeyShareGobable(t *testing.T) {
	share := (*EonSecretKeyShare)(big.NewInt(1111))
	shtest.EnsureGobable(t, share, new(EonSecretKeyShare), shtest.BigIntComparer)
}

func TestEonPublicKeyShareGobable(t *testing.T) {
	share := (*EonPublicKeyShare)(new(bn256.G2).ScalarBaseMult(big.NewInt(5)))
	shtest.EnsureGobable(t, share, new(EonPublicKeyShare))
}

func TestEonPublicKeyGobable(t *testing.T) {
	pubkey := (*EonPublicKey)(new(bn256.G2).ScalarBaseMult(big.NewInt(5)))
	shtest.EnsureGobable(t, pubkey, new(EonPublicKey))
}

func TestEpochIDGobable(t *testing.T) {
	epochid := ComputeEpochID(1111)
	shtest.EnsureGobable(t, epochid, new(EpochID))
}

func TestEpochSecretKeyShareGobable(t *testing.T) {
	share := (*EpochSecretKeyShare)(new(bn256.G1).ScalarBaseMult(big.NewInt(1111)))
	shtest.EnsureGobable(t, share, new(EpochSecretKeyShare))
}

func TestEpochSecretKeyGobable(t *testing.T) {
	key := (*EpochSecretKey)(new(bn256.G1).ScalarBaseMult(big.NewInt(1111)))
	shtest.EnsureGobable(t, key, new(EpochSecretKey), G1Comparer)
}
