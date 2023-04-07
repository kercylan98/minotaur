package feature

// RoomTeam 房间小队接口定义
type RoomTeam interface {
	// GetTeamMaximum 获取最大小队数量
	GetTeamMaximum() int
	// GetTeam 获取特定队伍
	GetTeam(guid int64) Team
	// GetTeams 获取所有队伍
	GetTeams() map[int64]Team
}

type Team interface {
	// GetGuid 获取小队 guid
	GetGuid() int64
	// GetPlayer 获取小队特定玩家
	GetPlayer(guid int64) Player
	// GetPlayers 获取小队玩家
	GetPlayers() map[int64]Player
}
