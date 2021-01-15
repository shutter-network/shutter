package keyper

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/rpc/client"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"golang.org/x/sync/errgroup"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper/shutterevents"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// BatchParams describes the parameters for single Batch identified by the BatchIndex
type BatchParams = contract.BatchParams

// BatchState is used to manage the key generation process for a single batch inside the keyper
type BatchState struct {
	BatchParams                   BatchParams
	KeyperConfig                  KeyperConfig
	MessageSender                 MessageSender
	ContractCaller                *ContractCaller
	decryptionSignatureAdded      chan shutterevents.DecryptionSignature
	cipherExecutionParams         chan CipherExecutionParams
	startBlockSeen                chan struct{}
	endBlockSeen                  chan struct{}
	executionTimeoutBlockSeen     chan struct{}
	startBlockSeenOnce            sync.Once
	endBlockSeenOnce              sync.Once
	executionTimeoutBlockSeenOnce sync.Once
}

// KeyperConfig contains validated configuration parameters for the keyper client
type KeyperConfig struct {
	ChainID                     string
	ShuttermintURL              string
	EthereumURL                 string
	DBDir                       string
	SigningKey                  *ecdsa.PrivateKey
	ValidatorKey                ed25519.PrivateKey
	EncryptionKey               *ecies.PrivateKey
	ConfigContractAddress       common.Address
	BatcherContractAddress      common.Address
	KeyBroadcastContractAddress common.Address
	ExecutorContractAddress     common.Address
	DepositContractAddress      common.Address
	ExecutionStaggering         uint64
}

// Keyper is used to run the keyper key generation
type Keyper struct {
	sync.Mutex

	Config KeyperConfig
	ethcl  *ethclient.Client
	shmcl  client.Client

	configContract       *contract.ConfigContract
	keyBroadcastContract *contract.KeyBroadcastContract
	batcherContract      *contract.BatcherContract
	executorContract     *contract.ExecutorContract
	depositContract      *contract.DepositContract

	batchConfigs          map[uint64]contract.BatchConfig
	batches               map[uint64]*BatchState
	startBlock            *big.Int
	checkedIn             bool
	txs                   <-chan coretypes.ResultEvent
	ctx                   context.Context
	newHeaders            chan *types.Header // start new batches when new block headers arrive
	group                 *errgroup.Group
	ms                    MessageSender
	executor              Executor
	cipherExecutionParams chan CipherExecutionParams
	keyperEncryptionKeys  map[common.Address]*ecies.PublicKey
	dkg                   map[uint64]*DKGInstance
}

// MessageSender defines the interface of sending messages to shuttermint.
type MessageSender interface {
	SendMessage(context.Context, *shmsg.Message) error
}

// RPCMessageSender signs messages and sends them via RPC to shuttermint.
type RPCMessageSender struct {
	rpcclient  client.Client
	chainID    string
	signingKey *ecdsa.PrivateKey
}

var _ MessageSender = &RPCMessageSender{}

// MockMessageSender sends all messages to a channel so that they can be checked for testing.
type MockMessageSender struct {
	Msgs chan *shmsg.Message
}

var _ MessageSender = &MockMessageSender{}

// ContractCaller interacts with the contracts on Ethereum.
type ContractCaller struct {
	Ethclient  *ethclient.Client
	signingKey *ecdsa.PrivateKey

	ConfigContract       *contract.ConfigContract
	KeyBroadcastContract *contract.KeyBroadcastContract
	BatcherContract      *contract.BatcherContract
	ExecutorContract     *contract.ExecutorContract
	DepositContract      *contract.DepositContract
}

// NewContractCaller creates a new ContractCaller.
func NewContractCaller(
	ethcl *ethclient.Client,
	signingKey *ecdsa.PrivateKey,
	configContract *contract.ConfigContract,
	keyBroadcastContract *contract.KeyBroadcastContract,
	batcherContract *contract.BatcherContract,
	executorContract *contract.ExecutorContract,
	depositContract *contract.DepositContract,
) ContractCaller {
	return ContractCaller{
		Ethclient:  ethcl,
		signingKey: signingKey,

		ConfigContract:       configContract,
		KeyBroadcastContract: keyBroadcastContract,
		BatcherContract:      batcherContract,
		ExecutorContract:     executorContract,
		DepositContract:      depositContract,
	}
}

// Executor is responsible for making sure batches are executed.
type Executor struct {
	ctx                   context.Context
	client                *ethclient.Client
	cc                    *ContractCaller
	cipherExecutionParams <-chan CipherExecutionParams
}

// CipherExecutionParams is the set of parameters necessary to execute a batch.
type CipherExecutionParams struct {
	BatchIndex              uint64
	CipherBatchHash         common.Hash
	DecryptionKey           *ecdsa.PrivateKey
	DecryptedTxs            [][]byte
	DecryptionSignerIndices []uint64
	DecryptionSignatures    [][]byte
}
