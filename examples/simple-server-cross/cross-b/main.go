package main

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/cross"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/timer"
	"go.uber.org/zap"
	"time"
)

func main() {
	srv := server.New(server.NetworkWebsocket, server.WithCross("nats", 2, cross.NewNats("127.0.0.1:4222")), server.WithTicker(10, false))
	srv.RegStartFinishEvent(func(srv *server.Server) {
		srv.Ticker().Loop("CROSS", timer.Instantly, time.Second, timer.Forever, func() {
			if err := srv.PushCrossMessage("nats", 1, []byte("I am cross 2")); err != nil {
				panic(err)
			}
		})
	})
	srv.RegReceiveCrossPacketEvent(func(srv *server.Server, senderServerId int64, packet []byte) {
		log.Info("Cross", zap.Int64("ServerID", senderServerId), zap.String("Packet", string(packet)))
	})
	if err := srv.Run(":19999"); err != nil {
		panic(err)
	}
}
