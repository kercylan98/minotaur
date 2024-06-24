package vivid_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"testing"
	"time"
)

type SupervisorActor struct {
}

func (s *SupervisorActor) OnAccident(system *vivid.ActorSystem, supervisor vivid.Supervisor, accident vivid.ActorRef, reason, message vivid.Message) {
	supervisor.Restart(accident)
}

func (s *SupervisorActor) OnReceive(ctx vivid.ActorContext) {
	switch ctx.Message().(type) {
	case vivid.OnLaunch:
		fmt.Println("supervisor launched")
		accident := ctx.ActorOf(func() vivid.Actor {
			return new(AccidentActor)
		})
		ctx.Tell(accident, "boom")
	}
}

type AccidentActor struct {
}

func (a *AccidentActor) OnReceive(ctx vivid.ActorContext) {
	switch ctx.Message().(type) {
	case vivid.OnLaunch:
		fmt.Println("accident launched")
	case vivid.OnTerminate:
		fmt.Println("accident terminated")
	case vivid.OnRestarting:
		fmt.Println("accident restarting")
	case string:
		fmt.Println("accident boom")
		panic("boom")
	}
}

func TestSupervisor(t *testing.T) {
	system := vivid.NewActorSystem("test")
	system.ActorOf(func() vivid.Actor {
		return new(SupervisorActor)
	})

	time.Sleep(time.Second)
}
