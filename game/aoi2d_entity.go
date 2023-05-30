package game

// AOIEntity2D 基于2D定义的AOI对象功能接口
//   - AOI 对象提供了 AOI 系统中常用的属性，诸如位置坐标和视野范围等
type AOIEntity2D interface {
	Actor
	// GetPosition 获取对象位置
	GetPosition() (x, y float64)
	// GetVision 获取视距
	GetVision() float64
}
