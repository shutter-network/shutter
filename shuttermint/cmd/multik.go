package cmd

import (
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/spf13/cobra"
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
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}

	var keys [3]*ecdsa.PrivateKey
	var keypers [3]common.Address
	for i := 0; i < 3; i++ {
		k, err := crypto.GenerateKey()
		if err != nil {
			panic(err)
		}
		keys[i] = k
		keypers[i] = crypto.PubkeyToAddress(k.PublicKey)
	}

	var cl client.Client
	cl, err = http.New(shuttermintURL, "/websocket")
	if err != nil {
		panic(err)
	}

	ms := keyper.NewMessageSender(cl, privateKey)
	bc := keyper.NewBatchConfig(keyper.NextBatchIndex(time.Now()),
		keypers[:],
		2)
	err = ms.SendMessage(bc)
	if err != nil {
		panic(err)
	}
	log.Print("Send new BatchConfig")
	for i := 0; i < 3; i++ {
		go func(key *ecdsa.PrivateKey) {
			k := keyper.NewKeyper(key, shuttermintURL)
			err = k.Run()
			if err != nil {
				panic(err)
			}
		}(keys[i])
	}
	time.Sleep(time.Hour)
}
