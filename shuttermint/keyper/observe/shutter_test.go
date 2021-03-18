package observe

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/internal/shtest"
)

// encryptionPublicKey generates an EncryptionPublicKey
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
