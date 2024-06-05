package vivid

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

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
	}
	s.core = new(_ActorSystemCore).init(&s)
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.BindDispatcher(new(_Dispatcher)) // default dispatcher
	s.BindMailboxFactory(NewFIFOFactory(s.onProcessMailboxMessage))
	s.BindMailboxFactory(NewPriorityFactory(s.onProcessMailboxMessage))
	var err error
	s.userGuard, err = generateActor(&s, new(UserGuardActor), parseActorOptions(NewActorOptions[*UserGuardActor]().WithName("user")))
	if err != nil {
		panic(err)
	}
	return s
}

type ActorSystem struct {
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
	userGuard         *_ActorCore        // 用户使用的顶级 Actor
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

func (s *ActorSystem) Context() context.Context {
	return s.ctx
}

func (s *ActorSystem) Shutdown() {
	defer s.cancel()
	s.unbindActor(s.userGuard)
	s.waitGroup.Wait()
}

func (s *ActorSystem) getSystem() *ActorSystem {
	return s
}

func (s *ActorSystem) GetDeadLetters() DeadLetterStream {
	return s.deadLetters
}

type actorOf interface {
	getSystem() *ActorSystem

	getContext() *_ActorCore
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

func (s *ActorSystem) unbindActor(actor ActorContext) {
	// 等待消息处理完毕后拒绝新消息
	core := actor.(*_ActorCore)
	core.Tell(OnDestroy{internal: true})
	core.messageGroup.Wait(func() {
		core.dispatcher.Detach(s.core, core)
		s.actorRW.Lock()
		delete(s.actors, core.GetId())
		s.actorRW.Unlock()
	})

	actor.getLockable().RLock()
	var children = actor.getChildren()
	actor.getLockable().RUnlock()

	for _, child := range children {
		s.unbindActor(child)
	}

	s.waitGroup.Done()
	//log.Debug("actor unbind", log.String("actor", core.GetId().String()))
}

func (s *ActorSystem) getActor(id ActorId) *_ActorCore {
	s.actorRW.RLock()
	defer s.actorRW.RUnlock()

	return s.actors[id]
}

func (s *ActorSystem) getContext() *_ActorCore {
	return s.userGuard
}

func (s *ActorSystem) sendToDispatcher(dispatcher Dispatcher, actor *_ActorCore, message MessageContext) {
	actor.messageGroup.Add(1)
	if !dispatcher.Send(s.core, actor, message) {
		actor.messageGroup.Done()
	}

	switch m := message.GetMessage().(type) {
	case OnDestroy:
		if !m.internal {
			s.unbindActor(actor)
		}
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
		if r := recover(); r != nil {
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error:   fmt.Errorf("%w: %v", ErrActorPanic, r),
				To:      core.GetId(),
				Message: message,
			}))
		}
	}()
	if core.messageHook != nil && !core.messageHook(message) {
		return
	}
	onReceive(core, message)
}

func generateActor[T Actor](system *ActorSystem, actor T, options *ActorOptions[T]) (*_ActorCore, error) {
	if options.Name == charproc.None {
		options.Name = uuid.NewString()
	}

	optionsNum := len(options.options)
	onReceive(actor, newMessageContext(system, OnOptionApply[T]{Options: options}, 0, false, false).withLocal(nil, nil))
	options.applyOption(options.options[optionsNum:]...)

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

	core.Tell(OnPreStart{})

	return core, nil
}
