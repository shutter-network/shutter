package cmd

// This has been copied from tendermint's own init command

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/brainbot-com/shutter/shuttermint/app"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

var (
	logger  = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	rootDir = ""
	devMode = false
	index   = 0
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a config file for a Shuttermint node",
	RunE:  initFiles,
}

func init() {
	initCmd.PersistentFlags().StringVar(&rootDir, "root", "", "root directory")
	initCmd.PersistentFlags().BoolVar(&devMode, "dev", false, "turn on devmode (disables validator set changes)")
	initCmd.PersistentFlags().IntVar(&index, "index", 0, "keyper index")
	initCmd.MarkPersistentFlagRequired("root")
}

func initFiles(cmd *cobra.Command, args []string) error {
	config := cfg.DefaultConfig()
	keyper0RPCAddress := config.RPC.ListenAddress
	rpcAddress, err := adjustPort(keyper0RPCAddress, index)
	if err != nil {
		return err
	}
	config.RPC.ListenAddress = rpcAddress

	keyper0P2PAddress := config.P2P.ListenAddress
	p2pAddress, err := adjustPort(keyper0P2PAddress, index)
	if err != nil {
		return err
	}
	config.P2P.ListenAddress = p2pAddress

	config.P2P.AllowDuplicateIP = true

	config.SetRoot(rootDir)
	if err := config.ValidateBasic(); err != nil {
		return fmt.Errorf("error in config file: %v", err)
	}
	cfg.EnsureRoot(config.RootDir)

	// EnsureRoot also write the config file but with the default config. We want our own, so
	// let's overwrite it.
	cfg.WriteConfigFile(filepath.Join(rootDir, "config", "config.toml"), config)
	appState := app.NewGenesisAppState(
		[]common.Address{
			crypto.PubkeyToAddress(sandbox.GanacheKey(sandbox.NumGanacheKeys() - 1).PublicKey),
		},
		1)

	return initFilesWithConfig(config, appState)
}

func adjustPort(address string, keyperIndex int) (string, error) {
	substrings := strings.Split(address, ":")
	if len(substrings) < 2 {
		return "", fmt.Errorf("address %s does not contain port", address)
	}
	portStr := substrings[len(substrings)-1]
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		return "", fmt.Errorf("port %s is not an integer", portStr)
	}
	portIntAdjusted := portInt + keyperIndex*2
	portStrAdjusted := strconv.Itoa(portIntAdjusted)
	return strings.Join(substrings[:len(substrings)-1], ":") + ":" + portStrAdjusted, nil
}

func initFilesWithConfig(config *cfg.Config, appState app.GenesisAppState) error {
	// private validator
	privValKeyFile := config.PrivValidatorKeyFile()
	privValStateFile := config.PrivValidatorStateFile()
	var pv *privval.FilePV
	if tmos.FileExists(privValKeyFile) {
		pv = privval.LoadFilePV(privValKeyFile, privValStateFile)
		logger.Info("Found private validator", "keyFile", privValKeyFile,
			"stateFile", privValStateFile)
	} else {
		pv = privval.GenFilePV(privValKeyFile, privValStateFile)
		pv.Save()
		logger.Info("Generated private validator", "keyFile", privValKeyFile,
			"stateFile", privValStateFile)
	}

	nodeKeyFile := config.NodeKeyFile()
	if tmos.FileExists(nodeKeyFile) {
		logger.Info("Found node key", "path", nodeKeyFile)
	} else {
		if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		logger.Info("Generated node key", "path", nodeKeyFile)
	}

	// genesis file
	genFile := config.GenesisFile()
	if tmos.FileExists(genFile) {
		logger.Info("Found genesis file", "path", genFile)
	} else {
		appStateBytes, err := amino.NewCodec().MarshalJSONIndent(appState, "", "    ")
		if err != nil {
			return err
		}
		genDoc := types.GenesisDoc{
			ChainID:         fmt.Sprintf("shutter-test-chain-%v", tmrand.Str(6)),
			GenesisTime:     tmtime.Now(),
			ConsensusParams: types.DefaultConsensusParams(),
			AppState:        appStateBytes,
		}
		pubKey, err := pv.GetPubKey()
		if err != nil {
			return errors.Wrap(err, "can't get pubkey")
		}
		genDoc.Validators = []types.GenesisValidator{{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   10,
		}}

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("Generated genesis file", "path", genFile)
	}
	a := app.NewShutterApp()
	a.Gobpath = filepath.Join(config.BaseConfig.DBDir(), "shutter.gob")
	a.DevMode = devMode
	err := a.PersistToDisk()
	if err != nil {
		return err
	}

	return nil
}
