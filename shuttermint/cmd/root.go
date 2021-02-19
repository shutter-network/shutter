package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/cmd/shversion"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "shuttermint",
	Short:   "shuttermint runs the shutter tendermint app or a shutter keyper",
	Version: shversion.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
