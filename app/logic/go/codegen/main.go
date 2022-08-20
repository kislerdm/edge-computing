//go:build gen
// +build gen

package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
	"text/template"
)

//go:embed data/model.json
var dataModel []byte

//go:embed data/colorsname.csv
var dataName []byte

func genTemplate(t, path string, v interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return template.Must(template.New("").Parse(t)).Execute(f, v)
}

func genColorNameDeps(dir string) error {
	const t = `{{- define "color" -}}{ Name: "{{ .Name }}", R: {{ .R }}, G: {{ .G }}, B: {{ .B }} },{{- end }}package logic

var colorNameLookupTable = []color{
{{- range . }}
	{{ template "color" . }}
{{- end }}
}`

	type color struct {
		Name, R, G, B string
	}

	r := csv.NewReader(bytes.NewReader(dataName))

	rowInd := 0
	var d []color
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		if rowInd == 0 {
			rowInd++
			continue
		}

		d = append(
			d, color{
				Name: row[0],
				R:    row[1],
				G:    row[2],
				B:    row[3],
			},
		)
	}

	return genTemplate(t, dir+"/colornamelookuptable.go", d)
}

func genColorTypeDeps(dir string) error {
	const t = `{{- define "node" -}}{
		ID:        {{ .ID }},
		{{- if .Children }}
		Depth:     {{ .Depth }},
		Feature:   "{{ .Feature }}",
		Threshold: {{ .Threshold }},
		Yes:       {{ .Yes }},
		No:        {{ .No }},
		Missing:   {{ .Missing }},
		Children:  []*node{
		{{- range .Children }}
			{{ template "node" . }}
		{{- end }}
	},
		{{- end }}
		{{- if not .Children }}
		Leaf:      {{ .Leaf }},
		{{- end }}
},
{{- end }}package logic

var colorTypeModel = model{
{{- range . }}
	{{ template "node" . }}
{{- end }}
}`

	type node struct {
		ID        int     `json:"nodeid"`
		Depth     int     `json:"depth,omitempty"`
		Feature   string  `json:"split,omitempty"`
		Threshold float64 `json:"split_condition,omitempty"`
		Yes       int     `json:"yes,omitempty"`
		No        int     `json:"no,omitempty"`
		Missing   int     `json:"missing,omitempty"`
		Leaf      float64 `json:"leaf,omitempty"`
		Children  []*node `json:"children,omitempty"`
	}

	var d []node
	if err := json.NewDecoder(bytes.NewReader(dataModel)).Decode(&d); err != nil {
		return err
	}

	return genTemplate(t, dir+"/colortypemodel.go", d)
}

func main() {
	p, _ := os.Getwd()
	dir := path.Dir(p) + "/logic"

	if err := genColorTypeDeps(dir); err != nil {
		log.Fatalln("cannot generate color type model:", err)
	}

	if err := genColorNameDeps(dir); err != nil {
		log.Fatalln("cannot generate color name lookup table:", err)
	}
}
