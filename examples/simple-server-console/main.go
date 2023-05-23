package main

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
)

func main() {
	srv := server.New(server.NetworkWebsocket)
	srv.RegConsoleCommandEvent("test", func(srv *server.Server) {
		log.Info("Console", zap.String("Info", "Test"))
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
