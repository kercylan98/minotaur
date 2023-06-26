package activity

type StartEventHandle[PlayerID comparable, Data any, PlayerData any] func(activity *Activity[PlayerID, Data, PlayerData])
type FinishEventHandle[PlayerID comparable, Data any, PlayerData any] func(activity *Activity[PlayerID, Data, PlayerData])
type NewDayEventHandle[PlayerID comparable, Data any, PlayerData any] func(activity *Activity[PlayerID, Data, PlayerData])
type PlayerNewDayEventHandle[PlayerID comparable, Data any, PlayerData any] func(activity *Activity[PlayerID, Data, PlayerData], playerId PlayerID)
