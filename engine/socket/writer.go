package socket

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

// Writer 是用于将数据写入到网络连接的写入器，它是一个函数类型。在该函数中，需要决定要如何将数据写入到网络连接中。
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
