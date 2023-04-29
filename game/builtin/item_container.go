package builtin

import (
	"minotaur/game"
	"minotaur/utils/huge"
	"minotaur/utils/synchronization"
)

func NewItemContainer[ItemID comparable, I game.Item[ItemID]]() *ItemContainer[ItemID, I] {
	return &ItemContainer[ItemID, I]{
		items:   synchronization.NewMap[ItemID, *synchronization.Map[int64, I]](),
		itemRef: synchronization.NewMap[int64, ItemID](),
	}
}

type ItemContainer[ItemID comparable, I game.Item[ItemID]] struct {
	guid    int64
	items   *synchronization.Map[ItemID, *synchronization.Map[int64, I]]
	itemRef *synchronization.Map[int64, ItemID]
}

func (slf *ItemContainer[ItemID, I]) GetItem(guid int64) I {
	id := slf.itemRef.Get(guid)
	return slf.items.Get(id).Get(guid)
}

func (slf *ItemContainer[ItemID, I]) GetItems() map[int64]I {
	items := make(map[int64]I)
	slf.items.Range(func(id ItemID, value *synchronization.Map[int64, I]) {
		if value != nil {
			value.Range(func(guid int64, value I) {
				items[guid] = value
			})
		}
	})
	return items
}

func (slf *ItemContainer[ItemID, I]) GetItemsWithId(id ItemID) map[int64]I {
	return slf.items.Get(id).Map()
}

func (slf *ItemContainer[ItemID, I]) AddItem(item I) error {
	id := item.GetID()
	items, exist := slf.items.GetExist(id)
	if !exist {
		items = synchronization.NewMap[int64, I]()
		slf.items.Set(id, items)
	}
	slf.guid++
	items.Set(slf.guid, item)
	slf.itemRef.Set(slf.guid, id)
	return nil
}

func (slf *ItemContainer[ItemID, I]) ChangeItemCount(guid int64, count *huge.Int) error {
	item := slf.GetItem(guid)
	item.ChangeStackCount(count)
	return nil
}

func (slf *ItemContainer[ItemID, I]) DeleteItem(guid int64) {
	id := slf.GetItem(guid).GetID()
	slf.items.Get(id).Delete(guid)
	slf.itemRef.Delete(guid)
}

func (slf *ItemContainer[ItemID, I]) DeleteItemsWithId(id ItemID) {
	for guid := range slf.items.Get(id).Map() {
		slf.itemRef.Delete(guid)
	}
	slf.items.Delete(id)
}
