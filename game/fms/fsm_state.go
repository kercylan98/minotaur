package fms

type (
	FSMStateEnterHandle[Data any]  func(data Data)
	FSMStateUpdateHandle[Data any] func(data Data)
	FSMStateExitHandle[Data any]   func(data Data)
)

func NewFSMState[State comparable, Data any](state State, enter FSMStateEnterHandle[Data], update FSMStateUpdateHandle[Data], exit FSMStateExitHandle[Data]) *FSMState[State, Data] {
	return &FSMState[State, Data]{
		enter:  enter,
		update: update,
		exit:   exit,
	}
}

type FSMState[State comparable, Data any] struct {
	state  State
	enter  FSMStateEnterHandle[Data]
	update FSMStateUpdateHandle[Data]
	exit   FSMStateExitHandle[Data]
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
