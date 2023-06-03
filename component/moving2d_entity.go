package component

import "github.com/kercylan98/minotaur/game"

// Moving2DEntity 2D移动对象接口定义
type Moving2DEntity interface {
	game.Actor
	game.Position2D
	game.Position2DSet

	// GetSpeed 获取移动速度
	GetSpeed() float64
}
