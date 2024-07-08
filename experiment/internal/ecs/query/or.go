package query

import "github.com/kercylan98/minotaur/experiment/internal/ecs"

type or struct {
	l, r ecs.QueryCondition
}

// Or 当实体组件满足 l 或 r 两个条件时，条件成立
func Or(l, r ecs.QueryCondition) ecs.QueryCondition {
	return &or{l: l, r: r}
}

func (f *or) Evaluate(mask *ecs.DynamicBitSet) bool {
	return f.l.Evaluate(mask) || f.r.Evaluate(mask)
}

func (f *or) String() string {
	return "(" + f.l.String() + " || " + f.r.String() + ")"
}
