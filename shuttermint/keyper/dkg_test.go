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

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

func TestDKGInstance(t *testing.T) {
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

	keyperIndex := 1
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

	go dkg.Run(context.Background())

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
			decryptedEval, err := keyperEncryptionPrivateKeys[receiver].Decrypt(msg.EncryptedEval, nil, nil)
			require.Nil(t, err)
			require.True(t, bytes.Equal(polyEval.Bytes(), decryptedEval))
		}
	})
}
