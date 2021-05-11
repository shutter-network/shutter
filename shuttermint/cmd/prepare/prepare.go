// Package prepare contains the implementation of the prepare subcommand
package prepare

import (
	"github.com/spf13/cobra"
)

var PrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare everything needed to test shutter.",
}

func init() {
	PrepareCmd.AddCommand(configCmd)
	PrepareCmd.AddCommand(fundCmd)

	initConfigFlags()
	initFundFlags()
}
