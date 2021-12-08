package main

import (
	"fmt"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

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
func verifyDecryptionKey(decryptionKey []byte, eonPublicKey []byte, epochID uint64) (bool, error) {
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

	return shcrypto.VerifyEpochSecretKey(epochSecretKey, eonPublicKeyPoint, epochID)
}
