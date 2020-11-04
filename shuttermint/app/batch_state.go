package app

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// FindEncryptionKeyAttestation returns the EncryptionKeyAttestation provided by the given address
func (bs *BatchState) FindEncryptionKeyAttestation(addr common.Address) (EncryptionKeyAttestation, error) {
	for _, a := range bs.EncryptionKeyAttestations {
		if a.Sender == addr {
			return a, nil
		}
	}
	return EncryptionKeyAttestation{}, errors.New("keyper did not provide an encryption key attestation")
}

// AddEncryptionKeyAttestation adds an EncryptionKeyAttestation to the batch.
func (bs *BatchState) AddEncryptionKeyAttestation(a EncryptionKeyAttestation) error {
	if a.ConfigContractAddress != bs.Config.ConfigContractAddress {
		return fmt.Errorf(
			"wrong config contract address %s instead of %s",
			a.ConfigContractAddress.Hex(),
			bs.Config.ConfigContractAddress.Hex(),
		)
	}

	if !bs.Config.IsKeyper(a.Sender) {
		return errors.New("not a keyper")
	}

	// if !bytes.Equal(crypto.FromECDSAPub(bs.PublicKey), a.EncryptionKey) {
	//	return errors.New("encryption key does not match")
	// }

	if !a.VerifySignature() {
		return errors.New("invalid signature")
	}

	for _, att := range bs.EncryptionKeyAttestations {
		if att.Sender == a.Sender {
			return errors.New("already have encyption key attestation")
		}
	}

	bs.EncryptionKeyAttestations = append(bs.EncryptionKeyAttestations, a)

	return nil
}

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
