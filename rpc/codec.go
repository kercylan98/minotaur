package rpc

// Codec 是一个 RPC 消息编解码器的接口，该接口用于定义一个 RPC 消息编解码器
type Codec interface {
	RequestCodec
	DataCodec
}

// RequestCodec 是一个 RPC 消息上下文编解码器的接口，该接口用于定义一个 RPC 消息上下文编解码器
type RequestCodec interface {
	// EncodeRequest 用于将一个 RPC 消息上下文编码为一个字节数据
	EncodeRequest(req *Request) ([]byte, error)

	// DecodeRequest 用于将一个字节数据解码为一个 RPC 消息上下文
	DecodeRequest(data []byte, dst *Request) error
}

// DataCodec 是一个 RPC 消息数据编解码器的接口，该接口用于定义一个 RPC 消息数据编解码器
type DataCodec interface {
	// EncodeData 用于将一个 RPC 消息数据编码为一个字节数据
	EncodeData(data any) ([]byte, error)

	// DecodeData 用于将一个字节数据解码为一个 RPC 消息数据
	DecodeData(data []byte, dst any) error
}
