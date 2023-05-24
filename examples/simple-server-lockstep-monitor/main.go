package main

import (
	"github.com/kercylan98/minotaur/component/components"
	"github.com/kercylan98/minotaur/game/builtin"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/log"
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
		srv.Ticker().Loop("monitor", timer.Instantly, time.Second, timer.Forever, func() {
			m := srv.GetMonitor()
			log.Info("Monitor.Message",
				zap.Any("MessageTotal", m.MessageTotal()),
				zap.Any("MessageSecond", m.MessageSecond()),
				zap.Any("MessageCost", m.MessageCost()),
				zap.Any("MessageDoneAvg", m.MessageDoneAvg()),
				zap.Any("MessageQPS", m.MessageQPS()),
				zap.Any("MessageTopQPS", m.MessageTopQPS()),
			)
			log.Info("Monitor.Cross",
				zap.Any("CrossMessageTotal", m.CrossMessageTotal()),
				zap.Any("CrossMessageSecond", m.CrossMessageSecond()),
				zap.Any("CrossMessageCost", m.CrossMessageCost()),
				zap.Any("CrossMessageDoneAvg", m.CrossMessageDoneAvg()),
				zap.Any("CrossMessageQPS", m.MessageQPS()),
				zap.Any("CrossMessageTopQPS", m.CrossMessageTopQPS()),
			)
			log.Info("Monitor.Packet",
				zap.Any("PacketMessageTotal", m.PacketMessageTotal()),
				zap.Any("PacketMessageSecond", m.PacketMessageSecond()),
				zap.Any("PacketMessageCost", m.PacketMessageCost()),
				zap.Any("PacketMessageDoneAvg", m.PacketMessageDoneAvg()),
				zap.Any("PacketMessageQPS", m.PacketMessageQPS()),
				zap.Any("PacketMessageTopQPS", m.PacketMessageTopQPS()),
			)
			log.Info("Monitor.Ticker",
				zap.Any("TickerMessageTotal", m.TickerMessageTotal()),
				zap.Any("TickerMessageSecond", m.TickerMessageSecond()),
				zap.Any("TickerMessageCost", m.TickerMessageCost()),
				zap.Any("TickerMessageDoneAvg", m.TickerMessageDoneAvg()),
				zap.Any("TickerMessageQPS", m.TickerMessageQPS()),
				zap.Any("TickerMessageTopQPS", m.TickerMessageTopQPS()),
			)
			log.Info("Monitor.Error",
				zap.Any("ErrorMessageTotal", m.ErrorMessageTotal()),
				zap.Any("ErrorMessageSecond", m.ErrorMessageSecond()),
				zap.Any("ErrorMessageCost", m.ErrorMessageCost()),
				zap.Any("ErrorMessageDoneAvg", m.ErrorMessageDoneAvg()),
				zap.Any("ErrorMessageQPS", m.ErrorMessageQPS()),
				zap.Any("ErrorMessageTopQPS", m.ErrorMessageTopQPS()),
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
			lockstep.AddCommand(&Command{CMD: 1, Data: string(packet)})
		}
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
