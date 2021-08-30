package app

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shlib/shcrypto"
	"github.com/shutter-network/shutter/shuttermint/medley"
	"github.com/shutter-network/shutter/shuttermint/shmsg"
)

func validateAddress(address []byte) (common.Address, error) {
	if len(address) != common.AddressLength {
		return common.Address{}, errors.Errorf(
			"address has invalid length (%d instead of %d bytes)",
			len(address),
			common.AddressLength,
		)
	}

	return common.BytesToAddress(address), nil
}

// ParsePolyEvalMsg converts a shmsg.PolyEvalMsg to an app.PolyEvalMsg.
func ParsePolyEvalMsg(msg *shmsg.PolyEval, sender common.Address) (*PolyEval, error) {
	if len(msg.Receivers) != len(msg.EncryptedEvals) {
		return nil, errors.Errorf("number of receivers %d does not match number of evals %d", len(msg.Receivers), len(msg.EncryptedEvals))
	}

	receivers := []common.Address{}
	for _, receiver := range msg.Receivers {
		address, err := validateAddress(receiver)
		if err != nil {
			return nil, err
		}
		receivers = append(receivers, address)
	}

	if err := medley.EnsureUniqueAddresses(receivers); err != nil {
		return nil, err
	}

	return &PolyEval{
		Sender:         sender,
		Eon:            msg.Eon,
		Receivers:      receivers,
		EncryptedEvals: msg.EncryptedEvals,
	}, nil
}

// ParsePolyCommitmentMsg converts a shmsg.PolyCommitmentMsg to an app.PolyCommitmentMsg.
func ParsePolyCommitmentMsg(msg *shmsg.PolyCommitment, sender common.Address) (*PolyCommitment, error) {
	gammas := shcrypto.Gammas{}
	for _, g := range msg.Gammas {
		g2 := new(bn256.G2)
		_, err := g2.Unmarshal(g)
		if err != nil {
			return nil, err
		}
		gammas = append(gammas, g2)
	}
	return &PolyCommitment{
		Sender: sender,
		Eon:    msg.Eon,
		Gammas: &gammas,
	}, nil
}

// ParseAccusationMsg converts a shmsg.AccusationMsg to an app.AccusationMsg.
func ParseAccusationMsg(msg *shmsg.Accusation, sender common.Address) (*Accusation, error) {
	accused := []common.Address{}
	for _, acc := range msg.Accused {
		address, err := validateAddress(acc)
		if err != nil {
			return nil, err
		}
		accused = append(accused, address)
	}

	if err := medley.EnsureUniqueAddresses(accused); err != nil {
		return nil, err
	}

	return &Accusation{
		Sender:  sender,
		Eon:     msg.Eon,
		Accused: accused,
	}, nil
}

// ParseApologyMsg converts a shmsg.ApologyMsg to an app.ApologyMsg.
func ParseApologyMsg(msg *shmsg.Apology, sender common.Address) (*Apology, error) {
	if len(msg.Accusers) != len(msg.PolyEvals) {
		return nil, errors.Errorf("number of accusers %d and apology evals %d not equal", len(msg.Accusers), len(msg.PolyEvals))
	}

	accusers := []common.Address{}

	for _, acc := range msg.Accusers {
		accuser, err := validateAddress(acc)
		if err != nil {
			return nil, err
		}

		accusers = append(accusers, accuser)
	}

	if err := medley.EnsureUniqueAddresses(accusers); err != nil {
		return nil, err
	}

	var polyEval []*big.Int
	for _, b := range msg.PolyEvals {
		e := new(big.Int)
		e.SetBytes(b)
		polyEval = append(polyEval, e)
	}

	return &Apology{
		Sender:   sender,
		Eon:      msg.Eon,
		Accusers: accusers,
		PolyEval: polyEval,
	}, nil
}

// ParseEpochSecretKeyShareMsg converts a shmsg.EpochSecretKeyShareMsg to an app.EpochSecretShareMsg.
func ParseEpochSecretKeyShareMsg(msg *shmsg.EpochSecretKeyShare, sender common.Address) (*EpochSecretKeyShare, error) {
	share := new(shcrypto.EpochSecretKeyShare)
	err := share.GobDecode(msg.Share)
	if err != nil {
		return nil, err
	}
	return &EpochSecretKeyShare{
		Sender: sender,
		Eon:    msg.Eon,
		Epoch:  msg.Epoch,
		Share:  share,
	}, nil
}
