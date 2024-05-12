package rpc

// Context 是一个 RPC 调用上下文的接口，该接口用于在 RPC 调用过程中传递上下文信息
type Context interface {
	// ReadTo 对于被调用方，ReadTo 方法用于将上下文信息读取到指定的结构体中
	ReadTo(dst any) error

	// MustReadTo 对于被调用方，MustReadTo 方法用于将上下文信息读取到指定的结构体中，如果读取失败则会 panic
	MustReadTo(dst any)

	// Write 对于调用方，Write 方法用于将上下文信息写入到上下文中
	Write(src any) error
}
