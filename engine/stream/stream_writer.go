package stream

import "github.com/kercylan98/minotaur/engine/vivid"

func newStreamWriter(stream Stream) *streamWriter {
	return &streamWriter{stream: stream}
}

type streamWriter struct {
	stream Stream
}

func (w *streamWriter) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *Packet:
		w.onPacket(ctx, m)
	case *vivid.OnTerminate:
		_ = w.stream.Close()
	}
}

func (w *streamWriter) onPacket(ctx vivid.ActorContext, m *Packet) {
	if err := w.stream.Write(m); err != nil {
		ctx.Tell(ctx.Parent(), err)
	}
}
