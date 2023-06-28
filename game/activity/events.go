package activity

type (
	StartEventHandle[PlayerID comparable, Data, PlayerData any]        func(activity *Activity[PlayerID, Data, PlayerData])
	StopEventHandle[PlayerID comparable, Data, PlayerData any]         func(activity *Activity[PlayerID, Data, PlayerData])
	StartShowEventHandle[PlayerID comparable, Data, PlayerData any]    func(activity *Activity[PlayerID, Data, PlayerData])
	StopShowEventHandle[PlayerID comparable, Data, PlayerData any]     func(activity *Activity[PlayerID, Data, PlayerData])
	PlayerJoinEventHandle[PlayerID comparable, Data, PlayerData any]   func(activity *Activity[PlayerID, Data, PlayerData], playerId PlayerID)
	PlayerLeaveEventHandle[PlayerID comparable, Data, PlayerData any]  func(activity *Activity[PlayerID, Data, PlayerData], playerId PlayerID)
	NewDayEventHandle[PlayerID comparable, Data, PlayerData any]       func(activity *Activity[PlayerID, Data, PlayerData])
	PlayerNewDayEventHandle[PlayerID comparable, Data, PlayerData any] func(activity *Activity[PlayerID, Data, PlayerData], playerId PlayerID)
)
