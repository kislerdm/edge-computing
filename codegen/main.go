package main

import (
	"context"
	_ "embed"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"genpages/speedtest"
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
	Path, Label                     string
	TotalSizeBitesLogic             uint32
	ExecutionMeanUS, ExecutionSemUS uint16
	CntLogicFiles                   uint8
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

func gatherReportData(p string, url string) ([]obj, error) {
	var o []obj

	if err := filepath.WalkDir(
		p, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() && filepath.Dir(path) == p {
				assets := obj{
					Path:  d.Name(),
					Label: d.Name(),
				}
				if err := imputeAssetsDetails(path+"/assets/logic", &assets); err != nil {
					return err
				}
				if err := imputeSpeedtest(url+"/"+d.Name(), &assets); err != nil {
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

func imputeSpeedtest(url string, assets *obj) error {
	o, err := speedtest.Run(url, 1000)
	if err != nil {
		return err
	}
	assets.ExecutionMeanUS = o.MeanUS
	assets.ExecutionSemUS = o.SemUM
	return nil
}

type server struct {
	l    net.Listener
	s    http.Server
	Addr string
}

func StartServer() (*server, error) {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		return nil, err
	}
	const address = "http://localhost:9090"
	return &server{
		l:    l,
		s:    http.Server{Addr: address},
		Addr: address,
	}, nil
}

func (s *server) Serve(path string) error {
	s.s.Handler = http.FileServer(http.Dir(path))
	return s.s.Serve(s.l)
}

func (s *server) Stop() error {
	return s.s.Shutdown(context.Background())
}

func main() {
	var p string
	flag.StringVar(&p, "path-pages", "", "path to keep pages data")
	flag.Parse()

	if p == "" {
		log.Fatalln("path to pages must be set")
	}

	s, err := StartServer()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := s.Serve(p); err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	r, err := gatherReportData(p, s.Addr)
	if err != nil {
		log.Fatalln("cannot gather report", err)
	}

	if err := s.Stop(); err != nil {
		log.Println(err)
	}

	if err := genTemplate(p+"/index.html", r); err != nil {
		log.Fatalln("error generating template", err)
	}
}
