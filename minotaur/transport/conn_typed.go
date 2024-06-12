package transport

type ConnTyped interface {
	Write(packet Packet)

	Close()
}
