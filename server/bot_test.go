package server_test

import (
	"github.com/kercylan98/minotaur/server"
	"io"
	"testing"
	"time"
)

type Writer struct {
	t   *testing.T
	bot *server.Bot
}

func (slf *Writer) Write(p []byte) (n int, err error) {
	slf.t.Log(string(p))
	switch string(p) {
	case "hello":
		slf.bot.SendPacket([]byte("world"))
	}
	return 0, nil
}

func TestNewBot(t *testing.T) {
	srv := server.New(server.NetworkWebsocket)

	srv.RegConnectionOpenedEvent(func(srv *server.Server, conn *server.Conn) {
		t.Logf("connection opened: %s", conn.GetID())
		conn.Close()
		conn.Write([]byte("hello"))
	})
	srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn, err any) {
		t.Logf("connection closed: %s", conn.GetID())
	})
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		t.Logf("connection %s receive packet: %s", conn.GetID(), string(packet))
		conn.Write([]byte("world"))
	})
	srv.RegStartFinishEvent(func(srv *server.Server) {
		bot := server.NewBot(srv, server.WithBotNetworkDelay(100, 20), server.WithBotWriter(func(bot *server.Bot) io.Writer {
			return &Writer{t: t, bot: bot}
		}))
		bot.JoinServer()
		time.Sleep(time.Second)
		bot.SendPacket([]byte("hello"))
	})

	srv.Run(":9600")
}
