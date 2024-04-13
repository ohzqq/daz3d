package cmd

import (
	"encoding/xml"
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
