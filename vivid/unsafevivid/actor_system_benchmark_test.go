package unsafevivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

type TestActor struct {
	*vivid.BasicActor
}

func BenchmarkActorSystem_ActorOf(b *testing.B) {
	var err error
	system := vivid.NewActorSystem("test")
	if err = system.Run(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err = system.ActorOf(new(TestActor)); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()

	if err = system.Shutdown(); err != nil {
		b.Fatal(err)
	}
}
