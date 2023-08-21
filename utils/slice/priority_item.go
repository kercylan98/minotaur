package slice

// NewPriorityItem 创建一个优先级切片元素
func NewPriorityItem[V any](v V, priority int) *PriorityItem[V] {
	return &PriorityItem[V]{v: v, p: priority}
}

// PriorityItem 是一个优先级切片元素
type PriorityItem[V any] struct {
	next *PriorityItem[V]
	prev *PriorityItem[V]
	v    V
	p    int
}

// Value 返回元素值
func (p *PriorityItem[V]) Value() V {
	return p.v
}

// Priority 返回元素优先级
func (p *PriorityItem[V]) Priority() int {
	return p.p
}

// Next 返回下一个元素
func (p *PriorityItem[V]) Next() *PriorityItem[V] {
	return p.next
}

// Prev 返回上一个元素
func (p *PriorityItem[V]) Prev() *PriorityItem[V] {
	return p.prev
}
