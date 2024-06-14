package pools_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/pools"
)

func ExampleNewObjectPool() {
	var p = pools.NewObjectPool[map[int]int](func() *map[int]int {
		return &map[int]int{}
	}, func(data *map[int]int) {
		collection.ClearMap(*data)
	})

	m := *p.Get()
	m[1] = 1
	p.Put(&m)
	fmt.Println(m)
	// Output:
	// map[]
}
