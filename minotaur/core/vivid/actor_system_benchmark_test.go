package vivid_test

import (
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"os"
	"runtime/pprof"
	"testing"
)

func BenchmarkActorContext_ActorOf(b *testing.B) {
	os.Remove("cpu.pprof")
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	system := vivid.NewActorSystem("benchmark")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		system.ActorOf(func(options *vivid.ActorOptions) vivid.Actor {
			return &vivid.WasteActor{}
		})
	}

	b.StopTimer()
	system.Shutdown()
}
