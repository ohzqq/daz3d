package cmd

import (
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/spf13/cobra"
)

type Pkg struct {
	path string
	base string
}

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

	dir = strings.TrimSuffix(dir, "/")

	path := filepath.Dir(dir)
	base := filepath.Base(dir)

	if base != "Content" {
		fmt.Printf("error: %s does not end with 'Content'\n", dir)
		os.Exit(1)
	}

	err = genMan(path)
	if err != nil {
		log.Fatalf("error generating manifest %s", err.Error())
	}

	err = genSup(path)
	if err != nil {
		log.Fatalf("error generating supplement %s", err.Error())
	}

	pn := pkgName(path)

	err = os.Mkdir(filepath.Join(path, pn), 0777)
	if err != nil {
		log.Fatalf("error generating supplement %s", err.Error())
	}
}

func genMan(path string) error {
	files := GetFiles(path + "/")
	man := NewManifest(files)

	d, err := xml.MarshalIndent(man, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path, manifest))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(xml.Header)
	if err != nil {
		return err
	}

	_, err = f.Write(d)
	if err != nil {
		return err
	}

	return nil
}

func genSup(path string) error {
	sup := NewSupplement(path)
	d, err := xml.MarshalIndent(sup, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path, supplement))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(xml.Header)
	if err != nil {
		return err
	}

	_, err = f.Write(d)
	if err != nil {
		return err
	}

	return nil
}

func pkgName(dir string) string {
	name := flect.Pascalize(filepath.Base(dir))
	sku := genSKU()
	return fmt.Sprintf("CH%08d-01_%s", sku, name)
}

func genSKU() int {
	return rand.Intn(100000000)
}
