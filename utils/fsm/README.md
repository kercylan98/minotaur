# Fsm



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/fsm)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewFSM](#NewFSM)|创建一个新的状态机
|[WithEnterBeforeEvent](#WithEnterBeforeEvent)|设置状态进入前的回调
|[WithEnterAfterEvent](#WithEnterAfterEvent)|设置状态进入后的回调
|[WithUpdateEvent](#WithUpdateEvent)|设置状态内刷新的回调
|[WithExitBeforeEvent](#WithExitBeforeEvent)|设置状态退出前的回调
|[WithExitAfterEvent](#WithExitAfterEvent)|设置状态退出后的回调


> 结构体定义

|结构体|描述
|:--|:--
|[FSM](#fsm)|状态机
|[Option](#option)|暂无描述...

</details>


#### func NewFSM(data Data)  *FSM[State, Data]
<span id="NewFSM"></span>
> 创建一个新的状态机
***
#### func WithEnterBeforeEvent(fn func (state *FSM[State, Data]))  Option[State, Data]
<span id="WithEnterBeforeEvent"></span>
> 设置状态进入前的回调
>   - 在首次设置状态时，状态机本身的当前状态为零值状态
***
#### func WithEnterAfterEvent(fn func (state *FSM[State, Data]))  Option[State, Data]
<span id="WithEnterAfterEvent"></span>
> 设置状态进入后的回调
***
#### func WithUpdateEvent(fn func (state *FSM[State, Data]))  Option[State, Data]
<span id="WithUpdateEvent"></span>
> 设置状态内刷新的回调
***
#### func WithExitBeforeEvent(fn func (state *FSM[State, Data]))  Option[State, Data]
<span id="WithExitBeforeEvent"></span>
> 设置状态退出前的回调
>   - 该阶段状态机的状态为退出前的状态，而非新的状态
***
#### func WithExitAfterEvent(fn func (state *FSM[State, Data]))  Option[State, Data]
<span id="WithExitAfterEvent"></span>
> 设置状态退出后的回调
>   - 该阶段状态机的状态为新的状态，而非退出前的状态
***
### FSM
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
### Option

```go
type Option[State comparable, Data any] struct{}
```
