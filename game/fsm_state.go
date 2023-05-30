package game

type (
	FSMStateEnterHandle[Data any]  func(data Data)
	FSMStateUpdateHandle[Data any] func(data Data)
	FSMStateExitHandle[Data any]   func(data Data)
)

// FSMState 有限状态机状态接口定义
//   - 描述了 FSM 中的状态行为
//
// 需要一个可比较的状态类型及相应的数据类型，数据类型在无需使用时支持传入 nil
//   - 内置实现：builtin.FSMState
//   - 构建函数：builtin.NewFSMState
type FSMState[State comparable, Data any] interface {
	// GetState 获取当前状态
	GetState() State
	// Enter 状态开始行为
	Enter(data Data)
	// Update 当前状态行为
	Update(data Data)
	// Exit 状态退出行为
	Exit(data Data)
}
