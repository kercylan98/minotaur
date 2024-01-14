# Dispatcher



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/dispatcher)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewDispatcher](#NewDispatcher)|创建一个新的消息分发器 Dispatcher 实例
|[NewManager](#NewManager)|生成消息分发器管理器


> 结构体定义

|结构体|描述
|:--|:--
|[Action](#action)|消息分发器操作器，用于暴露外部可操作的消息分发器函数
|[Handler](#handler)|消息处理器
|[Dispatcher](#dispatcher)|用于服务器消息处理的消息分发器
|[Manager](#manager)|消息分发器管理器
|[Message](#message)|暂无描述...
|[Producer](#producer)|暂无描述...

</details>


#### func NewDispatcher(bufferSize int, name string, handler Handler[P, M])  *Dispatcher[P, M]
<span id="NewDispatcher"></span>
> 创建一个新的消息分发器 Dispatcher 实例
***
#### func NewManager(bufferSize int, handler Handler[P, M])  *Manager[P, M]
<span id="NewManager"></span>
> 生成消息分发器管理器
***
### Action
消息分发器操作器，用于暴露外部可操作的消息分发器函数
```go
type Action[P Producer, M Message[P]] struct {
	unlock bool
	d      *Dispatcher[P, M]
}
```
### Handler
消息处理器
```go
type Handler[P Producer, M Message[P]] struct{}
```
### Dispatcher
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
### Manager
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
### Message

```go
type Message[P comparable] struct{}
```
### Producer

```go
type Producer struct{}
```
