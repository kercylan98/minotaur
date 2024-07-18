package supervision

import "time"

func NewAccidentState() *AccidentState {
	return &AccidentState{}
}

// AccidentState 事故状态
type AccidentState struct {
	accidentTimes []time.Time
}

// Solved 事故已解决
func (as *AccidentState) Solved() {
	as.accidentTimes = nil
}

// Record 记录一次事故发生
func (as *AccidentState) Record() {
	as.accidentTimes = append(as.accidentTimes, time.Now())
}

// AccidentCount 获取事故数量
func (as *AccidentState) AccidentCount() int {
	return len(as.accidentTimes)
}

// LastTime 获取最后一次产生事故的时间
func (as *AccidentState) LastTime() time.Time {
	if len(as.accidentTimes) == 0 {
		return time.Time{}
	}
	return as.accidentTimes[len(as.accidentTimes)-1]
}
