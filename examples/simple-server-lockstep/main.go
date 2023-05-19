package main

import (
	"github.com/kercylan98/minotaur/component/components"
	"github.com/kercylan98/minotaur/game/builtin"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/synchronization"
)

type Player struct {
	*builtin.Player[string]
}

type Command struct {
	CMD  int
	Data string
}

func main() {
	players := synchronization.NewMap[string, *Player]()

	srv := server.New(server.NetworkWebsocket, server.WithWebsocketWriteMessageType(server.WebsocketMessageTypeText))
	lockstep := components.NewLockstep[string, *Command]()

	srv.RegConnectionOpenedEvent(func(srv *server.Server, conn *server.Conn) {
		player := &Player{Player: builtin.NewPlayer[string](conn.GetID(), conn)}
		players.Set(conn.GetID(), player)
		lockstep.JoinClient(player)
	})
	srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn) {
		players.Delete(conn.GetID())
		lockstep.LeaveClient(conn.GetID())
		if players.Size() == 0 {
			lockstep.Stop()
		}
	})
	srv.RegConnectionReceiveWebsocketPacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte, messageType int) {
		switch string(packet) {
		case "start":
			lockstep.StartBroadcast()
		default:
			lockstep.AddCommand(&Command{CMD: 1, Data: string(packet)})
		}
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
