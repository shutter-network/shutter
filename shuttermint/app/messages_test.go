package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

func TestMessageParsing(t *testing.T) {
	eon := uint64(5)
	epoch := uint64(10)
	sender := common.BigToAddress(new(big.Int).SetUint64(123))
	anotherAddress := common.BigToAddress(new(big.Int).SetUint64(456))
	anotherAddressBytes := anotherAddress.Bytes()
	badAddressBytes := []byte("only nineteen bytes")
	data := []byte("some data")

	t.Run("ParsePolyEvalMsg", func(t *testing.T) {
		smsg := shmsg.PolyEvalMsg{
			Eon:           eon,
			Receiver:      anotherAddressBytes,
			EncryptedEval: data,
		}
		msg, err := ParsePolyEvalMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, anotherAddress, msg.Receiver)
		require.Equal(t, data, msg.EncryptedEval)

		// invalid receiver
		smsg = shmsg.PolyEvalMsg{
			Eon:           eon,
			Receiver:      badAddressBytes,
			EncryptedEval: data,
		}
		_, err = ParsePolyEvalMsg(&smsg, sender)
		require.NotNil(t, err)
	})

	t.Run("ParsePolyCommitmentMsg", func(t *testing.T) {
		smsg := shmsg.PolyCommitmentMsg{
			Eon:    eon,
			Gammas: [][]byte{},
		}
		msg, err := ParsePolyCommitmentMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
	})

	t.Run("ParseAccusationMsg", func(t *testing.T) {
		smsg := shmsg.AccusationMsg{
			Eon:     eon,
			Accused: anotherAddressBytes,
		}
		msg, err := ParseAccusationMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, anotherAddress, msg.Accused)

		// invalid accused
		smsg = shmsg.AccusationMsg{
			Eon:     eon,
			Accused: badAddressBytes,
		}
		_, err = ParseAccusationMsg(&smsg, sender)
		require.NotNil(t, err)
	})

	t.Run("ParseApologyMsg", func(t *testing.T) {
		smsg := shmsg.ApologyMsg{
			Eon:      eon,
			Accuser:  anotherAddressBytes,
			PolyEval: []byte{},
		}
		msg, err := ParseApologyMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, anotherAddress, msg.Accuser)

		// invalid accuser
		smsg = shmsg.ApologyMsg{
			Eon:      eon,
			Accuser:  badAddressBytes,
			PolyEval: []byte{},
		}
		_, err = ParseApologyMsg(&smsg, sender)
		require.NotNil(t, err)
	})

	t.Run("ParseEpochSKShareMsg", func(t *testing.T) {
		smsg := shmsg.EpochSKShareMsg{
			Eon:          eon,
			Epoch:        epoch,
			EpochSKShare: []byte{},
		}
		msg, err := ParseEpochSKShareMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, epoch, msg.Epoch)
	})
}
