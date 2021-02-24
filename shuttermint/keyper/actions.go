package keyper

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/shcrypto"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// IRunEnv is passed as a parameter to IAction's Run function.
type IRunEnv interface {
	MessageSender
	GetContractCaller(ctx context.Context) *contract.Caller
	WatchTransaction(tx *types.Transaction)
}

// IAction describes an action to run as determined by the Decider's Decide method.
type IAction interface {
	Run(ctx context.Context, runenv IRunEnv) error
}

var (
	_ IAction = SendShuttermintMessage{}
	_ IAction = ExecuteCipherBatch{}
	_ IAction = ExecutePlainBatch{}
	_ IAction = SkipCipherBatch{}
	_ IAction = Accuse{}
	_ IAction = Appeal{}
	_ IAction = EonKeyBroadcast{}
)

// SendShuttermintMessage is an Action that sends a message to shuttermint
type SendShuttermintMessage struct {
	description string
	msg         *shmsg.Message
}

func (a SendShuttermintMessage) Run(ctx context.Context, runenv IRunEnv) error {
	log.Printf("=====%s", a)
	return runenv.SendMessage(ctx, a.msg)
}

func (a SendShuttermintMessage) String() string {
	return fmt.Sprintf("=> shuttermint: %s", a.description)
}

// ExecuteCipherBatch is an Action that instructs the executor contract to execute a cipher batch.
type ExecuteCipherBatch struct {
	batchIndex      uint64
	cipherBatchHash [32]byte
	transactions    [][]byte
	keyperIndex     uint64
}

func (a ExecuteCipherBatch) Run(ctx context.Context, runenv IRunEnv) error {
	log.Printf("=====%s", a)

	cc := runenv.GetContractCaller(ctx)
	auth, err := cc.Auth()
	if err != nil {
		return err
	}

	tx, err := cc.ExecutorContract.ExecuteCipherBatch(auth, a.batchIndex, a.cipherBatchHash, a.transactions, a.keyperIndex)
	if err != nil {
		// XXX consider handling the error somehow
		log.Printf("Error creating cipher batch execution tx: %s", err)
		return nil
	}
	runenv.WatchTransaction(tx)

	return nil
}

func (a ExecuteCipherBatch) String() string {
	return fmt.Sprintf("=> executor contract: execute cipher batch %d", a.batchIndex)
}

// ExecutePlainBatch is an Action that instructs the executor contract to execute a plain batch.
type ExecutePlainBatch struct {
	batchIndex   uint64
	transactions [][]byte
}

func (a ExecutePlainBatch) Run(ctx context.Context, runenv IRunEnv) error {
	log.Printf("=====%s", a)

	cc := runenv.GetContractCaller(ctx)
	auth, err := cc.Auth()
	if err != nil {
		return err
	}

	tx, err := cc.ExecutorContract.ExecutePlainBatch(auth, a.batchIndex, a.transactions)
	if err != nil {
		// XXX consider handling the error somehow
		log.Printf("Error creating plain batch execution tx: %s", err)
		return nil
	}
	runenv.WatchTransaction(tx)

	return nil
}

func (a ExecutePlainBatch) String() string {
	return fmt.Sprintf("=> executor contract: execute plain batch %d", a.batchIndex)
}

// SkipCipherBatch is an Action that instructs the executor contract to skip a cipher batch
type SkipCipherBatch struct {
	batchIndex uint64
}

func (a SkipCipherBatch) Run(ctx context.Context, runenv IRunEnv) error {
	log.Printf("=====%s", a)

	cc := runenv.GetContractCaller(ctx)
	auth, err := cc.Auth()
	if err != nil {
		return err
	}

	tx, err := cc.ExecutorContract.SkipCipherExecution(auth, a.batchIndex)
	if err != nil {
		// XXX consider handling the error somehow
		log.Printf("Error creating skip cipher execution tx: %s", err)
		return nil
	}
	runenv.WatchTransaction(tx)

	return nil
}

func (a SkipCipherBatch) String() string {
	return fmt.Sprintf("=> executor contract: skip cipher batch %d", a.batchIndex)
}

// Accuse is an action accusing the executor of a given half step at the keyper slasher.
type Accuse struct {
	halfStep    uint64
	keyperIndex uint64 // index of the accuser, not the executor
}

func (a Accuse) Run(ctx context.Context, runenv IRunEnv) error {
	log.Printf("=====%s", a)

	cc := runenv.GetContractCaller(ctx)
	auth, err := cc.Auth()
	if err != nil {
		return err
	}

	tx, err := cc.KeyperSlasher.Accuse(auth, a.halfStep, a.keyperIndex)
	if err != nil {
		return err
	}
	runenv.WatchTransaction(tx)

	return nil
}

func (a Accuse) String() string {
	return fmt.Sprintf("=> keyper slasher: accuse for half step %d", a.halfStep)
}

// Appeal is an action countering an earlier invalid accusation.
type Appeal struct {
	authorization contract.Authorization
}

func (a Appeal) Run(ctx context.Context, runenv IRunEnv) error {
	log.Printf("=====%s", a)

	cc := runenv.GetContractCaller(ctx)
	auth, err := cc.Auth()
	if err != nil {
		return err
	}

	tx, err := cc.KeyperSlasher.Appeal(auth, a.authorization)
	if err != nil {
		return err
	}
	runenv.WatchTransaction(tx)

	return nil
}

func (a Appeal) String() string {
	return fmt.Sprintf("=> keyper slasher: appeal for half step %d", a.authorization.HalfStep)
}

// EonKeyBroadcast is an action sending a vote for an eon public key to the key broadcast contract.
type EonKeyBroadcast struct {
	keyperIndex     uint64
	startBatchIndex uint64
	eonPublicKey    *shcrypto.EonPublicKey
}

func (a EonKeyBroadcast) Run(ctx context.Context, runenv IRunEnv) error {
	log.Printf("=====%s", a)

	cc := runenv.GetContractCaller(ctx)
	auth, err := cc.Auth()
	if err != nil {
		return err
	}

	tx, err := cc.KeyBroadcastContract.Vote(
		auth,
		a.keyperIndex,
		a.startBatchIndex,
		a.eonPublicKey.Marshal(),
	)
	if err != nil {
		return err
	}
	runenv.WatchTransaction(tx)

	return nil
}

func (a EonKeyBroadcast) String() string {
	return fmt.Sprintf("=> key broadcast contract: voting for eon key with start batch %d", a.startBatchIndex)
}
