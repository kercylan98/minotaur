package vivid

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit"
	"sync"
	"time"
)

type ActorCore interface {
	Actor
	ActorContext
	ActorRef
	GetMailboxFactory() MailboxFactory
	GetMailbox() Mailbox
	BindMailbox(Mailbox)
	ModifyMessageCounter(delta int64)
}

// newActorCore 创建一个新的 ActorCore
func newActorCore[T Actor](system *ActorSystem, id ActorId, actor T, options *ActorOptions[T]) *_ActorCore {
	core := &_ActorCore{
		id:             id,
		Actor:          actor,
		_LocalActorRef: &_LocalActorRef{},
		_ActorContext: &_ActorContext{
			_internalActorContext: &_internalActorContext{},
		},
		system:              system,
		parent:              options.Parent,
		childrenRW:          new(sync.RWMutex),
		children:            make(map[ActorName]*_ActorCore),
		messageGroup:        toolkit.NewDynamicWaitGroup(),
		messageHook:         options.MessageHook,
		supervisor:          options.Supervisor,
		stopOnParentRestart: options.StopOnParentRestart,
		idleTimeout:         options.IdleTimeout,
		restartHandler: func(parent *_ActorCore) *_ActorCore {
			options.Parent = parent
			ctx, err := generateActor(system, actor, options, true)
			if err != nil {
				system.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeActorOf, DeadLetterEventActorOf{
					Error:  err,
					Parent: options.Parent,
					Name:   options.Name,
				}))
			}
			return ctx
		},
	}
	if core.parent != nil {
		core.Context = context.WithoutCancel(options.Parent)
	} else {
		core.Context = context.WithoutCancel(system.ctx)
	}

	core._LocalActorRef.core = core
	core._ActorContext._ActorCore = core
	core._ActorContext._internalActorContext._ActorCore = core

	return core
}

type _ActorCore struct {
	context.Context                                          // 上下文
	Actor                                                    // Actor 对象
	*_ActorContext                                           // ActorContext
	*_LocalActorRef                                          // ActorRef
	id                  ActorId                              // Actor ID
	parent              ActorContext                         // 父 ActorContext
	childrenRW          *sync.RWMutex                        // 保护 children 的读写锁
	children            map[ActorName]*_ActorCore            // 通过名字索引子 ActorContext，用于范围查找
	isEnd               bool                                 // 是否为终结节点
	system              *ActorSystem                         // Actor 所属的 Actor 系统
	dispatcher          Dispatcher                           // Actor 所绑定的 Dispatcher
	mailboxFactory      MailboxFactory                       // Actor 所绑定的 MailboxFactory
	mailbox             Mailbox                              // Actor 所绑定的 Mailbox
	messageGroup        *toolkit.DynamicWaitGroup            // 等待消息处理完毕的消息组
	messageHook         func(MessageContext) bool            // 消息钩子
	supervisor          Supervisor                           // 监管策略
	restartHandler      func(parent *_ActorCore) *_ActorCore // 重启处理器
	stopOnParentRestart bool                                 // 父 Actor 重启时是否停止
	idleTimeout         time.Duration                        // 空闲超时时间
}

func (a *_ActorCore) GetContext() ActorContext {
	return a._ActorContext
}

func (a *_ActorCore) GetMailboxFactory() MailboxFactory {
	return a.mailboxFactory
}

func (a *_ActorCore) GetMailbox() Mailbox {
	return a.mailbox
}

func (a *_ActorCore) BindMailbox(mailbox Mailbox) {
	a.mailbox = mailbox
}

func (a *_ActorCore) ModifyMessageCounter(delta int64) {
	a.core.messageGroup.Add(delta)
}

func (a *_ActorCore) GetSystem() *ActorSystem {
	return a.system
}

func (a *_ActorCore) refreshIdleTimeout() {
	if idleTimeout := a.idleTimeout; idleTimeout > 0 {
		id := a.id.String()
		a.system.scheduler.RegisterAfterTask(id, idleTimeout, func() {
			a.Terminate()
			a.system.GetLogger().Debug("ActorIdleTimeout", "actor", id, "timeout", idleTimeout)
		})
	}
}
