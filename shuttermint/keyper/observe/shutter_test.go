package observe

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	gocmp "github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shtest"
	"github.com/shutter-network/shutter/shuttermint/keyper/shutterevents"
)

// encryptionPublicKey generates an EncryptionPublicKey.
func encryptionPublicKey(t *testing.T) *EncryptionPublicKey {
	t.Helper()
	privkey, err := crypto.GenerateKey()
	assert.NilError(t, err)
	return (*EncryptionPublicKey)(ecies.ImportECDSAPublic(&privkey.PublicKey))
}

var encryptionPublicKeyComparer = gocmp.Comparer(func(x, y *EncryptionPublicKey) bool {
	return reflect.DeepEqual(x, y)
})

// TestGobSerializationIssue45 tests that we can serialize the encryption public key, see
// https://github.com/shutter-network/shutter/issues/45
func TestGobSerializationIssue45(t *testing.T) {
	sh := NewShutter()
	epk := encryptionPublicKey(t)
	sh.KeyperEncryptionKeys[common.Address{}] = epk
	shtest.EnsureGobable(t, sh, new(Shutter), shtest.BigIntComparer, encryptionPublicKeyComparer)
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

	assert.Equal(t, int64(0), sh.FindBatchConfigByBatchIndex(0).Height)
	assert.Equal(t, int64(0), sh.FindBatchConfigByBatchIndex(4).Height)
	assert.Equal(t, int64(1), sh.FindBatchConfigByBatchIndex(5).Height)
	assert.Equal(t, int64(1), sh.FindBatchConfigByBatchIndex(9).Height)
	assert.Equal(t, int64(2), sh.FindBatchConfigByBatchIndex(10).Height)
	assert.Equal(t, int64(2), sh.FindBatchConfigByBatchIndex(11).Height)
}
