package hub_test

import (
	"github.com/kercylan98/minotaur/utils/hub"
	"testing"
)

func BenchmarkPool_Get2Put(b *testing.B) {
	var pool = hub.NewObjectPool[map[string]int](func() *map[string]int {
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
