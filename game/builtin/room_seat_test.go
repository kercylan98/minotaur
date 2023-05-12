package builtin

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

var room = NewRoom[string, *Player[string]](1)

func TestNewRoomSeat(t *testing.T) {
	roomSeat := NewRoomSeat[string, *Player[string]](room)
	convey.Convey("TestNewRoomSeat.NewRoomSeat", t, func() {
		convey.So(roomSeat, convey.ShouldNotEqual, nil)
	})

	player1 := NewPlayer[string]("kiki", nil)
	player2 := NewPlayer[string]("john", nil)
	player3 := NewPlayer[string]("dave", nil)
	_ = roomSeat.Join(player1)
	_ = roomSeat.Join(player2)
	_ = roomSeat.Join(player3)
	player1Seat, _ := roomSeat.GetSeat(player1.GetID())
	player2Seat, _ := roomSeat.GetSeat(player2.GetID())
	player3Seat, _ := roomSeat.GetSeat(player3.GetID())
	convey.Convey("TestNewRoomSeat.GetSeat", t, func() {
		convey.So(player1Seat, convey.ShouldEqual, 0)
		convey.So(player2Seat, convey.ShouldEqual, 1)
		convey.So(player3Seat, convey.ShouldEqual, 2)
	})

	player1, _ = roomSeat.GetPlayerWithSeat(0)
	player2, _ = roomSeat.GetPlayerWithSeat(1)
	player3, _ = roomSeat.GetPlayerWithSeat(2)
	convey.Convey("TestNewRoomSeat.GetPlayerWithSeat", t, func() {
		convey.So(player1.GetID(), convey.ShouldEqual, "kiki")
		convey.So(player2.GetID(), convey.ShouldEqual, "john")
		convey.So(player3.GetID(), convey.ShouldEqual, "dave")
	})

	room.Leave(player2.GetID())
	player1Seat, _ = roomSeat.GetSeat(player1.GetID())
	player2Seat, err := roomSeat.GetSeat(player2.GetID())
	player3Seat, _ = roomSeat.GetSeat(player3.GetID())
	convey.Convey("TestNewRoomSeat.Leave", t, func() {
		convey.So(err, convey.ShouldNotEqual, nil)
		convey.So(player1Seat, convey.ShouldEqual, 0)
		convey.So(player3Seat, convey.ShouldEqual, 2)
	})
}
