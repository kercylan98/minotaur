package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
	"sync/atomic"
)

func NewItemContainer[ItemID comparable, Item game.Item[ItemID]](options ...ItemContainerOption[ItemID, Item]) *ItemContainer[ItemID, Item] {
	itemContainer := &ItemContainer[ItemID, Item]{
		items:         map[int64]*ItemContainerMember[ItemID, Item]{},
		itemIdGuidRef: map[ItemID]map[int64]bool{},
		stackLimits:   map[ItemID]*huge.Int{},
	}
	for _, option := range options {
		option(itemContainer)
	}
	return itemContainer
}

type ItemContainer[ItemID comparable, Item game.Item[ItemID]] struct {
	guid          atomic.Int64
	sizeLimit     int
	size          int
	expandSize    int
	items         map[int64]*ItemContainerMember[ItemID, Item]
	itemIdGuidRef map[ItemID]map[int64]bool
	sort          []*itemContainerSort[ItemID]
	stackLimits   map[ItemID]*huge.Int
}

func (slf *ItemContainer[ItemID, Item]) GetSize() int {
	return slf.size + slf.expandSize
}

func (slf *ItemContainer[ItemID, Item]) GetSizeLimit() int {
	return slf.sizeLimit
}

func (slf *ItemContainer[ItemID, Item]) SetExpandSize(size int) {
	slf.expandSize = size
}

func (slf *ItemContainer[ItemID, Item]) GetItem(guid int64) (game.ItemContainerMember[ItemID, Item], error) {
	item, exist := slf.items[guid]
	if !exist {
		return nil, ErrItemNotExist
	}
	return item, nil
}

func (slf *ItemContainer[ItemID, Item]) GetItems() []game.ItemContainerMember[ItemID, Item] {
	//TODO implement me
	panic("implement me")
}

func (slf *ItemContainer[ItemID, Item]) GetItemsMap() map[int64]game.ItemContainerMember[ItemID, Item] {
	var m = make(map[int64]game.ItemContainerMember[ItemID, Item])
	for k, v := range slf.items {
		m[k] = v
	}
	return m
}

func (slf *ItemContainer[ItemID, Item]) ExistItem(guid int64) bool {
	_, exist := slf.items[guid]
	return exist
}

func (slf *ItemContainer[ItemID, Item]) ExistItemWithID(id ItemID) bool {
	return len(slf.itemIdGuidRef[id]) > 0
}

func (slf *ItemContainer[ItemID, Item]) AddItem(item Item, count *huge.Int) error {
	if count.LessThanOrEqualTo(huge.IntZero) {
		return ErrCannotAddNegativeOrZeroItem
	}
	for guid := range slf.itemIdGuidRef[item.GetID()] {
		member := slf.items[guid]
		if member.GetItem().IsSame(item) {
			member.count.Add(count)
			return nil
		}
	}
	if slf.size >= slf.GetSizeLimit() {
		return ErrItemContainerIsFull
	}
	guid := slf.guid.Add(1)
	slf.items[guid] = &ItemContainerMember[ItemID, Item]{
		item:  item,
		guid:  guid,
		count: count.Copy(),
	}
	guids, exist := slf.itemIdGuidRef[item.GetID()]
	if !exist {
		guids = map[int64]bool{}
		slf.itemIdGuidRef[item.GetID()] = guids
	}
	guids[guid] = true
	slf.size++
	return nil
}

func (slf *ItemContainer[ItemID, Item]) DeductItem(guid int64, count *huge.Int) error {
	if !slf.ExistItem(guid) {
		return ErrItemNotExist
	}
	member := slf.items[guid]
	if member.count.GreaterThanOrEqualTo(count) {
		member.count.Sub(count)
		if member.count.EqualTo(huge.IntZero) {
			slf.size--
			delete(slf.items, guid)
		}
		return nil
	} else {
		var need = count.Copy()
		var handles []func()
		var guids = slf.itemIdGuidRef[member.GetID()]
		for guid := range guids {
			member := slf.items[guid]
			if need.GreaterThanOrEqualTo(member.count) {
				need.Sub(member.count)
				handles = append(handles, func() {
					member.count = huge.IntZero.Copy()
					slf.size--
					delete(guids, guid)
					delete(slf.items, guid)
					if len(guids) == 0 {
						delete(slf.itemIdGuidRef, member.GetID())
					}
				})
			} else {
				member.count.Sub(need)
				need = huge.IntZero
			}
		}
		if need.GreaterThan(huge.IntZero) {
			return ErrItemInsufficientQuantity
		}
		for _, handle := range handles {
			handle()
		}
		return nil
	}
}
