package vivid

import (
	"reflect"
)

var actorType = reflect.TypeOf((*Actor)(nil)).Elem()

// Actor 是 Actor 模型的接口，该接口用于定义一个 Actor
type Actor interface {
	// OnPreStart 在 Actor 启动之前执行的逻辑，适用于对 Actor 状态的初始化
	OnPreStart(ctx *ActorContext) error

	// OnReceived 当 Actor 接收到消息时执行的逻辑
	OnReceived(msg Message) error
}

// localActor 实现 Actor 模型的核心逻辑
type localActor struct {
	opts    *ActorOptions
	actor   Actor
	ctx     *ActorContext
	mailbox *Mailbox
}

func (a *localActor) init(opts *ActorOptions, id ActorId, actor Actor, systemGetter actorSystemGetter) *localActor {
	a.opts = opts
	a.actor = actor
	a.ctx = new(ActorContext).init(id, a, systemGetter)
	a.mailbox = opts.Mailbox()

	go a.mailbox.Start()
	return a
}

func (a *localActor) GetId() ActorId {
	return a.ctx.id
}

func (a *localActor) Tell(msg Message, opts ...MessageOption) error {
	system := a.ctx.GetSystem()
	return system.tell(a, msg, opts...)
}

func (a *localActor) Stop() error {
	a.mailbox.Stop()
	return nil
}

// remoteActor 实现 Actor 模型的远程调用逻辑
type remoteActor struct {
	id     ActorId
	system *ActorSystem // 仅用于调用 tell 方法，并非 Actor 真正的所属系统
}

func (a *remoteActor) init(system *ActorSystem, id ActorId) *remoteActor {
	a.id = id
	a.system = system
	return a
}

func (a *remoteActor) GetId() ActorId {
	return a.id
}

func (a *remoteActor) Tell(msg Message, opts ...MessageOption) error {
	return a.system.tell(a, msg, opts...)
}
