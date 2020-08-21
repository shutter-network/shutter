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

	pubkey, err := crypto.UnmarshalPubkey(commitment.Pubkey)
	if err != nil {
		return err
	}

	bk.Commitments = append(bk.Commitments, commitment)
	if len(bk.Commitments) == int(bk.Config.Threshhold) {
		bk.PublicKey = pubkey
	}

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

// FindSecretShare returns the SecretShare provided by the given address
func (bk *BatchKeys) FindSecretShare(addr common.Address) (SecretShare, error) {
	for _, s := range bk.SecretShares {
		if s.Sender == addr {
			return s, nil
		}
	}
	return SecretShare{}, errors.New("Keyper did not provide a secret share")
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

	if _, err = bk.FindSecretShare(share.Sender); err == nil {
		return errors.New("SecretShare already sent for this batch")
	}

	// Check that the secret key matches the public key. As a side effect this also makes sure
	// that share.Privkey is a valid private key
	privkey, err := crypto.ToECDSA(share.Privkey)
	if err != nil {
		return err
	}

	if !bytes.Equal(crypto.FromECDSAPub(&privkey.PublicKey), pkc.Pubkey) {
		return errors.New("Keys do not match")
	}
	bk.SecretShares = append(bk.SecretShares, share)

	if bk.PrivateKey == nil &&
		bk.PublicKey != nil &&
		len(bk.SecretShares) >= int(bk.Config.Threshhold) {
		ss, err := bk.FindSecretShare(bk.Commitments[bk.Config.Threshhold-1].Sender)
		if err == nil {
			privkey, err := crypto.ToECDSA(ss.Privkey)
			if err != nil {
				// If we end up here something is seriously wrong, because we call
				// ToECDSA on all shares we add
				panic(err)
			}

			bk.PrivateKey = privkey
		}
	}
	return nil
}
