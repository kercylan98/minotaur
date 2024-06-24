package vivid_test

import (
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"testing"
)

func BenchmarkActorContext_RootActorOf(b *testing.B) {
	system := vivid.NewActorSystem("benchmark")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		system.ActorOf(func() vivid.Actor {
			return &vivid.WasteActor{}
		})
	}

	b.StopTimer()
	system.Shutdown()
}