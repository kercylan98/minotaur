package arrangement

// AreaOption 编排区域选项
type AreaOption[ID comparable, AreaInfo any] func(area *Area[ID, AreaInfo])

type (
	AreaConstraintHandle[ID comparable, AreaInfo any] func(area *Area[ID, AreaInfo], item Item[ID]) (Item[ID], bool)
	AreaEvaluateHandle[ID comparable, AreaInfo any]   func(areaInfo AreaInfo, items map[ID]Item[ID]) float64
)

// WithAreaConstraint 设置编排区域的约束条件
//   - 该约束用于判断一个成员是否可以被添加到该编排区域中
func WithAreaConstraint[ID comparable, AreaInfo any](constraint AreaConstraintHandle[ID, AreaInfo]) AreaOption[ID, AreaInfo] {
	return func(area *Area[ID, AreaInfo]) {
		area.constraints = append(area.constraints, constraint)
	}
}

// WithAreaEvaluate 设置编排区域的评估函数
//   - 该评估函数将影响成员被编入区域的优先级
func WithAreaEvaluate[ID comparable, AreaInfo any](evaluate AreaEvaluateHandle[ID, AreaInfo]) AreaOption[ID, AreaInfo] {
	return func(area *Area[ID, AreaInfo]) {
		area.evaluate = evaluate
	}
}
