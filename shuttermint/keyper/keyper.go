// Package keyper contains the keyper implementation
package keyper

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/fx"
	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
)

const (
	// mainChainTimeout is the time after which we assume the connection to the main chain
	// node is lost if no new block is received
	mainChainTimeout           = 30 * time.Second
	mainChainReconnectInterval = 5 * time.Second // time between two reconnection attempts
	shuttermintTimeout         = 10 * time.Second
	shutterReconnectInterval   = 5 * time.Second
)

// IsWebsocketURL returns true iff the given URL is a websocket URL, i.e. if it starts with ws://
// or wss://. This is needed for the watchMainChainHeadBlock method.
func IsWebsocketURL(url string) bool {
	return strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://")
}

// Address returns the keyper's Ethereum address.
func (c *KeyperConfig) Address() common.Address {
	return crypto.PubkeyToAddress(c.SigningKey.PublicKey)
}

type Keyper struct {
	Config    KeyperConfig
	State     *State
	Shutter   *observe.Shutter
	MainChain *observe.MainChain

	ContractCaller contract.Caller
	shmcl          client.Client
	MessageSender  fx.MessageSender
	lastlogTime    time.Time
	runenv         *fx.RunEnv
}

func NewKeyper(kc KeyperConfig) Keyper {
	return Keyper{
		Config:    kc,
		State:     NewState(),
		Shutter:   observe.NewShutter(),
		MainChain: observe.NewMainChain(kc.MainChainFollowDistance),
	}
}

func NewContractCallerFromConfig(config KeyperConfig) (contract.Caller, error) {
	ethcl, err := ethclient.Dial(config.EthereumURL)
	if err != nil {
		return contract.Caller{}, err
	}
	configContract, err := contract.NewConfigContract(config.ConfigContractAddress, ethcl)
	if err != nil {
		return contract.Caller{}, err
	}

	keyBroadcastContract, err := contract.NewKeyBroadcastContract(config.KeyBroadcastContractAddress, ethcl)
	if err != nil {
		return contract.Caller{}, err
	}

	batcherContract, err := contract.NewBatcherContract(config.BatcherContractAddress, ethcl)
	if err != nil {
		return contract.Caller{}, err
	}

	executorContract, err := contract.NewExecutorContract(config.ExecutorContractAddress, ethcl)
	if err != nil {
		return contract.Caller{}, err
	}

	depositContract, err := contract.NewDepositContract(config.DepositContractAddress, ethcl)
	if err != nil {
		return contract.Caller{}, err
	}

	keyperSlasher, err := contract.NewKeyperSlasher(config.KeyperSlasherAddress, ethcl)
	if err != nil {
		return contract.Caller{}, err
	}

	return contract.NewContractCaller(
		ethcl,
		config.SigningKey,
		configContract,
		keyBroadcastContract,
		batcherContract,
		executorContract,
		depositContract,
		keyperSlasher,
	), nil
}

func (kpr *Keyper) init() error {
	if kpr.shmcl != nil {
		panic("internal error: already initialized")
	}
	var err error
	kpr.shmcl, err = http.New(kpr.Config.ShuttermintURL, "/websocket")
	if err != nil {
		return errors.Wrapf(err, "create shuttermint client at %s", kpr.Config.ShuttermintURL)
	}
	err = kpr.shmcl.Start()
	if err != nil {
		return errors.Wrapf(err, "start shuttermint client")
	}
	ms := fx.NewRPCMessageSender(kpr.shmcl, kpr.Config.SigningKey)
	kpr.MessageSender = &ms

	kpr.ContractCaller, err = NewContractCallerFromConfig(kpr.Config)
	kpr.runenv = fx.NewRunEnv(kpr.MessageSender, &kpr.ContractCaller, kpr.CurrentWorld)
	return err
}

func (kpr *Keyper) syncMain(ctx context.Context, mainChains chan<- *observe.MainChain, syncErrors chan<- error) error {
	headers := make(chan *types.Header)
	sub, err := kpr.ContractCaller.Ethclient.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}

	reconnect := func() {
		sub.Unsubscribe()
		for {
			log.Println("Attempting reconnection to main chain")
			sub, err = kpr.ContractCaller.Ethclient.SubscribeNewHead(ctx, headers)
			if err != nil {
				select {
				case <-time.After(mainChainReconnectInterval):
					continue
				case <-ctx.Done():
					return
				}
			} else {
				log.Println("Main chain connection regained")
				return
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			return nil
		case <-headers:
			newMainChain, err := kpr.MainChain.SyncToHead(ctx, &kpr.ContractCaller)
			if err != nil {
				syncErrors <- err
			} else {
				mainChains <- newMainChain
			}
		case err := <-sub.Err():
			log.Println("Main chain connection lost:", err)
			reconnect()
		case <-time.After(mainChainTimeout):
			log.Println("No main chain blocks received in a long time")
			reconnect()
		}
	}
}

