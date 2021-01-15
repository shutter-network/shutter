// Deploy the ConfigContract and KeyBroadcastContract to ganache. This uses a hard-coded private
// key that available when ganache is started with the -d flag.

package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

const (
	startBlockNumberOffset           = 30
	defaultGasLimit                  = 5000000
	defaultConfigChangeHeadsUpBlocks = 20
	ganacheKeyIdx                    = 9
	numKeypers                       = 3
	threshold                        = 2
	transactionSizeLimit             = 100
	batchSizeLimit                   = 100 * 100
	baseGasLimit                     = 21000
	fundAmount                       = 1000000000000000000 // 1 eth
	dialDefaultTimeout               = 5 * time.Second
	getconfigDefaultTimeout          = 10 * time.Second
	deployDefaultTimeout             = 30 * time.Second
	scheduleDefaultTimeout           = 60 * time.Second
	erc1820DeploymentAccountHex      = "0xa990077c3205cbDf861e17Fa532eeB069cE9fF96"
	erc1820DeploymentTransactionHex  = "0xf90a388085174876e800830c35008080b909e5608060405234801561001057600080fd5b506109c5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c010000000000000000000000000000000000000000000000000000000090048063a41e7d5111610078578063a41e7d51146101d4578063aabbb8ca1461020a578063b705676514610236578063f712f3e814610280576100a5565b806329965a1d146100aa5780633d584063146100e25780635df8122f1461012457806365ba36c114610152575b600080fd5b6100e0600480360360608110156100c057600080fd5b50600160a060020a038135811691602081013591604090910135166102b6565b005b610108600480360360208110156100f857600080fd5b5035600160a060020a0316610570565b60408051600160a060020a039092168252519081900360200190f35b6100e06004803603604081101561013a57600080fd5b50600160a060020a03813581169160200135166105bc565b6101c26004803603602081101561016857600080fd5b81019060208101813564010000000081111561018357600080fd5b82018360208201111561019557600080fd5b803590602001918460018302840111640100000000831117156101b757600080fd5b5090925090506106b3565b60408051918252519081900360200190f35b6100e0600480360360408110156101ea57600080fd5b508035600160a060020a03169060200135600160e060020a0319166106ee565b6101086004803603604081101561022057600080fd5b50600160a060020a038135169060200135610778565b61026c6004803603604081101561024c57600080fd5b508035600160a060020a03169060200135600160e060020a0319166107ef565b604080519115158252519081900360200190f35b61026c6004803603604081101561029657600080fd5b508035600160a060020a03169060200135600160e060020a0319166108aa565b6000600160a060020a038416156102cd57836102cf565b335b9050336102db82610570565b600160a060020a031614610339576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b6103428361092a565b15610397576040805160e560020a62461bcd02815260206004820152601a60248201527f4d757374206e6f7420626520616e204552433136352068617368000000000000604482015290519081900360640190fd5b600160a060020a038216158015906103b85750600160a060020a0382163314155b156104ff5760405160200180807f455243313832305f4143434550545f4d4147494300000000000000000000000081525060140190506040516020818303038152906040528051906020012082600160a060020a031663249cb3fa85846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182600160a060020a0316600160a060020a031681526020019250505060206040518083038186803b15801561047e57600080fd5b505afa158015610492573d6000803e3d6000fd5b505050506040513d60208110156104a857600080fd5b5051146104ff576040805160e560020a62461bcd02815260206004820181905260248201527f446f6573206e6f7420696d706c656d656e742074686520696e74657266616365604482015290519081900360640190fd5b600160a060020a03818116600081815260208181526040808320888452909152808220805473ffffffffffffffffffffffffffffffffffffffff19169487169485179055518692917f93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db15391a450505050565b600160a060020a03818116600090815260016020526040812054909116151561059a5750806105b7565b50600160a060020a03808216600090815260016020526040902054165b919050565b336105c683610570565b600160a060020a031614610624576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b81600160a060020a031681600160a060020a0316146106435780610646565b60005b600160a060020a03838116600081815260016020526040808220805473ffffffffffffffffffffffffffffffffffffffff19169585169590951790945592519184169290917f605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a43509190a35050565b600082826040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012090505b92915050565b6106f882826107ef565b610703576000610705565b815b600160a060020a03928316600081815260208181526040808320600160e060020a031996909616808452958252808320805473ffffffffffffffffffffffffffffffffffffffff19169590971694909417909555908152600284528181209281529190925220805460ff19166001179055565b600080600160a060020a038416156107905783610792565b335b905061079d8361092a565b156107c357826107ad82826108aa565b6107b85760006107ba565b815b925050506106e8565b600160a060020a0390811660009081526020818152604080832086845290915290205416905092915050565b6000808061081d857f01ffc9a70000000000000000000000000000000000000000000000000000000061094c565b909250905081158061082d575080155b1561083d576000925050506106e8565b61084f85600160e060020a031961094c565b909250905081158061086057508015155b15610870576000925050506106e8565b61087a858561094c565b909250905060018214801561088f5750806001145b1561089f576001925050506106e8565b506000949350505050565b600160a060020a0382166000908152600260209081526040808320600160e060020a03198516845290915281205460ff1615156108f2576108eb83836107ef565b90506106e8565b50600160a060020a03808316600081815260208181526040808320600160e060020a0319871684529091529020549091161492915050565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161590565b6040517f01ffc9a7000000000000000000000000000000000000000000000000000000008082526004820183905260009182919060208160248189617530fa90519096909550935050505056fea165627a7a72305820377f4a2d4301ede9949f163f319021a6e9c687c292a5e2b2c4734c126b524e6c00291ba01820182018201820182018201820182018201820182018201820182018201820a01820182018201820182018201820182018201820182018201820182018201820"
)

