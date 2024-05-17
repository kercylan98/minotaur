package server

import (
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/vivids"
)

func NewServer(network Network) *Server {
	return &Server{
		network: network,
	}
}

type Server struct {
	vivid.BasicActor
	network Network
}

func (s *Server) OnPreStart(ctx vivids.ActorContext) (err error) {
	_, err = ctx.ActorOf(s.network, vivids.NewActorOptions().WithName("network"))
	return err
}

func (s *Server) OnReceived(ctx vivids.MessageContext) (err error) {
	switch e := ctx.GetMessage().(type) {
	case NetworkConnectionOpenedMessage:
		return s.onNetworkConnectionOpened(ctx, e)
	case NetworkConnectionClosedMessage: // 向连接 Actor 转发关闭消息
		return e.Conn.Tell(e, vivids.WithMessageSender(ctx))
	}

	return
}

func (s *Server) onNetworkConnectionOpened(ctx vivids.MessageContext, e NetworkConnectionOpenedMessage) error {
	// 为连接创建一个新的 Actor，当创建失败时回应错误消息
	// 当回复失败时，调用方考虑超时机制来决定本次操作失败
	conn, err := ctx.ActorOf(newConn(e.Conn, e.ConnectionWriter), vivids.NewActorOptions().WithName(e.Conn.RemoteAddr().String()))
	return ctx.Reply(toolkit.NotNil(conn, err))
}
