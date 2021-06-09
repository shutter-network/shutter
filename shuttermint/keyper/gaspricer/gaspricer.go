// Package gaspricer is used to multiply the gas price by a configurable factor. This is needed
// because the give price returned from SuggestGasPrice is too low (at least on goerli)
package gaspricer

import (
	"math/big"

	"github.com/pkg/errors"
)

var gasPriceMultiplier = big.NewFloat(1.5)

// SetMultiplier sets the gas price multiplier. This is a global setting.
func SetMultiplier(f float64) error {
	if f < 0.0 {
		return errors.New("gas price multiplier must be non-negative")
	}
	gasPriceMultiplier = big.NewFloat(f)
	return nil
}

// Adjust multiplies the given gas price by the configured multiplier.
// place.
func Adjust(price *big.Int) *big.Int {
	p := new(big.Float).SetInt(price)
	p.Mul(p, gasPriceMultiplier)
	r, _ := p.Int(nil)
	return r
}
