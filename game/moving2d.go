package game

// Moving2D 2D移动功能接口定义
type Moving2D interface {
	MoveTo(entity Moving2DEntity, x float64, y float64)
}
