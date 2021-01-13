package keyper

import (
	"context"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

type RunEnv struct {
	MessageSender  MessageSender
	ContractCaller *ContractCaller
}

func (e RunEnv) SendMessage(ctx context.Context, msg *shmsg.Message) error {
	return e.MessageSender.SendMessage(ctx, msg)
}

func (e RunEnv) GetContractCaller(ctx context.Context) *ContractCaller {
	return e.ContractCaller
}
