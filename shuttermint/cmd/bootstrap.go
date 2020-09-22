package cmd

import (
	"context"
	"log"
	"math/big"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/spf13/cobra"
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
	configContract, err := contract.NewConfigContract(configContractAddress, ethcl)

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
	index := big.NewInt(int64(bootstrapFlags.BatchConfigIndex))
	bc, err := configContract.Configs(opts, index)
	if err != nil {
		log.Fatalf("Failed to fetch config at index %d: %v", bootstrapFlags.BatchConfigIndex, err)
	}

	if !bc.StartBatchIndex.IsUint64() {
		log.Fatalf("StartBatchIndex (%d) of config at index %d is too big", bc.StartBatchIndex, index)
	}
	startBatchIndex := bc.StartBatchIndex.Uint64()
	threshold := uint32(bc.Threshold.Uint64())
	if big.NewInt(int64(threshold)).Cmp(bc.Threshold) != 0 {
		log.Fatalf("Threshold (%d) of config at index %d is too big", bc.Threshold, index)
	}

	keypers, err := configContract.GetConfigKeypers(opts, index.Uint64())
	if err != nil {
		log.Fatalf("Failed to fetch keyper set: %s", err)
	}

	ms := keyper.NewMessageSender(shmcl, signingKey)
	message := keyper.NewBatchConfig(
		startBatchIndex,
		keypers,
		threshold,
	)

	err = ms.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send batch config message: %v", err)
	}

	log.Println("Submitted bootstrapping transaction")
	log.Printf("Config index: %d", index)
	log.Printf("StartBatchIndex: %d", startBatchIndex)
	log.Printf("Threshold: %d", threshold)
	log.Printf("Num Keypers: %d", len(keypers))
}
