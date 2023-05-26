package game

// Terrain2D 地形
type Terrain2D interface {
	GetBlock(x, y int) Terrain2DBlock
	GetBlocks() [][]Terrain2DBlock
	GetWidth() int
	GetHeight() int
}
