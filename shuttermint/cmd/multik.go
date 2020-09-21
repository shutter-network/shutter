package cmd

import (
	"context"
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/brainbot-com/shutter/shuttermint/contract"
	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/brainbot-com/shutter/shuttermint/sandbox"
)

// multikCmd represents the multik command
var multikCmd = &cobra.Command{
	Use:   "multik",
	Short: "Run multiple keypers in a single process for testing purposes",
	Run: func(cmd *cobra.Command, args []string) {
		multikMain()
	},
}

func init() {
	rootCmd.AddCommand(multikCmd)
}

func multikMain() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Printf("Starting multik version %s", version)
	shuttermintURL := "http://localhost:26657"
	ethereumURL := "ws://localhost:8545"
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}

	var keys [3]*ecdsa.PrivateKey
	var keypers [3]common.Address
	for i := 0; i < 3; i++ {
		k := sandbox.GanacheKey(i)
		keys[i] = k
		keypers[i] = crypto.PubkeyToAddress(k.PublicKey)
	}

	ethcl, err := ethclient.Dial(ethereumURL)
	addr := common.HexToAddress("0x07a457d878BF363E0Bb5aa0B096092f941e19962")
	configContract, err := contract.NewConfigContract(addr, ethcl)
	header, err := ethcl.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	headBlockNumber := header.Number.Uint64()

	startBatchIndex, err := configContract.NextBatchIndex(headBlockNumber)

	var shmcl client.Client
	shmcl, err = http.New(shuttermintURL, "/websocket")
	if err != nil {
		panic(err)
	}

	ms := keyper.NewMessageSender(shmcl, privateKey)
	bc := keyper.NewBatchConfig(startBatchIndex, keypers[:], 2)
	err = ms.SendMessage(bc)
	if err != nil {
		panic(err)
	}
	log.Printf("Send new BatchConfig (start batch index %d)", startBatchIndex)
	for i := 0; i < 3; i++ {
		go func(key *ecdsa.PrivateKey) {
			kpr := keyper.NewKeyper(key, shuttermintURL, ethereumURL)
			err = kpr.Run()
			if err != nil {
				panic(err)
			}
		}(keys[i])
	}
	time.Sleep(time.Hour)
}
