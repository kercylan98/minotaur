package extension

import (
	"github.com/kercylan98/minotaur/game"
	"time"
)

func NewPlayerLoginLauncher[PlayerID comparable](player game.Player[PlayerID]) *PlayerLoginLauncher[PlayerID] {
	return &PlayerLoginLauncher[PlayerID]{
		Player: player,
	}
}

type PlayerLoginLauncher[PlayerID comparable] struct {
	game.Player[PlayerID]
	loggedTime time.Time // 登录时间
	logoutTime time.Time // 登出时间
}

func (slf *PlayerLoginLauncher[PlayerID]) HasLogged() bool {
	return !slf.loggedTime.IsZero()
}

func (slf *PlayerLoginLauncher[PlayerID]) Logged() {
	slf.loggedTime = time.Now()
}

func (slf *PlayerLoginLauncher[PlayerID]) Logout() {
	slf.logoutTime = time.Now()
	slf.loggedTime = time.Time{}
}

func (slf *PlayerLoginLauncher[PlayerID]) GetLoggedTime() time.Time {
	return slf.loggedTime
}

func (slf *PlayerLoginLauncher[PlayerID]) GetLogoutTime() time.Time {
	return slf.logoutTime
}
