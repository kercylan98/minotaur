package standard

import (
	"minotaur/game/feature"
	"time"
)

// NewRoomGame 对普通房间附加游戏功能的实例
//
// startCountDown 游戏开始倒计时，大于0生效
func NewRoomGame[P feature.Player](room *Room[P], startCountDown time.Duration) *RoomGame[P] {
	return &RoomGame[P]{
		Room:           room,
		startCountDown: startCountDown,
	}
}

type RoomGame[P feature.Player] struct {
	*Room[P]
	*Timer
	startTimeSecond int64         // 游戏开始时间戳（秒）
	startCountDown  time.Duration // 游戏开始倒计时
}

func (slf *RoomGame[P]) GameStart(startTimeSecond int64, loop func(game feature.RoomGame[P])) {
	slf.startTimeSecond = startTimeSecond

	if slf.startCountDown > 0 {
		slf.After("RoomGame.GameStart.StartCountDown", slf.startCountDown, func() {
			loop(slf)
		})
	} else {
		loop(slf)
	}
}

func (slf *RoomGame[P]) GameOver() {
	slf.Release()
}
