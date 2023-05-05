package game

import (
	"minotaur/game"
	"minotaur/server"
	"minotaur/utils/sole"
)

func init() {
	Server.RegConnectionOpenedEvent(onConnectionOpened)
	Game.RegPlayerLeaveWorldEvent(onConnectionClosed)
	Server.RegConnectionReceivePacketEvent(onConnectionReceivePacket)
}

func onConnectionReceivePacket(srv *server.Server, conn *server.Conn, packet []byte) {
	player := Game.World.GetPlayerWithConnID(conn.GetID())
	if player == nil {
		return
	}

	player.RegGameplayStartEvent(player.onGameplayStart)
	player.RegGameplayOverEvent(player.onGameplayOver)

	switch string(packet) {
	case "start":
		player.Start()
	}
}

func onConnectionClosed(world game.World[int64, *Player], player *Player) {
	Game.Leave(player.GetID())
	player.Close()
}

func onConnectionOpened(srv *server.Server, conn *server.Conn) {
	player := NewPlayer(sole.GetSync(), conn)
	if err := Game.World.Join(player); err != nil {
		panic(err)
	}
}
