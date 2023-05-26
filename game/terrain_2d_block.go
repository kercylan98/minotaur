package game

// Terrain2DBlock 地形块
type Terrain2DBlock interface {
	// GetTerrain 获取归属的地形
	GetTerrain() Terrain2D
	// GetCost 获取移动消耗
	GetCost() float64
}
