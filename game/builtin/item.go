package builtin

import (
	"minotaur/utils/huge"
)

func NewItem[ID comparable](id ID, count *huge.Int) *Item[ID] {
	if count == nil || count.LessThan(huge.IntZero) {
		panic(ErrItemCountException)
	}
	return &Item[ID]{
		id:    id,
		count: count,
	}
}

type Item[ID comparable] struct {
	id    ID
	guid  int64
	count *huge.Int
}

func (slf *Item[ID]) GetID() ID {
	return slf.id
}

func (slf *Item[ID]) GetGuid() int64 {
	return slf.guid
}

func (slf *Item[ID]) SetGuid(guid int64) {
	slf.guid = guid
}

func (slf *Item[ID]) ChangeStackCount(count *huge.Int) *huge.Int {
	if count.IsZero() {
		return slf.count.Copy()
	}
	newCount := slf.count.Add(count)
	if newCount.LessThan(huge.IntZero) {
		slf.count = newCount.Set(huge.IntZero)
	}
	return slf.count
}

func (slf *Item[ID]) GetStackCount() *huge.Int {
	return slf.count.Copy()
}
