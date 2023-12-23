# Buffer

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/buffer)

该包提供了多种缓冲区实现，包括环形缓冲区和无界缓冲区。开发者可以使用它来快速构建和管理缓冲区。

## Ring [`环形缓冲区`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/buffer#Ring)

环形缓冲区是一种特殊的缓冲区，它的头尾是相连的。当缓冲区满时，新的元素会覆盖旧的元素。环形缓冲区在 `Minotaur` 中是一个泛型类型，可以容纳任意类型的元素。

## Unbounded [`无界缓冲区`](https://pkg.go.dev/github.com/kercylan98/minotaur/utils/buffer#Unbounded)

该缓冲区来源于 gRPC 的实现，用于在不使用额外 goroutine 的情况下实现无界缓冲区。无界缓冲区是一种特殊的缓冲区，它的大小可以动态扩展，不会出现溢出的情况。无界缓冲区在 `Minotaur` 中也是一个泛型类型，可以容纳任意类型的元素。

### 使用示例

环形缓冲区：

```go
package main

import (
    "fmt"
    "github.com/kercylan98/minotaur/utils/buffer"
)

func main() {
    ring := buffer.NewRing[int](5)
    for i := 0; i < 5; i++ {
        ring.Write(i)
    }

    for i := 0; i < 5; i++ {
        v, _ := ring.Read()
        fmt.Println(v) // 0 1 2 3 4
    }
}
```

无界缓冲区：

```go
package main

import (
    "fmt"
    "github.com/kercylan98/minotaur/utils/buffer"
)

func main() {
    unbounded := buffer.NewUnboundedN[int]()
    for i := 0; i < 10; i++ {
        unbounded.Put(i)
    }

    for {
        select {
            case v, ok := <-unbounded.Get():
                if !ok {
                    return
                }
                unbounded.Load()
                fmt.Println(v) // 0 1 2 3 4 5 6 7 8 9
                if v == 9 {
                    unbounded.Close()
                    return
                }
        } 
    }
}
```