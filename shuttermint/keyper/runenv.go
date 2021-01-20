package keyper

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

type RunEnv struct {
	MessageSender       MessageSender
	ContractCaller      *ContractCaller
	WatchedTransactions chan *types.Transaction
}

func (e RunEnv) SendMessage(ctx context.Context, msg *shmsg.Message) error {
	return e.MessageSender.SendMessage(ctx, msg)
}

func (e RunEnv) GetContractCaller(ctx context.Context) *ContractCaller {
	return e.ContractCaller
}

func (e RunEnv) WatchTransaction(tx *types.Transaction) {
	e.WatchedTransactions <- tx
}
