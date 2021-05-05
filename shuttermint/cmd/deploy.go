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

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/contract/erc1820"
	"github.com/brainbot-com/shutter/shuttermint/medley/txbatch"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), deployDefaultTimeout)
		defer cancel()
		var err error

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

func batchContractDeployments(batch *txbatch.TXBatch) (*sandbox.ContractsJSON, error) {
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

	executorAddress, tx, _, err := contract.DeployExecutorContract(auth, client, configAddress, batcherAddress)
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

	targetAddress, tx, _, err := contract.DeployTestTargetContract(auth, client, executorAddress)
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

	return &sandbox.ContractsJSON{
		ConfigContract:        configAddress,
		KeyBroadcastContract:  broadcastAddress,
		FeeBankContract:       feeAddress,
		BatcherContract:       batcherAddress,
		ExecutorContract:      executorAddress,
		TokenContract:         tokenAddress,
		DepositContract:       depositAddress,
		KeyperSlasherContract: keyperSlasherAddress,
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

	j, err := batchContractDeployments(batch)
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
	fmt.Println("Gas used:", totalGasUsed)

	fmt.Println("      ConfigContract:", j.ConfigContract.Hex())
	fmt.Println("KeyBroadcastContract:", j.KeyBroadcastContract.Hex())
	fmt.Println("     FeeBankContract:", j.FeeBankContract.Hex())
	fmt.Println("     BatcherContract:", j.BatcherContract.Hex())
	fmt.Println("    ExecutorContract:", j.ExecutorContract.Hex())
	fmt.Println("       TokenContract:", j.TokenContract.Hex())
	fmt.Println("     DepositContract:", j.DepositContract.Hex())
	fmt.Println("       KeyperSlasher:", j.KeyperSlasherContract.Hex())
	fmt.Println("      TargetContract:", j.TargetContract.Hex())

	if deployFlags.OutputFile != "" {
		outputFile := filepath.Clean(deployFlags.OutputFile)
		s, err := json.MarshalIndent(j, "", "    ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(outputFile, s, 0o644)
		if err != nil {
			return err
		}
		fmt.Println("addresses written to", outputFile)
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
