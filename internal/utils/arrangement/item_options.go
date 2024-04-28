package arrangement

// ItemOption 编排成员选项
type ItemOption[ID comparable, AreaInfo any] func(arrangement *Arrangement[ID, AreaInfo], item Item[ID])

type (
	ItemFixedAreaHandle[AreaInfo any]                     func(areaInfo AreaInfo) bool
	ItemPriorityHandle[ID comparable, AreaInfo any]       func(areaInfo AreaInfo, item Item[ID]) float64
	ItemNotAllowVerifyHandle[ID comparable, AreaInfo any] func(areaInfo AreaInfo, item Item[ID]) bool
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

// WithItemNotAllow 设置成员不允许的编排区域
func WithItemNotAllow[ID comparable, AreaInfo any](verify ItemNotAllowVerifyHandle[ID, AreaInfo]) ItemOption[ID, AreaInfo] {
	return func(arrangement *Arrangement[ID, AreaInfo], item Item[ID]) {
		arrangement.itemNotAllow[item.GetID()] = append(arrangement.itemNotAllow[item.GetID()], verify)
	}
}
