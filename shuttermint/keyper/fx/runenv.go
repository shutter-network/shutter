package fx

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/medley"
)

const (
	numMainChainWorkers = 20
)

type RunEnv struct {
	MessageSender  MessageSender
	ContractCaller *contract.Caller
	mainChainTXs   chan MainChainTX
	mutex          sync.Mutex
}

func NewRunEnv(messageSender MessageSender, contractCaller *contract.Caller) *RunEnv {
	return &RunEnv{
		MessageSender:  messageSender,
		ContractCaller: contractCaller,
		mainChainTXs:   make(chan MainChainTX),
	}
}

func (runenv RunEnv) GetContractCaller(ctx context.Context) *contract.Caller {
	return runenv.ContractCaller
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

	func() {
		// We lock the mutex in order to not reuse the same nonce
		runenv.mutex.Lock()
		defer runenv.mutex.Unlock()
		auth, err = runenv.ContractCaller.Auth()
		if err != nil {
			return
		}
		tx, err = act.SendTX(runenv.ContractCaller, auth)
	}()

	if err != nil {
		// XXX consider handling the error somehow
		log.Printf("Error: %s: %s", act, err)
		return nil
	}

	receipt, err := medley.WaitMined(ctx, runenv.ContractCaller.Ethclient, tx.Hash())
	if err != nil {
		log.Printf("Error waiting for transaction %s: %v", tx.Hash().Hex(), err)
	}
	if receipt == nil {
		// This happens if the context is canceled.
		return nil
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Printf("TX reverted: %s, hash=%s", act, tx.Hash().Hex())
	} else {
		log.Printf("TX success: %s, hash=%s", act, tx.Hash().Hex())
	}

	return nil
}

func (runenv *RunEnv) RunActions(ctx context.Context, actions []IAction) {
	var err error
	if len(actions) > 0 {
		log.Printf("Running %d actions", len(actions))
	}

	for _, act := range actions {
		switch a := act.(type) {
		case SendShuttermintMessage:
			err = runenv.sendShuttermintMessage(ctx, &a)
		case MainChainTX:
			runenv.mainChainTXs <- a
		default:
			err = fmt.Errorf("cannot run %s", a)
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
			_ = runenv.sendMainChainTX(ctx, a)
		case <-ctx.Done():
			return
		}
	}
}

func (runenv *RunEnv) StartBackgroundTasks(ctx context.Context, g *errgroup.Group) {
	for i := 0; i < numMainChainWorkers; i++ {
		g.Go(func() error {
			runenv.handleMainChainTXs(ctx)
			return nil
		})
	}
}
