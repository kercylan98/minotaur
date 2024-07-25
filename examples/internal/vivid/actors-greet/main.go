package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"sync"
)

type (
	GreetMessage      struct{ Name string }
	GreetMessageReply struct {
		Name    string
		Content string
	}
)

func NewGreetActor(wg *sync.WaitGroup, name string, neighbor ...vivid.ActorRef) *GreetActor {
	return &GreetActor{wg: wg, name: name, neighbor: neighbor}
}

type GreetActor struct {
	wg       *sync.WaitGroup
	name     string
	neighbor []vivid.ActorRef
}

func (g *GreetActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		g.onLaunch(ctx)
	case *GreetMessage:
		g.onGreet(ctx, m)
	case *GreetMessageReply:
		g.onGreetReply(ctx, m)
	case *vivid.OnTerminated:
		g.wg.Done()
	}
}

func (g *GreetActor) onLaunch(ctx vivid.ActorContext) {
	m := &GreetMessage{Name: g.name}
	for _, ref := range g.neighbor {
		ctx.Ask(ref, m)
	}
}

func (g *GreetActor) onGreet(ctx vivid.ActorContext, m *GreetMessage) {
	fmt.Println(fmt.Sprintf("%s 收到了来自 %s 的招呼", g.name, m.Name))
	ctx.Reply(&GreetMessageReply{Name: g.name, Content: fmt.Sprintf("Hi %s, I'm %s", m.Name, g.name)})

	ctx.Terminate(ctx.Ref(), true)
}

func (g *GreetActor) onGreetReply(ctx vivid.ActorContext, m *GreetMessageReply) {
	fmt.Println(fmt.Sprintf("%s 收到了来自 %s 的回复：%s", g.name, m.Name, m.Content))

	ctx.Terminate(ctx.Ref(), true)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	actor1 := system.ActorOfF(func() vivid.Actor {
		return NewGreetActor(wg, "王老五")
	})

	system.ActorOfF(func() vivid.Actor {
		return NewGreetActor(wg, "李老四", actor1)
	})

	wg.Wait()
}
