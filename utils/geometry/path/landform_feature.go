package path

var (
	LandformFeatureRoad  = NewLandformFeature(1, true)    // 道路
	LandformFeatureWater = NewLandformFeature(1.75, true) // 水
)

// NewLandformFeature 返回一个新的路径地貌特征， cost 将表示在该地貌上行走的消耗，walkable 则表示了该地貌是否支持行走
//   - 在基于 Terrain 使用时， LandformFeature 必须被声明为全局变量后再进行使用，例如 LandformFeatureRoad
func NewLandformFeature(cost float64, walkable bool) *LandformFeature {
	return &LandformFeature{cost: cost, walkable: walkable}
}

// LandformFeature 表示了路径地貌的特征
type LandformFeature struct {
	cost     float64 // 移动消耗
	walkable bool    // 适宜步行
}

// GetCost 获取在该路径地貌行走的消耗，这影响了该路径地貌是否是理想的选择
func (slf *LandformFeature) GetCost() float64 {
	return slf.cost
}

// Walkable 指示了该路径地貌是否可以在上面行走或者应该避免在上面行走
func (slf *LandformFeature) Walkable() bool {
	return slf.walkable
}
