package keyper

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	ganachePort               = 8545
	batchSpan                 = 5
	threshold                 = 2
	configChangeHeadsUpBlocks = 10
)

var (
	zeroDecryptionKey            *ecdsa.PrivateKey
	emptyBatchHash               common.Hash
	emptyDecryptionSignerIndices []uint64
	emptyDecryptionSignatures    [][]byte
)

var (
	cl                     *ethclient.Client
	rpcCl                  *rpc.Client
	cc                     *ContractCaller
	signer                 types.EIP155Signer
	batcherContractAddress common.Address
)

var (
	keyperKeys      []*ecdsa.PrivateKey
	keyperAddresses []common.Address
	deployKey       *ecdsa.PrivateKey
	deployAddress   common.Address
)

var batchConfig contract.BatchConfig

func init() {
	for i := 0; i < 3; i++ {
		keyperKeys = append(keyperKeys, sandbox.GanacheKey(i))
		keyperAddresses = append(keyperAddresses, crypto.PubkeyToAddress(keyperKeys[i].PublicKey))
	}
	deployKey = sandbox.GanacheKey(9)
	deployAddress = crypto.PubkeyToAddress(deployKey.PublicKey)

	emptyBatchHash = ComputeBatchHash([][]byte{})
	zeroKeyBytes := make([]byte, 32)
	copy(zeroKeyBytes[31:], []byte{1})
	zeroKey, err := crypto.ToECDSA(zeroKeyBytes)
	if err != nil {
		panic("failed to create zero decryption key")
	}
	zeroDecryptionKey = zeroKey
}

func runGanache(t *testing.T) {
	ganachePath, err := exec.LookPath("ganache-cli")
	require.Nil(t, err)
	ganacheCmd := exec.Command(ganachePath, "-d", "-p", strconv.Itoa(ganachePort))

	stdout, err := ganacheCmd.StdoutPipe()
	require.Nil(t, err)

	err = ganacheCmd.Start()
	require.Nil(t, err)
	t.Cleanup(func() {
		ganacheCmd.Process.Signal(os.Kill)
		ganacheCmd.Wait()
	})

	// wait for ganache to start
	stdoutBuf := bufio.NewReader(stdout)
	for {
		line, err := stdoutBuf.ReadString('\n')
		require.Nil(t, err)
		if strings.HasPrefix(line, "Listening on") {
			break
		}
	}

	rpcURL := "ws://127.0.0.1:" + strconv.Itoa(ganachePort)
	cl, err = ethclient.Dial(rpcURL)
	require.Nil(t, err)
	t.Cleanup(cl.Close)
	rpcCl, err = rpc.Dial(rpcURL)
	require.Nil(t, err)
	t.Cleanup(rpcCl.Close)

	chainID, err := cl.NetworkID(context.Background())
	require.Nil(t, err)
	signer = types.NewEIP155Signer(chainID)
}

func deployContracts(t *testing.T) {
	auth := bind.NewKeyedTransactor(deployKey)
	nonce, err := cl.PendingNonceAt(context.Background(), deployAddress)
	require.Nil(t, err)
	auth.Nonce = new(big.Int).SetUint64(nonce)

	configAddress, tx1, configContract, err := contract.DeployConfigContract(auth, cl, configChangeHeadsUpBlocks)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	_, tx2, keyBroadcastContract, err := contract.DeployKeyBroadcastContract(auth, cl, configAddress)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	feeAddress, tx3, _, err := contract.DeployFeeBankContract(auth, cl)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	batcherAddress, tx4, batcherContract, err := contract.DeployBatcherContract(auth, cl, configAddress, feeAddress)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)
	batcherContractAddress = batcherAddress

	_, tx5, executorContract, err := contract.DeployExecutorContract(auth, cl, configAddress, batcherAddress)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	for _, tx := range []*types.Transaction{tx1, tx2, tx3, tx4, tx5} {
		receipt, err := bind.WaitMined(context.Background(), cl, tx)
		require.Nil(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	}

	contractCaller := NewContractCaller(
		cl,
		keyperKeys[0],
		configContract,
		keyBroadcastContract,
		batcherContract,
		executorContract,
	)
	cc = &contractCaller

	for i, key := range keyperKeys {
		sig, err := ComputeDecryptionSignature(key, batcherAddress, emptyBatchHash, zeroDecryptionKey, emptyBatchHash)
		require.Nil(t, err)
		emptyDecryptionSignerIndices = append(emptyDecryptionSignerIndices, uint64(i))
		emptyDecryptionSignatures = append(emptyDecryptionSignatures, sig)
	}
}

