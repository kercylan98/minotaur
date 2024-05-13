package vivid

import (
	"bytes"
	goGob "encoding/gob"
)

type Codec interface {
	Encode(data any) ([]byte, error)

	Decode(data []byte, dst any) error
}

type gobCodec struct{}

func (c *gobCodec) Encode(data any) ([]byte, error) {
	var buf bytes.Buffer
	if err := goGob.NewEncoder(&buf).Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *gobCodec) Decode(data []byte, dst any) error {
	return goGob.NewDecoder(bytes.NewReader(data)).Decode(dst)
}
