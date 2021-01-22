package keyper

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/rpc/client"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

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
	KeyperSlasherAddress        common.Address
	ExecutionStaggering         uint64
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
	KeyperSlasher        *contract.KeyperSlasher
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
	keyperSlasher *contract.KeyperSlasher,
) ContractCaller {
	return ContractCaller{
		Ethclient:  ethcl,
		signingKey: signingKey,

		ConfigContract:       configContract,
		KeyBroadcastContract: keyBroadcastContract,
		BatcherContract:      batcherContract,
		ExecutorContract:     executorContract,
		DepositContract:      depositContract,
		KeyperSlasher:        keyperSlasher,
	}
}
