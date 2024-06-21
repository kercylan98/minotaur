package vivid

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/minotaur/cluster"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"math"
	"net"
	"path"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

const (
	logTypeRestart = "restart"
	logTypeStop    = "stop"
)

// NewActorSystem 创建一个 ActorSystem，ActorSystem 是 Actor 的容器，用于管理 Actor 的生命周期、消息分发等
func NewActorSystem(name string, options ...*ActorSystemOptions) ActorSystem {
	s := ActorSystem{
		options:         new(ActorSystemOptions).apply(options...),
		dispatchers:     make(map[DispatcherId]Dispatcher),
		dispatcherRW:    new(sync.RWMutex),
		mailboxFactors:  make(map[MailboxFactoryId]MailboxFactory),
		mailboxFactorRW: new(sync.RWMutex),
		actors:          make(map[ActorId]*_ActorCore),
		actorRW:         new(sync.RWMutex),
		deadLetters:     new(_DeadLetterStream),
		codec:           new(_GobCodec),
		askWaits:        make(map[uint64]chan<- Message),
		askWaitsLock:    new(sync.RWMutex),
		messageSeq:      new(atomic.Uint64),
		waitGroup:       toolkit.NewDynamicWaitGroup(),
		name:            name,
		logger:          new(atomic.Pointer[log.Logger]),
		eventSeq:        new(atomic.Int64),
		scheduler:       chrono.NewScheduler(chrono.DefaultSchedulerTick, chrono.DefaultSchedulerWheelSize),
	}
	//s.waitGroup.ChangeHook = func(before, delta, curr int64) {
	//	log.Info("ActorSystem.WG", log.Int64("before", before), log.Int64("delta", delta), log.Int64("curr", curr))
	//	//debug.PrintStack()
	//}
	if s.options.Logger == nil {
		//s.options.Logger = log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions().WithCallerSkip(5))).With(log.String("ActorSystem", name))
		s.options.Logger = log.NewSilentLogger()
	}
	s.logger.Store(s.options.Logger)
	s.core = new(_ActorSystemCore).init(&s)
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.BindDispatcher(new(_Dispatcher)) // default dispatcher
	s.BindMailboxFactory(NewFIFOFactory(s.onProcessMailboxMessage))
	s.BindMailboxFactory(NewPriorityFactory(s.onProcessMailboxMessage))
	var err error
	s.userGuardActor, err = generateActor(&s, new(userGuardActor), parseActorOptions(NewActorOptions[*userGuardActor]().WithName("user")), false)
	if err != nil {
		panic(err)
	}

	s.eventBusActor, err = generateActor(&s, new(EventBusActor), parseActorOptions(NewActorOptions[*EventBusActor]().WithName("event_bus")), false)
	if err != nil {
		panic(err)
	}

	if len(s.options.ClusterOptions) > 0 {
		s.options.ClusterOptions = append([]cluster.Option{cluster.WithLogger(s.GetLogger())}, s.options.ClusterOptions...)
		s.cluster, err = cluster.NewNode(s.options.ClusterOptions...)
		if err != nil {
			panic(err)
		}

		go func(ctx context.Context, system *ActorSystem) {
			for {
				select {
				case <-ctx.Done():
					return
				case b := <-system.cluster.Read():
					system.onProcessServerMessage(b)
				}
			}
		}(s.ctx, &s)
	}

	return s
}

// ActorSystem Actor 系统
type ActorSystem struct {
	options           *ActorSystemOptions
	cluster           *cluster.Node
	logger            *atomic.Pointer[log.Logger]
	core              *_ActorSystemCore
	ctx               context.Context
	cancel            context.CancelFunc
	dispatchers       map[DispatcherId]Dispatcher
	dispatcherGuid    DispatcherId
	dispatcherRW      *sync.RWMutex
	mailboxFactors    map[MailboxFactoryId]MailboxFactory // mailbox factory
	mailboxFactorGuid MailboxFactoryId
	mailboxFactorRW   *sync.RWMutex
	actors            map[ActorId]*_ActorCore // actor id -> actor core
	actorRW           *sync.RWMutex
	deadLetters       *_DeadLetterStream // 死信队列
	userGuardActor    *_ActorCore        // 用户使用的顶级 Actor
	eventBusActor     *_ActorCore        // 事件总线 Actor
	eventSeq          *atomic.Int64      // 全局事件序号，仅标记事件产生顺序
	codec             Codec
	messageSeq        *atomic.Uint64            // 全局消息序号，用于生成唯一消息标识
	askWaits          map[uint64]chan<- Message // 等待回复的消息
	askWaitsLock      *sync.RWMutex             // 等待回复的消息锁
	waitGroup         *toolkit.DynamicWaitGroup // Actor 等待组
	scheduler         *chrono.Scheduler         // 调度器

	name string // ActorSystem 名称
}

