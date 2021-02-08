package cmd

import (
	"context"

	"github.com/kr/pretty"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/brainbot-com/shutter/shuttermint/keyper/observe"
)

var showFlags struct {
	ShuttermintURL string
	Height         int64
}

// keyperCmd represents the keyper command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show shutter internal state",
	Run: func(cmd *cobra.Command, args []string) {
		showMain()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.PersistentFlags().StringVarP(
		&showFlags.ShuttermintURL,
		"shuttermint-url",
		"s",
		"http://localhost:26657",
		"Shuttermint RPC URL",
	)
	showCmd.PersistentFlags().Int64VarP(
		&showFlags.Height,
		"height",
		"",
		-1,
		"target height",
	)
}

func showShutter(shuttermintURL string, height int64) {
	var cl client.Client
	cl, err := http.New(shuttermintURL, "/websocket")
	if err != nil {
		panic(err)
	}

	s := observe.NewShutter()
	latestBlock, err := cl.Block(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	if latestBlock.Block == nil {
		panic("empty shutter blockchain")
	}
	if height == -1 {
		height = latestBlock.Block.Height
	}
	s, err = s.SyncToHeight(context.Background(), cl, height)
	if err != nil {
		panic(err)
	}
	pretty.Println("Synced:", s)
}

func showMain() {
	showShutter(showFlags.ShuttermintURL, showFlags.Height)
}
