package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/shutter-network/shutter/shuttermint/cmd/shversion"
	"github.com/shutter-network/shutter/shuttermint/keyper"
	"github.com/shutter-network/shutter/shuttermint/keyper/gaspricer"
)

// keyperCmd represents the keyper command.
var keyperCmd = &cobra.Command{
	Use:   "keyper",
	Short: "Run a Shutter keyper node",
	Long: `This command runs a keyper node. It will connect to both an Ethereum and a
Shuttermint node which have to be started separately in advance.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return keyperMain()
	},
}

func init() {
	keyperCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func readKeyperConfig() (keyper.Config, error) {
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
	viper.BindEnv("DepositContract")
	viper.BindEnv("KeyperSlasher")
	viper.BindEnv("MainChainFollowDistance")
	viper.BindEnv("ExecutionStaggering")
	viper.BindEnv("DKGPhaseLength")

	viper.SetDefault("ShuttermintURL", "http://localhost:26657")

	defer func() {
		if viper.ConfigFileUsed() != "" {
			log.Printf("Read config from %s", viper.ConfigFileUsed())
		}
	}()
	var err error
	config := keyper.Config{}

	viper.AddConfigPath("$HOME/.config/shutter")
	viper.SetConfigName("keyper")
	viper.SetConfigType("toml")
	viper.SetConfigFile(cfgFile)

	err = viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found
		if cfgFile != "" {
			return config, err
		}
	} else if err != nil {
		return config, err // Config file was found but another error was produced
	}
	err = config.Unmarshal(viper.GetViper())

	if err != nil {
		return config, err
	}

	if !filepath.IsAbs(config.DBDir) {
		r := filepath.Dir(viper.ConfigFileUsed())
		dbdir, err := filepath.Abs(filepath.Join(r, config.DBDir))
		if err != nil {
			return config, err
		}
		config.DBDir = dbdir
	}

	if !keyper.IsWebsocketURL(config.EthereumURL) {
		return config, errors.Errorf("field EthereumURL must start with ws:// or wss://")
	}

	return config, err
}

func keyperMain() error {
	kc, err := readKeyperConfig()
	if err != nil {
		return errors.WithMessage(err, "Please check your configuration")
	}
	err = gaspricer.SetMultiplier(kc.GasPriceMultiplier)
	if err != nil {
		return errors.WithMessage(err, "Please check your configuration")
	}

	log.Printf(
		"Starting keyper version %s with signing key %s, using %s for Shuttermint and %s for Ethereum",
		shversion.Version(),
		kc.Address().Hex(),
		kc.ShuttermintURL,
		kc.EthereumURL,
	)
	kpr := keyper.NewKeyper(kc)
	err = kpr.LoadState()
	if err != nil {
		return errors.WithMessage(err, "LoadState")
	}
	log.Printf("Loaded state with %d actions, %s", len(kpr.State.Actions), kpr.ShortInfo())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-termChan
		log.Printf("Received %s signal, shutting down", sig)
		cancel()
	}()

	err = kpr.Run(ctx)
	if err == context.Canceled {
		log.Printf("Bye.")
		return nil
	}
	return err
}
