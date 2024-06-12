package test

import (
	"errors"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
	"time"
)

func TestRestart(t *testing.T) {
	initTest(t)
	system := vivid.NewActorSystem("test")
	ref := system.ActorOf(vivid.OfO[*Actor](func(actorOptions *vivid.ActorOptions[*Actor]) {
		actorOptions.WithName("root")
	}))

	ref.Tell("child")

	ref.Tell(errors.New("error"))

	time.Sleep(time.Second)
	system.Shutdown()
}
