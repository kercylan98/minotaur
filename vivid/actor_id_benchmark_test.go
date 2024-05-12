package vivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

func BenchmarkActorId(b *testing.B) {
	id := vivid.NewActorId("127.0.0.1", 8080, "Connection", 1)
	b.Run("Generate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vivid.NewActorId("127.0.0.1", 8080, "Connection", 1)
		}
	})

	b.Run("GetHost", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.Host()
		}
	})

	b.Run("GetPort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.Port()
		}
	})

	b.Run("GetSystemName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.SystemName()
		}
	})

	b.Run("GetGuid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.Guid()
		}
	})
}

func BenchmarkNewActorId(b *testing.B) {
	for i := range uint64(b.N) {
		vivid.NewActorId("127.0.0.1", 8080, "Connection", i)
	}
}

func BenchmarkActorId_Get(b *testing.B) {
	id := vivid.NewActorId("127.0.0.1", 8080, "Connection", 1)
	b.Run("GetHost", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.Host()
		}
	})

	b.Run("GetPort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.Port()
		}
	})

	b.Run("GetSystemName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.SystemName()
		}
	})

	b.Run("GetGuid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id.Guid()
		}
	})
}
