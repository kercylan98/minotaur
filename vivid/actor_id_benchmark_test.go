package vivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

func BenchmarkNewActorId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")
	}
}

func BenchmarkActorIdInfo(b *testing.B) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	b.Run("Network", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actorId.Network()
		}
	})

	b.Run("Cluster", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actorId.Cluster()
		}
	})

	b.Run("Host", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actorId.Host()
		}
	})

	b.Run("Port", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actorId.Port()
		}
	})

	b.Run("System", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actorId.System()
		}
	})

	b.Run("Name", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actorId.Name()
		}
	})
}
