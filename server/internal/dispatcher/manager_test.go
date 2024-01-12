package dispatcher_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server/internal/dispatcher"
	"sync/atomic"
	"testing"
)

func TestNewManager(t *testing.T) {
	var cases = []struct {
		name        string
		bufferSize  int
		handler     dispatcher.Handler[string, *TestMessage]
		shouldPanic bool
	}{
		{name: "TestNewManager_BufferSize0AndHandlerNil", bufferSize: 0, handler: nil, shouldPanic: true},
		{name: "TestNewManager_BufferSize0AndHandlerNotNil", bufferSize: 0, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {}, shouldPanic: true},
		{name: "TestNewManager_BufferSize1AndHandlerNil", bufferSize: 1, handler: nil, shouldPanic: true},
		{name: "TestNewManager_BufferSize1AndHandlerNotNil", bufferSize: 1, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {}, shouldPanic: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !c.shouldPanic {
					t.Errorf("NewManager() should not panic, but panic: %v", r)
				}
			}()
			dispatcher.NewManager[string, *TestMessage](c.bufferSize, c.handler)
		})
	}
}

func TestManager_SetDispatcherClosedHandler(t *testing.T) {
	var cases = []struct {
		name            string
		setCloseHandler bool
	}{
		{name: "TestManager_SetDispatcherClosedHandler_Set", setCloseHandler: true},
		{name: "TestManager_SetDispatcherClosedHandler_NotSet", setCloseHandler: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var closed atomic.Bool
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			if c.setCloseHandler {
				m.SetDispatcherClosedHandler(func(name string) {
					closed.Store(true)
				})
			}
			m.BindProducer(c.name, c.name)
			m.UnBindProducer(c.name)
			m.Wait()
			if c.setCloseHandler && !closed.Load() {
				t.Errorf("SetDispatcherClosedHandler() should be called")
			}

		})
	}
}

func TestManager_SetDispatcherCreatedHandler(t *testing.T) {
	var cases = []struct {
		name              string
		setCreatedHandler bool
	}{
		{name: "TestManager_SetDispatcherCreatedHandler_Set", setCreatedHandler: true},
		{name: "TestManager_SetDispatcherCreatedHandler_NotSet", setCreatedHandler: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var created atomic.Bool
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			if c.setCreatedHandler {
				m.SetDispatcherCreatedHandler(func(name string) {
					created.Store(true)
				})
			}
			m.BindProducer(c.name, c.name)
			m.UnBindProducer(c.name)
			m.Wait()
			if c.setCreatedHandler && !created.Load() {
				t.Errorf("SetDispatcherCreatedHandler() should be called")
			}

		})
	}
}

func TestManager_HasDispatcher(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
		has      bool
	}{
		{name: "TestManager_HasDispatcher_Has", bindName: "TestManager_HasDispatcher_Has", has: true},
		{name: "TestManager_HasDispatcher_NotHas", bindName: "TestManager_HasDispatcher_NotHas", has: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			m.BindProducer(c.bindName, c.bindName)
			var cond string
			if c.has {
				cond = c.bindName
			}
			if m.HasDispatcher(cond) != c.has {
				t.Errorf("HasDispatcher() should return %v", c.has)
			}
		})
	}
}

func TestManager_GetDispatcherNum(t *testing.T) {
	var cases = []struct {
		name string
		num  int
	}{
		{name: "TestManager_GetDispatcherNum_N1", num: -1},
		{name: "TestManager_GetDispatcherNum_0", num: 0},
		{name: "TestManager_GetDispatcherNum_1", num: 1},
		{name: "TestManager_GetDispatcherNum_2", num: 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			switch {
			case c.num <= 0:
				return
			case c.num == 1:
				if m.GetDispatcherNum() != 1 {
					t.Errorf("GetDispatcherNum() should return 1")
				}
				return
			default:
				for i := 0; i < c.num-1; i++ {
					m.BindProducer(fmt.Sprintf("%s_%d", c.name, i), fmt.Sprintf("%s_%d", c.name, i))
				}
				if m.GetDispatcherNum() != c.num {
					t.Errorf("GetDispatcherNum() should return %v", c.num)
				}
			}
		})
	}
}

func TestManager_GetSystemDispatcher(t *testing.T) {
	var cases = []struct {
		name string
	}{
		{name: "TestManager_GetSystemDispatcher"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			if m.GetSystemDispatcher() == nil {
				t.Errorf("GetSystemDispatcher() should not return nil")
			}
		})
	}
}

func TestManager_GetDispatcher(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
	}{
		{name: "TestManager_GetDispatcher", bindName: "TestManager_GetDispatcher"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			m.BindProducer(c.bindName, c.bindName)
			if m.GetDispatcher(c.bindName) == nil {
				t.Errorf("GetDispatcher() should not return nil")
			}
		})
	}
}

func TestManager_BindProducer(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
	}{
		{name: "TestManager_BindProducer", bindName: "TestManager_BindProducer"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			m.BindProducer(c.bindName, c.bindName)
			if m.GetDispatcher(c.bindName) == nil {
				t.Errorf("GetDispatcher() should not return nil")
			}
		})
	}
}

func TestManager_UnBindProducer(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
	}{
		{name: "TestManager_UnBindProducer", bindName: "TestManager_UnBindProducer"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {})
			m.BindProducer(c.bindName, c.bindName)
			m.UnBindProducer(c.bindName)
			if m.GetDispatcher(c.bindName) != m.GetSystemDispatcher() {
				t.Errorf("GetDispatcher() should return SystemDispatcher")
			}
		})
	}
}
