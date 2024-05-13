package transporter

import (
	"github.com/kercylan98/minotaur/rpc"
	"net"
	goRPC "net/rpc"
)

const (
	goRPCName  = "Minotaur"
	GoRPCRoute = "Minotaur.OnHandle"
)

// NewGoRPC 创建一个基于 net/rpc 的传输器
func NewGoRPC() rpc.Transporter {
	t := &GoRpc{
		srv: goRPC.NewServer(),
	}
	return t
}

type GoRpc struct {
	srv    *goRPC.Server
	server rpc.Server
}

func (g *GoRpc) OnInit(srv rpc.Server) {
	g.server = srv
}

func (g *GoRpc) Serve(l net.Listener) error {
	if err := g.srv.RegisterName(goRPCName, g); err != nil {
		return err
	}

	g.srv.Accept(l)
	return nil
}

func (g *GoRpc) OnHandle(request []byte, response *[]byte) error {
	reply, err := g.server.OnReceived(request)
	if err != nil {
		return err
	}

	*response = reply
	return nil
}
