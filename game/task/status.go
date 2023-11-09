package task

const (
	StatusAccept   Status = iota + 1 // 已接受
	StatusFailed                     // 已失败
	StatusComplete                   // 已完成
	StatusReward                     // 已领取奖励
)

var (
	statusFormat = map[Status]string{
		StatusAccept:   "Accept",
		StatusComplete: "Complete",
		StatusReward:   "Reward",
		StatusFailed:   "Failed",
	}
)

type Status byte

func (slf Status) String() string {
	return statusFormat[slf]
}
