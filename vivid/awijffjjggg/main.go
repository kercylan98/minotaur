package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/client"
	"github.com/kercylan98/minotaur/rpc/codec"
	"github.com/kercylan98/minotaur/rpc/transporter"
	"github.com/kercylan98/minotaur/vivid"
	"time"
)

type AccountActor struct {
	vivid.BasicActor
}

func (a *AccountActor) OnSpawn(system *vivid.ActorSystem, terminated vivid.ActorTerminatedNotifier) error {
	return a.RegisterTell("onLogin", func(message vivid.Context) error {
		var pwd string
		message.MustReadTo(&pwd)
		fmt.Println("AccountActor.OnReceive", pwd)
		return nil
	})
}

type Discovery struct {
}

func (d *Discovery) GetInstance(name string) (rpc.Client, error) {
	return client.NewGoRPC("tcp", "127.0.0.1:9999", codec.NewGob())
}

func main() {
	srv := rpc.NewServer(transporter.NewGoRPC(), rpc.NewRouter(), codec.NewGob())
	system := vivid.NewActorSystem("127.0.0.1", 9999, "Account", vivid.WithDiscovery(
		srv, new(Discovery),
	))

	go func() {
		if err := srv.ListenAndServe("tcp", "127.0.0.1:9999"); err != nil {
			panic(err)
		}
	}()

	actorId, err := system.Spawn(new(AccountActor))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 1)

	if err := system.Tell(actorId, actorId, "onLogin", "123456"); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 100000)
	system.Destroy()
}