// ClusterEnabled 是否启用集群
func (s *ActorSystem) ClusterEnabled() bool {
	return s.cluster != nil
}

// JoinCluster 加入集群
func (s *ActorSystem) JoinCluster(addresses ...string) error {
	if s.cluster == nil {
		return ErrClusterNotEnabled
	}
	self := net.JoinHostPort(s.cluster.GetHost(), convert.Uint16ToString(s.cluster.GetPort()))
	hasSelf := false
	for _, address := range addresses {
		if address == self {
			hasSelf = true
			break
		}
	}
	if !hasSelf {
		addresses = append([]string{self}, addresses...)
	}
	return s.cluster.Join(addresses...)
}

// SetLogger 设置日志记录器
func (s *ActorSystem) SetLogger(logger *log.Logger) {
	s.logger.Store(logger)
}

// GetLogger 获取日志记录器
func (s *ActorSystem) GetLogger() *log.Logger {
	return s.logger.Load()
}

// ActorOf 创建一个 Actor
func (s *ActorSystem) ActorOf(ofo ActorOfO) ActorRef {
	return s.userGuardActor.ActorOf(ofo)
}

// Context 获取 Actor 系统上下文
func (s *ActorSystem) Context() context.Context {
	return s.ctx
}

// Shutdown 关闭 Actor 系统
func (s *ActorSystem) Shutdown() {
	defer s.cancel()
	s.unbindActor(s.userGuardActor, false, true)
	s.unbindActor(s.eventBusActor, false, true)
	s.waitGroup.Wait() // 等待 Actor 销毁结束
	s.scheduler.Close()
}

// AwaitShutdown 等待 Actor 系统关闭
func (s *ActorSystem) AwaitShutdown() {
	s.waitGroup.Wait() // 等待 ActorSystem 关闭
}

// GetSystem 获取 Actor 系统
func (s *ActorSystem) GetSystem() *ActorSystem {
	return s
}

// GetDeadLetters 获取死信队列
func (s *ActorSystem) GetDeadLetters() DeadLetterStream {
	return s.deadLetters
}

// BindMailboxFactory 绑定邮箱工厂
func (s *ActorSystem) BindMailboxFactory(f MailboxFactory) MailboxFactoryId {
	s.mailboxFactorRW.Lock()
	defer s.mailboxFactorRW.Unlock()

	s.mailboxFactorGuid++
	s.mailboxFactors[s.mailboxFactorGuid] = f

	return s.mailboxFactorGuid
}

// UnbindMailboxFactory 解绑邮箱工厂
func (s *ActorSystem) UnbindMailboxFactory(id MailboxFactoryId) {
	if id == FIFOMailboxFactoryId {
		return
	}
	s.mailboxFactorRW.Lock()
	defer s.mailboxFactorRW.Unlock()

	delete(s.mailboxFactors, id)
}

func (s *ActorSystem) getMailboxFactory(id MailboxFactoryId) MailboxFactory {
	s.mailboxFactorRW.RLock()
	defer s.mailboxFactorRW.RUnlock()

	return s.mailboxFactors[id]
}

func (s *ActorSystem) getDispatcher(id DispatcherId) Dispatcher {
	s.dispatcherRW.RLock()
	defer s.dispatcherRW.RUnlock()

	return s.dispatchers[id]
}

