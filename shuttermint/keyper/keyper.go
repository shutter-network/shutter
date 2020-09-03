package keyper

import (
	"context"
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"
	"golang.org/x/sync/errgroup"
)

// NewKeyper creates a new Keyper
func NewKeyper(SigningKey *ecdsa.PrivateKey, ShuttermintURL string) Keyper {
	return Keyper{
		SigningKey:          SigningKey,
		ShuttermintURL:      ShuttermintURL,
		events:              make(chan IEvent),
		batchIndexToChannel: make(map[uint64]chan IEvent)}
}

// Run runs the keyper process. It determines the next BatchIndex and runs the key generation
// process for this BatchIndex and all following batches.

func (k *Keyper) Run() error {
	group, ctx := errgroup.WithContext(context.Background())
	var cl client.Client
	cl, err := http.New(k.ShuttermintURL, "/websocket")
	if err != nil {
		return err
	}

	err = cl.Start()

	if err != nil {
		return err
	}

	defer cl.Stop()
	query := "tx.height > 3"

	txs, err := cl.Subscribe(ctx, "test-client", query)
	k.txs = txs
	if err != nil {
		return err
	}

	group.Go(k.dispatchTxs)
	group.Go(k.startNewBatches)
	group.Go(k.dispatchEvents)

	err = group.Wait()
	return err
}

func (k *Keyper) dispatchTxs() error {
	for tx := range k.txs {
		d := tx.Data.(types.EventDataTx)
		for _, ev := range d.TxResult.Result.Events {
			x, err := MakeEvent(ev)
			if err != nil {
				return err
			}
			k.dispatchEvent(x)
		}
	}
	return nil
}

func (k *Keyper) startNewBatches() error {
	var cl client.Client
	cl, err := http.New(k.ShuttermintURL, "/websocket")
	if err != nil {
		return err
	}

	for batchIndex := NextBatchIndex(time.Now()); ; batchIndex++ {
		bp := k.startBatch(batchIndex, cl)
		// The following waits for the start of the previous round. This is done on
		// purpose, because we generate keys in keyper.Run as a first step and then wait
		// for the start time
		SleepUntil(bp.PublicKeyGenerationStartTime)
	}
}

func (k *Keyper) startBatch(batchIndex uint64, cl client.Client) BatchParams {
	bp := NewBatchParams(batchIndex)
	ch := make(chan IEvent, 2)
	k.mux.Lock()
	k.batchIndexToChannel[batchIndex] = ch
	k.mux.Unlock()
	go Run(bp, NewMessageSender(cl, k.SigningKey), ch)
	return bp
}

func (k *Keyper) dispatchEvents() error {
	for ev := range k.events {
		k.dispatchEvent(ev)
	}
	return nil
}

func (k *Keyper) dispatchEventToBatch(BatchIndex uint64, ev IEvent) {
	k.mux.Lock()
	ch, ok := k.batchIndexToChannel[BatchIndex]
	k.mux.Unlock()

	if ok {
		ch <- ev
	}
}

func (k *Keyper) dispatchEvent(ev IEvent) {
	log.Printf("Dispatching event: %+v", ev)
	switch e := ev.(type) {
	case PubkeyGeneratedEvent:
		k.dispatchEventToBatch(e.BatchIndex, e)
	case PrivkeyGeneratedEvent:
		k.dispatchEventToBatch(e.BatchIndex, e)
	case BatchConfigEvent:
		_ = e
	default:
		panic("unknown type")
	}
}
