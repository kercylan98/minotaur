package times

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/generic"
	"strings"
	"time"
)

// NewStateLine 创建一个从左向右由早到晚的状态时间线
func NewStateLine[State generic.Basic](zero State) *StateLine[State] {
	return &StateLine[State]{
		states:  []State{zero},
		points:  []time.Time{{}},
		trigger: [][]func(){{}},
	}
}

// StateLine 状态时间线
type StateLine[State generic.Basic] struct {
	states  []State     // 每个时间点对应的状态
	points  []time.Time // 每个时间点
	trigger [][]func()  // 每个时间点对应的触发器
}

// Check 根据状态顺序检查时间线是否合法
//   - missingAllowed: 是否允许状态缺失，如果为 true，则状态可以不连续，如果为 false，则状态必须连续
//
// 状态不连续表示时间线中存在状态缺失，例如：
//   - 状态为 [1, 2, 3, 4, 5] 的时间线，如果 missingAllowed 为 true，则状态为 [1, 3, 5] 也是合法的
//   - 状态为 [1, 2, 3, 4, 5] 的时间线，如果 missingAllowed 为 false，则状态为 [1, 3, 5] 是不合法的
func (slf *StateLine[State]) Check(missingAllowed bool, states ...State) bool {
	var indexStored int
	var indexInput int

	for indexStored < len(slf.states) && indexInput < len(states) {
		if slf.states[indexStored] == states[indexInput] {
			indexStored++
			indexInput++
		} else if missingAllowed {
			indexInput++
		} else {
			return false
		}
	}

	//如果存储序列还有剩余, 而输入序列已经遍历完
	if indexStored != len(slf.states) && indexInput == len(states) {
		return false
	}

	// 如果输入序列还有剩余, 而存储序列已经遍历完
	if indexStored == len(slf.states) && indexInput != len(states) && !missingAllowed {
		return false
	}

	return true
}

// GetMissingStates 获取缺失的状态
func (slf *StateLine[State]) GetMissingStates(states ...State) []State {
	var missing = make([]State, 0, len(states))
	for _, state := range states {
		if !collection.InComparableSlice(slf.states, state) {
			missing = append(missing, state)
		}
	}
	return missing
}

// HasState 检查时间线中是否包含指定状态
func (slf *StateLine[State]) HasState(state State) bool {
	return collection.InComparableSlice(slf.states, state)
}

// String 获取时间线的字符串表示
func (slf *StateLine[State]) String() string {
	var parts []string
	for i := 0; i < len(slf.states); i++ {
		parts = append(parts, fmt.Sprintf("[%v] %v", slf.states[i], slf.points[i]))
	}
	return strings.Join(parts, " > ")
}

// AddState 添加一个状态到时间线中，状态不能与任一时间点重合，否则将被忽略
//   - onTrigger: 该状态绑定的触发器，该触发器不会被主动执行，需要主动获取触发器执行
func (slf *StateLine[State]) AddState(state State, t time.Time, onTrigger ...func()) *StateLine[State] {
	if collection.InComparableSlice(slf.states, state) {
		return slf
	}
	// 将 t 按照从左到右由早到晚的顺序插入到 points 中
	for i := 0; i < len(slf.points); i++ {
		if slf.points[i].After(t) {
			slf.points = append(slf.points[:i], append([]time.Time{t}, slf.points[i:]...)...)
			slf.states = append(slf.states[:i], append([]State{state}, slf.states[i:]...)...)
			slf.trigger = append(slf.trigger[:i], append([][]func(){onTrigger}, slf.trigger[i:]...)...)
			return slf
		}
	}
	slf.points = append(slf.points, t)
	slf.states = append(slf.states, state)
	slf.trigger = append(slf.trigger, onTrigger)
	return slf
}

// GetTimeByState 获取指定状态的时间点
func (slf *StateLine[State]) GetTimeByState(state State) time.Time {
	for i := 0; i < len(slf.states); i++ {
		if slf.states[i] == state {
			return slf.points[i]
		}
	}
	return time.Time{}
}

// GetNextTimeByState 获取指定状态的下一个时间点
func (slf *StateLine[State]) GetNextTimeByState(state State) time.Time {
	for i := 0; i < len(slf.states); i++ {
		if slf.states[i] == state && i+1 < len(slf.points) {
			return slf.points[i+1]
		}
	}
	return slf.points[0]
}

