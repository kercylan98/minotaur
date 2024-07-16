package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/core/vivid/supervisor"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"net"
)

func newFiberActor(fiber *Fiber, kit *FiberKit, addr string) *fiberActor {
	fa := &fiberActor{
		fiber: fiber,
		kit:   kit,
		addr:  addr,
	}
	return fa
}

type (
	fiberConnectionOpenedMessage fiberConnActor
	fiberConnectionClosedMessage fiberConnActor
)

type fiberActor struct {
	fiber    *Fiber
	kit      *FiberKit
	addr     string
	showAddr string
}

func (f *fiberActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		f.onLaunch(ctx)
	case vivid.FutureForwardMessage:
		f.onFutureForward(ctx, m)
	case vivid.OnTerminate:
		f.onTerminate()
	case *fiberConnectionOpenedMessage:
		f.onConnectionOpened(ctx, (*fiberConnActor)(m))
	case *fiberConnectionClosedMessage:
		f.onConnectionClosed(ctx, (*fiberConnActor)(m))
	}
}

func (f *fiberActor) onLaunch(ctx vivid.ActorContext) {
	host, port, err := net.SplitHostPort(f.addr)
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

	f.showAddr = fmt.Sprintf("http(s)://%s:%s", host, port)

	externalNetworkNum.Add(1)
	externalNetworkOnceLaunchInfo.Do(func() {
		f.fiber.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	})
	f.fiber.support.Logger().Info("", log.String("Minotaur", "enable network"), log.String("schema", "http(s)"), log.String("listen", f.showAddr))
	if externalNetworkLaunchedNum.Add(1) == externalNetworkNum.Load() {
		f.fiber.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	}

	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		return f.fiber.app.Listen(f.addr)
	})
}

func (f *fiberActor) onFutureForward(ctx vivid.ActorContext, m vivid.FutureForwardMessage) {
	if m.Error != nil {
		panic(m.Error) // 交由监督者重启
	}
}

func (f *fiberActor) onTerminate() {
	if err := f.fiber.app.Shutdown(); err != nil {
		f.fiber.support.Logger().Error("network", log.String("status", "shutdown"), log.String("listen", f.showAddr), log.Err(err))
	} else {
		f.fiber.support.Logger().Info("network", log.String("status", "shutdown"), log.String("listen", f.showAddr))
	}
}

func (f *fiberActor) onConnectionOpened(ctx vivid.ActorContext, m *fiberConnActor) {
	ref := ctx.ActorOf(func() vivid.Actor {
		return m
	}, func(options *vivid.ActorOptions) {
		options.WithName("conn-" + m.fiberConn.RemoteAddr().String())
		options.WithSupervisorStrategy(supervisor.Stop(), func(reason, message vivid.Message) {
			f.fiber.support.Logger().Error("network", log.String("status", "connection panic"), log.String("listen", f.showAddr), log.Err(reason.(error)))
		})
	})
	conn := NewConn(m.fiberConn, ctx.System(), ref)
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
	if err := f.kit.fws.connectionOpenedHook(f.kit, m.ctx, conn); err != nil {
		ctx.Tell(ref, err)
	}
	ctx.Reply(ref)
}

func (f *fiberActor) onConnectionClosed(ctx vivid.ActorContext, m *fiberConnActor) {
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

	f.kit.fws.connectionClosedHook(f.kit, m.ctx, NewConn(m.fiberConn, ctx.System(), m.ref), m.err)
	if m.status.Load() == fiberConnStatusOnline {
		ctx.TerminateGracefully(m.ref)
	}
}
