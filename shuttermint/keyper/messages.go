package keyper

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

func NewAccusationMessage(eon uint64, accused []common.Address) *shmsg.Message {
	accusedBytes := [][]byte{}
	for _, a := range accused {
		accusedBytes = append(accusedBytes, a.Bytes())
	}
	return &shmsg.Message{
		Payload: &shmsg.Message_AccusationMsg{
			AccusationMsg: &shmsg.AccusationMsg{
				Eon:     eon,
				Accused: accusedBytes,
			},
		},
	}
}

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
	receivers []common.Address,
	encryptedEvals [][]byte,
) *shmsg.Message {
	rs := [][]byte{}
	for _, receiver := range receivers {
		rs = append(rs, receiver.Bytes())
	}

	return &shmsg.Message{
		Payload: &shmsg.Message_PolyEvalMsg{
			PolyEvalMsg: &shmsg.PolyEvalMsg{
				Eon:            eon,
				Receivers:      rs,
				EncryptedEvals: encryptedEvals,
			},
		},
	}
}

// NewEonStartVoteMsg creates a new eon start vote message.
func NewEonStartVoteMsg(startBatchIndex uint64) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_EonStartVoteMsg{
			EonStartVoteMsg: &shmsg.EonStartVoteMsg{
				StartBatchIndex: startBatchIndex,
			},
		},
	}
}
