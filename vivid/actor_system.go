package vivid

import (
	"fmt"
	"sync"
)

// ActorTerminatedNotifier Actor 终止通知器，当 Actor 终止时务必调用该函数，以便 ActorSystem 可以正确清理资源
//   - 该函数多次调用不会产生副作用，但仅第一次调用有效
type ActorTerminatedNotifier func()

// NewActorSystem 创建一个 ActorSystem
func NewActorSystem(host string, port uint16, name string) *ActorSystem {
	s := &ActorSystem{
		host:   host,
		port:   port,
		name:   name,
		actors: make(map[ActorId]*actorCore),
	}
	return s
}

// ActorSystem 是维护 Actor 的容器，负责 Actor 的创建、销毁、消息分发等
//   - 通常推荐每个应用程序仅使用一个 ActorSystem
type ActorSystem struct {
	host    string
	port    uint16
	name    string
	actorRW sync.RWMutex
	guid    ActorGuid
	actors  map[ActorId]*actorCore
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
func (s *ActorSystem) Tell(sender ActorId, receiver ActorId, command any, params ...any) error {
	s.actorRW.RLock()
	actor, exist := s.actors[receiver]
	s.actorRW.RUnlock()
	if exist {
		actor.add(&Message{
			Sender:   sender,
			Receiver: receiver,
			Command:  command,
			Params:   params,
		})
		return nil
	}

	// TODO: RPC
	return fmt.Errorf("%w: %s", ErrActorNotFound, receiver)
}

// Ask 发送消息并等待回复
func (s *ActorSystem) Ask(sender ActorId, receiver ActorId, message Message) (*Message, error) {
	return nil, nil
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
