package actor

import (
	"context"
	"github.com/kercylan98/minotaur/utils/buffer"
	"github.com/kercylan98/minotaur/utils/super"
	"sync"
	"time"
)

// MessageHandler 定义了处理消息的函数类型
type MessageHandler[M any] func(message M)

// NewActor 创建一个新的 Actor，并启动其消息处理循环
func NewActor[M any](ctx context.Context, handler MessageHandler[M]) *Actor[M] {
	a := newActor(ctx, handler)
	a.counter = new(super.Counter[int])
	go a.run()
	return a
}

// newActor 创建一个新的 Actor
func newActor[M any](ctx context.Context, handler MessageHandler[M]) (actor *Actor[M]) {
	a := &Actor[M]{
		buf:     buffer.NewRing[M](1024),
		handler: handler,
	}
	a.cond = sync.NewCond(&a.rw)
	a.ctx, a.cancel = context.WithCancel(ctx)
	return a
}

// Actor 是一个消息驱动的并发实体
type Actor[M any] struct {
	idx     int                 // Actor 在其父 Actor 中的索引
	ctx     context.Context     // Actor 的上下文
	cancel  context.CancelFunc  // 用于取消 Actor 的函数
	buf     *buffer.Ring[M]     // 用于缓存消息的环形缓冲区
	handler MessageHandler[M]   // 处理消息的函数
	rw      sync.RWMutex        // 读写锁，用于保护 Actor 的并发访问
	cond    *sync.Cond          // 条件变量，用于触发消息处理流程
	counter *super.Counter[int] // 消息计数器，用于统计处理的消息数量
	dying   bool                // 标识 Actor 是否正在关闭中
	parent  *Actor[M]           // 父 Actor
	subs    []*Actor[M]         // 子 Actor 切片
	gap     []int               // 用于记录已经关闭的子 Actor 的索引位置，以便复用
}

// run 启动 Actor 的消息处理循环
func (a *Actor[M]) run() {
	var ctx = a.ctx
	var clearGap = time.NewTicker(time.Second * 30)
	defer func(a *Actor[M], clearGap *time.Ticker) {
		clearGap.Stop()
		a.cancel()
		a.parent.removeSub(a)
	}(a, clearGap)
	for {
		select {
		case <-a.ctx.Done():
			a.rw.Lock()
			if ctx == a.ctx {
				a.dying = true
			} else {
				ctx = a.ctx
			}
			a.rw.Unlock()
			a.cond.Signal()
		case <-clearGap.C:
			a.rw.Lock()
			for _, idx := range a.gap {
				a.subs = append(a.subs[:idx], a.subs[idx+1:]...)
			}
			for idx, sub := range a.subs {
				sub.idx = idx
			}
			a.gap = a.gap[:0]
			a.rw.Unlock()
		default:
			a.rw.Lock()
			if a.buf.IsEmpty() {
				if a.dying && a.counter.Val() == 0 {
					return
				}
				a.cond.Wait()
			}
			messages := a.buf.ReadAll()
			a.rw.Unlock()
			for _, message := range messages {
				a.handler(message)
			}
			a.counter.Add(-len(messages))
		}
	}
}

// Reuse 重用 Actor，Actor 会重新激活
func (a *Actor[M]) Reuse(ctx context.Context) {
	before := a.cancel
	defer before()

	a.rw.Lock()
	a.ctx, a.cancel = context.WithCancel(ctx)
	a.dying = false
	for _, sub := range a.subs {
		sub.Reuse(a.ctx)
	}
	a.rw.Unlock()
	a.cond.Signal()
}

// Send 发送消息
func (a *Actor[M]) Send(message M) {
	a.rw.Lock()
	a.counter.Add(1)
	a.buf.Write(message)
	a.rw.Unlock()
	a.cond.Signal()
}

// Sub 派生一个子 Actor，该子 Actor 生命周期将继承父 Actor 的生命周期
func (a *Actor[M]) Sub() {
	a.rw.Lock()
	defer a.rw.Unlock()

	sub := newActor(a.ctx, a.handler)
	sub.counter = a.counter.Sub()
	sub.parent = a
	if len(a.gap) > 0 {
		sub.idx = a.gap[0]
		a.gap = a.gap[1:]
	} else {
		sub.idx = len(a.subs)
	}
	a.subs = append(a.subs, sub)
	go sub.run()
}

// removeSub 从父 Actor 中移除指定的子 Actor
func (a *Actor[M]) removeSub(sub *Actor[M]) {
	if a == nil {
		return
	}

	a.rw.Lock()
	defer a.rw.Unlock()
	if sub.idx == len(a.subs)-1 {
		a.subs = a.subs[:sub.idx]
		return
	}
	a.subs[sub.idx] = nil
	a.gap = append(a.gap, sub.idx)
}
