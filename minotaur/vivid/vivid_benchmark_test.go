package vivid_test

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

func BenchmarkActorOf(b *testing.B) {
	sys := vivid.GetBenchmarkTestSystem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vivid.ActorOf[*vivid.IneffectiveActor](sys)
	}
	b.StopTimer()
	sys.Shutdown()
}

func BenchmarkActorOfParallel(b *testing.B) {
	sys := vivid.GetBenchmarkTestSystem()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			vivid.ActorOf[*vivid.IneffectiveActor](sys)
		}
	})
	b.StopTimer()
	sys.Shutdown()
}
