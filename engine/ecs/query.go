package ecs

import (
	"github.com/kercylan98/minotaur/toolkit"
	"strings"
)

type Query interface {
	Evaluate(mask *toolkit.DynamicBitSet) bool
	String() string
}

// And 当实体组件同时满足所有条件时，条件成立
func And(query ...Query) Query {
	return &and{q: query}
}

type and struct {
	q []Query
}

func (q *and) Evaluate(mask *toolkit.DynamicBitSet) bool {
	for _, query := range q.q {
		if !query.Evaluate(mask) {
			return false
		}
	}
	return true
}

func (q *and) String() string {
	var builder strings.Builder
	builder.WriteString("(")
	for i, query := range q.q {
		if i > 0 {
			builder.WriteString(" && ")
		}
		builder.WriteString(query.String())
	}
	builder.WriteString(")")
	return builder.String()
}

// Equal 当实体组件完全等于 componentIds 中的所有 ComponentId 时，条件成立
func Equal(componentIds ...ComponentId) Query {
	mask := new(toolkit.DynamicBitSet)
	for _, i := range componentIds {
		mask.Set(i)
	}
	return &equal{mask: mask}
}

type equal struct {
	mask *toolkit.DynamicBitSet
}

func (q *equal) Evaluate(mask *toolkit.DynamicBitSet) bool {
	return q.mask.Equal(mask)
}

func (q *equal) String() string {
	return "(" + q.mask.String() + " == " + q.mask.String() + ")"
}

type in struct {
	mask *toolkit.DynamicBitSet
}

// In 当实体组件包含 componentIds 中的所有 ComponentId 时，条件成立
func In(componentIds ...ComponentId) Query {
	mask := new(toolkit.DynamicBitSet)
	for _, i := range componentIds {
		mask.Set(i)
	}
	return &in{mask: mask}
}

func (q *in) Evaluate(mask *toolkit.DynamicBitSet) bool {
	return mask.In(q.mask)
}

func (q *in) String() string {
	return "in(" + q.mask.String() + ")"
}

type notIn struct {
	mask *toolkit.DynamicBitSet
}

// NotIn 当实体组件不包含 componentIds 中的所有 ComponentId 时，条件成立
func NotIn(componentIds ...ComponentId) Query {
	mask := new(toolkit.DynamicBitSet)
	for _, i := range componentIds {
		mask.Set(i)
	}
	return &notIn{mask: mask}
}

func (q *notIn) Evaluate(mask *toolkit.DynamicBitSet) bool {
	return mask.NotIn(q.mask)
}

func (q *notIn) String() string {
	return "not in(" + q.mask.String() + ")"
}

type or struct {
	q []Query
}

// Or 当实体组件满足 l 或 r 两个条件时，条件成立
func Or(query ...Query) Query {
	return &or{q: query}
}

func (q *or) Evaluate(mask *toolkit.DynamicBitSet) bool {
	for _, query := range q.q {
		if query.Evaluate(mask) {
			return true
		}
	}
	return false
}

func (q *or) String() string {
	var builder strings.Builder
	builder.WriteString("(")
	for i, query := range q.q {
		if i > 0 {
			builder.WriteString(" || ")
		}
		builder.WriteString(query.String())
	}
	builder.WriteString(")")
	return builder.String()
}
