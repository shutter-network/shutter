package app

import (
	"testing"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

func TestNewShutterApp(t *testing.T) {
	app := NewShutterApp()
	require.Equal(t, len(app.Configs), 1, "Configs should contain exactly one guard element")
	require.Equal(t, app.Configs[0], &BatchConfig{}, "Bad guard element")
}

func TestGetBatch(t *testing.T) {
	app := NewShutterApp()

	err := app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers:         addr,
	})
	require.Nil(t, err)

	err = app.addConfig(BatchConfig{
		ConfigIndex:     2,
		StartBatchIndex: 200,
		Threshold:       2,
		Keypers:         addr,
	})
	require.Nil(t, err)

	err = app.addConfig(BatchConfig{
		ConfigIndex:     3,
		StartBatchIndex: 300,
		Threshold:       3,
		Keypers:         addr,
	})
	require.Nil(t, err)

	require.Equal(t, uint64(0), app.getBatchState(0).Config.Threshold)
	require.Equal(t, uint64(0), app.getBatchState(99).Config.Threshold)
	require.Equal(t, uint64(1), app.getBatchState(100).Config.Threshold)
	require.Equal(t, uint64(1), app.getBatchState(101).Config.Threshold)
	require.Equal(t, uint64(1), app.getBatchState(199).Config.Threshold)
	require.Equal(t, uint64(2), app.getBatchState(200).Config.Threshold)
	require.Equal(t, uint64(3), app.getBatchState(1000).Config.Threshold)
}

func TestAddConfig(t *testing.T) {
	app := NewShutterApp()

	err := app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers:         addr,
	})
	require.Nil(t, err)

	err = app.addConfig(BatchConfig{
		ConfigIndex:     2,
		StartBatchIndex: 99,
		Threshold:       1,
		Keypers:         addr,
	})
	require.NotNil(t, err, "Expected error, StartBatchIndex must not decrease")

	err = app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers:         addr,
	})
	require.NotNil(t, err, "Expected error, ConfigIndex must increase")

	err = app.addConfig(BatchConfig{
		ConfigIndex:     2,
		StartBatchIndex: 100,
		Threshold:       2,
		Keypers:         addr,
	})
	require.Nil(t, err)
}

func TestKeyGeneration(t *testing.T) {
	app := NewShutterApp()
	keypers := addresses[:3]

	err := app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       2,
		Keypers:         keypers,
	})
	require.Nil(t, err)
	res1 := app.deliverPublicKeyCommitment(
		&shmsg.PublicKeyCommitment{
			BatchIndex: 200,
			Commitment: crypto.FromECDSAPub(&keys[0].PublicKey),
		},
		keypers[0])
	require.Equal(
		t,
		abcitypes.ResponseDeliverTx{Code: 0, Events: []abcitypes.Event(nil)},
		res1)

	res2 := app.deliverPublicKeyCommitment(
		&shmsg.PublicKeyCommitment{
			BatchIndex: 200,
			Commitment: crypto.FromECDSAPub(&keys[1].PublicKey),
		},
		keypers[1])
	// We've reached the threshold, there should be an event of Type "shutter.pubkey-generated"
	require.Equal(
		t,
		abcitypes.ResponseDeliverTx{
			Code: 0,
			Events: []abcitypes.Event{
				{
					Type: "shutter.pubkey-generated",
					Attributes: []kv.Pair{
						{
							Key:   []byte("BatchIndex"),
							Value: []byte("200"),
						},
						{
							Key:   []byte("Pubkey"),
							Value: []byte(encodePubkeyForEvent(&keys[1].PublicKey)),
						},
					},
				},
			},
		},
		res2)
	res3 := app.deliverPublicKeyCommitment(
		&shmsg.PublicKeyCommitment{
			BatchIndex: 200,
			Commitment: crypto.FromECDSAPub(&keys[2].PublicKey),
		},
		keypers[2])
	require.Equal(
		t,
		abcitypes.ResponseDeliverTx{Code: 0, Events: []abcitypes.Event(nil)},
		res3)

	// --- Now let's deliver the SecretShare's
	ss1 := app.deliverSecretShare(
		&shmsg.SecretShare{
			BatchIndex: 200,
			Privkey:    crypto.FromECDSA(keys[0]),
		},
		keypers[0])
	require.Equal(t, abcitypes.ResponseDeliverTx{Code: 0, Events: []abcitypes.Event(nil)}, ss1)
	ss2 := app.deliverSecretShare(
		&shmsg.SecretShare{
			BatchIndex: 200,
			Privkey:    crypto.FromECDSA(keys[1]),
		},
		keypers[1])
	require.Equal(
		t,
		abcitypes.ResponseDeliverTx{
			Code: 0,
			Events: []abcitypes.Event{
				{
					Type: "shutter.privkey-generated",
					Attributes: []kv.Pair{
						{
							Key:   []byte("BatchIndex"),
							Value: []byte("200"),
						},
						{
							Key:   []byte("Privkey"),
							Value: []byte(encodePrivkeyForEvent(keys[1])),
						},
					},
				},
			},
		},
		ss2)
	ss3 := app.deliverSecretShare(
		&shmsg.SecretShare{
			BatchIndex: 200,
			Privkey:    crypto.FromECDSA(keys[2]),
		},
		keypers[2])
	require.Equal(t, abcitypes.ResponseDeliverTx{Code: 0, Events: []abcitypes.Event(nil)}, ss3)

	// encryption key signature collection
	key := crypto.FromECDSAPub(&keys[1].PublicKey)
	configContractAddress := common.HexToAddress("0x")
	preimage := EncryptionKeyPreimage(key, 200, configContractAddress)
	hash := crypto.Keccak256Hash(preimage)
	sig, err := crypto.Sign(hash.Bytes(), keys[0])
	require.Nil(t, err)
	attMsg := shmsg.EncryptionKeyAttestation{
		BatchIndex:            200,
		Key:                   key,
		ConfigContractAddress: configContractAddress.Bytes(),
		Signature:             sig,
	}
	res4 := app.deliverEncryptionKeyAttestation(&attMsg, keypers[0])
	expectedEvent := MakeEncryptionKeySignatureAddedEvent(0, 200, key, sig)
	require.Equal(t, abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: []abcitypes.Event{expectedEvent},
	}, res4)
}

func TestEncodePubkeyForEvent(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.Nil(t, err, "Could not generate key")
	encoded := encodePubkeyForEvent(&key.PublicKey)
	t.Logf("Encoded: %s", encoded)
	require.True(t, utf8.ValidString(encoded))

	decoded, err := DecodePubkeyFromEvent(encoded)
	require.Nil(t, err, "could not decode pubkey")
	t.Logf("Decoded: %+v", decoded)
	require.Equal(t, key.PublicKey, *decoded)
}
