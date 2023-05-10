package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
	"github.com/kercylan98/minotaur/utils/super"
	"sync/atomic"
)

func NewItemContainer[ItemID comparable, Item game.Item[ItemID]](options ...ItemContainerOption[ItemID, Item]) *ItemContainer[ItemID, Item] {
	itemContainer := &ItemContainer[ItemID, Item]{
		items:         map[int64]*ItemContainerMember[ItemID, Item]{},
		itemIdGuidRef: map[ItemID]map[int64]bool{},
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
	sort          []*int64
	maxSort       int
	vacancy       []int
	stackLimit    map[ItemID]*huge.Int
}

func (slf *ItemContainer[ItemID, Item]) GetSize() int {
	return slf.size
}

func (slf *ItemContainer[ItemID, Item]) GetSizeLimit() int {
	return slf.sizeLimit + slf.expandSize
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
	var result = make([]game.ItemContainerMember[ItemID, Item], 0, len(slf.sort))
	for _, guid := range slf.sort {
		if guid == nil {
			continue
		}
		result = append(result, slf.items[*guid])
	}
	return result
}

func (slf *ItemContainer[ItemID, Item]) GetItemsFull() []game.ItemContainerMember[ItemID, Item] {
	var result = make([]game.ItemContainerMember[ItemID, Item], len(slf.sort), len(slf.sort))
	for i, guid := range slf.sort {
		if guid == nil {
			result[i] = nil
		} else {
			result[i] = slf.items[*guid]
		}
	}
	sizeLimit := slf.GetSizeLimit()
	for sizeLimit > 0 && len(result) > sizeLimit {
		result = append(result, nil)
	}
	return result
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

func (slf *ItemContainer[ItemID, Item]) AddItem(item Item, count *huge.Int) (guid int64, err error) {
	if count.LessThanOrEqualTo(huge.IntZero) {
		return 0, ErrCannotAddNegativeOrZeroItem
	}
	for guid := range slf.itemIdGuidRef[item.GetID()] {
		member := slf.items[guid]
		if member.GetItem().IsSame(item) {
			if stackLimit := slf.stackLimit[item.GetID()]; stackLimit != nil && member.count.Copy().Add(count).GreaterThan(stackLimit) {
				continue
			}
			member.count.Add(count)
			return guid, nil
		}
	}
	if slf.size >= slf.GetSizeLimit() {
		return 0, ErrItemContainerIsFull
	}
	guid = slf.guid.Add(1)
	slf.items[guid] = &ItemContainerMember[ItemID, Item]{
		item:  item,
		guid:  guid,
		count: count.Copy(),
		sort: super.If(len(slf.vacancy) == 0,
			func() int {
				sort := len(slf.sort)
				slf.sort = append(slf.sort, &guid)
				if sort > slf.maxSort {
					slf.maxSort = sort
				}
				return sort
			}(),
			func() int {
				sort := slf.vacancy[0]
				slf.vacancy = slf.vacancy[1:]
				slf.sort[sort] = &guid
				return sort
			}(),
		),
	}
	guids, exist := slf.itemIdGuidRef[item.GetID()]
	if !exist {
		guids = map[int64]bool{}
		slf.itemIdGuidRef[item.GetID()] = guids
	}
	guids[guid] = true
	slf.size++
	return guid, nil
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
			slf.sort[member.sort] = nil
			slf.vacancy = append(slf.vacancy, member.sort)
			delete(slf.items, guid)
			sizeLimit := slf.GetSizeLimit()
			for slf.sort[slf.maxSort] == nil && slf.maxSort > sizeLimit {
				slf.sort = append(slf.sort[0:slf.maxSort], slf.sort[slf.maxSort+1:]...)
				slf.maxSort--
			}
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
					slf.sort[member.sort] = nil
					slf.vacancy = append(slf.vacancy, member.sort)
					sizeLimit := slf.GetSizeLimit()
					for slf.sort[slf.maxSort] == nil && slf.maxSort > sizeLimit {
						slf.sort = append(slf.sort[0:slf.maxSort], slf.sort[slf.maxSort+1:]...)
						slf.maxSort--
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

func (slf *ItemContainer[ItemID, Item]) TransferTo(guid int64, count *huge.Int, target game.ItemContainer[ItemID, Item]) error {
	if target == nil {
		return ErrItemContainerNotExist
	}
	if count.LessThanOrEqualTo(huge.IntZero) {
		return ErrCannotAddNegativeOrZeroItem
	}
	member, err := slf.GetItem(guid)
	if err != nil {
		return err
	}
	if err = slf.CheckDeductItem(guid, count); err != nil {
		return err
	}
	if err = target.CheckAllowAdd(member.GetItem(), count); err != nil {
		return err
	}
	_ = slf.DeductItem(guid, count)
	_, _ = target.AddItem(member.GetItem(), count)
	return nil
}

func (slf *ItemContainer[ItemID, Item]) CheckAllowAdd(item Item, count *huge.Int) error {
	if count.LessThanOrEqualTo(huge.IntZero) {
		return ErrCannotAddNegativeOrZeroItem
	}
	for guid := range slf.itemIdGuidRef[item.GetID()] {
		member := slf.items[guid]
		if member.GetItem().IsSame(item) {
			if stackLimit := slf.stackLimit[item.GetID()]; stackLimit != nil && member.count.Copy().Add(count).GreaterThan(stackLimit) {
				continue
			}
			return nil
		}
	}
	if slf.size >= slf.GetSizeLimit() {
		return ErrItemContainerIsFull
	}
	return nil
}

func (slf *ItemContainer[ItemID, Item]) CheckDeductItem(guid int64, count *huge.Int) error {
	if !slf.ExistItem(guid) {
		return ErrItemNotExist
	}
	member := slf.items[guid]
	if member.count.GreaterThanOrEqualTo(count) {
		return nil
	} else {
		var need = count.Copy()
		var guids = slf.itemIdGuidRef[member.GetID()]
		for guid := range guids {
			member := slf.items[guid]
			if need.GreaterThanOrEqualTo(member.count) {
				need.Sub(member.count)
			} else {
				need = huge.IntZero
			}
		}
		if need.GreaterThan(huge.IntZero) {
			return ErrItemInsufficientQuantity
		}
		return nil
	}
}

func (slf *ItemContainer[ItemID, Item]) Remove(guid int64) {
	member, exist := slf.items[guid]
	if !exist {
		return
	}
	delete(slf.items, guid)
	delete(slf.itemIdGuidRef[member.GetID()], guid)
	slf.sort[member.sort] = nil
	slf.vacancy = append(slf.vacancy, member.sort)

	sizeLimit := slf.GetSizeLimit()
	for slf.sort[slf.maxSort] == nil && slf.maxSort > sizeLimit {
		slf.sort = append(slf.sort[0:slf.maxSort], slf.sort[slf.maxSort+1:]...)
		slf.maxSort--
	}
}

func (slf *ItemContainer[ItemID, Item]) RemoveWithID(id ItemID) {
	for guid := range slf.itemIdGuidRef[id] {
		member, exist := slf.items[guid]
		if !exist {
			continue
		}
		delete(slf.items, guid)
		slf.sort[member.sort] = nil
		slf.vacancy = append(slf.vacancy, member.sort)
	}
	delete(slf.itemIdGuidRef, id)

	sizeLimit := slf.GetSizeLimit()
	for slf.sort[slf.maxSort] == nil && slf.maxSort > sizeLimit {
		slf.sort = append(slf.sort[0:slf.maxSort], slf.sort[slf.maxSort+1:]...)
		slf.maxSort--
	}
}

func (slf *ItemContainer[ItemID, Item]) Clear() {
	slf.items = map[int64]*ItemContainerMember[ItemID, Item]{}
	slf.itemIdGuidRef = map[ItemID]map[int64]bool{}
	slf.sort = nil
	slf.maxSort = 0
	slf.vacancy = nil
}
