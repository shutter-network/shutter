package main

import (
	"encoding/binary"
	"fmt"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

func main() {}

//export decrypt
func decrypt(encryptedMessage []byte, decryptionKey []byte) ([]byte, error) {
	encrypted := new(shcrypto.EncryptedMessage)
	err := encrypted.Unmarshal(encryptedMessage)
	if err != nil {
		return nil, fmt.Errorf("invalid encrypted message: %s", err)
	}

	epochSecretKey := new(shcrypto.EpochSecretKey)
	err = epochSecretKey.Unmarshal(decryptionKey)
	if err != nil {
		return nil, fmt.Errorf("invalid decryption key: %s", err)
	}

	message, err := encrypted.Decrypt(epochSecretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt message: %s", err)
	}

	return message, nil
}

//export verifyDecryptionKey
func verifyDecryptionKey(decryptionKey []byte, eonPublicKey []byte, epochID []byte) (bool, error) {
	epochSecretKey := new(shcrypto.EpochSecretKey)
	err := epochSecretKey.Unmarshal(decryptionKey)
	if err != nil {
		return false, fmt.Errorf("invalid decryption key: %s", err)
	}

	eonPublicKeyPoint := new(shcrypto.EonPublicKey)
	err = eonPublicKeyPoint.Unmarshal(eonPublicKey)
	if err != nil {
		return false, fmt.Errorf("invalid eon public key: %s", err)
	}

	if len(epochID) != 8 {
		return false, fmt.Errorf("epoch id must be 8 bytes, got %d", len(epochID))
	}
	epochIDInt := binary.BigEndian.Uint64(epochID)

	return shcrypto.VerifyEpochSecretKey(epochSecretKey, eonPublicKeyPoint, epochIDInt)
}
