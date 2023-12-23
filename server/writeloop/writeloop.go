package writeloop

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/log"
)

// NewWriteLoop 创建写循环
//   - pool 用于管理 Message 对象的缓冲池，在创建 Message 对象时也应该使用该缓冲池，以便复用 Message 对象。 WriteLoop 会在写入完成后将 Message 对象放回缓冲池
//   - writeHandler 写入处理函数
//   - errorHandler 错误处理函数
//
// 传入 writeHandler 的消息对象是从 pool 中获取的，并且在 writeHandler 执行完成后会被放回 pool 中，因此 writeHandler 不应该持有消息对象的引用，同时也不应该主动释放消息对象
func NewWriteLoop[Message any](pool *concurrent.Pool[Message], writeHandler func(message Message) error, errorHandler func(err any)) *WriteLoop[Message] {
	wl := &WriteLoop[Message]{
		buf: buffer.NewUnbounded[Message](),
	}
	go func() {
		for {
			select {
			case message, ok := <-wl.buf.Get():
				if !ok {
					return
				}
				wl.buf.Load()

				err := writeHandler(message)
				pool.Release(message)
				if err != nil {
					if errorHandler == nil {
						log.Error("WriteLoop", log.Err(err))
						continue
					}
					errorHandler(err)
				}
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

// Put 将数据放入写循环，message 应该来源于 concurrent.Pool
func (slf *WriteLoop[Message]) Put(message Message) {
	slf.buf.Put(message)
}

// Close 关闭写循环
func (slf *WriteLoop[Message]) Close() {
	slf.buf.Close()
}
