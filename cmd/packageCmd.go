package cmd

import (
	"github.com/spf13/cobra"
)

func packageCmdRun(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("dir") {
		// todo
	}
	println(cmd.Name())
}
