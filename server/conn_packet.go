package server

type connPacket struct {
	websocketMessageType int
	packet               []byte
	callback             func(err error)
}
