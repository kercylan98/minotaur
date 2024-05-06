package rpc

type (
	// UnaryHandler 一对一调用处理器，该处理器的返回值将作为响应返回给调用方
	UnaryHandler func(reader Reader) any
	// UnaryNotifyHandler 一对一调用处理器，该处理器的调用方不需要响应
	UnaryNotifyHandler func(reader Reader)
)

// Route 表示一个 RPC 路由
type Route = string

// RouteBinder 表示一个 RPC 路由绑定器
type RouteBinder struct {
	route       []Route
	application *Application
}

// Unary 为指定的路由绑定一对一调用处理器
func (b RouteBinder) Unary(handler UnaryHandler) {
	b.application.router.Register(b.route...).Bind(handler)
}

// UnaryNotify 为指定的路由绑定一对一调用处理器
func (b RouteBinder) UnaryNotify(handler UnaryNotifyHandler) {
	b.application.router.Register(b.route...).Bind(handler)
}
