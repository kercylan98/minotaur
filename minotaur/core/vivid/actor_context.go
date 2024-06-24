package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"sync/atomic"
)

var (
	_ core.MessageProcessor = &actorContext{}
	_ ActorContext          = &actorContext{}
	_ Supervisor            = &actorContext{}
)

const (
	actorStatusAlive uint32 = iota
	actorStatusTerminating
	actorStatusTerminated
	actorStatusRestarting
)

func newActorContext(system *ActorSystem, options *ActorOptions, producer ActorProducer, ref ActorRef, container mappings.OrderInterface[core.Address, ActorRef]) *actorContext {
	ctx := &actorContext{
		actorSystem: system,
		options:     options,
		actor:       producer(),
		producer:    producer,
		ref:         ref,
		children:    container,
	}
	return ctx
}

type actorContext struct {
	actorSystem *ActorSystem
	options     *ActorOptions
	ref         ActorRef
	message     Message
	actor       Actor
	status      uint32 // atomic
	mailbox     Mailbox
	producer    ActorProducer
	children    mappings.OrderInterface[core.Address, ActorRef]
}

func (ctx *actorContext) ProcessRecover(reason core.Message) {
	ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onSuspendMailbox)

	ctx.Escalate(&accident{
		accidentActor:      ctx.ref,
		reason:             reason,
		message:            ctx.Message(),
		supervisorStrategy: ctx.options.SupervisorStrategy,
	})
}

func (ctx *actorContext) BehaviorOf() Behavior {
	return newBehavior()
}

func (ctx *actorContext) Reply(message Message) {
	rm, ok := ctx.message.(RegulatoryMessage)
	if !ok || rm.Sender == nil {
		rm.Sender = ctx.System().deadLetter.ref
	}
	ctx.System().sendUserMessage(ctx.ref, rm.Sender, message)
}

func (ctx *actorContext) System() *ActorSystem {
	return ctx.actorSystem
}

func (ctx *actorContext) Terminate(target ActorRef) {
	ctx.System().getProcess(target).Terminate(ctx.ref)
}

func (ctx *actorContext) ActorOf(producer ActorProducer, options ...ActorOptionDefiner) ActorRef {
	return ctx.actorSystem.internalActorOf(new(ActorOptions).WithParent(ctx.ref), producer, options, func(child *actorContext) {
		// 确保在第一个消息处理之前添加到父级的子级列表中
		ctx.children.Set(child.ref.Address(), child.ref)
	})
}

func (ctx *actorContext) Parent() ActorRef {
	if ctx.options.Parent == nil {
		panic("root actor has no parent")
	}
	return ctx.options.Parent
}

func (ctx *actorContext) Ref() ActorRef {
	if ctx.ref == nil {
		return nil
	}
	return ctx.ref
}

func (ctx *actorContext) Message() Message {
	switch m := ctx.message.(type) {
	case RegulatoryMessage:
		return m.Message
	default:
		return ctx.message
	}
}

func (ctx *actorContext) Tell(target ActorRef, message vivid.Message, options ...MessageOption) {
	opts := generateMessageOptions(options...)
	defer releaseMessageOptions(opts)

	ctx.System().sendUserMessage(ctx.ref, target, message)
}

func (ctx *actorContext) FutureAsk(target ActorRef, message vivid.Message, options ...MessageOption) Future {
	opts := generateMessageOptions(options...)
	defer releaseMessageOptions(opts)

	if len(opts.MessageHooks) > 0 {
		cover := func(cover Message) {
			message = cover
		}
		opts.hookMessage(message, cover)
	}

	f := NewFuture(ctx.System(), opts.FutureTimeout)
	m := RegulatoryMessage{
		Sender:  f.Ref(),
		Message: message,
	}

	opts.hookRegulatoryMessage(&m)
	ctx.System().sendUserMessage(ctx.ref, target, m)
	return f
}

func (ctx *actorContext) Ask(target ActorRef, message vivid.Message, options ...MessageOption) {
	opts := generateMessageOptions(options...)
	defer releaseMessageOptions(opts)

	if len(opts.MessageHooks) > 0 {
		cover := func(cover Message) {
			message = cover
		}
		opts.hookMessage(message, cover)
	}

	m := RegulatoryMessage{
		Sender:  ctx.ref,
		Message: message,
	}

	opts.hookAskRegulatoryMessage(&m)
	opts.hookRegulatoryMessage(&m)
	ctx.System().sendUserMessage(ctx.ref, target, m)
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
	case _OnTerminated:
		switch atomic.LoadUint32(&ctx.status) {
		case actorStatusTerminated:
			ctx.onTerminated(m)
		case actorStatusRestarting:
			ctx.onRestarted(m)
		default:
			panic("unexpected status")
		}
	case Accident:
		ctx.onAccident(m)
	case _OnRestart:
		ctx.onRestart(m)
	}
}

