package socket

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"time"
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

// Socket 是维护支持 vivid.Actor 的网络长连接包装接口，该接口无需进行实现，它将由内部的 socket 实现并进行维护
type Socket interface {
	// React 将数据包及其上下文通过 Actor.OnPacket 函数进行响应
	React(packet []byte, ctx any)

	// Write 将数据包及其上下文写入到客户端的网络连接中
	Write(packet []byte, ctx any)

	// WriteBytes 将字节数据写入到客户端的网络连接中
	WriteBytes(packet []byte)

	// WriteString 将字符串数据写入到客户端的网络连接中
	WriteString(packet string)

	// WritePacket 将数据包及其上下文写入到客户端的网络连接中
	WritePacket(packet *Packet)

	// DebounceWrite 将数据包及其上下文以指定的防抖延迟写入到客户端的网络连接中，当防抖时间窗口期间内收到多个相同名称的数据包时，仅写入最后一次的数据包
	DebounceWrite(name string, delay time.Duration, packet []byte, ctx any)

	// DebounceWriteBytes 将字节数据以指定的防抖延迟写入到客户端的网络连接中，当防抖时间窗口期间内收到多个相同名称的数据包时，仅写入最后一次的数据包
	DebounceWriteBytes(name string, delay time.Duration, packet []byte)

	// DebounceWriteString 将字符串数据以指定的防抖延迟写入到客户端的网络连接中，当防抖时间窗口期间内收到多个相同名称的数据包时，仅写入最后一次的数据包
	DebounceWriteString(name string, delay time.Duration, packet string)

	// DebounceWritePacket 将数据包及其上下文以指定的防抖延迟写入到客户端的网络连接中，当防抖时间窗口期间内收到多个相同名称的数据包时，仅写入最后一次的数据包
	DebounceWritePacket(name string, delay time.Duration, packet *Packet)

	// Close 可携带错误信息地关闭 Socket 连接，当包含错误且 Socket 绑定的 Actor 实现了 CloseActor 接口时，可在 CloseActor.OnClose 中接收到该错误
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

func (s *socket) write(packet []byte, ctx any) {
	s.ctx.Tell(s.writerRef, NewPacket(packet, ctx))
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
	s.ctx.Tell(s.ctx.Ref(), NewPacket(packet, ctx))
}

func (s *socket) WriteBytes(packet []byte) {
	s.write(packet, nil)
}

func (s *socket) WriteString(packet string) {
	s.write([]byte(packet), nil)
}

func (s *socket) WritePacket(packet *Packet) {
	s.write(packet.GetData(), packet.GetContext())
}

func (s *socket) Write(packet []byte, ctx any) {
	s.write(packet, ctx)
}

func (s *socket) debounceWrite(name string, delay time.Duration, packet []byte, ctx any) {
	s.ctx.AfterTask(name, delay, func(vivid.ActorContext) {
		s.write(packet, ctx)
	})
}

func (s *socket) DebounceWrite(name string, delay time.Duration, packet []byte, ctx any) {
	s.debounceWrite(name, delay, packet, ctx)
}

func (s *socket) DebounceWriteBytes(name string, delay time.Duration, packet []byte) {
	s.debounceWrite(name, delay, packet, nil)
}

func (s *socket) DebounceWriteString(name string, delay time.Duration, packet string) {
	s.debounceWrite(name, delay, []byte(packet), nil)
}

func (s *socket) DebounceWritePacket(name string, delay time.Duration, packet *Packet) {
	s.debounceWrite(name, delay, packet.GetData(), packet.GetContext())
}

func (s *socket) Close(err ...error) {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	s.ctx.Tell(s.ctx.Ref(), e)
}
