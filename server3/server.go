package server

import (
	"context"
	"github.com/kercylan98/minotaur/vivid"
)

type Server interface {
	vivid.Actor
	// Launch 启动服务器
	Launch(ctx ...context.Context) error
}

// NewServer 创建一个新的 Server
func NewServer(network Network) Server {
	return &_Server{
		network: network,
	}
}

// NewServerActorProps 创建一个新的 ServerActor
func NewServerActorProps(network Network) vivid.ActorOption[*_Server] {
	return func(opts *vivid.ActorOptions[*_Server]) {
		opts.WithProps(func(server *_Server) {
			server.network = network
		})
	}
}

type _Server struct {
	ctx     context.Context    // 上下文
	cancel  context.CancelFunc // 取消函数
	network Network            // 网络接口
}

func (s *_Server) OnReceive(ctx vivid.MessageContext) {
	switch ctx.GetMessage().(type) {
	case LaunchServerMessage:
		go s.Launch(ctx)
	}
}

func (s *_Server) Launch(ctx ...context.Context) (err error) {
	if len(ctx) == 0 {
		ctx = append(ctx, context.Background())
	}
	s.ctx, s.cancel = context.WithCancel(ctx[0])
	defer s.cancel()
	if err = s.network.OnSetup(s.ctx); err != nil {
		return
	}

	if err = s.network.OnRun(); err != nil {
		return
	}

	<-s.ctx.Done()
	return
}
