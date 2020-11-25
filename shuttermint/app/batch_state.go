package app

import (
	"fmt"
)

// AddDecryptionSignature adds a decryption signature to the batch
func (bs *BatchState) AddDecryptionSignature(ds DecryptionSignature) error {
	if !bs.Config.IsKeyper(ds.Sender) {
		return fmt.Errorf("sender %s is not a keyper", ds.Sender.Hex())
	}

	for _, sig := range bs.DecryptionSignatures {
		if sig.Sender == ds.Sender {
			return fmt.Errorf("already have decryption signature from %s", ds.Sender.Hex())
		}
	}

	bs.DecryptionSignatures = append(bs.DecryptionSignatures, ds)

	return nil
}
