package keyper

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

const (
	newHeadersSize       = 8
	balanceCheckInterval = 10 * time.Second
	warningGasThreshold  = 100 * 1000 * 1000
)

// NewKeyper creates a new Keyper
func NewKeyper(kc KeyperConfig) Keyper {
	return Keyper{
		Config:                kc,
		batchConfigs:          make(map[uint64]contract.BatchConfig),
		batches:               make(map[uint64]*BatchState),
		checkedIn:             false,
		newHeaders:            make(chan *types.Header, newHeadersSize),
		cipherExecutionParams: make(chan CipherExecutionParams),
		keyperEncryptionKeys:  make(map[common.Address]*ecies.PublicKey),
		dkg:                   nil,
	}
}

// NewBatchConfig creates a new BatchConfig message
func NewBatchConfig(
	startBatchIndex uint64,
	keypers []common.Address,
	threshold uint64,
	configContractAddress common.Address,
	configIndex uint64,
) *shmsg.Message {
	var addresses [][]byte
	for _, k := range keypers {
		addresses = append(addresses, k.Bytes())
	}
	return &shmsg.Message{
		Payload: &shmsg.Message_BatchConfig{
			BatchConfig: &shmsg.BatchConfig{
				StartBatchIndex:       startBatchIndex,
				Keypers:               addresses,
				Threshold:             threshold,
				ConfigContractAddress: configContractAddress.Bytes(),
				ConfigIndex:           configIndex,
			},
		},
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

// NewCheckIn creates a new CheckIn message
func NewCheckIn(validatorPublicKey []byte, encryptionKey *ecies.PublicKey) *shmsg.Message {
	encryptionKeyECDSA := encryptionKey.ExportECDSA()
	return &shmsg.Message{
		Payload: &shmsg.Message_CheckIn{
			CheckIn: &shmsg.CheckIn{
				ValidatorPublicKey:  validatorPublicKey,
				EncryptionPublicKey: crypto.CompressPubkey(encryptionKeyECDSA),
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

	shmcl, err := http.New(kpr.Config.ShuttermintURL, "/websocket")
	if err != nil {
		return errors.Wrapf(err, "create shuttermint client at %s", kpr.Config.ShuttermintURL)
	}
	kpr.shmcl = shmcl

	ms := NewRPCMessageSender(kpr.shmcl, kpr.Config.SigningKey)
	kpr.ms = &ms

	group, ctx := errgroup.WithContext(context.Background())
	kpr.ctx = ctx
	kpr.group = group

	cc, err := contract.NewConfigContract(kpr.Config.ConfigContractAddress, kpr.ethcl)
	if err != nil {
		return err
	}
	kpr.configContract = cc

	kbc, err := contract.NewKeyBroadcastContract(kpr.Config.KeyBroadcastContractAddress, kpr.ethcl)
	if err != nil {
		return err
	}
	kpr.keyBroadcastContract = kbc

	bc, err := contract.NewBatcherContract(kpr.Config.BatcherContractAddress, kpr.ethcl)
	if err != nil {
		return err
	}
	kpr.batcherContract = bc

	ec, err := contract.NewExecutorContract(kpr.Config.ExecutorContractAddress, kpr.ethcl)
	if err != nil {
		return err
	}
	kpr.executorContract = ec

	contractCaller := NewContractCaller(
		kpr.ethcl,
		kpr.Config.SigningKey,
		kpr.configContract,
		kpr.keyBroadcastContract,
		kpr.batcherContract,
		kpr.executorContract,
	)
	executor := Executor{
		ctx:                   kpr.ctx,
		client:                kpr.ethcl,
		cc:                    &contractCaller,
		cipherExecutionParams: kpr.cipherExecutionParams,
	}
	kpr.executor = executor

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

	err = kpr.shmcl.Start()
	if err != nil {
		return errors.Wrapf(err, "start shuttermint client at %s", kpr.Config.ShuttermintURL)
	}
	defer kpr.shmcl.Stop()

	query := "tx.height > 3"
	txs, err := kpr.shmcl.Subscribe(kpr.ctx, "test-client", query)
	kpr.txs = txs
	if err != nil {
		return err
	}

	err = kpr.fetchCurrentDKG(kpr.ctx)
	if err != nil {
		return err
	}

	startBlock, err := kpr.startUp()
	if err != nil {
		return err
	}
	kpr.startBlock = startBlock

	kpr.group.Go(kpr.dispatchTxs)
	kpr.group.Go(kpr.startNewBatches)
	kpr.group.Go(kpr.watchMainChainHeadBlock)
	kpr.group.Go(kpr.watchMainChainLogs)
	kpr.group.Go(kpr.watchBalance)
	kpr.group.Go(kpr.executor.Run)
	err = kpr.group.Wait()
	return err
}

// queries the current DKG Status and may start a new DKG process
func (kpr *Keyper) fetchCurrentDKG(ctx context.Context) error {
	page := 1
	perPage := 50
	res, err := kpr.shmcl.TxSearch(ctx, "shutter.new-dkg-instance.Eon EXISTS", false, &page, &perPage, "desc")
	if err != nil {
		return nil
	}

	for _, tx := range res.Txs {
		fmt.Printf("=== tx height=%d\n", tx.Height)
		events := tx.TxResult.GetEvents()
		for _, rawEvent := range events {
			e, err := MakeEvent(rawEvent)
			if err != nil {
				return err
			}
			switch event := e.(type) {
			case NewDKGInstanceEvent:
				// XXX Unconditionally starting the DKG for this eon is certainly
				// wrong. We need a way to decide if we should start it. For the
				// moment it helps me get a first version working.
				kpr.startNewDKGInstance(event)
			default:
				// even though we only ask for shutter.new-dk-instance events the transactions will also contain other events.
				// Let's ignore them
			}
		}
	}
	return nil
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

func (kpr *Keyper) startUp() (*big.Int, error) {
	startupHeader, err := kpr.ethcl.HeaderByNumber(kpr.ctx, nil)
	if err != nil {
		return nil, err
	}
	opts := &bind.CallOpts{
		BlockNumber: startupHeader.Number,
		Context:     kpr.ctx,
	}

	err = kpr.initializeConfigs(opts)
	if err != nil {
		return nil, err
	}

	err = kpr.doInitialCheckIn(opts)
	if err != nil {
		return nil, err
	}

	return startupHeader.Number, nil
}

// initializeConfigs fills batchConfigs with the current as well as all already scheduled future
// configs from the config contract.
func (kpr *Keyper) initializeConfigs(opts *bind.CallOpts) error {
	numConfigs, err := kpr.configContract.NumConfigs(opts)
	if err != nil {
		return err
	}

	for i := numConfigs; i >= 1; i-- {
		config, err := kpr.configContract.GetConfigByIndex(opts, i-1)
		if err != nil {
			return err
		}

		kpr.batchConfigs[i] = config

		if config.StartBlockNumber < opts.BlockNumber.Uint64() {
			break
		}
	}

	return nil
}

func (kpr *Keyper) doInitialCheckIn(opts *bind.CallOpts) error {
	configs, err := kpr.configContract.CurrentAndFutureConfigs(opts, opts.BlockNumber.Uint64())
	if err != nil {
		return err
	}

	// Check if we are a keyper now or in some scheduled config. If not, no need to check in.
	isKeyper := false
	for _, config := range configs {
		if config.IsKeyper(kpr.Config.Address()) {
			isKeyper = true
			break
		}
	}
	if !isKeyper {
		return nil
	}

	// Check that we're not already checked in. If not, no need to check in.
	checkedIn, err := queryCheckedIn(kpr.shmcl, kpr.Config.Address())
	if err != nil {
		return err
	}
	if checkedIn {
		log.Println("already checked in")
		kpr.checkedIn = true
		return nil
	}

	return kpr.sendCheckIn()
}

func (kpr *Keyper) watchMainChainHeadBlock() error {
	// We might have missed some blocks between startup (i.e. kpr.startBlock) and setting up the
	// header subscription. We don't rely on receiving all blocks though, so this is fine.
	headers := make(chan *types.Header)
	subscription, err := kpr.ethcl.SubscribeNewHead(kpr.ctx, headers)
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
		FromBlock: kpr.startBlock,
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

func (kpr *Keyper) maybeSendConfigVote() error {
	// don't vote if previous vote is still being processed
	_, voted, err := queryVote(kpr.shmcl, kpr.Config.Address())
	if err != nil {
		return err
	}
	if voted {
		return nil
	}

	// don't vote if last config has not started yet
	lastConfig, err := queryLastBatchConfig(kpr.shmcl)
	if err != nil {
		return err
	}
	if !lastConfig.Started {
		return nil
	}

	// don't vote if there's no newer config
	nextConfigIndex := lastConfig.ConfigIndex + 1
	ok := func() bool {
		kpr.Lock()
		defer kpr.Unlock()
		_, ok := kpr.batchConfigs[nextConfigIndex]
		return ok
	}()
	if !ok {
		return nil
	}

	// don't vote if we're not a keyper in current config
	if !lastConfig.IsKeyper(kpr.Config.Address()) {
		return nil
	}

	return kpr.sendConfigVote(nextConfigIndex)
}

func (kpr *Keyper) sendConfigVote(configIndex uint64) error {
	config, ok := func() (contract.BatchConfig, bool) {
		kpr.Lock()
		defer kpr.Unlock()
		c, ok := kpr.batchConfigs[configIndex]
		return c, ok
	}()
	if !ok {
		return fmt.Errorf("cannot vote on config %d as it is unknown", configIndex)
	}

	msg := NewBatchConfig(
		config.StartBatchIndex,
		config.Keypers,
		config.Threshold,
		kpr.Config.ConfigContractAddress,
		configIndex,
	)
	err := kpr.ms.SendMessage(kpr.ctx, msg)
	if err != nil {
		return err
	}
	log.Printf("voted for config %d", configIndex)
	return nil
}

func (kpr *Keyper) maybeSendStartVote(blockNumber uint64) error {
	lastConfig, err := queryLastBatchConfig(kpr.shmcl)
	if err != nil {
		return err
	}

	// don't vote to start if last config is already started
	if lastConfig.Started {
		return nil
	}

	// don't vote to start if last config is unknown
	localConfig, configKnown := func() (contract.BatchConfig, bool) {
		kpr.Lock()
		defer kpr.Unlock()
		localConfig, configKnown := kpr.batchConfigs[lastConfig.ConfigIndex]
		return localConfig, configKnown
	}()
	if !configKnown && lastConfig.ConfigIndex > 0 {
		log.Printf("not voting to start unknown config %d", lastConfig.ConfigIndex)
		return nil
	}

	// don't vote to start if start block number is still in the future
	if localConfig.StartBlockNumber > blockNumber {
		return nil
	}

	// don't vote to start if we're not a keyper
	if !localConfig.IsKeyper(kpr.Config.Address()) {
		return nil
	}

	// otherwise vote
	msg := NewBatchConfigStarted(lastConfig.ConfigIndex)
	err = kpr.ms.SendMessage(kpr.ctx, msg)
	if err != nil {
		return err
	}

	log.Printf("voted to start batch config with index %d", lastConfig.ConfigIndex)
	return nil
}

func (kpr *Keyper) handleNewBlockHeader(header *types.Header) {
	err := kpr.maybeSendConfigVote()
	if err != nil {
		log.Printf("error during send config vote step: %v", err)
	}

	err = kpr.maybeSendStartVote(header.Number.Uint64())
	if err != nil {
		log.Printf("error during send start vote step: %v", err)
	}
}

func (kpr *Keyper) dispatchNewBlockHeader(header *types.Header) {
	log.Printf("Dispatching new block #%d to %d batches\n", header.Number, len(kpr.batches))
	kpr.Lock()
	defer kpr.Unlock()

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
	index := ev.NumConfigs - 1
	config, err := kpr.configContract.GetConfigByIndex(nil, index)
	if err != nil {
		log.Printf("Failed to fetch config from main chain to vote on it: %v", err)
		return
	}

	func() {
		kpr.Lock()
		defer kpr.Unlock()
		kpr.batchConfigs[index] = config
	}()

	bc := NewBatchConfig(
		config.StartBatchIndex,
		config.Keypers,
		config.Threshold,
		kpr.Config.ConfigContractAddress,
		index,
	)
	err = kpr.ms.SendMessage(kpr.ctx, bc)
	if err != nil {
		log.Printf("Failed to send batch config vote: %v", err)
	}

	err = kpr.maybeSendCheckIn(config)
	if err != nil {
		log.Printf("Failed to send checkin message: %v", err)
	}
}

// send check in if we aren't checked in already and if we are a member of the keyper set in the
// given config.
func (kpr *Keyper) maybeSendCheckIn(config contract.BatchConfig) error {
	if kpr.checkedIn {
		return nil
	}

	if _, isKeyper := config.KeyperIndex(kpr.Config.Address()); !isKeyper {
		return nil
	}

	return kpr.sendCheckIn()
}

func (kpr *Keyper) sendCheckIn() error {
	validatorPublicKey, ok := kpr.Config.ValidatorKey.Public().(ed25519.PublicKey)
	if !ok {
		panic("Failed to assert type")
	}

	msg := NewCheckIn([]byte(validatorPublicKey), &kpr.Config.EncryptionKey.PublicKey)
	err := kpr.ms.SendMessage(kpr.ctx, msg)
	if err != nil {
		return err
	}

	log.Printf("Check in message sent")
	kpr.checkedIn = true
	return nil
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
				// XXX In case MakeEvent fails, returning with an error here, will
				// abort the whole program. At the moment a malicious keyper is
				// able to send a PolyCommitmentRegistered message with a malformed
				// Gammas field to trigger an error here.
				// We should fix that in shuttermint and/or ignore the error here.
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
	// We wait for block header that is not too old.
	header := kpr.waitCurrentBlockHeader()
	if header == nil {
		return nil
	}

	nextBatchIndex, err := kpr.configContract.NextBatchIndex(header.Number.Uint64())
	if err != nil {
		return err
	}
	lastBatchIndex := nextBatchIndex - 1

	for {
		// Start batches between lastBatchIndex and nextBatchIndex. Usually this will be only a
		// single one. If the batch is inactive (or we fail to query the batch config to check),
		// break out of the loop in order to wait for the next header and try again.
		for batchIndex := lastBatchIndex + 1; batchIndex <= nextBatchIndex; batchIndex++ {
			bp, err := kpr.configContract.QueryBatchParams(nil, batchIndex)
			if err != nil {
				log.Printf("Failed to query batch params for batch #%d", batchIndex)
				break
			}
			if !bp.BatchConfig.IsActive() {
				log.Printf("Batching is inactive at #%d", batchIndex)
				break
			}

			log.Printf("Starting batch #%d", batchIndex)
			err = kpr.startBatch(bp)
			lastBatchIndex = batchIndex
			if err != nil {
				log.Printf("Error: could not start batch #%d: %s", batchIndex, err)
			}
		}

		// Wait for next header and update nextBatchIndex accordingly.
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
	kpr.Lock()
	defer kpr.Unlock()
	delete(kpr.batches, batchIndex)
}

func (kpr *Keyper) startBatch(bp BatchParams) error {
	if !bp.BatchConfig.IsActive() {
		log.Printf("shutter is turned off right now, waiting...")
		return nil
	}

	cc := NewContractCaller(
		kpr.ethcl,
		kpr.Config.SigningKey,
		kpr.configContract,
		kpr.keyBroadcastContract,
		kpr.batcherContract,
		kpr.executorContract,
	)
	batch := NewBatchState(bp, kpr.Config, kpr.ms, &cc, kpr.cipherExecutionParams)

	kpr.Lock()
	defer kpr.Unlock()

	kpr.batches[bp.BatchIndex] = &batch

	go func() {
		defer kpr.removeBatch(bp.BatchIndex)

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
	kpr.Lock()
	defer kpr.Unlock()

	batch, ok := kpr.batches[batchIndex]

	if ok {
		batch.dispatchShuttermintEvent(ev)
	}
}

func (kpr *Keyper) dispatchEventToDKG(eon uint64, ev IEvent) {
	kpr.Lock()
	defer kpr.Unlock()

	// we should probably maintain a map of eon -> dkg
	if kpr.dkg == nil || kpr.dkg.Eon != eon {
		log.Printf("dkg is not active, cannot dispatch %+v", ev)
	}
	kpr.dkg.dispatchShuttermintEvent(ev)
}

func (kpr *Keyper) dispatchEvent(ev IEvent) {
	switch e := ev.(type) {
	case CheckInEvent:
		kpr.handleCheckInEvent(e)
	case BatchConfigEvent:
	case DecryptionSignatureEvent:
		kpr.dispatchEventToBatch(e.BatchIndex, e)
	case NewDKGInstanceEvent:
		kpr.startNewDKGInstance(e)
	case PolyCommitmentRegisteredEvent:
		kpr.dispatchEventToDKG(e.Eon, e)
	default:
		panic("unknown event type")
	}
}

func (kpr *Keyper) startNewDKGInstance(ev NewDKGInstanceEvent) {
	if kpr.dkg == nil {
		log.Printf("Starting DKG for eon %d", ev.Eon)
	} else {
		log.Printf("Aborting DKG for eon %d to start DKG for eon %d", kpr.dkg.Eon, ev.Eon)
		kpr.dkg = nil
	}

	config, err := kpr.configContract.GetConfigByIndex(nil, ev.ConfigIndex)
	if err != nil {
		log.Printf("Failed to fetch config from main chain to start DKG: %v", err)
		return
	}

	// copy keyper encryption keys to avoid concurrent access
	keyperEncryptionKeys := make(map[common.Address]*ecies.PublicKey)
	for k, v := range kpr.keyperEncryptionKeys {
		keyperEncryptionKeys[k] = v
	}

	dkg, err := NewDKGInstance(ev.Eon, config, kpr.Config, kpr.ms, keyperEncryptionKeys)
	if err != nil {
		log.Printf("Error starting DKG instance: %s", err)
	}
	kpr.dkg = dkg
	ctx, cancel := context.WithCancel(kpr.ctx)
	go func() {
		defer cancel()
		err = dkg.Run(ctx)
		if err != nil {
			log.Fatalf("DKG process for eon %d failed: %s", ev.Eon, err)
		}
	}()
}

func (kpr *Keyper) handleCheckInEvent(ev CheckInEvent) {
	kpr.Lock()
	defer kpr.Unlock()
	kpr.keyperEncryptionKeys[ev.Sender] = ev.EncryptionPublicKey
}

func (kpr *Keyper) watchBalance() error {
	for {
		balance, err := kpr.ethcl.BalanceAt(kpr.ctx, kpr.Config.Address(), nil)
		if err != nil {
			return fmt.Errorf("error checking keyper balance: %w", err)
		}

		gasPrice, err := kpr.ethcl.SuggestGasPrice(kpr.ctx)
		if err != nil {
			return fmt.Errorf("error checking current gas price: %w", err)
		}

		availableGas := new(big.Int).Div(balance, gasPrice)
		if availableGas.Cmp(new(big.Int).SetUint64(warningGasThreshold)) < 0 {
			balanceEth := new(big.Float).SetInt(balance)
			balanceEth.Quo(balanceEth, new(big.Float).SetFloat64(1e18))
			log.Printf(
				"Warning: Balance of keyper account is running low, only %s ETH left at address %s",
				balanceEth.String(),
				kpr.Config.Address().Hex(),
			)
		}

		select {
		case <-kpr.ctx.Done():
			return nil
		case <-time.After(balanceCheckInterval):
			continue
		}
	}
}

// Address returns the keyper's Ethereum address.
func (c *KeyperConfig) Address() common.Address {
	return crypto.PubkeyToAddress(c.SigningKey.PublicKey)
}
