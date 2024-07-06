package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/panjf2000/gnet/v2"
	"sync/atomic"
)

const (
	gnetConnStatusOnline uint32 = iota
	gnetConnStatusClosed
)

func newGNETConnActor(gnetActorRef vivid.ActorRef, status *atomic.Uint32, kit *GNETKit, conn gnet.Conn, writer func(packet Packet, callback func(err error))) *gnetConnActor {
	a := &gnetConnActor{
		gnetActorRef: gnetActorRef,
		kit:          kit,
		gnetConn:     conn,
		status:       status,
		writer:       writer,
	}
	return a
}

type gnetReceivePacketMessage struct {
	packet Packet
}

type gnetConnActor struct {
	gnetActorRef vivid.ActorRef
	kit          *GNETKit
	gnetConn     gnet.Conn
	conn         *Conn
	ref          vivid.ActorRef
	err          error
	status       *atomic.Uint32
	writer       func(packet Packet, callback func(err error))

	connectionPacketHook FiberConnectionPacketHook
}

func (f *gnetConnActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		f.ref = ctx.Ref()
		f.conn = NewConn(f.gnetConn, ctx.System(), ctx.Ref(), ctx.Ref())
	case gnetReceivePacketMessage:
		if f.err = f.kit.connectionPacketHook(f.kit, f.conn, m.packet); f.err != nil {
			ctx.Tell(f.gnetActorRef, (*gnetConnectionClosedMessage)(f))
			return
		}
	case Packet:
		f.writer(m, func(err error) {
			ctx.Tell(ctx.Ref(), err)
		})
	case error:
		f.err = m
		ctx.Tell(f.gnetActorRef, (*gnetConnectionClosedMessage)(f))
	case vivid.OnTerminate:
		if f.status.CompareAndSwap(gnetConnStatusOnline, gnetConnStatusClosed) {
			ctx.Tell(f.gnetActorRef, (*gnetConnectionClosedMessage)(f))
		}
		_ = f.gnetConn.Close()
	}
}
