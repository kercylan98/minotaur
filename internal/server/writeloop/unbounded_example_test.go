package writeloop_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server/writeloop"
	"github.com/kercylan98/minotaur/utils/hub"
	"sync"
)

func ExampleNewUnbounded() {
	pool := hub.NewObjectPool[Message](func() *Message {
		return &Message{}
	}, func(data *Message) {
		data.ID = 0
	})
	var wait sync.WaitGroup
	wait.Add(10)
	wl := writeloop.NewUnbounded(pool, func(message *Message) error {
		fmt.Println(message.ID)
		wait.Done()
		return nil
	}, func(err any) {
		fmt.Println(err)
	})

	for i := 0; i < 10; i++ {
		m := pool.Get()
		m.ID = i
		wl.Put(m)
	}

	wait.Wait()
	wl.Close()
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}
