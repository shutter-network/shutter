package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestRegisterMsgs(t *testing.T) {
	keypers := []common.Address{}
	for i := 0; i < 3; i++ {
		keypers = append(keypers, common.BigToAddress(big.NewInt(int64(i+10))))
	}
	nonKeyper := common.BigToAddress(big.NewInt(666))
	config := BatchConfig{
		Keypers: keypers,
	}

	t.Run("RegisterPolyEvalMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config)

		// fail if sender is not a keyper
		msg := PolyEvalMsg{
			Sender:   nonKeyper,
			Receiver: keypers[0],
		}
		err := dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.PolyEvalMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if receiver is not a keyper
		msg = PolyEvalMsg{
			Sender:   keypers[0],
			Receiver: nonKeyper,
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.PolyEvalMsgs[keypers[0]][nonKeyper]
		require.False(t, ok)

		// fail if sender and receiver are equal
		msg = PolyEvalMsg{
			Sender:   keypers[0],
			Receiver: keypers[0],
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.PolyEvalMsgs[keypers[0]][keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = PolyEvalMsg{
			Sender:   keypers[0],
			Receiver: keypers[1],
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.Nil(t, err)
		storedMsg, ok := dkg.PolyEvalMsgs[keypers[0]][keypers[1]]
		require.True(t, ok)
		require.Equal(t, msg, storedMsg)

		// adding twice should fail
		msg = PolyEvalMsg{Sender: keypers[0]}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterPolyCommitmentMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config)

		// fail if sender is not a keyper
		msg := PolyCommitmentMsg{
			Sender: nonKeyper,
		}
		err := dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.PolyCommitmentMsgs[nonKeyper]
		require.False(t, ok)

		// adding should work
		msg = PolyCommitmentMsg{Sender: keypers[0]}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.Nil(t, err)
		storedMsg, ok := dkg.PolyCommitmentMsgs[keypers[0]]
		require.True(t, ok)
		require.Equal(t, msg, storedMsg)

		// adding twice should fail
		msg = PolyCommitmentMsg{Sender: keypers[0]}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterAccusationMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config)

		// fail if sender is not a keyper
		msg := AccusationMsg{
			Sender:  nonKeyper,
			Accused: keypers[0],
		}
		err := dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.AccusationMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if accused is not a keyper
		msg = AccusationMsg{
			Sender:  keypers[0],
			Accused: nonKeyper,
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationMsgs[keypers[0]][nonKeyper]
		require.False(t, ok)

		// fail if sender and accused are equal
		msg = AccusationMsg{
			Sender:  keypers[0],
			Accused: keypers[0],
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationMsgs[keypers[0]][keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = AccusationMsg{
			Sender:  keypers[0],
			Accused: keypers[1],
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.Nil(t, err)
		_, ok = dkg.AccusationMsgs[keypers[0]]
		require.True(t, ok)

		// adding twice should fail
		msg = AccusationMsg{Sender: keypers[0]}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterApologyMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config)

		// fail if sender is not a keyper
		msg := ApologyMsg{
			Sender:  nonKeyper,
			Accuser: keypers[0],
		}
		err := dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.AccusationMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if accuser is not a keyper
		msg = ApologyMsg{
			Sender:  keypers[0],
			Accuser: nonKeyper,
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.ApologyMsgs[keypers[0]][nonKeyper]
		require.False(t, ok)

		// fail if sender and accused are equal
		msg = ApologyMsg{
			Sender:  keypers[0],
			Accuser: keypers[0],
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.ApologyMsgs[keypers[0]][keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = ApologyMsg{
			Sender:  keypers[0],
			Accuser: keypers[1],
		}
		err = dkg.RegisterApologyMsg(msg)
		require.Nil(t, err)
		_, ok = dkg.ApologyMsgs[keypers[0]]
		require.True(t, ok)

		// adding twice should fail
		msg = ApologyMsg{Sender: keypers[0]}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
	})
}