var (
	key    *ecdsa.PrivateKey
	client *ethclient.Client
)

var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Helper tools to deploy and interact with the Shutter contracts",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), dialDefaultTimeout)
		defer cancel()
		key = sandbox.GanacheKey(ganacheKeyIdx)

		cl, err := ethclient.DialContext(ctx, "http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}
		client = cl
	},
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy all contracts",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), deployDefaultTimeout)
		defer cancel()
		deploy(ctx)
	},
}

var scheduleFlags struct {
	ConfigContractAddress string
	StartBatchIndex       int
	BatchSpan             int
	StartBlockNumber      int
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a new config",
	Run: func(cmd *cobra.Command, args []string) {
		configContractAddress := common.HexToAddress(scheduleFlags.ConfigContractAddress)
		if scheduleFlags.ConfigContractAddress != configContractAddress.Hex() {
			log.Fatalf("Invalid config contract address %s", scheduleFlags.ConfigContractAddress)
		}

		ctx, cancel := context.WithTimeout(context.Background(), scheduleDefaultTimeout)
		defer cancel()

		schedule(
			ctx,
			configContractAddress,
			uint64(scheduleFlags.StartBatchIndex),
			uint64(scheduleFlags.BatchSpan),
			uint64(scheduleFlags.StartBlockNumber),
		)
	},
}

var getconfigFlags struct {
	ConfigContractAddress string
	ConfigIndex           int
}

var getconfigCmd = &cobra.Command{
	Use:   "getconfig",
	Short: "Get the config with the given index",
	Run: func(cmd *cobra.Command, args []string) {
		configContractAddress := common.HexToAddress(getconfigFlags.ConfigContractAddress)
		if getconfigFlags.ConfigContractAddress != configContractAddress.Hex() {
			log.Fatalf("Invalid config contract address %s", getconfigFlags.ConfigContractAddress)
		}
		ctx, cancel := context.WithTimeout(context.Background(), getconfigDefaultTimeout)
		defer cancel()
		getconfig(ctx, configContractAddress, getconfigFlags.ConfigIndex)
	},
}

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Fund accounts",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), getconfigDefaultTimeout)
		defer cancel()
		fund(ctx)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(getconfigCmd)
	rootCmd.AddCommand(fundCmd)

	initScheduleFlags()
	initGetconfigFlags()
}

func initScheduleFlags() {
	scheduleCmd.Flags().StringVarP(
		&scheduleFlags.ConfigContractAddress,
		"config-contract",
		"c",
		"",
		"address of config contract",
	)
	err := scheduleCmd.MarkFlagRequired("config-contract")
	if err != nil {
		panic(err)
	}

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBatchIndex,
		"start-batch-index",
		0,
		"the start batch index",
	)
	err = scheduleCmd.MarkFlagRequired("start-batch-index")
	if err != nil {
		panic(err)
	}

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.BatchSpan,
		"batch-span",
		0,
		"the batch span",
	)
	err = scheduleCmd.MarkFlagRequired("batch-span")
	if err != nil {
		panic(err)
	}

	scheduleCmd.Flags().IntVar(
		&scheduleFlags.StartBlockNumber,
		"start-block-number",
		0,
		"the start block number",
	)
	err = scheduleCmd.MarkFlagRequired("start-block-number")
	if err != nil {
		panic(err)
	}
}

func initGetconfigFlags() {
	getconfigCmd.Flags().StringVarP(
		&getconfigFlags.ConfigContractAddress,
		"config-contract",
		"c",
		"",
		"address of config contract",
	)
	err := getconfigCmd.MarkFlagRequired("config-contract")
	if err != nil {
		panic(err)
	}

	getconfigCmd.Flags().IntVarP(
		&getconfigFlags.ConfigIndex,
		"config-index",
		"i",
		0,
		"the config index",
	)
	err = getconfigCmd.MarkFlagRequired("config-index")
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func waitForTransactionReceipt(ctx context.Context, cl *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := cl.TransactionReceipt(ctx, txHash)
		if err == ethereum.NotFound {
			time.Sleep(time.Second)
			continue
		}
		return receipt, err
	}
}

