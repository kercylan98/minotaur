# Lockstep



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/lockstep)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewLockstep](#NewLockstep)|创建一个锁步（帧）同步默认实现的组件(Lockstep)进行返回
|[WithFrameLimit](#WithFrameLimit)|通过特定逻辑帧上限创建锁步（帧）同步组件
|[WithFrameRate](#WithFrameRate)|通过特定逻辑帧率创建锁步（帧）同步组件
|[WithSerialization](#WithSerialization)|通过特定的序列化方式将每一帧的数据进行序列化
|[WithInitFrame](#WithInitFrame)|通过特定的初始帧创建锁步（帧）同步组件


> 结构体定义

|结构体|描述
|:--|:--
|[Client](#client)|帧同步客户端接口定义
|[StoppedEventHandle](#stoppedeventhandle)|暂无描述...
|[Lockstep](#lockstep)|锁步（帧）同步默认实现
|[Option](#option)|暂无描述...

</details>


#### func NewLockstep(options ...Option[ClientID, Command])  *Lockstep[ClientID, Command]
<span id="NewLockstep"></span>
> 创建一个锁步（帧）同步默认实现的组件(Lockstep)进行返回
***
#### func WithFrameLimit(frameLimit int64)  Option[ClientID, Command]
<span id="WithFrameLimit"></span>
> 通过特定逻辑帧上限创建锁步（帧）同步组件
>   - 当达到上限时将停止广播
***
#### func WithFrameRate(frameRate int64)  Option[ClientID, Command]
<span id="WithFrameRate"></span>
> 通过特定逻辑帧率创建锁步（帧）同步组件
>   - 默认情况下为 15/s
***
#### func WithSerialization(handle func (frame int64, commands []Command)  []byte)  Option[ClientID, Command]
<span id="WithSerialization"></span>
> 通过特定的序列化方式将每一帧的数据进行序列化
> 
>   - 默认情况下为将被序列化为以下结构体的JSON字符串
> 
>     type Frame struct {
>     Frame int `json:"frame"`
>     Commands []Command `json:"commands"`
>     }
***
#### func WithInitFrame(initFrame int64)  Option[ClientID, Command]
<span id="WithInitFrame"></span>
> 通过特定的初始帧创建锁步（帧）同步组件
>   - 默认情况下为 0，即第一帧索引为 0
***
### Client
帧同步客户端接口定义
  - 客户端应该具备ID及写入数据包的实现
```go
type Client[ID comparable] struct{}
```
### StoppedEventHandle

```go
type StoppedEventHandle[ClientID comparable, Command any] struct{}
```
### Lockstep
锁步（帧）同步默认实现
  - 支持最大帧上限 WithFrameLimit
  - 自定逻辑帧频率，默认为每秒15帧(帧/66ms) WithFrameRate
  - 自定帧序列化方式 WithSerialization
  - 从特定帧开始追帧
  - 兼容各种基于TCP/UDP/Unix的网络类型，可通过客户端实现其他网络类型同步
```go
type Lockstep[ClientID comparable, Command any] struct {
	running                     bool
	runningLock                 sync.RWMutex
	initFrame                   int64
	frameRate                   int64
	frameLimit                  int64
	serialization               func(frame int64, commands []Command) []byte
	clients                     map[ClientID]Client[ClientID]
	clientFrame                 map[ClientID]int64
	clientLock                  sync.RWMutex
	currentFrame                int64
	currentCommands             []Command
	currentFrameLock            sync.RWMutex
	frameCache                  map[int64][]byte
	frameCacheLock              sync.RWMutex
	ticker                      *time.Ticker
	lockstepStoppedEventHandles []StoppedEventHandle[ClientID, Command]
}
```
### Option

```go
type Option[ClientID comparable, Command any] struct{}
```
