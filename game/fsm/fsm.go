package fsm

import (
	"fmt"
)

// NewFSM 创建一个新的状态机
func NewFSM[S comparable, Data any](data Data) *FSM[S, Data] {
	return &FSM[S, Data]{
		states: map[S]*State[S, Data]{},
		data:   data,
	}
}

// FSM 状态机
type FSM[S comparable, Data any] struct {
	current S
	data    Data
	states  map[S]*State[S, Data]
}

// Update 触发当前状态
func (slf *FSM[S, Data]) Update() {
	state := slf.states[slf.current]
	state.Update(slf.data)
}

// Register 注册状态
func (slf *FSM[S, Data]) Register(state *State[S, Data]) {
	slf.states[state.GetState()] = state
}

// Unregister 反注册状态
func (slf *FSM[S, Data]) Unregister(state S) {
	delete(slf.states, state)
}

// HasState 检查状态机是否存在特定状态
func (slf *FSM[S, Data]) HasState(state S) bool {
	_, has := slf.states[state]
	return has
}

// Change 改变状态机状态到新的状态
func (slf *FSM[S, Data]) Change(state S) {
	current := slf.states[slf.current]
	current.Exit(slf.data)

	next := slf.states[state]
	if next == nil {
		panic(fmt.Errorf("FSM object is attempting to switch to an invalid / undefined state: %v", state))
	}

	next.Enter(slf.data)
}
