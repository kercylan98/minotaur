package stream

type Conn interface {
	Write(packet Packet) error

	Read() (Packet, error)

	Close() error
}
