package dispatcher_test

import (
	"github.com/kercylan98/minotaur/server/internal/dispatcher"
	"sync"
	"sync/atomic"
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

func TestNewDispatcher(t *testing.T) {
	var cases = []struct {
		name        string
		bufferSize  int
		handler     dispatcher.Handler[string, *TestMessage]
		shouldPanic bool
	}{
		{name: "TestNewDispatcher_BufferSize0AndHandlerNil", bufferSize: 0, handler: nil, shouldPanic: true},
		{name: "TestNewDispatcher_BufferSize0AndHandlerNotNil", bufferSize: 0, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {}, shouldPanic: true},
		{name: "TestNewDispatcher_BufferSize1AndHandlerNil", bufferSize: 1, handler: nil, shouldPanic: true},
		{name: "TestNewDispatcher_BufferSize1AndHandlerNotNil", bufferSize: 1, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {}, shouldPanic: false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !c.shouldPanic {
					t.Errorf("NewDispatcher() should not panic, but panic: %v", r)
				}
			}()
			dispatcher.NewDispatcher(c.bufferSize, c.name, c.handler)
		})
	}
}

func TestDispatcher_SetProducerDoneHandler(t *testing.T) {
	var cases = []struct {
		name          string
		producer      string
		messageFinish *atomic.Bool
		cancel        bool
	}{
		{name: "TestDispatcher_SetProducerDoneHandlerNotCancel", producer: "producer", cancel: false},
		{name: "TestDispatcher_SetProducerDoneHandlerCancel", producer: "producer", cancel: true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c.messageFinish = &atomic.Bool{}
			w := new(sync.WaitGroup)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
				w.Done()
			})
			d.Put(&TestMessage{producer: c.producer})
			d.SetProducerDoneHandler(c.producer, func(p string, dispatcher *dispatcher.Action[string, *TestMessage]) {
				c.messageFinish.Store(true)
			})
			if c.cancel {
				d.SetProducerDoneHandler(c.producer, nil)
			}
			w.Add(1)
			d.Start()
			w.Wait()
			if c.cancel && c.messageFinish.Load() {
				t.Errorf("%s should cancel, but not", c.name)
			}
		})
	}
}

func TestDispatcher_SetClosedHandler(t *testing.T) {
	var cases = []struct {
		name                  string
		handlerFinishMsgCount *atomic.Int64
		msgTime               time.Duration
		msgCount              int
	}{
		{name: "TestDispatcher_SetClosedHandler_Normal", msgTime: 0, msgCount: 1},
		{name: "TestDispatcher_SetClosedHandler_MessageCount1024", msgTime: 0, msgCount: 1024},
		{name: "TestDispatcher_SetClosedHandler_MessageTime1sMessageCount3", msgTime: 1 * time.Second, msgCount: 3},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c.handlerFinishMsgCount = &atomic.Int64{}
			w := new(sync.WaitGroup)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
				time.Sleep(c.msgTime)
				c.handlerFinishMsgCount.Add(1)
			})
			d.SetClosedHandler(func(dispatcher *dispatcher.Action[string, *TestMessage]) {
				w.Done()
			})
			for i := 0; i < c.msgCount; i++ {
				d.Put(&TestMessage{producer: "producer"})
			}
			w.Add(1)
			d.Start()
			d.Expel()
			w.Wait()
			if c.handlerFinishMsgCount.Load() != int64(c.msgCount) {
				t.Errorf("%s should finish %d messages, but finish %d", c.name, c.msgCount, c.handlerFinishMsgCount.Load())
			}
		})
	}
}

func TestIncrCount(t *testing.T) {
	var cases = []struct {
		name        string
		producer    string
		messageDone *atomic.Bool
	}{
		{name: "TestIncrCount_Normal", producer: "producer"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c.messageDone = &atomic.Bool{}
			w := new(sync.WaitGroup)
			w.Add(1)
			var d *dispatcher.Dispatcher[string, *TestMessage]
			d = dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
				c.messageDone.Store(true)
				d.IncrCount(c.producer, -1)
			})
			d.SetClosedHandler(func(dispatcher *dispatcher.Action[string, *TestMessage]) {
				w.Done()
			})
			d.Start()
			d.IncrCount(c.producer, 1)
			d.Expel()
			d.Put(&TestMessage{producer: c.producer})
			w.Wait()
			if !c.messageDone.Load() {
				t.Errorf("%s should done, but not", c.name)
			}
		})
	}
}

