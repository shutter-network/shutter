// Package deploy contains the implementation of the deploy subcommand
package deploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/contract"
	"github.com/shutter-network/shutter/shuttermint/contract/erc1820"
	"github.com/shutter-network/shutter/shuttermint/medley/txbatch"
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

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy all Shutter contracts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
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
		if deployFlags.OutputFile != "" {
			deployFlags.OutputFile = filepath.Clean(deployFlags.OutputFile)
		}

		key, err = crypto.HexToECDSA(strings.TrimPrefix(deployFlags.OwnerKey, "0x"))
		if err != nil {
			return errors.WithMessage(err, "invalid --owner-key argument")
		}

		ctx, cancel := context.WithTimeout(context.Background(), deployDefaultTimeout)
		defer cancel()

		client, err := ethclient.DialContext(ctx, deployFlags.EthereumURL)
		if err != nil {
			return err
		}

		if gasPrice == nil {
			gasPrice, err = client.SuggestGasPrice(ctx)
			if err != nil {
				return err
			}
		}

		address := crypto.PubkeyToAddress(key.PublicKey)
		balance, err := client.BalanceAt(ctx, address, nil)
		if err != nil {
			return err
		}

		log.Printf("Deploy Address: %s", address)
		log.Printf("Balance: %f ETH", weiToEther(balance))
		log.Printf("Gas Price: %f GWei", weiToGwei(gasPrice))
		log.Printf("Available gas: %d", new(big.Int).Quo(balance, gasPrice))

		if !deployFlags.NoERC1820 {
			err := maybeDeployERC1820(ctx, client)
			if err != nil {
				return err
			}
		}
		return deploy(ctx, client)
	},
}

func init() {
	DeployCmd.PersistentFlags().StringVarP(
		&deployFlags.OwnerKey,
		"owner-key",
		"k",
		"",
		"private key of the deployer",
	)
	DeployCmd.MarkPersistentFlagRequired("owner-key")
	DeployCmd.PersistentFlags().StringVar(
		&deployFlags.GasPrice,
		"gas-price",
		"",
		"gas price in GWei (default: use suggested one)",
	)
	DeployCmd.PersistentFlags().StringVarP(
		&deployFlags.EthereumURL,
		"ethereum-url",
		"e",
		"",
		"Ethereum RPC URL",
	)
	DeployCmd.MarkPersistentFlagRequired("ethereum-url")
	DeployCmd.PersistentFlags().BoolVar(
		&deployFlags.NoERC1820,
		"no-erc1820",
		false,
		"don't deploy the ERC1820 contract",
	)
	DeployCmd.PersistentFlags().StringVarP(
		&deployFlags.OutputFile,
		"output",
		"o",
		"",
		"if given, store the contract addresses as a JSON file at this path",
	)
}

func maybeDeployERC1820(ctx context.Context, client *ethclient.Client) error {
	deployed, err := erc1820.IsDeployed(ctx, client)
	if err != nil {
		return err
	}
	if deployed {
		log.Print("erc1820 contract already deployed")
		return nil
	}
	log.Print("Deploying erc1820 contract")
	return erc1820.DeployContract(ctx, client, key)
}

func batchContractDeployments(batch *txbatch.TXBatch) (*Contracts, error) {
	var tx *types.Transaction
	var err error

	auth := batch.TransactOpts
	client := batch.Ethclient

	configAddress, tx, _, err := contract.DeployConfigContract(auth, client, defaultConfigChangeHeadsUpBlocks)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	broadcastAddress, tx, _, err := contract.DeployKeyBroadcastContract(auth, client, configAddress)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	feeAddress, tx, _, err := contract.DeployFeeBankContract(auth, client)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	batcherAddress, tx, _, err := contract.DeployBatcherContract(auth, client, configAddress, feeAddress)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	tokenAddress, tx, _, err := contract.DeployTestDepositTokenContract(auth, client)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	depositAddress, tx, _, err := contract.DeployDepositContract(auth, client, tokenAddress)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	executorAddress, tx, _, err := contract.DeployExecutorContract(auth, client, configAddress, batcherAddress, depositAddress)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	targetProxyAddress, tx, _, err := contract.DeployTargetProxyContract(auth, client, executorAddress)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	targetAddress, tx, _, err := contract.DeployTestTargetContract(auth, client, targetProxyAddress)
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

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
	if err != nil {
		return nil, err
	}
	batch.Add(tx)

	return &Contracts{
		ConfigContract:        configAddress,
		KeyBroadcastContract:  broadcastAddress,
		FeeBankContract:       feeAddress,
		BatcherContract:       batcherAddress,
		ExecutorContract:      executorAddress,
		TokenContract:         tokenAddress,
		DepositContract:       depositAddress,
		KeyperSlasherContract: keyperSlasherAddress,
		TargetProxyContract:   targetProxyAddress,
		TargetContract:        targetAddress,
	}, nil
}

func deploy(ctx context.Context, client *ethclient.Client) error {
	batch, err := txbatch.New(ctx, client, key)
	if err != nil {
		return err
	}

	batch.TransactOpts.GasPrice = gasPrice
	batch.TransactOpts.Context = ctx

	contracts, err := batchContractDeployments(batch)
	if err != nil {
		return err
	}

	receipts, err := batch.WaitMined(ctx)
	if err != nil {
		return err
	}
	totalGasUsed := uint64(0)
	for _, receipt := range receipts {
		totalGasUsed += receipt.GasUsed
	}
	fmt.Println("      ConfigContract:", contracts.ConfigContract.Hex())
	fmt.Println("KeyBroadcastContract:", contracts.KeyBroadcastContract.Hex())
	fmt.Println("     FeeBankContract:", contracts.FeeBankContract.Hex())
	fmt.Println("     BatcherContract:", contracts.BatcherContract.Hex())
	fmt.Println("    ExecutorContract:", contracts.ExecutorContract.Hex())
	fmt.Println("       TokenContract:", contracts.TokenContract.Hex())
	fmt.Println("     DepositContract:", contracts.DepositContract.Hex())
	fmt.Println("       KeyperSlasher:", contracts.KeyperSlasherContract.Hex())
	fmt.Println(" TargetProxyContract:", contracts.TargetProxyContract.Hex())
	fmt.Println("      TargetContract:", contracts.TargetContract.Hex())
	fmt.Println("")
	fmt.Println("            Gas used:", totalGasUsed)

	if deployFlags.OutputFile != "" {
		err := contracts.SaveJSON(deployFlags.OutputFile)
		if err != nil {
			return err
		}
		fmt.Println("addresses written to:", deployFlags.OutputFile)
	}
	return nil
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
