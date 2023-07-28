package room

import "github.com/kercylan98/minotaur/game"

type Option[PID comparable, P game.Player[PID], R Room] func(info *Info[PID, P, R])

// WithPlayerLimit 设置房间人数上限
func WithPlayerLimit[PID comparable, P game.Player[PID], R Room](limit int) Option[PID, P, R] {
	return func(info *Info[PID, P, R]) {
		info.playerLimit = limit
	}
}

// WithNotAutoJoinSeat 设置不自动加入座位
func WithNotAutoJoinSeat[PID comparable, P game.Player[PID], R Room]() Option[PID, P, R] {
	return func(info *Info[PID, P, R]) {
		info.seat.autoSitDown = false
	}
}
