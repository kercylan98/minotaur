package rpc

import (
	"fmt"
	"net"
)

// Server 是一个 RPC 服务端的接口，该接口用于定义一个 RPC 服务端，用于接收 RPC 调用
type Server interface {

	// GetRouter 获取路由器
	GetRouter() Router

	// OnReceived 用于接收 RPC 调用消息
	OnReceived(request []byte) ([]byte, error)

	// ListenAndServe 用于启动一个 RPC 服务端
	ListenAndServe(network, address string) error

	// Serve 用于启动一个 RPC 服务端
	Serve(l net.Listener) error
}

// NewServer 用于创建一个新的 RPC 服务端
func NewServer(transporter Transporter, router Router, codec Codec) Server {
	srv := &server{
		codec:  codec,
		router: router,
		trans:  transporter,
	}
	transporter.OnInit(srv)
	return srv
}

type server struct {
	codec  Codec
	router Router
	trans  Transporter
}

func (s *server) GetRouter() Router {
	return s.router
}

func (s *server) ListenAndServe(network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	return s.trans.Serve(l)
}

func (s *server) Serve(l net.Listener) error {
	return s.trans.Serve(l)
}

func (s *server) OnReceived(input []byte) ([]byte, error) {
	var req = new(Request)
	if err := s.codec.DecodeRequest(input, req); err != nil {
		return nil, err
	}

	handler := s.router.Match(req.Route)
	if handler == nil {
		return nil, fmt.Errorf("rpc: route %s not found", req.Route)
	}

	ctx := newContext(func(dst any) error {
		return s.codec.DecodeData(req.Data, dst)
	})
	if err := handler(ctx); err != nil {
		return nil, err
	}
	if ctx.reply == nil {
		return nil, nil
	}

	replyData, err := s.codec.EncodeData(ctx.reply)
	if err != nil {
		return nil, err
	}
	return replyData, nil
}
