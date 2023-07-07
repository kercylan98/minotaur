package server

import "errors"

var (
	ErrConstructed                 = errors.New("the Server must be constructed using the server.New function")
	ErrCanNotSupportNetwork        = errors.New("can not support network")
	ErrNetworkOnlySupportHttp      = errors.New("the current network mode is not compatible with HttpRouter, only NetworkHttp is supported")
	ErrNetworkOnlySupportGRPC      = errors.New("the current network mode is not compatible with RegGrpcServer, only NetworkGRPC is supported")
	ErrNetworkIncompatibleHttp     = errors.New("the current network mode is not compatible with NetworkHttp")
	ErrWebsocketIllegalMessageType = errors.New("illegal message type")
	ErrNoSupportCross              = errors.New("the server does not support GetID or PushCrossMessage, please use the WithCross option to create the server")
	ErrNoSupportTicker             = errors.New("the server does not support Ticker, please use the WithTicker option to create the server")
)
