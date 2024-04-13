package cmd

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Pkg struct {
	Name       string
	path       string
	base       string
	fs         fs.FS
	manifest   *Manifest
	supplement *Supplement
}

func NewPkg(path string) (*Pkg, error) {
	p := &Pkg{
		Name:       pkgName(path),
		path:       filepath.Dir(path),
		base:       filepath.Base(path),
		fs:         os.DirFS(path),
		manifest:   NewManifest(),
		supplement: NewSupplement(path),
	}

	err := p.GetFiles()
	if err != nil {
		return p, fmt.Errorf("error getting files %w\n", err)
	}

	return p, nil
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
	z, err := os.Create(pkg.Name + ".zip")
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
