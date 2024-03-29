package hub_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/hub"
)

func ExampleNewObjectPool() {
	var p = hub.NewObjectPool[map[int]int](func() *map[int]int {
		return &map[int]int{}
	}, func(data *map[int]int) {
		collection.ClearMap(*data)
	})

	m := *p.Get()
	m[1] = 1
	p.Release(&m)
	fmt.Println(m)
	// Output:
	// map[]
}
