package buffer_test

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"testing"
)

func BenchmarkUnbounded_Write(b *testing.B) {
	u := buffer.NewUnbounded[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Put(i)
	}
}

func BenchmarkUnbounded_Read(b *testing.B) {
	u := buffer.NewUnbounded[int]()
	for i := 0; i < b.N; i++ {
		u.Put(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-u.Get()
		u.Load()
	}
}

func BenchmarkUnbounded_Write_Parallel(b *testing.B) {
	u := buffer.NewUnbounded[int]()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			u.Put(1)
		}
	})
}

func BenchmarkUnbounded_Read_Parallel(b *testing.B) {
	u := buffer.NewUnbounded[int]()
	for i := 0; i < b.N; i++ {
		u.Put(i)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			<-u.Get()
			u.Load()
		}
	})
}
