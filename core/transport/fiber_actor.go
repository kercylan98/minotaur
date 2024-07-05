package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"net"
)

func newFiber(network *Fiber, addr string) *fiberActor {
	fa := &fiberActor{
		network: network,
		addr:    addr,
	}
	return fa
}

type fiberActor struct {
	network  *Fiber
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
		f.network.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	})
	f.network.support.Logger().Info("", log.String("Minotaur", "enable network"), log.String("schema", "http(s)"), log.String("listen", f.showAddr))
	if externalNetworkLaunchedNum.Add(1) == externalNetworkNum.Load() {
		f.network.support.Logger().Info("", log.String("Minotaur", "======================================================================="))
	}

	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		return f.network.app.Listen(f.addr)
	})
}

func (f *fiberActor) onFutureForward(ctx vivid.ActorContext, m vivid.FutureForwardMessage) {
	if m.Error != nil {
		panic(m.Error) // 交由监督者重启
	}
}

func (f *fiberActor) onTerminate() {
	if err := f.network.app.Shutdown(); err != nil {
		f.network.support.Logger().Error("network", log.String("status", "shutdown"), log.String("listen", f.showAddr), log.Err(err))
	} else {
		f.network.support.Logger().Info("network", log.String("status", "shutdown"), log.String("listen", f.showAddr))
	}
}
