package builtin

import "github.com/kercylan98/minotaur/game"

func NewFSMState[State comparable, Data any](state State, enter game.FSMStateEnterHandle[Data], update game.FSMStateUpdateHandle[Data], exit game.FSMStateExitHandle[Data]) *FSMState[State, Data] {
	return &FSMState[State, Data]{
		enter:  enter,
		update: update,
		exit:   exit,
	}
}

type FSMState[State comparable, Data any] struct {
	state  State
	enter  game.FSMStateEnterHandle[Data]
	update game.FSMStateUpdateHandle[Data]
	exit   game.FSMStateExitHandle[Data]
}

func (slf *FSMState[State, Data]) GetState() State {
	return slf.state
}

func (slf *FSMState[State, Data]) Enter(data Data) {
	slf.enter(data)
}

func (slf *FSMState[State, Data]) Update(data Data) {
	slf.update(data)
}

func (slf *FSMState[State, Data]) Exit(data Data) {
	slf.exit(data)
}
