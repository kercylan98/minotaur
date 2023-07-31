package fsm

type Option[State comparable, Data any] func(fsm *FSM[State, Data], state State)

// WithEnterBeforeEvent 设置状态进入前的回调
//   - 在首次设置状态时，状态机本身的当前状态为零值状态
func WithEnterBeforeEvent[State comparable, Data any](fn func(state *FSM[State, Data])) Option[State, Data] {
	return func(fsm *FSM[State, Data], state State) {
		if fsm.enterBeforeEventHandles == nil {
			fsm.enterBeforeEventHandles = map[State][]func(state *FSM[State, Data]){}
		}
		fsm.enterBeforeEventHandles[state] = append(fsm.enterBeforeEventHandles[state], fn)
	}
}

// WithEnterAfterEvent 设置状态进入后的回调
func WithEnterAfterEvent[State comparable, Data any](fn func(state *FSM[State, Data])) Option[State, Data] {
	return func(fsm *FSM[State, Data], state State) {
		if fsm.enterAfterEventHandles == nil {
			fsm.enterAfterEventHandles = map[State][]func(state *FSM[State, Data]){}
		}
		fsm.enterAfterEventHandles[state] = append(fsm.enterAfterEventHandles[state], fn)
	}
}

// WithUpdateEvent 设置状态内刷新的回调
func WithUpdateEvent[State comparable, Data any](fn func(state *FSM[State, Data])) Option[State, Data] {
	return func(fsm *FSM[State, Data], state State) {
		if fsm.updateEventHandles == nil {
			fsm.updateEventHandles = map[State][]func(state *FSM[State, Data]){}
		}
		fsm.updateEventHandles[state] = append(fsm.updateEventHandles[state], fn)
	}
}

// WithExitBeforeEvent 设置状态退出前的回调
//   - 该阶段状态机的状态为退出前的状态，而非新的状态
func WithExitBeforeEvent[State comparable, Data any](fn func(state *FSM[State, Data])) Option[State, Data] {
	return func(fsm *FSM[State, Data], state State) {
		if fsm.exitBeforeEventHandles == nil {
			fsm.exitBeforeEventHandles = map[State][]func(state *FSM[State, Data]){}
		}
		fsm.exitBeforeEventHandles[state] = append(fsm.exitBeforeEventHandles[state], fn)
	}
}

// WithExitAfterEvent 设置状态退出后的回调
//   - 该阶段状态机的状态为新的状态，而非退出前的状态
func WithExitAfterEvent[State comparable, Data any](fn func(state *FSM[State, Data])) Option[State, Data] {
	return func(fsm *FSM[State, Data], state State) {
		if fsm.exitAfterEventHandles == nil {
			fsm.exitAfterEventHandles = map[State][]func(state *FSM[State, Data]){}
		}
		fsm.exitAfterEventHandles[state] = append(fsm.exitAfterEventHandles[state], fn)
	}
}
