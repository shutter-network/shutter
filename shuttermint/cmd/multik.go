package cmd

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
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
	baseConfig := keyper.KeyperConfig{
		ShuttermintURL:                 "http://localhost:26657",
		EthereumURL:                    "ws://localhost:8545",
		ConfigContractAddress:          common.HexToAddress("0x07a457d878BF363E0Bb5aa0B096092f941e19962"),
		KeyBroadcastingContractAddress: common.HexToAddress("0xFA33c8EF8b5c4f3003361c876a298D1DB61ccA4e"),
	}

	privateKey := sandbox.GanacheKey(sandbox.NumGanacheKeys() - 1)

	var signingKeys [3]*ecdsa.PrivateKey
	var validatorKeys [3]ed25519.PrivateKey
	var keypers [3]common.Address
	for i := 0; i < 3; i++ {
		k := sandbox.GanacheKey(i)
		signingKeys[i] = k
		keypers[i] = crypto.PubkeyToAddress(k.PublicKey)

		validatorSeed := make([]byte, 32)
		copy(keypers[i].Bytes(), validatorSeed)
		validatorKeys[i] = ed25519.NewKeyFromSeed(validatorSeed)
	}

	ethcl, err := ethclient.Dial(baseConfig.EthereumURL)
	configContract, err := contract.NewConfigContract(baseConfig.ConfigContractAddress, ethcl)
	header, err := ethcl.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	headBlockNumber := header.Number.Uint64()

	startBatchIndex, err := configContract.NextBatchIndex(headBlockNumber)
	if err != nil {
		panic(err)
	}

	var shmcl client.Client
	shmcl, err = http.New(baseConfig.ShuttermintURL, "/websocket")
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
		go func(signingKey *ecdsa.PrivateKey, validatorKey ed25519.PrivateKey) {
			config := baseConfig
			config.SigningKey = signingKey
			config.ValidatorKey = validatorKey

			kpr := keyper.NewKeyper(config)
			err = kpr.Run()
			if err != nil {
				panic(err)
			}
		}(signingKeys[i], validatorKeys[i])
	}
	time.Sleep(time.Hour)
}
