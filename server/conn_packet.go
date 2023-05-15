package server

type connPacket struct {
	websocketMessageType int
	packet               []byte
}
