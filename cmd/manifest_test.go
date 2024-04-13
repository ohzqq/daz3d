package cmd

import (
	"encoding/xml"
	"fmt"
	"testing"
)

const testPath = `../testdata/Muscularity Morphs for Genesis 9`
const testSlash = `../testdata/Muscularity Morphs for Genesis 9/`

func TestBuildPackage(t *testing.T) {
	pkg, err := NewPkg(testPath)
	if err != nil {
		t.Fatal(err)
	}
	err = pkg.Build()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPaths(t *testing.T) {
	t.SkipNow()
	for _, d := range []string{testPath, testSlash} {
		pkg, err := NewPkg(d)
		if err != nil {
			t.Fatal(err)
		}

		println("path " + pkg.path)
		println("base " + pkg.base)
		println(pkg.Name)
	}
}

func TestNewManifest(t *testing.T) {
	t.SkipNow()
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
	t.SkipNow()
	files, err := GetFilesFS(testPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		fmt.Printf("%#v\n", f)
	}
}

func TestZip(t *testing.T) {
	t.SkipNow()
	pkg, err := NewPkg(testPath)
	if err != nil {
		t.Fatal(err)
	}

	err = pkg.Zip()
	if err != nil {
		t.Fatal(err)
	}
}
