package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func packageCmdRun(cmd *cobra.Command, args []string) {
	var dir string
	var err error
	if cmd.Flags().Changed("dir") {
		dir, err = cmd.Flags().GetString("dir")
		if err != nil {
			fmt.Errorf("dir error %s", err.Error())
			os.Exit(1)
		}
	}
	fmt.Fprintf(os.Stdout, "packaging %s\n", dir)

	pkg, err := NewPkg(dir)
	if err != nil {
		fmt.Errorf("new package error: %w\n", err)
		os.Exit(1)
	}
	err = pkg.Build()
	if err != nil {
		fmt.Errorf("package build error: %w\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
