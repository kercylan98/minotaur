package main

import (
	"github.com/kercylan98/minotaur/component/components"
	"github.com/kercylan98/minotaur/game/builtin"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"github.com/kercylan98/minotaur/utils/timer"
	"go.uber.org/zap"
	"time"
)

type Player struct {
	*builtin.Player[string]
}

type Command struct {
	CMD  int
	Data string
}

// 访问：http://www.websocket-test.com/
//   - 使用多个页面连接到服务器后，任一页面发送start即可开启帧同步
func main() {
	players := synchronization.NewMap[string, *Player]()

	srv := server.New(server.NetworkWebsocket,
		server.WithWebsocketWriteMessageType(server.WebsocketMessageTypeText),
		server.WithTicker(20, false),
		server.WithMonitor(),
	)
	lockstep := components.NewLockstep[string, *Command]()
	srv.RegStartFinishEvent(func(srv *server.Server) {
		srv.Ticker().Loop("monitor", timer.Instantly, time.Second/2, timer.Forever, func() {
			m := srv.GetMonitor()
			log.Info("Monitor.Message",
				zap.Any("Total", m.MessageTotal()),
				zap.Any("Second", m.MessageSecond()),
				zap.Any("Cost", m.MessageCost()),
				zap.Any("DoneAvg", m.MessageDoneAvg()),
				zap.Any("QPS", m.MessageQPS()),
				zap.Any("TopQPS", m.MessageTopQPS()),
			)
			log.Info("Monitor.Cross",
				zap.Any("Total", m.CrossMessageTotal()),
				zap.Any("Second", m.CrossMessageSecond()),
				zap.Any("Cost", m.CrossMessageCost()),
				zap.Any("DoneAvg", m.CrossMessageDoneAvg()),
				zap.Any("QPS", m.MessageQPS()),
				zap.Any("TopQPS", m.CrossMessageTopQPS()),
			)
			log.Info("Monitor.Packet",
				zap.Any("Total", m.PacketMessageTotal()),
				zap.Any("Second", m.PacketMessageSecond()),
				zap.Any("Cost", m.PacketMessageCost()),
				zap.Any("DoneAvg", m.PacketMessageDoneAvg()),
				zap.Any("QPS", m.PacketMessageQPS()),
				zap.Any("TopQPS", m.PacketMessageTopQPS()),
			)
			log.Info("Monitor.Ticker",
				zap.Any("Total", m.TickerMessageTotal()),
				zap.Any("Second", m.TickerMessageSecond()),
				zap.Any("Cost", m.TickerMessageCost()),
				zap.Any("DoneAvg", m.TickerMessageDoneAvg()),
				zap.Any("QPS", m.TickerMessageQPS()),
				zap.Any("TopQPS", m.TickerMessageTopQPS()),
			)
			log.Info("Monitor.Error",
				zap.Any("Total", m.ErrorMessageTotal()),
				zap.Any("Second", m.ErrorMessageSecond()),
				zap.Any("Cost", m.ErrorMessageCost()),
				zap.Any("DoneAvg", m.ErrorMessageDoneAvg()),
				zap.Any("QPS", m.ErrorMessageQPS()),
				zap.Any("TopQPS", m.ErrorMessageTopQPS()),
			)
		})
	})
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
			time.Sleep(random.Duration(1, 3) * time.Second)
			lockstep.AddCommand(&Command{CMD: 1, Data: string(packet)})
		}
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
