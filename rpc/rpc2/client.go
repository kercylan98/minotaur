package rpc

// Client 是一个 RPC 客户端的接口，该接口用于定义一个 RPC 客户端，用于发起 RPC 调用
type Client interface {
	// Tell 用于向指定的服务发起一个 RPC 调用，该调用不需要返回值
	Tell(route Route, data any) error

	// Ask 用于向指定的服务发起一个 RPC 调用，该调用需要返回值
	Ask(route Route, data any) (Context, error)
}
