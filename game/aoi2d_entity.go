package game

type AOIEntity2D interface {
	Actor
	// GetPosition 获取对象位置
	GetPosition() (x, y float64)
	// GetVision 获取视距
	GetVision() float64
}
