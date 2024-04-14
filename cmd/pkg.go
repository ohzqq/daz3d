package cmd

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/flect"
)

type Pkg struct {
	Name       string
	path       string
	dir        string
	base       string
	fs         fs.FS
	manifest   *Manifest
	supplement *Supplement
}

func NewPkg(path string) (*Pkg, error) {
	path = strings.TrimSuffix(path, "/")

	dfs := os.DirFS(path)

	_, err := fs.Stat(dfs, "Content")
	if err != nil {
		return nil, err
	}

	p := &Pkg{
		path:       path,
		dir:        filepath.Dir(path),
		base:       filepath.Base(path),
		fs:         dfs,
		manifest:   NewManifest(),
		supplement: NewSupplement(path),
	}
	p.Name = filepath.Join(p.dir, pkgName(p.base))

	err = p.GetFiles()
	if err != nil {
		return p, fmt.Errorf("error getting files %w\n", err)
	}

	return p, nil
}

func (pkg *Pkg) Build() error {
	err := pkg.WriteManifest()
	if err != nil {
		return err
	}

	err = pkg.WriteSupplement()
	if err != nil {
		return err
	}

	err = pkg.Zip()
	if err != nil {
		return fmt.Errorf("zip error: %w\n", err)
	}

	return nil
}

func (pkg *Pkg) WriteManifest() error {
	d, err := xml.MarshalIndent(pkg.manifest, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(pkg.path, manifest))
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

func (pkg *Pkg) WriteSupplement() error {
	d, err := xml.MarshalIndent(pkg.supplement, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(pkg.path, supplement))
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

func (pkg *Pkg) GetFiles() error {
	fn := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			f := filepath.Join(path)
			pkg.manifest.File = append(pkg.manifest.File, NewFile(f))
		}
		return nil
	}

	err := fs.WalkDir(pkg.fs, "Content", fn)
	if err != nil {
		return err
	}

	return nil
}

func (pkg *Pkg) Zip() error {
	z, err := os.Create(pkg.Name)
	if err != nil {
		return err
	}
	defer z.Close()

	w := zip.NewWriter(z)
	defer w.Close()

	err = w.AddFS(pkg.fs)
	if err != nil {
		return err
	}

	return nil
}

func pkgName(dir string) string {
	name := flect.Pascalize(dir)
	sku := genSKU()
	return fmt.Sprintf("CH%08d-01_%s.zip", sku, name)
}

func genSKU() int {
	return rand.Intn(100000000)
}
