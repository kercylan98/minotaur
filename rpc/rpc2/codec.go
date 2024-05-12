package rpc

// Codec 是一个 RPC 编解码器的接口，该接口用于定义一个 RPC 编解码器，用于在 RPC 调用过程中对数据进行编解码
type Codec interface {
	// Encode 用于对数据进行编码
	Encode(data any) ([]byte, error)

	// Decode 用于对数据进行解码
	Decode(data []byte, dst any) error
}
