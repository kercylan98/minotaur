package generator

import (
	"bytes"
	_ "embed"
	"github.com/kercylan98/minotaur/engine/vivid/typed/options"
	"strings"
	"text/template"
)

//go:embed typed.tmpl
var typedTemplate string

type typedServiceDesc struct {
	Name     string
	ModeName string
	Methods  []*typedServiceMethodDesc
}

type typedServiceMethodDesc struct {
	Name    string
	Input   string
	Output  string
	Index   int
	Options *options.MethodOptions
}

func (d *typedServiceDesc) render() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("typed").Parse(strings.TrimSpace(typedTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, d); err != nil {
		panic(err)
	}

	return strings.Trim(buf.String(), "\r\n")
}
