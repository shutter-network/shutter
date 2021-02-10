// package medley provides some functions that may be useful in various parts of shutter
package medley

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const receiptPollInterval = 500 * time.Millisecond

var errAddressNotFound = errors.New("address not found")

// FindAddressIndex returns the index of the given address inside the slice of addresses or returns
// an error, if the slice does not contain the given address
func FindAddressIndex(addresses []common.Address, addr common.Address) (int, error) {
	for i, a := range addresses {
		if a == addr {
			return i, nil
		}
	}
	return -1, errAddressNotFound
}

// WaitMined waits for a transaction to be mined and returns its receipt. It's a replacement for
// bind.WaitMined which doesn't seem to work with Ganache in some cases.
func WaitMined(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(receiptPollInterval)
				continue
			}
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
		return res, fmt.Errorf("some txs have failed, the first being %s", txHashes[firstFailed])
	}

	return res, nil
}

// EnsureUniqueAddresses makes sure the slice of addresses doesn't contain duplicate addresses
func EnsureUniqueAddresses(addrs []common.Address) error {
	seen := make(map[common.Address]struct{})
	for _, a := range addrs {
		if _, ok := seen[a]; ok {
			return fmt.Errorf("duplicate address: %s", a.Hex())
		}
		seen[a] = struct{}{}
	}
	return nil
}

// DedupAddresses returns a new slice containing only unique addresses
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

// CloneWithGob clones the given object by serializing/deserializing with gob
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
