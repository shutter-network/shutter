package app

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	keys      [10]*ecdsa.PrivateKey
	addresses [10]common.Address
)

func init() {
	for i := 0; i < 10; i++ {
		var d [32]byte
		d[31] = byte(i + 1)
		k, err := crypto.ToECDSA(d[:])
		if err != nil {
			panic(err)
		}
		keys[i] = k
		addresses[i] = crypto.PubkeyToAddress(k.PublicKey)
	}
}
