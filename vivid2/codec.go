package vivid

import (
	"bytes"
	"encoding/gob"
)

type Codec interface {
	Encode(v interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}

type _GobCodec struct{}

func (c *_GobCodec) Encode(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *_GobCodec) Decode(data []byte, v interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(v)
}
