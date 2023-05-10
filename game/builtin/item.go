package builtin

import "github.com/kercylan98/minotaur/game"

func NewItem[ItemID comparable](id ItemID) *Item[ItemID] {
	item := &Item[ItemID]{
		id: id,
	}
	return item
}

type Item[ItemID comparable] struct {
	id ItemID
}

func (slf *Item[ItemID]) GetID() ItemID {
	return slf.id
}

func (slf *Item[ItemID]) IsSame(item game.Item[ItemID]) bool {
	return slf.id == item.GetID()
}
