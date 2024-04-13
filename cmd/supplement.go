package cmd

import (
	"encoding/xml"
	"path/filepath"
)

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
