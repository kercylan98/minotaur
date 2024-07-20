package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid/internal/messages"
	"github.com/kercylan98/minotaur/engine/vivid/mailbox"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

const (
	DefaultFutureAskTimeout = time.Second
	futureNamePrefix        = "future-"
)

const (
	actorStatusAlive       uint32 = iota // Actor 存活状态
	actorStatusTerminating               // Actor 正在终止
	actorStatusTerminated                // Actor 已终止
	actorStatusRestarting                // Actor 正在重启
)

// ActorContext 是一个 Actor 完整的上下文，也是对外暴露的可用接口。
type ActorContext interface {
	mixinSpawner
	mixinDeliver
	mixinRecipient
	mixinWorker
	mixinScheduler
}

var _ ActorContext = (*actorContext)(nil)
var _ mailbox.Recipient = (*actorContext)(nil)
var _ supervision.Supervisor = (*actorContext)(nil)

func newActorContext(parent *actorContext, provider ActorProvider, descriptor *ActorDescriptor) (*actorContext, func(ref ActorRef)) {
	ctx := &actorContext{
		system:            parent.system,
		actor:             provider.Provide(),
		provider:          provider,
		parentRef:         parent.ref,
		children:          make(map[prc.LogicalAddress]ActorRef),
		accidentState:     supervision.NewAccidentState(),
		supervisorLoggers: descriptor.supervisionLoggers,
		idleDeadline:      descriptor.idleDeadline,
	}

	if descriptor.expireDuration > 0 {
		ctx.expireTime = time.Now().Add(descriptor.expireDuration)
	}

	if descriptor.internal != nil && descriptor.internal.parent != nil {
		ctx.parentRef = descriptor.internal.parent
	}

	if descriptor.supervisionStrategyProvider != nil {
		ctx.supervisorStrategy = descriptor.supervisionStrategyProvider.Provide()
		ctx.supervisorStrategyName = descriptor.supervisionStrategyProvider.GetStrategyProviderName()
	}

	return ctx, func(ref ActorRef) {
		if descriptor.internal == nil || descriptor.internal.parent == nil {
			parent.children[ref.LogicalAddress()] = ref
		}
		ctx.ref = ref
	}
}

type actorContext struct {
	system                 *ActorSystem                    // 所属 Actor 系统
	actor                  Actor                           // Actor 实例
	provider               ActorProvider                   // Actor 提供者
	ref                    ActorRef                        // 自身 Actor 引用
	parentRef              ActorRef                        // 父 Actor 引用
	children               map[prc.LogicalAddress]ActorRef // 子 Actor 引用表
	accidentState          *supervision.AccidentState      // Actor 事故状态
	status                 atomic.Uint32                   // Actor 状态
	childGuid              uint64                          // 子 Actor 自增 GUID 计数
	message                Message                         // 当前处理的消息
	sender                 ActorRef                        // 当前消息的发送者
	supervisorStrategy     supervision.Strategy            // 监管策略
	supervisorStrategyName supervision.StrategyName        // 监管策略名称
	supervisorLoggers      []supervision.Logger            // 监管记录器
	gracefullyTerminated   bool                            // 是否已优雅终止
	scheduler              *chrono.Scheduler               // 定时调度器
	schedulerInitializer   sync.Once                       // 调度器初始化锁
	expireTime             time.Time                       // 过期时间，即便是重启也不会被重置
	idleDeadline           time.Duration                   // 空闲截止时间
}

func (ctx *actorContext) initScheduler() {
	ctx.schedulerInitializer.Do(func() {
		ctx.scheduler = chrono.NewScheduler(chrono.DefaultSchedulerTick, chrono.DefaultSchedulerWheelSize)
	})
}

func (ctx *actorContext) Future(timeout ...time.Duration) future.Future {
	var t time.Duration
	if len(timeout) > 0 {
		t = timeout[0]
	}
	ctx.childGuid++
	return future.New(ctx.system.rc, ctx.ref.DerivationProcessId(futureNamePrefix+convert.Uint64ToString(ctx.childGuid)), t)
}

