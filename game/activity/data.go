package activity

func newData[PlayerID comparable, ActivityData any, PlayerData any]() *Data[PlayerID, ActivityData, PlayerData] {
	return &Data[PlayerID, ActivityData, PlayerData]{
		PlayerData:       make(map[PlayerID]PlayerData),
		PlayerLastNewDay: make(map[PlayerID]int64),
	}
}

// Data 活动数据信息
type Data[PlayerID comparable, ActivityData any, PlayerData any] struct {
	Data             ActivityData            // 活动全局数据
	PlayerData       map[PlayerID]PlayerData // 活动玩家数据
	LastNewDay       int64                   // 最后触发新的一天的时间戳
	PlayerLastNewDay map[PlayerID]int64      // 玩家最后触发新的一天的时间戳
}