func scheduleConfig(t *testing.T) {
	auth := bind.NewKeyedTransactor(deployKey)
	nonce, err := cl.PendingNonceAt(context.Background(), deployAddress)
	require.Nil(t, err)
	auth.Nonce = new(big.Int).SetUint64(nonce)

	tx1, err := cc.ConfigContract.NextConfigSetBatchSpan(auth, batchSpan)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	tx2, err := cc.ConfigContract.NextConfigAddKeypers(auth, keyperAddresses)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	tx3, err := cc.ConfigContract.NextConfigSetThreshold(auth, threshold)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	tx4, err := cc.ConfigContract.NextConfigSetExecutionTimeout(auth, batchSpan)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	tx5, err := cc.ConfigContract.NextConfigSetTransactionSizeLimit(auth, 100)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	tx6, err := cc.ConfigContract.NextConfigSetBatchSizeLimit(auth, 100*100)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	header, err := cl.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	startBlockNumber := header.Number.Uint64() + configChangeHeadsUpBlocks + 5
	tx7, err := cc.ConfigContract.NextConfigSetStartBlockNumber(auth, startBlockNumber)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)

	for _, tx := range []*types.Transaction{tx1, tx2, tx3, tx4, tx5, tx6, tx7} {
		waitSuccessful(t, tx)
	}

	tx8, err := cc.ConfigContract.ScheduleNextConfig(auth)
	require.Nil(t, err)
	auth.Nonce.SetUint64(auth.Nonce.Uint64() + 1)
	waitSuccessful(t, tx8)

	config, err := cc.ConfigContract.GetConfig(nil, 1)
	require.Nil(t, err)
	batchConfig = config
}

func waitSuccessful(t *testing.T, tx *types.Transaction) {
	receipt, err := bind.WaitMined(context.Background(), cl, tx)
	require.Nil(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
}

func mineBlock(t *testing.T) {
	var result hexutil.Uint64
	err := rpcCl.CallContext(context.Background(), &result, "evm_mine")
	require.Nil(t, err)
}

func mineUntilBlock(t *testing.T, blockNumber uint64) {
	for {
		header, err := cl.HeaderByNumber(context.Background(), nil)
		require.Nil(t, err)
		if header.Number.Uint64() >= blockNumber {
			return
		}
		mineBlock(t)
	}
}

func snapshot(t *testing.T) uint64 {
	var id hexutil.Uint64
	err := rpcCl.CallContext(context.Background(), &id, "evm_snapshot")
	require.Nil(t, err)
	return uint64(id)
}

func revert(t *testing.T, snapshotID uint64) uint64 {
	var result bool
	err := rpcCl.CallContext(context.Background(), &result, "evm_revert", hexutil.EncodeUint64(snapshotID))
	require.Nil(t, err)
	require.True(t, result)
	return snapshot(t) // reverting removes the snapshot, so recreate it
}

func TestExecutor(t *testing.T) {
	runGanache(t)
	deployContracts(t)
	scheduleConfig(t)

	snapshotID := snapshot(t)
	subTestExecutePlain(t)

	snapshotID = revert(t, snapshotID)
	subTestExecuteCipher(t)

	snapshotID = revert(t, snapshotID)
	subTestSkipCipher(t)

	snapshotID = revert(t, snapshotID)
	subTestFastForwardWait(t)

	snapshotID = revert(t, snapshotID)
	subTestFastForwardNoWait(t)

	snapshotID = revert(t, snapshotID)
	subTestExecuteFromChannel(t)
}

func subTestExecutePlain(t *testing.T) {
	cipherExecutionParams := make(chan CipherExecutionParams)
	ex := Executor{
		ctx:                   context.Background(),
		client:                cl,
		cc:                    cc,
		cipherExecutionParams: cipherExecutionParams,
	}
	auth := bind.NewKeyedTransactor(keyperKeys[0])

	// submit some plain txs
	mineUntilBlock(t, batchConfig.StartBlockNumber)
	for i := 0; i < 2; i++ {
		tx, err := ex.cc.BatcherContract.AddTransaction(auth, 0, contract.TransactionTypePlain, []byte{0, 1, 2})
		require.Nil(t, err)
		waitSuccessful(t, tx)
	}

	// execute cipher batch so that plain batch can be executed
	mineUntilBlock(t, batchConfig.StartBlockNumber+batchConfig.BatchSpan)
	tx, err := ex.cc.ExecutorContract.ExecuteCipherBatch(
		auth,
		emptyBatchHash,
		[][]byte{},
		0,
	)
	require.Nil(t, err)
	waitSuccessful(t, tx)

	// execute plain batch
	numHalfStepsBefore, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, uint64(1), numHalfStepsBefore%2)
	batchParams, err := contract.MakeBatchParams(&batchConfig, 0)
	require.Nil(t, err)
	err = ex.executePlainHalfStep(batchParams)
	require.Nil(t, err)
	numHalfStepsAfter, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, numHalfStepsBefore+1, numHalfStepsAfter)
}

