package arrangement

// AreaOption 编排区域选项
type AreaOption[ID comparable, AreaInfo any] func(area *Area[ID, AreaInfo])

type (
	AreaConstraintHandle[ID comparable, AreaInfo any] func(area *Area[ID, AreaInfo], item Item[ID]) error
	AreaConflictHandle[ID comparable, AreaInfo any]   func(area *Area[ID, AreaInfo], item Item[ID]) map[ID]Item[ID]
	AreaEvaluateHandle[ID comparable, AreaInfo any]   func(areaInfo AreaInfo, items map[ID]Item[ID]) float64
)

// WithAreaConstraint 设置编排区域的约束条件
//   - 该约束用于判断一个成员是否可以被添加到该编排区域中
//   - 与 WithAreaConflict 不同的是，约束通常用于非成员关系导致的硬性约束，例如：成员的等级过滤、成员的性别等
func WithAreaConstraint[ID comparable, AreaInfo any](constraint AreaConstraintHandle[ID, AreaInfo]) AreaOption[ID, AreaInfo] {
	return func(area *Area[ID, AreaInfo]) {
		area.constraints = append(area.constraints, constraint)
	}
}

// WithAreaConflict 设置编排区域的冲突条件，冲突处理函数需要返回造成冲突的成员列表
//   - 该冲突用于判断一个成员是否可以被添加到该编排区域中
//   - 与 WithAreaConstraint 不同的是，冲突通常用于成员关系导致的软性约束，例如：成员的职业唯一性、成员的种族唯一性等
func WithAreaConflict[ID comparable, AreaInfo any](conflict AreaConflictHandle[ID, AreaInfo]) AreaOption[ID, AreaInfo] {
	return func(area *Area[ID, AreaInfo]) {
		area.conflicts = append(area.conflicts, conflict)
	}
}

// WithAreaEvaluate 设置编排区域的评估函数
//   - 该评估函数将影响成员被编入区域的优先级
func WithAreaEvaluate[ID comparable, AreaInfo any](evaluate AreaEvaluateHandle[ID, AreaInfo]) AreaOption[ID, AreaInfo] {
	return func(area *Area[ID, AreaInfo]) {
		area.evaluate = evaluate
	}
}
