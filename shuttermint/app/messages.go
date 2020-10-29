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
func ParsePolyEvalMsg(msg *shmsg.PolyEvalMsg, sender common.Address) (*PolyEvalMsg, error) {
	receiver, err := validateAddress(msg.Receiver)
	if err != nil {
		return nil, err
	}
	return &PolyEvalMsg{
		Sender:        sender,
		Eon:           msg.Eon,
		Receiver:      receiver,
		EncryptedEval: msg.EncryptedEval,
	}, nil
}

// ParsePolyCommitmentMsg converts a shmsg.PolyCommitmentMsg to an app.PolyCommitmentMsg
func ParsePolyCommitmentMsg(msg *shmsg.PolyCommitmentMsg, sender common.Address) (*PolyCommitmentMsg, error) {
	return &PolyCommitmentMsg{
		Sender: sender,
		Eon:    msg.Eon,
		Gammas: msg.Gammas,
	}, nil
}

// ParseAccusationMsg converts a shmsg.AccusationMsg to an app.AccusationMsg
func ParseAccusationMsg(msg *shmsg.AccusationMsg, sender common.Address) (*AccusationMsg, error) {
	accused, err := validateAddress(msg.Accused)
	if err != nil {
		return nil, err
	}
	return &AccusationMsg{
		Sender:  sender,
		Eon:     msg.Eon,
		Accused: accused,
	}, nil
}

// ParseApologyMsg converts a shmsg.ApologyMsg to an app.ApologyMsg
func ParseApologyMsg(msg *shmsg.ApologyMsg, sender common.Address) (*ApologyMsg, error) {
	accuser, err := validateAddress(msg.Accuser)
	if err != nil {
		return nil, err
	}
	return &ApologyMsg{
		Sender:   sender,
		Eon:      msg.Eon,
		Accuser:  accuser,
		PolyEval: msg.PolyEval,
	}, nil
}

// ParseEpochSKShareMsg converts a shmsg.ESKShareMsg to an app.ESKShareMsg
func ParseEpochSKShareMsg(msg *shmsg.EpochSKShareMsg, sender common.Address) (*EpochSKShareMsg, error) {
	return &EpochSKShareMsg{
		Sender:       sender,
		Eon:          msg.Eon,
		Epoch:        msg.Epoch,
		EpochSKShare: msg.EpochSKShare,
	}, nil
}
