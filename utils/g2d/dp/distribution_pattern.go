package dp

import (
	"github.com/kercylan98/minotaur/utils/g2d"
)

// NewDistributionPattern 构建一个分布图实例
func NewDistributionPattern[Item any](sameKindVerifyHandle func(itemA, itemB Item) bool) *DistributionPattern[Item] {
	return &DistributionPattern[Item]{
		links:                map[int]map[int]Item{},
		sameKindVerifyHandle: sameKindVerifyHandle,
	}
}

// DistributionPattern 分布图
type DistributionPattern[Item any] struct {
	matrix               []Item
	links                map[int]map[int]Item
	sameKindVerifyHandle func(itemA, itemB Item) bool
	width                int
	usePos               bool
}

// GetLinks 获取关联的成员
//   - 其中包含传入的 pos 成员
func (slf *DistributionPattern[Item]) GetLinks(pos int) (result []Link[Item]) {
	for pos, item := range slf.links[pos] {
		result = append(result, Link[Item]{Pos: pos, Item: item})
	}
	return
}

// HasLink 检查一个位置是否包含除它本身外的其他关联成员
func (slf *DistributionPattern[Item]) HasLink(pos int) bool {
	links, exist := slf.links[pos]
	if !exist {
		return false
	}
	return len(links) > 1
}

// LoadMatrix 通过二维矩阵加载分布图
//   - 通过该函数加载的分布图使用的矩阵是复制后的矩阵，因此无法直接通过刷新(Refresh)来更新分布关系
//   - 需要通过直接刷新的方式请使用 LoadMatrixWithPos
func (slf *DistributionPattern[Item]) LoadMatrix(matrix [][]Item) {
	slf.LoadMatrixWithPos(g2d.MatrixToPosMatrix(matrix))
	slf.usePos = false
}

// LoadMatrixWithPos 通过二维矩阵加载分布图
func (slf *DistributionPattern[Item]) LoadMatrixWithPos(width int, matrix []Item) {
	slf.width = width
	slf.matrix = matrix
	slf.usePos = true
	for k := range slf.links {
		delete(slf.links, k)
	}
	for pos, item := range slf.matrix {
		slf.buildRelationships(pos, item)
	}
}

// Refresh 刷新特定位置的分布关系
//   - 由于 LoadMatrix 的矩阵是复制后的矩阵，所以任何外部的改动都不会影响到分布图的变化，在这种情况下，刷新将没有任何意义
//   - 需要通过直接刷新的方式请使用 LoadMatrixWithPos 加载矩阵，或者通过 RefreshWithItem 函数进行刷新
func (slf *DistributionPattern[Item]) Refresh(pos int) {
	if !slf.usePos {
		return
	}
	links, exist := slf.links[pos]
	if !exist {
		slf.buildRelationships(pos, slf.matrix[pos])
		return
	}
	var positions []int
	for tp := range links {
		positions = append(positions, tp)
		delete(slf.links, tp)
	}
	for _, tp := range positions {
		slf.buildRelationships(tp, slf.matrix[tp])
	}
}

// RefreshWithItem 通过特定的成员刷新特定位置的分布关系
//   - 如果矩阵通过 LoadMatrixWithPos 加载，将会重定向至 Refresh
func (slf *DistributionPattern[Item]) RefreshWithItem(pos int, item Item) {
	if slf.usePos {
		slf.Refresh(pos)
		return
	}

	slf.matrix[pos] = item
	links, exist := slf.links[pos]
	if !exist {
		slf.buildRelationships(pos, slf.matrix[pos])
		return
	}
	var positions []int
	for tp := range links {
		positions = append(positions, tp)
		delete(slf.links, tp)
	}
	for _, tp := range positions {
		slf.buildRelationships(tp, slf.matrix[tp])
	}
}

// 构建关系
func (slf *DistributionPattern[Item]) buildRelationships(pos int, item Item) {
	links, exist := slf.links[pos]
	if !exist {
		links = map[int]Item{pos: item}
		slf.links[pos] = links
	}

	for _, tp := range g2d.GetAdjacentCoordinatesWithPos(slf.matrix, slf.width, pos) {
		target := slf.matrix[tp]
		if _, exist := links[tp]; exist || !slf.sameKindVerifyHandle(item, target) {
			continue
		}

		slf.links[tp] = links
		links[tp] = target
		slf.buildRelationships(tp, target)
	}
}
