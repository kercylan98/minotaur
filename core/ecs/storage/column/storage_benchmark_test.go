package column_test

import (
	"github.com/kercylan98/minotaur/core/ecs/storage/column"
	"testing"
)

func CreateStorage(columnNum int) *column.Storage[int, int] {
	var storage = column.New[int, int]()
	for i := 0; i < columnNum; i++ {
		storage.SetColumn(i, func() any {
			return 0
		})
	}
	return storage
}

func BenchmarkStorage_Get(b *testing.B) {
	var storage = CreateStorage(5)
	for i := 0; i < 1000; i++ {
		storage.AddRow(i)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			storage.Get(500, 0)
		}
	})
}

func BenchmarkStorage_GetRow(b *testing.B) {
	var storage = CreateStorage(5)
	for i := 0; i < 1000; i++ {
		storage.AddRow(i)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			storage.GetRow(500)
		}
	})
}

func BenchmarkStorage_Add(b *testing.B) {
	var storage = CreateStorage(5)
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		storage.AddRow(j)
	}
}

func BenchmarkStorage_AddBatch(b *testing.B) {
	var storage = CreateStorage(5)
	var batches = make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		batches = append(batches, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		storage.AddRows(batches)
	}

}
