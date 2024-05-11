package exporter

import (
	"bytes"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/configuration/raw"
	"os"
)

// NewJSON 创建一个Json导出器
func NewJSON(writePath string) *JSON {
	return &JSON{
		writePath: writePath,
	}
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
	if j.writePath != "" {
		writer, err = os.OpenFile(j.writePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
	}

	_, err = writer.Write(buffer.Bytes())
	return err
}
