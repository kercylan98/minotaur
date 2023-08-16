package gateway

type Packet struct {
	ConnID        string
	WebsocketType int
	Data          []byte
}
