package keyper

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

func TestNewPolyCommitmentMsg(t *testing.T) {
	eon := uint64(10)
	threshold := uint64(5)
	poly, err := crypto.RandomPolynomial(rand.Reader, threshold)
	require.Nil(t, err)
	gammas := poly.Gammas()

	msgContainer := NewPolyCommitmentMsg(eon, gammas)
	msg := msgContainer.GetPolyCommitmentMsg()
	require.NotNil(t, msg)

	require.Equal(t, eon, msg.Eon)
	require.Equal(t, int(threshold)+1, len(msg.Gammas))
	for i := 0; i < int(threshold)+1; i++ {
		gammaBytes := msg.Gammas[i]
		require.True(t, bytes.Equal(gammaBytes, (*gammas)[i].Marshal()))
	}
}

func TestNewPolyEvalMsg(t *testing.T) {
	eon := uint64(10)
	receiver := common.BigToAddress(big.NewInt(0xaabbcc))
	encryptedEval := []byte("secret")

	msgContainer := NewPolyEvalMsg(eon, receiver, encryptedEval)
	msg := msgContainer.GetPolyEvalMsg()
	require.NotNil(t, msg)

	require.Equal(t, eon, msg.Eon)
	require.Equal(t, receiver.Bytes(), msg.Receiver)
}
