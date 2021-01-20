// Package shtest contains utility functions for use in our tests
package shtest

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/stretchr/testify/require"
)

func EnsureGobable(t *testing.T, src, dst interface{}) {
	buff := bytes.Buffer{}
	err := gob.NewEncoder(&buff).Encode(src)
	require.Nil(t, err)
	err = gob.NewDecoder(&buff).Decode(dst)
	require.Nil(t, err)
	require.Equal(t, src, dst)
}
