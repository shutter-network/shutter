package app

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// AddPublicKeyCommitment adds a PublicKeyCommitment to the batch. The PublicKeyCommitment must be
// sent from configured Keyper
func (bs *BatchState) AddPublicKeyCommitment(commitment PublicKeyCommitment) error {
	if !bs.Config.IsKeyper(commitment.Sender) {
		return errors.New("not a keyper")
	}

	for _, comm := range bs.Commitments {
		if comm.Sender == commitment.Sender {
			return errors.New("already have commitment")
		}
	}

	pubkey, err := crypto.UnmarshalPubkey(commitment.Pubkey)
	if err != nil {
		return err
	}

	bs.Commitments = append(bs.Commitments, commitment)
	if len(bs.Commitments) == int(bs.Config.Threshold) {
		bs.PublicKey = pubkey
	}

	return nil
}

// FindPublicKeyCommitment returns the PublicKeyCommitment provided by the given address
func (bs *BatchState) FindPublicKeyCommitment(addr common.Address) (PublicKeyCommitment, error) {
	for _, comm := range bs.Commitments {
		if comm.Sender == addr {
			return comm, nil
		}
	}
	return PublicKeyCommitment{}, errors.New("keyper did not provide a public key")
}

// FindSecretShare returns the SecretShare provided by the given address
func (bs *BatchState) FindSecretShare(addr common.Address) (SecretShare, error) {
	for _, s := range bs.SecretShares {
		if s.Sender == addr {
			return s, nil
		}
	}
	return SecretShare{}, errors.New("keyper did not provide a secret share")
}

// AddSecretShare adds a SecretShare to the batch.
func (bs *BatchState) AddSecretShare(share SecretShare) error {
	if !bs.Config.IsKeyper(share.Sender) {
		return errors.New("not a keyper")
	}
	pkc, err := bs.FindPublicKeyCommitment(share.Sender)
	if err != nil {
		return err
	}

	if _, err = bs.FindSecretShare(share.Sender); err == nil {
		return errors.New("secret share already sent for this batch")
	}

	// Check that the secret key matches the public key. As a side effect this also makes sure
	// that share.Privkey is a valid private key
	privkey, err := crypto.ToECDSA(share.Privkey)
	if err != nil {
		return err
	}

	if !bytes.Equal(crypto.FromECDSAPub(&privkey.PublicKey), pkc.Pubkey) {
		return errors.New("keys do not match")
	}
	bs.SecretShares = append(bs.SecretShares, share)

	if bs.PrivateKey == nil &&
		bs.PublicKey != nil &&
		len(bs.SecretShares) >= int(bs.Config.Threshold) {
		ss, err := bs.FindSecretShare(bs.Commitments[bs.Config.Threshold-1].Sender)
		if err == nil {
			privkey, err := crypto.ToECDSA(ss.Privkey)
			if err != nil {
				// If we end up here something is seriously wrong, because we call
				// ToECDSA on all shares we add
				panic(err)
			}

			bs.PrivateKey = privkey
		}
	}
	return nil
}

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

	if !bytes.Equal(crypto.FromECDSAPub(bs.PublicKey), a.EncryptionKey) {
		return errors.New("encryption key does not match")
	}

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
