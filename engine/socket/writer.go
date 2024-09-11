package socket

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

type Writer func(packet []byte, ctx any) error

func newWriterActor(writer Writer) *writerActor {
	return &writerActor{writer: writer}
}

type writerActor struct {
	writer Writer
}

func (w *writerActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *Packet:
		if err := w.writer(m.GetData(), m.GetContext()); err != nil {
			ctx.Tell(ctx.Parent(), err)
		}
	}
}
