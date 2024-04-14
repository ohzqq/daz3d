package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pre = "ZZ"

func packageCmdRun(cmd *cobra.Command, args []string) {
	var dir string
	var err error

	if len(args) > 0 {
		dir = args[0]
	}

	if dir == "" {
		err = errors.New("no dir provided")
	}

	if err != nil {
		fmt.Printf("dir error %s", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "packaging %s\n", dir)

	if cmd.Flags().Changed("prefix") {
		pre, err = cmd.Flags().GetString("prefix")
		if err != nil {
			pre = `"ZZ"`
		}
	}

	pkg, err := NewPkg(dir)
	if err != nil {
		fmt.Printf("new package error: %s\n", err)
		os.Exit(1)
	}

	err = pkg.Build()
	if err != nil {
		fmt.Errorf("package build error: %w\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
