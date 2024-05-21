package vivid

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"path"
	"reflect"
	"sync"
)

func ActorOf[T Actor](actorOf actorOf, options ...*ActorOptions[T]) ActorRef {
	var opts = parseActorOptions(options...)
	var tof = reflect.TypeOf((*T)(nil)).Elem()
	var ins = reflect.New(tof).Elem().Interface().(T)
	var system = actorOf.getSystem()

	if opts.Parent == nil {
		opts.Parent = system.userGuard
	}

	ctx, err := generateActor[T](actorOf.getSystem(), ins, opts)
	if err != nil {
		system.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeActorOf, DeadLetterEventActorOf{
			Error:  err,
			Parent: opts.Parent,
			Name:   opts.Name,
		}))
	}
	if ctx == nil {
		return &_DeadLetterActorRef{
			system: system,
		}
	}
	return ctx
}

func generateActor[T Actor](system *ActorSystem, actor T, options *ActorOptions[T]) (*_ActorCore, error) {
	if options.Name == charproc.None {
		options.Name = uuid.NewString()
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

	// 启动 Actor
	system.waitGroup.Add(1)
	core.dispatcher.Attach(core)
	core.Tell(OnPreStart{})

	return core, nil
}
