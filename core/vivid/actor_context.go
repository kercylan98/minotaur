package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/random"
	"sync/atomic"
	"time"
)

var (
	_ core.MessageProcessor = &actorContext{}
	_ ActorContext          = &actorContext{}
	_ Supervisor            = &actorContext{}
)

const (
	actorStatusAlive       uint32 = iota // Actor 存活状态
	actorStatusTerminating               // Actor 正在终止
	actorStatusTerminated                // Actor 已终止
	actorStatusRestarting                // Actor 正在重启
)

func newActorContext(system *ActorSystem, parent ActorRef, options *ActorOptions, producer ActorProducer, ref ActorRef, container mappings.OrderInterface[core.Address, ActorRef], kind Kind) *actorContext {
	ctx := &actorContext{
		actorSystem: system,
		childGuid:   new(atomic.Uint64),
		parent:      parent,
		actor:       producer(),
		producer:    producer,
		ref:         ref,
		children:    container,
		as:          &accidentState{},
		persistenceStatus: &persistenceStatus{
			persistenceStorage: options.PersistenceStorage,
			persistenceName:    options.PersistenceName,
			eventLimit:         options.PersistenceEventLimit,
		},
		supervisorStrategy: options.SupervisorStrategy,
		supervisorLogger:   options.SupervisorLogger,
		kind:               kind,
	}
	if options.Scheduler || options.Deadline+options.MessageDeadline > 0 {
		ctx.scheduler = chrono.NewScheduler(chrono.DefaultSchedulerTick, chrono.DefaultSchedulerWheelSize)
		ctx.scheduler.SetExecutor(func(name string, caller func(), separate func()) {
			caller()
		})
	}
	if options.Deadline > 0 {
		ctx.AfterTask("deadline"+random.HostName(), options.Deadline, func(ctx ActorContext) {
			system.opts.LoggerProvider().Debug("ActorContext", log.String("actor", ctx.Ref().Address().String()), log.String("status", "deadline"))
			ctx.TerminateGracefully(ctx.Ref())
		})
	}
	if options.MessageDeadline > 0 {
		ctx.messageDeadlineKey = "message_deadline" + random.HostName()
		ctx.messageDeadline = options.MessageDeadline
	}
	ctx.persistenceStatus.ctx = ctx
	if ctx.persistenceStatus.persistenceStorage == nil {
		ctx.persistenceStatus.persistenceStorage = defaultStorage
	}
	if ctx.persistenceStatus.persistenceName == "" {
		ctx.persistenceStatus.persistenceName = ref.Address().LogicPath()
	}
	if ctx.persistenceStatus.eventLimit <= 0 {
		ctx.persistenceStatus.eventLimit = DefaultPersistenceEventLimit
	}
	return ctx
}

type actorContext struct {
	actorSystem        *ActorSystem                                    // Actor 系统
	parent             ActorRef                                        // 父 Actor 引用
	childGuid          *atomic.Uint64                                  // 子 Actor 的 guid 当前值
	dispatcher         Dispatcher                                      // Actor 使用的调度器
	ref                ActorRef                                        // 该 Actor 引用
	message            Message                                         // 当前正在处理的消息（可能为包装）
	messageFail        bool                                            // 当前处理的消息是否失败
	actor              Actor                                           // Actor 实例
	mailbox            Mailbox                                         // Actor 使用的邮箱
	producer           ActorProducer                                   // Actor 生产者
	children           mappings.OrderInterface[core.Address, ActorRef] // 子 Actor
	as                 *accidentState                                  // 该 Actor 的事故状态
	persistenceStatus  *persistenceStatus                              // 该 Actor 的持久化状态
	scheduler          *chrono.Scheduler                               // Actor 使用的定时器
	supervisorStrategy SupervisorStrategy                              // Actor 使用的监督者策略
	supervisorLogger   []SupervisorLogger                              // Actor 使用的监督者日志
	kind               Kind                                            // Actor Kind 类型
	status             uint32                                          // 原子状态
	messageDeadlineKey string                                          // 消息处理超时任务的名称
	messageDeadline    time.Duration                                   // 消息处理空闲超时时间
}

