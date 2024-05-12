package rpc

// Message 是一个 RPC 消息的接口，该接口用于定义一个 RPC 消息，用于在 RPC 调用过程中传递
type Message interface {
	// GetRoute 用于获取消息的路由
	GetRoute() Route

	// GetBytes 用于获取消息的字节数据
	GetBytes() []byte
}
