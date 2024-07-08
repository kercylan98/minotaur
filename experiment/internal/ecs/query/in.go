package query

import "github.com/kercylan98/minotaur/experiment/internal/ecs"

type in struct {
	mask *ecs.DynamicBitSet
}

// In 当实体组件包含 componentIds 中的所有 ComponentId 时，条件成立
func In(componentIds ...ecs.ComponentId) ecs.QueryCondition {
	mask := new(ecs.DynamicBitSet)
	for _, i := range componentIds {
		mask.Set(i)
	}
	return &in{mask: mask}
}

func (f *in) Evaluate(mask *ecs.DynamicBitSet) bool {
	return mask.In(f.mask)
}

func (f *in) String() string {
	return "in(" + f.mask.String() + ")"
}
