package main

import (
	"fmt"
	"go.uber.org/zap"
	"minotaur/server"
	"minotaur/utils/log"
)

func main() {
	srv := server.New(server.NetworkTCP)
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		srv.GetConnections().RangeSkip(func(id string, c *server.Conn) bool {
			if id == conn.GetID() {
				return false
			}

			if err := c.Write([]byte(fmt.Sprintf("[%s]: %s", conn.GetID(), string(packet)))); err != nil {
				log.Debug("Message", zap.Error(err))
			}

			return true
		})
	})

	if err := srv.Run(":8888"); err != nil {
		panic(err)
	}
}
