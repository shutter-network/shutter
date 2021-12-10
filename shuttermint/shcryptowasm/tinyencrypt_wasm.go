package main

import (
	"encoding/binary"
	"fmt"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

func main() {}

//export encrypt
func encrypt(message []byte, eonPublicKey []byte, epochID []byte, sigma []byte) ([]byte, error) {
	eonPublicKeyPoint := new(shcrypto.EonPublicKey)
	err := eonPublicKeyPoint.Unmarshal(eonPublicKey)
	if err != nil {
		return nil, fmt.Errorf("invalid eon public key: %s", err)
	}

	if len(epochID) != 8 {
		return nil, fmt.Errorf("epoch id must be 8 bytes, got %d", len(epochID))
	}
	epochIDInt := binary.BigEndian.Uint64(epochID)
	epochIDPoint := shcrypto.ComputeEpochID(epochIDInt)

	if len(sigma) != shcrypto.BlockSize {
		return nil, fmt.Errorf("sigma must be %d bytes, got %d", shcrypto.BlockSize, len(sigma))
	}
	var sigmaBlock shcrypto.Block
	copy(sigmaBlock[:], sigma)

	encryptedMessage := shcrypto.Encrypt(message, eonPublicKeyPoint, epochIDPoint, sigmaBlock)
	return encryptedMessage.Marshal(), nil
}
