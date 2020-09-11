package cmd

import (
	"fmt"
	stdlog "log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"

	"github.com/brainbot-com/shutter/shuttermint/app"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the shutter tendermint app",
	Long:  `The run commands run the shutter tendermint app`,
	Run: func(cmd *cobra.Command, args []string) {
		runMain()
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			return fmt.Errorf("--config is required")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (required)")
}

func runMain() {
	stdlog.SetFlags(stdlog.LstdFlags | stdlog.Lshortfile | stdlog.Lmicroseconds)
	stdlog.Printf("Starting shuttermint version %s", version)

	shapp := app.NewShutterApp()

	node, err := newTendermint(shapp, cfgFile)
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
	<-c
	fmt.Println("Got signal. Exiting.")
	// Previously we had an os.Exit(0) call here, but now we do wait until the defer function
	// above is done
}

func newTendermint(shapp abcitypes.Application, configFile string) (*nm.Node, error) {
	// read config
	config := cfg.DefaultConfig()
	config.RootDir = filepath.Dir(filepath.Dir(configFile))
	config.SetRoot(config.RootDir)
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("viper failed to read config file: %w", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("viper failed to unmarshal config: %w", err)
	}
	if err := config.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("config is invalid: %w", err)
	}

	// create logger
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	var err error
	logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel())
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}

	// read private validator
	pv := privval.LoadFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	// read node key
	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, fmt.Errorf("failed to load node's key: %w", err)
	}

	// create node
	node, err := nm.NewNode(
		config,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(shapp),
		nm.DefaultGenesisDocProviderFunc(config),
		nm.DefaultDBProvider,
		nm.DefaultMetricsProvider(config.Instrumentation),
		logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create new Tendermint node: %w", err)
	}

	return node, nil
}
