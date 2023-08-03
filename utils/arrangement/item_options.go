package arrangement

// ItemOption 编排成员选项
type ItemOption[ID comparable, AreaInfo any] func(arrangement *Arrangement[ID, AreaInfo], item Item[ID])

type (
	ItemFixedAreaHandle[AreaInfo any]               func(areaInfo AreaInfo) bool
	ItemPriorityHandle[ID comparable, AreaInfo any] func(areaInfo AreaInfo, item Item[ID]) float64
)

// WithItemFixed 设置成员的固定编排区域
func WithItemFixed[ID comparable, AreaInfo any](matcher ItemFixedAreaHandle[AreaInfo]) ItemOption[ID, AreaInfo] {
	return func(arrangement *Arrangement[ID, AreaInfo], item Item[ID]) {
		arrangement.fixed[item.GetID()] = matcher
	}
}

// WithItemPriority 设置成员的优先级
func WithItemPriority[ID comparable, AreaInfo any](priority ItemPriorityHandle[ID, AreaInfo]) ItemOption[ID, AreaInfo] {
	return func(arrangement *Arrangement[ID, AreaInfo], item Item[ID]) {
		arrangement.priority[item.GetID()] = append(arrangement.priority[item.GetID()], priority)
	}
}
