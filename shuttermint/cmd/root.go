package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/cmd/config"
	"github.com/brainbot-com/shutter/shuttermint/cmd/shversion"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "shuttermint",
	Short:   "A collection of commands to run and interact with Shutter keyper nodes",
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

func init() {
	rootCmd.AddCommand(chainCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(keyperCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(bootstrapCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(deployCmd)
}
