package room_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/room"
	"testing"
)

func TestNewManager(t *testing.T) {
	m := room.NewManager[string, *Player, *Room]()
	r := &Room{}
	m.CreateRoom(r)
	helper := m.GetHelper(r)

	helper.Join(&Player{ID: "player_01"})
	helper.Join(&Player{ID: "player_02"})
	helper.Join(&Player{ID: "player_03"})
	helper.Leave(helper.GetPlayer("player_02"))
	helper.Join(&Player{ID: "player_02"})

	helper.BroadcastExcept(func(player *Player) {
		fmt.Println(player.GetID())
	}, func(player *Player) bool {
		return false
	})
}
