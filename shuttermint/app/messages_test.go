package app

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/crypto"
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
		smsg := shmsg.PolyEval{
			Eon:            eon,
			Receivers:      [][]byte{anotherAddressBytes},
			EncryptedEvals: [][]byte{data},
		}
		msg, err := ParsePolyEvalMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, anotherAddress, msg.Receivers[0])
		require.Equal(t, data, msg.EncryptedEvals[0])

		// invalid receiver
		smsg = shmsg.PolyEval{
			Eon:            eon,
			Receivers:      [][]byte{badAddressBytes},
			EncryptedEvals: [][]byte{data},
		}
		_, err = ParsePolyEvalMsg(&smsg, sender)
		require.NotNil(t, err)
	})

	t.Run("ParsePolyCommitmentMsg", func(t *testing.T) {
		smsg := shmsg.PolyCommitment{
			Eon:    eon,
			Gammas: [][]byte{},
		}
		msg, err := ParsePolyCommitmentMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
	})

	t.Run("ParseAccusationMsg", func(t *testing.T) {
		smsg := shmsg.Accusation{
			Eon:     eon,
			Accused: [][]byte{anotherAddressBytes},
		}
		msg, err := ParseAccusationMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, anotherAddress, msg.Accused[0])

		// invalid accused
		smsg = shmsg.Accusation{
			Eon:     eon,
			Accused: [][]byte{badAddressBytes},
		}
		_, err = ParseAccusationMsg(&smsg, sender)
		require.NotNil(t, err)
	})

	t.Run("ParseApologyMsg", func(t *testing.T) {
		smsg := shmsg.Apology{
			Eon:       eon,
			Accusers:  [][]byte{anotherAddressBytes},
			PolyEvals: [][]byte{[]byte{}},
		}
		msg, err := ParseApologyMsg(&smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, anotherAddress, msg.Accusers[0])

		// invalid accuser
		smsg = shmsg.Apology{
			Eon:       eon,
			Accusers:  [][]byte{badAddressBytes},
			PolyEvals: [][]byte{{}},
		}
		_, err = ParseApologyMsg(&smsg, sender)
		require.NotNil(t, err)
	})

	t.Run("ParseEpochSecretKeyShareMsg", func(t *testing.T) {
		share := (*crypto.EpochSecretKeyShare)(new(bn256.G1).ScalarBaseMult(big.NewInt(1111)))
		smsg := shmsg.NewEpochSecretKeyShare(eon, epoch, share).GetEpochSecretKeyShare()
		msg, err := ParseEpochSecretKeyShareMsg(smsg, sender)
		require.Nil(t, err)
		require.Equal(t, sender, msg.Sender)
		require.Equal(t, eon, msg.Eon)
		require.Equal(t, epoch, msg.Epoch)
		require.Equal(t, share, msg.Share)
	})
}
