package main

import (
	"time"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
)

func main() {
	for batchIndex := keyper.NextBatchIndex(time.Now()); ; batchIndex++ {
		bp := keyper.NewBatchParams(batchIndex)
		go keyper.Run(bp)
		// The following waits for the start of the previous round. This is done on
		// purpose, because we generate keys in keyper.Run as a first step and then wait
		// for the start time
		keyper.SleepUntil(bp.PublicKeyGenerationStartTime)
	}
}
