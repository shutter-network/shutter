package keyper

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

func TestDKGInstance(t *testing.T) {
	eon := uint64(10)
	threshold := uint64(2)
	keypers := []common.Address{
		common.BigToAddress(big.NewInt(0xaa)),
		common.BigToAddress(big.NewInt(0xbb)),
		common.BigToAddress(big.NewInt(0xcc)),
	}
	config := contract.BatchConfig{
		Threshold: threshold,
		Keypers:   keypers,
	}
	ms := NewMockMessageSender()
	dkg, err := NewDKGInstance(eon, config, ms)
	require.Nil(t, err)
	require.Equal(t, eon, dkg.Eon)
	require.NotNil(t, dkg.Polynomial)

	go dkg.Run() // TODO: add context to stop DKG

	t.Run("SendGammas", func(t *testing.T) {
		msgContainer := <-ms.Msgs
		msg := msgContainer.GetPolyCommitmentMsg()
		require.NotNil(t, msg)
		require.Equal(t, eon, msg.Eon)
		gammas := [][]byte{}
		for _, g := range *dkg.Polynomial.Gammas() {
			gammas = append(gammas, g.Marshal())
		}
		require.Equal(t, gammas, msg.Gammas)
	})

	t.Run("SendPolyEvals", func(t *testing.T) {
		for i, receiver := range keypers {
			msgContainer := <-ms.Msgs
			msg := msgContainer.GetPolyEvalMsg()
			require.NotNil(t, msg)
			require.Equal(t, eon, msg.Eon)
			require.Equal(t, receiver.Bytes(), msg.Receiver)
			polyEval := dkg.Polynomial.EvalForKeyper(i)
			require.Equal(t, polyEval.Bytes(), msg.EncryptedEval)
		}
	})
}
