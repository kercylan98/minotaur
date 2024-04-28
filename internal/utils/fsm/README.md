# Fsm

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
|[NewFSM](#NewFSM)|创建一个新的状态机
|[WithEnterBeforeEvent](#WithEnterBeforeEvent)|设置状态进入前的回调
|[WithEnterAfterEvent](#WithEnterAfterEvent)|设置状态进入后的回调
|[WithUpdateEvent](#WithUpdateEvent)|设置状态内刷新的回调
|[WithExitBeforeEvent](#WithExitBeforeEvent)|设置状态退出前的回调
|[WithExitAfterEvent](#WithExitAfterEvent)|设置状态退出后的回调


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[FSM](#struct_FSM)|状态机
|`STRUCT`|[Option](#struct_Option)|暂无描述...

</details>


***
## 详情信息
#### func NewFSM\[State comparable, Data any\](data Data) *FSM[State, Data]
<span id="NewFSM"></span>
> 创建一个新的状态机

***
#### func WithEnterBeforeEvent\[State comparable, Data any\](fn func (state *FSM[State, Data])) Option[State, Data]
<span id="WithEnterBeforeEvent"></span>
> 设置状态进入前的回调
>   - 在首次设置状态时，状态机本身的当前状态为零值状态

***
#### func WithEnterAfterEvent\[State comparable, Data any\](fn func (state *FSM[State, Data])) Option[State, Data]
<span id="WithEnterAfterEvent"></span>
> 设置状态进入后的回调

***
#### func WithUpdateEvent\[State comparable, Data any\](fn func (state *FSM[State, Data])) Option[State, Data]
<span id="WithUpdateEvent"></span>
> 设置状态内刷新的回调

***
#### func WithExitBeforeEvent\[State comparable, Data any\](fn func (state *FSM[State, Data])) Option[State, Data]
<span id="WithExitBeforeEvent"></span>
> 设置状态退出前的回调
>   - 该阶段状态机的状态为退出前的状态，而非新的状态

***
#### func WithExitAfterEvent\[State comparable, Data any\](fn func (state *FSM[State, Data])) Option[State, Data]
<span id="WithExitAfterEvent"></span>
> 设置状态退出后的回调
>   - 该阶段状态机的状态为新的状态，而非退出前的状态

***
<span id="struct_FSM"></span>
### FSM `STRUCT`
状态机
```go
type FSM[State comparable, Data any] struct {
	prev                    *State
	current                 *State
	data                    Data
	states                  map[State]struct{}
	enterBeforeEventHandles map[State][]func(state *FSM[State, Data])
	enterAfterEventHandles  map[State][]func(state *FSM[State, Data])
	updateEventHandles      map[State][]func(state *FSM[State, Data])
	exitBeforeEventHandles  map[State][]func(state *FSM[State, Data])
	exitAfterEventHandles   map[State][]func(state *FSM[State, Data])
}
```
<span id="struct_FSM_Update"></span>

#### func (*FSM) Update()
> 触发当前状态

***
<span id="struct_FSM_Register"></span>

#### func (*FSM) Register(state State, options ...Option[State, Data])
> 注册状态

***
<span id="struct_FSM_Unregister"></span>

#### func (*FSM) Unregister(state State)
> 反注册状态

***
<span id="struct_FSM_HasState"></span>

#### func (*FSM) HasState(state State)  bool
> 检查状态机是否存在特定状态

***
<span id="struct_FSM_Change"></span>

#### func (*FSM) Change(state State)
> 改变状态机状态到新的状态

***
<span id="struct_FSM_Current"></span>

#### func (*FSM) Current() (state State)
> 获取当前状态

***
<span id="struct_FSM_GetData"></span>

#### func (*FSM) GetData()  Data
> 获取状态机数据

***
<span id="struct_FSM_IsZero"></span>

#### func (*FSM) IsZero()  bool
> 检查状态机是否无状态

***
<span id="struct_FSM_PrevIsZero"></span>

#### func (*FSM) PrevIsZero()  bool
> 检查状态机上一个状态是否无状态

***
<span id="struct_Option"></span>
### Option `STRUCT`

```go
type Option[State comparable, Data any] func(fsm *FSM[State, Data], state State)
```
