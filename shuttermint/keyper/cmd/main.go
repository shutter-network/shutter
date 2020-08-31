package main

import (
	"time"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}

	var cl client.Client
	cl, err = http.New("http://localhost:26657", "/websocket")
	if err != nil {
		panic(err)
	}

	for batchIndex := keyper.NextBatchIndex(time.Now()); ; batchIndex++ {
		bp := keyper.NewBatchParams(batchIndex)
		go keyper.Run(bp, keyper.NewMessageSender(cl, privateKey))
		// The following waits for the start of the previous round. This is done on
		// purpose, because we generate keys in keyper.Run as a first step and then wait
		// for the start time
		keyper.SleepUntil(bp.PublicKeyGenerationStartTime)
	}
}
