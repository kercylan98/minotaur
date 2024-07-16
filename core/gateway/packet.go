package gateway

import (
	"google.golang.org/protobuf/proto"
	"time"
)

func wrapRequest(id MessageId, message proto.Message) (data []byte, err error) {
	var w = &RequestWrapper{
		Id:        id,
		Timestamp: time.Now().UnixMilli(),
	}
	if message != nil {
		w.Data, err = proto.Marshal(message)
		if err != nil {
			return
		}
	}
	data, err = proto.Marshal(w)
	return
}

func unwrapRequest(data []byte) (message *RequestWrapper, err error) {
	message = &RequestWrapper{}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return
	}
	message.ServerReceiveAt = time.Now().UnixMilli()
	return
}

func (r *RequestWrapper) read(dest proto.Message) error {
	return proto.Unmarshal(r.Data, dest)
}

func (r *RequestWrapper) wrapResponse(id MessageId, message proto.Message) (data []byte, err error) {
	var w = &ResponseWrapper{
		Id:              id,
		ClientSendAt:    r.Timestamp,
		ServerReceiveAt: r.ServerReceiveAt,
		ServerSendAt:    time.Now().UnixMilli(),
	}
	if message != nil {
		w.Data, err = proto.Marshal(message)
		if err != nil {
			return
		}
	}
	data, err = proto.Marshal(w)
	return
}
