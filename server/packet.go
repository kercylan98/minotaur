package server

type Packet struct {
	WebsocketType int
	Data          []byte
}

func (slf Packet) String() string {
	return string(slf.Data)
}
