package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"minotaur/game/protobuf/protobuf"
	"testing"
	"time"
)

func TestA(t *testing.T) {
run:
	{
	}
	for {
		ws := `ws://127.0.0.1:9000/test`
		c, _, err := websocket.DefaultDialer.Dial(ws, nil)
		if err != nil {
			continue
		}

		req := &protobuf.SystemHeartbeatClient{}
		d, _ := proto.Marshal(req)

		data, err := proto.Marshal(&protobuf.Message{
			Code: int32(protobuf.MessageCode_SystemHeartbeat),
			Data: d,
		})
		if err != nil {
			panic(err)
		}

		for {
			//fmt.Println(c, data)
			if err := c.WriteMessage(2, data); err != nil {
				fmt.Println(err)
				goto run
			}

			time.Sleep(3 * time.Second)
		}

	}
}
