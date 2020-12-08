package keyper

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Address returns the address of the account that is used to send transactions.
func (cc *ContractCaller) Address() common.Address {
	return crypto.PubkeyToAddress(cc.signingKey.PublicKey)
}

// Auth returns a new transactor with initialized key, nonce, and gas price.
func (cc *ContractCaller) Auth() (*bind.TransactOpts, error) {
	auth := bind.NewKeyedTransactor(cc.signingKey)

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
