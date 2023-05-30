package game

// FSM 有限状态机接口定义
//   - 有限状态机（Finite State Machine，简称FSM），它将系统看做一系列状态，而每个状态都有其对应的行为和转换条件
//   - 在游戏中，可以将各种游戏对象和游戏系统抽象为状态和转移，使用有限状态机来描述它们之间的关系，来实现游戏的逻辑和行为
//
// 使用场景
//   - 角色行为控制：游戏角色的行为通常会根据不同状态来变化，例如走、跑、跳、攻击等。通过使用有限状态机来描述角色的状态和状态之间的转移，可以实现角色行为的流畅切换和控制
//   - NPC行为控制：游戏中的 NPC 通常需要有一定的 AI 行为，例如巡逻、追击、攻击等。通过使用有限状态机来描述 NPC 的行为状态和转移条件，可以实现 NPC 行为的智能控制
//   - 游戏流程控制：游戏的各个流程可以看做状态，例如登录、主界面、游戏场景等。通过使用有限状态机来描述游戏的流程状态和转移条件，可以实现游戏流程的控制和管理
//
// 在服务端中使用 FSM 可以使游戏状态管理更加清晰和简单，清晰各状态职责
//   - 内置实现：builtin.FSM
//   - 构建函数：builtin.NewFSM
type FSM[State comparable, Data comparable] interface {
	// Update 执行当前状态行为
	Update()
	// Register 注册一个状态到当前状态机中
	Register(state FSMState[State, Data])
	// Unregister 取消一个状态的注册
	Unregister(state State)
	// HasState 检查是否某状态已注册
	HasState(state State) bool
	// Change 改变当前状态为输入的状态
	Change(state State)
}
