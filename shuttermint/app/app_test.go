package app

import (
	"encoding/base64"
	"testing"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

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

func TestAddDecryptionSignature(t *testing.T) {
	app := NewShutterApp()
	keypers := addresses[:3]
	err := app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       2,
		Keypers:         keypers,
	})
	require.Nil(t, err)

	// don't accept signature from non-keyper
	res1 := app.deliverDecryptionSignature(
		&shmsg.DecryptionSignature{
			BatchIndex: 200,
			Signature:  []byte("signature"),
		},
		addresses[3],
	)
	require.True(t, res1.IsErr())
	require.Empty(t, res1.Events)

	// accept signature from keyper
	res2 := app.deliverDecryptionSignature(
		&shmsg.DecryptionSignature{
			BatchIndex: 200,
			Signature:  []byte("signature"),
		},
		keypers[0],
	)
	require.True(t, res2.IsOK())
	require.Equal(t, 1, len(res2.Events))

	ev := res2.Events[0]
	require.Equal(t, "shutter.decryption-signature", ev.Type)
	require.Equal(t, []byte("BatchIndex"), ev.Attributes[0].Key)
	require.Equal(t, []byte("200"), ev.Attributes[0].Value)
	require.Equal(t, []byte("Sender"), ev.Attributes[1].Key)
	require.Equal(t, []byte(keypers[0].Hex()), ev.Attributes[1].Value)
	require.Equal(t, []byte("Signature"), ev.Attributes[2].Key)
	decodedSignature, _ := base64.RawURLEncoding.DecodeString(string(ev.Attributes[2].Value))
	require.Equal(t, []byte("signature"), decodedSignature)

	// don't accept another signature
	res3 := app.deliverDecryptionSignature(
		&shmsg.DecryptionSignature{
			BatchIndex: 200,
			Signature:  []byte("signature"),
		},
		keypers[0],
	)
	require.True(t, res3.IsErr())
	require.Empty(t, res1.Events)
}
