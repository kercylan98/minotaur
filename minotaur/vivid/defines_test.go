package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/log"
	"sync"
)

type (
	ExportUserGuardActor = userGuardActor
)

var (
	actorSystem     *ActorSystem
	actorSystemLock sync.Mutex
)

func GetTestActorSystem() *ActorSystem {
	actorSystemLock.Lock()
	defer actorSystemLock.Unlock()
	if actorSystem == nil {
		sys := NewActorSystem("test", NewActorSystemOptions().WithLogger(log.GetDefault()))
		actorSystem = &sys
	}
	return actorSystem
}

func GetBenchmarkTestSystem() *ActorSystem {
	actorSystemLock.Lock()
	defer actorSystemLock.Unlock()
	if actorSystem == nil {
		sys := NewActorSystem("benchmark", NewActorSystemOptions().WithLogger(log.NewSilentLogger()))
		actorSystem = &sys
	}
	return actorSystem
}

// IneffectiveActor 无效的 Actor，仅实现了 Actor 接口，但是没有任何行为。用于测试用途
type IneffectiveActor struct {
}

func (i *IneffectiveActor) OnReceive(ctx MessageContext) {}

// PrintlnActor 仅包含一个打印消息的行为的 Actor，用于测试用途
type PrintlnActor struct {
	ActorRef
}

func (p *PrintlnActor) OnReceive(ctx MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case OnBoot:
		p.ActorRef = ctx
	case string:
		fmt.Println(m)
	}
}

func (p *PrintlnActor) Println(message string) {
	p.Tell(message)
}

// PrintlnActorTyped 是 PrintlnActor 的类型化引用，定义了一个消息发送函数，用于测试用途
type PrintlnActorTyped interface {
	ActorRef

	Println(message string)
}

// PanicActor 在接收到 error 类型消息时将触发 panic，用于测试用途
type PanicActor struct {
	RestartHook func()
	PanicHook   func(err error)
}

func (p *PanicActor) OnReceive(ctx MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case OnRestart:
		if p.RestartHook != nil {
			p.RestartHook()
		}
		ctx.Stop()
	case error:
		if p.PanicHook != nil {
			p.PanicHook(m)
		}
		panic(fmt.Errorf("panic actor panic: %w", m))
	}
}
