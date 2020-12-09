package keyper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	shcrypto "github.com/brainbot-com/shutter/shuttermint/crypto"
	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
)

type testInstance struct {
	eon                         uint64
	threshold                   uint64
	keypers                     []common.Address
	keyperSigningKeys           map[common.Address]*ecdsa.PrivateKey
	keyperEncryptionPrivateKeys map[common.Address]*ecies.PrivateKey
	keyperEncryptionPublicKeys  map[common.Address]*ecies.PublicKey
	ms                          MockMessageSender
	dkg                         *DKGInstance
}

func setupTestInstance(t *testing.T) testInstance {
	eon := uint64(10)
	threshold := uint64(2)
	n := 3
	keypers := []common.Address{}
	keyperSigningKeys := make(map[common.Address]*ecdsa.PrivateKey)
	keyperEncryptionPrivateKeys := make(map[common.Address]*ecies.PrivateKey)
	keyperEncryptionPublicKeys := make(map[common.Address]*ecies.PublicKey)
	for i := 0; i < n; i++ {
		privateKey, err := crypto.GenerateKey()
		require.Nil(t, err)
		address := crypto.PubkeyToAddress(privateKey.PublicKey)

		encryptionPrivateKeyECDSA, err := crypto.GenerateKey()
		require.Nil(t, err)
		encryptionPrivateKey := ecies.ImportECDSA(encryptionPrivateKeyECDSA)

		keypers = append(keypers, address)
		keyperSigningKeys[address] = privateKey
		keyperEncryptionPrivateKeys[address] = encryptionPrivateKey
		keyperEncryptionPublicKeys[address] = &encryptionPrivateKey.PublicKey
	}

	keyperIndex := 2
	address := keypers[keyperIndex]
	batchConfig := contract.BatchConfig{
		Threshold: threshold,
		Keypers:   keypers,
	}
	keyperConfig := KeyperConfig{
		SigningKey:    keyperSigningKeys[address],
		EncryptionKey: keyperEncryptionPrivateKeys[address],
	}
	ms := NewMockMessageSender()
	dkg, err := NewDKGInstance(eon, batchConfig, keyperConfig, &ms, keyperEncryptionPublicKeys)
	require.Nil(t, err)
	require.Equal(t, eon, dkg.Eon)

	return testInstance{
		eon:                         eon,
		threshold:                   threshold,
		keypers:                     keypers,
		keyperSigningKeys:           keyperSigningKeys,
		keyperEncryptionPrivateKeys: keyperEncryptionPrivateKeys,
		keyperEncryptionPublicKeys:  keyperEncryptionPublicKeys,
		ms:                          ms,
		dkg:                         dkg,
	}
}

func TestDKGInstance(t *testing.T) {
	ti := setupTestInstance(t)
	err := ti.dkg.Run(context.Background())
	require.Nil(t, err)

	t.Run("SendGammas", func(t *testing.T) {
		select {
		case msgContainer := <-ti.ms.Msgs:
			msg := msgContainer.GetPolyCommitmentMsg()
			require.NotNil(t, msg)
			require.Equal(t, ti.eon, msg.Eon)
			gammas := [][]byte{}
			for _, g := range *ti.dkg.pure.Polynomial.Gammas() {
				gammas = append(gammas, g.Marshal())
			}
			require.Equal(t, gammas, msg.Gammas)
		default:
			require.FailNow(t, "no message on channel")
		}
	})

	t.Run("SendPolyEvals", func(t *testing.T) {
		select {
		case msgContainer := <-ti.ms.Msgs:
			msg := msgContainer.GetPolyEvalMsg()
			require.NotNil(t, msg)
			require.Equal(t, ti.eon, msg.Eon)
			require.Equal(t, len(ti.keypers)-1, len(msg.Receivers))
			for i, receiver := range ti.keypers {
				if i == 2 {
					continue
				}
				require.Equal(t, receiver.Bytes(), msg.Receivers[i])
				polyEval := ti.dkg.pure.Polynomial.EvalForKeyper(i)
				decryptedEval, err := ti.keyperEncryptionPrivateKeys[receiver].Decrypt(msg.EncryptedEvals[i], nil, nil)
				require.Nil(t, err)
				require.True(t, bytes.Equal(polyEval.Bytes(), decryptedEval))
			}
		default:
			require.FailNow(t, "no message on channel")
		}
	})
}

func TestDispatchPolyCommitmentRegistered(t *testing.T) {
	senderIndex := 1
	ti := setupTestInstance(t)
	sender := ti.keypers[senderIndex]
	polynomial, err := shcrypto.RandomPolynomial(rand.Reader, shcrypto.DegreeFromThreshold(ti.dkg.pure.Threshold))
	require.Nil(t, err)
	ev := shutterevents.PolyCommitmentRegisteredEvent{
		Eon:    ti.eon,
		Sender: sender,
		Gammas: polynomial.Gammas(),
	}
	ti.dkg.dispatchShuttermintEvent(ev)
	c := ti.dkg.pure.Commitments[senderIndex]
	require.Equal(t, ev.Gammas, c)
}
