package network

import (
	"context"
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	server "github.com/kercylan98/minotaur/server3"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/panjf2000/gnet/v2"
	"time"
)

var (
	schemaWebSocket = "ws"
	schemaTcp       = "tcp"
	schemaTcp4      = "tcp4"
	schemaTcp6      = "tcp6"
	schemaUdp       = "udp"
	schemaUdp4      = "udp4"
	schemaUdp6      = "udp6"
	schemaUnix      = "unix"
)

func newGnetEngine(schema, addr string, pattern ...string) server.Network {
	g := &gnetEngine{
		addr:    addr,
		schema:  schema,
		pattern: collection.FindFirstOrDefaultInSlice(pattern, "/"),
	}
	return g
}

type gnetEngine struct {
	addr     string
	schema   string
	pattern  string
	eng      gnet.Engine
	upgrader ws.Upgrader
	srv      server.Core
}

func (g *gnetEngine) Launch(ctx context.Context, srv server.Core) error {
	g.srv = srv

	var addr string
	switch g.schema {
	case schemaTcp, schemaWebSocket:
		addr = fmt.Sprintf("tcp://%s", g.addr)
		if g.schema == schemaWebSocket {
			g.initWebSocketUpgrader()
		}
	case schemaTcp4:
		addr = fmt.Sprintf("tcp4://%s", g.addr)
	case schemaTcp6:
		addr = fmt.Sprintf("tcp6://%s", g.addr)
	case schemaUdp:
		addr = fmt.Sprintf("udp://%s", g.addr)
	case schemaUdp4:
		addr = fmt.Sprintf("udp4://%s", g.addr)
	case schemaUdp6:
		addr = fmt.Sprintf("udp6://%s", g.addr)
	case schemaUnix:
		addr = fmt.Sprintf("unix://%s", g.addr)
	default:
		return fmt.Errorf("unsupported schema: %s", g.schema)
	}
	return gnet.Run(g, addr)
}

func (g *gnetEngine) Shutdown() error {
	return g.eng.Stop(context.TODO())
}

func (g *gnetEngine) Schema() string {
	return g.schema
}

func (g *gnetEngine) Address() string {
	return g.addr
}

func (g *gnetEngine) OnBoot(eng gnet.Engine) (action gnet.Action) {
	g.eng = eng
	return
}

func (g *gnetEngine) OnShutdown(eng gnet.Engine) {

}

func (g *gnetEngine) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	if g.schema == schemaWebSocket {
		c.SetContext(newWebsocketWrapper(c))
	} else {
		connActor := g.srv.BindConn(c, func(packet server.Packet) error {
			return c.AsyncWrite(packet.GetBytes(), func(c gnet.Conn, err error) error {
				return nil
			})
		})

		if connActor == nil {
			action = gnet.Close
			return
		}

		c.SetContext(connActor)
	}
	return
}

func (g *gnetEngine) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	g.srv.UnbindConn(c)
	return
}

func (g *gnetEngine) OnTraffic(c gnet.Conn) (action gnet.Action) {
	if g.schema == schemaWebSocket {
		wrapper := c.Context().(*websocketWrapper)

		if err := wrapper.readToBuffer(); err != nil {
			return gnet.Close
		}

		if err := wrapper.upgrade(g.upgrader, func() {
			// 协议升级成功后视为连接建立
			conn := g.srv.BindConn(c, func(packet server.Packet) error {
				return wsutil.WriteServerMessage(c, ws.OpText, packet.GetBytes())
			})

			if conn == nil {
				action = gnet.Close
				return
			}

			wrapper.ref = conn

		}); err != nil {
			return gnet.Close
		}
		wrapper.active = time.Now()

		// decode
		messages, err := wrapper.decode()
		if err != nil {
			return gnet.Close
		}

		for _, message := range messages {
			packet := server.NewPacket(message.Payload)
			packet.SetContext(message.OpCode)
			g.srv.ProcessPacket(wrapper.ref, packet)
		}
	} else {
		buf, err := c.Next(-1)
		if err != nil {
			return gnet.Close
		}

		var clone = make([]byte, len(buf))
		copy(clone, buf)

		g.srv.ProcessPacket(c.Context().(vivid.ActorRef), server.NewPacket(clone))
	}
	return
}

func (g *gnetEngine) OnTick() (delay time.Duration, action gnet.Action) {
	return
}

func (g *gnetEngine) initWebSocketUpgrader() {
	g.upgrader = ws.Upgrader{
		OnRequest: func(uri []byte) (err error) {
			if string(uri) != g.pattern {
				err = errors.New("bad request")
			}
			return
		},
	}
}
