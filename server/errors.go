package server

import "errors"

var (
	ErrConstructed                     = errors.New("the Server must be constructed using the server.New function")
	ErrCanNotSupportNetwork            = errors.New("can not support network")
	ErrMessageTypePacketAttrs          = errors.New("MessageTypePacket must contain *Conn and []byte")
	ErrWebsocketMessageTypePacketAttrs = errors.New("MessageTypePacket must contain *Conn and []byte and int(MessageType)")
	ErrMessageTypeErrorAttrs           = errors.New("MessageTypePacket must contain error and MessageErrorAction")
	ErrMessageTypeCrossErrorAttrs      = errors.New("MessageTypeCross must contain int64(server id) and []byte")
	ErrMessageTypeTickerErrorAttrs     = errors.New("MessageTypeTicker must contain func()")
	ErrNetworkOnlySupportHttp          = errors.New("the current network mode is not compatible with HttpRouter, only NetworkHttp is supported")
	ErrNetworkOnlySupportGRPC          = errors.New("the current network mode is not compatible with RegGrpcServer, only NetworkGRPC is supported")
	ErrNetworkIncompatibleHttp         = errors.New("the current network mode is not compatible with NetworkHttp")
	ErrWebsocketMessageTypeException   = errors.New("unknown message type, will not work")
	ErrNotWebsocketUseMessageType      = errors.New("message type filtering only supports websocket and does not take effect")
	ErrWebsocketIllegalMessageType     = errors.New("illegal message type")
	ErrPleaseUseWebsocketHandle        = errors.New("in Websocket mode, please use the RegConnectionReceiveWebsocketPacketEvent function to register")
	ErrPleaseUseOrdinaryPacketHandle   = errors.New("non Websocket mode, please use the RegConnectionReceivePacketEvent function to register")
	ErrNoSupportCross                  = errors.New("the server does not support GetID or PushCrossMessage, please use the WithCross option to create the server")
	ErrNoSupportMonitor                = errors.New("the server does not support GetMonitor, please use the WithMonitor option to create the server")
	ErrNoSupportTicker                 = errors.New("the server does not support Ticker, please use the WithTicker option to create the server")
	ErrUnregisteredCrossName           = errors.New("unregistered cross name, please use the WithCross option to create the server")
)
