package aoi

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
)

// TwoDimensionalEntity 基于2D定义的AOI对象功能接口
//   - AOI 对象提供了 AOI 系统中常用的属性，诸如位置坐标和视野范围等
type TwoDimensionalEntity[EID generic.Basic, PosType generic.SignedNumber] interface {
	// GetTwoDimensionalEntityID 获取 AOI 对象 ID
	GetTwoDimensionalEntityID() EID
	// GetVision 获取视距
	GetVision() float64
	// GetPosition 获取位置
	GetPosition() geometry.Point[PosType]
}
