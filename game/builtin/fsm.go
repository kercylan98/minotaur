package builtin

import (
	"fmt"
	"github.com/kercylan98/minotaur/game"
)

func NewFSM[State comparable, Data any](data Data) *FSM[State, Data] {
	return &FSM[State, Data]{
		states: map[State]game.FSMState[State, Data]{},
		data:   data,
	}
}

type FSM[State comparable, Data any] struct {
	current State
	data    Data
	states  map[State]game.FSMState[State, Data]
}

func (slf *FSM[State, Data]) Update() {
	state := slf.states[slf.current]
	state.Update(slf.data)
}

func (slf *FSM[State, Data]) Register(state game.FSMState[State, Data]) {
	slf.states[state.GetState()] = state
}

func (slf *FSM[State, Data]) Unregister(state State) {
	delete(slf.states, state)
}

func (slf *FSM[State, Data]) HasState(state State) bool {
	_, has := slf.states[state]
	return has
}

func (slf *FSM[State, Data]) Change(state State) {
	current := slf.states[slf.current]
	current.Exit(slf.data)

	next := slf.states[state]
	if next == nil {
		panic(fmt.Errorf("FSM object is attempting to switch to an invalid / undefined state: %v", state))
	}

	next.Enter(slf.data)
}
