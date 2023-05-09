package server

import "errors"

var (
	ErrConstructed                   = errors.New("the Server must be constructed using the server.New function")
	ErrCanNotSupportNetwork          = errors.New("can not support network")
	ErrMessageTypePacketAttrs        = errors.New("MessageTypePacket must contain *Conn and []byte")
	ErrMessageTypeErrorAttrs         = errors.New("MessageTypePacket must contain error and MessageErrorAction")
	ErrNetworkOnlySupportHttp        = errors.New("the current network mode is not compatible with HttpRouter, only NetworkHttp is supported")
	ErrNetworkIncompatibleHttp       = errors.New("the current network mode is not compatible with NetworkHttp")
	ErrWebsocketMessageTypeException = errors.New("unknown message type, will not work")
	ErrNotWebsocketUseMessageType    = errors.New("message type filtering only supports websocket and does not take effect")
	ErrWebsocketIllegalMessageType   = errors.New("illegal message type")
)
