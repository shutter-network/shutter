package keyper

import "log"
import "context"
import "crypto/ecdsa"
import "math/big"
import "github.com/ethereum/go-ethereum/accounts/abi/bind"
import "github.com/ethereum/go-ethereum/ethclient"
import "github.com/ethereum/go-ethereum/common"
import "github.com/ethereum/go-ethereum/crypto"
import "github.com/brainbot-com/shutter/shuttermint/contract"

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

	address := common.HexToAddress("0xD3e50BC9d970c3c07396842FCF9430C9Bd9ba2b9")
	return contract.NewKeyBroadcastContract(address, cl)
}

// BroadcastEncryptionKey makes sure that the encryption key is sent to the broadcasting contract.
func (cc *ContractCaller) BroadcastEncryptionKey(keyperIndex uint64, batchIndex uint64, encryptionKey []byte, signerIndices []uint64, signatures [][]byte) error {
	auth, err := cc.Auth()
	if err != nil {
		return err
	}

	c, err := cc.KeyBroadcastContract()
	if err != nil {
		return err
	}

	encryptionKeySized := [32]byte{}
	copy(encryptionKeySized[:], encryptionKey[:32])

	signerIndicesBig := []*big.Int{}
	for _, signerIndex := range signerIndices {
		signerIndexBig := big.NewInt(int64(signerIndex))
		signerIndicesBig = append(signerIndicesBig, signerIndexBig)
	}

	auth.GasLimit = 100000
	tx, err := c.BroadcastEncryptionKey(
		auth,
		big.NewInt(int64(keyperIndex)),
		big.NewInt(int64(batchIndex)),
		encryptionKeySized,
		signerIndicesBig,
		signatures,
	)
	if err != nil {
		return err
	}
	log.Printf("Sent tx: %s", tx.Hash().Hex())

	return nil
}
