package fsm

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
)

// NewFSM 创建一个新的状态机
func NewFSM[State comparable, Data any](data Data) *FSM[State, Data] {
	return &FSM[State, Data]{
		states: map[State]struct{}{},
		data:   data,
	}
}

// FSM 状态机
type FSM[State comparable, Data any] struct {
	prev    *State
	current *State
	data    Data
	states  map[State]struct{}

	enterBeforeEventHandles map[State][]func(state *FSM[State, Data])
	enterAfterEventHandles  map[State][]func(state *FSM[State, Data])
	updateEventHandles      map[State][]func(state *FSM[State, Data])
	exitBeforeEventHandles  map[State][]func(state *FSM[State, Data])
	exitAfterEventHandles   map[State][]func(state *FSM[State, Data])
}

// Update 触发当前状态
func (slf *FSM[State, Data]) Update() {
	if slf.current == nil {
		return
	}
	for _, event := range slf.updateEventHandles[*slf.current] {
		event(slf)
	}
}

// Register 注册状态
func (slf *FSM[State, Data]) Register(state State, options ...Option[State, Data]) {
	slf.states[state] = struct{}{}
	for _, option := range options {
		option(slf, state)
	}
}

// Unregister 反注册状态
func (slf *FSM[State, Data]) Unregister(state State) {
	if !slf.HasState(state) {
		return
	}
	delete(slf.states, state)
	delete(slf.enterBeforeEventHandles, state)
	delete(slf.enterAfterEventHandles, state)
	delete(slf.updateEventHandles, state)
	delete(slf.exitBeforeEventHandles, state)
	delete(slf.exitAfterEventHandles, state)
}

// HasState 检查状态机是否存在特定状态
func (slf *FSM[State, Data]) HasState(state State) bool {
	return collection.FindInMapKey(slf.states, state)
}

// Change 改变状态机状态到新的状态
func (slf *FSM[State, Data]) Change(state State) {
	if !slf.IsZero() {
		for _, event := range slf.exitBeforeEventHandles[*slf.current] {
			event(slf)
		}
	}

	slf.prev = slf.current
	slf.current = &state

	if !slf.PrevIsZero() {
		for _, event := range slf.exitAfterEventHandles[*slf.prev] {
			event(slf)
		}
	}

	if !slf.HasState(state) {
		panic(fmt.Errorf("FSM object is attempting to switch to an invalid / undefined state: %v", state))
	}

	for _, event := range slf.enterBeforeEventHandles[*slf.current] {
		event(slf)
	}

	for _, event := range slf.enterAfterEventHandles[*slf.current] {
		event(slf)
	}
}

// Current 获取当前状态
func (slf *FSM[State, Data]) Current() (state State) {
	if slf.current == nil {
		return
	}
	return *slf.current
}

// GetData 获取状态机数据
func (slf *FSM[State, Data]) GetData() Data {
	return slf.data
}

// IsZero 检查状态机是否无状态
func (slf *FSM[State, Data]) IsZero() bool {
	return slf.current == nil
}

// PrevIsZero 检查状态机上一个状态是否无状态
func (slf *FSM[State, Data]) PrevIsZero() bool {
	return slf.prev == nil
}
