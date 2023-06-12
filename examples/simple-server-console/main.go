// 该案例中延时了控制台服务器的实现，支持运行中根据控制台指令执行额外的功能逻辑
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
