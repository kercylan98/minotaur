package vivid

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/log"
	"reflect"
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
		waitGroup:       new(sync.WaitGroup),
		name:            name,
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.BindDispatcher(new(_Dispatcher)) // default dispatcher
	s.BindMailboxFactory(NewFIFOFactory(func(message MessageContext) {
		// received message
		core := message.GetReceiver().(*_ActorCore)
		defer core.group.Done()
		core.OnReceive(message)
	}))
	var err error
	s.userGuard, err = generateActor(&s, new(UserGuardActor), parseActorOptions(NewActorOptions[*UserGuardActor]().WithName("user")))
	if err != nil {
		panic(err)
	}
	return s
}

type ActorSystem struct {
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
	waitGroup         *sync.WaitGroup

	name    string
	network string
	host    string
	port    uint16
	cluster string
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
}

func (s *ActorSystem) BindMailboxFactory(f MailboxFactory) MailboxFactoryId {
	s.mailboxFactorRW.Lock()
	defer s.mailboxFactorRW.Unlock()

	s.mailboxFactorGuid++
	s.mailboxFactors[s.mailboxFactorGuid] = f

	return s.mailboxFactorGuid
}

func (s *ActorSystem) UnbindMailboxFactory(id MailboxFactoryId) {
	if id == DefaultMailboxFactoryId {
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
	actor.getLockable().RLock()
	var children = actor.getChildren()
	for _, child := range children {
		s.unbindActor(child)
	}
	actor.getLockable().RUnlock()

	core := actor.(*_ActorCore)
	core.group.Wait()
	core.dispatcher.Detach(core)

	s.actorRW.Lock()
	delete(s.actors, core.GetId())
	s.actorRW.Unlock()

	s.waitGroup.Done()
	log.Debug("actor unbind", log.String("actor", core.GetId().String()))
}

func (s *ActorSystem) sendMessage(receiver ActorRef, message Message, options ...MessageOption) Message {
	var opts = new(MessageOptions).apply(options)
	var seq = s.messageSeq.Add(1)
	var from, to ActorId

	switch ref := receiver.(type) {
	case *_LocalActorRef:
		if opts.replySeq > 0 {
			s.askWaitsLock.RLock()
			wait, exist := s.askWaits[opts.replySeq]
			s.askWaitsLock.RUnlock()
			if !exist {
				return nil
			}
			wait <- message
			return nil
		}

		ctx := &_LocalMessageContext{
			ActorContext: ref.core,
			message:      message,
			seq:          seq,
			network:      s.network,
			host:         s.host,
			port:         s.port,
		}
		if opts.Sender != nil {
			ctx.sender = opts.Sender
			from = opts.Sender.(*_LocalActorRef).core.GetId()
		}
		to = ref.core.GetId()

		ref.core.group.Add(1)
		ref.core.dispatcher.Send(ref.core, ctx)
	case *_RemoteActorRef:
		ctx := &_RemoteMessageContext{
			system:     s,
			ReceiverId: ref.actorId,
			Message:    message,
			Seq:        seq,
			ReplySeq:   opts.replySeq,
		}
		if opts.Sender != nil {
			ctx.SenderId = opts.Sender.(*_LocalActorRef).core.GetId()
		}
		from = ctx.SenderId
		to = ctx.ReceiverId

		data, err := s.codec.Encode(ctx)
		if err != nil {
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: err,
				From:  from,
				To:    to,
				Seq:   seq,
			}))
			return nil
		}

		// send data to remote actor
		if err = s.client.Exec(data); err != nil {
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: err,
				From:  from,
				To:    to,
				Seq:   seq,
			}))
		}
	default:
		panic(fmt.Errorf("unsupported actor ref type: %s", reflect.TypeOf(receiver).String()))
	}

	if opts.reply && opts.replySeq == 0 {
		if opts.ReplyTimeout == 0 {
			opts.ReplyTimeout = time.Second
		}

		waitChan := make(chan Message)
		ctx, cancel := context.WithTimeout(s.ctx, opts.ReplyTimeout)

		defer func(s *ActorSystem, seq uint64, cancel context.CancelFunc, wait chan Message) {
			cancel()
			close(wait)
			s.askWaitsLock.Lock()
			delete(s.askWaits, seq)
			s.askWaitsLock.Unlock()
		}(s, seq, cancel, waitChan)

		s.askWaitsLock.Lock()
		s.askWaits[seq] = waitChan
		s.askWaitsLock.Unlock()

		// wait for reply
		select {
		case <-ctx.Done():
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: ErrMessageReplyTimeout,
				From:  from,
				To:    to,
				Seq:   seq,
			}))
		case reply := <-waitChan:
			return reply
		}
	}

	return nil
}

func (s *ActorSystem) onProcessServerMessage() {
	for bytes := range s.server.C() {
		var remoteCtx = new(_RemoteMessageContext)
		if err := s.codec.Decode(bytes, remoteCtx); err != nil {
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: err,
			}))
			continue
		}

		s.actorRW.RLock()
		core, exist := s.actors[remoteCtx.ReceiverId]
		s.actorRW.RUnlock()
		if !exist {
			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: fmt.Errorf("%w: %s", ErrActorDeadOrNotExist, remoteCtx.ReceiverId),
				From:  remoteCtx.SenderId,
				To:    remoteCtx.ReceiverId,
			}))
			continue

		}

		if remoteCtx.ReplySeq > 0 {
			s.askWaitsLock.RLock()
			wait, existWait := s.askWaits[remoteCtx.ReplySeq]
			s.askWaitsLock.RUnlock()
			if existWait {
				wait <- remoteCtx.Message
				continue
			}

			s.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
				Error: fmt.Errorf("%w: %d", ErrAskWaitNotExist, remoteCtx.ReplySeq),
				From:  remoteCtx.SenderId,
				To:    remoteCtx.ReceiverId,
			}))
			continue
		}

		localCtx := &_LocalMessageContext{
			ActorContext: core,
			sender: &_RemoteActorRef{
				system:  s,
				actorId: remoteCtx.SenderId,
			},
			message: remoteCtx.Message,
			network: remoteCtx.Network,
			host:    remoteCtx.Host,
			port:    remoteCtx.Port,
		}

		core.group.Add(1)
		core.dispatcher.Send(core, localCtx)
	}
}
