package slice_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/slice"
	"testing"
)

func TestChunk(t *testing.T) {
	var collection = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var chunks = slice.Chunk(collection, 3)
	for _, chunk := range chunks {
		t.Log(chunk)
	}
}

func ExampleChunk() {
	var collection = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var chunks = slice.Chunk(collection, 3)
	for _, chunk := range chunks {
		fmt.Println(chunk)
	}
	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
}

func BenchmarkChunk(b *testing.B) {
	var collection = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.Chunk(collection, 3)
	}
}
