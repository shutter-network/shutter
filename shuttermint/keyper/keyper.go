package keyper

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

const newHeadersSize = 8

// NewKeyper creates a new Keyper
func NewKeyper(kc KeyperConfig) Keyper {
	return Keyper{
		Config:                kc,
		scheduledBatchConfigs: make(map[uint64]contract.BatchConfig),
		batches:               make(map[uint64]*BatchState),
		newHeaders:            make(chan *types.Header, newHeadersSize),
	}
}

// NewBatchConfigStarted creates a new BatchConfigStarted message
func NewBatchConfigStarted(configIndex uint64) *shmsg.Message {
	return &shmsg.Message{
		Payload: &shmsg.Message_BatchConfigStarted{
			BatchConfigStarted: &shmsg.BatchConfigStarted{
				BatchConfigIndex: configIndex,
			},
		},
	}
}

func (kpr *Keyper) init() error {
	ethcl, err := ethclient.Dial(kpr.Config.EthereumURL)
	if err != nil {
		return err
	}
	kpr.ethcl = ethcl
	group, ctx := errgroup.WithContext(context.Background())
	kpr.ctx = ctx
	kpr.group = group
	cc, err := contract.NewConfigContract(kpr.Config.ConfigContractAddress, kpr.ethcl)
	if err != nil {
		return err
	}
	kpr.configContract = cc
	return nil
}

// Run runs the keyper process. It determines the next BatchIndex and runs the key generation
// process for this BatchIndex and all following batches.
func (kpr *Keyper) Run() error {
	if err := kpr.checkEthereumWebsocketURL(); err != nil {
		return err
	}

	log.Printf("Running keyper with address %s", kpr.Config.Address().Hex())
	err := kpr.init()
	if err != nil {
		return err
	}

	var shmcl client.Client
	shmcl, err = http.New(kpr.Config.ShuttermintURL, "/websocket")
	if err != nil {
		return errors.Wrapf(err, "create shuttermint client at %s", kpr.Config.ShuttermintURL)
	}

	err = shmcl.Start()
	if err != nil {
		return errors.Wrapf(err, "start shuttermint client at %s", kpr.Config.ShuttermintURL)
	}
	defer shmcl.Stop()

	ms := NewMessageSender(shmcl, kpr.Config.SigningKey)
	kpr.ms = &ms

	query := "tx.height > 3"
	txs, err := shmcl.Subscribe(kpr.ctx, "test-client", query)
	kpr.txs = txs
	if err != nil {
		return err
	}

	kpr.group.Go(kpr.dispatchTxs)
	kpr.group.Go(kpr.startNewBatches)
	kpr.group.Go(kpr.watchMainChainHeadBlock)
	kpr.group.Go(kpr.watchMainChainLogs)
	err = kpr.group.Wait()
	return err
}

// IsWebsocketURL returns true iff the given URL is a websocket URL, i.e. if it starts with ws://
// or wss://. This is needed for the watchMainChainHeadBlock method.
func IsWebsocketURL(url string) bool {
	return strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://")
}

func (kpr *Keyper) checkEthereumWebsocketURL() error {
	if !IsWebsocketURL(kpr.Config.EthereumURL) {
		err := fmt.Errorf("must use ws:// or wss:// URL, have %s", kpr.Config.EthereumURL)
		log.Printf("Error: %s", err)
		return err
	}
	return nil
}

func (kpr *Keyper) watchMainChainHeadBlock() error {
	cl, err := ethclient.Dial(kpr.Config.EthereumURL)
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
		case <-kpr.ctx.Done():
			return nil
		case err := <-subscription.Err():
			return err
		case header := <-headers:
			kpr.newHeaders <- header
			kpr.handleNewBlockHeader(header)
			kpr.dispatchNewBlockHeader(header)
		}
	}
}

func (kpr *Keyper) watchMainChainLogs() error {
	cl, err := ethclient.Dial(kpr.Config.EthereumURL)
	if err != nil {
		return err
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			kpr.Config.ConfigContractAddress,
		},
	}
	logs := make(chan types.Log)
	sub, err := cl.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return err
	}

	for {
		select {
		case <-kpr.ctx.Done():
			return nil
		case err := <-sub.Err():
			return err
		case log := <-logs:
			kpr.dispatchMainChainLog(log)
		}
	}
}

