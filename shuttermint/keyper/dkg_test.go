package keyper

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

func TestNewDKGInstance(t *testing.T) {
	eon := uint64(10)
	config := contract.BatchConfig{
		Threshold: 2,
		Keypers: []common.Address{
			common.BigToAddress(big.NewInt(0xaa)),
			common.BigToAddress(big.NewInt(0xbb)),
			common.BigToAddress(big.NewInt(0xcc)),
		},
	}
	dkg, err := NewDKGInstance(eon, config, nil)
	require.Nil(t, err)
	require.Equal(t, eon, dkg.Eon)
	require.NotNil(t, dkg.Polynomial)
}
