package codec

import (
	"bytes"
	goGob "encoding/gob"
	"github.com/kercylan98/minotaur/rpc"
)

// NewGob 用于创建一个基于 Gob 的编解码器
func NewGob() rpc.Codec {
	codec := &gob{}
	return codec
}

type gob struct{}

func (c *gob) EncodeData(data any) ([]byte, error) {
	var buf bytes.Buffer
	if err := goGob.NewEncoder(&buf).Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *gob) DecodeData(data []byte, dst any) error {
	return goGob.NewDecoder(bytes.NewReader(data)).Decode(dst)
}

func (c *gob) EncodeRequest(req *rpc.Request) ([]byte, error) {
	var buf bytes.Buffer
	if err := goGob.NewEncoder(&buf).Encode(req); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *gob) DecodeRequest(data []byte, dst *rpc.Request) error {
	return goGob.NewDecoder(bytes.NewReader(data)).Decode(dst)
}
