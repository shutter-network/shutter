package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type KeyperConfig struct {
	ShuttermintURL string
	SigningKey     string
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

func readKeyperConfig() (KeyperConfig, error) {
	viper.SetEnvPrefix("KEYPER")
	viper.BindEnv("ShuttermintURL")
	viper.BindEnv("SigningKey")
	viper.SetDefault("ShuttermintURL", "http://localhost:26657")

	var err error
	kc := KeyperConfig{}

	viper.AddConfigPath("$HOME/.config/shutter")
	viper.SetConfigName("keyper")
	viper.SetConfigType("toml")

	if cfgFile == "" {
		err = viper.ReadInConfig()
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
		} else {
			return kc, err // Config file was found but another error was produced
		}
	} else {
		f, err := os.Open(cfgFile)
		if err != nil {
			return kc, err
		}
		defer f.Close()

		err = viper.ReadConfig(f)
		if err != nil {
			return kc, err
		}
	}

	err = viper.Unmarshal(&kc)
	return kc, err
}

func keyperMain() {
	kc, err := readKeyperConfig()
	if err != nil {
		panic(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	privateKey, err := crypto.HexToECDSA(kc.SigningKey)
	if err != nil {
		panic(fmt.Errorf("Bad signing key: %s", err))
	}

	addr := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	log.Printf("Starting keyper version %s with signing key %s, using %s", version, addr, kc.ShuttermintURL)
	k := keyper.NewKeyper(privateKey, kc.ShuttermintURL)
	err = k.Run()
	if err != nil {
		panic(err)
	}
}
