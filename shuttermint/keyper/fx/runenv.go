package fx

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/medley"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

const (
	// watchedTransactionsBufferSize is the maximum number of txs to watch. If this is too small,
	// actions will stall.
	watchedTransactionsBufferSize = 100
)

type RunEnv struct {
	MessageSender       MessageSender
	ContractCaller      *contract.Caller
	WatchedTransactions chan *types.Transaction
}

func NewRunEnv(messageSender MessageSender, contractCaller *contract.Caller) *RunEnv {
	return &RunEnv{
		MessageSender:       messageSender,
		ContractCaller:      contractCaller,
		WatchedTransactions: make(chan *types.Transaction, watchedTransactionsBufferSize),
	}
}

func (runenv RunEnv) SendMessage(ctx context.Context, msg *shmsg.Message) error {
	return runenv.MessageSender.SendMessage(ctx, msg)
}

func (runenv RunEnv) GetContractCaller(ctx context.Context) *contract.Caller {
	return runenv.ContractCaller
}

func (runenv RunEnv) WatchTransaction(tx *types.Transaction) {
	runenv.WatchedTransactions <- tx
}

func (runenv *RunEnv) sendShuttermintMessage(ctx context.Context, act *SendShuttermintMessage) error {
	log.Printf("=====%s", act)
	err := runenv.MessageSender.SendMessage(ctx, act.Msg)
	return err
}

func (runenv *RunEnv) sendMainChainTX(ctx context.Context, act MainChainTX) error {
	cc := runenv.GetContractCaller(ctx)
	auth, err := cc.Auth()
	if err != nil {
		return err
	}
	tx, err := act.SendTX(cc, auth)
	if err != nil {
		// XXX consider handling the error somehow
		log.Printf("Error: %s: %s", act, err)
		return nil
	}
	runenv.WatchTransaction(tx)
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
			err = runenv.sendMainChainTX(ctx, a)
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

func (runenv *RunEnv) StartBackgroundTasks(ctx context.Context, g *errgroup.Group) {
	g.Go(func() error { return runenv.watchTransactions(ctx) })
}

func (runenv *RunEnv) watchTransactions(ctx context.Context) error {
	for {
		select {
		case tx := <-runenv.WatchedTransactions:
			receipt, err := medley.WaitMined(ctx, runenv.ContractCaller.Ethclient, tx.Hash())
			if err != nil {
				log.Printf("Error waiting for transaction %s: %v", tx.Hash().Hex(), err)
			}
			if receipt == nil {
				// This happens if the context is canceled. By continuing, we end up in the
				// ctx.Done case
				continue
			}
			if receipt.Status != types.ReceiptStatusSuccessful {
				log.Printf("Tx %s has failed and was reverted", tx.Hash().Hex())
			} else {
				log.Printf("Tx %s was successful", tx.Hash().Hex())
			}
		case <-ctx.Done():
			return nil
		}
	}
}
