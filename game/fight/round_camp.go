package fight

// NewRoundCamp 创建一个新的回合制游戏阵营
func NewRoundCamp(campId, entity int, entities ...int) *RoundCamp {
	return &RoundCamp{
		campId:   campId,
		entities: append([]int{entity}, entities...),
	}
}

// RoundCamp 回合制游戏阵营
type RoundCamp struct {
	campId   int
	entities []int
}
