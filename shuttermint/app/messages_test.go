package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shcrypto"
	"github.com/shutter-network/shutter/shuttermint/shmsg"
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
		smsg := shmsg.PolyEval{
			Eon:            eon,
			Receivers:      [][]byte{anotherAddressBytes},
			EncryptedEvals: [][]byte{data},
		}
		msg, err := ParsePolyEvalMsg(&smsg, sender)
		assert.NilError(t, err)
		assert.DeepEqual(t, sender, msg.Sender)
		assert.Equal(t, eon, msg.Eon)
		assert.DeepEqual(t, anotherAddress, msg.Receivers[0])
		assert.DeepEqual(t, data, msg.EncryptedEvals[0])

		// invalid receiver
		smsg = shmsg.PolyEval{
			Eon:            eon,
			Receivers:      [][]byte{badAddressBytes},
			EncryptedEvals: [][]byte{data},
		}
		_, err = ParsePolyEvalMsg(&smsg, sender)
		assert.Assert(t, err != nil)
	})

	t.Run("ParsePolyCommitmentMsg", func(t *testing.T) {
		smsg := shmsg.PolyCommitment{
			Eon:    eon,
			Gammas: [][]byte{},
		}
		msg, err := ParsePolyCommitmentMsg(&smsg, sender)
		assert.NilError(t, err)
		assert.DeepEqual(t, sender, msg.Sender)
		assert.Equal(t, eon, msg.Eon)
	})

	t.Run("ParseAccusationMsg", func(t *testing.T) {
		smsg := shmsg.Accusation{
			Eon:     eon,
			Accused: [][]byte{anotherAddressBytes},
		}
		msg, err := ParseAccusationMsg(&smsg, sender)
		assert.NilError(t, err)
		assert.DeepEqual(t, sender, msg.Sender)
		assert.Equal(t, eon, msg.Eon)
		assert.DeepEqual(t, anotherAddress, msg.Accused[0])

		// invalid accused
		smsg = shmsg.Accusation{
			Eon:     eon,
			Accused: [][]byte{badAddressBytes},
		}
		_, err = ParseAccusationMsg(&smsg, sender)
		assert.Assert(t, err != nil)
	})

	t.Run("ParseApologyMsg", func(t *testing.T) {
		smsg := shmsg.Apology{
			Eon:       eon,
			Accusers:  [][]byte{anotherAddressBytes},
			PolyEvals: [][]byte{{}},
		}
		msg, err := ParseApologyMsg(&smsg, sender)
		assert.NilError(t, err)
		assert.DeepEqual(t, sender, msg.Sender)
		assert.Equal(t, eon, msg.Eon)
		assert.DeepEqual(t, anotherAddress, msg.Accusers[0])

		// invalid accuser
		smsg = shmsg.Apology{
			Eon:       eon,
			Accusers:  [][]byte{badAddressBytes},
			PolyEvals: [][]byte{{}},
		}
		_, err = ParseApologyMsg(&smsg, sender)
		assert.Assert(t, err != nil)
	})

	t.Run("ParseEpochSecretKeyShareMsg", func(t *testing.T) {
		share := (*shcrypto.EpochSecretKeyShare)(new(bn256.G1).ScalarBaseMult(big.NewInt(1111)))
		smsg := shmsg.NewEpochSecretKeyShare(eon, epoch, share).GetEpochSecretKeyShare()
		msg, err := ParseEpochSecretKeyShareMsg(smsg, sender)
		assert.NilError(t, err)
		assert.DeepEqual(t, sender, msg.Sender)
		assert.Equal(t, eon, msg.Eon)
		assert.Equal(t, epoch, msg.Epoch)
		assert.DeepEqual(t, share, msg.Share)
	})
}
