package dispatcher

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/utils/buffer"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/super"
	"sync"
	"sync/atomic"
)

var unique = struct{}{}

// Handler 消息处理器
type Handler[P Producer, M Message[P]] func(dispatcher *Dispatcher[P, M], message M)

// NewDispatcher 创建一个新的消息分发器 Dispatcher 实例
func NewDispatcher[P Producer, M Message[P]](bufferSize int, name string, handler Handler[P, M]) *Dispatcher[P, M] {
	if bufferSize <= 0 || handler == nil {
		panic(fmt.Errorf("bufferSize must be greater than 0 and handler must not be nil, but got bufferSize: %d, handler is nil: %v", bufferSize, handler == nil))
	}
	d := &Dispatcher[P, M]{
		name:    name,
		buf:     buffer.NewRingUnbounded[M](bufferSize),
		handler: handler,
		uniques: haxmap.New[string, struct{}](),
		pmc:     make(map[P]int64),
		pmcF:    make(map[P]func(p P, dispatcher *Action[P, M])),
		abort:   make(chan struct{}),
	}
	return d
}

// Dispatcher 用于服务器消息处理的消息分发器
//
// 这个消息分发器为并发安全的生产者和消费者模型，生产者可以是任意类型，消费者必须是 Message 接口的实现。
// 生产者可以通过 Put 方法并发安全地将消息放入消息分发器，消息执行过程不会阻塞到 Put 方法，同时允许在 Start 方法之前调用 Put 方法。
// 在执行 Start 方法后，消息分发器会阻塞地从消息缓冲区中读取消息，然后执行消息处理器，消息处理器的执行过程不会阻塞到消息的生产。
//
// 为了保证消息不丢失，内部采用了 buffer.RingUnbounded 作为缓冲区实现，并且消息分发器不提供 Close 方法。
// 如果需要关闭消息分发器，可以通过 Expel 方法设置驱逐计划，当消息分发器中没有任何消息时，将会被释放。
// 同时，也可以使用 UnExpel 方法取消驱逐计划。
//
// 为什么提供 Expel 和 UnExpel 方法：
//   - 在连接断开时，当需要执行一系列消息处理时，如果直接关闭消息分发器，可能会导致消息丢失。所以提供了 Expel 方法，可以在消息处理完成后再关闭消息分发器。
//   - 当消息还未处理完成时连接重连，如果没有取消驱逐计划，可能会导致消息分发器被关闭。所以提供了 UnExpel 方法，可以在连接重连后取消驱逐计划。
type Dispatcher[P Producer, M Message[P]] struct {
	buf           *buffer.RingUnbounded[M]
	uniques       *haxmap.Map[string, struct{}]
	handler       Handler[P, M]
	expel         bool
	mc            int64
	pmc           map[P]int64
	pmcF          map[P]func(p P, dispatcher *Action[P, M])
	lock          sync.RWMutex
	name          string
	closedHandler atomic.Pointer[func(dispatcher *Action[P, M])]
	abort         chan struct{}
}

// SetProducerDoneHandler 设置特定生产者所有消息处理完成时的回调函数
//   - 如果 handler 为 nil，则会删除该生产者的回调函数
//
// 需要注意的是，该 handler 中
func (d *Dispatcher[P, M]) SetProducerDoneHandler(p P, handler func(p P, dispatcher *Action[P, M])) *Dispatcher[P, M] {
	d.lock.Lock()
	if handler == nil {
		delete(d.pmcF, p)
	} else {
		d.pmcF[p] = handler
		if pmc := d.pmc[p]; pmc <= 0 {
			func(producer P, handler func(p P, dispatcher *Action[P, M])) {
				defer func(producer P) {
					if err := super.RecoverTransform(recover()); err != nil {
						log.Error("Dispatcher.ProducerDoneHandler", log.Any("producer", producer), log.Err(err))
					}
				}(p)
				handler(p, &Action[P, M]{d: d, unlock: true})
			}(p, handler)
		}
	}
	d.lock.Unlock()
	return d
}

// SetClosedHandler 设置消息分发器关闭时的回调函数
func (d *Dispatcher[P, M]) SetClosedHandler(handler func(dispatcher *Action[P, M])) *Dispatcher[P, M] {
	d.closedHandler.Store(&handler)
	return d
}

// Name 获取消息分发器名称
func (d *Dispatcher[P, M]) Name() string {
	return d.name
}

// Unique 设置唯一消息键，返回是否已存在
func (d *Dispatcher[P, M]) Unique(name string) bool {
	_, loaded := d.uniques.GetOrSet(name, unique)
	return loaded
}

// AntiUnique 取消唯一消息键
func (d *Dispatcher[P, M]) AntiUnique(name string) {
	d.uniques.Del(name)
}

// Expel 设置该消息分发器即将被驱逐，当消息分发器中没有任何消息时，会自动关闭
func (d *Dispatcher[P, M]) Expel() {
	d.lock.Lock()
	d.noLockExpel()
	d.lock.Unlock()
}

func (d *Dispatcher[P, M]) noLockExpel() {
	d.expel = true
	if d.mc <= 0 {
		d.abort <- struct{}{}
	}
}

// UnExpel 取消特定生产者的驱逐计划
func (d *Dispatcher[P, M]) UnExpel() {
	d.lock.Lock()
	d.noLockUnExpel()
	d.lock.Unlock()
}

func (d *Dispatcher[P, M]) noLockUnExpel() {
	d.expel = false
}

// IncrCount 主动增量设置特定生产者的消息计数，这在等待异步消息完成后再关闭消息分发器时非常有用
//   - 如果 i 为负数，则会减少消息计数
func (d *Dispatcher[P, M]) IncrCount(producer P, i int64) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.mc += i
	d.pmc[producer] += i
	if d.expel && d.mc <= 0 {
		d.abort <- struct{}{}
	}
}

// Put 将消息放入分发器
func (d *Dispatcher[P, M]) Put(message M) {
	d.lock.Lock()
	d.mc++
	d.pmc[message.GetProducer()]++
	d.lock.Unlock()
	d.buf.Write(message)
}

// Start 以非阻塞的方式开始进行消息分发，当消息分发器中没有任何消息并且处于驱逐计划 Expel 时，将会自动关闭
func (d *Dispatcher[P, M]) Start() *Dispatcher[P, M] {
	go func(d *Dispatcher[P, M]) {
	process:
		for {
			select {
			case <-d.abort:
				d.buf.Close()
				break process
			case message := <-d.buf.Read():
				// 先取出生产者信息，避免处理函数中将消息释放
				p := message.GetProducer()
				d.handler(d, message)
				d.lock.Lock()
				d.mc--
				pmc := d.pmc[p] - 1
				d.pmc[p] = pmc
				if f := d.pmcF[p]; f != nil && pmc <= 0 {
					func(producer P) {
						defer func(producer P) {
							if err := super.RecoverTransform(recover()); err != nil {
								log.Error("Dispatcher.ProducerDoneHandler", log.Any("producer", producer), log.Err(err))
							}
						}(p)
						f(p, &Action[P, M]{d: d, unlock: true})
					}(p)
				}
				if d.mc <= 0 && d.expel {
					d.buf.Close()
					d.lock.Unlock()
					break process
				}
				d.lock.Unlock()
			}
		}
		if ch := d.closedHandler.Load(); ch != nil {
			(*ch)(&Action[P, M]{d: d, unlock: true})
		}
		close(d.abort)
	}(d)
	return d
}

// Closed 判断消息分发器是否已关闭
func (d *Dispatcher[P, M]) Closed() bool {
	return d.buf.Closed()
}
