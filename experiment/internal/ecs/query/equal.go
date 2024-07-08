package query

import "github.com/kercylan98/minotaur/experiment/internal/ecs"

type equal struct {
	mask *ecs.DynamicBitSet
}

// Equal 当实体组件完全等于 componentIds 中的所有 ComponentId 时，条件成立
func Equal(componentIds ...ecs.ComponentId) ecs.QueryCondition {
	mask := new(ecs.DynamicBitSet)
	for _, i := range componentIds {
		mask.Set(i)
	}
	return &equal{mask: mask}
}

func (f *equal) Evaluate(mask *ecs.DynamicBitSet) bool {
	return f.mask.Equal(mask)
}

func (f *equal) String() string {
	return "(" + f.mask.String() + " == " + f.mask.String() + ")"
}
