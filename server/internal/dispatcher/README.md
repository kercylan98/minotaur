# Dispatcher

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

暂无介绍...


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewDispatcher](#NewDispatcher)|创建一个新的消息分发器 Dispatcher 实例
|[NewManager](#NewManager)|生成消息分发器管理器


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Action](#struct_Action)|消息分发器操作器，用于暴露外部可操作的消息分发器函数
|`STRUCT`|[Handler](#struct_Handler)|消息处理器
|`STRUCT`|[Dispatcher](#struct_Dispatcher)|用于服务器消息处理的消息分发器
|`STRUCT`|[Manager](#struct_Manager)|消息分发器管理器
|`INTERFACE`|[Message](#struct_Message)|暂无描述...
|`INTERFACE`|[Producer](#struct_Producer)|暂无描述...

</details>


***
## 详情信息
#### func NewDispatcher\[P Producer, M Message[P]\](bufferSize int, name string, handler Handler[P, M]) *Dispatcher[P, M]
<span id="NewDispatcher"></span>
> 创建一个新的消息分发器 Dispatcher 实例

**示例代码：**

```go

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
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewDispatcher(t *testing.T) {
	var cases = []struct {
		name        string
		bufferSize  int
		handler     dispatcher.Handler[string, *TestMessage]
		shouldPanic bool
	}{{name: "TestNewDispatcher_BufferSize0AndHandlerNil", bufferSize: 0, handler: nil, shouldPanic: true}, {name: "TestNewDispatcher_BufferSize0AndHandlerNotNil", bufferSize: 0, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
	}, shouldPanic: true}, {name: "TestNewDispatcher_BufferSize1AndHandlerNil", bufferSize: 1, handler: nil, shouldPanic: true}, {name: "TestNewDispatcher_BufferSize1AndHandlerNotNil", bufferSize: 1, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
	}, shouldPanic: false}}
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

```


</details>


***
#### func NewManager\[P Producer, M Message[P]\](bufferSize int, handler Handler[P, M]) *Manager[P, M]
<span id="NewManager"></span>
> 生成消息分发器管理器

**示例代码：**

```go

func ExampleNewManager() {
	mgr := dispatcher.NewManager[string, *TestMessage](10124*16, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
	})
	mgr.BindProducer("player_001", "shunt-001")
	mgr.BindProducer("player_002", "shunt-002")
	mgr.BindProducer("player_003", "shunt-sys")
	mgr.BindProducer("player_004", "shunt-sys")
	mgr.UnBindProducer("player_001")
	mgr.UnBindProducer("player_002")
	mgr.UnBindProducer("player_003")
	mgr.UnBindProducer("player_004")
	mgr.Wait()
	fmt.Println("done")
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewManager(t *testing.T) {
	var cases = []struct {
		name        string
		bufferSize  int
		handler     dispatcher.Handler[string, *TestMessage]
		shouldPanic bool
	}{{name: "TestNewManager_BufferSize0AndHandlerNil", bufferSize: 0, handler: nil, shouldPanic: true}, {name: "TestNewManager_BufferSize0AndHandlerNotNil", bufferSize: 0, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
	}, shouldPanic: true}, {name: "TestNewManager_BufferSize1AndHandlerNil", bufferSize: 1, handler: nil, shouldPanic: true}, {name: "TestNewManager_BufferSize1AndHandlerNotNil", bufferSize: 1, handler: func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
	}, shouldPanic: false}}
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

```


</details>


***
<span id="struct_Action"></span>
### Action `STRUCT`
消息分发器操作器，用于暴露外部可操作的消息分发器函数
```go
type Action[P Producer, M Message[P]] struct {
	unlock bool
	d      *Dispatcher[P, M]
}
```
#### func (*Action) Name()  string
> 获取消息分发器名称
***
#### func (*Action) UnExpel()
> 取消特定生产者的驱逐计划
***
#### func (*Action) Expel()
> 设置该消息分发器即将被驱逐，当消息分发器中没有任何消息时，会自动关闭
***
<span id="struct_Handler"></span>
### Handler `STRUCT`
消息处理器
```go
type Handler[P Producer, M Message[P]] func(dispatcher *Dispatcher[P, M], message M)
```
<span id="struct_Dispatcher"></span>
### Dispatcher `STRUCT`
用于服务器消息处理的消息分发器

这个消息分发器为并发安全的生产者和消费者模型，生产者可以是任意类型，消费者必须是 Message 接口的实现。
生产者可以通过 Put 方法并发安全地将消息放入消息分发器，消息执行过程不会阻塞到 Put 方法，同时允许在 Start 方法之前调用 Put 方法。
在执行 Start 方法后，消息分发器会阻塞地从消息缓冲区中读取消息，然后执行消息处理器，消息处理器的执行过程不会阻塞到消息的生产。

为了保证消息不丢失，内部采用了 buffer.RingUnbounded 作为缓冲区实现，并且消息分发器不提供 Close 方法。
如果需要关闭消息分发器，可以通过 Expel 方法设置驱逐计划，当消息分发器中没有任何消息时，将会被释放。
同时，也可以使用 UnExpel 方法取消驱逐计划。

为什么提供 Expel 和 UnExpel 方法：
  - 在连接断开时，当需要执行一系列消息处理时，如果直接关闭消息分发器，可能会导致消息丢失。所以提供了 Expel 方法，可以在消息处理完成后再关闭消息分发器。
  - 当消息还未处理完成时连接重连，如果没有取消驱逐计划，可能会导致消息分发器被关闭。所以提供了 UnExpel 方法，可以在连接重连后取消驱逐计划。
```go
type Dispatcher[P Producer, M Message[P]] struct {
	buf           *buffer.RingUnbounded[M]
	uniques       *haxmap.Map[string, struct{}]
	handler       Handler[P, M]
	expel         bool
	mc            int64
	pmc           map[P]int64
	pmcF          map[P]func(p P, dispatcher *Action[P, M])
	lock          sync.RWMutex
	name          string
	closedHandler atomic.Pointer[func(dispatcher *Action[P, M])]
	abort         chan struct{}
}
```
#### func (*Dispatcher) SetProducerDoneHandler(p P, handler func (p P, dispatcher *Action[P, M]))  *Dispatcher[P, M]
> 设置特定生产者所有消息处理完成时的回调函数
>   - 如果 handler 为 nil，则会删除该生产者的回调函数
> 
> 需要注意的是，该 handler 中
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_SetProducerDoneHandler(t *testing.T) {
	var cases = []struct {
		name          string
		producer      string
		messageFinish *atomic.Bool
		cancel        bool
	}{{name: "TestDispatcher_SetProducerDoneHandlerNotCancel", producer: "producer", cancel: false}, {name: "TestDispatcher_SetProducerDoneHandlerCancel", producer: "producer", cancel: true}}
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

```


</details>


***
#### func (*Dispatcher) SetClosedHandler(handler func (dispatcher *Action[P, M]))  *Dispatcher[P, M]
> 设置消息分发器关闭时的回调函数
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_SetClosedHandler(t *testing.T) {
	var cases = []struct {
		name                  string
		handlerFinishMsgCount *atomic.Int64
		msgTime               time.Duration
		msgCount              int
	}{{name: "TestDispatcher_SetClosedHandler_Normal", msgTime: 0, msgCount: 1}, {name: "TestDispatcher_SetClosedHandler_MessageCount1024", msgTime: 0, msgCount: 1024}, {name: "TestDispatcher_SetClosedHandler_MessageTime1sMessageCount3", msgTime: 1 * time.Second, msgCount: 3}}
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

```


</details>


***
#### func (*Dispatcher) Name()  string
> 获取消息分发器名称
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_Name(t *testing.T) {
	var cases = []struct{ name string }{{name: "TestDispatcher_Name_Normal"}}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
			if d.Name() != c.name {
				t.Errorf("%s should equal %s, but not", c.name, c.name)
			}
		})
	}
}

```


</details>


***
#### func (*Dispatcher) Unique(name string)  bool
> 设置唯一消息键，返回是否已存在
***
#### func (*Dispatcher) AntiUnique(name string)
> 取消唯一消息键
***
#### func (*Dispatcher) Expel()
> 设置该消息分发器即将被驱逐，当消息分发器中没有任何消息时，会自动关闭
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_Expel(t *testing.T) {
	var cases = []struct {
		name                  string
		handlerFinishMsgCount *atomic.Int64
		msgTime               time.Duration
		msgCount              int
	}{{name: "TestDispatcher_Expel_Normal", msgTime: 0, msgCount: 1}, {name: "TestDispatcher_Expel_MessageCount1024", msgTime: 0, msgCount: 1024}, {name: "TestDispatcher_Expel_MessageTime1sMessageCount3", msgTime: 1 * time.Second, msgCount: 3}}
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

```


</details>


***
#### func (*Dispatcher) UnExpel()
> 取消特定生产者的驱逐计划
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_UnExpel(t *testing.T) {
	var cases = []struct {
		name      string
		closed    *atomic.Bool
		isUnExpel bool
		expect    bool
	}{{name: "TestDispatcher_UnExpel_Normal", isUnExpel: true, expect: false}, {name: "TestDispatcher_UnExpel_NotExpel", isUnExpel: false, expect: true}}
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

```


</details>


***
#### func (*Dispatcher) IncrCount(producer P, i int64)
> 主动增量设置特定生产者的消息计数，这在等待异步消息完成后再关闭消息分发器时非常有用
>   - 如果 i 为负数，则会减少消息计数
***
#### func (*Dispatcher) Put(message M)
> 将消息放入分发器
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_Put(t *testing.T) {
	var cases = []struct {
		name        string
		producer    string
		messageDone *atomic.Bool
	}{{name: "TestDispatcher_Put_Normal", producer: "producer"}}
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

```


</details>


***
#### func (*Dispatcher) Start()  *Dispatcher[P, M]
> 以非阻塞的方式开始进行消息分发，当消息分发器中没有任何消息并且处于驱逐计划 Expel 时，将会自动关闭
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_Start(t *testing.T) {
	var cases = []struct {
		name        string
		producer    string
		messageDone *atomic.Bool
	}{{name: "TestDispatcher_Start_Normal", producer: "producer"}}
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

```


</details>


***
#### func (*Dispatcher) Closed()  bool
> 判断消息分发器是否已关闭
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestDispatcher_Closed(t *testing.T) {
	var cases = []struct{ name string }{{name: "TestDispatcher_Closed_Normal"}}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			w := new(sync.WaitGroup)
			w.Add(1)
			d := dispatcher.NewDispatcher(1024, c.name, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
			d.SetClosedHandler(func(dispatcher *dispatcher.Action[string, *TestMessage]) {
				w.Done()
			})
			d.Start()
			d.Expel()
			w.Wait()
			if !d.Closed() {
				t.Errorf("%s should closed, but not", c.name)
			}
		})
	}
}

```


</details>


***
<span id="struct_Manager"></span>
### Manager `STRUCT`
消息分发器管理器
```go
type Manager[P Producer, M Message[P]] struct {
	handler        Handler[P, M]
	sys            *Dispatcher[P, M]
	dispatchers    map[string]*Dispatcher[P, M]
	member         map[string]map[P]struct{}
	curr           map[P]*Dispatcher[P, M]
	lock           sync.RWMutex
	w              sync.WaitGroup
	size           int
	closedHandler  func(name string)
	createdHandler func(name string)
}
```
#### func (*Manager) Wait()
> 等待所有消息分发器关闭
***
#### func (*Manager) SetDispatcherClosedHandler(handler func (name string))  *Manager[P, M]
> 设置消息分发器关闭时的回调函数
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_SetDispatcherClosedHandler(t *testing.T) {
	var cases = []struct {
		name            string
		setCloseHandler bool
	}{{name: "TestManager_SetDispatcherClosedHandler_Set", setCloseHandler: true}, {name: "TestManager_SetDispatcherClosedHandler_NotSet", setCloseHandler: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var closed atomic.Bool
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
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

```


</details>


***
#### func (*Manager) SetDispatcherCreatedHandler(handler func (name string))  *Manager[P, M]
> 设置消息分发器创建时的回调函数
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_SetDispatcherCreatedHandler(t *testing.T) {
	var cases = []struct {
		name              string
		setCreatedHandler bool
	}{{name: "TestManager_SetDispatcherCreatedHandler_Set", setCreatedHandler: true}, {name: "TestManager_SetDispatcherCreatedHandler_NotSet", setCreatedHandler: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var created atomic.Bool
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
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

```


</details>


***
#### func (*Manager) HasDispatcher(name string)  bool
> 检查是否存在指定名称的消息分发器
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_HasDispatcher(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
		has      bool
	}{{name: "TestManager_HasDispatcher_Has", bindName: "TestManager_HasDispatcher_Has", has: true}, {name: "TestManager_HasDispatcher_NotHas", bindName: "TestManager_HasDispatcher_NotHas", has: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
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

```


</details>


***
#### func (*Manager) GetDispatcherNum()  int
> 获取当前正在工作的消息分发器数量
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_GetDispatcherNum(t *testing.T) {
	var cases = []struct {
		name string
		num  int
	}{{name: "TestManager_GetDispatcherNum_N1", num: -1}, {name: "TestManager_GetDispatcherNum_0", num: 0}, {name: "TestManager_GetDispatcherNum_1", num: 1}, {name: "TestManager_GetDispatcherNum_2", num: 2}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
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

```


</details>


***
#### func (*Manager) GetSystemDispatcher()  *Dispatcher[P, M]
> 获取系统消息分发器
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_GetSystemDispatcher(t *testing.T) {
	var cases = []struct{ name string }{{name: "TestManager_GetSystemDispatcher"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
			if m.GetSystemDispatcher() == nil {
				t.Errorf("GetSystemDispatcher() should not return nil")
			}
		})
	}
}

```


</details>


***
#### func (*Manager) GetDispatcher(p P)  *Dispatcher[P, M]
> 获取生产者正在使用的消息分发器，如果生产者没有绑定消息分发器，则会返回系统消息分发器
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_GetDispatcher(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
	}{{name: "TestManager_GetDispatcher", bindName: "TestManager_GetDispatcher"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
			m.BindProducer(c.bindName, c.bindName)
			if m.GetDispatcher(c.bindName) == nil {
				t.Errorf("GetDispatcher() should not return nil")
			}
		})
	}
}

```


</details>


***
#### func (*Manager) BindProducer(p P, name string)
> 绑定生产者使用特定的消息分发器，如果生产者已经绑定了消息分发器，则会先解绑
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_BindProducer(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
	}{{name: "TestManager_BindProducer", bindName: "TestManager_BindProducer"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
			m.BindProducer(c.bindName, c.bindName)
			if m.GetDispatcher(c.bindName) == nil {
				t.Errorf("GetDispatcher() should not return nil")
			}
		})
	}
}

```


</details>


***
#### func (*Manager) UnBindProducer(p P)
> 解绑生产者使用特定的消息分发器
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestManager_UnBindProducer(t *testing.T) {
	var cases = []struct {
		name     string
		bindName string
	}{{name: "TestManager_UnBindProducer", bindName: "TestManager_UnBindProducer"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := dispatcher.NewManager[string, *TestMessage](1024, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
			})
			m.BindProducer(c.bindName, c.bindName)
			m.UnBindProducer(c.bindName)
			if m.GetDispatcher(c.bindName) != m.GetSystemDispatcher() {
				t.Errorf("GetDispatcher() should return SystemDispatcher")
			}
		})
	}
}

```


</details>


***
<span id="struct_Message"></span>
### Message `INTERFACE`

```go
type Message[P comparable] interface {
	GetProducer() P
}
```
<span id="struct_Producer"></span>
### Producer `INTERFACE`

```go
type Producer interface {
	comparable
}
```
