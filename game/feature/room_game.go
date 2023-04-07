package feature

// RoomGame 房间游戏接口定义
type RoomGame[P Player] interface {
	Room[P]
	// GameStart 游戏开始
	GameStart(startTimeSecond int64, loop func(game RoomGame[P]))
	// GameOver 游戏结束
	GameOver()
}
