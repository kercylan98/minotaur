package socket

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"time"
)

// Writer 是用于将数据写入到网络连接的写入器，它是一个函数类型。在该函数中，需要决定要如何将数据写入到网络连接中。
type Writer func(packet []byte, ctx any) error

func newWriterActor(writer Writer) *writerActor {
	return &writerActor{writer: writer}
}

type writerActor struct {
	writer        Writer
	writeDeadline time.Duration
}

func (w *writerActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *Packet:
		ctx.StopTask(writeDeadlineTaskName)
		if err := w.writer(m.GetData(), m.GetContext()); err != nil {
			ctx.Tell(ctx.Parent(), err)
		}
		w.refreshWriteDeadline(ctx)
	case writeDeadline:
		w.writeDeadline = time.Duration(m)
		w.refreshWriteDeadline(ctx)
	}
}

func (w *writerActor) refreshWriteDeadline(ctx vivid.ActorContext) {
	if w.writeDeadline > 0 {
		ctx.AfterTask(writeDeadlineTaskName, w.writeDeadline, func(vivid.ActorContext) {
			ctx.Tell(ctx.Parent(), writeDeadlineError)
		})
	}
}
