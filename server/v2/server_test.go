package server_test

import (
	"github.com/gin-gonic/gin"
	"github.com/kercylan98/minotaur/server/v2"
	"github.com/kercylan98/minotaur/server/v2/traffickers"
	"net/http"
	"testing"
)

func TestNewServer(t *testing.T) {
	r := gin.New()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"ping": "pong",
		})
	})
	srv := server.NewServer(traffickers.WebSocket(r, func(handler *gin.Engine, upgradeHandler func(writer http.ResponseWriter, request *http.Request) error) {
		handler.GET("/ws", func(context *gin.Context) {
			if err := upgradeHandler(context.Writer, context.Request); err != nil {
				context.AbortWithError(500, err)
			}
		})
	}))

	if err := srv.Run("tcp://:8080"); err != nil {
		panic(err)
	}
}