func waitForTransactions(ctx context.Context, client *ethclient.Client, txs []*types.Transaction) ([]*types.Receipt, error) {
	defer fmt.Print("\n")
	var res []*types.Receipt

	failedTxs := []int{}
	for i, tx := range txs {
		receipt, err := waitForTransactionReceipt(ctx, client, tx.Hash())
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
		return res, errors.New("some txs have failed")
	}

	return res, nil
}

func makeKeypers() []common.Address {
	var keypers []common.Address
	for i := 0; i < numKeypers; i++ {
		keypers = append(keypers, crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey))
	}
	return keypers
}

func makeAuth(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = defaultGasLimit
	auth.GasPrice = gasPrice
	return auth, nil
}

func deploy(ctx context.Context) {
	deployERC1820Contract(ctx)

	auth, err := makeAuth(ctx, client, key)
	if err != nil {
		panic(err)
	}
	auth.Context = ctx
	var txs []*types.Transaction
	var tx *types.Transaction

	addTx := func() {
		if err != nil {
			panic(err)
		}
		txs = append(txs, tx)
		auth.Nonce.SetInt64(auth.Nonce.Int64() + 1)
	}

	configAddress, tx, _, err := contract.DeployConfigContract(
		auth,
		client,
		defaultConfigChangeHeadsUpBlocks,
	)
	addTx()

	broadcastAddress, tx, _, err := contract.DeployKeyBroadcastContract(auth, client, configAddress)
	addTx()

	feeAddress, tx, _, err := contract.DeployFeeBankContract(auth, client)
	addTx()

	batcherAddress, tx, _, err := contract.DeployBatcherContract(auth, client, configAddress, feeAddress)
	addTx()

	executorAddress, tx, _, err := contract.DeployExecutorContract(auth, client, configAddress, batcherAddress)
	addTx()

	tokenAddress, tx, _, err := contract.DeployTestDepositTokenContract(auth, client)
	addTx()

	depositAddress, tx, _, err := contract.DeployDepositContract(auth, client, tokenAddress)
	addTx()

	_, err = waitForTransactions(ctx, client, txs)
	if err != nil {
		panic(err)
	}
	fmt.Println("ConfigContract address:", configAddress.Hex())
	fmt.Println("KeyBroadcastContract address:", broadcastAddress.Hex())
	fmt.Println("FeeBankContract address:", feeAddress.Hex())
	fmt.Println("BatcherContract address:", batcherAddress.Hex())
	fmt.Println("ExecutorContract address:", executorAddress.Hex())
	fmt.Println("TokenContract address:", tokenAddress.Hex())
	fmt.Println("DepositContract address:", depositAddress.Hex())
}

// deployERC1820Contract deploys the ERC1820 contract as described in
// https://eips.ethereum.org/EIPS/eip-1820
func deployERC1820Contract(ctx context.Context) {
	// fund deployer account
	fundingAccount := crypto.PubkeyToAddress(key.PublicKey)
	chainID, err := client.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	signer := types.NewEIP155Signer(chainID)

	erc1820Deployer := common.HexToAddress(erc1820DeploymentAccountHex)
	nonce, err := client.PendingNonceAt(ctx, fundingAccount)
	if err != nil {
		panic(err)
	}
	amount := big.NewInt(80000000000000000) // 0.08 ETH
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}

	unsignedTx := types.NewTransaction(nonce, erc1820Deployer, amount, 21000, gasPrice, []byte{})
	tx, err := types.SignTx(unsignedTx, signer, key)
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(ctx, tx)
	if err != nil {
		panic(err)
	}

	// deploy contract
	txBytes, err := hexutil.Decode(erc1820DeploymentTransactionHex)
	if err != nil {
		panic(err)
	}
	txStream := rlp.NewStream(bytes.NewReader(txBytes), 0)
	tx = new(types.Transaction)
	err = tx.DecodeRLP(txStream)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		panic(err)
	}
}

func checkContractExists(ctx context.Context, configContractAddress common.Address) {
	code, err := client.CodeAt(ctx, configContractAddress, nil)
	if err != nil {
		panic(err)
	}
	if len(code) == 0 {
		log.Fatalf("No contract deployed at address %s", configContractAddress.Hex())
	}
}

