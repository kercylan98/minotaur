package network

import (
	"context"
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/kercylan98/minotaur/server/actor_srv/server"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/vivids"
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
	vivid.BasicActor
	addr     string
	schema   string
	pattern  string
	eng      gnet.Engine
	upgrader ws.Upgrader
	ctx      vivids.ActorContext
}

func (g *gnetEngine) OnPreStart(ctx vivids.ActorContext) (err error) {
	g.ctx = ctx
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

	ctx.Future(func() vivids.Message {
		return gnet.Run(g, addr)
	})

	return
}

func (g *gnetEngine) OnReceived(ctx vivids.MessageContext) (err error) {
	switch v := ctx.GetMessage().(type) {
	case error:
		ctx.NotifyTerminated(v)
	}

	return
}

func (g *gnetEngine) OnDestroy(ctx vivids.ActorContext) (err error) {
	return g.eng.Stop(context.TODO())
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
		connActor, err := g.ctx.ActorOf(new(server.Conn), vivids.NewActorOptions().
			WithName(c.RemoteAddr().String()).
			WithProps(server.ConnProps{
				Conn: c,
				Writer: func(packet server.Packet) error {
					return c.AsyncWrite(packet.GetBytes(), func(c gnet.Conn, err error) error {
						return g.ctx.GetParentActor().Tell(server.ConnectionAsyncWriteErrorEvent{
							Error: err,
						})
					})
				},
			}))
		if err != nil {
			action = gnet.Close
			return
		}

		g.ctx.PublishEvent(server.NetworkConnectionOpenedEvent{
			ActorRef: connActor,
		})

		c.SetContext(connActor)
	}
	return
}

func (g *gnetEngine) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	var conn vivids.ActorRef
	if g.schema == schemaWebSocket {
		conn = c.Context().(*websocketWrapper).ref
	} else {
		conn = c.Context().(vivids.ActorRef)
	}

	_ = g.ctx.GetParentActor().Tell(server.NetworkConnectionClosedEvent{
		ActorRef: conn,
	}, vivids.WithMessageSender(g.ctx))
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
			conn, err := g.ctx.ActorOf(new(server.Conn), vivids.NewActorOptions().
				WithName(c.RemoteAddr().String()).
				WithProps(server.ConnProps{
					Conn: c,
					Writer: func(packet server.Packet) error {
						return wsutil.WriteServerMessage(c, ws.OpText, packet.GetBytes())
					},
				}))
			if err != nil {
				action = gnet.Close
				return
			}

			wrapper.ref = conn
			g.ctx.PublishEvent(server.NetworkConnectionOpenedEvent{
				ActorRef: conn,
			})

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
			if err = wrapper.ref.Tell(server.NetworkConnectionReceivedMessage{
				Packet: packet,
			}, vivids.WithMessageSender(g.ctx)); err != nil {
				action = gnet.Close
				break
			}
		}
	} else {
		buf, err := c.Next(-1)
		if err != nil {
			return gnet.Close
		}

		var clone = make([]byte, len(buf))
		copy(clone, buf)

		if err = c.Context().(vivids.ActorRef).Tell(server.NetworkConnectionReceivedMessage{
			Packet: server.NewPacket(clone),
		}, vivids.WithMessageSender(g.ctx)); err != nil {
			action = gnet.Close
		}
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