func (ctx *actorContext) CronTask(name, expression string, function func(ctx ActorContext)) error {
	ctx.initScheduler()
	return ctx.scheduler.RegisterCronTask(name, expression, func() {
		ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onSchedulerFunc(func() {
			function(ctx)
		}))
	})
}

func (ctx *actorContext) ImmediateCronTask(name, expression string, function func(ctx ActorContext)) error {
	ctx.initScheduler()
	return ctx.scheduler.RegisterImmediateCronTask(name, expression, func() {
		ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onSchedulerFunc(func() {
			function(ctx)
		}))
	})
}

func (ctx *actorContext) AfterTask(name string, after time.Duration, function func(ctx ActorContext)) {
	ctx.initScheduler()
	ctx.scheduler.RegisterAfterTask(name, after, func() {
		ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onSchedulerFunc(func() {
			function(ctx)
		}))
	})
}

func (ctx *actorContext) RepeatedTask(name string, after, interval time.Duration, times int, function func(ctx ActorContext)) {
	ctx.initScheduler()
	ctx.scheduler.RegisterRepeatedTask(name, after, interval, times, func() {
		ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onSchedulerFunc(func() {
			function(ctx)
		}))
	})
}

func (ctx *actorContext) DayMomentTask(name string, lastExecuted time.Time, offset time.Duration, hour, min, sec int, function func(ctx ActorContext)) {
	ctx.initScheduler()
	ctx.scheduler.RegisterDayMomentTask(name, lastExecuted, offset, hour, min, sec, func() {
		ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onSchedulerFunc(func() {
			function(ctx)
		}))
	})
}

func (ctx *actorContext) StopTask(name string) {
	if ctx.scheduler == nil {
		return
	}
	ctx.scheduler.UnregisterTask(name)
}

func (ctx *actorContext) Parent() ActorRef {
	return ctx.parentRef
}

func (ctx *actorContext) ReportAbnormal(reason Message) {
	if ctx.status.Load() != actorStatusAlive {
		return
	}

	ctx.accidentState.Record()
	ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onSuspendMailbox)

	// 产生事故记录
	record := supervision.NewAccidentRecord(ctx.sender, ctx.ref, nil, ctx.message, reason, ctx.supervisorStrategy, ctx.accidentState)

	// 用户日志记录器
	for _, logger := range ctx.supervisorLoggers {
		logger.Log(record)
	}

	// 自身故障，由上级处理，升级故障
	// 默认监管者为空，父级收到时将成为监管者，如果穿透多级 Actor，那么监管者将会逐渐升级
	ctx.Escalate(record)
}

func (ctx *actorContext) Escalate(record *supervision.AccidentRecord) {
	if ctx.parentRef == nil {
		// 顶级 Actor 异常，可能是未考虑的情况
		panic(fmt.Errorf("the root actor should not continue to upgrade!, err: %v", record.Reason))
	}

	if !ctx.system.rc.Belong(ctx.parentRef) {
		// 远程的父级 Actor，需要通过网络传输
		car, err := record.CrossAccidentRecord(ctx.supervisorStrategyName, ctx.system.shared)
		if err != nil {
			panic(fmt.Errorf("cross accident record failed, err: %v", err))
		}
		ctx.system.rc.GetProcess(ctx.parentRef).DeliverySystemMessage(ctx.parentRef, ctx.ref, nil, car)
		return
	}

	ctx.system.rc.GetProcess(ctx.parentRef).DeliverySystemMessage(ctx.parentRef, ctx.ref, nil, record)
}

func (ctx *actorContext) onAccidentRecordProcess(m *supervision.AccidentRecord) {
	// 升级为监管者
	m.Supervisor = ctx

	// 如果指定了监管策略，或者责任人本身实现了监管策略，那么执行，否则继续升级
	if m.Strategy != nil {
		m.Strategy.OnPolicyDecision(m)
	} else if strategy, ok := ctx.actor.(supervision.Strategy); ok {
		strategy.OnPolicyDecision(m)
	} else {
		ctx.Escalate(m) // 继续升级
	}
}

