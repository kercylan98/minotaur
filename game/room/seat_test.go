package room_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/room"
	"github.com/kercylan98/minotaur/server"
	"testing"
)

type Player struct {
	ID string
}

func (slf *Player) GetID() string {
	return slf.ID
}

func (slf *Player) GetConn() *server.Conn {
	return nil
}

func (slf *Player) UseConn(conn *server.Conn) {

}

func (slf *Player) Close(err ...error) {

}

type Room struct {
}

func (slf *Room) GetGuid() int64 {
	return 0
}

func TestSeat_SetSeat(t *testing.T) {
	rm := room.NewManager[string, *Player, *Room]()

	r := &Room{}
	rm.CreateRoom(r)
	helper := rm.GetHelper(r)
	helper.AddSeat("a")
	helper.AddSeat("b")
	helper.RemoveSeat("a")
	helper.AddSeat("c")

	for i, s := range helper.GetSeatInfo() {
		if s == nil {
			fmt.Println(i, "nil")
		} else {
			fmt.Println(i, *s)
		}
	}
}
