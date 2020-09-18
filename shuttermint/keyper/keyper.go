package keyper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"golang.org/x/sync/errgroup"
)

// NewKeyper creates a new Keyper
func NewKeyper(signingKey *ecdsa.PrivateKey, shuttermintURL string, ethereumURL string) Keyper {
	return Keyper{
		SigningKey:     signingKey,
		ShuttermintURL: shuttermintURL,
		EthereumURL:    ethereumURL,
		batches:        make(map[uint64]*BatchState),
	}
}

// Run runs the keyper process. It determines the next BatchIndex and runs the key generation
// process for this BatchIndex and all following batches.
func (kpr *Keyper) Run() error {
	log.Printf("Running keyper with address %s", kpr.Address().Hex())

	group, ctx := errgroup.WithContext(context.Background())
	kpr.ctx = ctx

	var cl client.Client
	cl, err := http.New(kpr.ShuttermintURL, "/websocket")
	if err != nil {
		return errors.Wrapf(err, "create shuttermint client at %s", kpr.ShuttermintURL)
	}

	err = cl.Start()

	if err != nil {
		return errors.Wrapf(err, "start shuttermint client at %s", kpr.ShuttermintURL)
	}

	defer cl.Stop()
	query := "tx.height > 3"

	txs, err := cl.Subscribe(ctx, "test-client", query)
	kpr.txs = txs
	if err != nil {
		return err
	}

	group.Go(kpr.dispatchTxs)
	group.Go(kpr.startNewBatches)
	group.Go(kpr.watchMainChainHeadBlock)
	err = group.Wait()
	return err
}

func (kpr *Keyper) watchMainChainHeadBlock() error {
	const waitDuration = 2 * time.Second
	if !(strings.HasPrefix(kpr.EthereumURL, "ws://") || strings.HasPrefix(kpr.EthereumURL, "wss://")) {
		err := fmt.Errorf("must use ws:// or wss:// URL, have %s", kpr.EthereumURL)
		log.Printf("Error: %s", err)
		return err
	}

	cl, err := ethclient.Dial(kpr.EthereumURL)
	if err != nil {
		return err
	}
	headers := make(chan *types.Header)

	subscription, err := cl.SubscribeNewHead(kpr.ctx, headers)
	if err != nil {
		return err
	}
	for {
		select {
		case err := <-subscription.Err():
			return err
		case header := <-headers:
			kpr.dispatchNewBlockHeader(header)
		}
	}
}

func (kpr *Keyper) dispatchNewBlockHeader(header *types.Header) {
	log.Printf("Dispatching new block #%d to %d batches\n", header.Number, len(kpr.batches))
	kpr.mux.Lock()
	defer kpr.mux.Unlock()

	for _, batch := range kpr.batches {
		batch.NewBlockHeader(header)
	}
}

func (kpr *Keyper) dispatchTxs() error {
	// Unfortunately the channel isn't being closed for us, so we manually check if we should
	// cancel
	for {
		select {
		case tx := <-kpr.txs:
			d := tx.Data.(tmtypes.EventDataTx)
			for _, ev := range d.TxResult.Result.Events {
				x, err := MakeEvent(ev)
				if err != nil {
					return err
				}
				kpr.dispatchEvent(x)
			}
		case <-kpr.ctx.Done():
			return nil
		}
	}
}

func (kpr *Keyper) startNewBatches() error {
	var cl client.Client
	cl, err := http.New(kpr.ShuttermintURL, "/websocket")
	if err != nil {
		return err
	}

	for batchIndex := NextBatchIndex(time.Now()); ; batchIndex++ {
		bp := kpr.startBatch(batchIndex, cl)
		// The following waits for the start of the previous round. This is done on
		// purpose, because we generate keys in keyper.Run as a first step and then wait
		// for the start time
		select {
		case <-kpr.ctx.Done():
			return nil
		case <-time.After(time.Until(bp.PublicKeyGenerationStartTime)):
		}
	}
}

func (kpr *Keyper) removeBatch(batchIndex uint64) {
	log.Printf("Batch %d finished", batchIndex)
	kpr.mux.Lock()
	defer kpr.mux.Unlock()
	delete(kpr.batches, batchIndex)
}

func (kpr *Keyper) startBatch(batchIndex uint64, cl client.Client) BatchParams {
	bp := NewBatchParams(batchIndex)

	ms := NewMessageSender(cl, kpr.SigningKey)
	cc := NewContractCaller(
		kpr.EthereumURL,
		kpr.SigningKey,
		common.HexToAddress("0x791c3f20f865c582A204134E0A64030Fc22D2E38"),
	)
	batch := NewBatchState(bp, kpr.SigningKey, &ms, &cc)

	kpr.mux.Lock()
	defer kpr.mux.Unlock()

	kpr.batches[batchIndex] = &batch

	go func() {
		defer kpr.removeBatch(batchIndex)

		bc, err := queryBatchConfig(cl, batchIndex)
		if err != nil {
			log.Print("Error while trying to query batch config:", err)
			return
		}
		batch.BatchParams.BatchConfig = bc

		batch.Run()
	}()

	return bp
}

func (kpr *Keyper) dispatchEventToBatch(batchIndex uint64, ev IEvent) {
	kpr.mux.Lock()
	defer kpr.mux.Unlock()

	batch, ok := kpr.batches[batchIndex]

	if ok {
		batch.dispatchShuttermintEvent(ev)
	}
}

func (kpr *Keyper) dispatchEvent(ev IEvent) {
	log.Printf("Dispatching event: %+v", ev)
	switch e := ev.(type) {
	case PubkeyGeneratedEvent:
		kpr.dispatchEventToBatch(e.BatchIndex, e)
	case PrivkeyGeneratedEvent:
		kpr.dispatchEventToBatch(e.BatchIndex, e)
	case BatchConfigEvent:
		_ = e
	case EncryptionKeySignatureAddedEvent:
		kpr.dispatchEventToBatch(e.BatchIndex, e)
	default:
		panic("unknown type")
	}
}

// Address returns the keyper's Ethereum address.
func (kpr *Keyper) Address() common.Address {
	return crypto.PubkeyToAddress(kpr.SigningKey.PublicKey)
}
