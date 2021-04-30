// Package cmd implements the shuttermint subcommands
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/brainbot-com/shutter/shuttermint/cmd/config"
	"github.com/brainbot-com/shutter/shuttermint/cmd/shversion"
)

var (
	cfgFile   string
	logformat string
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "shuttermint",
	Short:   "A collection of commands to run and interact with Shutter keyper nodes",
	Version: shversion.Version(),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var flags int

		switch logformat {
		case "min":
		case "short":
			flags = log.Lshortfile
		case "long":
			flags = log.LstdFlags | log.Lshortfile | log.Lmicroseconds
		case "max":
			flags = log.LstdFlags | log.Llongfile | log.Lmicroseconds
		default:
			return fmt.Errorf("bad log value, possible values: min, short, long, max")
		}

		log.SetFlags(flags)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logformat, "log", "long", "set log format, possible values:  min, short, long, max")
	rootCmd.AddCommand(chainCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(keyperCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(txsearchCmd)
	rootCmd.AddCommand(bootstrapCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(deployCmd)
}
