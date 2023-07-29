package fsm

import (
	"fmt"
)

// NewFSM 创建一个新的状态机
func NewFSM[State comparable, Data any](data Data) *FSM[State, Data] {
	return &FSM[State, Data]{
		states: map[State]*State[State, Data]{},
		data:   data,
	}
}

// FSM 状态机
type FSM[State comparable, Data any] struct {
	current State
	data    Data
	states  map[State]*State[State, Data]
}

// Update 触发当前状态
func (slf *FSM[State, Data]) Update() {
	state := slf.states[slf.current]
	state.Update(slf.data)
}

// Register 注册状态
func (slf *FSM[State, Data]) Register(state *State[State, Data]) {
	slf.states[state.GetState()] = state
}

// Unregister 反注册状态
func (slf *FSM[State, Data]) Unregister(state State) {
	delete(slf.states, state)
}

// HasState 检查状态机是否存在特定状态
func (slf *FSM[State, Data]) HasState(state State) bool {
	_, has := slf.states[state]
	return has
}

// Change 改变状态机状态到新的状态
func (slf *FSM[State, Data]) Change(state State) {
	current := slf.states[slf.current]
	current.Exit(slf.data)

	next := slf.states[state]
	if next == nil {
		panic(fmt.Errorf("FSM object is attempting to switch to an invalid / undefined state: %v", state))
	}

	next.Enter(slf.data)
}
