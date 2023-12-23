package buffer_test

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"testing"
)

func BenchmarkUnboundedBuffer(b *testing.B) {
	ub := buffer.NewUnbounded[int]()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ub.Put(1)
			<-ub.Get()
			ub.Load()
		}
	})
}
