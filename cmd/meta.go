package cmd

import (
	"encoding/xml"
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

type Supplement struct {
	XMLName     xml.Name `xml:"ProductSupplement"`
	Version     string   `xml:"VERSION,attr"`
	ProductName struct {
		Value string `xml:"VALUE,attr"`
	}
	InstallTypes struct {
		Value string `xml:"VALUE,attr"`
	}
	ProductTags struct {
		Value string `xml:"VALUE,attr"`
	}
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

func NewFile(path string) *File {
	return &File{
		Target: "Content",
		Action: "Install",
		Value:  path,
	}
}

func NewSupplement(name string) *Supplement {
	return &Supplement{
		Version: "0.1",
		ProductName: struct {
			Value string `xml:"VALUE,attr"`
		}{
			Value: filepath.Base(name),
		},
		InstallTypes: struct {
			Value string `xml:"VALUE,attr"`
		}{
			Value: "Content",
		},
		ProductTags: struct {
			Value string `xml:"VALUE,attr"`
		}{
			Value: "DAZStudio4_5",
		},
	}
}
