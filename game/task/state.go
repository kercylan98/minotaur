package task

const (
	StateAccept State = iota // 已接受
	StateFinish              // 已完成
	StateReward              // 已领取
	StateFail                // 已失败
)

type State uint16
