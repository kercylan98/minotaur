package buffer_test

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"testing"
)

func BenchmarkRingUnbounded_Write(b *testing.B) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
}

func BenchmarkRingUnbounded_Read(b *testing.B) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-ring.Read()
	}
}

func BenchmarkRingUnbounded_Write_Parallel(b *testing.B) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ring.Write(1)
		}
	})
}

func BenchmarkRingUnbounded_Read_Parallel(b *testing.B) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			<-ring.Read()
		}
	})
}
