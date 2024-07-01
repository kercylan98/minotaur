package vivid_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
	"testing"
	"time"
)

type StateActor struct {
	counter int
}

func (s *StateActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case int:
		ctx.StatusChanged(m)
		s.counter++
	case vivid.OnPersistenceSnapshot:
		// 持久化快照
		ctx.PersistSnapshot(s)
	case *StateActor:
		// 恢复状态
		s.counter = m.counter
	case string:
		fmt.Println(s.counter)
	}
}

func TestActorContext_Persist(t *testing.T) {
	system := vivid.NewActorSystem()

	// 创建 Actor
	ref := system.ActorOf(func() vivid.Actor {
		return &StateActor{}
	}, func(options *vivid.ActorOptions) {
		options.WithPersistenceEventLimit(10).WithPersistence(vivid.NewMemoryStorage(), "state")
	})

	// 发送消息
	for i := 0; i < 11; i++ {
		system.Context().Tell(ref, i)
	}

	time.Sleep(time.Second)

	// 创建 Actor
	ref2 := system.ActorOf(func() vivid.Actor {
		return &StateActor{}
	}, func(options *vivid.ActorOptions) {
		options.WithPersistenceEventLimit(10).WithPersistence(vivid.NewMemoryStorage(), "state")
	})
	system.Context().Tell(ref2, "print")

	time.Sleep(time.Second)
}
