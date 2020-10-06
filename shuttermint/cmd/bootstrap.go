package cmd

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/brainbot-com/shutter/shuttermint/contracts/configcontract"
	"github.com/brainbot-com/shutter/shuttermint/keyper"
)

var bootstrapFlags struct {
	ShuttermintURL   string
	EthereumURL      string
	BatchConfigIndex int
	ConfigContract   string
	SigningKey       string
}

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap Shuttermint by submitting the initial batch config",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap()
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	bootstrapCmd.PersistentFlags().StringVarP(
		&bootstrapFlags.ShuttermintURL,
		"shuttermint-url",
		"s",
		"http://localhost:26657",
		"Shuttermint RPC URL",
	)
	bootstrapCmd.PersistentFlags().StringVarP(
		&bootstrapFlags.EthereumURL,
		"ethereum-url",
		"e",
		"ws://localhost:8545/websocket",
		"Ethereum RPC URL",
	)
	bootstrapCmd.PersistentFlags().IntVarP(
		&bootstrapFlags.BatchConfigIndex,
		"index",
		"i",
		1,
		"index of the batch config to bootstrap with",
	)

	bootstrapCmd.PersistentFlags().StringVarP(
		&bootstrapFlags.ConfigContract,
		"config-contract",
		"c",
		"",
		"address of the contract from which to fetch config",
	)
	bootstrapCmd.MarkPersistentFlagRequired("config-contract")

	bootstrapCmd.PersistentFlags().StringVarP(
		&bootstrapFlags.SigningKey,
		"signing-key",
		"k",
		"",
		"private key of the keyper to send the message with",
	)
	bootstrapCmd.MarkPersistentFlagRequired("signing-key")
}

func bootstrap() {
	ethcl, err := ethclient.Dial(bootstrapFlags.EthereumURL)
	if err != nil {
		log.Fatalf("Error connecting to Ethereum node: %v", err)
	}

	shmcl, err := http.New(bootstrapFlags.ShuttermintURL, "/websocket")
	if err != nil {
		log.Fatalf("Error connecting to Shuttermint node: %v", err)
	}

	signingKey, err := crypto.HexToECDSA(bootstrapFlags.SigningKey)
	if err != nil {
		log.Fatalf("Invalid signing key: %v", err)
	}

	configContractAddress := common.HexToAddress(bootstrapFlags.ConfigContract)
	if bootstrapFlags.ConfigContract != configContractAddress.Hex() {
		log.Fatalf("Invalid config contract address %s", bootstrapFlags.ConfigContract)
	}
	configContract, err := configcontract.NewConfigContract(configContractAddress, ethcl)
	if err != nil {
		log.Fatalf("Failed to instantiate ConfigContract: %v", err)
	}

	header, err := ethcl.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to fetch block header: %v", err)
	}
	opts := &bind.CallOpts{
		Pending:     false,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}

	if bootstrapFlags.BatchConfigIndex <= 0 {
		log.Fatalf("Batch config index must be at least 1")
	}
	indexBig := big.NewInt(int64(bootstrapFlags.BatchConfigIndex))
	bc, err := configContract.Configs(opts, indexBig)
	if err != nil {
		log.Fatalf("Failed to fetch config at index %d: %v", bootstrapFlags.BatchConfigIndex, err)
	}

	keypers, err := configContract.GetConfigKeypers(opts, uint64(bootstrapFlags.BatchConfigIndex))
	if err != nil {
		log.Fatalf("Failed to fetch keyper set: %s", err)
	}

	ms := keyper.NewMessageSender(shmcl, signingKey)
	batchConfigMsg := keyper.NewBatchConfig(
		bc.StartBatchIndex,
		keypers,
		bc.Threshold,
		configContractAddress,
		uint64(bootstrapFlags.BatchConfigIndex),
	)

	err = ms.SendMessage(batchConfigMsg)
	if err != nil {
		log.Fatalf("Failed to send batch config message: %v", err)
	}

	batchConfigStartedMsg := keyper.NewBatchConfigStarted(uint64(bootstrapFlags.BatchConfigIndex))
	err = ms.SendMessage(batchConfigStartedMsg)
	if err != nil {
		log.Fatalf("Failed to send start message: %v", err)
	}

	log.Println("Submitted bootstrapping transaction")
	log.Printf("Config index: %d", bootstrapFlags.BatchConfigIndex)
	log.Printf("StartBatchIndex: %d", bc.StartBatchIndex)
	log.Printf("Threshold: %d", bc.Threshold)
	log.Printf("Num Keypers: %d", len(keypers))
}
