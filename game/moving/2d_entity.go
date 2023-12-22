package moving

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
)

// TwoDimensionalEntity 2D移动对象接口定义
type TwoDimensionalEntity[EID generic.Basic, PosType generic.SignedNumber] interface {
	// GetTwoDimensionalEntityID 获取 Moving 对象 ID
	GetTwoDimensionalEntityID() EID
	// GetSpeed 获取移动速度
	GetSpeed() float64
	// GetPosition 获取位置
	GetPosition() geometry.Point[PosType]
	// SetPosition 设置位置
	SetPosition(geometry.Point[PosType])
}
