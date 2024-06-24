package vivid_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"testing"
)

type SupervisorActor struct {
}

func (s *SupervisorActor) OnAccident(system *vivid.ActorSystem, accident vivid.Accident) {
	accident.Responsible().Restart(accident.AccidentActor())
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
		ctx.System().Done()
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
	system.Add(2)
	system.ActorOf(func() vivid.Actor {
		return new(SupervisorActor)
	})

	system.Wait()
}
