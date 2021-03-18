package fx

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
	"github.com/brainbot-com/shutter/shuttermint/medley"
)

const (
	numMainChainWorkers = 20
)

type InFlightMainChainTX struct {
	MainChainTX MainChainTX
	Transaction *types.Transaction
}

type RunEnv struct {
	MessageSender        MessageSender
	ContractCaller       *contract.Caller
	mainChainTXs         chan MainChainTX
	inFlightMainChainTXs chan InFlightMainChainTX
	currentWorld         func() observe.World
}

func NewRunEnv(messageSender MessageSender, contractCaller *contract.Caller, currentWorld func() observe.World) *RunEnv {
	return &RunEnv{
		MessageSender:        messageSender,
		ContractCaller:       contractCaller,
		mainChainTXs:         make(chan MainChainTX, numMainChainWorkers),
		inFlightMainChainTXs: make(chan InFlightMainChainTX),
		currentWorld:         currentWorld,
	}
}

func (runenv *RunEnv) sendShuttermintMessage(ctx context.Context, act *SendShuttermintMessage) error {
	log.Printf("=====%s", act)
	err := runenv.MessageSender.SendMessage(ctx, act.Msg)
	return err
}

func (runenv *RunEnv) sendMainChainTX(ctx context.Context, act MainChainTX) error {
	var err error
	var tx *types.Transaction
	var auth *bind.TransactOpts

	auth, err = runenv.ContractCaller.Auth()

	if err != nil {
		return err
	}
	tx, err = act.SendTX(runenv.ContractCaller, auth)
	if err != nil {
		return err
	}

	runenv.inFlightMainChainTXs <- InFlightMainChainTX{MainChainTX: act, Transaction: tx}
	return nil
}

func (runenv *RunEnv) waitMined(ctx context.Context, inFlightMainChainTX InFlightMainChainTX) {
	tx := inFlightMainChainTX.Transaction
	act := inFlightMainChainTX.MainChainTX
	receipt, err := medley.WaitMined(ctx, runenv.ContractCaller.Ethclient, tx.Hash())
	if err != nil {
		log.Printf("Error waiting for transaction %s: %v", tx.Hash().Hex(), err)
		return
	}
	if receipt == nil {
		// This happens if the context is canceled.
		return
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		world := runenv.CurrentWorld() // XXX we should make sure our world includes the receipt's blocknumber
		expired := act.IsExpired(world)
		log.Printf("TX reverted: expired=%t, %s, hash=%s", expired, act, tx.Hash().Hex())
	} else {
		log.Printf("TX success: %s, hash=%s", act, tx.Hash().Hex())
	}
}

func (runenv *RunEnv) RunActions(ctx context.Context, actionCounter uint64, actions []IAction) {
	_ = actionCounter
	var err error
	if len(actions) > 0 {
		log.Printf("Running %d actions", len(actions))
	}

	for _, act := range actions {
		switch a := act.(type) {
		case *SendShuttermintMessage:
			err = runenv.sendShuttermintMessage(ctx, a)
		case MainChainTX:
			runenv.mainChainTXs <- a
		default:
			err = errors.Errorf("cannot run %s", a)
		}

		// XXX at the moment we just let the whole program die. We need a better strategy
		// here. We could retry the actions or feed the errors back into our state
		if err != nil {
			panic(err)
		}
	}
}

func (runenv *RunEnv) handleMainChainTXs(ctx context.Context) {
	for {
		select {
		case a := <-runenv.mainChainTXs:
			var err error
			for {
				if a.IsExpired(runenv.CurrentWorld()) {
					log.Printf("action expired: %s", a)
					break
				}
				if err != nil {
					log.Printf("retrying main chain tx %s; err=%s", a, err)
				}
				err := runenv.sendMainChainTX(ctx, a)
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
		case t := <-runenv.inFlightMainChainTXs:
			runenv.waitMined(ctx, t)
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

	for i := 0; i < numMainChainWorkers; i++ {
		g.Go(func() error {
			runenv.handleInFlightTXs(ctx)
			return nil
		})
	}
}
