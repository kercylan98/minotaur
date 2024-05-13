package vivid_test

import (
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/client"
	"github.com/kercylan98/minotaur/rpc/codec"
	"github.com/kercylan98/minotaur/rpc/transporter"
	"github.com/kercylan98/minotaur/vivid"
	"sync"
	"testing"
)

type TestActor struct {
	vivid.BasicActor
}

type Discovery struct {
}

func (d *Discovery) GetInstance(name string) (rpc.Client, error) {
	return client.NewGoRPC("tcp", "127.0.0.1:9999", codec.NewGob())
}

func (a *TestActor) OnSpawn(system *vivid.ActorSystem, terminated vivid.ActorTerminatedNotifier) error {
	return a.RegisterTell("test", func(message vivid.Context) error {
		return nil
	})
}

var once sync.Once
var system *vivid.ActorSystem
var actorId vivid.ActorId

func BenchmarkActorSystem_Tell(b *testing.B) {

	once.Do(func() {

		srv := rpc.NewServer(transporter.NewGoRPC(), rpc.NewRouter(), codec.NewGob())
		system = vivid.NewActorSystem("127.0.0.1", 9999, "Test", vivid.WithDiscovery(
			srv, new(Discovery),
		))

		var err error
		actorId, err = system.Spawn(new(TestActor))
		if err != nil {
			panic(err)
		}

		go func() {
			if err := srv.ListenAndServe("tcp", "127.0.0.1:9999"); err != nil {
				panic(err)
			}
		}()

	})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := system.Tell(actorId, actorId, "test", nil); err != nil {
				panic(err)
			}
		}
	})

}
