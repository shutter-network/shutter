// Package cmd implements the shuttermint subcommands
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/shutter-network/shutter/shuttermint/cmd/config"
	"github.com/shutter-network/shutter/shuttermint/cmd/deploy"
	"github.com/shutter-network/shutter/shuttermint/cmd/prepare"
	"github.com/shutter-network/shutter/shuttermint/cmd/shversion"
	"github.com/shutter-network/shutter/shuttermint/medley"
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
		err := medley.BindFlags(cmd)
		if err != nil {
			return err
		}
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
	Run:          medley.ShowHelpAndExit,
	SilenceUsage: true,
}

// Execute the shuttermint root command and exit the program afterwards. This is called from main.
func Execute() {
	status := 0

	if err := rootCmd.Execute(); err != nil {
		status = 1
	}

	os.Exit(status)
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
	rootCmd.AddCommand(deploy.DeployCmd)
	rootCmd.AddCommand(prepare.PrepareCmd)
	rootCmd.AddCommand(completionCmd)
}
