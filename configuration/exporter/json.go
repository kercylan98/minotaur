package exporter

import (
	"bytes"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/configuration/raw"
	"os"
)

// NewJSON 创建一个Json导出器
func NewJSON() *JSON {
	return &JSON{}
}

type JSON struct {
	writePath string
}

func (j *JSON) Export(config raw.Config, data any) error {
	buffer := &bytes.Buffer{}
	encoder := jsonIter.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(data)
	if err != nil {
		return err
	}

	var writer = os.Stdout
	_, err = buffer.WriteTo(writer)
	return err
}