func (kpr *Keyper) syncShutter(ctx context.Context, shutters chan<- *observe.Shutter, syncErrors chan<- error) error {
	name := "keyper"
	query := "tm.event = 'NewBlock'"
	events, err := kpr.shmcl.Subscribe(ctx, name, query)
	if err != nil {
		return err
	}

	reconnect := func() {
		for {
			log.Println("Attempting reconnection to Shuttermint")

			ctx2, cancel2 := context.WithTimeout(ctx, shutterReconnectInterval)
			events, err = kpr.shmcl.Subscribe(ctx2, name, query)
			cancel2()

			if err != nil {
				// try again, unless context is canceled
				select {
				case <-ctx.Done():
					return
				default:
					continue
				}
			} else {
				log.Println("Shuttermint connection regained")
				return
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-events:
			newShutter, err := kpr.Shutter.SyncToHead(ctx, kpr.shmcl)
			if err != nil {
				syncErrors <- err
			} else {
				shutters <- newShutter
			}
		case <-time.After(shuttermintTimeout):
			log.Println("No Shuttermint blocks received in a long time")
			reconnect()
		}
	}
}

func (kpr *Keyper) ShortInfo() string {
	var dkgInfo []string
	for _, dkg := range kpr.State.DKGs {
		dkgInfo = append(dkgInfo, dkg.ShortInfo())
	}
	return fmt.Sprintf(
		"shutter block %d, main chain %d, last eon started %d, num half steps: %d, DKGs: %s",
		kpr.Shutter.CurrentBlock,
		kpr.MainChain.CurrentBlock,
		kpr.State.LastEonStarted,
		kpr.MainChain.NumExecutionHalfSteps,
		strings.Join(dkgInfo, " - "),
	)
}

func (kpr *Keyper) syncOnce(ctx context.Context) error {
	newMain, err := kpr.MainChain.SyncToHead(ctx, &kpr.ContractCaller)
	if err != nil {
		return err
	}
	kpr.MainChain = newMain

	newShutter, err := kpr.Shutter.SyncToHead(ctx, kpr.shmcl)
	if err != nil {
		return err
	}
	kpr.Shutter = newShutter

	return nil
}

func (kpr *Keyper) Run() error {
	err := kpr.init()
	if err != nil {
		return err
	}
	g, ctx := errgroup.WithContext(context.Background())

	// Sync main and shutter chain once. Otherwise, the state of one of the two will be much more
	// recent than the other one when the first block appears.
	err = kpr.syncOnce(ctx)
	if err != nil {
		return err
	}

	mainChains := make(chan *observe.MainChain)
	shutters := make(chan *observe.Shutter)
	syncErrors := make(chan error)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1)
	g.Go(func() error { return kpr.syncMain(ctx, mainChains, syncErrors) })
	g.Go(func() error { return kpr.syncShutter(ctx, shutters, syncErrors) })
	kpr.runenv.StartBackgroundTasks(ctx, g)
	if len(kpr.State.Actions) > 0 {
		kpr.runActions(ctx)
	}

	for {
		select {
		case sig := <-signals:
			log.Printf("Received %s. Dumping internal state", sig)
			pretty.Println("Shutter:", kpr.Shutter)
			pretty.Println("Mainchain:", kpr.MainChain)
			pretty.Println("State:", kpr.State)
		case mainChain := <-mainChains:
			kpr.MainChain = mainChain
			kpr.runOneStep(ctx)
		case shutter := <-shutters:
			kpr.Shutter = shutter
			kpr.runOneStep(ctx)
		case err := <-syncErrors:
			return err
		case <-ctx.Done():
			return g.Wait()
		}
	}
}

func (kpr *Keyper) CurrentWorld() observe.World {
	return observe.World{Shutter: kpr.Shutter, MainChain: kpr.MainChain}
}

type storedState struct {
	State     *State
	Shutter   *observe.Shutter
	MainChain *observe.MainChain
	Actions   []fx.IAction
}

func (kpr *Keyper) gobpath() string {
	return filepath.Join(kpr.Config.DBDir, "state.gob")
}

func (kpr *Keyper) LoadState() error {
	gobpath := kpr.gobpath()

	gobfile, err := os.Open(gobpath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	log.Printf("Loading state from %s", gobpath)

	defer gobfile.Close()
	dec := gob.NewDecoder(gobfile)
	st := storedState{}
	err = dec.Decode(&st)
	if err != nil {
		return err
	}
	kpr.State = st.State
	kpr.Shutter = st.Shutter
	kpr.MainChain = st.MainChain

	// We don't store the pending actions, so for the moment it's better to reset the
	// PendingHalfStep. XXX
	kpr.State.PendingHalfStep = nil
	return nil
}

func (kpr *Keyper) saveState() error {
	gobpath := kpr.gobpath()
	tmppath := gobpath + ".tmp"
	file, err := os.Create(tmppath)
	if err != nil {
		return err
	}
	defer file.Close()
	st := storedState{
		State:     kpr.State,
		Shutter:   kpr.Shutter,
		MainChain: kpr.MainChain,
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(st)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}
	err = os.Rename(tmppath, gobpath)
	return err
}

func (kpr *Keyper) runActions(ctx context.Context) {
	now := time.Now()
	if len(kpr.State.Actions) > 0 || now.Sub(kpr.lastlogTime) > 10*time.Second {
		log.Println(kpr.ShortInfo())
		kpr.lastlogTime = now
	}

	kpr.runenv.RunActions(ctx, kpr.State.ActionCounter, kpr.State.Actions)
	kpr.State.ActionCounter += uint64(len(kpr.State.Actions))
	kpr.State.Actions = nil
}

func (kpr *Keyper) decide() []fx.IAction {
	decider := NewDecider(kpr)
	decider.Decide()
	return decider.Actions
}

func (kpr *Keyper) runOneStep(ctx context.Context) {
	if len(kpr.State.Actions) > 0 {
		panic("internal errror: kpr.State.Actions is not empty")
	}
	kpr.State.Actions = kpr.decide()
	if err := kpr.saveState(); err != nil {
		panic(err)
	}
	kpr.runActions(ctx)
}