func (ctx *actorContext) CronTask(name, expression string, function func(ctx ActorContext)) error {
	if ctx.scheduler != nil {
		return ctx.scheduler.RegisterCronTask(name, expression, func() {
			ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onSchedulerFunc(func() {
				function(ctx)
			}))
		})
	}
	return nil
}

func (ctx *actorContext) ImmediateCronTask(name, expression string, function func(ctx ActorContext)) error {
	if ctx.scheduler != nil {
		return ctx.scheduler.RegisterImmediateCronTask(name, expression, func() {
			ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onSchedulerFunc(func() {
				function(ctx)
			}))
		})
	}
	return nil
}

func (ctx *actorContext) AfterTask(name string, after time.Duration, function func(ctx ActorContext)) {
	if ctx.scheduler != nil {
		ctx.scheduler.RegisterAfterTask(name, after, func() {
			ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onSchedulerFunc(func() {
				function(ctx)
			}))
		})
	}
}

func (ctx *actorContext) RepeatedTask(name string, after, interval time.Duration, times int, function func(ctx ActorContext)) {
	if ctx.scheduler != nil {
		ctx.scheduler.RegisterRepeatedTask(name, after, interval, times, func() {
			ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onSchedulerFunc(func() {
				function(ctx)
			}))
		})
	}
}

func (ctx *actorContext) DayMomentTask(name string, lastExecuted time.Time, offset time.Duration, hour, min, sec int, function func(ctx ActorContext)) {
	if ctx.scheduler != nil {
		ctx.scheduler.RegisterDayMomentTask(name, lastExecuted, offset, hour, min, sec, func() {
			ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onSchedulerFunc(func() {
				function(ctx)
			}))
		})
	}
}

func (ctx *actorContext) StopTask(name string) {
	if ctx.scheduler != nil {
		ctx.scheduler.UnregisterTask(name)
	}
}

func (ctx *actorContext) IsKind() bool {
	return ctx.kind != charproc.None
}

func (ctx *actorContext) Kind() Kind {
	return ctx.kind
}

func (ctx *actorContext) DeadLetter() DeadLetter {
	return ctx.System().deadLetter
}

func (ctx *actorContext) PersistSnapshot(snapshot Message) {
	ctx.persistenceStatus.PersistSnapshot(snapshot)
	ctx.persistenceStatus.persistenceStorage.Persist(ctx.persistenceStatus.persistenceName, ctx.persistenceStatus)
	ctx.persistenceStatus.persistentDone = true
}

func (ctx *actorContext) StatusChanged(event Message) {
	ctx.persistenceStatus.StatusChanged(event)
	ctx.persistenceStatus.persistenceStorage.Persist(ctx.persistenceStatus.persistenceName, ctx.persistenceStatus)
}

func (ctx *actorContext) Sender() ActorRef {
	rm, ok := ctx.message.(RegulatoryMessage)
	if !ok || rm.Sender == nil {
		return ctx.System().deadLetter.ref
	}
	return rm.Sender
}

func (ctx *actorContext) HasSender() bool {
	rm, ok := ctx.message.(RegulatoryMessage)
	if !ok || rm.Sender == nil {
		return false
	}
	return true
}

func (ctx *actorContext) ReportAbnormal(reason Message) {
	if atomic.LoadUint32(&ctx.status) != actorStatusAlive {
		return
	}
	ctx.messageFail = true
	ctx.as.restartTimes = append(ctx.as.restartTimes, time.Now())
	ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onSuspendMailbox)

	message := ctx.Message()
	for _, logger := range ctx.supervisorLogger {
		logger(reason, message)
	}

	ctx.Escalate(&accident{
		accidentActor:      ctx.ref,
		reason:             reason,
		message:            message,
		supervisorStrategy: ctx.supervisorStrategy,
		state:              ctx.as,
	})
}

func (ctx *actorContext) ProcessRecover(reason core.Message) {
	ctx.ReportAbnormal(reason)
}

func (ctx *actorContext) BehaviorOf() Behavior {
	return newBehavior()
}

func (ctx *actorContext) Reply(message Message) {
	rm, ok := ctx.message.(RegulatoryMessage)
	if !ok || rm.Sender == nil {
		rm.Sender = ctx.System().deadLetter.ref
	}
	ctx.Ask(rm.Sender, message)
}

