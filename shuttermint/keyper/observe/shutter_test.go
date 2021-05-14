package observe

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/internal/shtest"
	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
)

// encryptionPublicKey generates an EncryptionPublicKey.
func encryptionPublicKey(t *testing.T) *EncryptionPublicKey {
	t.Helper()
	privkey, err := crypto.GenerateKey()
	require.Nil(t, err)
	return (*EncryptionPublicKey)(ecies.ImportECDSAPublic(&privkey.PublicKey))
}

// TestGobSerializationIssue45 tests that we can serialize the encryption public key, see
// https://github.com/brainbot-com/shutter/issues/45
func TestGobSerializationIssue45(t *testing.T) {
	sh := NewShutter()
	sh.KeyperEncryptionKeys[common.Address{}] = encryptionPublicKey(t)
	shtest.EnsureGobable(t, sh, new(Shutter))
}

func TestFindBatchConfigByBatchIndex(t *testing.T) {
	sh := NewShutter()

	sh.BatchConfigs = append(sh.BatchConfigs,
		shutterevents.BatchConfig{
			Height:          1,
			StartBatchIndex: 5,
		},
		shutterevents.BatchConfig{
			Height:          2,
			StartBatchIndex: 10,
		},
	)

	require.Equal(t, int64(0), sh.FindBatchConfigByBatchIndex(0).Height)
	require.Equal(t, int64(0), sh.FindBatchConfigByBatchIndex(4).Height)
	require.Equal(t, int64(1), sh.FindBatchConfigByBatchIndex(5).Height)
	require.Equal(t, int64(1), sh.FindBatchConfigByBatchIndex(9).Height)
	require.Equal(t, int64(2), sh.FindBatchConfigByBatchIndex(10).Height)
	require.Equal(t, int64(2), sh.FindBatchConfigByBatchIndex(11).Height)
}
