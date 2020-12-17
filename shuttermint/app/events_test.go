package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/app/evtype"
	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
)

func TestEvents(t *testing.T) {
	eon := uint64(5)
	sender := common.BigToAddress(new(big.Int).SetUint64(123))
	anotherAddress := common.BigToAddress(new(big.Int).SetUint64(456))
	data := []byte("some data")

	t.Run("MakeCheckInEvent", func(t *testing.T) {
		privateKeyECDSA, err := crypto.GenerateKey()
		publicKey := ecies.ImportECDSA(privateKeyECDSA).PublicKey

		require.Nil(t, err)
		ev := MakeCheckInEvent(sender, &publicKey)
		require.Equal(t, evtype.CheckIn, ev.Type)
		require.Equal(t, 2, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("EncryptionPublicKey"), ev.Attributes[1].Key)
		shev, err := shutterevents.MakeEvent(ev)
		require.Nil(t, err)
		checkInEv, ok := shev.(shutterevents.CheckIn)
		require.True(t, ok)
		require.Equal(t, &publicKey, checkInEv.EncryptionPublicKey)
	})

	t.Run("MakeEonStartedEvent", func(t *testing.T) {
		ev := MakeEonStartedEvent(uint64(10), uint64(20))
		require.Equal(t, evtype.EonStarted, ev.Type)
		require.Equal(t, 2, len(ev.Attributes))
		require.Equal(t, []byte("Eon"), ev.Attributes[0].Key)
		require.Equal(t, []byte("10"), ev.Attributes[0].Value)
		require.Equal(t, []byte("BatchIndex"), ev.Attributes[1].Key)
		require.Equal(t, []byte("20"), ev.Attributes[1].Value)
	})

	t.Run("MakePolyEvalRegisteredEvent", func(t *testing.T) {
		msg := &PolyEval{
			Sender:         sender,
			Eon:            eon,
			Receivers:      []common.Address{anotherAddress},
			EncryptedEvals: [][]byte{data},
		}
		ev := MakePolyEvalRegisteredEvent(msg)
		require.Equal(t, evtype.PolyEval, ev.Type)
		require.Equal(t, 4, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("Eon"), ev.Attributes[1].Key)
		require.Equal(t, []byte("5"), ev.Attributes[1].Value)
		require.Equal(t, []byte("Receivers"), ev.Attributes[2].Key)
		require.Equal(t, []byte(anotherAddress.Hex()), ev.Attributes[2].Value)
		require.Equal(t, []byte("EncryptedEvals"), ev.Attributes[3].Key)
		require.Equal(t, []byte(hexutil.Encode(data)), ev.Attributes[3].Value)
	})

	t.Run("MakePolyCommitmentRegisteredEvent", func(t *testing.T) {
		msg := &PolyCommitmentMsg{
			Sender: sender,
			Eon:    eon,
			Gammas: [][]byte{{0xa0, 0xa1, 0xa2}, {0xa3, 0xa4, 0xa5}},
		}
		ev := MakePolyCommitmentRegisteredEvent(msg)
		require.Equal(t, evtype.PolyCommitment, ev.Type)
		require.Equal(t, 3, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("Eon"), ev.Attributes[1].Key)
		require.Equal(t, []byte("5"), ev.Attributes[1].Value)
		require.Equal(t, []byte("Gammas"), ev.Attributes[2].Key)
		require.Equal(t, []byte("a0a1a2,a3a4a5"), ev.Attributes[2].Value)
	})

	t.Run("MakeAccusationRegisteredEvent", func(t *testing.T) {
		msg := &AccusationMsg{
			Sender:  sender,
			Eon:     eon,
			Accused: []common.Address{anotherAddress},
		}
		ev := MakeAccusationRegisteredEvent(msg)
		require.Equal(t, evtype.Accusation, ev.Type)
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
			Sender:    sender,
			Eon:       eon,
			Accusers:  []common.Address{anotherAddress},
			PolyEvals: [][]byte{data},
		}
		ev := MakeApologyRegisteredEvent(msg)
		require.Equal(t, evtype.Apology, ev.Type)
		require.Equal(t, 4, len(ev.Attributes))
		require.Equal(t, []byte("Sender"), ev.Attributes[0].Key)
		require.Equal(t, []byte(sender.Hex()), ev.Attributes[0].Value)
		require.Equal(t, []byte("Eon"), ev.Attributes[1].Key)
		require.Equal(t, []byte("5"), ev.Attributes[1].Value)
		require.Equal(t, []byte("Accusers"), ev.Attributes[2].Key)
		require.Equal(t, []byte(anotherAddress.Hex()), ev.Attributes[2].Value)
		require.Equal(t, []byte("PolyEvals"), ev.Attributes[3].Key)
		require.Equal(t, []byte(hexutil.Encode(data)), ev.Attributes[3].Value)
	})
}
