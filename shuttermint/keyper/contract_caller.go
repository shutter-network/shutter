package keyper

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

// Address returns the address of the account that is used to send transactions.
func (cc *ContractCaller) Address() common.Address {
	return crypto.PubkeyToAddress(cc.signingKey.PublicKey)
}

// Client creates a new ethclient.
func (cc *ContractCaller) Client() (*ethclient.Client, error) {
	return ethclient.Dial(cc.ethereumURL)
}

// Auth returns a new transactor with initialized key, nonce, and gas price.
func (cc *ContractCaller) Auth() (*bind.TransactOpts, error) {
	cl, err := cc.Client()
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(cc.signingKey)

	nonce, err := cl.PendingNonceAt(context.Background(), cc.Address())
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// gasPrice, err := cl.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// auth.GasPrice = gasPrice
	auth.GasPrice = big.NewInt(int64(0))

	return auth, nil
}

// KeyBroadcastContract returns a bound instance of the key broadcast contract.
func (cc *ContractCaller) KeyBroadcastContract() (*contract.KeyBroadcastContract, error) {
	cl, err := cc.Client()
	if err != nil {
		return nil, err
	}

	return contract.NewKeyBroadcastContract(cc.keyBroadcastContractAddress, cl)
}

// BroadcastEncryptionKey makes sure that the encryption key is sent to the broadcasting contract.
func (cc *ContractCaller) BroadcastEncryptionKey(keyperIndex uint64, batchIndex uint64, encryptionKey []byte, signerIndices []uint64, signatures [][]byte) error {
	auth, err := cc.Auth()
	if err != nil {
		return err
	}
	auth.GasLimit = 100000

	c, err := cc.KeyBroadcastContract()
	if err != nil {
		return err
	}

	tx, err := c.BroadcastEncryptionKey2(
		auth,
		keyperIndex,
		batchIndex,
		encryptionKey,
		signerIndices,
		signatures,
	)
	if err != nil {
		return err
	}
	log.Printf("Sent tx: %s", tx.Hash().Hex())

	return nil
}
