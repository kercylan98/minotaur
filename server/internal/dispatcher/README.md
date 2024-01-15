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
#### func NewDispatcher(bufferSize int, name string, handler Handler[P, M]) *Dispatcher[P, M]
<span id="NewDispatcher"></span>
> 创建一个新的消息分发器 Dispatcher 实例

示例代码：
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
#### func NewManager(bufferSize int, handler Handler[P, M]) *Manager[P, M]
<span id="NewManager"></span>
> 生成消息分发器管理器

示例代码：
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
