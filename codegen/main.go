package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
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

func genHTML(path string, v interface{}) error {
	return genTemplate(t, path, v)
}

func genDataJS(path string, v interface{}) error {
	d, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	const js = "(()=>{var d={},e=0;for(let[a,f]of(statsData.forEach(b=>{for(let[a,f]of Object.entries(b.output.elapsed))if(void 0===d[a]&&(d[a]={}),d[a][b.label]=f,null===document.getElementById(\"elapsed-time-${k}\")){var c=document.createElement(\"div\");c.setAttribute(\"id\",`elapsed-time-${a}`);let g=document.getElementsByTagName(\"main\")[0];g.appendChild(c)}e++}),Object.entries(d))){var b=[],c=0;for(let[g,h]of Object.entries(f))b.push({x:h.map(a=>1e3*a),name:g,type:\"histogram\",histfunc:\"count\",opacity:1/(1+c/e),xbins:{size:10,start:0}}),c++;Plotly.newPlot(`elapsed-time-${a}`,b,{title:`Logic execution elapsed time, color: ${a}`,yaxis:{title:\"Count\"},xaxis:{title:\"Elapsed time [usec.]\",range:[0,500]},boxmode:\"group\"})}})()"
	return genTemplate(fmt.Sprintf(`const statsData = %v;%s`, string(d), js), path, nil)
}

func genTemplate(t, path string, v interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return template.Must(template.New("").Parse(t)).Execute(f, v)
}

type obj struct {
	Path                string           `json:"_"`
	Label               string           `json:"label"`
	Raw                 speedtest.Output `json:"output"`
	TotalSizeBitesLogic uint32           `json:"totalSizeBitesLogic"`
	ExecutionMeanUS     uint16           `json:"executionMeanUS"`
	ExecutionSemUS      uint16           `json:"executionSemUS"`
	CntLogicFiles       uint8            `json:"cntLogicFiles"`
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
	assets.Raw = o
	assets.ExecutionMeanUS = o.StatsUS.Mean
	assets.ExecutionSemUS = o.StatsUS.SEM
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

	if err := genDataJS(p+"/stats.js", r); err != nil {
		log.Fatalln("error generating data", err)
	}

	if err := genHTML(p+"/index.html", r); err != nil {
		log.Fatalln("error generating template", err)
	}
}
