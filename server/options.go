package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
)

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

// WithMultiCore 通过特定核心数量运行服务器，默认为单核
//   - count > 1 的情况下，将会有对应数量的 goroutine 来处理消息
//   - 注意：HTTP和GRPC网络模式下不会生效
func WithMultiCore(count int) Option {
	return func(srv *Server) {
		srv.core = count
		if srv.core < 1 {
			log.Warn("WithMultiCore", zap.Int("count", count), zap.String("tips", "wrong core count configuration, corrected to 1, currently in single-core mode"))
			srv.core = 1
		}
	}
}
