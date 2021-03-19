package fx

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
	"github.com/brainbot-com/shutter/shuttermint/medley"
)

const (
	numMainChainWorkers = 20
)

type ActionWithID struct {
	Action IAction
	ID     ActionID
}

type RunEnv struct {
	PendingActions       *PendingActions
	PendingActionsPath   string
	MessageSender        MessageSender
	ContractCaller       *contract.Caller
	shuttermintMessages  chan ActionID
	mainChainTXs         chan ActionID
	inFlightMainChainTXs chan ActionID
	currentWorld         func() observe.World
}

func NewRunEnv(messageSender MessageSender, contractCaller *contract.Caller, currentWorld func() observe.World, path string) *RunEnv {
	return &RunEnv{
		PendingActions:       NewPendingActions(path),
		MessageSender:        messageSender,
		ContractCaller:       contractCaller,
		shuttermintMessages:  make(chan ActionID),
		mainChainTXs:         make(chan ActionID, numMainChainWorkers),
		inFlightMainChainTXs: make(chan ActionID),
		currentWorld:         currentWorld,
	}
}

func (runenv *RunEnv) sendShuttermintMessage(ctx context.Context, act *SendShuttermintMessage) error {
	log.Printf("=====%s", act)
	err := runenv.MessageSender.SendMessage(ctx, act.Msg)
	return err
}

func (runenv *RunEnv) sendMainChainTX(ctx context.Context, id ActionID, act MainChainTX) error {
	var err error
	var tx *types.Transaction
	var auth *bind.TransactOpts

	auth, err = runenv.ContractCaller.Auth()
	auth.Context = ctx

	if err != nil {
		return err
	}
	tx, err = act.SendTX(runenv.ContractCaller, auth)
	if err != nil {
		return err
	}
	runenv.PendingActions.SetMainChainTXHash(id, tx.Hash())
	runenv.inFlightMainChainTXs <- id
	return nil
}

func (runenv *RunEnv) waitMined(ctx context.Context, id ActionID) {
	act := runenv.PendingActions.GetAction(id)
	hash := runenv.PendingActions.GetMainChainTXHash(id)
	receipt, err := medley.WaitMined(ctx, runenv.ContractCaller.Ethclient, hash)
	if err != nil {
		log.Printf("Error waiting for transaction %s: %v", hash.Hex(), err)
		return
	}
	if receipt == nil {
		// This happens if the context is canceled.
		return
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		world := runenv.CurrentWorld() // XXX we should make sure our world includes the receipt's blocknumber
		expired := act.IsExpired(world)
		log.Printf("TX reverted: expired=%t, %s, hash=%s", expired, act, hash.Hex())
	} else {
		log.Printf("TX success: %s, hash=%s", act, hash.Hex())
	}
}

func (runenv *RunEnv) RunActions(ctx context.Context, actionCounter uint64, actions []IAction) {
	if len(actions) == 0 {
		return
	}

	log.Printf("Running %d actions", len(actions))
	startID, endID := runenv.PendingActions.AddActions(ActionID(actionCounter), actions)
	for id := startID; id < endID; id++ {
		runenv.scheduleAction(id)
	}
}

// scheduleAction schedules an action to be run. The given action must already be stored in the
// pending actions struct.
func (runenv *RunEnv) scheduleAction(id ActionID) {
	act := runenv.PendingActions.GetAction(id)
	switch a := act.(type) {
	case *SendShuttermintMessage:
		runenv.shuttermintMessages <- id
	case MainChainTX:
		nullhash := common.Hash{}
		txhash := runenv.PendingActions.GetMainChainTXHash(id)
		if txhash != nullhash {
			runenv.inFlightMainChainTXs <- id
		} else {
			runenv.mainChainTXs <- id
		}
	default:
		log.Fatalf("cannot run %s", a)
	}
}

// Load loads the pending actions from disk and schedules the actions to be run.
func (runenv *RunEnv) Load() error {
	err := runenv.PendingActions.Load()
	if err != nil {
		return err
	}

	for _, id := range runenv.PendingActions.SortedIDs() {
		runenv.scheduleAction(id)
	}
	return nil
}

func (runenv *RunEnv) handleShuttermintMessages(ctx context.Context) {
	for {
		select {
		case id := <-runenv.shuttermintMessages:
			a := runenv.PendingActions.GetAction(id).(*SendShuttermintMessage)
			var err error
			for {
				if a.IsExpired(runenv.CurrentWorld()) {
					log.Printf("action expired: %s", a)
					break
				}
				if err != nil {
					log.Printf("retrying main chain tx %s; err=%s", a, err)
				}
				err = runenv.sendShuttermintMessage(ctx, a)
				if err == nil {
					break
				}
				time.Sleep(time.Second)
			}
			runenv.PendingActions.RemoveAction(id)
		case <-ctx.Done():
			return
		}
	}
}

func (runenv *RunEnv) handleMainChainTXs(ctx context.Context) {
	for {
		select {
		case id := <-runenv.mainChainTXs:
			a := runenv.PendingActions.GetAction(id).(MainChainTX)
			var err error
			for {
				if a.IsExpired(runenv.CurrentWorld()) {
					log.Printf("action expired: %s", a)
					runenv.PendingActions.RemoveAction(id)
					break
				}
				if err != nil {
					log.Printf("retrying main chain tx %s; err=%s", a, err)
				}
				err := runenv.sendMainChainTX(ctx, id, a)
				if err == nil {
					break
				}
				time.Sleep(time.Second)
			}

		case <-ctx.Done():
			return
		}
	}
}

func (runenv *RunEnv) handleInFlightTXs(ctx context.Context) {
	for {
		select {
		case id := <-runenv.inFlightMainChainTXs:
			runenv.waitMined(ctx, id)
			runenv.PendingActions.RemoveAction(id)
		case <-ctx.Done():
			return
		}
	}
}

func (runenv *RunEnv) CurrentWorld() observe.World {
	return runenv.currentWorld()
}

func (runenv *RunEnv) StartBackgroundTasks(ctx context.Context, g *errgroup.Group) {
	// Start exactly one worker, because we like to make sure the actions are started in the
	// order given to us
	g.Go(func() error {
		runenv.handleMainChainTXs(ctx)
		return nil
	})

	g.Go(func() error {
		runenv.handleShuttermintMessages(ctx)
		return nil
	})
	for i := 0; i < numMainChainWorkers; i++ {
		g.Go(func() error {
			runenv.handleInFlightTXs(ctx)
			return nil
		})
	}
}
