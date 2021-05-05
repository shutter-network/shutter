package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/contract/erc1820"
	"github.com/brainbot-com/shutter/shuttermint/medley"
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
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		key, err = crypto.HexToECDSA(strings.TrimPrefix(deployFlags.OwnerKey, "0x"))
		if err != nil {
			return err
		}

		if deployFlags.GasPrice != "" {
			gasPriceGWei, ok := new(big.Int).SetString(deployFlags.GasPrice, 10)
			if !ok {
				return errors.Errorf("--gas-price: invalid gas price '%s'", deployFlags.GasPrice)
			}
			if gasPriceGWei.Sign() < 0 {
				return errors.Errorf("--gas-price: must be non-negative")
			}
			gasPrice = gweiToWei(gasPriceGWei)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), deployDefaultTimeout)
		defer cancel()
		var err error

		client, err := ethclient.DialContext(ctx, deployFlags.EthereumURL)
		failIfError(err)

		if gasPrice == nil {
			gasPrice, err = client.SuggestGasPrice(ctx)
			failIfError(err)
		}

		address := crypto.PubkeyToAddress(key.PublicKey)
		balance, err := client.BalanceAt(ctx, address, nil)
		failIfError(err)

		log.Printf("Deploy Address: %s", address)
		log.Printf("Balance: %f ETH", weiToEther(balance))
		log.Printf("Gas Price: %f GWei", weiToGwei(gasPrice))
		log.Printf("Available gas: %d", new(big.Int).Quo(balance, gasPrice))

		if !deployFlags.NoERC1820 {
			maybeDeployERC1820(ctx, client)
		}
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
	auth, err := sandbox.InitTransactOpts(ctx, client, key)
	failIfError(err)
	auth.GasPrice = gasPrice
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

	txHashes := []common.Hash{}
	for _, tx := range txs {
		txHashes = append(txHashes, tx.Hash())
	}
	receipts, err := medley.WaitMinedMany(ctx, client, txHashes)
	failIfError(err)

	totalGasUsed := uint64(0)
	for _, receipt := range receipts {
		totalGasUsed += receipt.GasUsed
	}
	fmt.Println("Gas used:", totalGasUsed)

	fmt.Println("      ConfigContract:", configAddress.Hex())
	fmt.Println("KeyBroadcastContract:", broadcastAddress.Hex())
	fmt.Println("     FeeBankContract:", feeAddress.Hex())
	fmt.Println("     BatcherContract:", batcherAddress.Hex())
	fmt.Println("    ExecutorContract:", executorAddress.Hex())
	fmt.Println("       TokenContract:", tokenAddress.Hex())
	fmt.Println("     DepositContract:", depositAddress.Hex())
	fmt.Println("       KeyperSlasher:", keyperSlasherAddress.Hex())
	fmt.Println("      TargetContract:", targetAddress.Hex())

	if deployFlags.OutputFile != "" {
		outputFile := filepath.Clean(deployFlags.OutputFile)
		j := sandbox.ContractsJSON{
			ConfigContract:        configAddress,
			KeyBroadcastContract:  broadcastAddress,
			FeeBankContract:       feeAddress,
			BatcherContract:       batcherAddress,
			ExecutorContract:      executorAddress,
			TokenContract:         tokenAddress,
			DepositContract:       depositAddress,
			KeyperSlasherContract: keyperSlasherAddress,
			TargetContract:        targetAddress,
		}
		s, err := json.MarshalIndent(j, "", "    ")
		failIfError(err)
		err = ioutil.WriteFile(outputFile, s, 0o644)
		failIfError(err)
		fmt.Println("addresses written to", outputFile)
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
