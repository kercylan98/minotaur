package moving

import "github.com/kercylan98/minotaur/game"

// TwoDimensionalEntity 2D移动对象接口定义
type TwoDimensionalEntity interface {
	game.Actor
	game.Position2D
	game.Position2DSet

	// GetSpeed 获取移动速度
	GetSpeed() float64
}