func schedule(
	ctx context.Context,
	configContractAddress common.Address,
	startBatchIndex, batchSpan, startBlockNumber uint64,
) {
	auth, err := makeAuth(ctx, client, sandbox.GanacheKey(ganacheKeyIdx))
	if err != nil {
		panic(err)
	}
	auth.Context = ctx

	var txs []*types.Transaction
	var tx *types.Transaction

	addTx := func() {
		if err != nil {
			panic(err)
		}
		txs = append(txs, tx)
		auth.Nonce.SetInt64(auth.Nonce.Int64() + 1)
	}

	checkContractExists(ctx, configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	if err != nil {
		panic(err)
	}

	tx, err = cc.NextConfigSetStartBatchIndex(auth, startBatchIndex)
	addTx()

	tx, err = cc.NextConfigSetBatchSpan(auth, batchSpan)
	addTx()

	tx, err = cc.NextConfigAddKeypers(auth, makeKeypers())
	addTx()

	tx, err = cc.NextConfigSetThreshold(auth, threshold)
	addTx()

	tx, err = cc.NextConfigSetExecutionTimeout(auth, batchSpan)
	addTx()

	tx, err = cc.NextConfigSetTransactionSizeLimit(auth, transactionSizeLimit)
	addTx()

	tx, err = cc.NextConfigSetBatchSizeLimit(auth, batchSizeLimit)
	addTx()

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		panic(err)
	}
	minStartBlockNumber := header.Number.Uint64() + startBlockNumberOffset
	if startBlockNumber < minStartBlockNumber {
		log.Fatalf(
			"Start block number %d is too close to current head %d (required offset %d)",
			startBlockNumber,
			header.Number,
			startBlockNumberOffset,
		)
	}
	tx, err = cc.NextConfigSetStartBlockNumber(auth, startBlockNumber)
	addTx()

	tx, err = cc.ScheduleNextConfig(auth)
	addTx()

	_, err = waitForTransactions(ctx, client, txs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("start block of config: %d\n", startBlockNumber)
}

func getconfig(ctx context.Context, configContractAddress common.Address, index int) {
	checkContractExists(ctx, configContractAddress)
	cc, err := contract.NewConfigContract(configContractAddress, client)
	if err != nil {
		panic(err)
	}
	c, err := cc.GetConfigByIndex(&bind.CallOpts{Context: ctx}, uint64(index))
	if err != nil {
		panic(err)
	}

	printConfig(c)
}

func printConfig(config contract.BatchConfig) {
	fmt.Printf("     Start batch index: %d\n", config.StartBatchIndex)
	fmt.Printf("    Start block number: %d\n", config.StartBlockNumber)
	fmt.Printf("            Batch span: %d\n", config.BatchSpan)
	fmt.Printf("           Num keypers: %d\n", len(config.Keypers))
	fmt.Printf("             Threshold: %d\n", config.Threshold)
	fmt.Printf("        BatchSizeLimit: %d\n", config.BatchSizeLimit)
	fmt.Printf("  TransactionSizeLimit: %d\n", config.TransactionSizeLimit)
	fmt.Printf("   TransactionGasLimit: %d\n", config.TransactionGasLimit)
	fmt.Printf("           FeeReceiver: %s\n", config.FeeReceiver.Hex())
	fmt.Printf("         TargetAddress: %s\n", config.TargetAddress.Hex())
	fmt.Printf("TargetFunctionSelector: %x\n", config.TargetFunctionSelector)
	fmt.Printf("      ExecutionTimeout: %d\n", config.ExecutionTimeout)
	fmt.Printf("\nKeypers:\n")
	if len(config.Keypers) == 0 {
		fmt.Printf("  None\n")
	} else {
		for i, k := range config.Keypers {
			fmt.Printf("  %d: %s\n", i, k.Hex())
		}
	}
}

func fund(ctx context.Context) {
	fromAddress := crypto.PubkeyToAddress(key.PublicKey)

	var txs []*types.Transaction
	var tx *types.Transaction

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		panic(err)
	}

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		panic(err)
	}

	amount := big.NewInt(fundAmount) // 1 eth
	gasLimit := uint64(baseGasLimit)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	addTx := func() {
		if err != nil {
			panic(err)
		}
		txs = append(txs, tx)
		nonce++
	}

	signer := types.NewEIP155Signer(chainID)

	for i := 0; i < sandbox.NumGanacheKeys()-1; i++ {
		receiver := crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey)
		var data []byte
		tx = types.NewTransaction(nonce, receiver, amount, gasLimit, gasPrice, data)

		tx, err = types.SignTx(tx, signer, key)
		if err != nil {
			panic(err)
		}

		err = client.SendTransaction(ctx, tx)
		if err != nil {
			panic(err)
		}
		addTx()
	}

	_, err = waitForTransactions(ctx, client, txs)
	if err != nil {
		panic(err)
	}

	for i := 0; i < sandbox.NumGanacheKeys(); i++ {
		addr := crypto.PubkeyToAddress(sandbox.GanacheKey(i).PublicKey)
		balance, err := client.BalanceAt(context.Background(), addr, nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %s\n", addr.Hex(), balance)
	}
}
