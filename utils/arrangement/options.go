package arrangement

// Option 编排选项
type Option[ID comparable, AreaInfo any] func(arrangement *Arrangement[ID, AreaInfo])

type (
	ConstraintHandle[ID comparable, AreaInfo any] func(editor *Editor[ID, AreaInfo], area *Area[ID, AreaInfo], item Item[ID], err error) error
	ConflictHandle[ID comparable, AreaInfo any]   func(editor *Editor[ID, AreaInfo], area *Area[ID, AreaInfo], item Item[ID], conflictItems map[ID]Item[ID]) map[ID]Item[ID]
)

// WithRetryThreshold 设置编排时的重试阈值
//   - 当每一轮编排结束任有成员未被编排时，将会进行下一轮编排，直到编排次数达到该阈值
//   - 默认的阈值为 10 次
func WithRetryThreshold[ID comparable, AreaInfo any](threshold int) Option[ID, AreaInfo] {
	return func(arrangement *Arrangement[ID, AreaInfo]) {
		if threshold <= 0 {
			threshold = 10
		}
		arrangement.threshold = threshold
	}
}

// WithConstraintHandle 设置编排时触发约束时的处理函数
//   - 当约束条件触发时，将会调用该函数。如果无法在该函数中处理约束，应该继续返回 err，尝试进行下一层的约束处理
//   - 当该函数的返回值为 nil 时，表示约束已经被处理，将会命中当前的编排区域
//   - 当所有的约束处理函数都无法处理约束时，将会进入下一个编排区域的尝试，如果均无法完成，将会将该成员加入到编排队列的末端，等待下一次编排
//
// 有意思的是，硬性约束应该永远是无解的，而当需要进行一些打破规则的操作时，则可以透过该函数传入的 editor 进行操作
func WithConstraintHandle[ID comparable, AreaInfo any](handle ConstraintHandle[ID, AreaInfo]) Option[ID, AreaInfo] {
	return func(arrangement *Arrangement[ID, AreaInfo]) {
		arrangement.constraintHandles = append(arrangement.constraintHandles, handle)
	}
}

// WithConflictHandle 设置编排时触发冲突时的处理函数
//   - 当冲突条件触发时，将会调用该函数。如果无法在该函数中处理冲突，应该继续返回这一批成员，尝试进行下一层的冲突处理
//   - 当该函数的返回值长度为 0 时，表示冲突已经被处理，将会命中当前的编排区域
//   - 当所有的冲突处理函数都无法处理冲突时，将会进入下一个编排区域的尝试，如果均无法完成，将会将该成员加入到编排队列的末端，等待下一次编排
func WithConflictHandle[ID comparable, AreaInfo any](handle ConflictHandle[ID, AreaInfo]) Option[ID, AreaInfo] {
	return func(arrangement *Arrangement[ID, AreaInfo]) {
		arrangement.conflictHandles = append(arrangement.conflictHandles, handle)
	}
}
