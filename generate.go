// +build ignore

package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Table struct {
	XMLName  xml.Name `xml:"table"`
	Elements []*struct {
		Name string `xml:"name,attr"`
		ID   string `xml:"id,attr"`
		Type string `xml:"type,attr"`
	} `xml:"element"`
}

func main() {
	log.Printf("Downloading specification XML ...")
	resp, err := http.Get("https://github.com/Matroska-Org/foundation-source/raw/master/spectool/specdata.xml")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	// data, err := ioutil.ReadFile("specdata.xml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Generating elements.go ...")
	table := Table{}
	err = xml.Unmarshal(data, &table)

	for _, v := range table.Elements {
		v.Name = strings.Replace(v.Name, "-", "", -1)
	}

	elementsFile, err := os.Create("elements.go")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer elementsFile.Close()

	elementsTemplate.Execute(elementsFile, table)
}

var elementsTemplate = template.Must(template.New("").Parse(`// GENERATED FILE. DO NOT EDIT.

package mkvparse

// Supported ElementIDs
const (
	{{- range .Elements }}
	{{ .Name }}Element ElementID = {{ .ID -}}
	{{end }}
)

var elementTypes = map[ElementID]elementType {
	{{- range .Elements }}
	{{ .Name }}Element: 
		{{- if eq .Type "master" -}}
			masterType
		{{- else if eq .Type "uinteger" -}}
			uintegerType
		{{- else if eq .Type "integer" -}}
			integerType
		{{- else if eq .Type "binary" -}}
			binaryType
		{{- else if eq .Type "utf-8" -}}
			utf8Type
		{{- else if eq .Type "string" -}}
			stringType
		{{- else if eq .Type "float" -}}
			floatType
		{{- else if eq .Type "date" -}}
			dateType
		{{- end -}},
	{{- end -}}
}

var elementNames = map[ElementID]string {
	{{- range .Elements }}
	{{ .Name }}Element: {{ printf "%q" .Name }},
	{{- end -}}
}

`))
