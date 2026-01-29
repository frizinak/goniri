package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/frizinak/goniri/ipc/generate"
)

const fn = "gen.actions.go"

func create() error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(
		f,
		`package ipc

import "github.com/frizinak/goniri/ipc/types"`)
	f.Close()
	return err
}

func simple(name string) error {
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(
		f,
		`
func Action%s() Request {
	return newSimpleAction(%#v)
}
`,
		name,
		[]byte(fmt.Sprintf("{\"Action\":{\"%s\":{}}}", name)),
	)
	f.Close()
	return err
}

func args(name string, args []string) error {
	var t = `
func Action{{ .Name }}(
{{- range .Fields }}
	{{ .Value }}{{ if .Type }} {{ .Type }}{{ end }},
{{- end }}
) Request {
	return newAction(map[string]map[string]{{ .CommonType }}{
		"{{ .Name }}": {
			{{- range .Fields }}
			"{{ .Key }}": {{ .Value }},
			{{- end }}
		},
	})
}
`

	tpl, err := template.New("main").Parse(t)
	if err != nil {
		return err
	}

	type field struct {
		Key, Value, Type string
	}

	type global struct {
		Name       string
		CommonType string
		Fields     []field
	}

	data := global{}
	data.Name = name
	data.CommonType = ""
	for _, arg := range args {
		p := strings.SplitN(arg, ":", 3)
		if len(p) == 2 {
			n := make([]string, 3)
			n[0], n[1], n[2] = p[0], p[0], p[1]
			p = n
		}

		f := field{
			Key:   p[0],
			Value: generate.Dash2Camel(p[1]),
			Type:  generate.Type(p[2]),
		}
		if data.CommonType == "" {
			data.CommonType = f.Type
		}
		if data.CommonType != f.Type {
			data.CommonType = "any"
		}
		data.Fields = append(data.Fields, f)
	}

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return tpl.Execute(f, data)
}

func main() {
	arg := os.Args[1]
	switch {
	case arg == "create":
		if err := create(); err != nil {
			panic(err)
		}
	case len(os.Args) == 2:
		if err := simple(arg); err != nil {
			panic(err)
		}
	default:
		if err := args(arg, os.Args[2:]); err != nil {
			panic(err)
		}

	}
}
