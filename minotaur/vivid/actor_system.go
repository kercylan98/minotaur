package vivid

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/log"
	"math"
	"os"
	"path"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// NewActorSystem 创建一个 ActorSystem，ActorSystem 是 Actor 的容器，用于管理 Actor 的生命周期、消息分发等
func NewActorSystem(name string) ActorSystem {
	s := ActorSystem{
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
	}
	s.logger.Store(log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions().WithCallerSkip(5))).With(log.String("ActorSystem", name)))
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
	return s
}

// ActorSystem Actor 系统
type ActorSystem struct {
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
	server            Server
	client            Client
	messageSeq        *atomic.Uint64
	askWaits          map[uint64]chan<- Message
	askWaitsLock      *sync.RWMutex
	waitGroup         *toolkit.DynamicWaitGroup

	name    string // ActorSystem 名称
	network string // 网络类型
	host    string // 主机地址
	port    uint16 // 端口
	cluster string // 集群名称
}

func (s *ActorSystem) SetLogger(logger *log.Logger) {
	s.logger.Store(logger)
}

func (s *ActorSystem) GetLogger() *log.Logger {
	return s.logger.Load()
}

func (s *ActorSystem) ActorOf(ofo ActorOfO) ActorRef {
	return s.userGuardActor.ActorOf(ofo)
}

func (s *ActorSystem) Context() context.Context {
	return s.ctx
}

func (s *ActorSystem) Shutdown() {
	defer s.cancel()
	s.unbindActor(s.userGuardActor, false, true)
	s.unbindActor(s.eventBusActor, false, true)
	s.waitGroup.Wait()
}

func (s *ActorSystem) GetSystem() *ActorSystem {
	return s
}

func (s *ActorSystem) GetDeadLetters() DeadLetterStream {
	return s.deadLetters
}

func (s *ActorSystem) BindMailboxFactory(f MailboxFactory) MailboxFactoryId {
	s.mailboxFactorRW.Lock()
	defer s.mailboxFactorRW.Unlock()

	s.mailboxFactorGuid++
	s.mailboxFactors[s.mailboxFactorGuid] = f

	return s.mailboxFactorGuid
}

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

func (s *ActorSystem) BindDispatcher(d Dispatcher) DispatcherId {
	s.dispatcherRW.Lock()
	defer s.dispatcherRW.Unlock()

	s.dispatcherGuid++
	s.dispatchers[s.dispatcherGuid] = d

	return s.dispatcherGuid
}

func (s *ActorSystem) UnbindDispatcher(id DispatcherId) {
	if id == DefaultDispatcherId {
		return
	}
	s.dispatcherRW.Lock()
	defer s.dispatcherRW.Unlock()

	delete(s.dispatchers, id)
}

func (s *ActorSystem) unbindActor(actor ActorContext, restart, root bool, deadCallback ...func(core *_ActorCore)) {
	// 等待消息处理完毕后拒绝新消息
	core := actor.(*_ActorCore)
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

	var logType string
	if restart {
		logType = "restart"
	} else {
		logType = "stop"
	}
	s.GetLogger().Debug("unbindActor", log.String("type", logType), log.String("actor", actor.GetId().String()))

	// 服务解绑
	if core.runtimeMods != nil {
		core.UnloadMod(core.currentMods...)
		core.ApplyMod()
	}

	// 宣布 Actor 死亡
	s.waitGroup.Done()
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

func (s *ActorSystem) getActor(id ActorId) *_ActorCore {
	s.actorRW.RLock()
	defer s.actorRW.RUnlock()

	return s.actors[id]
}

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
		// 异步重启
		s.waitGroup.Add(1)
		go func(s *ActorSystem, actor *_ActorCore, restart bool) {
			defer s.waitGroup.Done()
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
	receiver.send(ctx)

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
			// 不应存在远程发送者，其他类型增加消息总计数
			s.waitGroup.Add(1)
			defer s.waitGroup.Done()
		}

		if opts.ReplyTimeout == 0 {
			opts.ReplyTimeout = time.Second
		}

		waiter := make(chan Message)
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

		select {
		case <-timeoutCtx.Done():
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: ErrMessageReplyTimeout,
			}))
		case reply := <-waiter:
			return reply
		}
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
			return
		}
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
	var actorId = NewActorId(system.network, system.cluster, system.host, system.port, system.name, actorPath)
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
	system.waitGroup.Add(1)
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
