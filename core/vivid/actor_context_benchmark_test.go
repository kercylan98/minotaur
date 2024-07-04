package vivid_test

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"testing"
)

func BenchmarkActorContext_Tell(b *testing.B) {
	system := vivid.NewBenchmarkActorSystem()
	ref := system.ActorOf(func() vivid.Actor {
		return &vivid.WasteActor{}
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Context().Tell(ref, i)
	}
	b.StopTimer()

	system.Shutdown()
}

func BenchmarkActorContext_Ask(b *testing.B) {
	system := vivid.NewBenchmarkActorSystem()
	ref := system.ActorOf(func() vivid.Actor {
		return &vivid.StringEchoActor{}
	})
	var m = "hi"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Context().Ask(ref, m)
	}
	b.StopTimer()

	system.Shutdown()
}

func BenchmarkActorContext_FutureAsk(b *testing.B) {
	system := vivid.NewBenchmarkActorSystem()
	ref := system.ActorOf(func() vivid.Actor {
		return &vivid.StringEchoActor{}
	})
	var m = "hi"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := system.Context().FutureAsk(ref, m).Wait(); err != nil {
			panic(err)
		}
	}
	b.StopTimer()
	system.Shutdown()
}
