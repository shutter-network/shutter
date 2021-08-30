// Package shtest contains utility functions for use in our tests
package shtest

import (
	"bytes"
	"encoding/gob"
	"math/big"
	"testing"

	gocmp "github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

var BigIntComparer = gocmp.Comparer(func(x, y *big.Int) bool {
	return x.Cmp(y) == 0
})

func EnsureGobable(t *testing.T, src, dst interface{}, opts ...gocmp.Option) {
	t.Helper()
	buff := bytes.Buffer{}
	err := gob.NewEncoder(&buff).Encode(src)
	assert.NilError(t, err)
	err = gob.NewDecoder(&buff).Decode(dst)
	assert.NilError(t, err)
	assert.DeepEqual(t, src, dst, opts...)
}