func (ctx *actorContext) onCrossAccidentRecordProcess(m *supervision.CrossAccidentRecord) {
	var strategy supervision.Strategy
	if m.Strategy != charproc.None {
		provider, exist := ctx.system.config.supervisionStrategyProviderTable[m.Strategy]
		if !exist {
			panic(fmt.Errorf("supervisor strategy %s not found", m.Strategy))
		}
		strategy = provider.Provide()
	}
	ar, err := m.AccidentRecord(ctx, ctx.system.shared, strategy)
	if err != nil {
		panic(fmt.Errorf("cross accident record failed, err: %v", err))
	}
	ctx.onAccidentRecordProcess(ar)
}

func (ctx *actorContext) Restart(refs ...*prc.ProcessRef) {
	for _, ref := range refs {
		ctx.system.rc.GetProcess(ref).DeliverySystemMessage(ref, ctx.ref, nil, onRestart)
	}
}

func (ctx *actorContext) Stop(refs ...*prc.ProcessRef) {
	for _, ref := range refs {
		ctx.Terminate(ref, false)
	}
	ctx.tryTerminated()
}

func (ctx *actorContext) Resume(refs ...*prc.ProcessRef) {
	for _, ref := range refs {
		ctx.system.rc.GetProcess(ref).DeliverySystemMessage(ref, ctx.ref, nil, onResumeMailbox)
	}
}

func (ctx *actorContext) Message() Message {
	return ctx.message
}

func (ctx *actorContext) Sender() ActorRef {
	return ctx.sender
}

func (ctx *actorContext) Terminate(target ActorRef, gracefully bool) {
	if gracefully {
		ctx.Tell(target, onGracefullyTerminate)
	} else {
		ctx.system.rc.GetProcess(target).DeliverySystemMessage(target, ctx.ref, nil, onTerminate)
	}
}

func (ctx *actorContext) refreshIdleDeadline(reset bool) {
	if ctx.idleDeadline <= 0 {
		return
	}
	if reset {
		ctx.StopTask(":idle:")
	} else {
		ctx.AfterTask(":idle:", ctx.idleDeadline, func(ctx ActorContext) {
			ctx.Terminate(ctx.Ref(), true)
		})
	}
}

func (ctx *actorContext) processMessage(sender, receiver ActorRef, message Message, system bool) {
	ctx.refreshIdleDeadline(true)
	defer ctx.refreshIdleDeadline(false)

	ctx.message = message
	ctx.sender = sender
	if !system {
		switch m := message.(type) {
		case *OnTerminate:
			if m.Gracefully {
				ctx.gracefullyTerminated = true
				ctx.Terminate(ctx.ref, false)
				return
			}
		}
		ctx.actor.OnReceive(ctx)

		switch message.(type) {
		case *OnLaunch:
			ctx.accidentState.Solved()
		}
		return
	}

	switch m := message.(type) {
	case onSchedulerFunc:
		m()
	case *OnLaunch:
		ctx.processMessage(sender, receiver, m, false)
	case *OnRestarted:
		ctx.processMessage(sender, receiver, m, false)
	case *OnTerminate:
		ctx.onTerminate(m.Gracefully)
	case *Terminated: // 转换为 OnTerminated
		ctx.onTerminated(&OnTerminated{TerminatedActor: prc.NewProcessRef(m.TerminatedProcess)})
	case *onRestartMessage:
		ctx.onRestart()
	case *supervision.AccidentRecord:
		ctx.onAccidentRecordProcess(m)
	case *supervision.CrossAccidentRecord:
		ctx.onCrossAccidentRecordProcess(m)
	case *messages.GenerateRemoteActor:
		ctx.onGenerateRemoteActor(m)
	}
}

func (ctx *actorContext) ProcessUserMessage(message prc.Message) {
	sender, receiver, message := unwrapMessage(message)
	ctx.processMessage(sender, receiver, message, false)
}

func (ctx *actorContext) ProcessSystemMessage(message prc.Message) {
	sender, receiver, message := unwrapMessage(message)
	ctx.processMessage(sender, receiver, message, true)
}

