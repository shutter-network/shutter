package shcrypto

import (
	"math/big"
	"testing"

	"github.com/shutter-network/shutter/shlib/shtest"
)

func TestEonSecretKeyShareGobable(t *testing.T) {
	share := (*EonSecretKeyShare)(big.NewInt(1111))
	shtest.EnsureGobable(t, share, new(EonSecretKeyShare), shtest.BigIntComparer)
}

func TestEonPublicKeyShareGobable(t *testing.T) {
	share := (*EonPublicKeyShare)(makeTestG2(5))
	shtest.EnsureGobable(t, share, new(EonPublicKeyShare))
}

func TestEonPublicKeyGobable(t *testing.T) {
	pubkey := (*EonPublicKey)(makeTestG2(5))
	shtest.EnsureGobable(t, pubkey, new(EonPublicKey))
}

func TestEpochIDGobable(t *testing.T) {
	epochid := ComputeEpochID([]byte("epoch1"))
	shtest.EnsureGobable(t, epochid, new(EpochID))
}

func TestEpochSecretKeyShareGobable(t *testing.T) {
	share := (*EpochSecretKeyShare)(makeTestG1(1111))
	shtest.EnsureGobable(t, share, new(EpochSecretKeyShare))
}

func TestEpochSecretKeyGobable(t *testing.T) {
	key := (*EpochSecretKey)(makeTestG1(1111))
	shtest.EnsureGobable(t, key, new(EpochSecretKey), g1Comparer)
}
