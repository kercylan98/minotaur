package writeloop

import (
	"github.com/kercylan98/minotaur/utils/hub"
	"github.com/kercylan98/minotaur/utils/log"
)

// NewChannel 创建基于 Channel 的写循环
//   - pool 用于管理 Message 对象的缓冲池，在创建 Message 对象时也应该使用该缓冲池，以便复用 Message 对象。 Channel 会在写入完成后将 Message 对象放回缓冲池
//   - channelSize Channel 的大小
//   - writeHandler 写入处理函数
//   - errorHandler 错误处理函数
//
// 传入 writeHandler 的消息对象是从 Channel 中获取的，因此 writeHandler 不应该持有消息对象的引用，同时也不应该主动释放消息对象
func NewChannel[Message any](pool *hub.ObjectPool[Message], channelSize int, writeHandler func(message Message) error, errorHandler func(err any)) *Channel[Message] {
	wl := &Channel[Message]{
		c: make(chan Message, channelSize),
	}
	go func() {
		for {
			select {
			case message, ok := <-wl.c:
				if !ok {
					return
				}

				err := writeHandler(message)
				pool.Release(message)
				if err != nil {
					if errorHandler == nil {
						log.Error("Channel", log.Err(err))
						continue
					}
					errorHandler(err)
				}
			}
		}
	}()

	return wl
}

// Channel 基于 chan 的写循环，与 Unbounded 相同，但是使用 Channel 实现
type Channel[T any] struct {
	c chan T
}

// Put 将数据放入写循环，message 应该来源于 hub.ObjectPool
func (slf *Channel[T]) Put(message T) {
	slf.c <- message
}

// Close 关闭写循环
func (slf *Channel[T]) Close() {
	close(slf.c)
}
