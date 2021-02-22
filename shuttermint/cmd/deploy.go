package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/contract/erc1820"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	defaultConfigChangeHeadsUpBlocks = 20
	defaultAppealBlocks              = 20
	deployDefaultTimeout             = 300 * time.Second
)

var (
	key      *ecdsa.PrivateKey
	gasPrice *big.Int
)

var deployFlags struct {
	OwnerKey    string
	EthereumURL string
	GasPrice    string
	NoERC1820   bool
	OutputFile  string
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy all Shutter contracts",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), deployDefaultTimeout)
		defer cancel()
		var err error

		key, err = crypto.HexToECDSA(strings.TrimPrefix(deployFlags.OwnerKey, "0x"))
		failIfError(err)

		client, err := ethclient.DialContext(ctx, deployFlags.EthereumURL)
		failIfError(err)

		if deployFlags.GasPrice == "" {
			gasPrice, err = client.SuggestGasPrice(ctx)
			failIfError(err)
		} else {
			gasPriceGWei, ok := new(big.Int).SetString(deployFlags.GasPrice, 10)
			if !ok {
				log.Fatalf("Invalid gas price %s", deployFlags.GasPrice)
			}
			gasPrice = gweiToWei(gasPriceGWei)
		}

		address := crypto.PubkeyToAddress(key.PublicKey)
		balance, err := client.BalanceAt(ctx, address, nil)
		failIfError(err)

		log.Printf("Deploy Address: %s", address)
		log.Printf("Balance: %f ETH", weiToEther(balance))
		log.Printf("Gas Price: %f GWei", weiToGwei(gasPrice))
		log.Printf("Available gas: %d", new(big.Int).Quo(balance, gasPrice))

		deploy(ctx, client)
	},
}

func init() {
	deployCmd.PersistentFlags().StringVarP(
		&deployFlags.OwnerKey,
		"owner-key",
		"k",
		"",
		"private key of the deployer",
	)
	deployCmd.MarkPersistentFlagRequired("owner-key")
	deployCmd.PersistentFlags().StringVar(
		&deployFlags.GasPrice,
		"gas-price",
		"",
		"gas price in GWei (default: use suggested one)",
	)
	deployCmd.PersistentFlags().StringVarP(
		&deployFlags.EthereumURL,
		"ethereum-url",
		"e",
		"",
		"Ethereum RPC URL",
	)
	deployCmd.MarkPersistentFlagRequired("ethereum-url")
	deployCmd.PersistentFlags().BoolVar(
		&deployFlags.NoERC1820,
		"no-erc1820",
		false,
		"don't deploy the ERC1820 contract",
	)
	deployCmd.PersistentFlags().StringVarP(
		&deployFlags.OutputFile,
		"output",
		"o",
		"",
		"if given, store the contract addresses as a JSON file at this path",
	)
}

func waitForTransactionReceipt(ctx context.Context, cl *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := cl.TransactionReceipt(ctx, txHash)
		if err == ethereum.NotFound {
			time.Sleep(time.Second)
			continue
		}
		return receipt, err
	}
}

func waitForTransactions(ctx context.Context, client *ethclient.Client, txs []*types.Transaction) ([]*types.Receipt, error) {
	defer fmt.Print("\n")
	var res []*types.Receipt

	failedTxs := []int{}
	for i, tx := range txs {
		receipt, err := waitForTransactionReceipt(ctx, client, tx.Hash())
		if err != nil {
			return res, err
		}
		res = append(res, receipt)
		if receipt.Status != 1 {
			fmt.Print("X")
			failedTxs = append(failedTxs, i)
		} else {
			fmt.Print(".")
		}
	}

	if len(failedTxs) > 0 {
		log.Printf("The following transactions have failed:")
		for _, i := range failedTxs {
			log.Printf("%s", txs[i].Hash().Hex())
		}
		return res, errors.New("some txs have failed")
	}

	return res, nil
}

func makeAuth(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	return auth, nil
}

func maybeDeployERC1820(ctx context.Context, client *ethclient.Client) {
	deployed, err := erc1820.IsDeployed(ctx, client)
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	if deployed {
		log.Print("erc1820 contract already deployed")
		return
	}
	log.Print("Deploying erc1820 contract")
	err = erc1820.DeployContract(ctx, client, key)
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
}

