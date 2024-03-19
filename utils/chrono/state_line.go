package chrono

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

// StateLine 表示一个状态时间线，它记录了一系列时间点及其对应的状态和触发器。
// 在时间线中，每个时间点都与一个状态和一组触发器相关联，可以通过时间点查找状态，并触发与之相关联的触发器。
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
func (s *StateLine[State]) Check(missingAllowed bool, states ...State) bool {
	var indexStored int
	var indexInput int

	for indexStored < len(s.states) && indexInput < len(states) {
		if s.states[indexStored] == states[indexInput] {
			indexStored++
			indexInput++
		} else if missingAllowed {
			indexInput++
		} else {
			return false
		}
	}

	//如果存储序列还有剩余, 而输入序列已经遍历完
	if indexStored != len(s.states) && indexInput == len(states) {
		return false
	}

	// 如果输入序列还有剩余, 而存储序列已经遍历完
	if indexStored == len(s.states) && indexInput != len(states) && !missingAllowed {
		return false
	}

	return true
}

// GetMissingStates 获取缺失的状态
func (s *StateLine[State]) GetMissingStates(states ...State) []State {
	var missing = make([]State, 0, len(states))
	for _, state := range states {
		if !collection.InComparableSlice(s.states, state) {
			missing = append(missing, state)
		}
	}
	return missing
}

// HasState 检查时间线中是否包含指定状态
func (s *StateLine[State]) HasState(state State) bool {
	return collection.InComparableSlice(s.states, state)
}

// String 获取时间线的字符串表示
func (s *StateLine[State]) String() string {
	var parts []string
	for i := 0; i < len(s.states); i++ {
		parts = append(parts, fmt.Sprintf("[%v] %v", s.states[i], s.points[i]))
	}
	return strings.Join(parts, " > ")
}

// AddState 添加一个状态到时间线中，状态不能与任一时间点重合，否则将被忽略
//   - onTrigger: 该状态绑定的触发器，该触发器不会被主动执行，需要主动获取触发器执行
func (s *StateLine[State]) AddState(state State, t time.Time, onTrigger ...func()) *StateLine[State] {
	if collection.InComparableSlice(s.states, state) {
		return s
	}
	// 将 t 按照从左到右由早到晚的顺序插入到 points 中
	for i := 0; i < len(s.points); i++ {
		if s.points[i].After(t) {
			s.points = append(s.points[:i], append([]time.Time{t}, s.points[i:]...)...)
			s.states = append(s.states[:i], append([]State{state}, s.states[i:]...)...)
			s.trigger = append(s.trigger[:i], append([][]func(){onTrigger}, s.trigger[i:]...)...)
			return s
		}
	}
	s.points = append(s.points, t)
	s.states = append(s.states, state)
	s.trigger = append(s.trigger, onTrigger)
	return s
}

// GetTimeByState 获取指定状态的时间点
func (s *StateLine[State]) GetTimeByState(state State) time.Time {
	for i := 0; i < len(s.states); i++ {
		if s.states[i] == state {
			return s.points[i]
		}
	}
	return time.Time{}
}

// GetNextTimeByState 获取指定状态的下一个时间点
func (s *StateLine[State]) GetNextTimeByState(state State) time.Time {
	for i := 0; i < len(s.states); i++ {
		if s.states[i] == state && i+1 < len(s.points) {
			return s.points[i+1]
		}
	}
	return s.points[0]
}

// GetLastState 获取最后一个状态
func (s *StateLine[State]) GetLastState() State {
	return s.states[len(s.states)-1]
}

// GetPrevTimeByState 获取指定状态的上一个时间点
func (s *StateLine[State]) GetPrevTimeByState(state State) time.Time {
	for i := len(s.states) - 1; i >= 0; i-- {
		if s.states[i] == state && i > 0 {
			return s.points[i-1]
		}
	}
	return time.Time{}
}

// GetIndexByState 获取指定状态的索引
func (s *StateLine[State]) GetIndexByState(state State) int {
	for i := 0; i < len(s.states); i++ {
		if s.states[i] == state {
			return i
		}
	}
	return -1
}

// GetStateByTime 获取指定时间点的状态
func (s *StateLine[State]) GetStateByTime(t time.Time) State {
	for i := len(s.points) - 1; i >= 0; i-- {
		point := s.points[i]
		if point.Before(t) || point.Equal(t) {
			return s.states[i]
		}
	}
	return s.states[len(s.points)-1]
}

// GetTimeByIndex 获取指定索引的时间点
func (s *StateLine[State]) GetTimeByIndex(index int) time.Time {
	return s.points[index]
}

// Move 时间线整体移动
func (s *StateLine[State]) Move(d time.Duration) *StateLine[State] {
	for i := 0; i < len(s.points); i++ {
		s.points[i] = s.points[i].Add(d)
	}
	return s
}

// GetNextStateTimeByIndex 获取指定索引的下一个时间点
func (s *StateLine[State]) GetNextStateTimeByIndex(index int) time.Time {
	return s.points[index+1]
}

// GetPrevStateTimeByIndex 获取指定索引的上一个时间点
func (s *StateLine[State]) GetPrevStateTimeByIndex(index int) time.Time {
	return s.points[index-1]
}

// GetStateIndexByTime 获取指定时间点的索引
func (s *StateLine[State]) GetStateIndexByTime(t time.Time) int {
	for i := len(s.points) - 1; i >= 0; i-- {
		var point = s.points[i]
		if point.Before(t) || point.Equal(t) {
			return i
		}
	}
	return -1
}

// GetStateCount 获取状态数量
func (s *StateLine[State]) GetStateCount() int {
	return len(s.states)
}

// GetStateByIndex 获取指定索引的状态
func (s *StateLine[State]) GetStateByIndex(index int) State {
	return s.states[index]
}

// GetTriggerByTime 获取指定时间点的触发器
func (s *StateLine[State]) GetTriggerByTime(t time.Time) []func() {
	for i := len(s.points) - 1; i >= 0; i-- {
		var point = s.points[i]
		if point.Before(t) || point.Equal(t) {
			return s.trigger[i]
		}
	}
	return nil
}

// GetTriggerByIndex 获取指定索引的触发器
func (s *StateLine[State]) GetTriggerByIndex(index int) []func() {
	return s.trigger[index]
}

// GetTriggerByState 获取指定状态的触发器
func (s *StateLine[State]) GetTriggerByState(state State) []func() {
	for i := 0; i < len(s.states); i++ {
		if s.states[i] == state {
			return s.trigger[i]
		}
	}
	return nil
}

// AddTriggerToState 给指定状态添加触发器
func (s *StateLine[State]) AddTriggerToState(state State, onTrigger ...func()) *StateLine[State] {
	for i := 0; i < len(s.states); i++ {
		if s.states[i] == state {
			s.trigger[i] = append(s.trigger[i], onTrigger...)
			return s
		}
	}
	return s
}

// Iterate 按照时间顺序遍历时间线
func (s *StateLine[State]) Iterate(handler func(index int, state State, t time.Time) bool) {
	for i := 0; i < len(s.points); i++ {
		if !handler(i, s.states[i], s.points[i]) {
			return
		}
	}
}

// IterateReverse 按照时间逆序遍历时间线
func (s *StateLine[State]) IterateReverse(handler func(index int, state State, t time.Time) bool) {
	for i := len(s.points) - 1; i >= 0; i-- {
		if !handler(i, s.states[i], s.points[i]) {
			return
		}
	}
}
