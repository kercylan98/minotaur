package vivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"sync"
	"testing"
)

func BenchmarkActorSystem_ActorOf(b *testing.B) {
	sys := vivid.NewActorSystem("test")
	if err := sys.Run(); err != nil {
		b.Fatal(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sys.ActorOf(&vivid.BasicActor{})
	}
	b.StopTimer()

	if err := sys.Shutdown(); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkActorSystem_ActorOf_Parallel(b *testing.B) {
	sys := vivid.NewActorSystem("test")
	if err := sys.Run(); err != nil {
		b.Fatal(err)
	}
	var wait sync.WaitGroup

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		wait.Add(1)
		defer wait.Done()
		for pb.Next() {
			_, _ = sys.ActorOf(&vivid.BasicActor{})
		}
	})
	wait.Wait()
	b.StopTimer()

	if err := sys.Shutdown(); err != nil {
		b.Fatal(err)
	}
}
