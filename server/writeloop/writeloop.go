package writeloop

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/log"
	"runtime/debug"
)

// NewWriteLoop 创建写循环
//   - pool 用于管理 Message 对象的缓冲池，在创建 Message 对象时也应该使用该缓冲池，以便复用 Message 对象。 WriteLoop 会在写入完成后将 Message 对象放回缓冲池
func NewWriteLoop[Message any](pool *concurrent.Pool[Message], writeHandle func(message Message) error, errorHandle func(err any)) *WriteLoop[Message] {
	wl := &WriteLoop[Message]{
		buf: buffer.NewUnboundedN[Message](),
	}
	go func() {
		for !wl.buf.IsClosed() {
			select {
			case message, ok := <-wl.buf.Get():
				if !ok {
					return
				}
				wl.buf.Load()
				func() {
					defer func() {
						pool.Release(message)
						if err := recover(); err != nil {
							if errorHandle == nil {
								log.Error("WriteLoop", log.Any("err", err))
								debug.PrintStack()
								return
							}
							errorHandle(err)
						}
					}()
					err := writeHandle(message)
					if err != nil {
						panic(err)
					}

				}()
			}
		}
	}()

	return wl
}

// WriteLoop 写循环
//   - 用于将数据并发安全的写入到底层连接
type WriteLoop[Message any] struct {
	buf *buffer.Unbounded[Message]
}

// Put 将数据放入写循环
func (slf *WriteLoop[Message]) Put(message Message) {
	slf.buf.Put(message)
}

// Close 关闭写循环
func (slf *WriteLoop[Message]) Close() {
	slf.buf.Close()
}
