package keyper

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// NewPolyCommitmentMsg creates a new poly commitment message containing gamma values.
func NewPolyCommitmentMsg(eon uint64, gammas *crypto.Gammas) *shmsg.Message {
	gammaBytes := [][]byte{}
	for _, gamma := range *gammas {
		gammaBytes = append(gammaBytes, gamma.Marshal())
	}

	return &shmsg.Message{
		Payload: &shmsg.Message_PolyCommitmentMsg{
			PolyCommitmentMsg: &shmsg.PolyCommitmentMsg{
				Eon:    eon,
				Gammas: gammaBytes,
			},
		},
	}
}

// NewPolyEvalMsg creates a new poly eval message.
func NewPolyEvalMsg(
	eon uint64,
	receiver common.Address,
	encryptedEval []byte,
) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_PolyEvalMsg{
			PolyEvalMsg: &shmsg.PolyEvalMsg{
				Eon:           eon,
				Receiver:      receiver.Bytes(),
				EncryptedEval: encryptedEval,
			},
		},
	}
}
