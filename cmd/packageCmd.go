package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func packageCmdRun(cmd *cobra.Command, args []string) {
	var dir string
	var err error
	if cmd.Flags().Changed("dir") {
		dir, err = cmd.Flags().GetString("dir")
		if err != nil {
			log.Fatalf("dir error %s", err.Error())
		}
	}
	fmt.Fprintf(os.Stdout, "packaging %s\n", dir)

	pkg, err := NewPkg(dir)
	if err != nil {
		log.Fatal(err)
	}
	err = pkg.Build()
	if err != nil {
		log.Fatal(err)
	}
}
