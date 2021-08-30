package shmsg

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

func TestNewPolyCommitmentMsg(t *testing.T) {
	eon := uint64(10)
	threshold := uint64(5)
	poly, err := shcrypto.RandomPolynomial(rand.Reader, threshold)
	assert.NilError(t, err)
	gammas := poly.Gammas()

	msgContainer := NewPolyCommitment(eon, gammas)
	msg := msgContainer.GetPolyCommitment()
	assert.Assert(t, msg != nil)

	assert.Equal(t, eon, msg.Eon)
	assert.Equal(t, int(threshold)+1, len(msg.Gammas))
	for i := 0; i < int(threshold)+1; i++ {
		gammaBytes := msg.Gammas[i]
		assert.DeepEqual(t, gammaBytes, (*gammas)[i].Marshal())
	}
}

func TestNewPolyEvalMsg(t *testing.T) {
	eon := uint64(10)
	receiver := common.BigToAddress(big.NewInt(0xaabbcc))
	encryptedEval := []byte("secret")

	msgContainer := NewPolyEval(eon, []common.Address{receiver}, [][]byte{encryptedEval})
	msg := msgContainer.GetPolyEval()
	assert.Assert(t, msg != nil)

	assert.Equal(t, eon, msg.Eon)
	assert.DeepEqual(t, receiver.Bytes(), msg.Receivers[0])
}