func (ctx *actorContext) TryReply(message Message) bool {
	rm, ok := ctx.message.(RegulatoryMessage)
	if !ok || rm.Sender == nil {
		return false
	}
	ctx.Ask(rm.Sender, message)
	return true
}

func (ctx *actorContext) System() *ActorSystem {
	return ctx.actorSystem
}

func (ctx *actorContext) Terminate(target ActorRef) {
	ctx.System().getProcess(target).Terminate(target)
}

func (ctx *actorContext) TerminateGracefully(target ActorRef) {
	ctx.System().Tell(target, onTerminateGracefully)
}

func (ctx *actorContext) KindOf(kind Kind, parent ...ActorRef) ActorRef {
	if len(parent) > 0 {
		if ctx.ref.Address().PhysicalAddress() == parent[0].Address().PhysicalAddress() {
			return ctx.localKindOf(kind, parent...)
		}

		f := NewFuture(ctx.System(), time.Second)

		ctx.System().sendSystemMessage(f.Ref(), parent[0], wrapRegulatoryMessage(f.Ref(), parent[0], &KindOf{Kind: kind, ParentAddress: []byte(parent[0].Address())}))
		result, err := f.Result()
		if err != nil {
			return ctx.System().deadLetter.Ref()
		}
		var addr = result.(*ActorRefAddress).Address
		return core.NewProcessRef(core.Address(addr))
	}
	return ctx.localKindOf(kind)
}

func (ctx *actorContext) localKindOf(kind Kind, parent ...ActorRef) ActorRef {
	ctx.actorSystem.kindRw.RLock()
	defer ctx.actorSystem.kindRw.RUnlock()
	kindInfo, exist := ctx.actorSystem.kinds[kind]
	if !exist {
		return ctx.System().deadLetter.Ref()
	}

	var parentRef = ctx.ref
	if len(parent) > 0 {
		parentRef = parent[0]
	}

	opts := actorOptionsPool.Get().WithParent(parentRef)
	defer actorOptionsPool.Put(opts)
	ref, err := ctx.actorSystem.internalActorOf(
		opts,
		kindInfo.producer,
		[]ActorOptionDefiner{func(options *ActorOptions) {
			options.options = append(options.options, kindInfo.options.options...)
		}},
		func(child *actorContext) {
			// 确保在第一个消息处理之前添加到父级的子级列表中
			ctx.children.Set(child.ref.Address(), child.ref)
		},
		ctx.childGuid,
		kind,
	)
	if err != nil && !opts.ConflictReuse {
		panic(err)
	}
	return ref
}

func (ctx *actorContext) ActorOf(producer ActorProducer, options ...ActorOptionDefiner) ActorRef {
	opts := actorOptionsPool.Get().WithParent(ctx.ref)
	defer actorOptionsPool.Put(opts)
	ref, err := ctx.actorSystem.internalActorOf(
		opts,
		producer,
		options,
		func(child *actorContext) { // 确保在第一个消息处理之前添加到父级的子级列表中
			ctx.System().sendSystemMessage(ctx.ref, opts.Parent, onBindChildren{child.ref})
		},
		ctx.childGuid,
		charproc.None,
	)
	if err != nil && !opts.ConflictReuse {
		panic(err)
	}
	return ref
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
	_, _, m := unwrapRegulatoryMessage(ctx.message)
	return m
}

func (ctx *actorContext) Tell(target ActorRef, message Message, options ...MessageOption) {
	opts := generateMessageOptions(options...)
	defer releaseMessageOptions(opts)

	if len(opts.MessageHooks) > 0 {
		cover := func(cover Message) {
			message = cover
		}
		opts.hookMessage(message, cover)
	}

	m := wrapRegulatoryMessage(nil, target, message)

	opts.hookRegulatoryMessage(&m)
	ctx.System().sendUserMessage(ctx.ref, target, m)
}

func (ctx *actorContext) FutureAsk(target ActorRef, message Message, options ...MessageOption) Future {
	opts := generateMessageOptions(options...)
	defer releaseMessageOptions(opts)

	if len(opts.MessageHooks) > 0 {
		cover := func(cover Message) {
			message = cover
		}
		opts.hookMessage(message, cover)
	}

	f := NewFuture(ctx.System(), opts.FutureTimeout)

	m := wrapRegulatoryMessage(f.Ref(), target, message)

	opts.hookRegulatoryMessage(&m)
	ctx.System().sendUserMessage(ctx.ref, target, m)
	return f
}

