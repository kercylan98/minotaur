package dispatcher_test

import (
	"github.com/kercylan98/minotaur/server/internal/dispatcher"
	"sync"
	"testing"
	"time"
)

type TestMessage struct {
	producer string
	v        int
}

func (m *TestMessage) GetProducer() string {
	return m.producer
}

func TestDispatcher_PutStartClose(t *testing.T) {
	// 写入完成后，关闭分发器再开始分发，确保消息不会丢失
	w := new(sync.WaitGroup)
	cw := new(sync.WaitGroup)
	cw.Add(1)
	d := dispatcher.NewDispatcher[string, *TestMessage](1024*16, "test", func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
		t.Log(message)
		w.Done()
	}).SetClosedHandler(func(dispatcher *dispatcher.Dispatcher[string, *TestMessage]) {
		t.Log("closed")
		cw.Done()
	})

	for i := 0; i < 100; i++ {
		w.Add(1)
		d.Put(&TestMessage{
			producer: "test",
			v:        i,
		})
	}

	d.Start()
	d.Expel()
	d.UnExpel()
	w.Wait()
	time.Sleep(time.Second)
	d.Expel()
	cw.Wait()
	t.Log("done")
}
