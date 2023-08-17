package server

// NewHttpHandleWrapper 创建一个新的 http 处理程序包装器
//   - 默认使用 server.HttpContext 作为上下文，如果需要依赖其作为新的上下文，可以通过 NewHttpContext 创建
func NewHttpHandleWrapper[Context any](srv *Server, packer ContextPacker[Context]) *Http[Context] {
	return &Http[Context]{
		HttpRouter: &HttpRouter[Context]{
			srv:    srv,
			group:  srv.ginServer,
			packer: packer,
		},
	}
}

// Http 基于 gin.Engine 包装的 http 服务器
type Http[Context any] struct {
	srv *Server
	*HttpRouter[Context]
}
