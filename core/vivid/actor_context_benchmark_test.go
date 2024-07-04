package vivid_test

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit"
	"testing"
)

func BenchmarkActorContext_Tell(b *testing.B) {
	toolkit.EnableHttpPProf(":9999", "/debug/pprof")
	system := vivid.NewActorSystem()
	ref := system.ActorOf(func() vivid.Actor {
		return &vivid.WasteActor{}
	})
	b.Log("start", b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Context().Tell(ref, i)
	}
	b.StopTimer()

	system.Shutdown()
	b.Log("end", b.N)
}

func BenchmarkActorContext_Ask(b *testing.B) {
	system := vivid.NewActorSystem()
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
	system := vivid.NewActorSystem()
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
