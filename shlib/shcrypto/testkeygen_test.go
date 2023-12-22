package shcrypto

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestTestKeyGen(t *testing.T) {
	g, err := NewTestKeyGen()
	assert.NilError(t, err)

	epochIDBytes := []byte("epochid")
	epochID := ComputeEpochID(epochIDBytes)
	epochSecretKey, err := g.ComputeEpochSecretKey(epochID)
	assert.NilError(t, err)

	ok, err := VerifyEpochSecretKey(epochSecretKey, g.EonPublicKey, epochIDBytes)
	assert.NilError(t, err)
	assert.Assert(t, ok)
}
