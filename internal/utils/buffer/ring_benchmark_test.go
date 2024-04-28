package buffer_test

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"testing"
)

func BenchmarkRing_Write(b *testing.B) {
	ring := buffer.NewRing[int](1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
}

func BenchmarkRing_Read(b *testing.B) {
	ring := buffer.NewRing[int](1024)
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ring.Read()
	}
}
