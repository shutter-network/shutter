package fx

import (
	"encoding/gob"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
	"github.com/brainbot-com/shutter/shuttermint/shcrypto"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// IAction describes an action to run as determined by the Decider's Decide method.
type IAction interface {
	// IsExpired checks if the action expired because some time limit is reached or because the
	// result of the action has been achieved (we or someone else may have performed the same
	// action). We could think about having a dedicated IsDone method.
	IsExpired(world observe.World) bool
}

// MainChainTX is an action that sends a transaction to the main chain.
type MainChainTX interface {
	IAction
	SendTX(caller *contract.Caller, auth *bind.TransactOpts) (*types.Transaction, error)
}

var (
	_ IAction     = SendShuttermintMessage{}
	_ MainChainTX = ExecuteCipherBatch{}
	_ MainChainTX = ExecutePlainBatch{}
	_ MainChainTX = SkipCipherBatch{}
	_ MainChainTX = Accuse{}
	_ MainChainTX = Appeal{}
	_ MainChainTX = EonKeyBroadcast{}
)

func init() {
	for _, a := range []IAction{
		&SendShuttermintMessage{},
		&ExecuteCipherBatch{},
		&ExecutePlainBatch{},
		&SkipCipherBatch{},
		&Accuse{},
		&Appeal{},
		&EonKeyBroadcast{},
	} {
		gob.Register(a)
	}
}

// SendShuttermintMessage is an Action that sends a message to shuttermint.
type SendShuttermintMessage struct {
	Description string
	Msg         *shmsg.Message
}

func (a SendShuttermintMessage) String() string {
	return fmt.Sprintf("=> shuttermint: %s", a.Description)
}

func (a SendShuttermintMessage) IsExpired(world observe.World) bool {
	return false
}

// ExecuteCipherBatch is an Action that instructs the executor contract to execute a cipher batch.
type ExecuteCipherBatch struct {
	BatchIndex      uint64
	CipherBatchHash [32]byte
	Transactions    [][]byte
	KeyperIndex     uint64
}

func (a ExecuteCipherBatch) SendTX(caller *contract.Caller, auth *bind.TransactOpts) (*types.Transaction, error) {
	return caller.ExecutorContract.ExecuteCipherBatch(
		auth, a.BatchIndex, a.CipherBatchHash, a.Transactions, a.KeyperIndex,
	)
}

func (a ExecuteCipherBatch) String() string {
	return fmt.Sprintf("=> executor contract: execute cipher batch %d with %d txs", a.BatchIndex, len(a.Transactions))
}

func (a ExecuteCipherBatch) IsExpired(world observe.World) bool {
	halfStep := 2 * a.BatchIndex
	return world.MainChain.NumExecutionHalfSteps > halfStep
}

// ExecutePlainBatch is an Action that instructs the executor contract to execute a plain batch.
type ExecutePlainBatch struct {
	BatchIndex   uint64
	Transactions [][]byte
}

func (a ExecutePlainBatch) SendTX(caller *contract.Caller, auth *bind.TransactOpts) (*types.Transaction, error) {
	return caller.ExecutorContract.ExecutePlainBatch(auth, a.BatchIndex, a.Transactions)
}

func (a ExecutePlainBatch) String() string {
	return fmt.Sprintf("=> executor contract: execute plain batch %d with %d txs", a.BatchIndex, len(a.Transactions))
}

func (a ExecutePlainBatch) IsExpired(world observe.World) bool {
	halfStep := 2*a.BatchIndex + 1
	return world.MainChain.NumExecutionHalfSteps > halfStep
}

// SkipCipherBatch is an Action that instructs the executor contract to skip a cipher batch.
type SkipCipherBatch struct {
	BatchIndex uint64
}

func (a SkipCipherBatch) SendTX(caller *contract.Caller, auth *bind.TransactOpts) (*types.Transaction, error) {
	return caller.ExecutorContract.SkipCipherExecution(auth, a.BatchIndex)
}

func (a SkipCipherBatch) String() string {
	return fmt.Sprintf("=> executor contract: skip cipher batch %d", a.BatchIndex)
}

func (a SkipCipherBatch) IsExpired(world observe.World) bool {
	halfStep := 2 * a.BatchIndex
	return world.MainChain.NumExecutionHalfSteps > halfStep
}

// Accuse is an action accusing the executor of a given half step at the keyper slasher.
type Accuse struct {
	HalfStep    uint64
	KeyperIndex uint64 // index of the accuser, not the executor
}

func (a Accuse) SendTX(caller *contract.Caller, auth *bind.TransactOpts) (*types.Transaction, error) {
	return caller.KeyperSlasher.Accuse(auth, a.HalfStep, a.KeyperIndex)
}

func (a Accuse) String() string {
	return fmt.Sprintf("=> keyper slasher: accuse for half step %d", a.HalfStep)
}

func (a Accuse) IsExpired(world observe.World) bool {
	return false
}

// Appeal is an action countering an earlier invalid accusation.
type Appeal struct {
	Authorization contract.Authorization
}

func (a Appeal) SendTX(caller *contract.Caller, auth *bind.TransactOpts) (*types.Transaction, error) {
	return caller.KeyperSlasher.Appeal(auth, a.Authorization)
}

func (a Appeal) String() string {
	return fmt.Sprintf("=> keyper slasher: appeal for half step %d", a.Authorization.HalfStep)
}

func (a Appeal) IsExpired(world observe.World) bool {
	return false
}

// EonKeyBroadcast is an action sending a vote for an eon public key to the key broadcast contract.
type EonKeyBroadcast struct {
	KeyperIndex     uint64
	StartBatchIndex uint64
	EonPublicKey    *shcrypto.EonPublicKey
}

func (a EonKeyBroadcast) SendTX(caller *contract.Caller, auth *bind.TransactOpts) (*types.Transaction, error) {
	return caller.KeyBroadcastContract.Vote(
		auth,
		a.KeyperIndex,
		a.StartBatchIndex,
		a.EonPublicKey.Marshal(),
	)
}

func (a EonKeyBroadcast) String() string {
	return fmt.Sprintf("=> key broadcast contract: voting for eon key with start batch %d", a.StartBatchIndex)
}

func (a EonKeyBroadcast) IsExpired(world observe.World) bool {
	return false
}
