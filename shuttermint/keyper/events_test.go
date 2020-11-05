package keyper

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/app"
)

func TestMakeEventPrivkeyGenerated(t *testing.T) {
	privkey, err := crypto.GenerateKey()
	require.Nil(t, err)
	appEvent := app.MakePrivkeyGeneratedEvent(111, privkey)
	ev, err := MakeEvent(appEvent)
	require.Nil(t, err)
	require.Equal(t, PrivkeyGeneratedEvent{BatchIndex: 111, Privkey: privkey}, ev)
}

func TestMakeEventBatchConfig(t *testing.T) {
	var addresses []common.Address = []common.Address{
		common.BigToAddress(big.NewInt(1)),
		common.BigToAddress(big.NewInt(2)),
		common.BigToAddress(big.NewInt(3)),
	}

	appEvent := app.MakeBatchConfigEvent(111, 2, addresses)
	ev, err := MakeEvent(appEvent)
	require.Nil(t, err)
	require.Equal(t,
		BatchConfigEvent{
			StartBatchIndex: 111,
			Threshold:       2,
			Keypers:         addresses,
		},
		ev)
}

func TestMakeEventEncryptionSignatureAddedEvent(t *testing.T) {
	var keyperIndex uint64 = 3
	var batchIndex uint64 = 111
	key := []byte("key")
	sig := []byte("sig")
	appEvent := app.MakeEncryptionKeySignatureAddedEvent(keyperIndex, batchIndex, key, sig)
	ev, err := MakeEvent(appEvent)
	expectedEvent := EncryptionKeySignatureAddedEvent{
		KeyperIndex:   keyperIndex,
		BatchIndex:    batchIndex,
		EncryptionKey: key,
		Signature:     sig,
	}
	require.Nil(t, err)
	require.Equal(t, ev, expectedEvent)
}

func TestMakeNewDKGInstanceEvent(t *testing.T) {
	var eon uint64 = 10
	var configIndex uint64 = 20
	appEv := app.MakeNewDKGInstanceEvent(eon, configIndex)
	ev, err := MakeEvent(appEv)
	expectedEv := NewDKGInstanceEvent{
		Eon:         eon,
		ConfigIndex: configIndex,
	}
	require.Nil(t, err)
	require.Equal(t, expectedEv, ev)
}
