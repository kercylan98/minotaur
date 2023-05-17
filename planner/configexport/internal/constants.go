package internal

var TemplateGo = `// Code generated DO NOT EDIT.

package %s

import (
	jsonIter "github.com/json-iterator/go"
	"os"
)

var json = jsonIter.ConfigCompatibleWithStandardLibrary

%s

func LoadConfig(handle func(filename string, config any) error) {
	%s
}

func Replace() {
	%s
}


func DefaultLoad(filepath string) {
	LoadConfig(func(filename string, config any) error {
		bytes, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		return json.Unmarshal(bytes, &config)
	})
}

`

var TemplateStructGo = `// Code generated DO NOT EDIT.

package %s

%s
`
