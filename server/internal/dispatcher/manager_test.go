package dispatcher_test

import (
	"github.com/kercylan98/minotaur/server/internal/dispatcher"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	var mgr *dispatcher.Manager[string, *TestMessage]
	var onHandler = func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
		t.Log(dispatcher.Name(), message, mgr.GetDispatcherNum())
		switch message.v {
		case 4:
			mgr.UnBindProducer("test")
			t.Log("UnBindProducer")
		case 6:
			mgr.BindProducer(message.GetProducer(), "test-dispatcher")
			t.Log("BindProducer")
		case 9:
			dispatcher.Put(&TestMessage{
				producer: "test",
				v:        10,
			})
		case 10:
			mgr.UnBindProducer("test")
			t.Log("UnBindProducer", mgr.GetDispatcherNum())
		}

	}
	mgr = dispatcher.NewManager[string, *TestMessage](1024*16, onHandler)

	mgr.BindProducer("test", "test-dispatcher")
	for i := 0; i < 10; i++ {
		d := mgr.GetDispatcher("test").SetClosedHandler(func(dispatcher *dispatcher.Dispatcher[string, *TestMessage]) {
			t.Log("closed")
		})
		d.Put(&TestMessage{
			producer: "test",
			v:        i,
		})
	}

	time.Sleep(times.Day)
}
