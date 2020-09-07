package cmd

import (
	"log"

	"github.com/brainbot-com/shutter/shuttermint/keyper"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
)

// keyperCmd represents the keyper command
var keyperCmd = &cobra.Command{
	Use:   "keyper",
	Short: "Run a shutter keyper",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		keyperMain()
	},
}

func init() {
	rootCmd.AddCommand(keyperCmd)
}

func keyperMain() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Printf("Starting keyper version %s", version)

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		panic(err)
	}
	k := keyper.NewKeyper(privateKey, "http://localhost:26657")
	err = k.Run()
	if err != nil {
		panic(err)
	}
}
