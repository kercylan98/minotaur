package server

type Option func(srv *Server)

// WithProd 通过生产模式运行服务器
func WithProd() Option {
	return func(srv *Server) {
		srv.prod = true
	}
}

// WithMessageBufferSize 通过特定的消息缓冲池大小运行服务器
//   - 默认大小为 1024
//   - 消息数量超出这个值的时候，消息处理将会造成更大的开销（频繁创建新的结构体），同时服务器将输出警告内容
func WithMessageBufferSize(size int) Option {
	return func(srv *Server) {
		srv.messagePoolSize = size
	}
}
