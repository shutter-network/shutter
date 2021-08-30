package app

import (
	"encoding/base64"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"

	"github.com/shutter-network/shutter/shlib/shtest"
	"github.com/shutter-network/shutter/shuttermint/shmsg"
)

func TestNewShutterApp(t *testing.T) {
	app := NewShutterApp()
	assert.Equal(t, len(app.Configs), 1, "Configs should contain exactly one guard element")
	assert.Assert(t, is.DeepEqual(app.Configs[0], &BatchConfig{}), "Bad guard element")
}

func TestGetBatch(t *testing.T) {
	app := NewShutterApp()

	err := app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers:         addr,
	})
	assert.NilError(t, err)

	err = app.addConfig(BatchConfig{
		ConfigIndex:     2,
		StartBatchIndex: 200,
		Threshold:       2,
		Keypers:         addr,
	})
	assert.NilError(t, err)

	err = app.addConfig(BatchConfig{
		ConfigIndex:     3,
		StartBatchIndex: 300,
		Threshold:       3,
		Keypers:         addr,
	})
	assert.NilError(t, err)

	assert.Equal(t, uint64(0), app.getBatchState(0).Config.Threshold)
	assert.Equal(t, uint64(0), app.getBatchState(99).Config.Threshold)
	assert.Equal(t, uint64(1), app.getBatchState(100).Config.Threshold)
	assert.Equal(t, uint64(1), app.getBatchState(101).Config.Threshold)
	assert.Equal(t, uint64(1), app.getBatchState(199).Config.Threshold)
	assert.Equal(t, uint64(2), app.getBatchState(200).Config.Threshold)
	assert.Equal(t, uint64(3), app.getBatchState(1000).Config.Threshold)
}

func TestAddConfig(t *testing.T) {
	app := NewShutterApp()

	err := app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers:         addr,
	})
	assert.NilError(t, err)

	err = app.addConfig(BatchConfig{
		ConfigIndex:     2,
		StartBatchIndex: 99,
		Threshold:       1,
		Keypers:         addr,
	})
	assert.Assert(t, err != nil, "Expected error, StartBatchIndex must not decrease")

	err = app.addConfig(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers:         addr,
	})
	assert.Assert(t, err != nil, "Expected error, ConfigIndex must increase")

	err = app.addConfig(BatchConfig{
		ConfigIndex:     2,
		StartBatchIndex: 100,
		Threshold:       2,
		Keypers:         addr,
	})
	assert.NilError(t, err)
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
	assert.NilError(t, err)

	// don't accept signature from non-keyper
	res1 := app.deliverDecryptionSignature(
		&shmsg.DecryptionSignature{
			BatchIndex: 200,
			Signature:  []byte("signature"),
		},
		addresses[3],
	)
	assert.Assert(t, res1.IsErr())
	assert.Assert(t, is.Len(res1.Events, 0))

	// accept signature from keyper
	res2 := app.deliverDecryptionSignature(
		&shmsg.DecryptionSignature{
			BatchIndex: 200,
			Signature:  []byte("signature"),
		},
		keypers[0],
	)
	assert.Assert(t, res2.IsOK())
	assert.Equal(t, 1, len(res2.Events))

	ev := res2.Events[0]
	assert.Equal(t, "shutter.decryption-signature", ev.Type)
	assert.DeepEqual(t, []byte("BatchIndex"), ev.Attributes[0].Key)
	assert.DeepEqual(t, []byte("200"), ev.Attributes[0].Value)
	assert.DeepEqual(t, []byte("Sender"), ev.Attributes[1].Key)
	assert.DeepEqual(t, []byte(keypers[0].Hex()), ev.Attributes[1].Value)
	assert.DeepEqual(t, []byte("Signature"), ev.Attributes[2].Key)
	decodedSignature, _ := base64.RawURLEncoding.DecodeString(string(ev.Attributes[2].Value))
	assert.DeepEqual(t, []byte("signature"), decodedSignature)

	// don't accept another signature
	res3 := app.deliverDecryptionSignature(
		&shmsg.DecryptionSignature{
			BatchIndex: 200,
			Signature:  []byte("signature"),
		},
		keypers[0],
	)
	assert.Assert(t, res3.IsErr())
	assert.Assert(t, is.Len(res1.Events, 0))
}

func TestGobDKG(t *testing.T) {
	var eon uint64 = 201
	var err error
	keypers := addr
	dkg := NewDKGInstance(BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers:         keypers,
	}, eon)

	err = dkg.RegisterAccusationMsg(Accusation{
		Sender:  keypers[0],
		Eon:     eon,
		Accused: []common.Address{keypers[1]},
	})
	assert.NilError(t, err)

	err = dkg.RegisterApologyMsg(Apology{
		Sender:   keypers[0],
		Eon:      eon,
		Accusers: []common.Address{keypers[1]},
	})
	assert.NilError(t, err)

	err = dkg.RegisterPolyCommitmentMsg(PolyCommitment{
		Sender: keypers[0],
		Eon:    eon,
	})
	assert.NilError(t, err)

	err = dkg.RegisterPolyEvalMsg(PolyEval{
		Sender:         keypers[0],
		Eon:            eon,
		Receivers:      []common.Address{keypers[1]},
		EncryptedEvals: [][]byte{{}},
	})
	assert.NilError(t, err)

	shtest.EnsureGobable(t, &dkg, new(DKGInstance))
}
