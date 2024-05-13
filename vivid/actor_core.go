package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

// newActorCore 创建一个 actorCore
func newActorCore(id ActorId, actor Actor) *actorCore {
	a := &actorCore{
		ActorId:  id,
		actor:    actor,
		messages: buffer.NewRing[*context](1024),
		cond:     sync.NewCond(&sync.Mutex{}),
		closed:   make(chan struct{}),
	}
	return a
}

// actorCore Actor 的核心结构体，负责 Actor 的消息分发、状态管理等
type actorCore struct {
	ActorId                         // 唯一标识
	actor    Actor                  // Actor 状态
	status   actorStatus            // Actor 运行状态
	cond     *sync.Cond             // 消息队列条件变量
	messages *buffer.Ring[*context] // 消息缓冲区
	closed   chan struct{}          // 关闭信号
}

func (a *actorCore) getId() ActorId {
	return a.ActorId
}

func (a *actorCore) getActor() Actor {
	return a.actor
}

func (a *actorCore) start() {
	if !atomic.CompareAndSwapInt32(&a.status, actorStatusNone, actorStatusRunning) {
		return
	}
	defer func(a *actorCore) {
		close(a.closed)
	}(a)

	for {
		a.cond.L.Lock()
		messages := a.messages.ReadAll()
		if len(messages) == 0 {
			// 此刻队列没有消息且处于停止中状态，将会关闭 Actor
			if atomic.CompareAndSwapInt32(&a.status, actorStatusStopping, actorStatusStopped) {
				a.cond.L.Unlock()
				return
			}
			a.cond.Wait()

			// 重新读取消息
			messages = a.messages.ReadAll()
		}
		a.cond.L.Unlock()

		for i := 0; i < len(messages); i++ {
			m := messages[i]
			if err := a.onReceive(m); err != nil {
				m.err = err
			}
		}
	}
}

func (a *actorCore) add(message *context) {
	if atomic.LoadInt32(&a.status) == actorStatusStopped {
		return
	}

	a.cond.L.Lock()
	a.messages.Write(message)
	a.cond.Signal()
	a.cond.L.Unlock()
}

func (a *actorCore) stop() {
	a.cond.L.Lock()
	if !atomic.CompareAndSwapInt32(&a.status, actorStatusRunning, actorStatusStopping) {
		a.cond.L.Unlock()
		return
	}
	// 避免循环位于等待状态且一直没有新消息进入无法退出循环
	a.cond.Signal()
	a.cond.L.Unlock()

	<-a.closed
	return
}

func (a *actorCore) onReceive(message *context) error {
	defer func() {
		close(message.done)
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	return a.getActor().OnReceive(message)
}
