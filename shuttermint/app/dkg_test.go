package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var polyEval = []*big.Int{new(big.Int).SetBytes([]byte{})}

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
		msg := PolyEval{
			Sender:         keypers[0],
			Eon:            eon + 1,
			Receivers:      []common.Address{keypers[1]},
			EncryptedEvals: [][]byte{{}},
		}
		err := dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)

		// fail if sender is not a keyper
		msg = PolyEval{
			Sender:         nonKeyper,
			Eon:            eon,
			Receivers:      []common.Address{keypers[0]},
			EncryptedEvals: [][]byte{{}},
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)

		// fail if receiver is not a keyper
		msg = PolyEval{
			Sender:         keypers[0],
			Eon:            eon,
			Receivers:      []common.Address{nonKeyper},
			EncryptedEvals: [][]byte{{}},
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)

		// fail if sender and receiver are equal
		msg = PolyEval{
			Sender:         keypers[0],
			Eon:            eon,
			Receivers:      []common.Address{keypers[0]},
			EncryptedEvals: [][]byte{{}},
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)

		// adding should work
		msg = PolyEval{
			Sender:         keypers[0],
			Eon:            eon,
			Receivers:      []common.Address{keypers[1]},
			EncryptedEvals: [][]byte{{}},
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.Nil(t, err)

		// adding twice should fail
		msg = PolyEval{
			Sender:         keypers[0],
			Eon:            eon,
			Receivers:      []common.Address{keypers[1]},
			EncryptedEvals: [][]byte{{}},
		}
		err = dkg.RegisterPolyEvalMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterPolyCommitmentMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)

		// fail if wrong eon
		msg := PolyCommitment{
			Sender: keypers[0],
			Eon:    eon + 1,
		}
		err := dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.PolyCommitmentsSeen[nonKeyper]
		require.False(t, ok)

		// fail if sender is not a keyper
		msg = PolyCommitment{
			Sender: nonKeyper,
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.PolyCommitmentsSeen[nonKeyper]
		require.False(t, ok)

		// adding should work
		msg = PolyCommitment{
			Sender: keypers[0],
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.Nil(t, err)
		_, ok = dkg.PolyCommitmentsSeen[keypers[0]]
		require.True(t, ok)

		// adding twice should fail
		msg = PolyCommitment{
			Sender: keypers[0],
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterAccusationMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)

		// fail if wrong eon
		msg := Accusation{
			Sender:  keypers[0],
			Eon:     eon + 1,
			Accused: []common.Address{keypers[1]},
		}
		err := dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.AccusationsSeen[nonKeyper]
		require.False(t, ok)

		// fail if sender is not a keyper
		msg = Accusation{
			Sender:  nonKeyper,
			Eon:     eon,
			Accused: []common.Address{keypers[0]},
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationsSeen[nonKeyper]
		require.False(t, ok)

		// fail if accused is not a keyper
		msg = Accusation{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: []common.Address{nonKeyper},
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationsSeen[keypers[0]]
		require.False(t, ok)

		// fail if sender and accused are equal
		msg = Accusation{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: []common.Address{keypers[0]},
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationsSeen[keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = Accusation{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: []common.Address{keypers[1]},
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.Nil(t, err)
		_, ok = dkg.AccusationsSeen[keypers[0]]
		require.True(t, ok)

		// adding twice should fail
		msg = Accusation{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: []common.Address{keypers[1]},
		}
		err = dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)
	})

	t.Run("RegisterApologyMsg", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)

		// fail if wrong eon
		msg := Apology{
			Sender:   keypers[0],
			Eon:      eon + 1,
			Accusers: []common.Address{keypers[1]},
			PolyEval: polyEval,
		}
		err := dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok := dkg.AccusationsSeen[nonKeyper]
		require.False(t, ok)

		// fail if sender is not a keyper
		msg = Apology{
			Sender:   nonKeyper,
			Eon:      eon,
			Accusers: []common.Address{keypers[0]},
			PolyEval: polyEval,
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.AccusationsSeen[nonKeyper]
		require.False(t, ok)

		// fail if accuser is not a keyper
		msg = Apology{
			Sender:   keypers[0],
			Eon:      eon,
			Accusers: []common.Address{nonKeyper},
			PolyEval: polyEval,
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.ApologiesSeen[keypers[0]]
		require.False(t, ok)

		// fail if sender and accused are equal
		msg = Apology{
			Sender:   keypers[0],
			Eon:      eon,
			Accusers: []common.Address{keypers[0]},
			PolyEval: polyEval,
		}
		err = dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
		_, ok = dkg.ApologiesSeen[keypers[0]]
		require.False(t, ok)

		// adding should work
		msg = Apology{
			Sender:   keypers[0],
			Eon:      eon,
			Accusers: []common.Address{keypers[1]},
			PolyEval: polyEval,
		}
		err = dkg.RegisterApologyMsg(msg)
		require.Nil(t, err)
		_, ok = dkg.ApologiesSeen[keypers[0]]
		require.True(t, ok)

		// adding twice should fail
		msg = Apology{
			Sender:   keypers[0],
			Eon:      eon,
			Accusers: []common.Address{keypers[1]},
			PolyEval: polyEval,
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

		msg1 := PolyEval{
			Sender:         keypers[0],
			Eon:            eon,
			Receivers:      []common.Address{keypers[1]},
			EncryptedEvals: [][]byte{},
		}
		err := dkg.RegisterPolyEvalMsg(msg1)
		require.NotNil(t, err)

		msg2 := PolyCommitment{
			Sender: keypers[0],
			Eon:    eon,
		}
		err = dkg.RegisterPolyCommitmentMsg(msg2)
		require.NotNil(t, err)

		// accusations and apologies still work
		msg3 := Accusation{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: []common.Address{keypers[1]},
		}
		err = dkg.RegisterAccusationMsg(msg3)
		require.Nil(t, err)
		msg4 := Apology{
			Sender:   keypers[0],
			Eon:      eon,
			Accusers: []common.Address{keypers[1]},
		}
		err = dkg.RegisterApologyMsg(msg4)
		require.Nil(t, err)
	})

	t.Run("CloseAccusations", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)
		dkg.CloseAccusations()

		msg := Accusation{
			Sender:  keypers[0],
			Eon:     eon,
			Accused: []common.Address{keypers[1]},
		}
		err := dkg.RegisterAccusationMsg(msg)
		require.NotNil(t, err)

		// Apologies still work
		msg2 := Apology{
			Sender:   keypers[0],
			Eon:      eon,
			Accusers: []common.Address{keypers[1]},
		}
		err = dkg.RegisterApologyMsg(msg2)
		require.Nil(t, err)
	})

	t.Run("CloseApologies", func(t *testing.T) {
		dkg := NewDKGInstance(config, eon)
		dkg.CloseApologies()

		msg := Apology{
			Sender:   keypers[0],
			Eon:      eon,
			Accusers: []common.Address{keypers[1]},
			PolyEval: polyEval,
		}
		err := dkg.RegisterApologyMsg(msg)
		require.NotNil(t, err)
	})
}
