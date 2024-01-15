# Lockstep

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
|[NewLockstep](#NewLockstep)|创建一个锁步（帧）同步默认实现的组件(Lockstep)进行返回
|[WithFrameLimit](#WithFrameLimit)|通过特定逻辑帧上限创建锁步（帧）同步组件
|[WithFrameRate](#WithFrameRate)|通过特定逻辑帧率创建锁步（帧）同步组件
|[WithSerialization](#WithSerialization)|通过特定的序列化方式将每一帧的数据进行序列化
|[WithInitFrame](#WithInitFrame)|通过特定的初始帧创建锁步（帧）同步组件


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`INTERFACE`|[Client](#struct_Client)|帧同步客户端接口定义
|`STRUCT`|[StoppedEventHandle](#struct_StoppedEventHandle)|暂无描述...
|`STRUCT`|[Lockstep](#struct_Lockstep)|锁步（帧）同步默认实现
|`STRUCT`|[Option](#struct_Option)|暂无描述...

</details>


***
## 详情信息
#### func NewLockstep\[ClientID comparable, Command any\](options ...Option[ClientID, Command]) *Lockstep[ClientID, Command]
<span id="NewLockstep"></span>
> 创建一个锁步（帧）同步默认实现的组件(Lockstep)进行返回

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewLockstep(t *testing.T) {
	ls := lockstep.NewLockstep[string, int](lockstep.WithInitFrame[string, int](1))
	ls.JoinClient(&Cli{id: "player_1"})
	ls.JoinClient(&Cli{id: "player_2"})
	count := 0
	ls.StartBroadcast()
	endChan := make(chan bool)
	go func() {
		for {
			ls.AddCommand(random.Int(1, 9999))
			count++
			if count >= 10 {
				break
			}
			time.Sleep(time.Millisecond * time.Duration(random.Int(10, 200)))
		}
		ls.StopBroadcast()
		endChan <- true
	}()
	<-endChan
	time.Sleep(time.Second)
	fmt.Println("end")
}

```


</details>


***
#### func WithFrameLimit\[ClientID comparable, Command any\](frameLimit int64) Option[ClientID, Command]
<span id="WithFrameLimit"></span>
> 通过特定逻辑帧上限创建锁步（帧）同步组件
>   - 当达到上限时将停止广播

***
#### func WithFrameRate\[ClientID comparable, Command any\](frameRate int64) Option[ClientID, Command]
<span id="WithFrameRate"></span>
> 通过特定逻辑帧率创建锁步（帧）同步组件
>   - 默认情况下为 15/s

***
#### func WithSerialization\[ClientID comparable, Command any\](handle func (frame int64, commands []Command)  []byte) Option[ClientID, Command]
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
#### func WithInitFrame\[ClientID comparable, Command any\](initFrame int64) Option[ClientID, Command]
<span id="WithInitFrame"></span>
> 通过特定的初始帧创建锁步（帧）同步组件
>   - 默认情况下为 0，即第一帧索引为 0

***
<span id="struct_Client"></span>
### Client `INTERFACE`
帧同步客户端接口定义
  - 客户端应该具备ID及写入数据包的实现
```go
type Client[ID comparable] interface {
	GetID() ID
	Write(packet []byte, callback ...func(err error))
}
```
<span id="struct_StoppedEventHandle"></span>
### StoppedEventHandle `STRUCT`

```go
type StoppedEventHandle[ClientID comparable, Command any] func(lockstep *Lockstep[ClientID, Command])
```
<span id="struct_Lockstep"></span>
### Lockstep `STRUCT`
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
<span id="struct_Lockstep_JoinClient"></span>

#### func (*Lockstep) JoinClient(client Client[ClientID])
> 将客户端加入到广播队列中，通常在开始广播前使用
>   - 如果客户端在开始广播后加入，将丢失之前的帧数据，如要从特定帧开始追帧请使用 JoinClientWithFrame

***
<span id="struct_Lockstep_JoinClientWithFrame"></span>

#### func (*Lockstep) JoinClientWithFrame(client Client[ClientID], frameIndex int64)
> 加入客户端到广播队列中，并从特定帧开始追帧
>   - 可用于重连及状态同步、帧同步混用的情况
>   - 混用：服务端记录指令时同时做一次状态计算，新客户端加入时直接同步当前状态，之后从特定帧开始广播

***
<span id="struct_Lockstep_GetClientCount"></span>

#### func (*Lockstep) GetClientCount()  int
> 获取客户端数量

***
<span id="struct_Lockstep_DropCache"></span>

#### func (*Lockstep) DropCache(handler func (frame int64)  bool)
> 丢弃特定帧的缓存，当 handler 返回 true 时将丢弃缓存

***
<span id="struct_Lockstep_LeaveClient"></span>

#### func (*Lockstep) LeaveClient(clientId ClientID)
> 将客户端从广播队列中移除

***
<span id="struct_Lockstep_StartBroadcast"></span>

#### func (*Lockstep) StartBroadcast()
> 开始广播
>   - 在开始广播后将持续按照设定的帧率进行帧数推进，并在每一帧推进时向客户端进行同步，需提前将客户端加入广播队列 JoinClient
>   - 广播过程中使用 AddCommand 将该帧数据追加到当前帧中

***
<span id="struct_Lockstep_StopBroadcast"></span>

#### func (*Lockstep) StopBroadcast()
> 停止广播

***
<span id="struct_Lockstep_IsRunning"></span>

#### func (*Lockstep) IsRunning()  bool
> 是否正在广播

***
<span id="struct_Lockstep_AddCommand"></span>

#### func (*Lockstep) AddCommand(command Command)
> 添加命令到当前帧

***
<span id="struct_Lockstep_AddCommands"></span>

#### func (*Lockstep) AddCommands(commands []Command)
> 添加命令到当前帧

***
<span id="struct_Lockstep_GetCurrentFrame"></span>

#### func (*Lockstep) GetCurrentFrame()  int64
> 获取当前帧

***
<span id="struct_Lockstep_GetClientCurrentFrame"></span>

#### func (*Lockstep) GetClientCurrentFrame(clientId ClientID)  int64
> 获取客户端当前帧

***
<span id="struct_Lockstep_GetFrameLimit"></span>

#### func (*Lockstep) GetFrameLimit()  int64
> 获取帧上限
>   - 未设置时将返回0

***
<span id="struct_Lockstep_GetCurrentCommands"></span>

#### func (*Lockstep) GetCurrentCommands()  []Command
> 获取当前帧还未结束时的所有指令

***
<span id="struct_Lockstep_RegLockstepStoppedEvent"></span>

#### func (*Lockstep) RegLockstepStoppedEvent(handle StoppedEventHandle[ClientID, Command])
> 当广播停止时将触发被注册的事件处理函数

***
<span id="struct_Lockstep_OnLockstepStoppedEvent"></span>

#### func (*Lockstep) OnLockstepStoppedEvent()

***
<span id="struct_Option"></span>
### Option `STRUCT`

```go
type Option[ClientID comparable, Command any] func(lockstep *Lockstep[ClientID, Command])
```
