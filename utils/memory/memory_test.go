package memory_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/memory"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

type Player struct {
	ID       int64
	Name     string
	Account  string
	Password string
}

var (
	QueryPlayer = memory.BindAction("QueryPlayer", func(playerId int64) *Player {
		return &Player{ID: playerId}
	})
	QueryPlayerPersist = memory.BindPersistCacheProgram("QueryPlayer", func(player *Player) {
		fmt.Println(player)
	}, memory.NewOption().WithPeriodicity(timer.GetTicker(10), timer.Instantly, time.Second*10, time.Second))
)

func TestBindAction(t *testing.T) {
	var player *Player
	player = QueryPlayer(1)
	fmt.Println(player.ID)
	player.ID = 666
	player = QueryPlayer(1)
	fmt.Println(player.ID)
	player = QueryPlayer(2)
	fmt.Println(player.ID)

	QueryPlayerPersist()

	time.Sleep(times.Week)
}
