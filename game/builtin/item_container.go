package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
	"github.com/kercylan98/minotaur/utils/slice"
)

func NewItemContainer[ItemID comparable, Item game.Item[ItemID]](options ...ItemContainerOption[ItemID, Item]) *ItemContainer[ItemID, Item] {
	itemContainer := &ItemContainer[ItemID, Item]{
		items: map[ItemID]map[int64]*ItemContainerMember[ItemID]{},
	}
	for _, option := range options {
		option(itemContainer)
	}
	return itemContainer
}

type ItemContainer[ItemID comparable, Item game.Item[ItemID]] struct {
	sizeLimit  int
	size       int
	expandSize int
	items      map[ItemID]map[int64]*ItemContainerMember[ItemID]
	sort       []*itemContainerSort[ItemID]
}

func (slf *ItemContainer[ItemID, Item]) GetSize() int {
	return slf.size
}

func (slf *ItemContainer[ItemID, Item]) SetExpandSize(size int) {
	slf.expandSize = size
}

func (slf *ItemContainer[ItemID, Item]) GetSizeLimit() int {
	return slf.sizeLimit + slf.expandSize
}

func (slf *ItemContainer[ItemID, Item]) GetItem(id ItemID) (game.ItemContainerMember[ItemID], error) {
	for _, member := range slf.items[id] {
		return member, nil
	}
	return nil, ErrItemNotExist
}

func (slf *ItemContainer[ItemID, Item]) GetItemWithGuid(id ItemID, guid int64) (game.ItemContainerMember[ItemID], error) {
	member, exist := slf.items[id][guid]
	if !exist {
		return nil, ErrItemNotExist
	}
	return member, nil
}

func (slf *ItemContainer[ItemID, Item]) GetItems() []game.ItemContainerMember[ItemID] {
	var items []game.ItemContainerMember[ItemID]
	for _, sort := range slf.sort {
		items = append(items, slf.items[sort.id][sort.guid])
	}
	return items
}

func (slf *ItemContainer[ItemID, Item]) AddItem(item Item, count *huge.Int) error {
	if count.LessThanOrEqualTo(huge.IntZero) {
		return ErrCannotAddNegativeItem
	}
	members, exist := slf.items[item.GetID()]
	if !exist {
		members = map[int64]*ItemContainerMember[ItemID]{}
		slf.items[item.GetID()] = members

	}
	member, exist := members[item.GetGUID()]
	if !exist {
		if slf.GetSizeLimit() >= slf.GetSize() {
			return ErrItemContainerIsFull
		}
		members[item.GetGUID()] = &ItemContainerMember[ItemID]{
			sort:  len(slf.sort),
			Item:  item,
			count: count,
		}
		slf.sort = append(slf.sort, &itemContainerSort[ItemID]{
			id:   item.GetID(),
			guid: item.GetGUID(),
		})
		slf.size++
	} else {
		member.count = member.count.Add(count)
	}
	return nil
}

func (slf *ItemContainer[ItemID, Item]) DeductItem(id ItemID, count *huge.Int) error {
	members, exist := slf.items[id]
	if !exist || len(members) == 0 {
		return ErrItemNotExist
	}

	var backupMembers = make(map[int64]*ItemContainerMember[ItemID])
	var pending = count.Copy()
	var deductMembers []*ItemContainerMember[ItemID]
	for guid, member := range members {
		member.bakCount = member.count.Copy()
		backupMembers[guid] = member

		if pending.GreaterThanOrEqualTo(member.count) {
			pending = pending.Sub(member.count)
			member.count = huge.IntZero
		} else {
			member.count = member.count.Sub(pending)
			pending = huge.IntZero
		}

		if member.count.EqualTo(huge.IntZero) {
			delete(members, guid)
			deductMembers = append(deductMembers, member)
		}
		if pending.EqualTo(huge.IntZero) {
			break
		}
	}
	if pending.GreaterThan(huge.IntZero) {
		for guid, member := range backupMembers {
			members[guid] = member
			member.count = member.bakCount
			member.bakCount = nil
		}
		return ErrItemInsufficientQuantity
	}
	slf.size -= len(deductMembers)
	for _, member := range deductMembers {
		slice.Del(&slf.sort, member.sort)
	}
	return nil
}

func (slf *ItemContainer[ItemID, Item]) DeductItemWithGuid(id ItemID, guid int64, count *huge.Int) error {
	member, exist := slf.items[id][guid]
	if !exist {
		return ErrItemNotExist
	}
	if count.GreaterThan(member.count) {
		return ErrItemInsufficientQuantity
	} else {
		member.count = member.count.Sub(count)
	}
	if member.count.EqualTo(huge.IntZero) {
		delete(slf.items[id], guid)
		slice.Del(&slf.sort, member.sort)
		slf.size--
	}
	return nil
}