func TestDispatcher_Expel(t *testing.T) {
	var cases = []struct {
		name                  string
		handlerFinishMsgCount *atomic.Int64
		msgTime               time.Duration
		msgCount              int
	}{
		{name: "TestDispatcher_Expel_Normal", msgTime: 0, msgCount: 1},
		{name: "TestDispatcher_Expel_MessageCount1024", msgTime: 0, msgCount: 1024},
		{name: "TestDispatcher_Expel_MessageTime1sMessageCount3", msgTime: 1 * time.Second, msgCount: 3},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c.handlerFinishMsgCount = &atomic.Int64{}
			w := new(sync.WaitGroup)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
				time.Sleep(c.msgTime)
				c.handlerFinishMsgCount.Add(1)
			})
			d.SetClosedHandler(func(dispatcher *dispatcher.Action[string, *TestMessage]) {
				w.Done()
			})
			for i := 0; i < c.msgCount; i++ {
				d.Put(&TestMessage{producer: "producer"})
			}
			w.Add(1)
			d.Start()
			d.Expel()
			w.Wait()
			if c.handlerFinishMsgCount.Load() != int64(c.msgCount) {
				t.Errorf("%s should finish %d messages, but finish %d", c.name, c.msgCount, c.handlerFinishMsgCount.Load())
			}
		})
	}
}

func TestDispatcher_UnExpel(t *testing.T) {
	var cases = []struct {
		name      string
		closed    *atomic.Bool
		isUnExpel bool
		expect    bool
	}{
		{name: "TestDispatcher_UnExpel_Normal", isUnExpel: true, expect: false},
		{name: "TestDispatcher_UnExpel_NotExpel", isUnExpel: false, expect: true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c.closed = &atomic.Bool{}
			w := new(sync.WaitGroup)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
				w.Done()
			})
			d.SetClosedHandler(func(dispatcher *dispatcher.Action[string, *TestMessage]) {
				c.closed.Store(true)
			})
			d.Put(&TestMessage{producer: "producer"})
			w.Add(1)
			if c.isUnExpel {
				d.Expel()
				d.UnExpel()
			} else {
				d.Expel()
			}
			d.Start()
			w.Wait()
			if c.closed.Load() != c.expect {
				t.Errorf("%s should %v, but %v", c.name, c.expect, c.closed.Load())
			}
		})
	}
}

func TestDispatcher_Put(t *testing.T) {
	var cases = []struct {
		name        string
		producer    string
		messageDone *atomic.Bool
	}{
		{name: "TestDispatcher_Put_Normal", producer: "producer"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c.messageDone = &atomic.Bool{}
			w := new(sync.WaitGroup)
			w.Add(1)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
				c.messageDone.Store(true)
				w.Done()
			})
			d.Start()
			d.Put(&TestMessage{producer: c.producer})
			d.Expel()
			w.Wait()
			if !c.messageDone.Load() {
				t.Errorf("%s should done, but not", c.name)
			}
		})
	}
}

func TestDispatcher_Start(t *testing.T) {
	var cases = []struct {
		name        string
		producer    string
		messageDone *atomic.Bool
	}{
		{name: "TestDispatcher_Start_Normal", producer: "producer"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c.messageDone = &atomic.Bool{}
			w := new(sync.WaitGroup)
			w.Add(1)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
				c.messageDone.Store(true)
				w.Done()
			})
			d.Start()
			d.Put(&TestMessage{producer: c.producer})
			d.Expel()
			w.Wait()
			if !c.messageDone.Load() {
				t.Errorf("%s should done, but not", c.name)
			}
		})
	}
}

func TestDispatcher_Name(t *testing.T) {
	var cases = []struct {
		name string
	}{
		{name: "TestDispatcher_Name_Normal"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			if d.Name() != c.name {
				t.Errorf("%s should equal %s, but not", c.name, c.name)
			}
		})
	}
}

func TestDispatcher_Closed(t *testing.T) {
	var cases = []struct {
		name string
	}{
		{name: "TestDispatcher_Closed_Normal"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			w := new(sync.WaitGroup)
			w.Add(1)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			d.SetClosedHandler(func(dispatcher *dispatcher.Action[string, *TestMessage]) { w.Done() })
			d.Start()
			d.Expel()
			w.Wait()
			if !d.Closed() {
				t.Errorf("%s should closed, but not", c.name)
			}
		})
	}
}
