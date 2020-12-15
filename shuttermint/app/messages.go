package app

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

func validateAddress(address []byte) (common.Address, error) {
	if len(address) != common.AddressLength {
		return common.Address{}, fmt.Errorf(
			"address has invalid length (%d instead of %d bytes)",
			len(address),
			common.AddressLength,
		)
	}

	return common.BytesToAddress(address), nil
}

// ParsePolyEvalMsg converts a shmsg.PolyEvalMsg to an app.PolyEvalMsg
func ParsePolyEvalMsg(msg *shmsg.PolyEval, sender common.Address) (*PolyEvalMsg, error) {
	if len(msg.Receivers) != len(msg.EncryptedEvals) {
		return nil, fmt.Errorf("number of receivers %d does not match number of evals %d", len(msg.Receivers), len(msg.EncryptedEvals))
	}

	receivers := []common.Address{}
	receiverMap := make(map[common.Address]bool)
	for _, receiver := range msg.Receivers {
		address, err := validateAddress(receiver)
		if err != nil {
			return nil, err
		}

		if receiverMap[address] {
			return nil, fmt.Errorf("duplicate receiver address %s", address.Hex())
		}
		receiverMap[address] = true

		receivers = append(receivers, address)
	}

	return &PolyEvalMsg{
		Sender:         sender,
		Eon:            msg.Eon,
		Receivers:      receivers,
		EncryptedEvals: msg.EncryptedEvals,
	}, nil
}

// ParsePolyCommitmentMsg converts a shmsg.PolyCommitmentMsg to an app.PolyCommitmentMsg
func ParsePolyCommitmentMsg(msg *shmsg.PolyCommitment, sender common.Address) (*PolyCommitmentMsg, error) {
	return &PolyCommitmentMsg{
		Sender: sender,
		Eon:    msg.Eon,
		Gammas: msg.Gammas,
	}, nil
}

// ParseAccusationMsg converts a shmsg.AccusationMsg to an app.AccusationMsg
func ParseAccusationMsg(msg *shmsg.Accusation, sender common.Address) (*AccusationMsg, error) {
	accused := []common.Address{}
	accusedMap := make(map[common.Address]bool)
	for _, acc := range msg.Accused {
		address, err := validateAddress(acc)
		if err != nil {
			return nil, err
		}

		if accusedMap[address] {
			return nil, fmt.Errorf("duplicate accusation from %s against %s", sender.Hex(), address.Hex())
		}
		accusedMap[address] = true

		accused = append(accused, address)
	}

	return &AccusationMsg{
		Sender:  sender,
		Eon:     msg.Eon,
		Accused: accused,
	}, nil
}

// ParseApologyMsg converts a shmsg.ApologyMsg to an app.ApologyMsg
func ParseApologyMsg(msg *shmsg.Apology, sender common.Address) (*ApologyMsg, error) {
	if len(msg.Accusers) != len(msg.PolyEvals) {
		return nil, fmt.Errorf("number of accusers %d and apology evals %d not equal", len(msg.Accusers), len(msg.PolyEvals))
	}

	accusers := []common.Address{}
	accuserMap := make(map[common.Address]bool)
	for _, acc := range msg.Accusers {
		accuser, err := validateAddress(acc)
		if err != nil {
			return nil, err
		}

		if accuserMap[accuser] {
			return nil, fmt.Errorf("duplicate accuser %s", accuser.Hex())
		}
		accuserMap[accuser] = true

		accusers = append(accusers, accuser)
	}

	return &ApologyMsg{
		Sender:    sender,
		Eon:       msg.Eon,
		Accusers:  accusers,
		PolyEvals: msg.PolyEvals,
	}, nil
}

// ParseEpochSKShareMsg converts a shmsg.ESKShareMsg to an app.ESKShareMsg
func ParseEpochSKShareMsg(msg *shmsg.EpochSKShare, sender common.Address) (*EpochSKShareMsg, error) {
	return &EpochSKShareMsg{
		Sender:       sender,
		Eon:          msg.Eon,
		Epoch:        msg.Epoch,
		EpochSKShare: msg.EpochSkShare,
	}, nil
}
