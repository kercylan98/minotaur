package main

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/vivid"
	"time"
)

type AccountActor struct {
	vivid.BasicActor
}

func (a *AccountActor) OnSpawn(system *vivid.ActorSystem, terminated vivid.ActorTerminatedNotifier) error {
	return a.RegisterTell(a.onLogin)
}

func (a *AccountActor) onLogin(account string, password string) {
	log.Info("AccountActor.onLogin", account, password)
}

func main() {
	system := vivid.NewActorSystem("127.0.0.1", 9999, "Account")
	defer system.Destroy()

	actorId, err := system.Spawn(new(AccountActor))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 1)

	if err := system.Tell(actorId, actorId, new(AccountActor).onLogin, "kercylan", "123456"); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 100000)
}
