package collection

import "testing"

func BenchmarkTopological(b *testing.B) {
	type Item struct {
		ID      int
		Depends []int
	}

	var items = []Item{
		{ID: 2, Depends: []int{4}},
		{ID: 1, Depends: []int{2, 3}},
		{ID: 3, Depends: []int{4}},
		{ID: 4, Depends: []int{5}},
		{ID: 5, Depends: []int{}},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := TopologicalSort(items, func(item Item) int {
			return item.ID
		}, func(item Item) []int {
			return item.Depends
		})
		if err != nil {
			b.Error(err)
			return
		}
	}
}