func (ctx *actorContext) ProcessAccident(reason prc.Message) {
	ctx.ReportAbnormal(reason)
}

func (ctx *actorContext) Tell(target ActorRef, message Message) {
	ctx.system.rc.GetProcess(target).DeliveryUserMessage(target, nil, nil, message)
}

func (ctx *actorContext) Ask(target ActorRef, message Message) {
	ctx.system.rc.GetProcess(target).DeliveryUserMessage(target, ctx.ref, nil, message)
}

func (ctx *actorContext) FutureAsk(target ActorRef, message Message, timeout ...time.Duration) future.Future {
	var t = DefaultFutureAskTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}
	ctx.childGuid++
	f := ctx.Future(t)
	ctx.system.rc.GetProcess(target).DeliveryUserMessage(target, f.Ref(), nil, message)
	return f
}

func (ctx *actorContext) Broadcast(message Message) {
	for _, child := range ctx.children {
		ctx.Ask(child, message)
	}
}

func (ctx *actorContext) Reply(message Message) {
	ctx.Ask(ctx.sender, message)
}

func (ctx *actorContext) System() *ActorSystem {
	return ctx.system
}

func (ctx *actorContext) Ref() ActorRef {
	return ctx.ref
}

// findClusterNode 寻找符合条件的集群节点，如果为空，那么没有满足的
func (ctx *actorContext) findClusterNode(provider ActorProvider, descriptor *ActorDescriptor) *prc.DiscoverNode {
	if descriptor.fixedLocal {
		return nil
	}
	system := ctx.system
	if system.discoverer == nil {
		return nil
	}
	if descriptor.internal != nil {
		return nil
	}
	if descriptor.supervisionStrategyProvider != nil && descriptor.supervisionStrategyProvider.GetStrategyProviderName() != charproc.None {
		return nil
	}
	if len(descriptor.supervisionLoggers) > 0 {
		return nil
	}

	var matches []*prc.DiscoverNode
	for _, node := range system.discoverer.GetNodes() {
		md := node.Metadata()

		actorCheck := md.UserMetadata[actorSystemMetadataKeyActorProviderTable] == provider.GetActorProviderName()
		dispatcherCheck := actorCheck && md.UserMetadata[actorSystemMetadataKeyDispatcherProviderTable] == descriptor.dispatcherProvider.GetDispatcherProviderName()
		mailboxCheck := dispatcherCheck && md.UserMetadata[actorSystemMetadataKeyMailboxProviderTable] == descriptor.mailboxProvider.GetMailboxProviderName()
		strategyCheck := mailboxCheck && (descriptor.supervisionStrategyProvider == nil || md.UserMetadata[actorSystemMetadataKeySupervisionStrategyProviderTable] == descriptor.supervisionStrategyProvider.GetStrategyProviderName())

		if !strategyCheck {
			continue
		}

		matches = append(matches, node)
	}

	if len(matches) == 0 {
		return nil
	}

	// 排序，暂时随机
	collection.Shuffle(&matches)
	return matches[0]
}

