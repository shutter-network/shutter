package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestRegisterMsgs(t *testing.T) {
	eon := uint64(10)
	keypers := []common.Address{}
	for i := 0; i < 3; i++ {
		keypers = append(keypers, common.BigToAddress(big.NewInt(int64(i+10))))
	}
	nonKeyper := common.BigToAddress(big.NewInt(666))
	config := BatchConfig{
		Keypers: keypers,
	}

	t.Run("RegisterPolyEvalMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)

		// fail if wrong eon
		msg := PolyEvalMsg{
			Sender:   keypers[0],
			Eon:      eon + 1,
			Receiver: keypers[1],
		}
		err := dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.PolyEvalMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if sender is not a keyper
		msg = PolyEvalMsg{
			Sender:   nonKeyper,
			Eon:      eon,
			Receiver: keypers[0],
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.PolyEvalMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if receiver is not a keyper
		msg = PolyEvalMsg{
			Sender:   keypers[0],
			Eon:      eon,
			Receiver: nonKeyper,
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.PolyEvalMsgs[keypers[0]][nonKeyper]
		require.False(t, ok)

		// fail if sender and receiver are equal
		msg = PolyEvalMsg{
			Sender:   keypers[0],
			Eon:      eon,
			Receiver: keypers[0],
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.PolyEvalMsgs[keypers[0]][keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = PolyEvalMsg{
			Sender:   keypers[0],
			Eon:      eon,
			Receiver: keypers[1],
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.Nil(t, err)
		storedMsg, ok := dkg.PolyEvalMsgs[keypers[0]][keypers[1]]
		require.True(t, ok)
		require.Equal(t, msg, storedMsg)

		// adding twice should fail
		msg = PolyEvalMsg{
			Sender:   keypers[0],
			Eon:      eon,
			Receiver: keypers[1],
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterPolyCommitmentMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)

		// fail if wrong eon
		msg := PolyCommitmentMsg{
			Sender: keypers[0],
			Eon:    eon + 1,
		}
		err := dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.PolyCommitmentMsgs[nonKeyper]
		require.False(t, ok)

		// fail if sender is not a keyper
		msg = PolyCommitmentMsg{
			Sender: nonKeyper,
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.PolyCommitmentMsgs[nonKeyper]
		require.False(t, ok)

		// adding should work
		msg = PolyCommitmentMsg{
			Sender: keypers[0],
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.Nil(t, err)
		storedMsg, ok := dkg.PolyCommitmentMsgs[keypers[0]]
		require.True(t, ok)
		require.Equal(t, msg, storedMsg)

		// adding twice should fail
		msg = PolyCommitmentMsg{
			Sender: keypers[0],
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterAccusationMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)

		// fail if wrong eon
		msg := AccusationMsg{
			Sender:  keypers[0],
			Eon:     eon + 1,
			Accused: keypers[1],
		}
		err := dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.AccusationMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if sender is not a keyper
		msg = AccusationMsg{
			Sender:  nonKeyper,
			Eon:     eon,
			Accused: keypers[0],
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if accused is not a keyper
		msg = AccusationMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: nonKeyper,
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationMsgs[keypers[0]][nonKeyper]
		require.False(t, ok)

		// fail if sender and accused are equal
		msg = AccusationMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: keypers[0],
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationMsgs[keypers[0]][keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = AccusationMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: keypers[1],
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.Nil(t, err)
		_, ok = dkg.AccusationMsgs[keypers[0]]
		require.True(t, ok)

		// adding twice should fail
		msg = AccusationMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: keypers[1],
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterApologyMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)

		// fail if wrong eon
		msg := ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon + 1,
			Accuser: keypers[1],
		}
		err := dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.AccusationMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if sender is not a keyper
		msg = ApologyMsg{
			Sender:  nonKeyper,
			Eon:     eon,
			Accuser: keypers[0],
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationMsgs[nonKeyper][nonKeyper]
		require.False(t, ok)

		// fail if accuser is not a keyper
		msg = ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accuser: nonKeyper,
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.ApologyMsgs[keypers[0]][nonKeyper]
		require.False(t, ok)

		// fail if sender and accused are equal
		msg = ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accuser: keypers[0],
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.ApologyMsgs[keypers[0]][keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accuser: keypers[1],
		}
		err = dkg.RegisterApologyMsg(msg)
		require.Nil(t, err)
		_, ok = dkg.ApologyMsgs[keypers[0]]
		require.True(t, ok)

		// adding twice should fail
		msg = ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accuser: keypers[1],
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
	})
}

func TestClosing(t *testing.T) {
	keypers := []common.Address{}
	for i := 0; i < 3; i++ {
		keypers = append(keypers, common.BigToAddress(big.NewInt(int64(i+10))))
	}
	eon := uint64(10)
	config := BatchConfig{
		Keypers: keypers,
	}

	t.Run("CloseSubmissions", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)
		dkg.CloseSubmissions()

		msg1 := PolyEvalMsg{
			Sender:   keypers[0],
			Eon:      eon,
			Receiver: keypers[1],
		}
		err := dkg.RegisterPolyEvalMsg(msg1)
		require.NotNil(t, err)

		msg2 := PolyCommitmentMsg{
			Sender: keypers[0],
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg2)
		require.NotNil(t, err)

		// accusations and apologies still work
		msg3 := AccusationMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: keypers[1],
		}
		err = dkg.RegisterAccusationMsg(msg3)
		require.Nil(t, err)
		msg4 := ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accuser: keypers[1],
		}
		err = dkg.RegisterApologyMsg(msg4)
		require.Nil(t, err)
	})

	t.Run("CloseAccusations", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)
		dkg.CloseAccusations()

		msg := AccusationMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: keypers[1],
		}
		err := dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)

		// Apologies still work
		msg2 := ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accuser: keypers[1],
		}
		err = dkg.RegisterApologyMsg(msg2)
		require.Nil(t, err)
	})

	t.Run("CloseApologies", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)
		dkg.CloseApologies()

		msg := ApologyMsg{
			Sender:  keypers[0],
			Eon:     eon,
			Accuser: keypers[1],
		}
		err := dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
	})
}
