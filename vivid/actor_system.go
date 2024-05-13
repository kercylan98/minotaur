package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/rpc"
	"sync"
)

// ActorTerminatedNotifier Actor 终止通知器，当 Actor 终止时务必调用该函数，以便 ActorSystem 可以正确清理资源
//   - 该函数多次调用不会产生副作用，但仅第一次调用有效
type ActorTerminatedNotifier func()

// NewActorSystem 创建一个 ActorSystem
func NewActorSystem(host string, port uint16, name string, opts ...ActorSystemOption) *ActorSystem {
	s := &ActorSystem{
		opt:    defaultActorSystemOptions().apply(opts...),
		host:   host,
		port:   port,
		name:   name,
		actors: make(map[ActorId]*actorCore),
	}

	if s.opt.rpcSrv != nil {
		s.opt.rpcSrv.GetRouter().Register("/actor/context", s.onActorMessage)
	}
	return s
}

// ActorSystem 是维护 Actor 的容器，负责 Actor 的创建、销毁、消息分发等
//   - 通常推荐每个应用程序仅使用一个 ActorSystem
type ActorSystem struct {
	opt     *ActorSystemOptions
	host    string
	port    uint16
	name    string
	actorRW sync.RWMutex
	guid    ActorGuid
	actors  map[ActorId]*actorCore
}

func (s *ActorSystem) onActorMessage(ctx rpc.Context) error {
	var msg context
	if err := ctx.ReadTo(&msg); err != nil {
		return err
	}

	s.actorRW.RLock()
	actor, exist := s.actors[msg.Receiver]
	s.actorRW.RUnlock()
	if !exist {
		return fmt.Errorf("%w: %s", ErrActorNotFound, msg.Receiver)
	}

	msg.reader = func(dst any) error {
		return s.opt.codec.Decode(msg.Data, dst)
	}
	msg.done = make(chan struct{})
	actor.add(&msg)
	<-msg.done
	return msg.err
}

// Spawn 创建一个 Actor
func (s *ActorSystem) Spawn(actor Actor) (ActorId, error) {
	s.actorRW.Lock()

	s.guid++
	id := NewActorId(s.host, s.port, s.name, s.guid)
	core := newActorCore(id, actor)
	if err := actor.OnSpawn(s, func() {
		s.onDestroy(id)
	}); err != nil {
		s.guid--
		s.actorRW.Unlock()
		return "", err
	}
	go core.start()
	s.actors[id] = core

	s.actorRW.Unlock()
	return id, nil
}

// Tell 发送消息
func (s *ActorSystem) Tell(sender ActorId, receiver ActorId, command string, data any) error {
	ctx := &context{
		Sender:   sender,
		Receiver: receiver,
		Command:  command,
	}
	if data != nil {
		d, err := s.opt.codec.Encode(data)
		if err != nil {
			return err
		}
		ctx.Data = d
	}

	s.actorRW.RLock()
	actor, exist := s.actors[receiver]
	s.actorRW.RUnlock()
	if exist {
		ctx.reader = func(dst any) error {
			return s.opt.codec.Decode(ctx.Data, dst)
		}
		ctx.done = make(chan struct{})
		actor.add(ctx)
		<-ctx.done
		return ctx.err
	}

	// 通过服务发现发送消息
	if s.opt.discovery != nil {
		cli, err := s.opt.discovery.GetInstance(s.name)
		if err != nil {
			return err
		}
		return cli.Tell("/actor/context", ctx)
	}

	return fmt.Errorf("%w: %s", ErrActorNotFound, receiver)
}

// Ask 发送消息并等待回复
func (s *ActorSystem) Ask(sender ActorId, receiver ActorId, message context) (Result, error) {
	s.actorRW.RLock()
	actor, exist := s.actors[receiver]
	s.actorRW.RUnlock()
	if exist {
		message.Sender = sender
		message.Receiver = receiver
		message.reader = func(dst any) error {
			return s.opt.codec.Decode(message.Data, dst)
		}
		message.done = make(chan struct{})
		actor.add(&message)
		<-message.done
		if message.err != nil {
			return nil, message.err
		}

		return func(dst any) error {
			result, err := s.opt.codec.Encode(message.reply)
			if err != nil {
				return err
			}
			return s.opt.codec.Decode(result, dst)
		}, nil
	}

	// 通过服务发现发送消息
	if s.opt.discovery != nil {
		cli, err := s.opt.discovery.GetInstance(s.name)
		if err != nil {
			return nil, err
		}
		resp, err := cli.Ask("/actor/context", message)
		if err != nil {
			return nil, err
		}
		return func(dst any) error {
			return resp.ReadTo(dst)
		}, nil
	}

	return nil, fmt.Errorf("%w: %s", ErrActorNotFound, receiver)
}

// onDestroy
func (s *ActorSystem) onDestroy(id ActorId) {
	s.actorRW.Lock()
	actor, exist := s.actors[id]
	if exist {
		actor.stop()
		delete(s.actors, id)
	}
	s.actorRW.Unlock()
}

// Destroy 销毁整个 ActorSystem 下当前所有的 Actor，重置 ActorSystem 状态
func (s *ActorSystem) Destroy() {
	defer s.actorRW.RUnlock() // 跳出循环必定是锁定状态，重置资源后确保解锁
	for {
		// Actor 会等待所有消息处理完毕后再退出，在退出期间也会接收新消息，需要多轮检查退出。故不适用 context.Context
		s.actorRW.RLock()
		if len(s.actors) == 0 {
			break
		}

		// 使用 WaitGroup 等待异步停止，避免此刻无法产生新的 Actor
		var wg = new(sync.WaitGroup)
		for _, a := range s.actors {
			wg.Add(1)
			go func(wg *sync.WaitGroup, s *ActorSystem, a *actorCore) {
				defer wg.Done()
				a.getActor().OnDestroy()
				a.stop()
				s.actorRW.Lock()
				delete(s.actors, a.getId())
				s.actorRW.Unlock()
			}(wg, s, a)
		}

		s.actorRW.RUnlock()
		wg.Wait()
	}

	// 重置状态
	s.guid = 0
	s.actors = make(map[ActorId]*actorCore)
}