func (kpr *Keyper) handleNewBlockHeader(header *types.Header) {
	for i, c := range kpr.scheduledBatchConfigs {
		if header.Number.Cmp(c.StartBlockNumber) >= 0 {
			msg := NewBatchConfigStarted(i)
			err := kpr.ms.SendMessage(msg)
			if err == nil {
				log.Printf("BatchConfigStarted message sent for config %d", i)
			} else {
				log.Printf("Failed to send start message for batch config %d: %v", i, err)
			}

			delete(kpr.scheduledBatchConfigs, i)
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

func (kpr *Keyper) dispatchMainChainLog(l types.Log) {
	switch l.Address {
	case kpr.Config.ConfigContractAddress:
		// TODO: figure out how to effectively distinguish event types
		ev, err := kpr.configContract.ParseConfigScheduled(l)
		if err != nil {
			log.Printf("Failed to parse ConfigScheduled event: %v", err)
			return
		}
		kpr.handleConfigScheduledEvent(ev)
	default:
		log.Printf("Received unexpected event from %s", l.Address.Hex())
	}
}

func (kpr *Keyper) handleConfigScheduledEvent(ev *contract.ConfigContractConfigScheduled) {
	index := ev.NumConfigs.Uint64() - 1
	config, err := kpr.configContract.GetConfigByIndex(nil, index)
	if err != nil {
		log.Printf("Failed to fetch config from main chain to vote on it: %v", err)
		return
	}

	bc := NewBatchConfig(config.StartBatchIndex.Uint64(), config.Keypers, config.Threshold.Uint64())
	err = kpr.ms.SendMessage(bc)
	if err != nil {
		log.Printf("Failed to send batch config vote: %v", err)
	}

	kpr.scheduledBatchConfigs[index] = config
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

// waitCurrentBlockHeaders waits for a block header that is at max 1 Minute in the past. We do want
// to ignore block headers we may get when syncing.
func (kpr *Keyper) waitCurrentBlockHeader() *types.Header {
	for {
		select {
		case <-kpr.ctx.Done():
			return nil
		case header := <-kpr.newHeaders:
			if time.Since(time.Unix(int64(header.Time), 0)) < time.Minute {
				return header
			}
		}
	}
}

func (kpr *Keyper) startNewBatches() error {
	var cl client.Client
	cl, err := http.New(kpr.Config.ShuttermintURL, "/websocket")
	if err != nil {
		return err
	}
	_ = cl

	// We wait for block header that is not too old.

	header := kpr.waitCurrentBlockHeader()
	if header == nil {
		return nil
	}

	nextBatchIndex, err := kpr.configContract.NextBatchIndex(header.Number.Uint64())
	if err != nil {
		return err
	}
	lastBatchIndex := nextBatchIndex - 1 // we'll start only the batch for nextBatchIndex
	for {
		for batchIndex := lastBatchIndex + 1; batchIndex <= nextBatchIndex; batchIndex++ {
			log.Printf("Starting batch #%d", batchIndex)
			err = kpr.startBatch(nextBatchIndex, cl)
			if err != nil {
				log.Printf("Error: could not start batch #%d: %s", nextBatchIndex, err)
			}
		}
		lastBatchIndex = nextBatchIndex

		select {
		case <-kpr.ctx.Done():
			return nil
		case header := <-kpr.newHeaders:
			nextBatchIndex, err = kpr.configContract.NextBatchIndex(header.Number.Uint64())
			if err != nil {
				return err
			}
		}
	}
}

func (kpr *Keyper) removeBatch(batchIndex uint64) {
	log.Printf("Batch %d finished", batchIndex)
	kpr.mux.Lock()
	defer kpr.mux.Unlock()
	delete(kpr.batches, batchIndex)
}

func (kpr *Keyper) startBatch(batchIndex uint64, cl client.Client) error {
	bp, err := kpr.configContract.QueryBatchParams(nil, batchIndex)
	if err != nil {
		return err
	}

	ms := NewMessageSender(cl, kpr.Config.SigningKey)
	cc := NewContractCaller(
		kpr.Config.EthereumURL,
		kpr.Config.SigningKey,
		kpr.Config.KeyBroadcastingContractAddress,
	)
	batch := NewBatchState(bp, kpr.Config, &ms, &cc)

	kpr.mux.Lock()
	defer kpr.mux.Unlock()

	kpr.batches[batchIndex] = &batch

	go func() {
		defer kpr.removeBatch(batchIndex)

		// XXX we should check against what is configured in shuttermint
		// Here is the code we used to get the config from shuttermint:

		// bc, err := queryBatchConfig(cl, batchIndex)
		// if err != nil {
		//	log.Print("Error while trying to query batch config:", err)
		//	return
		// }
		// batch.BatchParams.BatchConfig = bc

		batch.Run()
	}()

	return nil
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
func (c *KeyperConfig) Address() common.Address {
	return crypto.PubkeyToAddress(c.SigningKey.PublicKey)
}
