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
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	"golang.org/x/sync/errgroup"

	"github.com/shutter-network/shutter/shuttermint/contract"
	"github.com/shutter-network/shutter/shuttermint/keyper/fx"
	"github.com/shutter-network/shutter/shuttermint/keyper/observe"
)

// IsWebsocketURL returns true iff the given URL is a websocket URL, i.e. if it starts with ws://
// or wss://. This is needed for the watchMainChainHeadBlock method.
func IsWebsocketURL(url string) bool {
	return strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://")
}

type Keyper struct {
	// Keep the atomic.Value as first field in order to make sure it's 64-bit aligned. Visit
	// https://golang.org/pkg/sync/atomic/#pkg-note-BUG for more information
	world atomic.Value // holds an observe.World struct

	Config Config // Configuration of the keyper client read from the config file
	State  *State // keyper's internal state

	ContractCaller contract.Caller
	shmcl          client.Client
	MessageSender  fx.MessageSender
	lastlogTime    time.Time
	runenv         *fx.RunEnv

	mainChainCh     chan *observe.MainChain    // observed main chain updates
	shutterCh       chan *observe.Shutter      // observed shutter updates
	signalCh        chan os.Signal             // signals received
	shutterFilterCh chan observe.ShutterFilter // new shutter filter for garbage collecting the shutter state
}

func NewKeyper(kc Config) Keyper {
	world := atomic.Value{}
	world.Store(observe.World{
		Shutter:   observe.NewShutter(),
		MainChain: observe.NewMainChain(kc.MainChainFollowDistance),
	})

	return Keyper{
		Config: kc,
		State:  NewState(),
		world:  world,
	}
}

func NewContractCallerFromConfig(config Config) (contract.Caller, error) {
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

	return contract.NewCaller(
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
	if err != nil {
		return err
	}
	kpr.runenv = fx.NewRunEnv(kpr.MessageSender, &kpr.ContractCaller, kpr.CurrentWorld, kpr.pathActionsGob())
	kpr.mainChainCh = make(chan *observe.MainChain)
	kpr.shutterCh = make(chan *observe.Shutter)
	kpr.signalCh = make(chan os.Signal, 1)
	kpr.shutterFilterCh = make(chan observe.ShutterFilter, 3)
	signal.Notify(kpr.signalCh, syscall.SIGUSR1)

	return nil
}

func (kpr *Keyper) dkginfo() string {
	var ds []string
	for i := len(kpr.State.DKGs) - 1; i >= 0; i-- {
		dkg := kpr.State.DKGs[i]
		if dkg.IsFinalized() {
			break
		}
		ds = append(ds, dkg.ShortInfo())
	}

	if len(ds) == 0 {
		return ""
	}
	return fmt.Sprintf(", DKGs: %s", strings.Join(ds, " - "))
}

func (kpr *Keyper) ShortInfo() string {
	world := kpr.CurrentWorld()
	var notAKeyper string
	if len(world.MainChain.BatchConfigs) > 0 {
		configIndex := world.MainChain.ActiveConfigIndex(world.MainChain.CurrentBlock)
		batchConfig := world.MainChain.BatchConfigs[configIndex]
		if !batchConfig.IsKeyper(kpr.Config.Address()) {
			notAKeyper = fmt.Sprintf("Not configured as keyper in config %d, ", configIndex)
		}
	}
	return fmt.Sprintf(
		"%sshutter block %d, main chain %d, %s, last eon started %d, num half steps: %d%s",
		notAKeyper,
		world.Shutter.CurrentBlock,
		world.MainChain.CurrentBlock,
		kpr.runenv.ShortInfo(),
		kpr.State.LastEonStarted,
		world.MainChain.NumExecutionHalfSteps,
		kpr.dkginfo(),
	)
}

func (kpr *Keyper) dumpInternalState() {
	log.Printf("Received signal. Dumping internal state")
	world := kpr.CurrentWorld()
	pretty.Println("Shutter:", world.Shutter)
	pretty.Println("Mainchain:", world.MainChain)
	pretty.Println("State:", kpr.State)
}

// syncOnce syncs the main and shutter chain at least once. Otherwise, the state of one of the two
// will be much more recent than the other one when the first block appears.
func (kpr *Keyper) syncOnce(ctx context.Context) {
	var world observe.World
	for world.MainChain == nil || world.Shutter == nil {
		select {
		case <-kpr.signalCh:
			kpr.dumpInternalState()
			continue
		case mainChain := <-kpr.mainChainCh:
			world.MainChain = mainChain
		case shutter := <-kpr.shutterCh:
			world.Shutter = shutter
		case <-ctx.Done():
			return
		}
	}
	kpr.world.Store(world)
}

func (kpr *Keyper) syncLoop(ctx context.Context) error {
	var world observe.World = kpr.CurrentWorld()

	for {
		select {
		case <-kpr.signalCh:
			kpr.dumpInternalState()
			continue
		case <-ctx.Done():
			return ctx.Err()
		case mainChain := <-kpr.mainChainCh:
			world.MainChain = mainChain
		case shutter := <-kpr.shutterCh:
			world.Shutter = shutter
		}
		kpr.world.Store(world)
		err := kpr.runOneStep(ctx)
		if err != nil {
			return err
		}
		select {
		case kpr.shutterFilterCh <- kpr.State.GetShutterFilter(world.MainChain):
		default:
		}
	}
}

func (kpr *Keyper) startSyncTasks(ctx context.Context, g *errgroup.Group) {
	g.Go(func() error {
		return observe.SyncMain(ctx, &kpr.ContractCaller, kpr.CurrentWorld().MainChain, kpr.mainChainCh)
	})
	g.Go(func() error {
		return observe.SyncShutter(ctx, kpr.shmcl, kpr.CurrentWorld().Shutter, kpr.shutterCh, kpr.shutterFilterCh)
	})
}

func (kpr *Keyper) loadRunenv(ctx context.Context) error {
	havePendingActions, err := kpr.runenv.Load(ctx)
	if err != nil {
		return err
	}

	if !havePendingActions && len(kpr.State.Actions) == 0 && kpr.State.PendingHalfStep != nil {
		log.Printf("Fixing State: PendingHalfStep should be nil!")
		kpr.State.PendingHalfStep = nil
	}
	return nil
}

func (kpr *Keyper) run(ctx context.Context, g *errgroup.Group) error {
	kpr.startSyncTasks(ctx, g)
	kpr.syncOnce(ctx)
	kpr.runenv.StartBackgroundTasks(ctx, g)
	if err := kpr.loadRunenv(ctx); err != nil {
		return err
	}

	if len(kpr.State.Actions) > 0 {
		err := kpr.runActions(ctx)
		if err != nil {
			return err
		}
	}
	return kpr.syncLoop(ctx)
}

func (kpr *Keyper) Run(ctx context.Context) error {
	if err := kpr.init(); err != nil {
		return err
	}
	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return kpr.run(groupCtx, g)
	})

	return g.Wait()
}

