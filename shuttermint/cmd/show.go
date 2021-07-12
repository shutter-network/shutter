package cmd

import (
	"context"

	"github.com/kr/pretty"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/shutter-network/shutter/shuttermint/keyper/observe"
)

var showFlags struct {
	ShuttermintURL string
	Height         int64
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the internal state of a Shuttermint node",
	Long: `This command queries transactions from a running shuttermint node and rebuilds the
internal shutter state object according to the results. It then prints the result to stdout.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		showMain()
	},
}

func init() {
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
	if height == -1 {
		height, err = s.GetLastCommittedHeight(context.Background(), cl)
		if err != nil {
			panic(err)
		}
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
