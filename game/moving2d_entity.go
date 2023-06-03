package game

// Moving2DEntity 2D移动对象接口定义
type Moving2DEntity interface {
	Actor
	Position2D
	Position2DSet

	// GetSpeed 获取移动速度
	GetSpeed() float64
}
