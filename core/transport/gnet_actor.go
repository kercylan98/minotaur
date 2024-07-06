package transport

import (
	"context"
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/core/vivid/supervisorstategy"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"github.com/panjf2000/gnet/v2"
	gnetErrors "github.com/panjf2000/gnet/v2/pkg/errors"
	"net"
	"sync/atomic"
	"time"
)

const (
	schemaWebSocket = "ws"
	schemaTcp       = "tcp"
	schemaTcp4      = "tcp4"
	schemaTcp6      = "tcp6"
	schemaUdp       = "udp"
	schemaUdp4      = "udp4"
	schemaUdp6      = "udp6"
	schemaUnix      = "unix"
)

var _ gnet.EventHandler = &gnetActor{}

func newGNETActor(gnet *GNET, kit *GNETKit, schema, addr string, pattern ...string) *gnetActor {
	g := &gnetActor{
		gnet:    gnet,
		kit:     kit,
		addr:    addr,
		schema:  schema,
		pattern: "/",
	}
	if len(pattern) > 0 {
		g.pattern = pattern[0]
	}
	return g
}

type (
	gnetConnectionOpenedMessage gnetConnActor
	gnetConnectionClosedMessage gnetConnActor
)

type gnetActor struct {
	ctx      vivid.ActorContext
	gnet     *GNET
	kit      *GNETKit
	addr     string
	schema   string
	pattern  string
	engine   gnet.Engine
	upgrader ws.Upgrader
	showAddr string
}

func (g *gnetActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		g.onLaunch(ctx)
	case vivid.FutureForwardMessage:
		g.onFutureForward(ctx, m)
	case vivid.OnTerminate:
		g.onTerminate()
	case *gnetConnectionOpenedMessage:
		g.onConnectionOpened(ctx, (*gnetConnActor)(m))
	case *gnetConnectionClosedMessage:
		g.onConnectionClosed(ctx, (*gnetConnActor)(m))
	}
}

func (g *gnetActor) initWebSocketUpgrader() {
	g.upgrader = ws.Upgrader{
		OnRequest: func(uri []byte) (err error) {
			if string(uri) != g.pattern {
				err = errors.New("bad request")
			}
			return
		},
	}
}

func (g *gnetActor) OnBoot(eng gnet.Engine) (action gnet.Action) {
	g.engine = eng
	return
}

func (g *gnetActor) OnShutdown(eng gnet.Engine) {
	_ = eng.Stop(context.Background())
}

func (g *gnetActor) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	if g.schema == schemaWebSocket {
		c.SetContext(newGNETConnWebsocketContext(c))
	} else {
		ctx := &gnetConnContext{
			status: new(atomic.Uint32),
		}

		ctx.actor = newGNETConnActor(g.ctx.Ref(), ctx.status, g.kit, c, func(packet Packet, callback func(err error)) {
			err := c.AsyncWrite(packet.GetBytes(), func(c gnet.Conn, err error) error {
				if err != nil {
					callback(err)
				}
				return nil
			})
			if err != nil {
				callback(err)
			}
		})

		result, err := g.ctx.FutureAsk(g.ctx.Ref(), (*gnetConnectionOpenedMessage)(ctx.actor)).Result()
		if err != nil {
			return
		}

		ctx.ref = result.(vivid.ActorRef)
		c.SetContext(ctx)
	}
	return
}

func (g *gnetActor) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	switch ctx := c.Context().(type) {
	case *gnetConnContext:
		if err != nil {
			g.ctx.Tell(ctx.ref, err)
		}
		g.ctx.Tell(g.ctx.Ref(), (*gnetConnectionClosedMessage)(ctx.actor))
	case *gnetWebSocketConnContext:
		if err != nil {
			g.ctx.Tell(ctx.ref, err)
		}
		g.ctx.Tell(g.ctx.Ref(), (*gnetConnectionClosedMessage)(ctx.actor))
	}
	return
}

