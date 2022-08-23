package main

import (
	_ "embed"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed index.html.template
var t string

func genTemplate(path string, v interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return template.Must(template.New("").Parse(t)).Execute(f, v)
}

type obj struct {
	Path, Label         string
	TotalSizeBitesLogic uint32
	CntLogicFiles       uint8
}

func imputeAssetsDetails(p string, o *obj) error {
	if err := filepath.WalkDir(
		p, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				o.CntLogicFiles += 1

				i, err := d.Info()
				if err != nil {
					return err
				}
				o.TotalSizeBitesLogic += uint32(i.Size())
			}
			return nil
		},
	); err != nil {
		return err
	}
	return nil
}

func gatherReportData(p string) ([]obj, error) {
	var o []obj

	if err := filepath.WalkDir(
		p, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() && filepath.Dir(path) == p {
				assets := obj{
					Path:  d.Name(),
					Label: strings.ToUpper(d.Name()),
				}
				if err := imputeAssetsDetails(path+"/assets/logic", &assets); err != nil {
					return err
				}
				o = append(o, assets)
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return o, nil
}

func main() {
	var p string
	flag.StringVar(&p, "path-pages", "", "path to keep pages data")
	flag.Parse()

	if p == "" {
		log.Fatalln("path to pages must be set")
	}

	r, err := gatherReportData(p)
	if err != nil {
		log.Fatalln("cannot gather report", err)
	}

	if err := genTemplate(p+"/index.html", r); err != nil {
		log.Fatalln("error generating template", err)
	}
}
