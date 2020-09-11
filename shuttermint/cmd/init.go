package cmd

// This has been copied from tendermint's own init command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

var (
	logger  = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	rootDir = ""
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize shuttermint",
	RunE:    initFiles,
	PreRunE: initCheckRootDir,
}

func initCheckRootDir(cmd *cobra.Command, args []string) error {
	if rootDir == "" {
		return fmt.Errorf("root argument missing")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVar(&rootDir, "root", "", "root directory")
}

func initFiles(cmd *cobra.Command, args []string) error {
	config := cfg.DefaultConfig()
	config.TxIndex.IndexAllKeys = true
	config.SetRoot(rootDir)
	if err := config.ValidateBasic(); err != nil {
		return fmt.Errorf("error in config file: %v", err)
	}
	cfg.EnsureRoot(config.RootDir)

	// EnsureRoot also write the config file but with the default config. We want our own, so
	// let's overwrite it.
	cfg.WriteConfigFile(filepath.Join(rootDir, "config", "config.toml"), config)

	return initFilesWithConfig(config)
}

func initFilesWithConfig(config *cfg.Config) error {
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
		genDoc := types.GenesisDoc{
			ChainID:         fmt.Sprintf("shutter-test-chain-%v", tmrand.Str(6)),
			GenesisTime:     tmtime.Now(),
			ConsensusParams: types.DefaultConsensusParams(),
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

	return nil
}
