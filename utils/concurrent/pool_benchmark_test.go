package concurrent_test

import (
	"github.com/kercylan98/minotaur/utils/concurrent"
	"testing"
)

func BenchmarkPool_Get2Put(b *testing.B) {
	var pool = concurrent.NewPool[map[string]int](func() *map[string]int {
		return &map[string]int{}
	}, func(data *map[string]int) {
		for k := range *data {
			delete(*data, k)
		}
	})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			msg := pool.Get()
			pool.Release(msg)
		}
	})
	b.StopTimer()
	b.ReportAllocs()
}
