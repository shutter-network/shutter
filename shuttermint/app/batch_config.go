package app

import "github.com/ethereum/go-ethereum/common"

// IsKeyper checks if the given address is a keyper
func (bc *BatchConfig) IsKeyper(candidate common.Address) bool {
	for _, k := range bc.Keypers {
		if k == candidate {
			return true
		}
	}
	return false
}
