package app

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// NewDKGInstance creates a new DKGInstance.
func NewDKGInstance(config BatchConfig, eon uint64) DKGInstance {
	return DKGInstance{
		Config:              config,
		Eon:                 eon,
		PolyEvalsSeen:       make(map[SenderReceiverPair]struct{}),
		PolyCommitmentsSeen: make(map[common.Address]struct{}),
		AccusationsSeen:     make(map[common.Address]struct{}),
		ApologiesSeen:       make(map[common.Address]struct{}),
	}
}

// RegisterPolyEvalMsg adds a polynomial evaluation message to the instance. It makes sure the
// message meets the basic requirements, i.e. the sender and receivers are keypers and we do not
// send multiple messages from one sender to one receiver.
func (dkg *DKGInstance) RegisterPolyEvalMsg(msg PolyEval) error {
	if msg.Eon != dkg.Eon {
		return fmt.Errorf("msg is from eon %d, not %d", msg.Eon, dkg.Eon)
	}
	if dkg.SubmissionsClosed {
		return fmt.Errorf("submissions are already closed")
	}

	sender := msg.Sender
	if !dkg.Config.IsKeyper(sender) {
		return fmt.Errorf("sender %s is not a keyper", sender.Hex())
	}

	for _, receiver := range msg.Receivers {
		if !dkg.Config.IsKeyper(receiver) {
			return fmt.Errorf("receiver %s is not a keyper", msg.Sender.Hex())
		}
		if receiver == sender {
			return fmt.Errorf("receiver %s is also the sender", msg.Sender.Hex())
		}
		_, ok := dkg.PolyEvalsSeen[SenderReceiverPair{sender, receiver}]
		if ok {
			return fmt.Errorf("polynomial evaluation from keyper %s for receiver %s already present", sender.Hex(), receiver.Hex())
		}
	}

	for _, receiver := range msg.Receivers {
		dkg.PolyEvalsSeen[SenderReceiverPair{sender, receiver}] = struct{}{}
	}

	return nil
}

// RegisterPolyCommitmentMsg adds a polynomial commitment message to the instance.
func (dkg *DKGInstance) RegisterPolyCommitmentMsg(msg PolyCommitment) error {
	if msg.Eon != dkg.Eon {
		return fmt.Errorf("msg is from eon %d, not %d", msg.Eon, dkg.Eon)
	}
	if dkg.SubmissionsClosed {
		return fmt.Errorf("submissions are already closed")
	}
	if !dkg.Config.IsKeyper(msg.Sender) {
		return fmt.Errorf("sender %s is not a keyper", msg.Sender.Hex())
	}

	if _, ok := dkg.PolyCommitmentsSeen[msg.Sender]; ok {
		return fmt.Errorf("polynomial commitment from keyper %s already present", msg.Sender.Hex())
	}
	dkg.PolyCommitmentsSeen[msg.Sender] = struct{}{}

	return nil
}

// RegisterAccusationMsg adds an accusation message to the instance.
func (dkg *DKGInstance) RegisterAccusationMsg(msg Accusation) error {
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

	if _, ok := dkg.AccusationsSeen[msg.Sender]; ok {
		return fmt.Errorf("accusation from keyper %s already present", msg.Sender.Hex())
	}
	dkg.AccusationsSeen[msg.Sender] = struct{}{}

	return nil
}

// RegisterApologyMsg adds an apology message to the instance.
func (dkg *DKGInstance) RegisterApologyMsg(msg Apology) error {
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

	if _, ok := dkg.ApologiesSeen[msg.Sender]; ok {
		return fmt.Errorf("apology from keyper %s already present", msg.Sender.Hex())
	}
	dkg.ApologiesSeen[msg.Sender] = struct{}{}

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