// GetLastState 获取最后一个状态
func (slf *StateLine[State]) GetLastState() State {
	return slf.states[len(slf.states)-1]
}

// GetPrevTimeByState 获取指定状态的上一个时间点
func (slf *StateLine[State]) GetPrevTimeByState(state State) time.Time {
	for i := len(slf.states) - 1; i >= 0; i-- {
		if slf.states[i] == state && i > 0 {
			return slf.points[i-1]
		}
	}
	return time.Time{}
}

// GetIndexByState 获取指定状态的索引
func (slf *StateLine[State]) GetIndexByState(state State) int {
	for i := 0; i < len(slf.states); i++ {
		if slf.states[i] == state {
			return i
		}
	}
	return -1
}

// GetStateByTime 获取指定时间点的状态
func (slf *StateLine[State]) GetStateByTime(t time.Time) State {
	for i := len(slf.points) - 1; i >= 0; i-- {
		point := slf.points[i]
		if point.Before(t) || point.Equal(t) {
			return slf.states[i]
		}
	}
	return slf.states[len(slf.points)-1]
}

// GetTimeByIndex 获取指定索引的时间点
func (slf *StateLine[State]) GetTimeByIndex(index int) time.Time {
	return slf.points[index]
}

// Move 时间线整体移动
func (slf *StateLine[State]) Move(d time.Duration) *StateLine[State] {
	for i := 0; i < len(slf.points); i++ {
		slf.points[i] = slf.points[i].Add(d)
	}
	return slf
}

// GetNextStateTimeByIndex 获取指定索引的下一个时间点
func (slf *StateLine[State]) GetNextStateTimeByIndex(index int) time.Time {
	return slf.points[index+1]
}

// GetPrevStateTimeByIndex 获取指定索引的上一个时间点
func (slf *StateLine[State]) GetPrevStateTimeByIndex(index int) time.Time {
	return slf.points[index-1]
}

// GetStateIndexByTime 获取指定时间点的索引
func (slf *StateLine[State]) GetStateIndexByTime(t time.Time) int {
	for i := len(slf.points) - 1; i >= 0; i-- {
		var point = slf.points[i]
		if point.Before(t) || point.Equal(t) {
			return i
		}
	}
	return -1
}

// GetStateCount 获取状态数量
func (slf *StateLine[State]) GetStateCount() int {
	return len(slf.states)
}

// GetStateByIndex 获取指定索引的状态
func (slf *StateLine[State]) GetStateByIndex(index int) State {
	return slf.states[index]
}

// GetTriggerByTime 获取指定时间点的触发器
func (slf *StateLine[State]) GetTriggerByTime(t time.Time) []func() {
	for i := len(slf.points) - 1; i >= 0; i-- {
		var point = slf.points[i]
		if point.Before(t) || point.Equal(t) {
			return slf.trigger[i]
		}
	}
	return nil
}

// GetTriggerByIndex 获取指定索引的触发器
func (slf *StateLine[State]) GetTriggerByIndex(index int) []func() {
	return slf.trigger[index]
}

// GetTriggerByState 获取指定状态的触发器
func (slf *StateLine[State]) GetTriggerByState(state State) []func() {
	for i := 0; i < len(slf.states); i++ {
		if slf.states[i] == state {
			return slf.trigger[i]
		}
	}
	return nil
}

// AddTriggerToState 给指定状态添加触发器
func (slf *StateLine[State]) AddTriggerToState(state State, onTrigger ...func()) *StateLine[State] {
	for i := 0; i < len(slf.states); i++ {
		if slf.states[i] == state {
			slf.trigger[i] = append(slf.trigger[i], onTrigger...)
			return slf
		}
	}
	return slf
}

// Range 按照时间顺序遍历时间线
func (slf *StateLine[State]) Range(handler func(index int, state State, t time.Time) bool) {
	for i := 0; i < len(slf.points); i++ {
		if !handler(i, slf.states[i], slf.points[i]) {
			return
		}
	}
}

// RangeReverse 按照时间逆序遍历时间线
func (slf *StateLine[State]) RangeReverse(handler func(index int, state State, t time.Time) bool) {
	for i := len(slf.points) - 1; i >= 0; i-- {
		if !handler(i, slf.states[i], slf.points[i]) {
			return
		}
	}
}
