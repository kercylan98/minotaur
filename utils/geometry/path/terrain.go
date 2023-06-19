package path

import (
	"container/heap"
	"github.com/kercylan98/minotaur/utils/geometry"
)

// NewTerrain 返回一个大小为 width 和 height 的新的路径覆盖信息，landformWidth 和 landformHeight 将对每
// 个路径地貌的尺寸进行描述，用于世界位置的转化，例如当 landformWidth 和 landformHeight 为 2, 2 时，路径
// 地貌位于 2, 3 的情况下，它的世界位置将是 4, 6
//   - 地貌特征将默认统一为 PathLandformFeatureRoad
func NewTerrain(width, height, landformWidth, landformHeight int) *Terrain {
	path := &Terrain{
		features:       map[*LandformFeature]map[int]struct{}{},
		width:          width,
		height:         height,
		landformWidth:  landformWidth,
		landformHeight: landformHeight,
	}
	path.matrix = make([]*Landform, width*height)
	for pos := 0; pos < len(path.matrix); pos++ {
		path.matrix[pos] = NewLandform(pos, LandformFeatureRoad)
	}
	path.refreshFeatures()
	return path
}

// NewPathWithMatrix 基于特定的 matrix 返回一个新的路径覆盖信息
//   - 可参照于 NewTerrain
func NewPathWithMatrix(matrix []*Landform, width, height, landformWidth, landformHeight int) *Terrain {
	path := &Terrain{
		matrix:         matrix,
		width:          width,
		height:         height,
		landformWidth:  landformWidth,
		landformHeight: landformHeight,
	}
	path.refreshFeatures()
	return path
}

type Terrain struct {
	matrix                        []*Landform                           // 矩阵
	width, height                 int                                   // 矩阵宽高
	landformWidth, landformHeight int                                   // 地貌宽高
	features                      map[*LandformFeature]map[int]struct{} // 标注了特定特征的路径地貌位置
}

// Get 返回 x 和 y 指向的地貌信息
//   - 通常更建议使用 GetWithPos 进行获取，因为这样可以减少一次转换
func (slf *Terrain) Get(x, y int) *Landform {
	return slf.matrix[geometry.CoordinateToPos(slf.width, x, y)]
}

// GetWithPos 返回 pos 指向的地貌信息
func (slf *Terrain) GetWithPos(pos int) *Landform {
	return slf.matrix[pos]
}

// GetHeight 返回这个路径覆盖的范围高度
func (slf *Terrain) GetHeight() int {
	return slf.height
}

// GetWidth 返回这个路径覆盖的范围宽度
func (slf *Terrain) GetWidth() int {
	return slf.width
}

// GetAll 获取所有地貌信息
func (slf *Terrain) GetAll() []*Landform {
	return slf.matrix
}

