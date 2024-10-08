package main

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
)

type CounterChangeEvent struct {
	Delta int
}

type CounterSnapshot struct {
	Count int
}

type MyActor struct {
	counter *CounterSnapshot
}

func (a *MyActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		// 初始化状态
		a.counter = &CounterSnapshot{}
	case *CounterSnapshot:
		a.counter = m
	case *CounterChangeEvent:
		a.counter.Count += m.Delta // 状态改变
		ctx.StateChanged(m)
	case int:
		// 模拟业务逻辑校验
		if m == 0 {
			return
		}
		// 校验通过，状态改变
		ctx.StateChangeEventApply(&CounterChangeEvent{Delta: m})
	case *vivid.OnPersistenceSnapshot:
		ctx.SaveSnapshot(a.counter)
	case error:
		ctx.ReportAbnormal(m) // 模拟崩溃
	case string:
		// 测试当前值
		fmt.Println("当前值", a.counter.Count)
	}
}

func main() {
	system := vivid.NewActorSystem()
	ref := system.ActorOfF(func() vivid.Actor {
		return new(MyActor)
	})

	for i := 0; i < 10000; i++ {
		system.Tell(ref, i)
	}
	system.Tell(ref, errors.New("panic"))
	system.Tell(ref, "test")

	system.Shutdown(true)
}
