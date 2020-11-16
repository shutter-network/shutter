package cmd

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
)

// RawKeyperConfig contains raw, unvalidated configuration parameters
type RawKeyperConfig struct {
	ShuttermintURL       string
	EthereumURL          string
	SigningKey           string
	ValidatorSeed        string
	EncryptionKey        string
	ConfigContract       string
	BatcherContract      string
	KeyBroadcastContract string
	ExecutorContract     string
}

// keyperCmd represents the keyper command
var keyperCmd = &cobra.Command{
	Use:   "keyper",
	Short: "Run a shutter keyper",
	Run: func(cmd *cobra.Command, args []string) {
		keyperMain()
	},
}

func init() {
	rootCmd.AddCommand(keyperCmd)
	keyperCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func readKeyperConfig() (RawKeyperConfig, error) {
	viper.SetEnvPrefix("KEYPER")
	viper.BindEnv("ShuttermintURL")
	viper.BindEnv("EthereumURL")
	viper.BindEnv("SigningKey")
	viper.BindEnv("ValidatorSeed")
	viper.BindEnv("EncryptionKey")
	viper.BindEnv("ConfigContract")
	viper.BindEnv("BatcherContract")
	viper.BindEnv("KeyBroadcastContract")
	viper.BindEnv("ExecutorContract")

	viper.SetDefault("ShuttermintURL", "http://localhost:26657")
	viper.SetDefault("EthereumURL", "ws://localhost:8545/websocket")

	defer func() {
		if viper.ConfigFileUsed() != "" {
			log.Printf("Read config from %s", viper.ConfigFileUsed())
		}
	}()
	var err error
	rkc := RawKeyperConfig{}

	viper.AddConfigPath("$HOME/.config/shutter")
	viper.SetConfigName("keyper")
	viper.SetConfigType("toml")
	viper.SetConfigFile(cfgFile)

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found
		if cfgFile != "" {
			return rkc, err
		}
	} else if err != nil {
		return rkc, err // Config file was found but another error was produced
	}

	err = viper.Unmarshal(&rkc)
	return rkc, err
}

func ValidateKeyperConfig(r RawKeyperConfig) (keyper.KeyperConfig, error) {
	emptyConfig := keyper.KeyperConfig{}

	signingKey, err := crypto.HexToECDSA(r.SigningKey)
	if err != nil {
		return emptyConfig, fmt.Errorf("bad signing key: %w", err)
	}

	validatorSeed, err := hex.DecodeString(r.ValidatorSeed)
	if err != nil {
		return emptyConfig, fmt.Errorf("invalid validator seed: %w", err)
	}
	if len(validatorSeed) != ed25519.SeedSize {
		return emptyConfig, fmt.Errorf("invalid validator seed length %d (must be %d)", len(validatorSeed), ed25519.SeedSize)
	}
	validatorKey := ed25519.NewKeyFromSeed(validatorSeed)

	encryptionKeyECDSA, err := crypto.HexToECDSA(r.EncryptionKey)
	if err != nil {
		return emptyConfig, fmt.Errorf("bad encryption key: %w", err)
	}
	encryptionKey := ecies.ImportECDSA(encryptionKeyECDSA)

	if !keyper.IsWebsocketURL(r.EthereumURL) {
		return emptyConfig, fmt.Errorf("EthereumURL must start with ws:// or wss://")
	}

	configContractAddress := common.HexToAddress(r.ConfigContract)
	if r.ConfigContract != configContractAddress.Hex() {
		return emptyConfig, fmt.Errorf("ConfigContract must be a valid checksummed address")
	}

	batcherContractAddress := common.HexToAddress(r.BatcherContract)
	if r.BatcherContract != batcherContractAddress.Hex() {
		return emptyConfig, fmt.Errorf("BatcherContract must be a valid checksummed address")
	}

	keyBroadcastContractAddress := common.HexToAddress(r.KeyBroadcastContract)
	if r.KeyBroadcastContract != keyBroadcastContractAddress.Hex() {
		return emptyConfig, fmt.Errorf(
			"KeyBroadcastContract must be a valid checksummed address",
		)
	}

	executorContractAddress := common.HexToAddress(r.ExecutorContract)
	if r.ExecutorContract != executorContractAddress.Hex() {
		return emptyConfig, fmt.Errorf(
			"ExecutorContract must be a valid checksummed address",
		)
	}

	return keyper.KeyperConfig{
		ShuttermintURL:              r.ShuttermintURL,
		EthereumURL:                 r.EthereumURL,
		SigningKey:                  signingKey,
		ValidatorKey:                validatorKey,
		EncryptionKey:               encryptionKey,
		ConfigContractAddress:       configContractAddress,
		BatcherContractAddress:      batcherContractAddress,
		KeyBroadcastContractAddress: keyBroadcastContractAddress,
		ExecutorContractAddress:     executorContractAddress,
	}, nil
}

func keyperMain() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	rkc, err := readKeyperConfig()
	if err != nil {
		log.Fatalf("Error reading the configuration file: %s\nPlease check your configuration.", err)
	}

	kc, err := ValidateKeyperConfig(rkc)
	if err != nil {
		log.Fatalf("Error: %s\nPlease check your configuration", err)
	}

	log.Printf(
		"Starting keyper version %s with signing key %s, using %s for Shuttermint and %s for Ethereum",
		version,
		kc.Address().Hex(),
		kc.ShuttermintURL,
		kc.EthereumURL,
	)
	kpr := keyper.NewKeyper(kc)
	err = kpr.Run()
	if err != nil {
		panic(err)
	}
}