// BindDispatcher 绑定分发器
func (s *ActorSystem) BindDispatcher(d Dispatcher) DispatcherId {
	s.dispatcherRW.Lock()
	defer s.dispatcherRW.Unlock()

	s.dispatcherGuid++
	s.dispatchers[s.dispatcherGuid] = d

	return s.dispatcherGuid
}

// UnbindDispatcher 解绑分发器
func (s *ActorSystem) UnbindDispatcher(id DispatcherId) {
	if id == DefaultDispatcherId {
		return
	}
	s.dispatcherRW.Lock()
	defer s.dispatcherRW.Unlock()

	delete(s.dispatchers, id)
}

func (s *ActorSystem) unbindActor(actor ActorContext, restart, root bool, deadCallback ...func(core *_ActorCore)) {
	core := actor.getCore()
	if !atomic.CompareAndSwapInt32(&core.status, actorStatusRunning, actorStatusTerminated) {
		return
	} else {
		core.Stop()
	}

	// 解绑空闲超时调度器
	s.scheduler.UnregisterTask(actor.GetId().String())

	// 等待消息处理完毕后拒绝新消息
	core.messageGroup.Wait(func() {
		core.dispatcher.Detach(s.core, core)
		s.actorRW.Lock()
		delete(s.actors, core.GetId())
		s.actorRW.Unlock()
	})

	// 重启顺序
	var restartOrder []*_ActorCore
	if restart {
		deadCallback = append(deadCallback, func(core *_ActorCore) {
			restartOrder = append([]*_ActorCore{core}, restartOrder...)
		})
	}

	// 递归解绑子 Actor
	func(actor ActorContext) {
		actor.getLockable().Lock()
		defer actor.getLockable().Unlock()
		var children = actor.getChildren()
		for name, child := range children {
			if child.stopOnParentRestart {
				s.unbindActor(child, false, false)
			} else {
				s.unbindActor(child, false, false, deadCallback...)
			}
			actor.unbindChild(name)
		}
	}(actor)

	// 服务解绑
	if core.runtimeMods != nil {
		core.UnloadMod(core.currentMods...)
		core.ApplyMod()
	}

	var logType string
	if restart {
		logType = logTypeRestart
	} else {
		logType = logTypeStop
	}
	s.GetLogger().Debug("unbindActor", log.String("type", logType), log.String("actor", actor.GetId().String()))

	// 宣布 Actor 死亡
	s.waitGroup.Done() // Actor 生命周期结束
	if !root {
		for _, f := range deadCallback {
			f(core)
		}
	} else {
		if parent := actor.getParent(); parent != nil {
			func(parent ActorContext) {
				parent.getLockable().Lock()
				defer parent.getLockable().Unlock()
				parent.unbindChild(actor.GetId().Name())
			}(parent)
		}
	}

	// 此刻 Actor 及其子 Actor 已经全部解绑
	if root && restart {
		restartOrder = append([]*_ActorCore{core}, restartOrder...)
		var restarted = make(map[ActorId]*_ActorCore)
		for i, c := range restartOrder {
			var parent = c.getParent().(*_internalActorContext).core
			if i > 0 {
				parent = restarted[parent.GetId()]
			}
			ctx := c.restartHandler(parent)
			c._LocalActorRef.core = ctx
			restarted[ctx.GetId()] = ctx
		}
	}
}

// GetActorRef 根据 ActorId 获取 Actor 引用，当 Actor 不存在时将会返回一个远程 Actor 引用
func (s *ActorSystem) GetActorRef(id ActorId) ActorRef {
	if core := s.getActor(id); core != nil {
		return core._LocalActorRef
	}
	return newRemoteActorRef(s, id)
}

func (s *ActorSystem) getActor(id ActorId) *_ActorCore {
	s.actorRW.RLock()
	defer s.actorRW.RUnlock()

	return s.actors[id]
}

// GetContext 获取 Actor 系统用户 Actor 上下文
func (s *ActorSystem) GetContext() ActorContext {
	return s.userGuardActor
}

