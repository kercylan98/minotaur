package slice_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
)

func TestDrop(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	t.Log(s, slice.Drop(1, 3, s))
}

func ExampleDrop() {
	fmt.Println(slice.Drop(1, 3, []int{1, 2, 3, 4, 5}))
	// Output:
	// [1 5]
}

func BenchmarkDrop(b *testing.B) {
	s := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.Drop(1, 3, s)
	}
}
