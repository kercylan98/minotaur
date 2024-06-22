package vivid

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"sync/atomic"
)

var (
	_ core.MessageProcessor = &actorContext{}
	_ ActorContext          = &actorContext{}
)

const (
	actorStatusAlive uint32 = iota
	actorStatusTerminating
	actorStatusTerminated
)

func newActorContext(system *ActorSystem, actor Actor, parent ActorRef, ref ActorRef, mailbox *defaultMailbox) *actorContext {
	ctx := &actorContext{
		actorSystem: system,
		actor:       actor,
		parent:      parent,
		ref:         ref,
		mailbox:     mailbox,
		children:    haxmap.New[core.Address, ActorRef](),
	}
	return ctx
}

type actorContext struct {
	actorSystem *ActorSystem
	actor       Actor
	parent      ActorRef
	ref         ActorRef
	message     Message
	status      uint32 // atomic
	mailbox     Mailbox
	children    *haxmap.Map[core.Address, ActorRef]
}

func (ctx *actorContext) Terminate(target ActorRef) {
	ctx.actorSystem.processes.GetProcess(target).Terminate(ctx.ref)
}

func (ctx *actorContext) ActorOf(producer ActorProducer) ActorRef {
	return ctx.actorSystem.internalActorOf(new(ActorOptions).WithParent(ctx.ref), producer, func(child *actorContext) {
		// 确保在第一个消息处理之前添加到父级的子级列表中
		ctx.children.Set(child.ref.Address(), child.ref)
	})
}

func (ctx *actorContext) Parent() ActorRef {
	return ctx.parent
}

func (ctx *actorContext) Ref() ActorRef {
	if ctx.ref == nil {
		return nil
	}
	return ctx.ref
}

func (ctx *actorContext) Message() Message {
	return ctx.message
}

func (ctx *actorContext) Tell(target ActorRef, message vivid.Message) {
	ctx.actorSystem.processes.GetProcess(target).SendUserMessage(ctx.ref, message)
}

func (ctx *actorContext) ProcessUserMessage(msg core.Message) {
	if atomic.LoadUint32(&ctx.status) == actorStatusTerminated {
		return
	}

	ctx.message = msg
	ctx.actor.OnReceive(ctx)
}

func (ctx *actorContext) ProcessSystemMessage(msg core.Message) {
	switch m := msg.(type) {
	case OnTerminate:
		ctx.onTerminate()
	case OnTerminated:
		ctx.onTerminated(m)
	}
}

func (ctx *actorContext) String() string {
	return ctx.ref.Address().String()
}

func (ctx *actorContext) onTerminate() {
	if !atomic.CompareAndSwapUint32(&ctx.status, actorStatusAlive, actorStatusTerminating) {
		return
	}

	ctx.children.ForEach(func(address core.Address, ref ActorRef) bool {
		ctx.Terminate(ref)
		return true
	})

	ctx.onTerminated(OnTerminated{
		TerminatedActor: ctx.ref,
	})
}

func (ctx *actorContext) onTerminated(message OnTerminated) {
	ctx.children.Del(message.TerminatedActor.Address())
	if ctx.children.Len() > 0 {
		return
	}

	if !atomic.CompareAndSwapUint32(&ctx.status, actorStatusTerminating, actorStatusTerminated) {
		return
	}

	ctx.actorSystem.processes.Unregister(ctx.ref.Address())
	if ctx.parent != nil {
		ctx.actorSystem.processes.GetProcess(ctx.parent).SendSystemMessage(ctx.ref, OnTerminated{TerminatedActor: ctx.ref})
	}

	if dmb, ok := ctx.mailbox.(*defaultMailbox); ok {
		releaseDefaultMailbox(dmb)
	}

	fmt.Println("actor terminated", ctx.ref.Address().String())
}
