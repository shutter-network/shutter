package keyper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDKGInstance(t *testing.T) {
	eon := uint64(10)
	dkg, err := NewDKGInstance(eon)
	require.Nil(t, err)
	require.Equal(t, eon, dkg.Eon)
	require.NotNil(t, dkg.PolyBase)
}
