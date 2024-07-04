package transport

import (
	"context"
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/core/vivid/supervisorstategy"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"github.com/panjf2000/gnet/v2"
	gnetErrors "github.com/panjf2000/gnet/v2/pkg/errors"
	"net"
	"sync"
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

var _ gnet.EventHandler = &gnetEngine{}
var gnetEngineNum atomic.Int32
var gnetEngineLaunchedNum atomic.Int32
var gnetOnceLaunchInfo sync.Once

func newGnetEngine(network *ExternalNetwork, schema, addr string, pattern ...string) *gnetEngine {
	g := &gnetEngine{
		network: network,
		addr:    addr,
		schema:  schema,
		pattern: collection.FindFirstOrDefaultInSlice(pattern, "/"),
	}
	return g
}

type gnetEngine struct {
	network  *ExternalNetwork
	addr     string
	schema   string
	pattern  string
	eng      gnet.Engine
	upgrader ws.Upgrader
	ref      vivid.ActorRef
	showAddr string
}

func (g *gnetEngine) OnShutdown(eng gnet.Engine) {
	_ = eng.Stop(context.Background())
}

func (g *gnetEngine) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		g.onLaunch(ctx)
	case vivid.FutureForwardMessage:
		g.onFutureForward(ctx, m)
	case vivid.OnTerminate:
		g.onTerminate()
	}
}

func (g *gnetEngine) Shutdown() error {
	return g.eng.Stop(context.TODO())
}

func (g *gnetEngine) OnBoot(eng gnet.Engine) (action gnet.Action) {
	g.eng = eng
	return
}

func (g *gnetEngine) createWriterActor(c gnet.Conn) vivid.Actor {
	return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
		switch m := ctx.Message().(type) {
		case Packet:
			c.AsyncWrite(m.GetBytes(), func(c gnet.Conn, err error) error {
				return nil
			})
		}
	})
}
func (g *gnetEngine) createWebsocketWriterActor(c gnet.Conn) vivid.Actor {
	return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
		switch m := ctx.Message().(type) {
		case Packet:
			wsutil.WriteServerMessage(c, m.GetContext().(ws.OpCode), m.GetBytes())
		}
	})
}

func (g *gnetEngine) createReaderActor(c gnet.Conn) vivid.Actor {
	var conn *Conn
	var err error
	return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
		switch m := ctx.Message().(type) {
		case vivid.OnLaunch:
			writerRef := ctx.ActorOf(func() vivid.Actor {
				if g.schema == schemaWebSocket {
					return g.createWebsocketWriterActor(c)
				}
				return g.createWriterActor(c)
			}, func(options *vivid.ActorOptions) {
				options.WithName("writer")
				options.WithSupervisorStrategy(supervisorstategy.OneForOne(func(reason, message vivid.Message) vivid.Directive {
					return vivid.DirectiveStop
				}, 0))
			})
			conn = NewConn(c, ctx, writerRef)
			g.network.connOpenedHandler(conn)
		case Packet:
			g.network.packetHandler(conn, m)
		case error:
			err = m
		case vivid.OnTerminate:
			g.network.connClosedHandler(conn, err)
		}
	})
}

func (g *gnetEngine) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	if g.schema == schemaWebSocket {
		c.SetContext(newWebsocketWrapper(c))
	} else {
		ref := g.network.support.System().ActorOf(func() vivid.Actor {
			return g.createReaderActor(c)
		}, func(options *vivid.ActorOptions) {
			options.WithName("conn-" + c.RemoteAddr().String())
			options.WithSupervisorStrategy(supervisorstategy.OneForOne(func(reason, message vivid.Message) vivid.Directive {
				return vivid.DirectiveStop
			}, 0))
		})

		c.SetContext(ref)
	}
	return
}

func (g *gnetEngine) onLaunch(ctx vivid.ActorContext) {
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

	gnetEngineNum.Add(1)
	gnetOnceLaunchInfo.Do(func() {
		g.network.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	})
	g.network.support.Logger().Info("", log.String("Minotaur", "Start network"), log.String("schema", g.schema), log.String("listen", g.showAddr))
	if gnetEngineLaunchedNum.Add(1) == gnetEngineNum.Load() {
		g.network.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	}

	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		return gnet.Run(g, addr, gnet.WithLogger(log.NewGNetLogger(log.GetDefault())))
	})
}

func (g *gnetEngine) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	var ref vivid.ActorRef
	switch ctx := c.Context().(type) {
	case *websocketWrapper:
		ref = ctx.ref
	case vivid.ActorRef:
		ref = ctx
	}
	if ref != nil {
		g.network.support.System().Context().Tell(ref, err)
		g.network.support.System().Terminate(ref)
	}
	return
}

func (g *gnetEngine) OnTraffic(c gnet.Conn) (action gnet.Action) {
	if g.schema == schemaWebSocket {
		wrapper := c.Context().(*websocketWrapper)

		if err := wrapper.readToBuffer(); err != nil {
			return gnet.Close
		}

		if err := wrapper.upgrade(g.upgrader, func() {
			ref := g.network.support.System().ActorOf(func() vivid.Actor {
				return g.createReaderActor(c)
			}, func(options *vivid.ActorOptions) {
				options.WithName("conn-" + c.RemoteAddr().String())
				options.WithSupervisorStrategy(supervisorstategy.OneForOne(func(reason, message vivid.Message) vivid.Directive {
					return vivid.DirectiveStop
				}, 0))
			})

			wrapper.process = g.network.support.GetProcess(ref.Address())
			wrapper.ref = core.NewProcessRef(wrapper.process.GetAddress())
			c.SetContext(wrapper)

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
			p := NewPacket(message.Payload)
			p.SetContext(message.OpCode)
			wrapper.process.SendUserMessage(g.network.support.System().Context().Ref(), p)
		}
	} else {
		buf, err := c.Next(-1)
		if err != nil {
			return gnet.Close
		}

		var clone = make([]byte, len(buf))
		copy(clone, buf)

		g.network.support.System().Context().Tell(c.Context().(vivid.ActorRef), clone)
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

func (g *gnetEngine) onFutureForward(ctx vivid.ActorContext, m vivid.FutureForwardMessage) {
	if m.Error != nil {
		panic(m.Error) // 交由监督者重启
	}
}

func (g *gnetEngine) onTerminate() {
	if err := g.eng.Stop(context.Background()); err != nil && !errors.Is(err, gnetErrors.ErrEmptyEngine) {
		g.network.support.Logger().Error("network", log.String("status", "shutdown"), log.String("listen", g.showAddr), log.Err(err))
	} else {
		g.network.support.Logger().Info("network", log.String("status", "shutdown"), log.String("listen", g.showAddr))
	}
}
