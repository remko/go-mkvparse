// +build ignore

package main

import (
	"bytes"
	"encoding/xml"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Table struct {
	XMLName         xml.Name   `xml:"table"`
	Elements        []*Element `xml:"element"`
	ChildHierarchy  map[string][]string
	ParentHierarchy map[string]string
}

type Element struct {
	Name  string `xml:"name,attr"`
	Level int    `xml:"level,attr"`
	ID    string `xml:"id,attr"`
	Type  string `xml:"type,attr"`
}

// The specification is ordered in a way that, given consecutive elements A and
// B:
//
// * If A.Level < B.Level, then B is a child of A.
// * If A.Level == B.Level, then B is a sibling of A.
// * If A.Level > B.Level, then B does not have a direct relationship with A.
//
// Additionally, if the level is less than 0, then it means the element is
// allowed as a child anywhere. This is not represented in the resulting
// hierarchy.
//
// This function will compute direct child and parent mappings.
func (t *Table) ComputeHierarchy() (map[string][]string, map[string]string) {
	stack := []Element{}

	chierarchy := make(map[string][]string)
	phierarchy := make(map[string]string)

	for _, v := range t.Elements {
		// Compute effective levels to handle -1. Negative level will
		// be handled as if the element was level 0.
		v_level := v.Level
		if v_level < 0 {
			v_level = 0
		}

		var p_level int
		if len(stack) > 0 {
			p_level = stack[len(stack)-1].Level
			if p_level < 0 {
				p_level = 0
			}
		}

		if len(stack) == 0 {
			stack = append(stack, *v)
		} else if p_level < v_level {
			stack = append(stack, *v)
		} else if p_level == v_level {
			stack[len(stack)-1] = *v
		} else {
			for len(stack) != 0 && p_level >= v_level {
				stack = stack[:len(stack)-1]

				if len(stack) > 0 {
					p_level = stack[len(stack)-1].Level
					if p_level < 0 {
						p_level = 0
					}
				}
			}
			stack = append(stack, *v)
		}

		if len(stack) >= 2 {
			p := stack[len(stack)-2]
			c := stack[len(stack)-1]

			ch, ok := chierarchy[p.Name]
			if !ok {
				ch = []string{}
			}

			ch = append(ch, c.Name)
			chierarchy[p.Name] = ch

			phierarchy[c.Name] = p.Name
		}
	}

	t.ChildHierarchy = chierarchy
	t.ParentHierarchy = phierarchy

	return chierarchy, phierarchy
}

func main() {
	var data []byte
	var err error

	// If a spec wasn't provided, download it. Otherwise read it from the
	// path provided.
	if len(os.Args) == 1 {
		log.Printf("Downloading specification XML ...")

		resp, err := http.Get("https://github.com/Matroska-Org/foundation-source/raw/master/spectool/specdata.xml")
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer resp.Body.Close()

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else if len(os.Args) == 2 {
		data, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		log.Fatalf("Invalid number of args: %d", len(os.Args))
	}

	log.Printf("Generating elements.go ...")

	table := Table{}
	err = xml.Unmarshal(data, &table)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Normalize element names to Go conventions.
	for _, v := range table.Elements {
		v.Name = strings.Replace(v.Name, "-", "", -1)
	}

	table.ComputeHierarchy()

	elementsFile, err := os.Create("elements.go")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer elementsFile.Close()

	var buf bytes.Buffer

	err = elementsTemplate.Execute(&buf, table)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmtd, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("%v", err)
	}

	_, err = elementsFile.Write(fmtd)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

var elementsTemplate = template.Must(template.New("").Parse(`// GENERATED FILE. DO NOT EDIT.

package mkvparse

// Supported ElementIDs
const (
	{{- range .Elements }}
	{{ .Name }}Element = {{ .ID -}}
	{{ end }}
)

var elementTypes = map[ElementID]ElementType {
	{{- range .Elements }}
	{{ .Name }}Element: {{ if eq .Type "master" -}}
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
	{{- end }}
}

var elementNames = map[ElementID]string {
	{{- range .Elements }}
	{{ .Name }}Element: {{ printf "%q" .Name }},
	{{- end }}
}

var elementLevels = map[ElementID]int {
	{{- range .Elements }}
	{{ .Name }}Element: {{ .Level }},
	{{- end }}
}

var elementChildren = map[ElementID]map[ElementID]bool {
	{{- range $key, $value := .ChildHierarchy }}
	{{ $key }}Element: map[ElementID]bool{
		{{- range $value }}
		{{ . }}Element: true,
		{{- end }}
	},
	{{- end }}
}

var elementParent = map[ElementID]ElementID {
	{{- range $key, $value := .ParentHierarchy }}
	{{ $key }}Element: {{ $value }}Element,
	{{- end }}
}
`))
