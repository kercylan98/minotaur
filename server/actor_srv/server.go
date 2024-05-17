package server

import (
	"context"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/vivid"
)

// New 创建一个新的服务器
func New() *Server {
	return &Server{}
}

type Server struct {
	ctx     context.Context
	network server.Network
}

func (s *Server) OnPreStart(ctx vivid.ActorContext) (err error) {
	if err = s.network.OnSetup(s.ctx, nil); err != nil {
		return
	}

	ctx.Future(func() vivid.Message {
		return s.network.OnRun
	})
	return
}

func (s *Server) OnReceived(ctx vivid.MessageContext) (err error) {
	switch m := ctx.GetMessage().(type) {
	case error:
		if err = s.network.OnShutdown(); err != nil {
			panic(err)
		}
		ctx.NotifyTerminated(m)
	}

	return
}

func (s *Server) OnDestroy(ctx vivid.ActorContext) (err error) {

}

func (s *Server) OnSaveSnapshot(ctx vivid.ActorContext) (snapshot []byte, err error) {

}

func (s *Server) OnRecoverSnapshot(ctx vivid.ActorContext, snapshot []byte) (err error) {

}

func (s *Server) OnChildTerminated(ctx vivid.ActorContext, child vivid.ActorTerminatedContext) {

}
