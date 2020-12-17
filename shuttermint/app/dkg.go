package app

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// NewDKGInstance creates a new DKGInstance.
func NewDKGInstance(config BatchConfig, eon uint64) DKGInstance {
	polyEvalMsgs := make(map[common.Address]PolyEval)
	polyCommitmentMsgs := make(map[common.Address]PolyCommitmentMsg)
	accusationMsgs := make(map[common.Address]AccusationMsg)
	apologyMsgs := make(map[common.Address]ApologyMsg)

	return DKGInstance{
		Config: config,
		Eon:    eon,

		PolyEvalMsgs:       polyEvalMsgs,
		PolyCommitmentMsgs: polyCommitmentMsgs,
		AccusationMsgs:     accusationMsgs,
		ApologyMsgs:        apologyMsgs,
	}
}

// RegisterPolyEvalMsg adds a polynomial evaluation message to the instance.
func (dkg *DKGInstance) RegisterPolyEvalMsg(msg PolyEval) error {
	if msg.Eon != dkg.Eon {
		return fmt.Errorf("msg is from eon %d, not %d", msg.Eon, dkg.Eon)
	}
	if dkg.SubmissionsClosed {
		return fmt.Errorf("submissions are already closed")
	}
	if !dkg.Config.IsKeyper(msg.Sender) {
		return fmt.Errorf("sender %s is not a keyper", msg.Sender.Hex())
	}
	for _, receiver := range msg.Receivers {
		if !dkg.Config.IsKeyper(receiver) {
			return fmt.Errorf("receiver %s is not a keyper", msg.Sender.Hex())
		}
		if msg.Sender == receiver {
			return fmt.Errorf("receiver %s is also the sender", msg.Sender.Hex())
		}
	}

	if _, ok := dkg.PolyEvalMsgs[msg.Sender]; ok {
		return fmt.Errorf("polynomial evaluation from keyper %s already present", msg.Sender.Hex())
	}
	dkg.PolyEvalMsgs[msg.Sender] = msg

	return nil
}

// RegisterPolyCommitmentMsg adds a polynomial commitment message to the instance.
func (dkg *DKGInstance) RegisterPolyCommitmentMsg(msg PolyCommitmentMsg) error {
	if msg.Eon != dkg.Eon {
		return fmt.Errorf("msg is from eon %d, not %d", msg.Eon, dkg.Eon)
	}
	if dkg.SubmissionsClosed {
		return fmt.Errorf("submissions are already closed")
	}
	if !dkg.Config.IsKeyper(msg.Sender) {
		return fmt.Errorf("sender %s is not a keyper", msg.Sender.Hex())
	}

	if _, ok := dkg.PolyCommitmentMsgs[msg.Sender]; ok {
		return fmt.Errorf("polynomial commitment from keyper %s already present", msg.Sender.Hex())
	}
	dkg.PolyCommitmentMsgs[msg.Sender] = msg

	return nil
}

// RegisterAccusationMsg adds an accusation message to the instance.
func (dkg *DKGInstance) RegisterAccusationMsg(msg AccusationMsg) error {
	if msg.Eon != dkg.Eon {
		return fmt.Errorf("msg is from eon %d, not %d", msg.Eon, dkg.Eon)
	}
	if dkg.AccusationsClosed {
		return fmt.Errorf("accusations are already closed")
	}
	if !dkg.Config.IsKeyper(msg.Sender) {
		return fmt.Errorf("sender %s is not a keyper", msg.Sender.Hex())
	}
	for _, accused := range msg.Accused {
		if !dkg.Config.IsKeyper(accused) {
			return fmt.Errorf("accused %s is not a keyper", accused.Hex())
		}
		if msg.Sender == accused {
			return fmt.Errorf("sender %s is accusing themselves", msg.Sender.Hex())
		}
	}

	if _, ok := dkg.AccusationMsgs[msg.Sender]; ok {
		return fmt.Errorf("accusation from keyper %s already present", msg.Sender.Hex())
	}
	dkg.AccusationMsgs[msg.Sender] = msg

	return nil
}

// RegisterApologyMsg adds an apology message to the instance.
func (dkg *DKGInstance) RegisterApologyMsg(msg ApologyMsg) error {
	if msg.Eon != dkg.Eon {
		return fmt.Errorf("msg is from eon %d, not %d", msg.Eon, dkg.Eon)
	}
	if dkg.ApologiesClosed {
		return fmt.Errorf("apologies are already closed")
	}
	if !dkg.Config.IsKeyper(msg.Sender) {
		return fmt.Errorf("sender %s is not a keyper", msg.Sender.Hex())
	}
	for _, accuser := range msg.Accusers {
		if !dkg.Config.IsKeyper(accuser) {
			return fmt.Errorf("accuser %s is not a keyper", msg.Sender.Hex())
		}
		if msg.Sender == accuser {
			return fmt.Errorf("sender %s sends apology for accusation against themselves", msg.Sender.Hex())
		}
	}

	if _, ok := dkg.ApologyMsgs[msg.Sender]; ok {
		return fmt.Errorf("apology from keyper %s already present", msg.Sender.Hex())
	}
	dkg.ApologyMsgs[msg.Sender] = msg

	return nil
}

// CloseSubmissions prevents new polynomial evaluations or commitments from being registered.
func (dkg *DKGInstance) CloseSubmissions() {
	dkg.SubmissionsClosed = true
}

// CloseAccusations prevents new accusations from being registered.
func (dkg *DKGInstance) CloseAccusations() {
	dkg.AccusationsClosed = true
}

// CloseApologies prevents new apologies from being registered.
func (dkg *DKGInstance) CloseApologies() {
	dkg.ApologiesClosed = true
}
