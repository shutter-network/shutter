package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestEvents(t *testing.T) {
	eon := uint64(5)
	sender := common.BigToAddress(new(big.Int).SetUint64(123))
	anotherAddress := common.BigToAddress(new(big.Int).SetUint64(456))
	data := []byte("some data")

	t.Run("MakeNewDKGInstanceEvent", func(t *testing.T) {
		ev := MakeNewDKGInstanceEvent(uint64(10), uint64(20))
		require.Equal(t, "shutter.new-dkg-instance", ev.Type)
		require.Equal(t, 2, len(ev.Attributes))
		require.Equal(t, []byte("Eon"), ev.Attributes[0].Key)
		require.Equal(t, []byte("10"), ev.Attributes[0].Value)
		require.Equal(t, []byte("ConfigIndex"), ev.Attributes[1].Key)
		require.Equal(t, []byte("20"), ev.Attributes[1].Value)
	})

	t.Run("MakePolyEvalRegisteredEvent", func(t *testing.T) {
		msg := &PolyEvalMsg{
			Sender:        sender,
			Eon:           eon,
			Receiver:      anotherAddress,
			EncryptedEval: data,
		}
		ev := MakePolyEvalRegisteredEvent(msg)
		require.Equal(t, ev.Type, "shutter.poly-eval-registered")
		require.Equal(t, 4, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("Eon"), ev.Attributes[1].Key)
		require.Equal(t, []byte("5"), ev.Attributes[1].Value)
		require.Equal(t, []byte("Receiver"), ev.Attributes[2].Key)
		require.Equal(t, []byte(anotherAddress.Hex()), ev.Attributes[2].Value)
		require.Equal(t, []byte("EncryptedEval"), ev.Attributes[3].Key)
		require.Equal(t, data, ev.Attributes[3].Value)
	})

	t.Run("MakePolyCommitmentRegisteredEvent", func(t *testing.T) {
		msg := &PolyCommitmentMsg{
			Sender: sender,
			Eon:    eon,
		}
		ev := MakePolyCommitmentRegisteredEvent(msg)
		require.Equal(t, ev.Type, "shutter.poly-commitment-registered")
		require.Equal(t, 2, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("Eon"), ev.Attributes[1].Key)
		require.Equal(t, []byte("5"), ev.Attributes[1].Value)
	})

	t.Run("MakeAccusationRegisteredEvent", func(t *testing.T) {
		msg := &AccusationMsg{
			Sender:  sender,
			Eon:     eon,
			Accused: anotherAddress,
		}
		ev := MakeAccusationRegisteredEvent(msg)
		require.Equal(t, ev.Type, "shutter.accusation-registered")
		require.Equal(t, 3, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("Eon"), ev.Attributes[1].Key)
		require.Equal(t, []byte("5"), ev.Attributes[1].Value)
		require.Equal(t, []byte("Accused"), ev.Attributes[2].Key)
		require.Equal(t, []byte(anotherAddress.Hex()), ev.Attributes[2].Value)
	})

	t.Run("MakeApologyRegisteredEvent", func(t *testing.T) {
		msg := &ApologyMsg{
			Sender:   sender,
			Eon:      eon,
			Accuser:  anotherAddress,
			PolyEval: data,
		}
		ev := MakeApologyRegisteredEvent(msg)
		require.Equal(t, ev.Type, "shutter.apology-registered")
		require.Equal(t, 4, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("Eon"), ev.Attributes[1].Key)
		require.Equal(t, []byte("5"), ev.Attributes[1].Value)
		require.Equal(t, []byte("Accuser"), ev.Attributes[2].Key)
		require.Equal(t, []byte(anotherAddress.Hex()), ev.Attributes[2].Value)
		require.Equal(t, []byte("PolyEval"), ev.Attributes[3].Key)
		require.Equal(t, data, ev.Attributes[3].Value)
	})
}
