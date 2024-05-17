package unsafevivid

import (
	"bytes"
	goGob "encoding/gob"
)

var gob = new(gobCodec)

// Codec 用于远程消息的编解码器的接口
type Codec interface {
	// Encode 用于编码一个消息
	Encode(v any) ([]byte, error)

	// Decode 用于解码一个消息
	Decode(data []byte, v any) error
}

type gobCodec struct{}

func (c *gobCodec) Encode(v any) ([]byte, error) {
	var buf bytes.Buffer
	if err := goGob.NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *gobCodec) Decode(data []byte, v any) error {
	return goGob.NewDecoder(bytes.NewReader(data)).Decode(v)
}
