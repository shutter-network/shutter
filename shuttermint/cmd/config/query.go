package config

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/contract"
)

const (
	nextConfigIndexPlaceholder = "next"
	lastConfigIndexPlaceholder = "last"
	allConfigsIndexPlaceholder = "all"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Download a batch config and print it as JSON",
	Args:  cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		switch queryFlags.Index {
		case nextConfigIndexPlaceholder, lastConfigIndexPlaceholder, allConfigsIndexPlaceholder:
			return nil
		default:
			var err error
			queryFlags.IndexValue, err = strconv.ParseUint(queryFlags.Index, 10, 64)
			if err != nil {
				return errors.Wrap(err, "--index argument not valid")
			}
			return err
		}
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		err := processConfigFlags(ctx)
		if err != nil {
			return err
		}
		return query(ctx)
	},
}

var queryFlags struct {
	Index      string // the --index argument
	IndexValue uint64 // --index argument if user passed an integer
}

func init() {
	placeholders := []string{}
	for _, p := range []string{nextConfigIndexPlaceholder, lastConfigIndexPlaceholder, allConfigsIndexPlaceholder} {
		placeholders = append(placeholders, strconv.Quote(p))
	}
	queryCmd.PersistentFlags().StringVarP(
		&queryFlags.Index,
		"index",
		"i",
		allConfigsIndexPlaceholder,
		fmt.Sprintf("the index of the config to query or %s", strings.Join(placeholders, ", ")),
	)
}

func printjson(d interface{}) error {
	s, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return err
	}
	os.Stdout.Write(s)
	os.Stdout.Write([]byte{'\n'})
	return nil
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

	numConfigs, err := configContract.NumConfigs(callOpts)
	if err != nil {
		return errors.Wrapf(err, "failed to call config contract")
	}

	var index uint64

	switch queryFlags.Index {
	case nextConfigIndexPlaceholder:
		config, err = configContract.GetNextConfig(callOpts)
		if err != nil {
			return err
		}
		return printjson(config)
	case allConfigsIndexPlaceholder:
		configs := []contract.BatchConfig{}
		for index = 0; index < numConfigs; index++ {
			config, err = configContract.GetConfigByIndex(callOpts, index)
			if err != nil {
				return err
			}
			configs = append(configs, config)
		}
		return printjson(configs)
	case lastConfigIndexPlaceholder:
		if numConfigs == 0 {
			return errors.Errorf("no configs scheduled in contract")
		}
		index = numConfigs - 1
	default:
		index = queryFlags.IndexValue
	}

	config, err = configContract.GetConfigByIndex(callOpts, index)
	if err != nil {
		return err
	}
	return printjson(config)
}
