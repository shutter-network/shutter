// Package medley provides some functions that may be useful in various parts of shutter
package medley

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	pkgErrors "github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const receiptPollInterval = 500 * time.Millisecond

var errAddressNotFound = errors.New("address not found")

// FindAddressIndex returns the index of the given address inside the slice of addresses or returns
// an error, if the slice does not contain the given address.
func FindAddressIndex(addresses []common.Address, addr common.Address) (int, error) {
	for i, a := range addresses {
		if a == addr {
			return i, nil
		}
	}
	return -1, pkgErrors.WithStack(errAddressNotFound)
}

// Sleep pauses the current goroutine for the given duration.
func Sleep(ctx context.Context, d time.Duration) {
	if d <= 0 {
		return
	}
	select {
	case <-ctx.Done():
		return
	case <-time.After(d):
	}
}

// WaitMined waits for a transaction to be mined and returns its receipt. It's a replacement for
// bind.WaitMined which doesn't seem to work with Ganache in some cases.
func WaitMined(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == ethereum.NotFound {
			Sleep(ctx, receiptPollInterval)
			continue
		}
		if err != nil {
			return nil, err
		}
		return receipt, nil
	}
}

func WaitMinedMany(ctx context.Context, client *ethclient.Client, txHashes []common.Hash) ([]*types.Receipt, error) {
	defer fmt.Print("\n")
	var res []*types.Receipt

	failedTxs := []int{}
	for i, txHash := range txHashes {
		receipt, err := WaitMined(ctx, client, txHash)
		if err != nil {
			return res, err
		}
		res = append(res, receipt)
		if receipt.Status != 1 {
			fmt.Print("X")
			failedTxs = append(failedTxs, i)
		} else {
			fmt.Print(".")
		}
	}

	if len(failedTxs) > 0 {
		firstFailed := failedTxs[0]
		return res, pkgErrors.Errorf("some txs have failed, the first being %s", txHashes[firstFailed])
	}

	return res, nil
}

// EnsureUniqueAddresses makes sure the slice of addresses doesn't contain duplicate addresses.
func EnsureUniqueAddresses(addrs []common.Address) error {
	seen := make(map[common.Address]struct{})
	for _, a := range addrs {
		if _, ok := seen[a]; ok {
			return pkgErrors.Errorf("duplicate address: %s", a.Hex())
		}
		seen[a] = struct{}{}
	}
	return nil
}

// DedupAddresses returns a new slice containing only unique addresses.
func DedupAddresses(addrs []common.Address) []common.Address {
	var res []common.Address
	seen := make(map[common.Address]struct{})

	for _, a := range addrs {
		if _, ok := seen[a]; ok {
			continue
		}
		seen[a] = struct{}{}
		res = append(res, a)
	}

	return res
}

// CloneWithGob clones the given object by serializing/deserializing with gob.
func CloneWithGob(src, dst interface{}) {
	buff := bytes.Buffer{}
	err := gob.NewEncoder(&buff).Encode(src)
	if err != nil {
		panic(err)
	}
	err = gob.NewDecoder(&buff).Decode(dst)
	if err != nil {
		panic(err)
	}
}

func normName(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, "-", "_"))
}

func argumentFromEnvironmentVariable(cmd *cobra.Command, f *pflag.Flag) error {
	if f.Changed {
		return nil
	}

	candidates := []string{}
	if cmd.Parent() != nil {
		candidates = append(candidates, normName(fmt.Sprintf("SHUTTER_%s_%s", cmd.Name(), f.Name)))
	}
	candidates = append(candidates, normName(fmt.Sprintf("SHUTTER_%s", f.Name)))

	for _, envvar := range candidates {
		val, ok := os.LookupEnv(envvar)
		if !ok {
			continue
		}
		err := cmd.Flags().Set(f.Name, val)
		if err != nil {
			return pkgErrors.Wrapf(err, "argument from environment variable %s", envvar)
		}
		return nil
	}
	return nil
}

// BindFlags automatically sets options to command line flags from environment variables.
func BindFlags(cmd *cobra.Command) error {
	var err error
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err != nil {
			return
		}
		err = argumentFromEnvironmentVariable(cmd, f)
	})
	return err
}

// ShowHelpAndExit shows the commands help message and exits the program with status 1.
func ShowHelpAndExit(cmd *cobra.Command, args []string) {
	_ = args
	_ = cmd.Help()
	os.Exit(1)
}
