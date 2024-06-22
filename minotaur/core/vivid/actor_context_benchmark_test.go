package vivid_test

import (
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"testing"
)

func BenchmarkActorContext_Tell(b *testing.B) {
	system := vivid.NewActorSystem("benchmark")
	ref := system.ActorOf(func(options *vivid.ActorOptions) vivid.Actor {
		return &vivid.WasteActor{}
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Context().Tell(ref, 1)
	}
	b.StopTimer()

	system.Shutdown()
}

func BenchmarkActorContext_Ask(b *testing.B) {
	system := vivid.NewActorSystem("benchmark")
	ref := system.ActorOf(func(options *vivid.ActorOptions) vivid.Actor {
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
	system := vivid.NewActorSystem("benchmark")
	ref := system.ActorOf(func(options *vivid.ActorOptions) vivid.Actor {
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
