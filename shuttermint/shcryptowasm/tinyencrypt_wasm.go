package main

import (
	"fmt"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

//export encrypt
func encrypt(message []byte, eonPublicKey []byte, epochID uint64, sigma []byte) ([]byte, error) {
	eonPublicKeyPoint := new(shcrypto.EonPublicKey)
	err := eonPublicKeyPoint.Unmarshal(eonPublicKey)
	if err != nil {
		return nil, fmt.Errorf("invalid eon public key: %s", err)
	}

	epochIDPoint := shcrypto.ComputeEpochID(epochID)

	if len(sigma) != shcrypto.BlockSize {
		return nil, fmt.Errorf("sigma must be %d bytes, got %d", shcrypto.BlockSize, len(sigma))
	}
	var sigmaBlock shcrypto.Block
	copy(sigmaBlock[:], sigma)

	encryptedMessage := shcrypto.Encrypt(message, eonPublicKeyPoint, epochIDPoint, sigmaBlock)
	return encryptedMessage.Marshal(), nil
}
