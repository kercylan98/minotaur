package space_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/space"
)

type Room struct {
	Id int64
}

func (r *Room) GetId() int64 {
	return r.Id
}

type Player struct {
	Id string
}

func (p *Player) GetId() string {
	return p.Id
}

func ExampleNewRoomManager() {
	var rm = space.NewRoomManager[string, int64, *Player, *Room]()
	fmt.Println(rm == nil)

	// Output:
	// false
}

func ExampleRoomManager_AssumeControl() {
	var rm = space.NewRoomManager[string, int64, *Player, *Room]()
	var room = &Room{Id: 1}
	var controller = rm.AssumeControl(room)

	if err := controller.AddEntity(&Player{Id: "1"}); err != nil {
		// 房间密码不匹配或者房间已满
		panic(err)
	}

	fmt.Println(controller.GetEntityCount())

	// Output:
	// 1
}