// GetPath 返回一个从起点位置到目标位置的路径
//   - 可以通过参数 diagonals 控制是否支持沿对角线进行移动
//   - 可以通过参数 wallsBlockDiagonals 控制在进行对角线移动时是否允许穿越不可通行的区域
func (slf *Terrain) GetPath(startPos, endPos int, diagonals, wallsBlockDiagonals bool) *Path {
	start, end := slf.GetWithPos(startPos), slf.GetWithPos(endPos)

	var nodes h
	var checkedLandforms = make(map[int]struct{})
	var path = new(Path)
	heap.Push(&nodes, &Node{landform: end, cost: end.GetTotalCost()})

	if !start.Walkable() || !end.Walkable() {
		return nil
	}

	for {
		if len(nodes) == 0 {
			break
		}

		node := heap.Pop(&nodes).(*Node)
		if node.landform == start {
			var t = node
			for true {
				path.points = append(path.points, t.landform)
				t = t.parent
				if t == nil {
					break
				}
			}
			break
		}

		for _, adjacent := range geometry.GetAdjacentTranslatePos(slf.matrix, slf.width, node.landform.pos) {
			landform := slf.GetWithPos(adjacent)
			n := &Node{landform: landform, parent: node, cost: landform.GetTotalCost() + node.cost}
			if _, exist := checkedLandforms[adjacent]; n.landform.Walkable() && !exist {
				heap.Push(&nodes, n)
				checkedLandforms[adjacent] = struct{}{}
			}
		}

		if diagonals {
			var up, down, left, right bool
			if upPos := node.landform.pos - slf.width; upPos >= 0 {
				up = slf.GetWithPos(upPos).Walkable()
			}
			if downPos := node.landform.pos + slf.width; downPos < len(slf.matrix) {
				down = slf.GetWithPos(downPos).Walkable()
			}
			row := node.landform.pos / slf.width
			if leftPos := node.landform.pos - 1; row == (leftPos / slf.width) {
				left = slf.GetWithPos(leftPos).Walkable()
			}
			if rightPos := node.landform.pos + 1; row == (rightPos / slf.width) {
				right = slf.GetWithPos(rightPos).Walkable()
			}

			diagonalCost := .414

			size := len(slf.matrix)
			currentRow := node.landform.pos / slf.width
			if topLeft := node.landform.pos - slf.width - 1; topLeft >= 0 && currentRow-1 == (topLeft/slf.width) {
				landform := slf.GetWithPos(topLeft)
				n := &Node{landform: landform, parent: node, cost: landform.GetTotalCost() + node.cost + diagonalCost}
				if _, exist := checkedLandforms[topLeft]; n.landform.Walkable() && !exist && (!wallsBlockDiagonals || (left && up)) {
					heap.Push(&nodes, n)
					checkedLandforms[topLeft] = struct{}{}
				}
			}
			if topRight := node.landform.pos - slf.width + 1; topRight >= 0 && currentRow-1 == (topRight/slf.width) {
				landform := slf.GetWithPos(topRight)
				n := &Node{landform: landform, parent: node, cost: landform.GetTotalCost() + node.cost + diagonalCost}
				if _, exist := checkedLandforms[topRight]; n.landform.Walkable() && !exist && (!wallsBlockDiagonals || (right && up)) {
					heap.Push(&nodes, n)
					checkedLandforms[topRight] = struct{}{}
				}
			}
			if bottomLeft := node.landform.pos + slf.width - 1; bottomLeft < size && currentRow+1 == (bottomLeft/slf.width) {
				landform := slf.GetWithPos(bottomLeft)
				n := &Node{landform: landform, parent: node, cost: landform.GetTotalCost() + node.cost + diagonalCost}
				if _, exist := checkedLandforms[bottomLeft]; n.landform.Walkable() && !exist && (!wallsBlockDiagonals || (left && down)) {
					heap.Push(&nodes, n)
					checkedLandforms[bottomLeft] = struct{}{}
				}
			}
			if bottomRight := node.landform.pos + slf.width + 1; bottomRight < size && currentRow+1 == (bottomRight/slf.width) {
				landform := slf.GetWithPos(bottomRight)
				n := &Node{landform: landform, parent: node, cost: landform.GetTotalCost() + node.cost + diagonalCost}
				if _, exist := checkedLandforms[bottomRight]; n.landform.Walkable() && !exist && (!wallsBlockDiagonals || (right && down)) {
					heap.Push(&nodes, n)
					checkedLandforms[bottomRight] = struct{}{}
				}
			}

		}
	}

	return path

}

// 刷新地貌特征标注信息
//   - 冗余：已使用过的地貌特征类型在未使用后不会被删除
func (slf *Terrain) refreshFeatures() {
	for _, positions := range slf.features {
		for pos := range positions {
			delete(positions, pos)
		}
	}

	for pos, landform := range slf.matrix {
		for _, feature := range landform.GetFeatures() {
			positions, exist := slf.features[feature]
			if !exist {
				positions = map[int]struct{}{}
				slf.features[feature] = positions
			}
			positions[pos] = struct{}{}
		}
	}
}