func (ctx *actorContext) String() string {
	return ctx.ref.Address().String()
}

func (ctx *actorContext) onTerminate() {
	if !atomic.CompareAndSwapUint32(&ctx.status, actorStatusAlive, actorStatusTerminating) {
		return
	}

	ctx.ProcessUserMessage(onTerminate)

	ctx.children.Range(func(address core.Address, ref ActorRef) bool {
		ctx.Terminate(ref)
		return true
	})
}

func (ctx *actorContext) onTerminated(m _OnTerminated) {
	ctx.children.Del(m.TerminatedActor.Address())
	if ctx.children.Len() > 0 {
		return
	}

	if !atomic.CompareAndSwapUint32(&ctx.status, actorStatusTerminating, actorStatusTerminated) {
		return
	}

	ctx.actorSystem.processes.Unregister(ctx.ref)
	if ctx.Parent() != nil {
		ctx.System().sendSystemMessage(ctx.ref, ctx.Parent(), _OnTerminated{TerminatedActor: ctx.ref})
	}
}

func (ctx *actorContext) Children() []ActorRef {
	var children = make([]ActorRef, 0, ctx.children.Len())
	ctx.children.Range(func(address core.Address, ref ActorRef) bool {
		children = append(children, ref)
		return true
	})
	return children
}

func (ctx *actorContext) Stop(children ...ActorRef) {
	if len(children) == 0 {
		children = ctx.Children()
	}

	for _, ref := range children {
		ctx.System().sendSystemMessage(ctx.ref, ref, onTerminate)
	}

	// 确保没有子级的情况下能完成计数归零的逻辑
	ctx.onTerminated(_OnTerminated{TerminatedActor: ctx.ref})
}

func (ctx *actorContext) Resume(children ...ActorRef) {
	if len(children) == 0 {
		children = ctx.Children()
	}

	for _, ref := range children {
		ctx.System().sendSystemMessage(ctx.ref, ref, onResumeMailbox)
	}
}

func (ctx *actorContext) Escalate(accident Accident) {
	if parent := ctx.Parent(); parent != nil {
		ctx.System().sendSystemMessage(ctx.ref, parent, accident)
	} else {
		panic("the root actor should not continue to upgrade!")
	}
}

func (ctx *actorContext) Restart(children ...ActorRef) {
	if len(children) == 0 {
		children = ctx.Children()
	}

	for _, ref := range children {
		ctx.System().sendSystemMessage(ctx.ref, ref, onRestart)
	}
}

func (ctx *actorContext) onRestart(m _OnRestart) {
	atomic.StoreUint32(&ctx.status, actorStatusRestarting)
	ctx.System().sendUserMessage(ctx.ref, ctx.ref, onRestarting)
	ctx.children.Range(func(address core.Address, ref ActorRef) bool {
		ctx.Terminate(ref)
		return true
	})

	// 确保没有子级的情况下能完成计数归零的逻辑
	ctx.onRestarted(_OnTerminated{TerminatedActor: ctx.ref})
}

func (ctx *actorContext) onRestarted(m _OnTerminated) {
	ctx.children.Del(m.TerminatedActor.Address())
	if ctx.children.Len() > 0 {
		return
	}

	ctx.actor = ctx.producer()
	ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onResumeMailbox)
	ctx.System().sendUserMessage(ctx.ref, ctx.ref, onLaunch)
}

func (ctx *actorContext) onAccident(m Accident) {
	// 当责任人为空时，该父级理应是第一责任人
	m.trySetResponsible(ctx)

	// 如果指定了监管策略，那么执行
	if m.trySupervisorStrategy(ctx.System()) {
		return
	}

	// 如果该 Actor 本身实现了监管策略，那么执行
	if strategy, ok := ctx.actor.(SupervisorStrategy); ok {
		strategy.OnAccident(ctx.System(), m)
		return
	}

	// 否则继续升级
	ctx.Escalate(m)
}
