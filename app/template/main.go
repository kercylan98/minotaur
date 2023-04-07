package main

import (
	"fmt"
	"minotaur/app/template/app"
	"minotaur/game"
	"minotaur/game/conn"
	"minotaur/game/protobuf/protobuf"
)

func init() {
	app.State.RegMessagePlayer(int32(protobuf.MessageCode_SystemHeartbeat), func(player *game.Player) {
		fmt.Println("hhha")
	})
}

func main() {
	app.State.Run("/test", 8888, conn.NewOrdinary())
}
