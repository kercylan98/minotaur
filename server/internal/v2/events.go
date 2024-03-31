package server

import "github.com/kercylan98/minotaur/utils/collection/listings"

type (
	ConnectionReceivePacketEventHandler func(srv Server, conn Conn, packet Packet)
)

type Events interface {
	RegisterConnectionReceivePacketEvent(handler ConnectionReceivePacketEventHandler, priority ...int)
}

type events struct {
	*server

	connectionReceivePacketEventHandlers listings.PrioritySlice[ConnectionReceivePacketEventHandler]
}

func (s *events) init(srv *server) *events {
	s.server = srv
	return s
}

func (s *events) RegisterConnectionReceivePacketEvent(handler ConnectionReceivePacketEventHandler, priority ...int) {
	s.connectionReceivePacketEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionReceivePacket(conn Conn, packet Packet) {
	s.connectionReceivePacketEventHandlers.RangeValue(func(index int, value ConnectionReceivePacketEventHandler) bool {
		value(s.server, conn, packet)
		return true
	})
}
