package config

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	nextConfigIndexPlaceholder = "next"
	lastConfigIndexPlaceholder = "last"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Download a config and print it as JSON",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		sandbox.ExitIfError(processConfigFlags(ctx))
		if flag, err := validateQueryFlags(); err != nil {
			sandbox.ExitIfError(errors.Wrapf(err, "invalid value for flag %s", flag))
		}
		sandbox.ExitIfError(query(ctx))
	},
}

var queryFlags struct {
	Index string
}

func init() {
	queryCmd.PersistentFlags().StringVarP(
		&queryFlags.Index,
		"index",
		"i",
		"last",
		"the index of the config to query or `last` or `next` (for the next config)",
	)
}

func validateQueryFlags() (string, error) {
	if queryFlags.Index != nextConfigIndexPlaceholder && queryFlags.Index != lastConfigIndexPlaceholder {
		if _, err := strconv.ParseInt(queryFlags.Index, 10, 64); err != nil {
			return "index", errors.Wrapf(err, "not a valid index or `next`")
		}
	}

	return "", nil
}

func query(ctx context.Context) error {
	var err error
	var config contract.BatchConfig

	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed to query block number")
	}
	callOpts := &bind.CallOpts{
		Pending:     false,
		BlockNumber: new(big.Int).SetUint64(blockNumber),
		Context:     ctx,
	}

	if queryFlags.Index == nextConfigIndexPlaceholder {
		config, err = configContract.GetNextConfig(callOpts)
	} else {
		var index uint64
		if queryFlags.Index == lastConfigIndexPlaceholder {
			numConfigs, err := configContract.NumConfigs(callOpts)
			if err != nil {
				return errors.Wrapf(err, "failed to call config contract")
			}
			if numConfigs == 0 {
				return fmt.Errorf("no configs scheduled in contract")
			}
			index = numConfigs - 1
		} else {
			index, err = strconv.ParseUint(queryFlags.Index, 10, 64)
			if err != nil {
				return err // should already be catched during argument validation
			}
		}

		config, err = configContract.GetConfigByIndex(callOpts, index)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}
	configJSON := sandbox.ConfigToJSON(&config)
	s, err := json.MarshalIndent(configJSON, "", "    ")
	if err != nil {
		return err
	}
	os.Stdout.Write(s)
	os.Stdout.Write([]byte{'\n'})

	return nil
}
