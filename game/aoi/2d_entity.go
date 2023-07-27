package aoi

import "github.com/kercylan98/minotaur/game"

// TwoDimensionalEntity 基于2D定义的AOI对象功能接口
//   - AOI 对象提供了 AOI 系统中常用的属性，诸如位置坐标和视野范围等
type TwoDimensionalEntity interface {
	game.Actor
	game.Position2D
	// GetVision 获取视距
	GetVision() float64
}