func (s *ActorSystem) sendToDispatcher(dispatcher Dispatcher, actor *_ActorCore, message MessageContext) {
	actor.messageGroup.Add(1)

	if !dispatcher.Send(s.core, actor, message) {
		actor.messageGroup.Done()
	}

	switch m := message.GetMessage().(type) {
	case OnTerminate:
		// 异步停止
		s.waitGroup.Add(1) // 等待 Actor 销毁结束
		go func(s *ActorSystem, actor *_ActorCore, restart bool) {
			defer s.waitGroup.Done() // Actor 销毁结束
			s.unbindActor(actor, restart, true)
		}(s, actor, m.restart)
	}
}

func (s *ActorSystem) sendMessage(receiver ActorRef, message Message, options ...MessageOption) Message {
	var opts = new(MessageOptions).apply(options)

	ctx := newMessageContext(s, message, opts.Priority, opts.Instantly, opts.reply)
	switch ref := receiver.(type) {
	case *_LocalActorRef:
		ctx = ctx.withLocal(ref.core, opts.Sender)
	case *_RemoteActorRef:
		var senderId ActorId
		if sender, ok := opts.Sender.(*_LocalActorRef); ok {
			senderId = sender.core.GetId()
		}
		ctx = ctx.withRemote(ref.actorId, senderId)
	}
	if opts.ContextHook != nil {
		opts.ContextHook(ctx)
	}

	// 等待回复
	if opts.reply {
		// 如果等待回复，需要增加发送人等待计数
		switch sender := opts.Sender.(type) {
		case *_ActorCore:
			sender.messageGroup.Add(1)
			defer sender.messageGroup.Done()
		case *_LocalActorRef:
			sender.core.messageGroup.Add(1)
			defer sender.core.messageGroup.Done()
		default:
			// 不存在远程发送者，其他类型增加消息总计数
			s.waitGroup.Add(1)       // 等待 Ask 消息响应
			defer s.waitGroup.Done() // Ask 消息响应结束
		}

		if opts.ReplyTimeout == 0 {
			opts.ReplyTimeout = time.Second
		}

		waiter := make(chan Message, 1)
		timeoutCtx, cancel := context.WithTimeout(s.ctx, opts.ReplyTimeout)
		defer func(seq uint64, waiter chan Message, cancel context.CancelFunc) {
			cancel()
			close(waiter)
			s.askWaitsLock.Lock()
			delete(s.askWaits, seq)
			s.askWaitsLock.Unlock()
		}(ctx.GetSeq(), waiter, cancel)

		s.askWaitsLock.Lock()
		s.askWaits[ctx.GetSeq()] = waiter
		s.askWaitsLock.Unlock()
		receiver.send(ctx)

		select {
		case <-timeoutCtx.Done():
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: ErrMessageReplyTimeout,
			}))
		case reply := <-waiter:
			return reply
		}
	} else {
		receiver.send(ctx)
	}

	return nil
}

func (s *ActorSystem) onProcessServerMessage(bytes []byte) {
	var ctx = new(_MessageContext)
	if err := s.codec.Decode(bytes, ctx); err != nil {
		s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
			Error: err,
		}))
		return
	}
	ctx.system = s

	// 远程回复
	if ctx.RemoteReplySeq != 0 {
		s.askWaitsLock.RLock()
		wait, existWait := s.askWaits[ctx.RemoteReplySeq]
		s.askWaitsLock.RUnlock()
		if existWait {
			wait <- ctx.Message
		}
		return
	}

	// 查找接收者
	receiver := s.getActor(ctx.ReceiverId)
	if receiver == nil {
		s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
			Error: fmt.Errorf("%w: %s", ErrActorDeadOrNotExist, ctx.ReceiverId),
			From:  ctx.SenderId,
			To:    ctx.ReceiverId,
		}))
		return
	}
	ctx.actorContext = receiver

	// 远程消息增加计数，该计数将在消息处理完毕后减少
	s.sendToDispatcher(receiver.dispatcher, receiver, ctx)
}

func (s *ActorSystem) onProcessMailboxMessage(message MessageContext) {
	// received message
	core := message.GetRef().(*_LocalActorRef).core
	defer func() {
		core.messageGroup.Done()
	}()
	if core.messageHook != nil && !core.messageHook(message) {
		return
	}
	onReceive(core, message)
}

