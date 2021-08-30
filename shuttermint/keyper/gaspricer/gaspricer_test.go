package gaspricer

import (
	"math/big"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shtest"
)

func TestGasPricerAdjust(t *testing.T) {
	old := gasPriceMultiplier
	defer func() {
		gasPriceMultiplier = old
	}()

	_ = SetMultiplier(1.0)
	price := Adjust(big.NewInt(1e18))
	assert.DeepEqual(t, price, big.NewInt(1e18), shtest.BigIntComparer)

	_ = SetMultiplier(2.0)
	price = Adjust(big.NewInt(1e18))
	assert.DeepEqual(t, price, big.NewInt(2e18), shtest.BigIntComparer)

	_ = SetMultiplier(2.5)
	price = Adjust(big.NewInt(1e18))
	assert.DeepEqual(t, price, big.NewInt(2.5e18), shtest.BigIntComparer)

	err := SetMultiplier(-1.0)
	assert.Error(t, err, "gas price multiplier must be non-negative")
}
