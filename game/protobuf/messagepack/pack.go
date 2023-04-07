package messagepack

import (
	"google.golang.org/protobuf/proto"
	"minotaur/game/protobuf/protobuf"
)

// Pack 消息打包
func Pack(messageCode int32, message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	msg := &protobuf.Message{
		Code: messageCode,
		Data: data,
	}

	return proto.Marshal(msg)
}
