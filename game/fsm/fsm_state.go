package fsm

type (
	StateEnterHandle[Data any]  func(data Data)
	StateUpdateHandle[Data any] func(data Data)
	StateExitHandle[Data any]   func(data Data)
)

func NewFSMState[S comparable, Data any](state S, enter StateEnterHandle[Data], update StateUpdateHandle[Data], exit StateExitHandle[Data]) *State[S, Data] {
	return &State[S, Data]{
		enter:  enter,
		update: update,
		exit:   exit,
	}
}

type State[S comparable, Data any] struct {
	state  S
	enter  StateEnterHandle[Data]
	update StateUpdateHandle[Data]
	exit   StateExitHandle[Data]
}

func (slf *State[S, Data]) GetState() S {
	return slf.state
}

func (slf *State[S, Data]) Enter(data Data) {
	slf.enter(data)
}

func (slf *State[S, Data]) Update(data Data) {
	slf.update(data)
}

func (slf *State[S, Data]) Exit(data Data) {
	slf.exit(data)
}
