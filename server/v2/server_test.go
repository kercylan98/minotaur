package server_test

import (
	"github.com/gin-gonic/gin"
	"github.com/kercylan98/minotaur/server/v2"
	"github.com/kercylan98/minotaur/server/v2/network"
	"net/http"
	"testing"
)

func TestNewServer(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"ping": "pong",
		})
	})
	srv := server.NewServer(network.WebSocketWithHandler(":9999", r, func(handler *gin.Engine, ws http.HandlerFunc) {
		handler.GET("/ws", func(context *gin.Context) {
			ws(context.Writer, context.Request)
		})
	}))
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
