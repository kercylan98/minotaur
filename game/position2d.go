package game

// Position2D 2D位置接口定义
type Position2D interface {
	// GetPosition 获取对象位置
	GetPosition() (x, y float64)
}
