package poker

import "github.com/kercylan98/minotaur/utils/generic"

// Card 扑克牌
type Card[P, C generic.Number] interface {
	// GetGuid 获取扑克牌的唯一标识
	GetGuid() int64
	// GetPoint 获取扑克牌的点数
	GetPoint() P
	// GetColor 获取扑克牌的花色
	GetColor() C
}
