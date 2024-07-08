package query

import "github.com/kercylan98/minotaur/experiment/internal/ecs"

type and struct {
	l, r ecs.QueryCondition
}

// And 当实体组件同时满足 l 和 r 两个条件时，条件成立
func And(l, r ecs.QueryCondition) ecs.QueryCondition {
	return &and{l: l, r: r}
}

func (f *and) Evaluate(mask *ecs.DynamicBitSet) bool {
	return f.l.Evaluate(mask) && f.r.Evaluate(mask)
}

func (f *and) String() string {
	return "(" + f.l.String() + " && " + f.r.String() + ")"
}
