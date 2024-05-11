package vivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

func BenchmarkNewActorId(b *testing.B) {
	for i := range uint64(b.N) {
		vivid.NewActorId("127.0.0.1", 8080, "Connection", i)
	}
}
