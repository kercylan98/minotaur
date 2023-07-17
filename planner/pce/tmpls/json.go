package tmpls

import (
	"bytes"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/utils/str"
)

func NewJSON() *JSON {
	return &JSON{
		API: jsonIter.ConfigCompatibleWithStandardLibrary,
	}
}

type JSON struct {
	jsonIter.API
}

func (slf *JSON) Render(data map[any]any) (string, error) {
	buffer := &bytes.Buffer{}
	encoder := jsonIter.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(data)
	if err != nil {
		return str.None, err
	}
	return buffer.String(), nil
}
