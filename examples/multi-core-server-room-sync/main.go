package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/builtin"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"strconv"
	"strings"
)

type message struct {
	conn   *server.Conn
	packet []byte
}

type Room struct {
	*builtin.Room[string, *builtin.Player[string]]
	channel chan *message // need close
	pool    *synchronization.Pool[*message]
}

func newRoom(guid int64) *Room {
	room := &Room{
		Room: builtin.NewRoom[string, *builtin.Player[string]](guid),
	}
	room.pool = synchronization.NewPool[*message](1024*100, func() *message {
		return new(message)
	}, func(data *message) {
		data.conn = nil
		data.packet = nil
	})
	room.channel = make(chan *message, 1024*100)
	go func() {
		for msg := range room.channel {
			room.handePacket(msg.conn, msg.packet)
			room.pool.Release(msg)
		}
	}()
	return room
}

func (slf *Room) PushMessage(conn *server.Conn, packet []byte) {
	msg := slf.pool.Get()
	msg.conn = conn
	msg.packet = packet
	slf.channel <- msg
}

func (slf *Room) handePacket(conn *server.Conn, packet []byte) {
	conn.WriteString(fmt.Sprintf("[%d] %s", slf.GetGuid(), string(packet)))
}

// 以房间为核心玩法的多核服务器实现
//   - 服务器消息处理为异步执行
//   - 由房间分发具体消息，在房间内所有消息为同步执行
func main() {
	rooms := synchronization.NewMap[int64, *Room]()

	srv := server.New(server.NetworkWebsocket,
		server.WithWebsocketWriteMessageType(server.WebsocketMessageTypeText),
		server.WithMultiCore(10),
	)

	srv.RegConnectionReceiveWebsocketPacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte, messageType int) {
		p := strings.SplitN(string(packet), ":", 2)
		roomId, err := strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			conn.WriteString(fmt.Sprintf("wrong room id, err: %s", err.Error()))
			return
		}
		// 假定命令格式 ${房间ID}:命令
		switch p[1] {
		case "create":
			if !rooms.Exist(roomId) {
				rooms.Set(roomId, newRoom(roomId))
				conn.WriteString(fmt.Sprintf("create room[%d] success", roomId))
			} else {
				conn.WriteString(fmt.Sprintf("room[%d] existed", roomId))
			}
		default:
			room, exist := rooms.GetExist(roomId)
			if !exist {
				rooms.Set(roomId, room)
				conn.WriteString(fmt.Sprintf("room[%d] does not exist, create room please use ${roomId}:create", roomId))
			} else {
				room.PushMessage(conn, []byte(p[1]))
			}
		}
	})

	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
