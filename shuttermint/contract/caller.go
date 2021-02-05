package contract

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractCaller interacts with the contracts on Ethereum.
type Caller struct {
	Ethclient  *ethclient.Client
	signingKey *ecdsa.PrivateKey

	ConfigContract       *ConfigContract
	KeyBroadcastContract *KeyBroadcastContract
	BatcherContract      *BatcherContract
	ExecutorContract     *ExecutorContract
	DepositContract      *DepositContract
	KeyperSlasher        *KeyperSlasher
}

// NewContractCaller creates a new ContractCaller.
func NewContractCaller(
	ethcl *ethclient.Client,
	signingKey *ecdsa.PrivateKey,
	configContract *ConfigContract,
	keyBroadcastContract *KeyBroadcastContract,
	batcherContract *BatcherContract,
	executorContract *ExecutorContract,
	depositContract *DepositContract,
	keyperSlasher *KeyperSlasher,
) Caller {
	return Caller{
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

// Address returns the address of the account that is used to send transactions.
func (cc *Caller) Address() common.Address {
	return crypto.PubkeyToAddress(cc.signingKey.PublicKey)
}

// Auth returns a new transactor with initialized key, nonce, and gas price.
func (cc *Caller) Auth() (*bind.TransactOpts, error) {
	chainID, err := cc.Ethclient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(cc.signingKey, chainID)
	if err != nil {
		return nil, err
	}

	nonce, err := cc.Ethclient.PendingNonceAt(context.Background(), cc.Address())
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	gasPrice, err := cc.Ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.GasPrice = gasPrice
	return auth, nil
}
