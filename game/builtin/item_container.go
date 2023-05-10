package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
)

func NewItemContainer[ItemID comparable, Item game.Item[ItemID]]() *ItemContainer[ItemID, Item] {
	return &ItemContainer[ItemID, Item]{
		items: map[ItemID]map[int64]*itemContainerInfo[ItemID, Item]{},
	}
}

type ItemContainer[ItemID comparable, Item game.Item[ItemID]] struct {
	sizeLimit  int
	size       int
	expandSize int
	items      map[ItemID]map[int64]*itemContainerInfo[ItemID, Item]
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

func (slf *ItemContainer[ItemID, Item]) AddItem(item Item, count *huge.Int) error {
	if count.LessThanOrEqualTo(huge.IntZero) {
		return ErrCannotAddNegativeItem
	}
	infos, exist := slf.items[item.GetID()]
	if !exist {
		infos = map[int64]*itemContainerInfo[ItemID, Item]{}
		slf.items[item.GetID()] = infos

	}
	info, exist := infos[item.GetGUID()]
	if !exist {
		if slf.GetSizeLimit() >= slf.GetSize() {
			return ErrItemContainerIsFull
		}
		infos[item.GetGUID()] = &itemContainerInfo[ItemID, Item]{
			item:  item,
			count: count,
		}
		slf.size++
	} else {
		info.count = info.count.Add(count)
	}
	return nil
}

func (slf *ItemContainer[ItemID, Item]) DeductItem(id ItemID, count *huge.Int) error {
	return slf.DeductItemWithGuid(id, -1, count)
}

func (slf *ItemContainer[ItemID, Item]) DeductItemWithGuid(id ItemID, guid int64, count *huge.Int) error {
	infos, exist := slf.items[id]
	if !exist || len(infos) == 0 {
		return ErrItemNotExist
	}

	var backupInfos = make(map[int64]*itemContainerInfo[ItemID, Item])
	var pending = count.Copy()
	var deductSize int
	for g, info := range infos {
		if g != guid && guid >= 0 {
			continue
		}
		info.bakCount = info.count.Copy()
		backupInfos[g] = info

		if pending.GreaterThanOrEqualTo(info.count) {
			pending = pending.Sub(info.count)
			info.count = huge.IntZero
		} else {
			info.count = info.count.Sub(pending)
			pending = huge.IntZero
		}

		if info.count.EqualTo(huge.IntZero) {
			delete(infos, guid)
			deductSize++
		}
		if pending.EqualTo(huge.IntZero) {
			break
		}
	}
	if pending.GreaterThan(huge.IntZero) {
		for g, info := range backupInfos {
			infos[g] = info
			info.count = info.bakCount
			info.bakCount = nil
		}
		return ErrItemInsufficientQuantity
	}
	slf.size -= deductSize
	return nil
}