func (ctx *actorContext) ActorOf(provider ActorProvider, configurator ...ActorDescriptorConfigurator) ActorRef {
	// 生成描述并配置
	descriptor := newActorDescriptor()
	for _, c := range configurator {
		c.Configure(descriptor)
	}

	// 名称及前缀初始化
	if descriptor.name == charproc.None {
		ctx.childGuid++
		descriptor.name = convert.Uint64ToString(ctx.childGuid)
	}
	if descriptor.namePrefix != charproc.None {
		descriptor.name = descriptor.namePrefix + "-" + descriptor.name
	}

	if node := ctx.findClusterNode(provider, descriptor); node != nil && ctx.system.rc.GetPhysicalAddress() != node.Metadata().GetRcPhysicalAddress() {
		// 集群中生成
		msg := &messages.GenerateRemoteActor{
			ParentPid:              ctx.ref.GetId(),
			ProviderName:           provider.GetActorProviderName(),
			ActorName:              descriptor.name,
			ActorNamePrefix:        descriptor.namePrefix,
			MailboxProviderName:    descriptor.mailboxProvider.GetMailboxProviderName(),
			DispatcherProviderName: descriptor.dispatcherProvider.GetDispatcherProviderName(),
			ExpireDuration:         int64(descriptor.expireDuration),
			IdleDeadline:           int64(descriptor.idleDeadline),
		}
		if descriptor.supervisionStrategyProvider != nil {
			msg.SupervisionStrategyName = descriptor.supervisionStrategyProvider.GetStrategyProviderName()
		}

		remoteRef := prc.NewProcessRef(prc.NewClusterProcessId(node.Metadata().UserMetadata[actorSystemClusterName], node.Metadata().RcPhysicalAddress, "/"))

		f := ctx.Future()
		ctx.system.rc.GetProcess(remoteRef).DeliverySystemMessage(remoteRef, f.Ref(), nil, msg)

		ref, err := f.Result()
		if err != nil {
			panic(fmt.Errorf("generate remote actor failed, err: %v", err))
		}
		return prc.NewProcessRef(ref.(*messages.GenerateRemoteActorResult).Pid)
	}

	// 进程 Id 初始化
	var processId *prc.ProcessId
	if descriptor.internal != nil && descriptor.internal.parent != nil {
		processId = descriptor.internal.parent.DerivationProcessId(descriptor.name)
		processId.ClusterName = ctx.system.Name()
		processId.PhysicalAddress = ctx.system.rc.GetPhysicalAddress()
	} else {
		processId = ctx.ref.DerivationProcessId(descriptor.name)
	}

	// 创建上下文
	ctx, refBinder := newActorContext(ctx, provider, descriptor)

	// 初始化分发器及邮箱
	mb := descriptor.mailboxProvider.Provide(descriptor.dispatcherProvider.Provide(), ctx)

	// 创建进程
	process := newActorProcess(mb)
	ref, exist := ctx.system.rc.Register(processId, process)
	if exist {
		panic(fmt.Errorf("actor %s already exists", processId.LogicalAddress))
	}

	// 绑定 ActorRef
	refBinder(ref)

	// 内部钩子
	if descriptor.internal != nil && descriptor.internal.actorContextHook != nil {
		descriptor.internal.actorContextHook(ctx)
	}

	ctx.system.logger().Debug("ActorSystem", log.String("event", "launch"), log.String("type", reflect.TypeOf(ctx.actor).String()), log.String("actor", processId.LogicalAddress), log.Int("child", len(ctx.children)))

	// 第一条消息
	ctx.system.rc.GetProcess(ref).DeliverySystemMessage(ref, ctx.ref, nil, onLaunch)

	ctx.setExpireDuration()
	return ref
}

func (ctx *actorContext) setExpireDuration() {
	if ctx.expireTime.IsZero() {
		return
	}
	ctx.AfterTask(":expire:", ctx.expireTime.Sub(time.Now()), func(ctx ActorContext) {
		ctx.Terminate(ctx.Ref(), true)
	})
}

func (ctx *actorContext) ActorOfF(provider FunctionalActorProvider, configurator ...FunctionalActorDescriptorConfigurator) ActorRef {
	var c = make([]ActorDescriptorConfigurator, len(configurator))
	for i, f := range configurator {
		c[i] = f
	}
	return ctx.ActorOf(provider, c...)
}

func (ctx *actorContext) Children() []ActorRef {
	return collection.ConvertMapValuesToSlice(ctx.children)
}

func (ctx *actorContext) onRestart() {
	ctx.status.Store(actorStatusRestarting)

	ctx.processMessage(ctx.sender, ctx.ref, onRestarting, false)

	for _, ref := range ctx.children {
		ctx.Terminate(ref, false)
	}

	ctx.tryRestarted()
}

