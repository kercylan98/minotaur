package path

import (
	"github.com/kercylan98/minotaur/utils/g2d"
)

func NewLandform(pos int, features ...*LandformFeature) *Landform {
	if len(features) == 0 {
		panic("path landforms without any features")
	}
	pl := &Landform{
		pos:      pos,
		features: features,
	}

	return pl
}

// Landform 路径地貌表示路径地图上的一个点
//   - 末尾特征将决定该路径地貌是否可行走
type Landform struct {
	width     int                // 所在路径覆盖宽度
	pos       int                // 位置
	features  []*LandformFeature // 地貌特征
	original  []*LandformFeature // 原始地貌
	totalCost float64            // 总消耗
}

// GetCoordinate 获取这个路径地貌指向的 x 和 y 坐标
//   - 建议通过 GetPos 来进行获取，这样可以避免一次转换
func (slf *Landform) GetCoordinate() (x, y int) {
	return g2d.PosToCoordinate(slf.width, slf.pos)
}

// GetPos 获取这个路径地貌指向的 pos 位置
func (slf *Landform) GetPos() int {
	return slf.pos
}

// GetTotalCost 获取这个路径地貌的总特征消耗
func (slf *Landform) GetTotalCost() float64 {
	return slf.totalCost
}

// Walkable 指示了该路径地貌是否可以在上面行走或者应该避免在上面行走
func (slf *Landform) Walkable() bool {
	return slf.features[len(slf.features)-1].Walkable()
}

// GetFeatures 获取这个路径地貌的特征
func (slf *Landform) GetFeatures() []*LandformFeature {
	return slf.features
}

// SetFeatures 设置这个路径地貌的特征
func (slf *Landform) SetFeatures(features ...*LandformFeature) {
	slf.features = features
	slf.original = nil
	slf.totalCost = 0
	for _, feature := range slf.features {
		slf.totalCost += feature.GetCost()
	}
}

// SetFeaturesWithRecoverable 通过可恢复的方式设置这个路径地貌的特征
//   - 使用该函数设置地貌特征后续也应该继续使用该函数进行地貌特征修改，如果中途使用 SetFeatures 修改地貌特征后将导致 Recover 不可用
func (slf *Landform) SetFeaturesWithRecoverable(features ...*LandformFeature) {
	if slf.original == nil {
		slf.original = slf.features
	}
	slf.features = features
	slf.totalCost = 0
	for _, feature := range slf.features {
		slf.totalCost += feature.GetCost()
	}
}

// Recover 恢复这个路径地貌的特征
func (slf *Landform) Recover() {
	if slf.original == nil {
		return
	}
	slf.features = slf.original
	slf.original = nil
	slf.totalCost = 0
	for _, feature := range slf.features {
		slf.totalCost += feature.GetCost()
	}
}
