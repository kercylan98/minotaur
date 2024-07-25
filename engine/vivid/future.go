package vivid

import (
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"time"
)

// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 future.Future 对象，可被用于获取响应消息
//   - 当 timeout 参数为空时，将会使用默认的超时时间 DefaultFutureAskTimeout
func FutureAsk[M Message](ctx mixinDeliver, target ActorRef, message Message, timeout ...time.Duration) future.Future[M] {
	var t = DefaultFutureAskTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}
	var system *ActorSystem
	var c *actorContext
	switch v := ctx.(type) {
	case *ActorSystem:
		c = v.guard
		system = v
	case *actorContext:
		c = v
		system = c.system
	}

	f := future.New[M](c.system.rc, c.ref.DerivationProcessId(futureNamePrefix+convert.Uint64ToString(c.nextChildGuid())), t)
	system.rc.GetProcess(target).DeliveryUserMessage(target, f.Ref(), nil, message)
	return f
}
