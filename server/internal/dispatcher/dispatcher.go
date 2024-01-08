package dispatcher

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/utils/buffer"
	"sync"
	"sync/atomic"
)

var unique = struct{}{}

// Handler 消息处理器
type Handler[P Producer, M Message[P]] func(dispatcher *Dispatcher[P, M], message M)

// NewDispatcher 生成消息分发器
func NewDispatcher[P Producer, M Message[P]](bufferSize int, name string, handler Handler[P, M]) *Dispatcher[P, M] {
	d := &Dispatcher[P, M]{
		name:    name,
		buf:     buffer.NewRingUnbounded[M](bufferSize),
		handler: handler,
		uniques: haxmap.New[string, struct{}](),
		pmc:     make(map[P]int64),
		pmcF:    make(map[P]func(p P, dispatcher *Dispatcher[P, M])),
		abort:   make(chan struct{}),
	}
	return d
}

// Dispatcher 消息分发器
type Dispatcher[P Producer, M Message[P]] struct {
	buf           *buffer.RingUnbounded[M]
	uniques       *haxmap.Map[string, struct{}]
	handler       Handler[P, M]
	expel         bool
	mc            int64
	pmc           map[P]int64
	pmcF          map[P]func(p P, dispatcher *Dispatcher[P, M])
	lock          sync.RWMutex
	name          string
	closedHandler atomic.Pointer[func(dispatcher *Dispatcher[P, M])]
	abort         chan struct{}
}

// SetProducerDoneHandler 设置特定生产者所有消息处理完成时的回调函数
func (d *Dispatcher[P, M]) SetProducerDoneHandler(p P, handler func(p P, dispatcher *Dispatcher[P, M])) *Dispatcher[P, M] {
	d.lock.Lock()
	if handler == nil {
		delete(d.pmcF, p)
	} else {
		d.pmcF[p] = handler
	}
	d.lock.Unlock()
	return d
}

// SetClosedHandler 设置消息分发器关闭时的回调函数
func (d *Dispatcher[P, M]) SetClosedHandler(handler func(dispatcher *Dispatcher[P, M])) *Dispatcher[P, M] {
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
	d.expel = true
	if d.mc <= 0 {
		d.abort <- struct{}{}
	}
	d.lock.Unlock()
}

// UnExpel 取消特定生产者的驱逐计划
func (d *Dispatcher[P, M]) UnExpel() {
	d.lock.Lock()
	d.expel = false
	d.lock.Unlock()
}

// IncrCount 主动增量设置特定生产者的消息计数，这在等待异步消息完成后再关闭消息分发器时非常有用
func (d *Dispatcher[P, M]) IncrCount(producer P, i int64) {
	d.lock.Lock()
	d.mc += i
	d.pmc[producer] += i
	if d.expel && d.mc <= 0 {
		d.abort <- struct{}{}
	}
	d.lock.Unlock()
}

// Put 将消息放入分发器
func (d *Dispatcher[P, M]) Put(message M) {
	d.lock.Lock()
	d.mc++
	d.pmc[message.GetProducer()]++
	d.lock.Unlock()
	d.buf.Write(message)
}

// Start 以阻塞的方式开始进行消息分发，当消息分发器中没有任何消息时，会自动关闭
func (d *Dispatcher[P, M]) Start() *Dispatcher[P, M] {
	go func(d *Dispatcher[P, M]) {
	process:
		for {
			select {
			case <-d.abort:
				d.buf.Close()
				break process
			case message := <-d.buf.Read():
				d.handler(d, message)
				d.lock.Lock()
				d.mc--
				p := message.GetProducer()
				pmc := d.pmc[p] - 1
				d.pmc[p] = pmc
				if f := d.pmcF[p]; f != nil && pmc <= 0 {
					go f(p, d)
				}
				if d.mc <= 0 && d.expel {
					d.buf.Close()
					break process
				}
				d.lock.Unlock()
			}
		}
		closedHandler := *(d.closedHandler.Load())
		if closedHandler != nil {
			closedHandler(d)
		}
		close(d.abort)
	}(d)
	return d
}

// Closed 判断消息分发器是否已关闭
func (d *Dispatcher[P, M]) Closed() bool {
	return d.buf.Closed()
}
