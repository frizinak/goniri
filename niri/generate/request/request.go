package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/frizinak/goniri/niri/generate"
)

const fn = "gen.requests.go"

func create() error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(
		f,
		`package niri

import "github.com/frizinak/goniri/niri/types"`)
	f.Close()
	return err
}

func add(name, prop, propType string) error {
	var t = `
type {{ .N }}Request struct {
	r *{{ .N }}Response
}

func Request{{ .U }}() {{ .N }}Request {
	return {{ .N }}Request{r: new({{ .N }}Response)}
}

func (r {{ .N }}Request) MarshalJSON() ([]byte, error) {
	return []byte(` + "`\"{{ .U }}\"`" + `), nil
}

func (r {{ .N }}Request) response() Response {
	return r.r
}

{{ if .Prop -}}
func (r {{ .N }}Request) Response() {{ .PropType }} {
	return r.r.{{ .Prop }}
}
{{ else -}}
func (r {{ .N }}Request) Response() *{{ .N }}Response {
	return r.r
}
{{ end -}}`

	tpl, err := template.New("main").Parse(t)
	if err != nil {
		return err
	}

	data := struct {
		N        string
		U        string
		Prop     string
		PropType string
	}{
		name,
		generate.UCFirst(name),
		prop,
		generate.Type(propType),
	}

	// do not use an underscore or gen_windows.go wont compile (on non-windows)
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return tpl.Execute(f, data)
}

func main() {
	arg := os.Args[1]
	switch arg {
	case "create":
		if err := create(); err != nil {
			panic(err)
		}
	default:
		var prop, propType string
		if len(os.Args) > 2 {
			prop = os.Args[2]
			propType = os.Args[3]
		}
		if err := add(arg, prop, propType); err != nil {
			panic(err)
		}

	}
}
