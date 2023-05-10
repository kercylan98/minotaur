package game

import "github.com/kercylan98/minotaur/utils/huge"

// ItemContainerMember 物品容器成员信息
type ItemContainerMember[ItemID comparable, I Item[ItemID]] interface {
	// GetID 获取物品ID
	GetID() ItemID
	// GetGUID 获取物品GUID
	GetGUID() int64
	// GetCount 获取物品数量
	GetCount() *huge.Int
	// GetItem 获取物品
	GetItem() I
}