func subTestExecuteCipher(t *testing.T) {
	cipherExecutionParams := make(chan CipherExecutionParams)
	ex := Executor{
		ctx:                   context.Background(),
		client:                cl,
		cc:                    cc,
		cipherExecutionParams: cipherExecutionParams,
	}
	auth := bind.NewKeyedTransactor(keyperKeys[0])

	// submit some cipher txs
	mineUntilBlock(t, batchConfig.StartBlockNumber)
	cipherTxs := [][]byte{}
	decryptedTxs := [][]byte{}
	for i := 0; i < 2; i++ {
		dataCipher := make([]byte, 3)
		dataCipher[0] = byte(i)
		cipherTxs = append(cipherTxs, dataCipher)

		dataDecrypted := make([]byte, 3)
		dataDecrypted[1] = byte(i)
		decryptedTxs = append(decryptedTxs, dataDecrypted)

		tx, err := ex.cc.BatcherContract.AddTransaction(auth, 0, contract.TransactionTypeCipher, dataCipher)
		require.Nil(t, err)
		waitSuccessful(t, tx)
	}
	cipherBatchHash := ComputeBatchHash(cipherTxs)
	decryptedBatchHash := ComputeBatchHash(decryptedTxs)

	// compute signatures
	decryptionSignerIndices := []uint64{}
	decryptionSignatures := [][]byte{}
	for i := 0; i < threshold; i++ {
		decryptionSignerIndices = append(decryptionSignerIndices, uint64(i))
		sig, err := ComputeDecryptionSignature(
			keyperKeys[i],
			batcherContractAddress,
			cipherBatchHash,
			zeroDecryptionKey,
			decryptedBatchHash,
		)
		require.Nil(t, err)
		decryptionSignatures = append(decryptionSignatures, sig)
	}

	cipherParams := CipherExecutionParams{
		BatchIndex:              0,
		CipherBatchHash:         cipherBatchHash,
		DecryptionKey:           zeroDecryptionKey,
		DecryptedTxs:            decryptedTxs,
		DecryptionSignerIndices: decryptionSignerIndices,
		DecryptionSignatures:    decryptionSignatures,
	}
	batchParams, err := contract.MakeBatchParams(&batchConfig, 0)
	require.Nil(t, err)

	// execute cipher batch
	mineUntilBlock(t, batchConfig.StartBlockNumber+batchConfig.BatchSpan)
	numHalfStepsBefore, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, uint64(0), numHalfStepsBefore%2)
	err = ex.executeCipherHalfStep(batchParams, cipherParams)
	require.Nil(t, err)
	numHalfStepsAfter, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, numHalfStepsBefore+1, numHalfStepsAfter)
}

func subTestSkipCipher(t *testing.T) {
	cipherExecutionParams := make(chan CipherExecutionParams)
	ex := Executor{
		ctx:                   context.Background(),
		client:                cl,
		cc:                    cc,
		cipherExecutionParams: cipherExecutionParams,
	}
	auth := bind.NewKeyedTransactor(keyperKeys[0])

	// submit some cipher txs
	mineUntilBlock(t, batchConfig.StartBlockNumber)
	for i := 0; i < 2; i++ {
		dataCipher := make([]byte, 3)
		dataCipher[0] = byte(i)

		tx, err := ex.cc.BatcherContract.AddTransaction(auth, 0, contract.TransactionTypeCipher, dataCipher)
		require.Nil(t, err)
		waitSuccessful(t, tx)
	}

	// skip cipher batch
	mineUntilBlock(t, batchConfig.StartBlockNumber+batchConfig.BatchSpan+batchConfig.ExecutionTimeout)
	numHalfStepsBefore, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, uint64(0), numHalfStepsBefore%2)
	batchParams, err := contract.MakeBatchParams(&batchConfig, 0)
	require.Nil(t, err)
	err = ex.skipCipherHalfStep(batchParams)
	require.Nil(t, err)
	numHalfStepsAfter, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, numHalfStepsBefore+1, numHalfStepsAfter)
}

