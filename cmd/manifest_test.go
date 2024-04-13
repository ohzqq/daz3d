package cmd

import (
	"encoding/xml"
	"fmt"
	"testing"
)

const testPath = `../testdata/Muscularity Morphs for Genesis 9/`

func TestNewManifest(t *testing.T) {
	pkg, err := NewPkg(testPath)
	if err != nil {
		t.Fatal(err)
	}
	println(pkg.Name)

	d, err := xml.MarshalIndent(pkg.supplement, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	println(string(d))
}

func TestDirFS(t *testing.T) {
	files, err := GetFilesFS(testPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		fmt.Printf("%#v\n", f)
	}
}

//func TestZip(t *testing.T) {
//name := pkgName()
//}
