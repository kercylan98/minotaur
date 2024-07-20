package benchmark

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
)

func BenchmarkActorSystemActorOf(b *testing.B) {
	system := NewBenchmarkActorSystem()
	provider := vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.ActorOf(provider)
	}
	b.StopTimer()
	system.Shutdown(true)
}
