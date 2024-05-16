package common

import (
	"fmt"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/components"
)

type UserActor struct {
	components.BehaviorAutoExecutor
}

func (u *UserActor) OnPreStart(ctx vivid.ActorContext) error {
	u.BehaviorAutoExecutor.Init(ctx)
	vivid.RegisterBehavior(ctx, u.onHello)
	vivid.RegisterBehavior(ctx, u.onEcho)
	return nil
}

func (u *UserActor) OnDestroy(ctx vivid.ActorContext) error {
	return nil
}

func (u *UserActor) OnChildTerminated(ctx vivid.ActorContext, child vivid.ActorTerminatedContext) {

}

func (u *UserActor) onHello(ctx vivid.MessageContext, msg string) error {
	fmt.Println(msg)
	return nil
}

func (u *UserActor) onEcho(ctx vivid.MessageContext, msg int) error {
	fmt.Println("ECHO", msg)
	return ctx.Reply(msg)
}
