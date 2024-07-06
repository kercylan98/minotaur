package transport

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"sync/atomic"
)

const (
	fiberConnStatusOnline uint32 = iota
	fiberConnStatusClosed
)

func newFiberConnActor(fiberActorRef vivid.ActorRef, status *atomic.Uint32, kit *FiberKit, ctx *FiberContext, conn *websocket.Conn) *fiberConnActor {
	a := &fiberConnActor{
		fiberActorRef: fiberActorRef,
		kit:           kit,
		ctx:           ctx,
		fiberConn:     &fiberConnWrapper{conn},
		status:        status,
	}
	return a
}

type fiberReceivePacketMessage struct {
	packet Packet
}

type fiberConnActor struct {
	fiberActorRef vivid.ActorRef
	kit           *FiberKit
	ctx           *FiberContext
	fiberConn     *fiberConnWrapper
	conn          *Conn
	ref           vivid.ActorRef
	err           error
	status        *atomic.Uint32
}

func (f *fiberConnActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		f.ref = ctx.Ref()
		f.conn = NewConn(f.fiberConn, ctx.System(), ctx.Ref())
	case fiberReceivePacketMessage:
		if f.err = f.kit.fws.connectionPacketHook(f.kit, f.ctx, f.conn, m.packet); f.err != nil {
			ctx.Tell(f.fiberActorRef, (*fiberConnectionClosedMessage)(f))
			return
		}
	case Packet:
		if f.err = f.fiberConn.WriteMessage(m.GetContext().(int), m.GetBytes()); f.err != nil {
			ctx.Tell(f.fiberActorRef, (*fiberConnectionClosedMessage)(f))
		}
	case error:
		f.err = m
		ctx.Tell(f.fiberActorRef, (*fiberConnectionClosedMessage)(f))
	case vivid.OnTerminate:
		if f.status.CompareAndSwap(fiberConnStatusOnline, fiberConnStatusClosed) {
			ctx.Tell(f.fiberActorRef, (*fiberConnectionClosedMessage)(f))
		}
		if f.err == nil {
			_ = f.fiberConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, charproc.None))
		} else {
			_ = f.fiberConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseAbnormalClosure, f.err.Error()))
		}
		_ = f.fiberConn.Close()

	}
}
