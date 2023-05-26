package game

type (
	FSMStateEnterHandle[Data any]  func(data Data)
	FSMStateUpdateHandle[Data any] func(data Data)
	FSMStateExitHandle[Data any]   func(data Data)
)

// FSMState 有限状态机状态
type FSMState[State comparable, Data any] interface {
	GetState() State
	Enter(data Data)
	Update(data Data)
	Exit(data Data)
}