func (ctx *actorContext) Ask(target ActorRef, message Message, options ...MessageOption) {
	opts := generateMessageOptions(options...)
	defer releaseMessageOptions(opts)

	if len(opts.MessageHooks) > 0 {
		cover := func(cover Message) {
			message = cover
		}
		opts.hookMessage(message, cover)
	}

	m := wrapRegulatoryMessage(ctx.ref, target, message)

	opts.hookRegulatoryMessage(&m)
	ctx.System().sendUserMessage(ctx.ref, target, m)
}

func (ctx *actorContext) Broadcast(message Message, options ...MessageOption) {
	opts := generateMessageOptions(options...)
	defer releaseMessageOptions(opts)

	if len(opts.MessageHooks) > 0 {
		cover := func(cover Message) {
			message = cover
		}
		opts.hookMessage(message, cover)
	}

	m := wrapRegulatoryMessage(ctx.ref, nil, message)

	for _, ref := range ctx.Children() {
		m.Receiver = ref
		opts.hookRegulatoryMessage(&m)
		ctx.System().sendUserMessage(ctx.ref, ref, m)
	}
}

func (ctx *actorContext) AwaitForward(target ActorRef, blockFunc func() Message) {
	f := NewFuture(ctx.System(), 0)
	f.Forward(target)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				ctx.System().sendUserMessage(f.Ref(), f.Ref(), err)
			}
		}()
		message := blockFunc()
		ctx.System().sendUserMessage(f.Ref(), f.Ref(), message)
	}()
}

func (ctx *actorContext) clearMessageDeadline() {
	if ctx.messageDeadline > 0 {
		if _, ok := ctx.message.(onSchedulerFunc); ok {
			return
		}
		ctx.StopTask(ctx.messageDeadlineKey)
	}
}

func (ctx *actorContext) setMessageDeadline() {
	if ctx.messageDeadline > 0 {
		if _, ok := ctx.message.(onSchedulerFunc); ok {
			return
		}
		ctx.AfterTask(ctx.messageDeadlineKey, ctx.messageDeadline, func(ctx ActorContext) {
			ctx.TerminateGracefully(ctx.Ref())
			ctx.System().opts.LoggerProvider().Debug("ActorContext", log.String("actor", ctx.Ref().Address().String()), log.String("status", "message_deadline"))
		})
	}
}

func (ctx *actorContext) ProcessUserMessage(msg core.Message, recoveryMessage ...Message) {
	ctx.clearMessageDeadline()
	defer ctx.setMessageDeadline()
	if atomic.LoadUint32(&ctx.status) == actorStatusTerminated {
		return
	}

	defer func() {
		if len(recoveryMessage) > 0 {
			ctx.message = recoveryMessage[0]
		}
	}()

	ctx.message = msg
	ctx.messageFail = false
	ctx.actor.OnReceive(ctx)

	switch msg.(type) {
	case OnLaunch:
		if !ctx.messageFail {
			ctx.as.restartTimes = ctx.as.restartTimes[:0]
		}
	case TerminateGracefully:
		ctx.onTerminate(true)
	default:
		switch ctx.Message().(type) {
		case TerminateGracefully:
			ctx.onTerminate(true)
		}
	}
}

