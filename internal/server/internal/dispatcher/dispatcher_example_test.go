package dispatcher_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server/internal/dispatcher"
	"sync"
	"sync/atomic"
)

func ExampleNewDispatcher() {
	m := new(atomic.Int64)
	fm := new(atomic.Int64)
	w := new(sync.WaitGroup)
	w.Add(1)
	d := dispatcher.NewDispatcher(1024, "example-dispatcher", func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
		m.Add(1)
	})
	d.SetClosedHandler(func(dispatcher *dispatcher.Action[string, *TestMessage]) {
		w.Done()
	})
	var producers = []string{"producer1", "producer2", "producer3"}
	for i := 0; i < len(producers); i++ {
		p := producers[i]
		for i := 0; i < 10; i++ {
			d.Put(&TestMessage{producer: p})
		}
		d.SetProducerDoneHandler(p, func(p string, dispatcher *dispatcher.Action[string, *TestMessage]) {
			fm.Add(1)
		})
	}
	d.Start()
	d.Expel()
	w.Wait()
	fmt.Println(fmt.Sprintf("producer num: %d, producer done: %d, finished: %d", len(producers), fm.Load(), m.Load()))
	// Output:
	// producer num: 3, producer done: 3, finished: 30
}
