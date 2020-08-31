package keyper

import (
	"time"

	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
)

func (k Keyper) Run() error {
	var cl client.Client
	cl, err := http.New(k.ShuttermintURL, "/websocket")
	if err != nil {
		return err
	}

	for batchIndex := NextBatchIndex(time.Now()); ; batchIndex++ {
		bp := NewBatchParams(batchIndex)
		go Run(bp, NewMessageSender(cl, k.SigningKey))
		// The following waits for the start of the previous round. This is done on
		// purpose, because we generate keys in keyper.Run as a first step and then wait
		// for the start time
		SleepUntil(bp.PublicKeyGenerationStartTime)
	}

}
