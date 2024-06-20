package vivid_test

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"sync"
	"testing"
	"time"
)

type TestPriorityEventSubscriberActor struct {
	Name     string
	Priority int64
	Output   *[]string
	wg       *sync.WaitGroup
}

func (t *TestPriorityEventSubscriberActor) OnReceive(ctx vivid.MessageContext) {
	switch ctx.GetMessage().(type) {
	case vivid.OnBoot:
		t.wg.Add(1)
		ctx.Subscribe("", vivid.WithSubscribePriority(t.Priority, time.Second))
	case string:
		*t.Output = append(*t.Output, t.Name)
		ctx.Reply(1)
		t.wg.Done()
	}
}

func TestPriorityEventActor_OnReceive(t *testing.T) {
	var result = make([]string, 0, 3)
	var wg = new(sync.WaitGroup)
	defer func() {
		wg.Wait()
		vivid.TestActorSystem.Shutdown()
		for _, s := range result {
			t.Log(s)
		}

		if len(result) != 3 || result[0] != "sub3" || result[1] != "sub2" || result[2] != "sub1" {
			t.Fail()
		}
	}()

	vivid.ActorOfF(&vivid.TestActorSystem, func(options *vivid.ActorOptions[*TestPriorityEventSubscriberActor]) {
		options.WithInit(func(actor *TestPriorityEventSubscriberActor) {
			actor.Priority = 3
			actor.Output = &result
			actor.Name = "sub1"
			actor.wg = wg
		})
	})
	vivid.ActorOfF(&vivid.TestActorSystem, func(options *vivid.ActorOptions[*TestPriorityEventSubscriberActor]) {
		options.WithInit(func(actor *TestPriorityEventSubscriberActor) {
			actor.Priority = 2
			actor.Output = &result
			actor.Name = "sub2"
			actor.wg = wg
		})
	})
	vivid.ActorOfF(&vivid.TestActorSystem, func(options *vivid.ActorOptions[*TestPriorityEventSubscriberActor]) {
		options.WithInit(func(actor *TestPriorityEventSubscriberActor) {
			actor.Priority = 1
			actor.Output = &result
			actor.Name = "sub3"
			actor.wg = wg
		})
	})

	vivid.TestActorSystem.Publish(vivid.TestActorSystem.GetContext(), "")
}
