package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	// WebsocketMessageTypeText 表示文本数据消息。文本消息负载被解释为 UTF-8 编码的文本数据
	WebsocketMessageTypeText = 1
	// WebsocketMessageTypeBinary 表示二进制数据消息
	WebsocketMessageTypeBinary = 2
	// WebsocketMessageTypeClose 表示关闭控制消息。可选消息负载包含数字代码和文本。使用 FormatCloseMessage 函数来格式化关闭消息负载
	WebsocketMessageTypeClose = 8
	// WebsocketMessageTypePing 表示 ping 控制消息。可选的消息负载是 UTF-8 编码的文本
	WebsocketMessageTypePing = 9
	// WebsocketMessageTypePong 表示一个 pong 控制消息。可选的消息负载是 UTF-8 编码的文本
	WebsocketMessageTypePong = 10
)

type Option func(srv *Server)

// WithTLS 通过安全传输层协议TLS创建服务器
//   - 支持：Http、Websocket
func WithTLS(certFile, keyFile string) Option {
	return func(srv *Server) {
		switch srv.network {
		case NetworkHttp, NetworkWebsocket, NetworkTCP, NetworkTCP4, NetworkTCP6:
			srv.certFile = certFile
			srv.keyFile = keyFile
		}
	}
}

// WithGRPCServerOptions 通过GRPC的可选项创建GRPC服务器
func WithGRPCServerOptions(options ...grpc.ServerOption) Option {
	return func(srv *Server) {
		if srv.network != NetworkGRPC {
			return
		}
		srv.grpcServer = grpc.NewServer(options...)
	}
}

// WithProd 通过生产模式运行服务器
func WithProd() Option {
	return func(srv *Server) {
		srv.prod = true
	}
}

// WithWebsocketMessageType 设置仅支持特定类型的Websocket消息
func WithWebsocketMessageType(messageTypes ...int) Option {
	return func(srv *Server) {
		if srv.network != NetworkWebsocket {
			log.Warn("WitchWebsocketMessageType", zap.String("Network", string(srv.network)), zap.Error(ErrNotWebsocketUseMessageType))
			return
		}
		var supports = make(map[int]bool)
		for _, messageType := range messageTypes {
			switch messageType {
			case WebsocketMessageTypeText, WebsocketMessageTypeBinary, WebsocketMessageTypeClose, WebsocketMessageTypePing, WebsocketMessageTypePong:
				supports[messageType] = true
			default:
				log.Warn("WitchWebsocketMessageType", zap.Int("MessageType", messageType), zap.Error(ErrWebsocketMessageTypeException))
			}
		}
		srv.supportMessageTypes = supports
	}
}

// WithMessageBufferSize 通过特定的消息缓冲池大小运行服务器
//   - 默认大小为 4096 * 1024
//   - 消息数量超出这个值的时候，消息处理将会造成更大的开销（频繁创建新的结构体），同时服务器将输出警告内容
func WithMessageBufferSize(size int) Option {
	return func(srv *Server) {
		if size <= 0 {
			return
		}
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
