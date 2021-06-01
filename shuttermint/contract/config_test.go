package contract

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"gotest.tools/v3/assert"
)

func makeExampleBatchConfig() BatchConfig {
	return BatchConfig{
		StartBatchIndex:  123,
		StartBlockNumber: 456,
		Keypers: []common.Address{
			common.HexToAddress("0x6ab87f620cb46764C6466e9Ca7Ac193711855D09"),
			common.HexToAddress("0xd64c1dBD939ED3A9947CDB812e706342Dc4A15FC"),
		},
		Threshold:              2,
		BatchSpan:              10,
		BatchSizeLimit:         100000,
		TransactionSizeLimit:   1000,
		TransactionGasLimit:    10000,
		FeeReceiver:            common.HexToAddress("0x1111111111111111111111111111111111111111"),
		TargetAddress:          common.HexToAddress("0x0CF414a5c990e859327E2bB19537EBf5684E58Ce"),
		TargetFunctionSelector: [4]byte{0x01, 0x02, 0x03, 0x04},
		ExecutionTimeout:       15,
	}
}

func TestBatchConfigJSONMarshaling(t *testing.T) {
	bc := makeExampleBatchConfig()

	d, err := json.MarshalIndent(bc, "", "    ")
	assert.NilError(t, err)
	bc2 := BatchConfig{}
	err = json.Unmarshal(d, &bc2)
	assert.NilError(t, err)
	assert.DeepEqual(t, bc, bc2)
}

func TestBatchConfigUnmarshaling(t *testing.T) {
	bc := makeExampleBatchConfig()
	d, err := json.MarshalIndent(bc, "", "    ")
	assert.NilError(t, err)
	fmt.Println(string(d))
	bc2 := BatchConfig{}
	replaceUnmarshal := func(t *testing.T, old, new string) {
		t.Helper()
		err = json.Unmarshal(bytes.ReplaceAll(d, []byte(old), []byte(new)), &bc2)
		assert.Assert(t, err != nil)
		fmt.Println("ERROR:", err)
	}

	t.Run("malformed fee receiver address",
		func(t *testing.T) {
			replaceUnmarshal(t, "0x1111", "0xABCD")
			replaceUnmarshal(t, "0x1111", "0x11")
		},
	)
	t.Run("malformed keyper address",
		func(t *testing.T) {
			replaceUnmarshal(t, "0x6ab87f", "0X6AB87F")
		},
	)
	t.Run("malformed target function selector",
		func(t *testing.T) {
			replaceUnmarshal(t, "0x01020304", "0x010203")
			replaceUnmarshal(t, "0x01020304", "0x0102030405")
			replaceUnmarshal(t, "0x01020304", "0xX1020304")
		},
	)
}

func TestNextBatchIndex(t *testing.T) {
	key, err := crypto.GenerateKey()
	assert.NilError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	assert.NilError(t, err)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 8000000)
	defer blockchain.Close()

	_, tx, cc, err := DeployConfigContract(
		auth,
		blockchain,
		10,
	)
	assert.NilError(t, err)

	blockchain.Commit()

	var txs []*types.Transaction

	addTx := func() {
		assert.NilError(t, err)
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
		assert.NilError(t, err)
		assert.Equal(t, uint64(1), receipt.Status)
	}

	var batchIndex uint64
	for i := uint64(0); i < 1000; i++ {
		batchIndex, err = cc.NextBatchIndex(i)
		assert.NilError(t, err)
		assert.Equal(t, uint64(0), batchIndex)
	}

	for i := uint64(1000); i < 1005; i++ {
		batchIndex, err = cc.NextBatchIndex(i)
		assert.NilError(t, err)
		assert.Equal(t, uint64(1), batchIndex)
	}

	batchIndex, err = cc.NextBatchIndex(1005)
	assert.NilError(t, err)
	assert.Equal(t, uint64(2), batchIndex)
}

func TestBatchConfig(t *testing.T) {
	bc := BatchConfig{
		StartBatchIndex:  10,
		StartBlockNumber: 500,
		BatchSpan:        5,
	}
	t.Run("BatchStartBlock", func(t *testing.T) {
		assert.Equal(t, uint64(500), bc.BatchStartBlock(10))
		assert.Equal(t, uint64(505), bc.BatchStartBlock(11))
	})
	t.Run("BatchIndex", func(t *testing.T) {
		assert.Equal(t, uint64(10), bc.BatchIndex(500))
		assert.Equal(t, uint64(10), bc.BatchIndex(504))
		assert.Equal(t, uint64(11), bc.BatchIndex(505))
	})
}
