package game

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"minotaur/game/protobuf/protobuf"
	"minotaur/utils/gin/middlewares"
	"minotaur/utils/log"
	"minotaur/utils/timer"
	"net/http"
	"strings"
	"time"
)

const (
	loginTimeoutTimerName = "player_login_timeout_timer"
)

type OnCreateConnHandleFunc func() Conn
type gateway struct {
	server *http.Server
}

func (slf *gateway) run(stateMachine *StateMachine, appName string, port int, onCreateConnHandleFunc OnCreateConnHandleFunc) *gateway {

	// Gin WebSocket
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middlewares.Logger(log.Logger), middlewares.Cors())
	pprof.Register(router) // pprof 可视化分析

	var upgrade = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	router.GET(fmt.Sprintf("/%s", appName), func(context *gin.Context) {
		ip := context.GetHeader("X-Real-IP")
		ws, err := upgrade.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Error("Gateway", zap.Error(err))
			return
		}
		if len(ip) == 0 {
			addr := ws.RemoteAddr().String()
			if index := strings.LastIndex(addr, ":"); index != -1 {
				ip = addr[0:index]
			}
		}
		conn := onCreateConnHandleFunc()
		if err := context.ShouldBind(conn); err != nil {
			log.Error("Gateway", zap.Error(err))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		playerGuid, err := stateMachine.sonyflake.NextID()
		if err != nil {
			log.Error("Gateway", zap.Error(err))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		player := &Player{
			guid:      int64(playerGuid),
			conn:      conn,
			ws:        ws,
			ip:        ip,
			Manager:   timer.GetManager(64),
			GameTimer: stateMachine.Manager,
		}

		channelId, size := stateMachine.channelStrategy()
		channel := stateMachine.channel(channelId, size)
		player, err = channel.join(player)
		if err != nil {
			player.exit(err.Error())
			return
		}

		player.channel.push(new(Message).Init(MessageTypeEvent, EventTypeGuestJoin, player))
		if stateMachine.loginTimeout > 0 {
			player.After(loginTimeoutTimerName, stateMachine.loginTimeout, func() {
				player.exit("login timeout")
			})
		}

		defer func() {
			if err := recover(); err != nil {
				player.exit()
			}
		}()

		for {
			if err := player.ws.SetReadDeadline(time.Now().Add(time.Second * 30)); err != nil {
				panic(err)
			}

			_, data, err := player.ws.ReadMessage()
			if err != nil {
				panic(err)
			}

			var msg = new(protobuf.Message)
			if err = proto.Unmarshal(data, msg); err != nil {
				continue
			}

			player.channel.push(new(Message).Init(MessageTypePlayer, player, msg.Code, msg.Data))

		}

	})

	// HttpServer
	slf.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	log.Info("GatewayStart", zap.String("listen", slf.server.Addr))

	go func() {
		if err := slf.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			stateMachine.errChannel <- err
			return
		}
	}()

	return slf
}

func (slf *gateway) shutdown(context context.Context) error {
	log.Info("GatewayShutdown", zap.String("stateMachine", "start"))
	if err := slf.server.Shutdown(context); err != nil {
		return err
	}
	log.Info("GatewayShutdown", zap.String("stateMachine", "normal"))
	return nil
}
