package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
	"unicode"
)

func do(raw, prop, propType string) error {
	var snake string
	var ucfirst string
	{
		a := []rune(raw)
		b := make([]rune, 0, len(a)+5)
		for i, ch := range a {
			if unicode.IsUpper(ch) {
				if i != 0 {
					b = append(b, '_')
				}
				ch = unicode.ToLower(ch)
			}
			b = append(b, ch)
		}
		c := make([]rune, len(a))
		copy(c, a)
		c[0] = unicode.ToUpper(c[0])
		snake, ucfirst = string(b), string(c)
	}

	var t = `package niri

{{ if .Import -}}
import "github.com/frizinak/goniri/niri/types"

{{ end -}}
type {{ .N }}Request struct {
	r *{{ .N }}Response
}

func New{{ .U }}Request() {{ .N }}Request {
	return {{ .N }}Request{r: &{{ .N }}Response{}}
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
		Import   bool
	}{raw, ucfirst, prop, propType, strings.Contains(propType, "types.")}

	// do not use an underscore or gen_windows.go wont compile (on non-windows)
	f, err := os.Create(fmt.Sprintf("gen.%s.go", snake))
	if err != nil {
		return err
	}
	defer f.Close()
	return tpl.Execute(f, data)
}

func main() {
	flag.Parse()
	if err := do(flag.Arg(0), flag.Arg(1), flag.Arg(2)); err != nil {
		panic(err)
	}
}
