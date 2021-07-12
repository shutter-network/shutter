package cmd

import (
	"fmt"
	stdlog "log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cfg "github.com/tendermint/tendermint/config"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"

	"github.com/shutter-network/shutter/shuttermint/app"
	"github.com/shutter-network/shutter/shuttermint/cmd/shversion"
)

var chainCmd = &cobra.Command{
	Use:   "chain",
	Short: "Run a node for Shutter's Tendermint chain",
	Long:  `This command runs a node that will connect to Shutter's Tendermint chain.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		chainMain()
	},
}

func init() {
	chainCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (required)")
	chainCmd.MarkPersistentFlagRequired("config")
}

func chainMain() {
	stdlog.Printf("Starting shuttermint version %s", shversion.Version())

	node, err := newTendermint(cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	err = node.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = node.Stop()
		if err != nil {
			panic(err)
		}
		node.Wait()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c
	stdlog.Printf("Got signal '%s'. Exiting.", sig)
	// Previously we had an os.Exit(0) call here, but now we do wait until the defer function
	// above is done
}

func newTendermint(configFile string) (*nm.Node, error) {
	// read config
	config := cfg.DefaultConfig()
	config.RootDir = filepath.Dir(filepath.Dir(configFile))
	config.SetRoot(config.RootDir)
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper failed to read config file")
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "viper failed to unmarshal config")
	}
	if err := config.ValidateBasic(); err != nil {
		return nil, errors.Wrap(err, "config is invalid")
	}

	// create logger
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	var err error
	logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level")
	}

	shapp, err := app.LoadShutterAppFromFile(
		filepath.Join(config.BaseConfig.DBDir(), "shutter.gob"))
	if err != nil {
		return nil, err
	}

	// read private validator
	pv := privval.LoadFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	// read node key
	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, errors.Wrap(err, "failed to load node's key")
	}

	// create node
	node, err := nm.NewNode(
		config,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(&shapp),
		nm.DefaultGenesisDocProviderFunc(config),
		nm.DefaultDBProvider,
		nm.DefaultMetricsProvider(config.Instrumentation),
		logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new Tendermint node")
	}

	return node, nil
}
