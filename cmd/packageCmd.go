package cmd

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	path := filepath.Dir(dir)
	base := filepath.Base(dir)

	if base != "Content" {
		fmt.Printf("error: %s does not end with 'Content'\n", dir)
		os.Exit(1)
	}

	//os.Exit(0)
	files := GetFiles(path + "/")
	//for _, f := range files {
	//fmt.Printf("%#v\n", f)
	//}
	man := NewManifest(files)

	d, err := xml.MarshalIndent(man, "", "  ")
	if err != nil {
		log.Fatalf("dir error %s", err.Error())
	}

	err = os.WriteFile(filepath.Join(path, manifest), d, 0666)
	if err != nil {
		log.Fatalf("dir error %s", err.Error())
	}
}
