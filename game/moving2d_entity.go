package game

type Moving2DEntity interface {
	Actor
	Position2D
	Position2DSet

	// GetSpeed 获取移动速度
	GetSpeed() float64
}
