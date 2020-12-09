// package medley provides some functions that may be useful in various parts of shutter
package medley

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

var errAddressNotFound = errors.New("address not found")

// FindAddressIndex returns the index of the given address inside the slice of addresses or returns
// an error, if the slice does not contain the given address
func FindAddressIndex(addresses []common.Address, addr common.Address) (int, error) {
	for i, a := range addresses {
		if a == addr {
			return i, nil
		}
	}
	return -1, errAddressNotFound
}
