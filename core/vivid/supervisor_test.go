package vivid_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
	"testing"
	"time"
)

type SupervisorActor struct {
}

func (s *SupervisorActor) OnReceive(ctx vivid.ActorContext) {
	switch ctx.Message().(type) {
	case vivid.OnLaunch:
		fmt.Println("supervisor launched")
		ctx.ActorOf(func() vivid.Actor {
			return new(AccidentActor)
		})
	}
}

type AccidentActor struct {
}

func (a *AccidentActor) OnReceive(ctx vivid.ActorContext) {
	switch ctx.Message().(type) {
	case vivid.OnLaunch:
		fmt.Println("accident launched")
		panic("boom")
	case vivid.OnTerminate:
		fmt.Println("accident terminated")
	case vivid.OnRestarting:
		fmt.Println("accident restarting")
	case string:
		fmt.Println("accident boom")
	}
}

func TestSupervisor(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithName("test")
	})
	//system.Add(2)
	system.ActorOf(func() vivid.Actor {
		return new(SupervisorActor)
	})

	time.Sleep(time.Second * 3333)
	//system.Wait()
}
