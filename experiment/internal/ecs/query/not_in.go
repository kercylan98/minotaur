package query

import "github.com/kercylan98/minotaur/experiment/internal/ecs"

type notIn struct {
	mask *ecs.DynamicBitSet
}

// NotIn 当实体组件不包含 componentIds 中的所有 ComponentId 时，条件成立
func NotIn(componentIds ...ecs.ComponentId) ecs.QueryCondition {
	mask := new(ecs.DynamicBitSet)
	for _, i := range componentIds {
		mask.Set(i)
	}
	return &notIn{mask: mask}
}

func (f *notIn) Evaluate(mask *ecs.DynamicBitSet) bool {
	return mask.NotIn(f.mask)
}

func (f *notIn) String() string {
	return "not in(" + f.mask.String() + ")"
}
