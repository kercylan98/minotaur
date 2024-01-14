# Writeloop



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/writeloop)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewChannel](#NewChannel)|创建基于 Channel 的写循环
|[NewUnbounded](#NewUnbounded)|创建写循环


> 结构体定义

|结构体|描述
|:--|:--
|[Channel](#channel)|基于 chan 的写循环，与 Unbounded 相同，但是使用 Channel 实现
|[Unbounded](#unbounded)|写循环
|[WriteLoop](#writeloop)|暂无描述...

</details>


#### func NewChannel(pool *hub.ObjectPool[Message], channelSize int, writeHandler func (message Message)  error, errorHandler func (err any))  *Channel[Message]
<span id="NewChannel"></span>
> 创建基于 Channel 的写循环
>   - pool 用于管理 Message 对象的缓冲池，在创建 Message 对象时也应该使用该缓冲池，以便复用 Message 对象。 Channel 会在写入完成后将 Message 对象放回缓冲池
>   - channelSize Channel 的大小
>   - writeHandler 写入处理函数
>   - errorHandler 错误处理函数
> 
> 传入 writeHandler 的消息对象是从 Channel 中获取的，因此 writeHandler 不应该持有消息对象的引用，同时也不应该主动释放消息对象
***
#### func NewUnbounded(pool *hub.ObjectPool[Message], writeHandler func (message Message)  error, errorHandler func (err any))  *Unbounded[Message]
<span id="NewUnbounded"></span>
> 创建写循环
>   - pool 用于管理 Message 对象的缓冲池，在创建 Message 对象时也应该使用该缓冲池，以便复用 Message 对象。 Unbounded 会在写入完成后将 Message 对象放回缓冲池
>   - writeHandler 写入处理函数
>   - errorHandler 错误处理函数
> 
> 传入 writeHandler 的消息对象是从 pool 中获取的，并且在 writeHandler 执行完成后会被放回 pool 中，因此 writeHandler 不应该持有消息对象的引用，同时也不应该主动释放消息对象
***
### Channel
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
### Unbounded
写循环
  - 用于将数据并发安全的写入到底层连接
```go
type Unbounded[Message any] struct {
	buf *buffer.Unbounded[Message]
}
```
#### func (*Unbounded) Put(message Message)
> 将数据放入写循环，message 应该来源于 hub.ObjectPool
***
#### func (*Unbounded) Close()
> 关闭写循环
***
### WriteLoop

```go
type WriteLoop[Message any] struct{}
```
