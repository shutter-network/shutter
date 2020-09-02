package keyper

import (
	"math/big"
	"testing"

	"github.com/brainbot-com/shutter/shuttermint/app"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestMakeEventPrivkeyGenerated(t *testing.T) {
	privkey, err := crypto.GenerateKey()
	require.Nil(t, err)
	appEvent := app.MakePrivkeyGeneratedEvent(111, privkey)
	ev, err := MakeEvent(appEvent)
	require.Nil(t, err)
	require.Equal(t, PrivkeyGeneratedEvent{BatchIndex: 111, Privkey: privkey}, ev)
}

func TestMakeEventPubkeyGenerated(t *testing.T) {
	privkey, err := crypto.GenerateKey()
	require.Nil(t, err)
	appEvent := app.MakePubkeyGeneratedEvent(111, &privkey.PublicKey)
	ev, err := MakeEvent(appEvent)
	require.Nil(t, err)
	require.Equal(t, PubkeyGeneratedEvent{BatchIndex: 111, Pubkey: &privkey.PublicKey}, ev)
}

func TestMakeEventBatchConfig(t *testing.T) {
	var addresses []common.Address = []common.Address{common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2)), common.BigToAddress(big.NewInt(3))}

	appEvent := app.MakeBatchConfigEvent(111, 2, addresses)
	ev, err := MakeEvent(appEvent)
	require.Nil(t, err)
	require.Equal(t,
		BatchConfigEvent{
			StartBatchIndex: 111,
			Threshhold:      2,
			Keypers:         addresses},
		ev)
}
