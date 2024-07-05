package core_test

import (
	"github.com/kercylan98/minotaur/core"
	"testing"
)

func BenchmarkNewAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.NewAddress("tcp", "benchmark", "localhost", 8080, "/path")
	}
}

func BenchmarkAddress(b *testing.B) {
	b.Run("Network", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = address.Network()
		}
	})

	b.Run("System", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = address.System()
		}
	})

	b.Run("Host", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = address.Host()
		}
	})

	b.Run("Port", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = address.Port()
		}
	})

	b.Run("LogicPath", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = address.LogicPath()
		}
	})

	b.Run("PhysicalAddress", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = address.PhysicalAddress()
		}
	})

	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = address.String()
		}
	})
}