func (ctx *actorContext) tryRestarted() {
	if len(ctx.children) > 0 || ctx.status.Load() != actorStatusRestarting {
		return
	}

	ctx.processMessage(ctx.sender, ctx.ref, onTerminate, false)
	ctx.processMessage(ctx.sender, ctx.ref, &OnTerminated{ctx.ref}, false)

	ctx.actor = ctx.provider.Provide()
	if ctx.scheduler != nil {
		ctx.scheduler.Clear()
	}
	ctx.status.Store(actorStatusAlive)

	ctx.system.logger().Debug("ActorSystem", log.String("event", "restarted"), log.String("type", reflect.TypeOf(ctx.actor).String()), log.String("actor", ctx.ref.LogicalAddress()), log.Int("child", len(ctx.children)))

	ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onResumeMailbox)
	ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onRestarted)
	ctx.system.rc.GetProcess(ctx.ref).DeliverySystemMessage(ctx.ref, ctx.ref, nil, onLaunch)

	ctx.setExpireDuration()
}

func (ctx *actorContext) onTerminate(gracefully bool) {
	if !ctx.status.CompareAndSwap(actorStatusAlive, actorStatusTerminating) {
		return
	}
	ctx.processMessage(ctx.sender, ctx.ref, onTerminate, false)

	for _, ref := range ctx.children {
		ctx.Terminate(ref, gracefully || ctx.gracefullyTerminated)
	}

	ctx.tryTerminated()
}

func (ctx *actorContext) onTerminated(terminated *OnTerminated) {
	delete(ctx.children, terminated.TerminatedActor.LogicalAddress())

	ctx.processMessage(ctx.sender, ctx.ref, terminated, false)
	switch ctx.status.Load() {
	case actorStatusTerminating:
		ctx.tryTerminated()
	case actorStatusRestarting:
		ctx.tryRestarted()
	default:
	}
}

func (ctx *actorContext) tryTerminated() {
	if len(ctx.children) > 0 {
		return
	}

	if !ctx.status.CompareAndSwap(actorStatusTerminating, actorStatusTerminated) {
		return
	}

	terminatedMessage := &OnTerminated{TerminatedActor: ctx.ref}
	ctx.processMessage(ctx.sender, ctx.ref, terminatedMessage, false)
	ctx.system.rc.Unregister(ctx.sender, ctx.ref)
	if ctx.scheduler != nil {
		ctx.scheduler.Close()
	}

	ctx.system.logger().Debug("ActorSystem", log.String("event", "terminated"), log.String("type", reflect.TypeOf(ctx.actor).String()), log.String("actor", ctx.ref.LogicalAddress()), log.Int("child", len(ctx.children)))

	if ctx.parentRef != nil {
		ctx.system.rc.GetProcess(ctx.parentRef).DeliverySystemMessage(ctx.parentRef, ctx.ref, nil, &Terminated{TerminatedProcess: ctx.ref.GetId()})
	} else {
		close(ctx.system.closed)
	}
}

func (ctx *actorContext) onGenerateRemoteActor(m *messages.GenerateRemoteActor) {
	provider, exist := ctx.system.config.actorProviderTable[m.ProviderName]
	if !exist {
		panic(fmt.Errorf("actor provider %s not found", m.ProviderName))
	}

	descriptor := newActorDescriptor()
	descriptor.name = m.ActorName
	descriptor.namePrefix = m.ActorNamePrefix
	descriptor.mailboxProvider = ctx.system.config.mailboxProviderTable[m.MailboxProviderName]
	descriptor.dispatcherProvider = ctx.system.config.dispatcherProviderTable[m.DispatcherProviderName]
	descriptor.supervisionStrategyProvider = ctx.system.config.supervisionStrategyProviderTable[m.SupervisionStrategyName]
	descriptor.idleDeadline = time.Duration(m.IdleDeadline)
	descriptor.expireDuration = time.Duration(m.ExpireDuration)

	ref := ctx.ActorOf(provider, FunctionalActorDescriptorConfigurator(func(descriptor *ActorDescriptor) {
		descriptor.withInternalDescriptor(&actorInternalDescriptor{
			useDescriptor: descriptor,
			parent:        prc.NewProcessRef(m.ParentPid),
		})
	}))

	ctx.Reply(&messages.GenerateRemoteActorResult{Pid: ref.GetId()})
}
