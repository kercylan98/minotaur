package game

type Item[ID comparable] interface {
	// GetID 获取物品ID
	GetID() ID
	// GetGUID 获取物品GUID
	//  - 用于标识同一件物品不同的特征
	//  - 负数的GUID在内置功能中可能会被用于特殊判定，如果需要负数建议另外对特殊功能进行实现
	GetGUID() int64
}
