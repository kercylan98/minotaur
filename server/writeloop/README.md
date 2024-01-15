# Writeloop

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
|[NewChannel](#NewChannel)|创建基于 Channel 的写循环
|[NewUnbounded](#NewUnbounded)|创建写循环


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Channel](#channel)|基于 chan 的写循环，与 Unbounded 相同，但是使用 Channel 实现
|`STRUCT`|[Unbounded](#unbounded)|写循环
|`INTERFACE`|[WriteLoop](#writeloop)|暂无描述...

</details>


***
## 详情信息
#### func NewChannel(pool *hub.ObjectPool[Message], channelSize int, writeHandler func (message Message)  error, errorHandler func (err any)) *Channel[Message]
<span id="NewChannel"></span>
> 创建基于 Channel 的写循环
>   - pool 用于管理 Message 对象的缓冲池，在创建 Message 对象时也应该使用该缓冲池，以便复用 Message 对象。 Channel 会在写入完成后将 Message 对象放回缓冲池
>   - channelSize Channel 的大小
>   - writeHandler 写入处理函数
>   - errorHandler 错误处理函数
> 
> 传入 writeHandler 的消息对象是从 Channel 中获取的，因此 writeHandler 不应该持有消息对象的引用，同时也不应该主动释放消息对象

***
#### func NewUnbounded(pool *hub.ObjectPool[Message], writeHandler func (message Message)  error, errorHandler func (err any)) *Unbounded[Message]
<span id="NewUnbounded"></span>
> 创建写循环
>   - pool 用于管理 Message 对象的缓冲池，在创建 Message 对象时也应该使用该缓冲池，以便复用 Message 对象。 Unbounded 会在写入完成后将 Message 对象放回缓冲池
>   - writeHandler 写入处理函数
>   - errorHandler 错误处理函数
> 
> 传入 writeHandler 的消息对象是从 pool 中获取的，并且在 writeHandler 执行完成后会被放回 pool 中，因此 writeHandler 不应该持有消息对象的引用，同时也不应该主动释放消息对象

示例代码：
```go

func ExampleNewUnbounded() {
	pool := hub.NewObjectPool[Message](func() *Message {
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

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewUnbounded(t *testing.T) {
	wl := writeloop.NewUnbounded(wp, func(message *Message) error {
		t.Log(message.ID)
		return nil
	}, func(err any) {
		t.Log(err)
	})
	assert.NotNil(t, wl)
	wl.Close()
}

```


</details>


***
### Channel `STRUCT`
基于 chan 的写循环，与 Unbounded 相同，但是使用 Channel 实现
```go
type Channel[T any] struct {
	c chan T
}
```
#### func (*Channel) Put(message T)
> 将数据放入写循环，message 应该来源于 hub.ObjectPool
***
#### func (*Channel) Close()
> 关闭写循环
***
### Unbounded `STRUCT`
写循环
  - 用于将数据并发安全的写入到底层连接
```go
type Unbounded[Message any] struct {
	buf *buffer.Unbounded[Message]
}
```
#### func (*Unbounded) Put(message Message)
> 将数据放入写循环，message 应该来源于 hub.ObjectPool
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestUnbounded_Put(t *testing.T) {
	wl := writeloop.NewUnbounded(wp, func(message *Message) error {
		t.Log(message.ID)
		return nil
	}, func(err any) {
		t.Log(err)
	})
	assert.NotNil(t, wl)
	for i := 0; i < 100; i++ {
		m := wp.Get()
		m.ID = i
		wl.Put(m)
	}
	wl.Close()
}

```


</details>


<details>
<summary>查看 / 收起基准测试</summary>


```go

func BenchmarkUnbounded_Put(b *testing.B) {
	wl := writeloop.NewUnbounded(wp, func(message *Message) error {
		return nil
	}, nil)
	defer func() {
		wl.Close()
	}()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wl.Put(wp.Get())
		}
	})
	b.StopTimer()
}

```


</details>


***
#### func (*Unbounded) Close()
> 关闭写循环
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestUnbounded_Close(t *testing.T) {
	wl := writeloop.NewUnbounded(wp, func(message *Message) error {
		t.Log(message.ID)
		return nil
	}, func(err any) {
		t.Log(err)
	})
	assert.NotNil(t, wl)
	for i := 0; i < 100; i++ {
		m := wp.Get()
		m.ID = i
		wl.Put(m)
	}
	wl.Close()
}

```


</details>


***
### WriteLoop `INTERFACE`

```go
type WriteLoop[Message any] interface {
	Put(message Message)
	Close()
}
```
