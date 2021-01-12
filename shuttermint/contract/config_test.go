package contract

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestNextBatchIndex(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.Nil(t, err)
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 8000000)
	defer blockchain.Close()

	_, tx, cc, err := DeployConfigContract(
		auth,
		blockchain,
		10,
	)
	require.Nil(t, err)

	blockchain.Commit()

	var txs []*types.Transaction

	addTx := func() {
		require.Nil(t, err)
		txs = append(txs, tx)
	}

	tx, err = cc.NextConfigSetStartBlockNumber(auth, uint64(1000))
	addTx()

	tx, err = cc.NextConfigSetStartBatchIndex(auth, uint64(0))
	addTx()

	tx, err = cc.NextConfigSetBatchSpan(auth, 5)
	addTx()

	tx, err = cc.NextConfigAddKeypers(auth, []common.Address{common.BytesToAddress([]byte("foo"))})
	addTx()

	tx, err = cc.ScheduleNextConfig(auth)
	addTx()

	blockchain.Commit()

	for _, tx := range txs {
		var receipt *types.Receipt
		receipt, err = blockchain.TransactionReceipt(context.Background(), tx.Hash())
		require.Nil(t, err)
		require.Equal(t, uint64(1), receipt.Status)
	}

	var batchIndex uint64
	for i := uint64(0); i < 1000; i++ {
		batchIndex, err = cc.NextBatchIndex(i)
		require.Nil(t, err)
		require.Equal(t, uint64(0), batchIndex)
	}

	for i := uint64(1000); i < 1005; i++ {
		batchIndex, err = cc.NextBatchIndex(i)
		require.Nil(t, err)
		require.Equal(t, uint64(1), batchIndex)
	}

	batchIndex, err = cc.NextBatchIndex(1005)
	require.Nil(t, err)
	require.Equal(t, uint64(2), batchIndex)
}

func TestBatchConfig(t *testing.T) {
	bc := BatchConfig{
		StartBatchIndex:  10,
		StartBlockNumber: 500,
		BatchSpan:        5,
	}
	t.Run("BatchStartBlock", func(t *testing.T) {
		require.Equal(t, uint64(500), bc.BatchStartBlock(10))
		require.Equal(t, uint64(505), bc.BatchStartBlock(11))
	})
	t.Run("BatchIndex", func(t *testing.T) {
		require.Equal(t, uint64(10), bc.BatchIndex(500))
		require.Equal(t, uint64(10), bc.BatchIndex(504))
		require.Equal(t, uint64(11), bc.BatchIndex(505))
	})
}
