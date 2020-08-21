package app

import (
	"bytes"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// AddPublicKeyCommitment adds a PublicKeyCommitment to the batch. The PublicKeyCommitment must be
// sent from configured Keyper
func (bk *BatchKeys) AddPublicKeyCommitment(commitment PublicKeyCommitment) error {
	if !bk.Config.IsKeyper(commitment.Sender) {
		return errors.New("Not a keyper")
	}

	for _, comm := range bk.Commitments {
		if comm.Sender == commitment.Sender {
			return errors.New("Already have commitment")
		}
	}
	bk.Commitments = append(bk.Commitments, commitment)
	return nil
}

// FindPublicKeyCommitment returns the PublicKeyCommitment provided by the given address
func (bk *BatchKeys) FindPublicKeyCommitment(addr common.Address) (PublicKeyCommitment, error) {
	for _, comm := range bk.Commitments {
		if comm.Sender == addr {
			return comm, nil
		}
	}
	return PublicKeyCommitment{}, errors.New("Keyper did not provide a public key")
}

// AddSecretShare adds a SecretShare to the batch.
func (bk *BatchKeys) AddSecretShare(share SecretShare) error {
	if !bk.Config.IsKeyper(share.Sender) {
		return errors.New("Not a keyper")
	}
	pkc, err := bk.FindPublicKeyCommitment(share.Sender)
	if err != nil {
		return err
	}
	for _, s := range bk.SecretShares {
		if s.Sender == share.Sender {
			return errors.New("SecretShare already sent for this batch")
		}
	}

	// check that the secret key matches the public key
	privkey, err := crypto.ToECDSA(share.Privkey)
	if err != nil {
		return err
	}

	if !bytes.Equal(crypto.FromECDSAPub(&privkey.PublicKey), pkc.Pubkey) {
		return errors.New("pubkeys do not match")
	}
	bk.SecretShares = append(bk.SecretShares, share)
	return nil
}
