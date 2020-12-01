package keyper

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/kr/pretty"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/app"
	"github.com/brainbot-com/shutter/shuttermint/crypto"
)

var (
	polynomial *crypto.Polynomial
	gammas     crypto.Gammas
)

func init() {
	var err error
	polynomial, err = crypto.RandomPolynomial(rand.Reader, 3)
	if err != nil {
		panic(err)
	}

	gammas = *polynomial.Gammas()
}

func TestCheckInEvent(t *testing.T) {
	sender := common.BigToAddress(big.NewInt(1))
	privateKeyECDSA, err := ethcrypto.GenerateKey()
	publicKey := ecies.ImportECDSAPublic(&privateKeyECDSA.PublicKey)
	require.Nil(t, err)
	appEv := app.MakeCheckInEvent(sender, publicKey)
	evInt, err := MakeEvent(appEv)
	require.Nil(t, err)
	ev, ok := evInt.(CheckInEvent)
	require.True(t, ok)
	require.Equal(t, sender, ev.Sender)
	require.True(t, ev.EncryptionPublicKey.ExportECDSA().Equal(&privateKeyECDSA.PublicKey))
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

func TestMakeEonStartedEvent(t *testing.T) {
	var eon uint64 = 10
	var batchIndex uint64 = 20
	appEv := app.MakeEonStartedEvent(eon, batchIndex)
	ev, err := MakeEvent(appEv)
	expectedEv := EonStartedEvent{
		Eon:        eon,
		BatchIndex: batchIndex,
	}
	require.Nil(t, err)
	require.Equal(t, expectedEv, ev)
}

func TestMakePolyCommitmentRegisteredEvent(t *testing.T) {
	var eon uint64 = 10
	sender := common.BytesToAddress([]byte("foo"))

	appEv := app.MakePolyCommitmentRegisteredEvent(&app.PolyCommitmentMsg{
		Sender: sender,
		Eon:    eon,
		Gammas: gammasToMsg(gammas),
	})
	pretty.Println(appEv)
	ev, err := MakeEvent(appEv)
	require.Nil(t, err)

	expectedEv := PolyCommitmentRegisteredEvent{
		Eon:    eon,
		Sender: sender,
		Gammas: &gammas,
	}
	require.Equal(t, expectedEv, ev)
}

// gammasToMsg converts the gammas to what the keyper sends to shuttermint
func gammasToMsg(gammas crypto.Gammas) [][]byte {
	// original implementation in NewPolyCommitmentMsg
	gammaBytes := [][]byte{}
	for _, gamma := range gammas {
		gammaBytes = append(gammaBytes, gamma.Marshal())
	}
	return gammaBytes
}

// gammasToEvent converts the gammas to what we get in a shuttermint event
func gammasToEvent(gammas crypto.Gammas) []byte {
	data := gammasToMsg(gammas) // this is what the keyper sends to shuttermint

	// Convert it to event data like newGammas defined in app/events.go
	var encoded []string
	for _, i := range data {
		encoded = append(encoded, hex.EncodeToString(i))
	}
	return []byte(strings.Join(encoded, ","))
}

func TestDecodeGammasFromEvent(t *testing.T) {
	eventValue := gammasToEvent(gammas)
	decoded, err := decodeGammasFromEvent(eventValue)
	require.Nil(t, err)
	require.Equal(t, gammas, decoded)
}
