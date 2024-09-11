package socket

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

var _ Socket = (*socket)(nil)

type Closer = func() error

func newSocket(actor Actor, writer Writer, closer Closer) *socket {
	return &socket{
		Actor:  actor,
		writer: writer,
		closer: closer,
	}
}

type Socket interface {
	React(packet []byte, ctx any)

	WriteBytes(packet []byte)

	WriteString(packet string)

	Write(packet []byte, ctx any)

	WritePacket(packet *Packet)

	Close(err ...error)
}

type socket struct {
	Actor
	writer    Writer
	closer    Closer
	ctx       vivid.ActorContext
	writerRef vivid.ActorRef
	err       error
}

func (s *socket) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		s.onLaunch(ctx, m)
		switch v := s.Actor.(type) {
		case OpenedActor:
			v.OnOpened(ctx, s)
		}
	case error:
		s.onError(ctx, m)
	case *Packet:
		s.Actor.OnPacket(ctx, s, m)
	case *vivid.OnTerminate:
		s.Actor.OnReceive(ctx)
		switch v := s.Actor.(type) {
		case CloseActor:
			v.OnClose(ctx, s, s.err)
		}
	case *vivid.OnTerminated:
		if m.TerminatedActor.Equal(s.ctx.Ref()) {
			_ = s.closer()
		}
		return
	}
	s.Actor.OnReceive(ctx)
}

func (s *socket) onLaunch(ctx vivid.ActorContext, m *vivid.OnLaunch) {
	s.ctx = ctx
	s.writerRef = ctx.ActorOfF(func() vivid.Actor {
		return newWriterActor(s.writer)
	})
}

func (s *socket) onError(ctx vivid.ActorContext, err error) {
	if err != nil {
		s.err = err
	}
	ctx.Terminate(ctx.Ref(), false)
}

func (s *socket) React(packet []byte, ctx any) {
	s.ctx.Tell(s.ctx.Ref(), newPacket(packet, ctx))
}

func (s *socket) WriteBytes(packet []byte) {
	s.Write(packet, nil)
}

func (s *socket) WriteString(packet string) {
	s.Write([]byte(packet), nil)
}

func (s *socket) WritePacket(packet *Packet) {
	s.Write(packet.GetData(), packet.GetContext())
}

func (s *socket) Write(packet []byte, ctx any) {
	s.ctx.Tell(s.writerRef, newPacket(packet, ctx))
}

func (s *socket) Close(err ...error) {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	s.ctx.Tell(s.ctx.Ref(), e)
}
