package game

import "github.com/kercylan98/minotaur/utils/huge"

// ItemContainerMember 物品容器成员信息
type ItemContainerMember[ItemID comparable] interface {
	Item[ItemID]
	// GetCount 获取物品数量
	GetCount() *huge.Int
}
