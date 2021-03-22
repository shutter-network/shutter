// Package erc1820 defines methods to deploy the erc1820 registry contract, see
// https://eips.ethereum.org/EIPS/eip-1820
package erc1820

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

const (
	contractAddressHex   = "0x1820a4B7618BdE71Dce8cdc73aAB6C95905faD24"
	deploymentAddressHex = "0xa990077c3205cbDf861e17Fa532eeB069cE9fF96"
	deployerAmountInt    = 80000000000000000 // 0.08 ETH
)

// DeploymentTransaction is the transaction deploying the ERC1820 contract.
var DeploymentTransaction *types.Transaction

// DeployerAddress is the sender address of the ERC1820 deployment transaction.
var DeployerAddress common.Address

// DeployerAmount is the amount the ERC1820 deployer is expected to be funded with before
// the deployment transaction is sent.
var DeployerAmount *big.Int

// ContractAddress is the address of the deployed ERC1820 contract.
var ContractAddress common.Address

func init() {
	initDeploymentTransaction()

	DeployerAddress = common.HexToAddress(deploymentAddressHex)
	DeployerAmount = big.NewInt(deployerAmountInt)
	ContractAddress = common.HexToAddress(contractAddressHex)
}

func initDeploymentTransaction() {
	txBytes, err := hexutil.Decode(deploymentTransactionHex)
	if err != nil {
		panic(err)
	}
	txStream := rlp.NewStream(bytes.NewReader(txBytes), 0)
	DeploymentTransaction = new(types.Transaction)
	err = DeploymentTransaction.DecodeRLP(txStream)
	if err != nil {
		panic(err)
	}
}

// NewFundingTransaction creates a new transaction sending the required funds from the
// account given by a private key to the ERC1820 deployer address.
func NewFundingTransaction(ctx context.Context, client *ethclient.Client, key *ecdsa.PrivateKey) (*types.Transaction, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	from := crypto.PubkeyToAddress(key.PublicKey)
	signer := types.NewEIP155Signer(chainID)

	nonce, err := client.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	unsignedTx := types.NewTransaction(nonce, DeployerAddress, DeployerAmount, 21000, gasPrice, []byte{})
	return types.SignTx(unsignedTx, signer, key)
}

// IsDeployed checks if the ERC1820 contract is deployed.
func IsDeployed(ctx context.Context, client *ethclient.Client) (bool, error) {
	code, err := client.CodeAt(ctx, ContractAddress, nil)
	if err != nil {
		return false, errors.Wrap(err, "cannot get erc1820 bytecode")
	}
	return len(code) > 0, nil
}

// DeployContract deploys the ERC1820 contract after funding the deployer account according
// to https://eips.ethereum.org/EIPS/eip-1820.
func DeployContract(ctx context.Context, client *ethclient.Client, key *ecdsa.PrivateKey) error {
	tx, err := NewFundingTransaction(ctx, client, key)
	if err != nil {
		return err
	}
	err = client.SendTransaction(ctx, tx)
	if err != nil {
		return err
	}
	err = waitMinedSuccessful(ctx, client, tx)
	if err != nil {
		return err
	}

	err = client.SendTransaction(ctx, DeploymentTransaction)
	if err != nil {
		return err
	}
	err = waitMinedSuccessful(ctx, client, DeploymentTransaction)
	if err != nil {
		return err
	}

	return nil
}

func waitMinedSuccessful(ctx context.Context, client *ethclient.Client, tx *types.Transaction) error {
	// bind.WaitMined doesn't work for some reason, at least not with Ganache
	const sleepDuration = time.Millisecond * 500
	var receipt *types.Receipt
	var err error
	for receipt == nil {
		receipt, err = client.TransactionReceipt(ctx, tx.Hash())
		if err != nil && err != ethereum.NotFound {
			return err
		}
		time.Sleep(sleepDuration)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.Errorf("deployment of ERC1820 contract failed")
	}
	return nil
}