func (g *gnetActor) OnTraffic(c gnet.Conn) (action gnet.Action) {
	// 非 WebSocket 消息处理
	if g.schema != schemaWebSocket {
		ctx := c.Context().(*gnetConnContext)
		buf, err := c.Next(-1)
		if err != nil {
			return gnet.Close
		}

		var clone = make([]byte, len(buf))
		copy(clone, buf)
		g.ctx.Tell(ctx.ref, gnetReceivePacketMessage{packet: NewPacket(clone)})
		return
	}

	// 获取 WebSocket 上下文
	ctx := c.Context().(*gnetWebSocketConnContext)
	if err := ctx.readToBuffer(); err != nil {
		return gnet.Close
	}

	// 如果没有升级，那么执行升级
	if err := ctx.upgrade(g.upgrader, func() {
		ctx.actor = newGNETConnActor(g.ctx.Ref(), ctx.status, g.kit, c, func(packet Packet, callback func(err error)) {
			if err := wsutil.WriteServerMessage(c, packet.GetContext().(ws.OpCode), packet.GetBytes()); err != nil {
				callback(err)
			}
		})

		result, err := g.ctx.FutureAsk(g.ctx.Ref(), (*gnetConnectionOpenedMessage)(ctx.actor)).Result()
		if err != nil {
			return
		}

		ctx.ref = result.(vivid.ActorRef)
	}); err != nil {
		return gnet.Close
	}
	ctx.active = time.Now()

	// WebSocket 消息解码
	messages, err := ctx.decode()
	if err != nil {
		return gnet.Close
	}

	for _, message := range messages {
		p := NewPacket(message.Payload)
		p.SetContext(message.OpCode)
		g.ctx.Tell(ctx.ref, gnetReceivePacketMessage{packet: p})
	}
	return
}

func (g *gnetActor) OnTick() (delay time.Duration, action gnet.Action) {
	return
}

func (g *gnetActor) onLaunch(ctx vivid.ActorContext) {
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
		ctx.Tell(ctx.Ref(), fmt.Errorf("unsupported schema: %s", g.schema))
		return
	}

	host, port, err := net.SplitHostPort(g.addr)
	if err != nil {
		ctx.Tell(ctx.Ref(), err)
		return
	}
	if host == "" {
		ip, err := network.IP()
		if err != nil {
			ctx.Tell(ctx.Ref(), err)
			return
		}
		host = ip.String()
	}

	g.showAddr = fmt.Sprintf("%s://%s:%s", g.schema, host, port)
	if g.schema == schemaWebSocket {
		g.showAddr = fmt.Sprintf("%s://%s:%s%s", g.schema, host, port, g.pattern)
	}

	externalNetworkNum.Add(1)
	externalNetworkOnceLaunchInfo.Do(func() {
		g.gnet.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	})
	g.gnet.support.Logger().Info("", log.String("Minotaur", "enable network"), log.String("schema", g.schema), log.String("listen", g.showAddr))
	if externalNetworkLaunchedNum.Add(1) == externalNetworkNum.Load() {
		g.gnet.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	}

	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		return gnet.Run(g, addr, gnet.WithLogger(log.NewGNetLogger(log.GetDefault())))
	})
}

func (g *gnetActor) onFutureForward(ctx vivid.ActorContext, m vivid.FutureForwardMessage) {
	if m.Error != nil {
		panic(m.Error) // 交由监督者重启
	}
}

func (g *gnetActor) onTerminate() {
	if err := g.engine.Stop(context.Background()); err != nil && !errors.Is(err, gnetErrors.ErrEmptyEngine) {
		g.gnet.support.Logger().Error("network", log.String("status", "shutdown"), log.String("listen", g.showAddr), log.Err(err))
	} else {
		g.gnet.support.Logger().Info("network", log.String("status", "shutdown"), log.String("listen", g.showAddr))
	}
}

func (g *gnetActor) onConnectionOpened(ctx vivid.ActorContext, m *gnetConnActor) {
	ref := ctx.ActorOf(func() vivid.Actor {
		return m
	}, func(options *vivid.ActorOptions) {
		options.WithName("conn-" + m.gnetConn.RemoteAddr().String())
		options.WithSupervisorStrategy(supervisorstategy.OneForOne(func(reason, message vivid.Message) vivid.Directive {
			g.gnet.support.Logger().Error("network", log.String("status", "connection panic"), log.String("listen", g.showAddr), log.Err(reason.(error)))
			return vivid.DirectiveStop
		}, 0))
	})
	conn := NewConn(m.gnetConn, ctx.System(), ref)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				ctx.Tell(ref, err)
			default:
				ctx.Tell(ref, fmt.Errorf("connection opened panic: %v", err))
			}
		}
	}()
	if err := g.kit.connectionOpenedHook(g.kit, conn); err != nil {
		ctx.Tell(ref, err)
	}
	ctx.Reply(ref)
}

func (g *gnetActor) onConnectionClosed(ctx vivid.ActorContext, m *gnetConnActor) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				ctx.Tell(m.ref, err)
			default:
				ctx.Tell(m.ref, fmt.Errorf("connection opened panic: %v", err))
			}
		}
	}()

	if m.status.CompareAndSwap(gnetConnStatusOnline, gnetConnStatusClosed) {
		g.kit.connectionClosedHook(g.kit, NewConn(m.gnetConn, ctx.System(), m.ref), m.err)
		ctx.TerminateGracefully(m.ref)
	}
}