func generateActor[T Actor](system *ActorSystem, actor T, options *ActorOptions[T], restart bool) (*_ActorCore, error) {
	if !restart {
		if options.Name == charproc.None {
			options.Name = uuid.NewString()
		}

		optionsNum := len(options.options)
		onReceive(actor, newMessageContext(system, OnInit[T]{Options: options}, 0, false, false).withLocal(nil, nil))
		options.applyOption(options.options[optionsNum:]...)
	}

	var actorPath = options.Name
	if options.Parent != nil {
		actorPath = path.Join(options.Parent.GetId().Path(), options.Name)
	} else {
		actorPath = path.Clean(options.Name)
	}

	// 绝大多数情况均会成功，提前创建资源，减少锁粒度
	var actorId = NewActorId(system.cluster.GetClusterName(), system.cluster.GetHost(), system.cluster.GetPort(), system.name, actorPath)
	var core = newActorCore(system, actorId, actor, options)

	// 检查是否重名
	var parentLock *sync.RWMutex
	if options.Parent != nil {
		parentLock = options.Parent.getLockable()
		parentLock.Lock()
		if options.Parent.hasChild(options.Name) {
			parentLock.Unlock()
			return nil, fmt.Errorf("%w: %s", ErrActorAlreadyExists, options.Name)
		}
	}

	// 绑定分发器
	core.dispatcher = system.getDispatcher(options.DispatcherId)
	if core.dispatcher == nil {
		if parentLock != nil {
			parentLock.Unlock()
		}
		return nil, fmt.Errorf("%w: %d", ErrDispatcherNotFound, options.DispatcherId)
	}

	// 绑定邮箱
	core.mailboxFactory = system.getMailboxFactory(options.MailboxFactoryId)
	if core.mailboxFactory == nil {
		if parentLock != nil {
			parentLock.Unlock()
		}
		return nil, fmt.Errorf("%w: %d", ErrMailboxFactoryNotFound, options.MailboxFactoryId)
	}

	// 启动 Actor
	system.waitGroup.Add(1) // 等待 Actor 生命周期结束
	core.dispatcher.Attach(system.core, core)

	if options.Init != nil {
		options.Init(actor)
	}

	// 绑定父 Actor 并注册到系统
	system.actorRW.Lock()
	if options.Parent != nil {
		options.Parent.bindChild(options.Name, core)
	}
	system.actors[actorId] = core
	system.actorRW.Unlock()
	if parentLock != nil {
		parentLock.Unlock()
	}

	var logType string
	if restart {
		logType = "restart"
	} else {
		logType = "generate"
	}

	system.GetLogger().Debug("generateActor", log.String("type", logType), log.String("actor", actorId.String()))

	// 生命周期
	if restart {
		core.Tell(OnRestart{}, WithInstantly(true))
	}
	core.Tell(OnBoot{}, WithInstantly(true))

	if options.ActorContextHook != nil {
		system.waitGroup.Add(1)       // 等待上下文钩子执行完毕
		defer system.waitGroup.Done() // 上下文钩子执行完毕
		options.ActorContextHook(core)
	}

	return core, nil
}

func (s *ActorSystem) Subscribe(subscriber Subscriber, event Event, options ...SubscribeOption) {
	opts := new(SubscribeOptions).apply(options)
	s.eventBusActor.Tell(SubscribeMessage{
		Producer:        opts.Producer,
		Subscriber:      subscriber,
		Event:           reflect.TypeOf(event),
		Priority:        opts.Priority,
		PriorityTimeout: opts.PriorityTimeout,
	}, WithPriority(math.MinInt64), WithInstantly(true))
}

func (s *ActorSystem) Unsubscribe(subscriber Subscriber, event Event) {
	s.eventBusActor.Tell(UnsubscribeMessage{
		Event:      reflect.TypeOf(event),
		Subscriber: subscriber,
	}, WithPriority(math.MinInt64), WithInstantly(true))
}

func (s *ActorSystem) Publish(producer Producer, event Event) {
	s.eventBusActor.Tell(PublishMessage{
		Producer: producer,
		Event:    event,
	}, WithPriority(s.eventSeq.Add(1)))
}
