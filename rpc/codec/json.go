package codec

import (
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/toolkit"
)

// NewJSON 用于创建一个基于 JSON 的编解码器
func NewJSON() rpc.Codec {
	codec := &json{}
	return codec
}

type json struct{}

func (c *json) EncodeData(data any) ([]byte, error) {
	return toolkit.MarshalJSONE(data)
}

func (c *json) DecodeData(data []byte, dst any) error {
	return toolkit.UnmarshalJSONE(data, dst)
}

func (c *json) EncodeRequest(req *rpc.Request) ([]byte, error) {
	return toolkit.MarshalJSONE(req)
}

func (c *json) DecodeRequest(data []byte, dst *rpc.Request) error {
	return toolkit.UnmarshalJSONE(data, dst)
}
