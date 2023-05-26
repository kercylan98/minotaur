package game

// FSM 有限状态机
type FSM[State comparable, Data comparable] interface {
	Update()
	Register(state FSMState[State, Data])
	Unregister(state State)
	HasState(state State) bool
	Change(state State)
}
