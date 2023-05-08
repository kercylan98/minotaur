package builtin

import "github.com/kercylan98/minotaur/utils/huge"

func NewItem[ID comparable](id ID, options ...ItemOption[ID]) *Item[ID] {
	item := &Item[ID]{
		id: id,
	}
	for _, option := range options {
		option(item)
	}
	return item
}

type Item[ID comparable] struct {
	id         ID
	stackLimit *huge.Int
	count      *huge.Int
}

func (slf *Item[ID]) GetID() ID {
	return slf.id
}

func (slf *Item[ID]) GetCount() *huge.Int {
	return slf.count
}

func (slf *Item[ID]) GetStackLimit() *huge.Int {
	return slf.stackLimit
}

func (slf *Item[ID]) SetCount(count *huge.Int) {
	if count.LessThan(huge.IntZero) {
		slf.count = huge.IntZero.Copy()
		return
	}
	slf.count = count.Copy()
}

func (slf *Item[ID]) ChangeCount(count *huge.Int) error {
	newCount := slf.count.Copy().Add(count)
	if newCount.LessThan(huge.IntZero) {
		return ErrItemInsufficientQuantityDeduction
	}
	if newCount.GreaterThan(slf.stackLimit) {
		return ErrItemStackLimit
	}
	slf.count = newCount
	return nil
}
