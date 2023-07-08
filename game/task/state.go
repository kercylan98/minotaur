package task

const (
	StateDoing  State = iota // 进行中
	StateDone                // 已完成
	StateReward              // 已领取
)

// State 任务状态
type State byte