func (kpr *Keyper) CurrentWorld() observe.World {
	return kpr.world.Load().(observe.World)
}

type storedState struct {
	State     *State
	Shutter   *observe.Shutter
	MainChain *observe.MainChain
}

func (kpr *Keyper) pathStateGob() string {
	return filepath.Join(kpr.Config.DBDir, "state.gob")
}

func (kpr *Keyper) pathActionsGob() string {
	return filepath.Join(kpr.Config.DBDir, "actions.gob")
}

func (kpr *Keyper) LoadState() error {
	gobpath := kpr.pathStateGob()

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

	if st.State.SyncHeight == 0 && st.Shutter.CurrentBlock > 0 {
		log.Printf("Fixing SyncHeight: %d", st.Shutter.CurrentBlock)
		st.State.SyncHeight = st.Shutter.CurrentBlock // We didn't have this field in older versions
	}
	kpr.State = st.State
	world := observe.World{
		Shutter:   st.Shutter,
		MainChain: st.MainChain,
	}
	kpr.world.Store(world)

	return nil
}

func (kpr *Keyper) saveState() error {
	gobpath := kpr.pathStateGob()
	tmppath := gobpath + ".tmp"
	file, err := os.Create(tmppath)
	if err != nil {
		return err
	}
	defer file.Close()
	world := kpr.CurrentWorld()
	st := storedState{
		State:     kpr.State,
		Shutter:   world.Shutter,
		MainChain: world.MainChain,
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

func (kpr *Keyper) runActions(ctx context.Context) error {
	now := time.Now()
	if len(kpr.State.Actions) > 0 || now.Sub(kpr.lastlogTime) > 10*time.Second {
		log.Println(kpr.ShortInfo())
		kpr.lastlogTime = now
	}

	err := kpr.runenv.RunActions(ctx, kpr.State.ActionCounter, kpr.State.Actions)
	if err != nil {
		return err
	}
	kpr.State.ActionCounter += uint64(len(kpr.State.Actions))
	kpr.State.Actions = nil
	return nil
}

func (kpr *Keyper) decide() []fx.IAction {
	decider := NewDecider(kpr)
	decider.Decide()
	return decider.Actions
}

func (kpr *Keyper) runOneStep(ctx context.Context) error {
	if len(kpr.State.Actions) > 0 {
		panic("internal errror: kpr.State.Actions is not empty")
	}
	kpr.State.Actions = kpr.decide()
	if err := kpr.saveState(); err != nil {
		panic(err)
	}
	return kpr.runActions(ctx)
}
