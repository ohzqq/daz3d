package cmd

import (
	"fmt"
	"testing"
)

const testPath = `../testdata/Muscularity Morphs for Genesis 9/`

func TestDirFS(t *testing.T) {
	files, err := GetFilesFS(testPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		fmt.Printf("%#v\n", f)
	}
}

func TestZip(t *testing.T) {
}
