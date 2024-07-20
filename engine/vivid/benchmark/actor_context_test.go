package benchmark

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
)

func BenchmarkActorContextTell(b *testing.B) {
	system := NewBenchmarkActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
	}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Tell(ref, i)
	}
	b.StopTimer()
	system.Shutdown(true)
}

func BenchmarkActorContextAsk(b *testing.B) {
	system := NewBenchmarkActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
	}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Ask(ref, i)
	}
	b.StopTimer()
	system.Shutdown(true)
}

func BenchmarkActorContextFutureAsk(b *testing.B) {
	system := NewBenchmarkActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch v := ctx.Message().(type) {
			case int:
				ctx.Reply(v)
			}
		})
	}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.FutureAsk(ref, i).AssertWait()
	}
	b.StopTimer()
	system.Shutdown(true)
}

func BenchmarkActorContextBroadcast_10(b *testing.B) {
	system := NewBenchmarkActorSystem()
	for i := 0; i < 10; i++ {
		system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
		}))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Broadcast(i)
	}
	b.StopTimer()
	system.Shutdown(true)
}

func BenchmarkActorContextBroadcast_100(b *testing.B) {
	system := NewBenchmarkActorSystem()
	for i := 0; i < 100; i++ {
		system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
		}))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Broadcast(i)
	}
	b.StopTimer()
	system.Shutdown(true)
}

func BenchmarkActorContextBroadcast_1000(b *testing.B) {
	system := NewBenchmarkActorSystem()
	for i := 0; i < 1000; i++ {
		system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
		}))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Broadcast(i)
	}
	b.StopTimer()
	system.Shutdown(true)
}
