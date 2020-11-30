package keyper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
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
	dkg, err := NewDKGInstance(eon, batchConfig, keyperConfig, ms, keyperEncryptionPublicKeys)
	require.Nil(t, err)
	require.Equal(t, eon, dkg.Eon)
	require.NotNil(t, dkg.Polynomial)

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	group, _ := errgroup.WithContext(ctx)
	group.Go(func() error {
		return ti.dkg.Run(ctx)
	})

	t.Run("SendGammas", func(t *testing.T) {
		msgContainer := <-ti.ms.Msgs
		msg := msgContainer.GetPolyCommitmentMsg()
		require.NotNil(t, msg)
		require.Equal(t, ti.eon, msg.Eon)
		gammas := [][]byte{}
		for _, g := range *ti.dkg.Polynomial.Gammas() {
			gammas = append(gammas, g.Marshal())
		}
		require.Equal(t, gammas, msg.Gammas)
	})

	t.Run("SendPolyEvals", func(t *testing.T) {
		msgContainer := <-ti.ms.Msgs
		msg := msgContainer.GetPolyEvalMsg()
		require.NotNil(t, msg)
		require.Equal(t, ti.eon, msg.Eon)
		require.Equal(t, len(ti.keypers)-1, len(msg.Receivers))
		for i, receiver := range ti.keypers {
			if i == 2 {
				continue
			}
			require.Equal(t, receiver.Bytes(), msg.Receivers[i])
			polyEval := ti.dkg.Polynomial.EvalForKeyper(i)
			decryptedEval, err := ti.keyperEncryptionPrivateKeys[receiver].Decrypt(msg.EncryptedEvals[i], nil, nil)
			require.Nil(t, err)
			require.True(t, bytes.Equal(polyEval.Bytes(), decryptedEval))
		}
	})

	err := group.Wait()
	require.Nil(t, err)
}

func TestDispatchPolyCommitmentRegistered(t *testing.T) {
	ti := setupTestInstance(t)
	sender := ti.keypers[1]
	ti.dkg.Polynomial.Gammas()
	ev := PolyCommitmentRegisteredEvent{
		Eon:    ti.eon,
		Sender: sender,
		Gammas: ti.dkg.Polynomial.Gammas(),
	}
	ti.dkg.dispatchShuttermintEvent(ev)
	c, ok := ti.dkg.Commitment[sender]
	require.True(t, ok)
	require.Equal(t, *ev.Gammas, c)
}
