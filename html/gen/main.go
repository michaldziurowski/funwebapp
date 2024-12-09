package main

import (
	"encoding/json"
	"os"
	"strings"
	"text/template"
)

type Element struct {
	Attributes []string `json:"attributes"`
}

func main() {
	content, err := os.ReadFile("html-elements.json")
	if err != nil {
		panic(err)
	}

	f := map[string]*Element{}
	err = json.Unmarshal(content, &f)
	if err != nil {
		panic(err)
	}

	for k, v := range f {
		if k != "*" {
			for _, av := range f["*"].Attributes {
				v.Attributes = append(v.Attributes, av)
			}
		}
	}

	funcsMap := template.FuncMap{
		"title": strings.Title,
		"removedash": func(s string) string {
			return strings.ReplaceAll(s, "-", "")
		},
	}
	t, err := template.New("html.gotmpl").Funcs(funcsMap).ParseFiles("html.gotmpl")
	if err != nil {
		panic(err)
	}

	t.Execute(os.Stdout, f)
}
