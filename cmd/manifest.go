package cmd

import (
	"encoding/xml"
	"io/fs"
	"log"
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
	GlobalID *GlobalID
	File     []*File
}

type GlobalID struct {
	XMLName xml.Name `xml:"GlobalID"`
	Value   string   `xml:"VALUE,attr"`
}

type File struct {
	XMLName xml.Name `xml:"File"`
	Target  string   `xml:"TARGET,attr"`
	Action  string   `xml:"ACTION,attr"`
	Value   string   `xml:"VALUE,attr"`
}

func NewManifest(files []*File) *Manifest {
	return &Manifest{
		Version:  "0.1",
		GlobalID: NewGlobalID(),
		File:     files,
	}
}

func NewGlobalID() *GlobalID {
	return &GlobalID{
		Value: strings.ToUpper(uuid.New().String()),
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
		f := filepath.Join(strings.TrimPrefix(path, root))
		files = append(files, NewFile(f))
		return nil
	}

	err := filepath.WalkDir(filepath.Join(root, "Content"), fn)
	if err != nil {
		log.Fatal(err)
	}

	return files
}
