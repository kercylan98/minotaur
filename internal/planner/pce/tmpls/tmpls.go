package tmpls

import (
	"bytes"
	"github.com/kercylan98/minotaur/utils/str"
	"text/template"
)

func render(temp string, o any) (string, error) {
	tmpl, err := template.New(temp).Parse(temp)
	if err != nil {
		return str.None, err
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, o); err != nil {
		return str.None, err
	}
	return buf.String(), nil
}