func deploy(ctx context.Context, client *ethclient.Client) {
	if !deployFlags.NoERC1820 {
		maybeDeployERC1820(ctx, client)
	}

	auth, err := makeAuth(ctx, client, key)
	failIfError(err)
	auth.Context = ctx
	var txs []*types.Transaction
	var tx *types.Transaction

	addTx := func() {
		failIfError(err)
		txs = append(txs, tx)
		auth.Nonce.SetInt64(auth.Nonce.Int64() + 1)
	}

	configAddress, tx, _, err := contract.DeployConfigContract(
		auth,
		client,
		defaultConfigChangeHeadsUpBlocks,
	)
	addTx()

	broadcastAddress, tx, _, err := contract.DeployKeyBroadcastContract(auth, client, configAddress)
	addTx()

	feeAddress, tx, _, err := contract.DeployFeeBankContract(auth, client)
	addTx()

	batcherAddress, tx, _, err := contract.DeployBatcherContract(auth, client, configAddress, feeAddress)
	addTx()

	executorAddress, tx, _, err := contract.DeployExecutorContract(auth, client, configAddress, batcherAddress)
	addTx()

	tokenAddress, tx, _, err := contract.DeployTestDepositTokenContract(auth, client)
	addTx()

	depositAddress, tx, _, err := contract.DeployDepositContract(auth, client, tokenAddress)
	addTx()

	targetAddress, tx, _, err := contract.DeployTestTargetContract(auth, client, executorAddress)
	addTx()

	// The keyper slasher requires the deposit contract to exist. Since at this point it doesn't,
	// gas estimation will fail, so we simply set it manually.
	auth.GasLimit = 2000000
	keyperSlasherAddress, tx, _, err := contract.DeployKeyperSlasher(
		auth,
		client,
		big.NewInt(defaultAppealBlocks),
		configAddress,
		executorAddress,
		depositAddress,
	)
	auth.GasLimit = 0
	addTx()

	receipts, err := waitForTransactions(ctx, client, txs)
	failIfError(err)

	totalGasUsed := uint64(0)
	for _, receipt := range receipts {
		totalGasUsed += receipt.GasUsed
	}
	fmt.Println("Gas used:", totalGasUsed)

	fmt.Println("ConfigContract address:", configAddress.Hex())
	fmt.Println("KeyBroadcastContract address:", broadcastAddress.Hex())
	fmt.Println("FeeBankContract address:", feeAddress.Hex())
	fmt.Println("BatcherContract address:", batcherAddress.Hex())
	fmt.Println("ExecutorContract address:", executorAddress.Hex())
	fmt.Println("TokenContract address:", tokenAddress.Hex())
	fmt.Println("DepositContract address:", depositAddress.Hex())
	fmt.Println("KeyperSlasher address:", keyperSlasherAddress.Hex())
	fmt.Println("TargetContract address:", targetAddress.Hex())

	if deployFlags.OutputFile != "" {
		j := sandbox.ContractsJSON{
			ConfigContract:        configAddress.Hex(),
			KeyBroadcastContract:  broadcastAddress.Hex(),
			FeeBankContract:       feeAddress.Hex(),
			BatcherContract:       batcherAddress.Hex(),
			ExecutorContract:      executorAddress.Hex(),
			TokenContract:         tokenAddress.Hex(),
			DepositContract:       depositAddress.Hex(),
			KeyperSlasherContract: keyperSlasherAddress.Hex(),
			TargetContract:        targetAddress.Hex(),
		}
		s, err := json.MarshalIndent(j, "", "    ")
		failIfError(err)
		err = ioutil.WriteFile(deployFlags.OutputFile, s, 0o644)
		failIfError(err)
		fmt.Println("addresses written to", deployFlags.OutputFile)
	}
}

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

func weiToGwei(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.GWei))
}

func gweiToWei(gwei *big.Int) *big.Int {
	return new(big.Int).Mul(gwei, big.NewInt(params.GWei))
}

func failIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
