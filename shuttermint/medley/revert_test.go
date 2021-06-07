package medley

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"gotest.tools/v3/assert"
)

func TestUnpackErrorEmpty(t *testing.T) {
	err := unpackError([]byte{0x01})
	assert.ErrorContains(t, err, "TX result not of type Error(string)")
}

func TestUnpackError(t *testing.T) {
	packed, err := abi.Arguments{{Type: abiString}}.Pack("foo bar baz")
	assert.NilError(t, err)
	err = unpackError(bytes.Join([][]byte{errorSig, packed}, []byte{}))
	assert.ErrorContains(t, err, "foo bar baz")
}

func TestUnpackErrorOnlySig(t *testing.T) {
	err := unpackError(errorSig)
	assert.ErrorContains(t, err, "unpack revert reason")
}