func subTestFastForwardNoWait(t *testing.T) {
	cipherExecutionParams := make(chan CipherExecutionParams)
	ex := Executor{
		ctx:                   context.Background(),
		client:                cl,
		cc:                    cc,
		cipherExecutionParams: cipherExecutionParams,
	}

	mineUntilBlock(
		t,
		batchConfig.StartBlockNumber+
			batchConfig.BatchSpan*3+
			batchConfig.ExecutionTimeout+uint64(len(keyperAddresses))*kickOffBlockStagger,
	)

	numHalfStepsBefore, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, uint64(0), numHalfStepsBefore%2)

	err = ex.fastForward(10, false)
	require.Nil(t, err)
	numHalfStepsAfter, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	// Due to the keyper staggering, the number of executed half steps depends on the keyper's
	// position in the priority queue which we don't know.
	require.GreaterOrEqual(t, numHalfStepsAfter, uint64(12))
	require.LessOrEqual(t, numHalfStepsAfter, uint64(20))
}

func subTestFastForwardWait(t *testing.T) {
	cipherExecutionParams := make(chan CipherExecutionParams)
	ex := Executor{
		ctx:                   context.Background(),
		client:                cl,
		cc:                    cc,
		cipherExecutionParams: cipherExecutionParams,
	}

	numHalfStepsBefore, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, uint64(0), numHalfStepsBefore%2)

	go func() {
		time.Sleep(1 * time.Second)
		mineUntilBlock(
			t,
			batchConfig.StartBlockNumber+
				batchConfig.BatchSpan*3+
				batchConfig.ExecutionTimeout+uint64(len(keyperAddresses))*kickOffBlockStagger,
		)
	}()

	err = ex.fastForward(3, true) // this should block until enough blocks have been mined
	require.Nil(t, err)
	numHalfStepsAfter, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, uint64(6), numHalfStepsAfter)
}

func subTestExecuteFromChannel(t *testing.T) {
	cipherExecutionParams := make(chan CipherExecutionParams)
	executorContext, cancelExecutor := context.WithCancel(context.Background())
	defer cancelExecutor()
	ex := Executor{
		ctx:                   executorContext,
		client:                cl,
		cc:                    cc,
		cipherExecutionParams: cipherExecutionParams,
	}

	numHalfStepsBefore, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(nil)
	require.Nil(t, err)
	require.Equal(t, uint64(0), numHalfStepsBefore%2)

	params := CipherExecutionParams{
		BatchIndex:              1, // we only send the params for batch #1, so #0 should be skipped automatically
		CipherBatchHash:         emptyBatchHash,
		DecryptionKey:           zeroDecryptionKey,
		DecryptedTxs:            [][]byte{},
		DecryptionSignerIndices: emptyDecryptionSignerIndices,
		DecryptionSignatures:    emptyDecryptionSignatures,
	}

	batchExecutedEvents := make(chan *contract.ExecutorContractBatchExecuted)
	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   nil,
	}
	batchExecutedSub, err := cc.ExecutorContract.WatchBatchExecuted(watchOpts, batchExecutedEvents)
	require.Nil(t, err)
	defer batchExecutedSub.Unsubscribe()

	// mine far enough so that batch 0 can be skipped and batch 1 can be executed
	staggerOffset := uint64(len(keyperAddresses)) * kickOffBlockStagger
	mineUntilBlock(t, batchConfig.StartBlockNumber+batchConfig.BatchSpan+batchConfig.ExecutionTimeout+staggerOffset)
	mineUntilBlock(t, batchConfig.StartBlockNumber+2*batchConfig.BatchSpan+staggerOffset)

	// Sending params should skip batch 0 and execute batch 1, so that the final number of
	// executed half steps is 4.
	go ex.Run()
	cipherExecutionParams <- params
	for {
		select {
		case err := <-batchExecutedSub.Err():
			require.Nil(t, err)
		case event := <-batchExecutedEvents:
			log.Printf("asfd %+v", event)
			if event.NumExecutionHalfSteps >= 4 {
				return
			}
		case <-time.After(5 * time.Second):
			require.FailNow(t, "timeout")
		}
	}
}
