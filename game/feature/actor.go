package feature

// Actor 玩家演员接口定义
type Actor interface {
	// GetGuid 获取演员 guid
	GetGuid() int64
	// GetPlayer 获取所属玩家
	GetPlayer() Player
}
