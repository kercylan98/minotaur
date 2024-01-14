# Buffer

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

buffer 提供了缓冲区相关的实用程序。

包括创建、读取和写入缓冲区的函数。

这个包还提供了一个无界缓冲区的实现，可以在不使用额外 goroutine 的情况下实现无界缓冲区。

无界缓冲区的所有方法都是线程安全的，除了用于同步的互斥锁外，不会阻塞任何东西。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewRing](#NewRing)|创建一个并发不安全的环形缓冲区
|[NewRingUnbounded](#NewRingUnbounded)|创建一个并发安全的基于环形缓冲区实现的无界缓冲区
|[NewUnbounded](#NewUnbounded)|创建一个无界缓冲区


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Ring](#ring)|环形缓冲区
|`STRUCT`|[RingUnbounded](#ringunbounded)|基于环形缓冲区实现的无界缓冲区
|`STRUCT`|[Unbounded](#unbounded)|是无界缓冲区的实现

</details>


***
## 详情信息
#### func NewRing(initSize ...int)  *Ring[T]
<span id="NewRing"></span>
> 创建一个并发不安全的环形缓冲区
>   - initSize: 初始容量
> 
> 当初始容量小于 2 或未设置时，将会使用默认容量 2

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewRing(t *testing.T) {
	ring := buffer.NewRing[int]()
	for i := 0; i < 100; i++ {
		ring.Write(i)
		t.Log(ring.Read())
	}
}

```


</details>


***
#### func NewRingUnbounded(bufferSize int)  *RingUnbounded[T]
<span id="NewRingUnbounded"></span>
> 创建一个并发安全的基于环形缓冲区实现的无界缓冲区

***
#### func NewUnbounded()  *Unbounded[V]
<span id="NewUnbounded"></span>
> 创建一个无界缓冲区
>   - generateNil: 生成空值的函数，该函数仅需始终返回 nil 即可
> 
> 该缓冲区来源于 gRPC 的实现，用于在不使用额外 goroutine 的情况下实现无界缓冲区
>   - 该缓冲区的所有方法都是线程安全的，除了用于同步的互斥锁外，不会阻塞任何东西

***
### Ring `STRUCT`
环形缓冲区
```go
type Ring[T any] struct {
	buf      []T
	initSize int
	size     int
	r        int
	w        int
}
```
#### func (*Ring) Read()  T,  error
> 读取数据
<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkRing_Read(b *testing.B) {
	ring := buffer.NewRing[int](1024)
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ring.Read()
	}
}

```


</details>


***
#### func (*Ring) ReadAll()  []T
> 读取所有数据
***
#### func (*Ring) Peek() (t T, err error)
> 查看数据
***
#### func (*Ring) Write(v T)
> 写入数据
<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkRing_Write(b *testing.B) {
	ring := buffer.NewRing[int](1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
}

```


</details>


***
#### func (*Ring) IsEmpty()  bool
> 是否为空
***
#### func (*Ring) Cap()  int
> 返回缓冲区容量
***
#### func (*Ring) Len()  int
> 返回缓冲区长度
***
#### func (*Ring) Reset()
> 重置缓冲区
***
### RingUnbounded `STRUCT`
基于环形缓冲区实现的无界缓冲区
```go
type RingUnbounded[T any] struct {
	ring         *Ring[T]
	rrm          sync.Mutex
	cond         *sync.Cond
	rc           chan T
	closed       bool
	closedMutex  sync.RWMutex
	closedSignal chan struct{}
}
```
#### func (*RingUnbounded) Write(v T)
> 写入数据
<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkRingUnbounded_Write(b *testing.B) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
}

```


</details>


***
#### func (*RingUnbounded) Read()  chan T
> 读取数据
<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkRingUnbounded_Read(b *testing.B) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	for i := 0; i < b.N; i++ {
		ring.Write(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-ring.Read()
	}
}

```


</details>


***
#### func (*RingUnbounded) Closed()  bool
> 判断缓冲区是否已关闭
***
#### func (*RingUnbounded) Close()  chan struct {}
> 关闭缓冲区，关闭后将不再接收新数据，但是已有数据仍然可以读取
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestRingUnbounded_Close(t *testing.T) {
	ring := buffer.NewRingUnbounded[int](1024 * 16)
	for i := 0; i < 100; i++ {
		ring.Write(i)
	}
	t.Log("write done")
	ring.Close()
	t.Log("close done")
	for v := range ring.Read() {
		ring.Write(v)
		t.Log(v)
	}
	t.Log("read done")
}

```


</details>


***
### Unbounded `STRUCT`
是无界缓冲区的实现
```go
type Unbounded[V any] struct {
	c       chan V
	closed  bool
	mu      sync.Mutex
	backlog []V
}
```
#### func (*Unbounded) Put(t V)
> 将数据放入缓冲区
***
#### func (*Unbounded) Load()
> 将缓冲区中的数据发送到读取通道中，如果缓冲区中没有数据，则不会发送
>   - 在每次 Get 后都应该执行该函数
***
#### func (*Unbounded) Get()  chan V
> 获取读取通道
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestUnbounded_Get(t *testing.T) {
	ub := buffer.NewUnbounded[int]()
	for i := 0; i < 100; i++ {
		ub.Put(i + 1)
		fmt.Println(<-ub.Get())
		ub.Load()
	}
}

```


</details>


***
#### func (*Unbounded) Close()
> 关闭
***
#### func (*Unbounded) IsClosed()  bool
> 是否已关闭
***
