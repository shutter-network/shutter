package keyper

import (
	"crypto/ecdsa"
	"crypto/ed25519"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/ecies"
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
	MainChainFollowDistance     uint64 // in main chain blocks
	ExecutionStaggering         uint64 // in main chain blocks
	DKGPhaseLength              uint64 // in shuttermint blocks
}
