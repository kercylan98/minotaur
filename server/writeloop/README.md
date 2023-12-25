# WriteLoop

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop)

该包提供了一个并发安全的写循环实现。开发者可以使用它来快速构建和管理写入操作。

写循环是一种特殊的循环，它可以并发安全地将数据写入到底层连接。写循环在 `Minotaur` 中是一个泛型类型，可以处理任意类型的消息。

## Unbounded [`基于无界缓冲区的写循环`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Unbounded)

基于无界缓冲区的写循环实现，它可以处理任意数量的消息。它使用 [`Pool`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/concurrent#Pool) 来管理消息对象，使用 [`Unbounded`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/buffer#Unbounded) 来管理消息队列。

> [`Unbounded`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Unbounded) 使用了 [`Pool`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/concurrent#Pool) 和 [`Unbounded`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/buffer#Unbounded) 进行实现。
> 通过 [`Pool`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/concurrent#Pool) 创建的消息对象无需手动释放，它会在写循环处理完消息后自动回收。

## Channel [`基于 chan 的写循环`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Channel)

基于 [`chan`](https://pkg.go.dev/builtin#chan) 的写循环实现，拥有极高的吞吐量。它使用 [`Pool`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/concurrent#Pool) 来管理消息对象，使用 [`Channel`](https://pkg.go.dev/builtin#chan) 来管理消息队列。

## 如何选择
- [`Unbounded`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Unbounded) 适用于消息体较小、消息量庞大、消息处理时间较慢的场景，由于没有消息数量限制，所以它可以处理任意数量的消息，但是它的吞吐量较低，并且会占用大量的内存。
- [`Channel`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Channel) 适用于消息体较大、消息量较小、消息处理时间较快的场景，由于使用了 [`Channel`](https://pkg.go.dev/builtin#chan) 来管理消息队列，所以它的吞吐量较高，但是它的消息数量是有限制的，如果消息数量超过了限制，那么写循环会阻塞，造成无法及时响应消息（例如心跳中断等）。

> 通常来说，建议使用 [`Channel`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Channel) 来实现写循环，因为它的吞吐量更高，并且绝大多数情况是很难达到阻塞的。
### 使用示例
> 该示例由 [`Unbounded`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Unbounded) 实现，[`Channel`](https://pkg.go.dev/github.com/kercylan98/minotaur/server/writeloop#Channel) 的使用方法与之类似。

```go
package main

import (
    "fmt"
	"github.com/kercylan98/minotaur/server/writeloop"
	"github.com/kercylan98/minotaur/utils/concurrent"
)

func main() {
	pool := concurrent.NewPool[Message](func() *Message {
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
}
```

在这个示例中，我们创建了一个写循环，然后将一些消息放入写循环。每个消息都会被 `writeHandle` 函数处理，如果在处理过程中发生错误，`errorHandle` 函数会被调用。在使用完写循环后，我们需要关闭它。