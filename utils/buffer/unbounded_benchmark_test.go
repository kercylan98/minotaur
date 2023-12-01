package buffer_test

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"testing"
)

func BenchmarkUnboundedBuffer(b *testing.B) {
	ub := buffer.NewUnboundedN[int]()

	b.Run("Put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.Put(i)
		}
	})

	b.Run("Load", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.Load()
		}
	})

	// 先填充数据以防止 Get 被阻塞
	for i := 0; i < b.N; i++ {
		ub.Put(i)
	}

	b.Run("Get", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ub.Put(i)
			<-ub.Get()
			ub.Load()
		}
	})
}