func (ctx *actorContext) ProcessSystemMessage(msg core.Message) {
	ctx.clearMessageDeadline()
	defer ctx.setMessageDeadline()
	ctx.message = msg
	switch m := msg.(type) {
	case OnLaunch:
		if status := ctx.persistenceStatus.persistenceStorage.Load(ctx.persistenceStatus.persistenceName); status != nil {
			ctx.persistenceStatus.recovery = true
			defer func() {
				ctx.persistenceStatus.recovery = false
			}()
			ctx.ProcessUserMessage(status.GetSnapshot(), msg)
			for _, event := range status.GetEvents() {
				ctx.ProcessUserMessage(event, msg)
			}
			ctx.ProcessUserMessage(m, msg)
		} else {
			ctx.ProcessUserMessage(m, msg)
		}
	case onBindChildren:
		ctx.children.Set(m.ChildrenRef.Address(), m.ChildrenRef)
	case OnTerminate:
		ctx.onTerminate(false)
	case OnTerminated:
		switch atomic.LoadUint32(&ctx.status) {
		case actorStatusAlive, actorStatusTerminating:
			ctx.onTerminated(m)
		case actorStatusRestarting:
			ctx.onRestarted(m)
		default:
			// 自身为其他状态时无需等待子 Actor 全部关闭或重启，忽略
		}
	case Accident:
		ctx.onAccident(m)
	case OnRestart:
		ctx.onRestart(m)
	case OnPersistenceSnapshot:
		ctx.ProcessUserMessage(m, msg)
	case onSchedulerFunc:
		m()
	default:
		switch m := ctx.Message().(type) {
		case *KindOf:
			parentRef := core.NewProcessRef(ActorId(m.ParentAddress))
			ref := ctx.localKindOf(m.Kind, parentRef)

			ctx.Reply(&ActorRefAddress{Address: []byte(ref.Address())})
		case OnTerminate:
			ctx.ProcessSystemMessage(m)
		}
	}
}

func (ctx *actorContext) String() string {
	return ctx.ref.Address().String()
}

func (ctx *actorContext) onTerminate(gracefully bool) {
	if !atomic.CompareAndSwapUint32(&ctx.status, actorStatusAlive, actorStatusTerminating) {
		return
	}

	ctx.ProcessUserMessage(onTerminate)

	ctx.children.Range(func(address core.Address, ref ActorRef) bool {
		if gracefully {
			ctx.TerminateGracefully(ref)
		} else {
			ctx.Terminate(ref)
		}
		return true
	})

	// 确保没有子级的情况下能完成计数归零的逻辑
	ctx.onTerminated(OnTerminated{TerminatedActor: ctx.ref})
}

func (ctx *actorContext) onTerminated(m OnTerminated) {
	ctx.children.Del(m.TerminatedActor.Address())
	if ctx.children.Len() > 0 {
		return
	}

	ctx.actorSystem.processes.Unregister(ctx.ref, func() {
		ctx.ProcessUserMessage(m, OnTerminated{TerminatedActor: ctx.ref})
	})

	if !atomic.CompareAndSwapUint32(&ctx.status, actorStatusTerminating, actorStatusTerminated) {
		return
	}

	system := ctx.System()
	system.opts.LoggerProvider().Debug("ActorContext", log.String("actor", ctx.ref.Address().String()), log.String("status", "terminated"))
	if ctx.scheduler != nil {
		ctx.scheduler.Close()
	}
	if parent := ctx.Parent(); parent != nil {
		system.sendSystemMessage(ctx.ref, parent, OnTerminated{TerminatedActor: ctx.ref})
	} else {
		close(system.closed)
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
	ctx.onTerminated(OnTerminated{TerminatedActor: ctx.ref})
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
		panic(fmt.Errorf("the root actor should not continue to upgrade!, err: %v", accident.Reason()))
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

func (ctx *actorContext) onRestart(m OnRestart) {
	atomic.StoreUint32(&ctx.status, actorStatusRestarting)
	defer func() {
		if atomic.LoadUint32(&ctx.status) == actorStatusRestarting {
			if err := recover(); err != nil {
				ctx.System().opts.LoggerProvider().Error("ActorContext", log.String("actor", ctx.ref.Address().String()), log.String("status", "restart panic"), log.Any("err", err))
			}
		}
	}()
	ctx.ProcessUserMessage(onTerminate)
	ctx.ProcessUserMessage(onRestarting)
	ctx.children.Range(func(address core.Address, ref ActorRef) bool {
		ctx.Terminate(ref)
		return true
	})

	// 确保没有子级的情况下能完成计数归零的逻辑
	ctx.onRestarted(OnTerminated{TerminatedActor: ctx.ref})
}

func (ctx *actorContext) onRestarted(m OnTerminated) {
	ctx.children.Del(m.TerminatedActor.Address())
	if ctx.children.Len() > 0 {
		return
	}

	ctx.actor = ctx.producer()
	if ctx.scheduler != nil {
		ctx.scheduler.Clear()
	}
	atomic.StoreUint32(&ctx.status, actorStatusAlive)
	ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onResumeMailbox)
	ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onLaunch)
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
