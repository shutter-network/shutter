package shutterevents_test

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/ecies"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/brainbot-com/shutter/shuttermint/app"
	"github.com/brainbot-com/shutter/shuttermint/crypto"
	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
)

var (
	polynomial *crypto.Polynomial
	gammas     crypto.Gammas
	eon        = uint64(64738)
	sender     = common.BytesToAddress([]byte("foo"))
	addresses  = []common.Address{
		common.BigToAddress(big.NewInt(1)),
		common.BigToAddress(big.NewInt(2)),
		common.BigToAddress(big.NewInt(3)),
	}
)

func init() {
	var err error
	polynomial, err = crypto.RandomPolynomial(rand.Reader, 3)
	if err != nil {
		panic(err)
	}

	gammas = *polynomial.Gammas()
}

// mkeq ensures that calling MakeEvent on the given app event returns the expected IEvent
func mkeq(t *testing.T, appEv abcitypes.Event, expected shutterevents.IEvent) {
	ev, err := shutterevents.MakeEvent(appEv)
	require.Nil(t, err)
	require.Equal(t, expected, ev)
}

func TestAccusation(t *testing.T) {
	appEv := app.MakeAccusationRegisteredEvent(&app.Accusation{
		Eon:     eon,
		Sender:  sender,
		Accused: addresses,
	})
	mkeq(t, appEv, shutterevents.Accusation{
		Eon:     eon,
		Sender:  sender,
		Accused: addresses,
	})
}

func TestApology(t *testing.T) {
	accusers := addresses
	var polyEval []*big.Int
	var polyEvalsBytes [][]byte
	for i := 0; i < len(accusers); i++ {
		eval := big.NewInt(int64(100 + i))
		polyEval = append(polyEval, eval)
		polyEvalsBytes = append(polyEvalsBytes, eval.Bytes())
	}
	appEv := app.MakeApologyRegisteredEvent(&app.ApologyMsg{
		Eon:       eon,
		Sender:    sender,
		Accusers:  addresses,
		PolyEvals: polyEvalsBytes,
	})
	mkeq(t, appEv, shutterevents.Apology{
		Eon:      eon,
		Sender:   sender,
		Accusers: addresses,
		PolyEval: polyEval,
	})
}

func TestMakeEventBatchConfig(t *testing.T) {
	configIndex := uint64(0xffffffffffffffff)
	appEv := app.MakeBatchConfigEvent(111, 2, addresses, configIndex)
	mkeq(t, appEv, shutterevents.BatchConfig{
		StartBatchIndex: 111,
		Threshold:       2,
		Keypers:         addresses,
		ConfigIndex:     configIndex,
	})
}

func TestCheckInEvent(t *testing.T) {
	privateKeyECDSA, err := ethcrypto.GenerateKey()
	require.Nil(t, err)
	publicKey := ecies.ImportECDSAPublic(&privateKeyECDSA.PublicKey)
	appEv := app.MakeCheckInEvent(sender, publicKey)
	mkeq(t, appEv, shutterevents.CheckIn{Sender: sender, EncryptionPublicKey: publicKey})
}

func TestMakeEonStartedEvent(t *testing.T) {
	var batchIndex uint64 = 20
	appEv := app.MakeEonStartedEvent(eon, batchIndex)

	mkeq(t, appEv, shutterevents.EonStarted{Eon: eon, BatchIndex: batchIndex})
}

func TestMakePolyCommitmentRegisteredEvent(t *testing.T) {
	appEv := app.MakePolyCommitmentRegisteredEvent(&app.PolyCommitmentMsg{
		Sender: sender,
		Eon:    eon,
		Gammas: gammasToMsg(gammas),
	})
	mkeq(t, appEv, shutterevents.PolyCommitment{
		Eon:    eon,
		Sender: sender,
		Gammas: &gammas,
	})
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

// // gammasToEvent converts the gammas to what we get in a shuttermint event
// func gammasToEvent(gammas crypto.Gammas) []byte {
//	data := gammasToMsg(gammas) // this is what the keyper sends to shuttermint

//	// Convert it to event data like newGammas defined in app/events.go
//	var encoded []string
//	for _, i := range data {
//		encoded = append(encoded, hex.EncodeToString(i))
//	}
//	return []byte(strings.Join(encoded, ","))
// }

// func TestDecodeGammasFromEvent(t *testing.T) {
//	eventValue := gammasToEvent(gammas)
//	decoded, err := decodeGammas(eventValue)
//	require.Nil(t, err)
//	require.Equal(t, gammas, decoded)
// }
