package stream_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/stream"
)

func ExampleWithMap() {
	m := stream.WithMap(map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "d"}).Filter(func(key int, value string) bool {
		return key > 3
	})
	fmt.Println(len(m))

	// Output:
	// 2
}

func ExampleWithMapCopy() {
	m := stream.WithMapCopy(map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "d"}).Filter(func(key int, value string) bool {
		return key > 3
	})
	fmt.Println(len(m))

	// Output:
	// 2
}
