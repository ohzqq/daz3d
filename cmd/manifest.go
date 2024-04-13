package cmd

import (
	"encoding/xml"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const (
	manifest   = `Manifest.dsx`
	supplement = `Supplement.dsx`
)

type Manifest struct {
	XMLName  xml.Name `xml:"DAZInstallManifest"`
	Version  string   `xml:"VERSION,attr"`
	GlobalID struct {
		Value string `xml:"VALUE,attr"`
	}
	File []*File
}

type File struct {
	XMLName xml.Name `xml:"File"`
	Target  string   `xml:"TARGET,attr"`
	Action  string   `xml:"ACTION,attr"`
	Value   string   `xml:"VALUE,attr"`
}

func NewManifest() *Manifest {
	return &Manifest{
		Version: "0.1",
		GlobalID: struct {
			Value string `xml:"VALUE,attr"`
		}{
			Value: strings.ToUpper(uuid.New().String()),
		},
	}
}

func NewMan(files []*File) *Manifest {
	return &Manifest{
		Version: "0.1",
		GlobalID: struct {
			Value string `xml:"VALUE,attr"`
		}{
			Value: strings.ToUpper(uuid.New().String()),
		},
		File: files,
	}
}

func NewFile(path string) *File {
	return &File{
		Target: "Content",
		Action: "Install",
		Value:  path,
	}
}

func GetFiles(root string) []*File {
	var files []*File

	fn := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			f := filepath.Join(strings.TrimPrefix(path, root))
			files = append(files, NewFile(f))
		}
		return nil
	}

	err := filepath.WalkDir(filepath.Join(root, "Content"), fn)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func GetFilesFS(root string) ([]*File, error) {
	var files []*File

	fn := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			f := filepath.Join(path)
			files = append(files, NewFile(f))
		}
		return nil
	}

	dir := os.DirFS(root)

	err := fs.WalkDir(dir, "Content", fn)
	if err != nil {
		return nil, err
	}

	return files, nil
}
